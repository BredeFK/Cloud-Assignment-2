package main

import (
	"net/http"
	"encoding/json"
	"strings"
)

func HandlePOST(w http.ResponseWriter, r *http.Request){
	/*
	client := http.Client{
		Timeout: time.Second * 2,
	}
	*/

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

}

func HandleAverage (w http.ResponseWriter, t *http.Request) {

}