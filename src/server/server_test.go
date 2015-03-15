package server

import (
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
	"lru_cache"
)

func DummyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Works")
}

func TestGetOnly(t *testing.T) {
	//setup
	recorder := httptest.NewRecorder()
	url := "http://foo-bar-blahblah.com"
	request, _ := http.NewRequest("GET", url, nil)

	//test
	GetOnly(DummyHandler)(recorder, request)

	//assertions
	if(recorder.Body.String() != "Works") {
		t.Error("should allow get, is failing")
	}

}

func TestDataHandlerSuccess(t *testing.T) {
	//setup
	cache := lru_cache.CreateCache(10)
	serverCacheStore = &cache
	serverCacheStore.StoreValueForKey("some value", 1)
	
	recorder := httptest.NewRecorder()
	url := "http://foo-bar-blahblah.com/data/1"
	request, _ := http.NewRequest("GET", url, nil)

	//test
	DataHandler(recorder, request)

	//assertions
	if(recorder.Body.String() != "Data for ID: 1 is some value") {
		t.Error("Incorrect Data received")
	}
}

func TestDataHandlerDailure(t *testing.T) {
	//setup
	cache := lru_cache.CreateCache(10)
	serverCacheStore = &cache
	serverCacheStore.StoreValueForKey("some value", 1)
	
	recorder := httptest.NewRecorder()
	url := "http://foo-bar-blahblah.com/data/2"
	request, _ := http.NewRequest("GET", url, nil)

	//test
	DataHandler(recorder, request)

	//assertions
	if(recorder.Body.String() != "404 page not found\n") {
		t.Error("Expecting Body as 404 page not found")
	}
	if(recorder.Code != http.StatusNotFound) {
		t.Error("Incorrect Status Code expecting 404")
	}
}

// func TestMain(t *testing.T) {
// 	var cache2 lru_cache.CacheStore
// 	cache := lru_cache.CreateCache(10)
// 	fmt.Println("Main")
// 	fmt.Println(reflect.TypeOf(cache))
// 	fmt.Println(reflect.TypeOf(cache2))

// 	t.Error("should allow get, is failing")
// }