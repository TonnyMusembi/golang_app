package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Connect to the database
    db, err := sql.Open("mysql", "user:password@tcp(hostname:port)/database_name")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer db.Close()

    // Test the connection
    err = db.Ping()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Connected to the database.")

    // Perform database operations here
}
