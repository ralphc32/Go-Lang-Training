package main

import "fmt"

type LeadRec struct {
	name, position, company string
	salary                  int
}

func main() {
	var x LeadRec
	x.name = "Joe Rizal"
	x.position = "Editor-in-Chief"
	x.company = "La Solidaridad"
	y := LeadRec{"Andy Bonifacio", "Supremo", "KKK", 50}
	fmt.Println(x, y)
}
