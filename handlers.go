package main

import (
	"net/http"
	"encoding/json"
)

func HandlePOST(w http.ResponseWriter, r *http.Request){
	/*
	client := http.Client{
		Timeout: time.Second * 2,
	}
	*/

	Payload := Payload{}

	err := json.NewDecoder(r.Body).Decode(&Payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

}

func HandleGET(w http.ResponseWriter, r *http.Request) {

}

func HandleDELETE( w http.ResponseWriter, r *http.Request) {

}

func HandleWebhook (w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		HandlePOST(w, r)

	case "GET":
		HandleGET(w, r)

	case "DELETE":
		HandleDELETE(w, r)

	}
}