//=========================================================================================================//
// 		   AUTHOR: 	Brede Fritjof Klausen                                                                  //
// 		  SUBJECT: 	IMT2681 Cloud Technologies                                                             //
//=========================================================================================================//
//	SOURCES:                                                                                               //
// * https://stackoverflow.com/questions/38127583/get-last-inserted-element-from-mongodb-in-golang         //
// * https://elithrar.github.io/article/testing-http-handlers-go/                                          //
// * https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format //
//=========================================================================================================//

package gofiles

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlePOST(t *testing.T) {
	payload := Payload{"", "www.imgur.com/", "EUR", "NOK", 0.123, 21.0}
	testDB := SetupDB()
	testDB.Init()
	testDB.Add(payload)
}

func TestHandleGET(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	HandleGET(resp, req, "59fb10ef9d57c471e2a465e7")

	if status := resp.Code; status != http.StatusBadRequest {
		t.Errorf("Object id should not work, status code was %v, wanted %v",
			status, http.StatusBadRequest)
	}
}

func TestHandleWebhook(t *testing.T) {
	// Method can be GET, DELETE or POST | But none worked for this tester :/
	req, err := http.NewRequest("VIEW", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(HandleWebhook).ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v wanted %v",
			status, http.StatusBadRequest)
	}
}

func TestHandleLatest(t *testing.T) {
	// Only get and post will pass HandleLatest()
	reqTest, err := http.NewRequest("DELETE", "/latest", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	respTest := httptest.NewRecorder()

	HandleLatest(respTest, reqTest)

	if status := respTest.Code; status != http.StatusBadRequest {
		t.Errorf("Method has to be POST (or GET)")
	}

	// Testing with the correct method
	req, err := http.NewRequest("GET", "/latest", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	resp := httptest.NewRecorder()
	HandleLatest(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Could not get latest, status code was %v, expected %v",
			status, http.StatusOK)
	}
}

func TestHandleAverage(t *testing.T) {
	// Only get and post will pass HandleLatest()
	reqTest, err := http.NewRequest("DELETE", "/latest", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	respTest := httptest.NewRecorder()

	HandleAverage(respTest, reqTest)

	if status := respTest.Code; status != http.StatusBadRequest {
		t.Errorf("Method has to be POST (or GET)")
	}

	// Testing with the correct method
	req, err := http.NewRequest("GET", "/latest", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	resp := httptest.NewRecorder()
	HandleAverage(resp, req)
}

func TestHandleDELETE(t *testing.T) {
	// First test the handler
	req, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	HandleDELETE(resp, req, "59fb10ef9d57c471e2a465e7")

	if status := resp.Code; status != http.StatusBadRequest {
		t.Errorf("Object id should not work, status code was %v, wanted %v",
			status, http.StatusBadRequest)
	}

	// Then delete what was added in TestHandlePost()
	testDB := SetupDB()

	session, err := mgo.Dial(testDB.DatabaseURL)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer session.Close()

	err = session.DB(testDB.DatabaseName).C(testDB.ColWebHook).Remove(bson.M{"webhookurl": "www.imgur.com/"})
	if err != nil {
		t.Fatal("Could not delete payload with webhookurl: www.imgur.com/", err.Error())
	}
}
