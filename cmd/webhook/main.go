//==================================================================================================\\
// 		   AUTHOR: 	Brede Fritjof Klausen		  				  								    \\
// 		  SUBJECT: 	IMT2681 Cloud Technologies													    \\
//==================================================================================================\\
//	SOURCES:												 									    \\
// * https://stackoverflow.com/questions/38127583/get-last-inserted-element-from-mongodb-in-golang  \\
// * https://elithrar.github.io/article/testing-http-handlers-go/								    \\
//==================================================================================================\\

package main

import (
	"Oblig2_Heroku/gofiles"
)

func main() {

	// Heroku scheduler =  15:30 UTC

	gofiles.DailyCurrencyAdder()
	gofiles.CheckTrigger()
}
