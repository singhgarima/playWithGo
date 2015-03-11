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

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/data/{id:[0-9]+}", DataHandler)
	http.ListenAndServe(":4321", nil)
}