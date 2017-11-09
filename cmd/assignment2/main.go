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
	"net/http"

)

func main() {
//	port := os.Getenv("PORT")
	http.HandleFunc("/", gofiles.HandleWebhook)
	http.HandleFunc("/latest", gofiles.HandleLatest)
	http.HandleFunc("/average", gofiles.HandleAverage)
	http.HandleFunc("/evaluationtrigger", gofiles.HandleTestTrigger)
//	http.ListenAndServe(":"+port, nil)
	http.ListenAndServe("localhost:8080", nil)
}
