package main

import (
	"Oblig2_Heroku/gofiles"
)

func main(){

	// Heroku scheduler =  17:00 UTC

	/*
	const triggerTime = "17"

	tempTime := time.Now().UTC()
	timeNow := tempTime.Format("15")	// format to 2400

	if timeNow == triggerTime {
	*/
		gofiles.DailyCurrencyAdder()
		gofiles.CheckTrigger()
	// }

}


