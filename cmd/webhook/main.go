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
<<<<<<< HEAD
	"Oblig2_Heroku/gofiles"
=======
	"bitbucket.org/Brede_F_Klausen/assignment2_cloud/gofiles"
>>>>>>> 0d9578bbeb1cb90d40d1a0ab9bb405a052d78ed0
)

func main() {

	// Heroku scheduler =  15:30 UTC

	gofiles.DailyCurrencyAdder()
	gofiles.CheckTrigger()
}
