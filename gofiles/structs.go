//=========================================================================================================//
// 		   AUTHOR: 	Brede Fritjof Klausen                                                                  //
// 		  SUBJECT: 	IMT2681 Cloud Technologies                                                             //
//=========================================================================================================//
//	SOURCES:                                                                                               //
// * https://stackoverflow.com/questions/38127583/get-last-inserted-element-from-mongodb-in-golang         //
// * https://elithrar.github.io/article/testing-http-handlers-go/                                          //
// * https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format //
//=========================================================================================================//

package gofiles

import (
	"gopkg.in/mgo.v2/bson"
)

// WebhookInfo struct
type WebhookInfo struct {
	Content string `json:"content"`
}

// Payload struct
type Payload struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	WebhookURL      string        `json:"webhookURL"`
	BaseCurrency    string        `json:"baseCurrency"`
	TargetCurrency  string        `json:"targetCurrency"`
	MinTriggerValue float64       `json:"minTriggerValue"`
	MaxTriggerValue float64       `json:"maxTriggerValue"`
}

// Currency struct
type Currency struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

// MongoDB struct
type MongoDB struct {
	DatabaseURL  string
	DatabaseName string
	ColWebHook   string
	ColCurrency  string
}
