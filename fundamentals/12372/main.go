package main

import "fmt"

func main() {
	var l, w, h, x int

	const case_size = 20

	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		if fmt.Scan(&l, &w, &h); l <= case_size && w <= case_size && h <= case_size {
			fmt.Printf("Case %d: good\n", i)
		} else {
			fmt.Printf("Case %d: bad\n", i)
		}
	}
}
