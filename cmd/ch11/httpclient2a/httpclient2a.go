package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Dial: (&net.Dialer{
				Timeout: 30 * time.Second,
			}).Dial,
		},
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
