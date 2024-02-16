package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("./iopipe.data")
	if err != nil {
		fmt.Println("File creation failed:", err)
		os.Exit(1)
	}

	pr, pw := io.Pipe() // pipe reader/writer
	go func() {
		fmt.Fprint(pw, "Pipe streaming")
		pw.Close()
	}()

	wait := make(chan struct{})
	go func() {
		_, err = io.Copy(file, pr)
		if err != nil {
			log.Fatal(err)
		}
		pr.Close()
		close(wait)
	}()
	<-wait
}
