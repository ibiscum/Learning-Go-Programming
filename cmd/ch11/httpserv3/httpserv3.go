package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	hello := func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)
		fmt.Fprint(resp, "Hello from Above!")
	}

	goodbye := func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)
		fmt.Fprint(resp, "Goodbye, it's been real!")
	}

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/goodbye", goodbye)

	err := http.ListenAndServe(":4040", mux)
	if err != nil {
		log.Fatal(err)
	}
}
