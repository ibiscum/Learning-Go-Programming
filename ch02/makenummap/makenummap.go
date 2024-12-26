// This program geneartes a bitmap string.
// The string is generated by looping from numbers 0...max.
// It randomly selects ith number and adds it to the map.
// If selected, a 1 is added to the bitmap list otherwise zero.
// The result is saved in file nummap.txt
//
// Run this to generate file, then run nummap.go.
//
// Example based on book Program Pearls (Jon Bentley), Column 1.

package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var max, fileMode int
var mapFileName string

func makeBitMap() string {
	var numbuff bytes.Buffer

	i := 0
	for i = 0; i < max; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		selected := r.Intn(2)
		if selected == 1 {
			numbuff.WriteString("1")
		} else {
			numbuff.WriteString("0")
		}
	}
	fmt.Println("Mapped", i, "values")
	return numbuff.String()
}

func main() {
	max = 10
	fileMode = 4000
	mapFileName = "nummap.txt"

	nummap := makeBitMap()
	err := os.WriteFile(mapFileName, []byte(nummap), os.FileMode(fileMode))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Bitmap file", mapFileName, "created OK")
}