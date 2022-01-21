package main

import (
	"fmt"
)

func getCriteria(a, b, c int) bool {
	switch {
	case a >= b && a >= c:
		return b*b+c*c == a*a
	case b >= a && b >= c:
		return a*a+c*c == b*b
	default:
		return a*a+b*b == c*c
	}
}

func main() {

	var a, b, c int
	for {
		if fmt.Scan(&a, &b, &c); a == 0 && b == 0 && c == 0 {
			break
		}
		if getCriteria(a, b, c) {
			fmt.Println("right")
		} else {
			fmt.Println("wrong")
		}
	}
}
