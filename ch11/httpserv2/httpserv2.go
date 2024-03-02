package main

import (
	"fmt"
	"log"
	"net/http"
)

// simple handler, no multiplex
type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "text/html")
	resp.WriteHeader(http.StatusOK)
	fmt.Fprint(resp, m)
}

// using default server
func main() {
	msgHandler := msg("Hello from high above!")
	err := http.ListenAndServe(":4040", msgHandler)
	if err != nil {
		log.Fatal(err)
	}
}
