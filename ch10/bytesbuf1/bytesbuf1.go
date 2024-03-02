package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	var books bytes.Buffer
	books.WriteString("The Great Gatsby\n")
	books.WriteString("1984\n")
	books.WriteString("A Take of Two Cities\n")
	books.WriteString("Les Miserables\n")
	books.WriteString("The Call of the Wild\n")

	file, err := os.Create("./books.data")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}
	defer file.Close()
	_, err = books.WriteTo(file)
	if err != nil {
		log.Fatal(err)
	}
}
