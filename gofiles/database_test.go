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
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestMongoDB_GetLatest(t *testing.T) {

	today := "2017-11-07"

	testDB := SetupDB()
	currency, ok := testDB.GetLatest(today, 0)

	if ok == false {
		t.Fatalf("Couldn't get any data from " + today + "!")
	}

	base := "EUR"
	target := "THB"
	rate := 38.305

	if currency.Base != base {
		t.Fatalf("Error! got '%s' instead of '%s'", currency.Base, base)
	}

	if currency.Rates[target] != rate {
		t.Fatalf("Error! got '%v' instead of '%v'", currency.Rates[target], rate)
	}

	if currency.Date != today {
		t.Fatalf("Error! got '%s' instead of '%s'", currency.Date, today)
	}
}

func TestMongoDB_AddCurrency(t *testing.T) {
	p := Payload{"", "www.webhookURL.com/", "EUR", "RUB", 40.1, 100}

	testDB := SetupDB()
	testDB.Init()
	err := testDB.Add(p)
	if err != nil {
		t.Fatal("Error! Could not add new payload", err.Error())
	}
}

func TestMongoDB_Get(t *testing.T) {
	testDB := SetupDB()

	session, err := mgo.Dial(testDB.DatabaseURL)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer session.Close()

	payload := Payload{}

	err = session.DB(testDB.DatabaseName).C(testDB.ColWebHook).Find(bson.M{"webhookurl": "www.webhookURL.com/"}).One(&payload)
	if err != nil {
		t.Fatal("Could not get payload with webhookurl: www.webhookURL.com/", err.Error())
	}

	if payload.WebhookURL != "www.webhookURL.com/" || payload.BaseCurrency != "EUR" ||
		payload.TargetCurrency != "RUB" || payload.MinTriggerValue != 40.1 ||
		payload.MaxTriggerValue != 100 {
		t.Error("payload doesn't match!")
	}
}

func TestMongoDB_Delete(t *testing.T) {
	testDB := SetupDB()

	session, err := mgo.Dial(testDB.DatabaseURL)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer session.Close()

	err = session.DB(testDB.DatabaseName).C(testDB.ColWebHook).Remove(bson.M{"webhookurl": "www.webhookURL.com/"})
	if err != nil {
		t.Fatal("Could not delete payload with webhookurl: www.webhookURL.com/", err.Error())
	}
}
