package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"os"
	"strconv"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	// Open the CSV file and read the contents
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the list of users and send an OTP to each one
	for _, record := range records {
		// Extract the user's email address from the CSV record
		email := record[0]

		// Generate a random 6-digit code for the OTP
		rand.Seed(time.Now().UnixNano())
		code := strconv.Itoa(rand.Intn(900000) + 100000)

		// Send an email with the OTP to the user
		err := sendOTPEmail(email, code)
		if err != nil {
			log.Printf("Error sending OTP email to %s: %v\n", email, err)
		} else {
			log.Printf("OTP email sent to %s\n", email)
		}
	}
}

func sendOTPEmail(email string, code string) error {
	// Set up the email message
	from := mail.NewEmail("Sender Name", "sender@example.com")
	subject := "Your OTP Code"
	to := mail.NewEmail("", email)
	content := mail.NewContent("text/plain", fmt.Sprintf("Your OTP code is %s", code))
	message := mail.NewV3MailInit(from, subject, to, content)

	// Set up the SendGrid client and API request
	client := sendgrid.NewSendClient("<SENDGRID_API_KEY>")
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}
