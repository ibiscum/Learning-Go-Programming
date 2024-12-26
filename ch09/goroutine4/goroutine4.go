package main

import (
	"fmt"
	"log"
)

func main() {

	starts := []int{10, 40, 70, 100}

	for _, j := range starts {
		j := j
		go func() {
			count(j, j+20, 10)
		}()
	}

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
