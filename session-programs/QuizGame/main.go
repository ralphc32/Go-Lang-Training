package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	database_n := flag.String("csv", "questions.csv", "database")
	flag.Parse()

	fp, err := os.Open(*database_n)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	lines, _ := r.ReadAll()
	for _, item := range lines {
		var ans string
		fmt.Printf("%s = ", item[0])
		fmt.Scan(&ans)
		if ans == strings.TrimLeft(item[1], " ") {
			fmt.Println("Correct!")
		}
	}
}
