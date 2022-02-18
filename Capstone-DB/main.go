package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Link struct {
	ID         int    `json:"id"`
	Links      string `json:"Links"`
	Date_Added string `json:"DATE_ADDED"`
}

func retrieveData() []*Link {
	db, err := sql.Open("mysql", "default_user:1234@tcp(db:3306)/list_db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	output, err := db.Query("SELECT * FROM DIM_LYRICS")

	if err != nil {
		panic(err)
	}

	var links []*Link

	for output.Next() {
		var u Link

		err = output.Scan(&u.ID, &u.Links, &u.Date_Added)
		if err != nil {
			panic(err)
		}

		links = append(links, &u)
	}

	return links
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Capstone DB homepage")
	fmt.Printf("Accessed homepage")
}

func ViewLinks(w http.ResponseWriter, r *http.Request) {
	links := retrieveData()

	fmt.Println("Accessed View")
	json.NewEncoder(w).Encode(links)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/links", ViewLinks)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
