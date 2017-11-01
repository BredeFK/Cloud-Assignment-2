//==================================================================================================\\
// 		   AUTHOR: 	Brede Fritjof Klausen		  				  								    \\
// 		  SUBJECT: 	IMT2681 Cloud Technologies													    \\
//==================================================================================================\\
//	SOURCES:												 									    \\
// * https://stackoverflow.com/questions/38127583/get-last-inserted-element-from-mongodb-in-golang  \\														 //
// * 																							    \\
//==================================================================================================\\
package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", HandleWebhook)
	http.HandleFunc("/latest", HandleLatest)
	http.HandleFunc("/add", HandleAdd)		// TODO : Remove this to automatic
	http.HandleFunc("/average", HandleAverage)
	http.HandleFunc("/evaluationtrigger", HandleTestTrigger)
	go http.ListenAndServe(":" + port, nil)
//	go http.ListenAndServe("localhost:8080", nil)

/*		// TODO : deploy mongodb and the uncomment this
	for {
		delay := time.Minute * 15
		time.Sleep(delay)

		DailyCurrencyAdder()
		CheckTrigger()
	}
*/
}