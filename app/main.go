package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ", r.URL.Path)

	db, err := sql.Open("postgres", "postgres://postgres/gotest?port=5432&user=postgres&sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var data string
	db.QueryRow("SELECT 'hello from postgres'").Scan(&data)

	fmt.Fprintf(w, "%v", data)
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
