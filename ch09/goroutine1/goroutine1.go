package main

import (
	"fmt"
	"log"
)

func main() {
	go count(10, 30, 10)
	go count(40, 60, 10)
	go count(70, 120, 20)
	_, err := fmt.Scanln() // blocks for kb input
	if err != nil {
		log.Fatal(err)
	}
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}
