// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic:", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	fileName := os.Args[1]
	//fmt.Println(fileName)
	content, err := os.ReadFile(fileName)

	// convert the bytes to a string
	var text string
	if err != nil {
		panic(err)
	} else {
		text = string(content)
	}

	// initialize a map to store the counts
	counts := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(text)

	// capture the length of the words slice
	counts["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, character := range word {
			if unicode.IsLetter(character) {
				counts["letters"]++
			}
			if unicode.IsSymbol(character) {
				counts["symbols"]++
			}
			if unicode.IsNumber(character) {
				counts["numbers"]++
			}
		}

	}

	// dump the map to the console using the spew package
	spew.Dump(counts)
}
