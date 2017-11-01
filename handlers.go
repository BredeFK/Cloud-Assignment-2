package main

import (
	"net/http"
	"encoding/json"
	"strings"
	"time"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

func HandlePOST(w http.ResponseWriter, r *http.Request){

	payload := Payload{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	db := SetupDB()
	db.Init()
	db.Add(payload)

}

func HandleGET(w http.ResponseWriter, r *http.Request, getID string) {

	db := SetupDB()
	payload, ok := db.Get(getID)
	if ok == false{
		http.Error(w, "ObjectID not found", http.StatusBadRequest)
		return
	}

	json.Marshal(&payload)
	json.NewEncoder(w).Encode(payload)
}

func HandleDELETE( w http.ResponseWriter, r *http.Request, getID string) {

	db := SetupDB()
	ok := db.Delete(getID)
	if ok == false{
		http.Error(w, "ObjectID not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleWebhook (w http.ResponseWriter, r *http.Request) {

	URL := strings.Split(r.URL.Path, "/")
	ObjectID := URL[1]

	switch r.Method {

	case "POST":
		HandlePOST(w, r)

	case "GET":
		HandleGET(w, r, ObjectID)

	case "DELETE":
		HandleDELETE(w, r, ObjectID)
	}
}

func HandleLatest (w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "Method has to be POST (or GET)", http.StatusMethodNotAllowed)
		return
	}

	tempToday := time.Now().Local()
	today := tempToday.Format("2006-01-02")

	db := SetupDB()
	currency, ok := db.GetLatest(today)

	if ok == false{
		http.Error(w, "There isn't any data from today yet", 404)
		return
	}
	rate := currency.Rates[payload.TargetCurrency]
	fmt.Fprint(w, rate)
	/*
	rateString := fmt.Sprint(rate)
	text := "The rate between " + payload.BaseCurrency + " and " + payload.TargetCurrency + " is: " + rateString
	DiscordOperator( text , DiscordURL_notAbot)
	*/
}

func HandleAverage (w http.ResponseWriter, r *http.Request) {

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
			http.Error(w, payload.BaseCurrency+" is not implemented to be baseCurrency", http.StatusNotImplemented)
			return
		}
	case "GET":
		// Setting a default value if the user opens the program in a webbrowser
		payload.BaseCurrency = "EUR"
		payload.TargetCurrency = "NOK"

	default:
		http.Error(w, "Method has to be POST (or GET)", http.StatusMethodNotAllowed)
		return
	}

	tempToday := time.Now().Local()
	today := tempToday.Format("2006-01-02")

	db := SetupDB()
	sum := 0.0000

	for i := 0; i < totalDays; i++{
		tempToday = time.Now().Local().AddDate(0,0,-i)
		today = tempToday.Format("2006-01-02")

		currency, ok := db.GetLatest(today)
		if ok == false{
			http.Error(w, "There isn't any data for all the days yet", 404)
			return
		}

		sum += currency.Rates[payload.TargetCurrency]
	}
	average := sum / tdFloat
	fmt.Fprint(w, average)

	/*
	days := strconv.Itoa(totalDays)
	averageString := fmt.Sprint(average)
	text := "The average rate over the " + days + " days between " + payload.BaseCurrency + " and " + payload.TargetCurrency + " is: " + averageString
	DiscordOperator( text , DiscordURL_notAbot)
	*/
}

func HandleTestTrigger (w http.ResponseWriter, r *http.Request) {
	db := SetupDB()
	count := db.Count()

	if count > 0{
		webhook := Payload{}
		session, err := mgo.Dial(db.DatabaseURL)
		if err != nil{
			panic(err)
		}
		defer session.Close()

		tempToday := time.Now().Local()
		today := tempToday.Format("2006-01-02")

		currency, ok := db.GetLatest(today)
		if ok == false{
			http.Error(w, "There isn't any currency data from today", 404)
			return
		}

		for i := 1; i<= count; i++{
			err = session.DB(db.DatabaseName).C(db.ColWebHook).Find(nil).Skip(count-i).One(&webhook)
			if err != nil{
				log.Printf("Error in HandleTestTrigger() | Can not get one or more webhook data", err)
				return
			}

			rate := currency.Rates[webhook.TargetCurrency]
			rateString := fmt.Sprint(rate)
			min := fmt.Sprint(webhook.MinTriggerValue)
			max := fmt.Sprint(webhook.MaxTriggerValue)
			text := "baseCurrency: " + webhook.BaseCurrency + "\ntargetCurrency: " + webhook.TargetCurrency + "\ncurrent: " + rateString + "\nminTriggerValue: " + min + "\nmaxTriggerValue: " + max

			DiscordOperator(text, webhook.WebhookURL)


		}


	}else{
		http.Error(w, "There isn't any data yet", 404)
		return
	}
}

func HandleAdd (w http.ResponseWriter, r *http.Request) {	// TODO : make automatic
	DailyCurrencyAdder()
}