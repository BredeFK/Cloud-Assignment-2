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
