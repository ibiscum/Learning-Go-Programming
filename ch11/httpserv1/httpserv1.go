package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// simple server, no multiplex
type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "text/html")
	resp.WriteHeader(http.StatusOK)
	fmt.Fprint(resp, m)
}

func main() {
	msgHandler := msg("Hello from high above!")
	server := http.Server{
		Addr:         ":4040",
		Handler:      msgHandler,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 3,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
