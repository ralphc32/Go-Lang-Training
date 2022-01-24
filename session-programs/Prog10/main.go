package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %s", err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("failed to parse csv file: %s", err)
	}
	fmt.Println(lines)

	fp2, _ := os.Create("output.csv")
	w := csv.NewWriter(fp2)
	for _, row := range lines {
		_ = w.Write(row)
	}
	w.Flush()
	fp2.Close()
}
