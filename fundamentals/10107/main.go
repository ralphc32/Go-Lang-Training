package main

import (
	"fmt"
	"sort"
)

func main() {

	var x int
	var arrays []int
	for {
		if _, err := fmt.Scan(&x); err != nil {
			break
		}
		arrays = append(arrays, x)
		sort.Ints(arrays)
		if length := len(arrays); length%2 == 0 {
			fmt.Println((arrays[length/2-1] + arrays[length/2]) / 2)
		} else {
			fmt.Println(arrays[length/2])
		}
	}
}
