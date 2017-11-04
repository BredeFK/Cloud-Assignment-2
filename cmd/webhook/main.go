package main

import (
	"Oblig2_Heroku/gofiles"
	"time"
)

func main(){

	// Heroku scheduler =  16:30

	triggerTime := "12"

	tempTime := time.Now().Local()
	timeNow := tempTime.Format("15")	// format to 2400

	if timeNow == triggerTime {
		gofiles.DailyCurrencyAdder()
		gofiles.CheckTrigger()
	}

}


