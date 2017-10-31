package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)


func SetupDB() *MongoDB {
	db := MongoDB{
		"mongodb://localhost",		// TODO : deploy mongodb
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

func DailyCurrencyAdder(){	// TODO : Add a timer to make this automatic daily


	currency := GetCurrency()

	db := SetupDB()
	db.Init()
	db.AddCurrency(currency)

}

