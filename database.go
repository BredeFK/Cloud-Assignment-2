package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
)


func SetupDB() *MongoDB {
	db := MongoDB{
		"mongodb://localhost",
		"currencyDB",
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