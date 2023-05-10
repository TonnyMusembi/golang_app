package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connection string for the PostgreSQL database
	connStr := "host=localhost port=5433 user=postgres password=yourpassword dbname=postgres1 sslmode=disable"
	// Connect to the database
	    // log.SetOutput(connStr);

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the PostgreSQL database!")
}
