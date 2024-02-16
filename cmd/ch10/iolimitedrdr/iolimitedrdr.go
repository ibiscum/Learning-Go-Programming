package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	str := strings.NewReader("The quick brown " +
		"fox jumps over the lazy dog")
	limited := &io.LimitedReader{R: str, N: 19}
	_, err := io.Copy(os.Stdout, limited)
	if err != nil {
		log.Fatal(err)
	}
}
