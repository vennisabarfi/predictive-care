package cron

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
	"github.com/joho/godotenv"
)

func SendEmail() {
	// List of recipients

	var list []struct {
		Name    string //username
		Address string //email address
	}

	// Using MailHog (SMTP server on port 1025). Docker Run!
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// convert port number to int
	smtp_port := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(smtp_port) // port number
	if err != nil {
		panic(err)
	}

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, "", "")
	s, err := d.Dial()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("SMTP Server is up and running")
	}

	m := gomail.NewMessage()
	for _, r := range list {
		m.SetHeader("From", "no-reply@example.com")
		m.SetAddressHeader("To", r.Address, r.Name)
		m.SetHeader("Subject", "Newsletter #1")
		m.SetBody("text/html", fmt.Sprintf("Hello %s!", r.Name))

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Address, err)
		}
		m.Reset()
	}
}
