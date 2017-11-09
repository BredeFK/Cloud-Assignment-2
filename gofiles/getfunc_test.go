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
	"net/http"
	"testing"
	"time"
)

func TestGetCurrency(t *testing.T) {
	testURL := "http://api.fixer.io/2016-12-30?base=EUR"

	client := http.Client{
		Timeout: time.Second * 10,
	}

	currency := GetCurrency(testURL, client)

	base := "EUR"
	date := "2016-12-30"
	target := "NOK"
	rate := 9.0863

	if currency.Base != base {
		t.Fatalf("Error! got '%s' instead of '%s'", currency.Base, base)
	}

	if currency.Date != date {
		t.Fatalf("Error! got '%s' instead of '%s'", currency.Date, date)
	}

	if currency.Rates[target] != rate {
		t.Fatalf("Error! got '%v' instead of '%v'", currency.Rates[target], rate)
	}
}
