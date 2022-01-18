package main

import "fmt"

func main() {
	//var a, b, i int
	// fmt.Scan(&x, &y)
	// for i := x; i <= y; i = x + 2 {
	// 	z += i
	// }
	// fmt.Printf("The sum is %v", x)
	var x, a, b, sum int

	fmt.Scan(&x)

	for i := 1; i <= x; i++ {

		sum = 0
		fmt.Scan(&a, &b)

		for j := 1; j <= b; j++ {
			if j%2 == 1 {
				sum += j
			}
		}
		fmt.Printf("Case %d: %d\n", i, sum)
	}
}
