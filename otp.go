package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

func main() {
    // Open a database connection
    db, err := sql.Open("postgres", "user=your_username password=your_password dbname=your_database sslmode=disable")
    if err != nil {
        fmt.Printf("Failed to connect to database: %v\n", err)
        return
    }
    defer db.Close()

    // Prepare the SQL statement
    stmt := "SELECT (SELECT COUNT(*) FROM table1) + (SELECT COUNT(*) FROM table2) AS total"

    // Execute the query and get the result
    var count int
    if err := db.QueryRow(stmt).Scan(&count); err != nil {
        fmt.Printf("Failed to query count: %v\n", err)
        return
    }

    // Print the result
    fmt.Printf("Total users: %d\n", count)
}
