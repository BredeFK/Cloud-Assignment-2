package gofiles

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func DiscordOperator(someText string, discordURL string) {
	info := WebhookInfo{}
	info.Content = someText + "\n"
	raw, _ := json.Marshal(info)
	resp, err := http.Post(discordURL, "application/json", bytes.NewBuffer(raw))
	if err != nil {
		log.Println(err)
		log.Println(ioutil.ReadAll(resp.Body))
	}

	log.Println(resp.StatusCode)
}
