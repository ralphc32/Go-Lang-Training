package main

import (
	"fmt"
)

var wordMap = map[string]string{"HELLO": "ENGLISH", "HOLA": "SPANISH", "HALLO": "GERMAN", "BONJOUR": "FRENCH", "CIAO": "ITALIAN", "ZDRAVSTVUJTE": "RUSSIAN"}

func solve(word string) string {
	if country, ok := wordMap[word]; ok {
		return country
	}
	return "UNKNOWN"
}

func main() {

	var word string
	for i := 1; ; i++ {
		if fmt.Scan("%s", &word); word == "#" {
			break
		}
		fmt.Printf("Case %d: %s\n", i, solve(word))
	}
}
