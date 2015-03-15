package server

import (
	"fmt"
	"net/http"
	"strconv"
	"lru_cache"
)

var serverCacheStore *lru_cache.CacheStore

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello WOrld!")
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	id, conversionErr := strconv.Atoi(r.URL.Path[len("/data/"):])
	if ( conversionErr == nil ) {
		data, searchErr := serverCacheStore.GetValueForKey(id)
		if ( searchErr == nil ) {
	    	fmt.Fprintf(w, "Data for ID: %d is %s", id, data)
		} else {
			http.NotFound(w, r)
		}
	} else {
		http.NotFound(w, r)
	}
}

func GetOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "Only supports GET", http.StatusForbidden)
	}
}

func Start(cache *lru_cache.CacheStore) {
	serverCacheStore = cache

	http.HandleFunc("/", GetOnly(HomeHandler))
	http.HandleFunc("/data/", GetOnly(DataHandler))
	http.ListenAndServe(":4321", nil)
}