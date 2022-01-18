package main

import "fmt"

func main() {
	var a, b, i int
	// fmt.Scan(&x, &y)
	// for i := x; i <= y; i = x + 2 {
	// 	z += i
	// }
	// fmt.Printf("The sum is %v", x)
	if a%2 == 0 {

		// then the smallest odd integer we want is one greater than a
		a++
	}

	// if b is even
	if b%2 == 0 {

		// then the largest odd integer we want is one less than b
		b--
	}

	// find the index of b in the sequence
	i = (b - a + 2) / 2

	// calculate the sum of the first i numbers in the sequence
	fmt.Println(i * (a + (a + 2*(i-1))) / 2)
}
