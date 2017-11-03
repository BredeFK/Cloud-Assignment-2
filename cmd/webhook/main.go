package main

import (
	"Oblig2_Heroku/gofiles"
	"time"
)

func main(){

	// Heroku scheduler =  16:30

	triggerTime := "2300"

	for {
		tempTime := time.Now().Local()
		timeNow := tempTime.Format("1504")	// format to 2400

		if timeNow == triggerTime {
			gofiles.DailyCurrencyAdder()
			gofiles.CheckTrigger()
		}
		delay := time.Minute
		time.Sleep(delay)
	}
}


