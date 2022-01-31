package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "betty bought the butter , the butter was bitter , " +
		"betty bought more butter to make the bitter butter better "
	for index, element := range repetition(input) {
		fmt.Printf("%v %d\n", index, element)
	}
}

func repetition(st string, ch chan string) map[string]int {

	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	ch <- fmt.Printf("%v %d\n", index, element)
}
