package main

import (
	"fmt"
	"log"
)

func main() {

	start := 0
	stop := 50
	step := 5

	go func() {
		count(start, stop, step)
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
