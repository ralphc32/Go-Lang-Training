package main

import (
	"fmt"
	/* if input output */ // "os"
)

func main() {
	/* if input output */
	// in, _ := os.Open("11172.in")
	// defer in.Close()
	// out, _ := os.Create("11172.out")
	// defer out.Close()

	var x, x1, x2, x3 int
	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		fmt.Scan(&x1, &x2, &x3)

		if (x1 > x2 && x1 < x3) || (x1 > x3 && x1 < x2) {
			fmt.Println(x1)
		} else if (x2 > x1 && x2 < x3) || (x2 < x1 && x2 > x3) {
			fmt.Println(x2)
		} else {
			fmt.Println(x3)
		}
	}
}
