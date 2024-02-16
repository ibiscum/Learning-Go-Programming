package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data := [][]byte{
		[]byte("The quick brown fox\n"),
		[]byte("jumps over the lazy dog\n"),
	}

	fout, err := os.Create("./filewrite.data")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fout.Close()

	for _, out := range data {
		_, err = fout.Write(out)
		if err != nil {
			log.Fatal(err)
		}
	}
}
