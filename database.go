package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)


func SetupDB() *MongoDB {
	db := MongoDB{
		"mongodb://localhost",
		"2imt2681",
		"webhookCollection",
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
		DropDups:	true,
		Background: true,
		Sparse: 	true,
	}

	err = session.DB(db.DatabaseName).C(db.CollectionName).EnsureIndex(index)
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

	err = session.DB(db.DatabaseName).C(db.CollectionName).Insert(p)

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
	allWasGood := true

	err = session.DB(db.DatabaseName).C(db.CollectionName).Find(bson.M{"ID": keyID}).One(&payload)
	if err != nil {
		allWasGood = false
	}

	return payload, allWasGood
}