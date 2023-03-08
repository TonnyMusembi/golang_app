package main

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Get the database credentials from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Set up the database connection
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Test the database connection
    err = db.Ping()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Connected to MySQL database!")

    // Close the database connection when finished
    defer db.Close()
}
