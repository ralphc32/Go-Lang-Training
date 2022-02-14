package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	ID        int    `json:"id"`
	Cust_Name string `json:"Customer_Name"`
	City_Name string `json:"City"`
}

func retrieveData() []*Customer {
	db, err := sql.Open("mysql", "default_user:1234@tcp(db:3306)/list_db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	output, err := db.Query("SELECT * FROM DIM_CUSTOMER")

	if err != nil {
		panic(err)
	}

	var customers []*Customer

	for output.Next() {
		var u Customer

		err = output.Scan(&u.ID, &u.Cust_Name, &u.City_Name)
		if err != nil {
			panic(err)
		}

		customers = append(customers, &u)
	}

	return customers
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Capstone DB homepage")
	fmt.Printf("Accessed homepage")
}

func ViewCustomers(w http.ResponseWriter, r *http.Request) {
	customers := retrieveData()

	fmt.Println("Accessed View")
	json.NewEncoder(w).Encode(customers)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/customers", ViewCustomers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
