package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

const DiscordURL_FY = "https://discordapp.com/api/webhooks/370393359900082200/_RASdjfNlTsFm9QMprDIFukfV05u7_vfN8nBjgoJ7y0_D_JmLXYdoWVbY8guoCkbOAVx"
const DiscordURL_Bender = "https://discordapp.com/api/webhooks/374593069007503360/AFWh44UGmY5lWubuRd_GfiIag9JQ6tUsHr-UMioob2vdLixIray1lkmgxAcv6HKtOiDc"
const DiscordURL_notAbot = "https://discordapp.com/api/webhooks/374710634161504256/YjnETeldqzORKMLUFDU-u5Ocz8SEKyeCh-k1TKhXx_c7mrkgLTyLq199Ko5nP8c4sN4J"



func DiscordOperator (someText string, discordURL string) {
	info := WebhookInfo{}
	info.Content = someText + "\n"
	raw, _ := json.Marshal(info)
	resp, err := http.Post(discordURL, "application/json", bytes.NewBuffer(raw))
	if err != nil {
		fmt.Println(err)
		fmt.Println(ioutil.ReadAll(resp.Body))
	}
}