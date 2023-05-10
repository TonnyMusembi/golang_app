package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Establish a database connection
	db, err := sql.Open("postgres", "host=localhost port=5432 user=my-postgres password=password dbname=yourdbname sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Perform a SELECT query
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
