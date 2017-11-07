package main

import (
	"Oblig2_Heroku/gofiles"
)

func main() {

	// Heroku scheduler =  15:30 UTC

	gofiles.DailyCurrencyAdder()
	gofiles.CheckTrigger()
}
