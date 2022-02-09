package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Record struct {
	ID          int
	LastName    string
	FirstName   string
	Company_Add string
	Country     string
	Position    string
}

type Database struct {
	nextID int
	mu     sync.Mutex
	recs   []Record
}

func main() {
	db := &Database{recs: []Record{}}
	http.ListenAndServe(":8080", db.handler())
}

func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		if r.URL.Path == "/contacts" {
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/contacts/%d", &id); n == 1 {
			db.processID(id, w, r)
		}
	}
}

func (db *Database) process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var rec Record
		if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.mu.Lock()
		rec.ID = db.nextID
		db.nextID++
		db.recs = append(db.recs, rec)
		db.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{\"status\":successful}")
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(db.recs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (db *Database) processID(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		exists := false
		db.mu.Lock()
		for j, item := range db.recs {
			if id == item.ID {
				db.recs = append(db.recs[:j], db.recs[j+1:]...)
				exists = true
				break
			}
		}
		db.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		if exists {
			fmt.Fprintln(w, "{\"status\":successful}")
		} else {
			fmt.Fprintln(w, "{\"status\":unsuccessful}")
		}
	case "PUT":
		exists := false
		db.mu.Lock()
		for x, item := range db.recs {
			if id == item.ID {
				db.recs = append(db.recs[:x], db.recs[x+1:]...)
				exists = true
				var rec Record
				if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				db.recs = append(db.recs, rec)
				break
			}
		}
		db.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		if exists {
			fmt.Fprintln(w, "{\"status\":successful}")
		} else {
			fmt.Fprintln(w, "{\"status\":unsuccessful}")
		}
	}
}
