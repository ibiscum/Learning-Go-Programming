package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	var books bytes.Buffer
	books.WriteString("The Great Gatsby\n")
	books.WriteString("1984\n")
	books.WriteString("A Tale of Two Cities\n")
	books.WriteString("Les Miserables\n")
	books.WriteString("The Call of the Wild")

	_, err := books.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
