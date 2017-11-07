//==================================================================================================\\
// 		   AUTHOR: 	Brede Fritjof Klausen		  				  								    \\
// 		  SUBJECT: 	IMT2681 Cloud Technologies													    \\
//==================================================================================================\\
//	SOURCES:												 									    \\
// * https://stackoverflow.com/questions/38127583/get-last-inserted-element-from-mongodb-in-golang  \\
// * https://elithrar.github.io/article/testing-http-handlers-go/								    \\
//==================================================================================================\\

package gofiles

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

// HandlePOST handles post
func HandlePOST(w http.ResponseWriter, r *http.Request) {

	payload := Payload{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)

	db := SetupDB()
	db.Init()
	db.Add(payload)

}

// HandleGET handles get
func HandleGET(w http.ResponseWriter, r *http.Request, getID string) {

	db := SetupDB()
	payload, ok := db.Get(getID)
	if ok == false {
		http.Error(w, "ObjectID not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.Marshal(&payload)
	json.NewEncoder(w).Encode(payload)
}

// HandleDELETE handles delete
func HandleDELETE(w http.ResponseWriter, r *http.Request, getID string) {

	db := SetupDB()
	ok := db.Delete(getID)
	if ok == false {
		http.Error(w, "ObjectID not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandleWebhook handles post, get or delete
func HandleWebhook(w http.ResponseWriter, r *http.Request) {

	URL := strings.Split(r.URL.Path, "/")
	length := len(URL[1])
	objectID := URL[1]

	switch r.Method {

	case "POST":
		if length == 0 {
			HandlePOST(w, r)
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}

	case "GET":
		if length == 24 {
			HandleGET(w, r, objectID)
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}

	case "DELETE":
		if length == 24 {
			HandleDELETE(w, r, objectID)
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}

	default:
		http.Error(w, "Method has to be GET, POST or DELETE", http.StatusBadRequest)
	}
}

// HandleLatest handles latest
func HandleLatest(w http.ResponseWriter, r *http.Request) {

	payload := Payload{}

	switch r.Method {
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if payload.BaseCurrency != "EUR" {
			http.Error(w, payload.BaseCurrency+" is not implemented to be baseCurrency", http.StatusNotImplemented)
			return
		}
	case "GET":
		// Setting a default value if the user opens the program in a webbrowser
		payload.BaseCurrency = "EUR"
		payload.TargetCurrency = "NOK"

	default:
		http.Error(w, "Method has to be POST (or GET)", http.StatusBadRequest)
		return
	}

	tempToday := time.Now().Local()
	today := tempToday.Format("2006-01-02")

	db := SetupDB()
	currency, ok := db.GetLatest(today)

	if ok == false {
		tempToday = time.Now().Local().AddDate(0, 0, -1)
		yesterday := tempToday.Format("2006-01-02")
		currency, ok = db.GetLatest(yesterday)
		if ok == false {
			http.Error(w, "Could not get any data from today or yesterday :/", 404)
		}
	}

	rate := currency.Rates[payload.TargetCurrency]
	fmt.Fprint(w, rate)
}

// HandleAverage handles average of 3 days
func HandleAverage(w http.ResponseWriter, r *http.Request) {

	const totalDays = 3
	tdFloat := float64(totalDays)
	payload := Payload{}

	switch r.Method {
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if payload.BaseCurrency != "EUR" {
			http.Error(w, payload.BaseCurrency+" is not implemented to be baseCurrency", http.StatusBadRequest)
			return
		}
	case "GET":
		// Setting a default value if the user opens the program in a webbrowser
		payload.BaseCurrency = "EUR"
		payload.TargetCurrency = "NOK"

	default:
		http.Error(w, "Method has to be POST (or GET)", http.StatusBadRequest)
		return
	}

	tempToday := time.Now().Local()
	today := tempToday.Format("2006-01-02")

	db := SetupDB()
	sum := 0.0000

	for i := 0; i < totalDays; i++ {
		tempToday = time.Now().Local().AddDate(0, 0, -i)
		today = tempToday.Format("2006-01-02")

		currency, ok := db.GetLatest(today)
		if ok == false {
			http.Error(w, "There isn't any data for all the days yet", 404)
			return
		}

		sum += currency.Rates[payload.TargetCurrency]
	}
	average := sum / tdFloat
	fmt.Fprint(w, average)
}

// HandleTestTrigger overrides triggercheck and sends to discord
func HandleTestTrigger(w http.ResponseWriter, r *http.Request) {
	db := SetupDB()
	count := db.Count()

	if count > 0 {
		webhook := Payload{}
		session, err := mgo.Dial(db.DatabaseURL)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		tempToday := time.Now().Local()
		today := tempToday.Format("2006-01-02")

		currency, ok := db.GetLatest(today)
		if ok == false {
			http.Error(w, "There isn't any currency data from today", 404)
			return
		}

		for i := 1; i <= count; i++ {
			err = session.DB(db.DatabaseName).C(db.ColWebHook).Find(nil).Skip(count - i).One(&webhook)
			if err != nil {
				log.Println("Error in HandleTestTrigger() | Can not get one or more webhook data", err.Error())
				return
			}

			rate := currency.Rates[webhook.TargetCurrency]
			rateString := fmt.Sprint(rate)
			min := fmt.Sprint(webhook.MinTriggerValue)
			max := fmt.Sprint(webhook.MaxTriggerValue)
			text := "baseCurrency: " + webhook.BaseCurrency + "\ntargetCurrency: " + webhook.TargetCurrency + "\ncurrent: " + rateString + "\nminTriggerValue: " + min + "\nmaxTriggerValue: " + max

			DiscordOperator(text, webhook.WebhookURL)
		}
	} else {
		http.Error(w, "There isn't any data yet", 404)
		return
	}
}
