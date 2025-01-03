package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f1, err := os.OpenFile("./file1.go", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer f1.Close()

	f2, err := os.Create("./file1.bkp")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer f2.Close()

	f3, err := os.OpenFile("./file1.bkp", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer f3.Close()

	n, err := io.Copy(f2, f3)
	if err != nil {
		fmt.Println("Failed to copy:", err)
		os.Exit(1)
	}

	fmt.Printf("Copied %d bytes from %s to %s\n", n, f1.Name(), f2.Name())
}
