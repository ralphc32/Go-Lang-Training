package main

import "fmt"

func compute(a, b int) int {
	var channel int
	if a < b {
		channel = b - a
	} else {
		channel = a - b
	}
	if channel > 50 {
		channel = 100 - channel
	}
	return channel
}

func main() {
	var a, b int

	for {
		if fmt.Scan(&a, &b); a == -1 && b == -1 {
			break
		} else {
			fmt.Println(compute(a, b))
		}
	}
}
