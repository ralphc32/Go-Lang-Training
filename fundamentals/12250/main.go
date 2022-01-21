package main

import (
	"fmt"
	"os"
)

var keywords = map[string]string{
	"HELLO":        "ENGLISH",
	"HOLA":         "SPANISH",
	"HALLO":        "GERMAN",
	"BONJOUR":      "FRENCH",
	"CIAO":         "ITALIAN",
	"ZDRAVSTVUJTE": "RUSSIAN"}

func solve(word string) string {
	if country, ok := keywords[word]; ok {
		return country
	}
	return "UNKNOWN"
}

func main() {
	in, _ := os.Open("12250.in")
	defer in.Close()
	out, _ := os.Create("12250.out")
	defer out.Close()

	var word string
	for i := 1; ; i++ {
		if fmt.Fscan(in, &word); word == "#" {
			break
		}
		fmt.Fprintf(out, "Case %d: %s\n", i, solve(word))
	}
}
