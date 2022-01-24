package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("file", "input.txt", "program database")
	flag.Parse()
	fmt.Println(filename)
}
