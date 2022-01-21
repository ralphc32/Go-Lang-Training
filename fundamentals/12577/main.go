package main

import "fmt"

func main() {
	var input string

	for i := 1; ; i++ {
		if fmt.Scan(&input); input == "*" {
			break
		} else if string(input[0]) == "H" {
			fmt.Printf("Case %d: Hajj-e-Akbar\n", i)
		} else {
			fmt.Printf("Case %d: Hajj-e-Asghar\n", i)
		}
	}
}
