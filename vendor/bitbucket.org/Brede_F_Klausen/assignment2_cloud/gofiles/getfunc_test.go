package gofiles

import "testing"

func TestGetCurrency(t *testing.T) {
	testURL := "http://api.fixer.io/2016-12-30"

	currency := GetCurrency(testURL)

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
