package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 21 * time.Second,
	}
	resp, err := client.Get("http://tools.ietf.org/rfc/rfc7540.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
