package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	str := strings.NewReader("The quick brown" +
		"fox jumps over the lazy dog")
	section := io.NewSectionReader(str, 19, 23)
	_, err := io.Copy(os.Stdout, section)
	if err != nil {
		log.Fatal(err)
	}
}
