package main

import (
	"fmt"
	"log"
	"net/http"

	"gonum.org/v1/gonum/mat"
)

func main() {
	http.HandleFunc("/solve", solver)
	http.ListenAndServe(":8080", nil)
}

func solver(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	v, ok := r.Form["coef"]

	fmt.Fprintf(w, "%v %v\n", v, ok)

	var a1, a2, a3, a4, b1, b2, b3, b4, c1, c2, c3, c4 float64

	if n, _ := fmt.Sscanf(v[0], "%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f", &a1, &a2, &a3, &a4,
		&b1, &b2, &b3, &b4,
		&c1, &c2, &c3, &c4); n == 12 {

		// Create 3 variables gomatrix for 3 equations
		var variables = mat.NewDense(3, 3, []float64{
			a1, a2, a3,
			b1, b2, b3,
			c1, c2, c3,
		})

		// Setup constants array
		var con = []float64{a4, b4, c4}

		// Setup placeholder array for coefficient answers
		ans := make([]float64, len(con))

		b := make([]float64, len(con))
		d := mat.Det(variables)

		// Compute for determinants
		for c := range con {
			mat.Col(b, c, variables)
			variables.SetCol(c, con)
			ans[c] = mat.Det(variables) / d
			variables.SetCol(c, b)
		}

		fmt.Fprintf(w, "x = %.2f, y = %.2f, z = %.2f\n", ans[0], ans[1], ans[2])
	}
}
