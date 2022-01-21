package main

import "fmt"

func main() {
	var z, i, j, c, a int

	fmt.Scan(&z)

	for i = 0; i < z; i++ {
		d := 0
		fmt.Scan(&a)

		for j = 0; j < a; j++ {
			fmt.Scan(&c)
			if c > d {
				d = c
			}
		}

		fmt.Printf("Case %d: %d\n", i+1, d)
	}
}
