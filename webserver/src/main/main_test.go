package main_test

import (
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
	. "main"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Works")
}

type MyWriter http.ResponseWriter

func TestGetOnly(t *testing.T) {
	recorder := httptest.NewRecorder()
	
	url := "http://foo-bar-blahblah.com"	
	request, _ := http.NewRequest("GET", url, nil)
	
	GetOnly(TestHandler)(recorder, request)
	
	if(recorder.Body.String() != "Works") {
		t.Error("should allow get, is failing")
	}
	
}