package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com/")
	defer resp.Body.Close()
	if err == nil {
		fmt.Printf(resp.Status)
	}
	
}