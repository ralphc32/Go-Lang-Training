package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	ch := make(chan map[string]int)
	for _, input := range os.Args[1:] {
		go WordCount(input, ch)
	}
	for range os.Args[1:] {
		for x, y := range <-ch {
			fmt.Printf("%v %d\n", x, y)
		}
	}
}

func WordCount(a string, ch chan map[string]int) {

	text := strings.Fields(a)

	// create map variable
	per_text := make(map[string]int)

	// evaluate input as individual text and count words
	for _, words := range text {
		_, condition := per_text[words]
		if condition {
			per_text[words] += 1
		} else {
			per_text[words] = 1
		}
	}

	// store map into channel
	ch <- per_text
}
