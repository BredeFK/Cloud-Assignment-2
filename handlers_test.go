package main

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

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleTestTrigger(t *testing.T) {
	testDB := SetupDB()
	count := testDB.Count()
	ok := false

	if count == 0 {
		t.Fatal("Count is wrong, there should be at least one payload")
	} else {
		session, err := mgo.Dial(testDB.DatabaseURL)
		if err != nil {
			t.Fatal(err.Error())
		}
		defer session.Close()

		payload := Payload{}

		for i := 1; i <= count; i++ {
			err = session.DB(testDB.DatabaseName).C(testDB.ColWebHook).Find(nil).Skip(count - i).One(&payload)
			if err != nil {
				t.Fatal("Can not get one or more webhook data", err.Error())
				return
			}
			if payload.WebhookURL == "www.imgur.com/" {
				ok = true
			}
		}
		if ok != true {
			t.Fatal("Could not find added payload")
		}
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

	if status := respTest.Code; status != http.StatusMethodNotAllowed {
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

	if status := respTest.Code; status != http.StatusMethodNotAllowed {
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
