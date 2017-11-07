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
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", gofiles.HandleWebhook)
	http.HandleFunc("/latest", gofiles.HandleLatest)
	http.HandleFunc("/average", gofiles.HandleAverage)
	http.HandleFunc("/evaluationtrigger", gofiles.HandleTestTrigger)
	http.ListenAndServe(":"+port, nil)
}
