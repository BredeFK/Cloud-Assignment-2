/////////////////////////////////////////////////////////////////////
// 		   AUTHOR: 	Brede Fritjof Klausen		  				  //
// STUDENT NUMBER: 	473211						 				 //
// 		  SUBJECT: 	IMT2681 Cloud Technologies					//
//=============================================================//
//	SOURCES:												  //
// * 														 //
// *														//
/////////////////////////////////////////////////////////////
package main

import (
	"net/http"
	"os"
)


func main() {
	http.HandleFunc("/", HandleWebhook)
	http.HandleFunc("/latest", HandleLatest)
	http.HandleFunc("/add", HandleAdd)		// TODO : Remove this to automatic
	http.HandleFunc("/average", HandleAverage)
//	http.ListenAndServe("localhost:8080", nil)

	port := os.Getenv("PORT")
	http.ListenAndServe(":" + port, nil)

}