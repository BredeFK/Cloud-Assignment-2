package main

import (
	"bitbucket.org/Brede_F_Klausen/assignment2_cloud/gofiles"
)

func main() {

	// Heroku scheduler =  15:30 UTC

	gofiles.DailyCurrencyAdder()
	gofiles.CheckTrigger()
}
