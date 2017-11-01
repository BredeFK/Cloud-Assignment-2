package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

func DiscordOperator (someText string, discordURL string) {
	info := WebhookInfo{}
	info.Content = someText + "\n"
	raw, _ := json.Marshal(info)
	resp, err := http.Post(discordURL, "application/json", bytes.NewBuffer(raw))
	if err != nil {
		fmt.Println(err)
		fmt.Println(ioutil.ReadAll(resp.Body))
	}

	fmt.Println(resp.StatusCode)
}