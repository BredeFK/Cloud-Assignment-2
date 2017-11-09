//=========================================================================================================//
// 		   AUTHOR: 	Brede Fritjof Klausen                                                                  //
// 		  SUBJECT: 	IMT2681 Cloud Technologies                                                             //
//=========================================================================================================//
//	SOURCES:                                                                                               //
// * https://stackoverflow.com/questions/38127583/get-last-inserted-element-from-mongodb-in-golang         //
// * https://elithrar.github.io/article/testing-http-handlers-go/                                          //
// * https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format //
//=========================================================================================================//

package main

import (
	"bitbucket.org/Brede_F_Klausen/assignment2_cloud/gofiles"
	"time"
)

func main() {
	// Heroku scheduler =  15:30 UTC

	// Get weekday for today
	tempDay := time.Now().Local()

	// Convert to string
	day := tempDay.Format("Monday")

	// If it is weekend, don't get currency(fixer.io does not update in the weekends)
	if day != "Saturday" && day != "Sunday" {

		// Add currency for today
		gofiles.DailyCurrencyAdder()

		// Check if it gets triggered
		gofiles.CheckTrigger()
	}
}
