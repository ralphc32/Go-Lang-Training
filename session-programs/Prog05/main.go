package main

import "fmt"

func intseq(start int) func() int {
	i := start - 1
	return func() int {
		i = i + 2
		return i
	}
}

func main() {
	nextID := intseq(1001)
	fmt.Println(nextID())
	fmt.Println(nextID())
	anotherNextID := intseq(100001)
	fmt.Println(anotherNextID())

}
