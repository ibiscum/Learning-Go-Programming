package main

import (
	"bytes"
	"os"
)

func main() {
	var books bytes.Buffer
	books.WriteString("The Great Gatsby\n")
	books.WriteString("1984\n")
	books.WriteString("A Tale of Two Cities\n")
	books.WriteString("Les Miserables\n")
	books.WriteString("The Call of the Wild")

	books.WriteTo(os.Stdout)
}
