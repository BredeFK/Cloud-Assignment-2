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
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", HandleWebhook)
	http.HandleFunc("/latest", HandleLatest)
	http.HandleFunc("/average", HandleAverage)
	http.HandleFunc("/evaluationtrigger", HandleTestTrigger)
	http.ListenAndServe(":"+port, nil)
}
