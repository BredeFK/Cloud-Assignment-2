package main

import (
	"Oblig2_Heroku/gofiles"
	"time"
)

func main(){

	// Heroku scheduler =  11:00 UTC

	const triggerTime = "17"

	tempTime := time.Now().UTC()
	timeNow := tempTime.Format("15")	// format to 2400

	if timeNow == triggerTime {
		gofiles.DailyCurrencyAdder()
		gofiles.CheckTrigger()
	}

}


