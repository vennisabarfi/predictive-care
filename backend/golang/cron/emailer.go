package cron

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"user_auth/models"
	"user_auth/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App struct to hold dependencies
var db *gorm.DB

func FindUsers(c *gin.Context) {

	// Establish PostgreSQL connection
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	//Refactor this to call on storage
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode),
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Connected to Database!.")

	// call user struct from models
	type User = models.User

	// initialize slice of emails
	var emails []string

	// retrieve emails from database and append into an array (slice)
	result := db.Model(&User{}).Pluck("email", &emails)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	c.JSON(http.StatusOK, emails)

}

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
