package main

import(
	"gopkg.in/mgo.v2/bson"
)

type WebhookInfo struct {
	Content string `json:"content"`
}

type Payload struct {
	ID 				bson.ObjectId `bson:"_id,omitempty"`
	WebhookURL  	string 		  `json:"webhookURL"`
	BaseCurrency 	string		  `json:"baseCurrency"`
	TargetCurrency	string		  `json:"targetCurrency"`
//	CurrentRate		float64		  `json:"currentRate"`	// TODO : Maybe remove
	MinTriggerValue float64		  `json:"minTriggerValue"`
	MaxTriggerValue float64		  `json:"maxTriggerValue"`
}

type Currency struct {
	Base 	string				`json:"base"`
	Date 	string 				`json:"date"`
	Rates	map[string]float64	`json:"rates"`
}

type MongoDB struct {
	DatabaseURL    string
	DatabaseName   string
	ColWebHook 	   string
	ColCurrency	   string
}