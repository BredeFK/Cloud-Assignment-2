package main

import (
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"encoding/json"
)

func GetCurrency(URL string) Currency{

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil{
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Assignment")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil{
		log.Fatal(readErr)
	}

	currency := Currency{}
	jsonErr := json.Unmarshal(body, &currency)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return currency
}
