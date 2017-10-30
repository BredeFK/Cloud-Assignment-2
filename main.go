/////////////////////////////////////////////////////////////////////
// 		   AUTHOR: 	Brede Fritjof Klausen		  				  //
//	STUDENTNUMBER: 	473211						 				 //
// 		  SUBJECT: 	IMT2681 Cloud Technologies					//
//=============================================================//
//	SOURCES:												  //
// * 														 //
// *														//
/////////////////////////////////////////////////////////////
package main

import (
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type Payload struct {
	ID 				bson.ObjectId `bson:"_id,omitempty"`
	WebhookURL  	string 		  `json:"webhookURL"`
	BaseCurrency 	string		  `json:"baseCurrency"`
	TargetCurrency	string		  `json:"targetCurrency"`
//	CurrentRate		float64		  `json:"currentRate"`
	MinTriggerValue float64		  `json:"minTriggerValue"`
	MaxTriggerValue float64		  `json:"maxTriggerValue"`
}

type Currency struct {
	Base 	string				`json:"base"`
	Date 	string 				`json:"date"`
	Rates	map[string]float64	`json:"rates"`
}



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
	}

	defer r.Body.Close()



	/* TESTING
	json.Marshal(&payload)
	json.NewEncoder(w).Encode
	*/
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


func main() {

//	port := os.Getenv("PORT")
//	http.HandleFunc("/webhook", HandleWebhook)
	http.HandleFunc("/", HandleWebhook)
//	http.ListenAndServe(":" + port, nil)
	http.ListenAndServe("localhost:8080", nil)
}