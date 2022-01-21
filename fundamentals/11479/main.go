package main

import (
	"fmt"
)

func main() {

	var cases, x, y, z int
	fmt.Scan(&cases)
	for i := 1; i <= cases; i++ {
		if _, err := fmt.Scan(&x, &y, &z); err != nil {
			break
		}
		fmt.Printf("Case %d: ", i)
		switch {
		case x <= 0 || y <= 0 || z <= 0 || x+y <= z || y+z <= x || x+z <= y:
			fmt.Println("Invalid")
		case x == y && x == z:
			fmt.Println("Equilateral")
		case x == y || x == z || y == z:
			fmt.Println("Isosceles")
		default:
			fmt.Println("Scalene")
		}
	}
}
