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
	"net/http"
	"time"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", HandleWebhook)
	http.HandleFunc("/latest", HandleLatest)
	http.HandleFunc("/daily/add", HandleAdd)		// Manual testing
	http.HandleFunc("/average", HandleAverage)
	http.HandleFunc("/evaluationtrigger", HandleTestTrigger)
	go http.ListenAndServe(":" + port, nil)
//	go http.ListenAndServe("localhost:8080", nil)		// for local testing

	cycle := 0
	for {
		delay := time.Minute * 10
		time.Sleep(delay)
		cycle += 10

		if cycle == 1440{
			cycle = 0
			DailyCurrencyAdder()
			CheckTrigger()
		}
	}
}