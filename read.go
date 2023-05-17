package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    // Open the CSV file
    file, err := os.Open("datas.csv")
   if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()


    // Create a new CSV reader
    reader := csv.NewReader(file)

    // Read in all the CSV records
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Loop through the records and print them out
    for _, record := range records {
        fmt.Println(record)
    }
}
