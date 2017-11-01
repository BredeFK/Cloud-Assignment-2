package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
)


func SetupDB() *MongoDB {
	db := MongoDB{
		"mongodb://fritjof:mlab123@ds241395.mlab.com:41395/2imt2681",		// TODO : deploy mongodb
		"2imt2681",
		"webhookCollection",
		"currencyCollection",
	}

	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()

	if err != nil{
		panic(err)
	}

	return &db
}

func (db * MongoDB) Init() {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil  {
		panic(err)
	}
	defer session.Close()

	index := mgo.Index{
		Key:		[]string{"currencyid"},
		Unique: 	true,
		DropDups:	false,
		Background: true,
		Sparse: 	true,
	}

	err = session.DB(db.DatabaseName).C(db.ColWebHook).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func (db *MongoDB) Add(p Payload) error {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.ColWebHook).Insert(p)

	if err != nil {
		fmt.Printf("Could not add to db, error in Insert(): %v", err.Error())
		return err
	}

	return nil
}

func (db *MongoDB) Get(keyID string) (Payload, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	payload := Payload{}
	ok := true

	err = session.DB(db.DatabaseName).C(db.ColWebHook).Find(bson.M{"_id": bson.ObjectIdHex(keyID)}).One(&payload)
	if err != nil {
		ok = false
	}

	return payload, ok
}

func (db *MongoDB) GetLatest(date string) (Currency, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	currency := Currency{}
	notToday := true

	err = session.DB(db.DatabaseName).C(db.ColCurrency).Find(bson.M{"date": date}).One(&currency)
	if err != nil {
		notToday = false
	}

	return currency, notToday
}

func (db *MongoDB) Delete(keyID string) bool {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	ok := true

	err = session.DB(db.DatabaseName).C(db.ColWebHook).Remove(bson.M{"_id": bson.ObjectIdHex(keyID)})
	if err != nil {
		ok = false
	}

	return ok
}

func (db *MongoDB) AddCurrency(c Currency) error {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.ColCurrency).Insert(c)

	if err != nil {
		fmt.Printf("Could not add currency to db, error in Insert(): %v", err.Error())
		return err
	}

	return nil
}

func (db *MongoDB) Count() int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	count, err := session.DB(db.DatabaseName).C(db.ColWebHook).Count()
	if err != nil{
		fmt.Printf("Error in Count(): %v", err.Error())
		return -1
	}

	return count
}

func DailyCurrencyAdder(){
	currency := GetCurrency()
	db := SetupDB()
	db.Init()
	db.AddCurrency(currency)
}

func CheckTrigger() {

	db := SetupDB()
	count := db.Count()

	if count > 0{
		webHook := Payload{}
		session, err := mgo.Dial(db.DatabaseURL)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		tempToday := time.Now().Local()
		today := tempToday.Format("2006-01-02")

		currency, ok := db.GetLatest(today)
		if ok == false{
			log.Printf("Error in CheckTrigger() | There isn't any data for today yet")
			return
		}

		for i := 1; i <= count; i++ {


			err = session.DB(db.DatabaseName).C(db.ColWebHook).Find(nil).Skip(count-i).One(&webHook)
			if err != nil{
				log.Printf("Error in CheckTrigger() | Can not get one or more webhooks", err)
				return
			}

			rate := currency.Rates[webHook.TargetCurrency]
			rateString := fmt.Sprint(rate)
			min := fmt.Sprint(webHook.MinTriggerValue)
			max := fmt.Sprint(webHook.MaxTriggerValue)


			if rate > webHook.MaxTriggerValue || rate < webHook.MinTriggerValue{
				text := "baseCurrency: " + webHook.BaseCurrency + "\ntargetCurrency: " + webHook.TargetCurrency + "\ncurrent: " + rateString + "\nminTriggerValue: " + min + "\nmaxTriggerValue: " + max
				DiscordOperator(text, webHook.WebhookURL)
			}

		}
	}else{
		fmt.Printf("Error in CheckTrigger() | There is no recorded data in the webhook collection")
	}

}