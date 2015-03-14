package server

import (
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func DummyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Works")
}

func TestGetOnly(t *testing.T) {
	recorder := httptest.NewRecorder()

	url := "http://foo-bar-blahblah.com"
	request, _ := http.NewRequest("GET", url, nil)

	GetOnly(DummyHandler)(recorder, request)

	if(recorder.Body.String() != "Works") {
		t.Error("should allow get, is failing")
	}

}

func TestDataHandler(t *testing.T) {
	recorder := httptest.NewRecorder()

	url := "http://foo-bar-blahblah.com/data/1"
	request, _ := http.NewRequest("GET", url, nil)

	DataHandler(recorder, request)

	if(recorder.Body.String() != "Data ID: 1") {
		t.Error("should allow get, is failing")
	}
}