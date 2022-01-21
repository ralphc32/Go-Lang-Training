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

	var x, x1, x2 int

	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		fmt.Scan(&x1, &x2)
		if x1 > x2 {
			fmt.Println(">")
		} else if x1 < x2 {
			fmt.Println("<")
		} else {
			fmt.Println("=")
		}
	}
}
