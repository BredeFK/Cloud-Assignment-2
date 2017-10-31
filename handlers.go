package main

import (
	"net/http"
	"encoding/json"
	"strings"
	"time"
	"fmt"
)

func HandlePOST(w http.ResponseWriter, r *http.Request){

	payload := Payload{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// Remove before turning in
	json.Marshal(&payload)
	json.NewEncoder(w).Encode(payload)

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
		payload.BaseCurrency = "EUR"
		payload.TargetCurrency = "NOK"

	default:
		http.Error(w, "Method has to be POST", http.StatusMethodNotAllowed)
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
	/*
	rateString := fmt.Sprint(rate)
	text := "The rate between " + payload.BaseCurrency + " and " + payload.TargetCurrency + " is: " + rateString
	DiscordOperator( text , DiscordURL_notAbot)
	*/
	fmt.Fprint(w, rate)
}

func HandleAverage (w http.ResponseWriter, r *http.Request) {

}

func HandleAdd (w http.ResponseWriter, r *http.Request) {	// TODO : make automatic
	DailyCurrencyAdder()
}