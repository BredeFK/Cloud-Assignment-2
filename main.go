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
)


func main() {

//	port := os.Getenv("PORT")
// 	http.ListenAndServe(":" + port, nil)

	http.HandleFunc("/", HandleWebhook)
	http.HandleFunc("/latest", HandleLatest)
	http.HandleFunc("/average", HandleAverage)
	http.ListenAndServe("localhost:8080", nil)

//	text := "Sup mah dudes!"
//	DiscordOperator(text, DiscordURL_notAbot)
}