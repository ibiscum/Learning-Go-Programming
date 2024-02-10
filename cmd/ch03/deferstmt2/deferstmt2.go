//go:build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func printHtml(address string) {
	rsp, err := http.Get(address)
	if err != nil {
		fmt.Println("Unable to reach", address)
		return
	}

	if body, err := ioutil.ReadAll(rsp.Body); err != nil {
		return
	} else {
		fmt.Println(string(body))
	}
}
