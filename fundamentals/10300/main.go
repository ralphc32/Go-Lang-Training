package main

import "fmt"

func main() {
	var x, y int
	var a, b, c uint64

	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		var d uint64

		fmt.Scan(&y)
		for j := 1; j <= y; j++ {
			fmt.Scan(&a, &b, &c)
			d += a * c
		}
		fmt.Println(d)
	}
}
