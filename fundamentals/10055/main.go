package main

import (
	"fmt"
	// "os"
)

func main() {
	// x, _ := os.Open("10055.in")
	// defer x.Close()
	// y, _ := os.Create("10055.out")
	// defer y.Close()

	var a, b uint64
	for {
		if _, err := fmt.Scan(&a, &b); err != nil {
			break
		}
		fmt.Println(b - a)
	}
}
