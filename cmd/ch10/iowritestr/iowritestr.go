package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fout, err := os.Create("./iowritestr.data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fout.Close()
	_, err = io.WriteString(fout, "Hello there!\n")
	if err != nil {
		log.Fatal(err)
	}
}
