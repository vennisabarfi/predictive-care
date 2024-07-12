package cron

import (
	"fmt"
	"log"

	"github.com/go-gomail/gomail"
)

func SendEmail() {
	// List of recipients

	var list []struct {
		Name    string //username
		Address string //email address
	}

	// Using MailHog (SMTP server on port 1025)
	d := gomail.NewDialer("localhost", 3000, "", "")
	s, err := d.Dial()
	if err != nil {
		panic(err)
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
