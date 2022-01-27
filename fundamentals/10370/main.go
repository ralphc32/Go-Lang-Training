package main

import "fmt"

func main() {
	var x, y int

	// how many test cases
	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		// how many variables
		fmt.Scan(&y)

		// initialize group of numbers as input
		numbers := make([]int, y)

		// preset variables for counting
		count := 0
		sum := 0

		for j := range numbers {
			fmt.Scan(&numbers[j])

			sum += numbers[j]
		}

		// get mean of numbers
		mean := sum / y

		for _, z := range numbers {
			if z > mean {
				count++
			}
		}
		fmt.Printf("%.3f%%\n", float32(count)*100/float32(y))
	}
}
