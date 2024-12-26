package main

import (
	"fmt"
	"log"
)

func main() {

	go func() {
		count(60, 100, 10)
	}()

	_, err := fmt.Scanln() // wait for enter
	if err != nil {
		log.Fatal(err)
	}
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}
