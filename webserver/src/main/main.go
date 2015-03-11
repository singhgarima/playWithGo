package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello WOrld!")
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/data/"):]
    fmt.Fprintf(w, "Data ID: %s", id)
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

func main() {
	http.HandleFunc("/", GetOnly(HomeHandler))
	http.HandleFunc("/data/{id:[0-9]+}", GetOnly(DataHandler))
	http.ListenAndServe(":4321", nil)
}