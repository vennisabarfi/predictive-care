package cron

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	"user_auth/handlers"
	"user_auth/models"
	"user_auth/storage"

	"github.com/go-gomail/gomail"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// generates a random number between one and the given number
func RandomId(num int64) int64 {
	rand.NewSource(time.Now().UnixNano())
	var min int64
	var result int64
	min = 1
	max := num

	result = rand.Int63n(max-min+1) + min //returns int64

	return result

}

func SendMail() {

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

	//choose random proverb from database
	var proverb []handlers.Proverb
	var count int64

	// find the total number of rows in database
	db.Table("proverbs").Count(&count)
	// db.Raw("SELECT COUNT(*) FROM public.proverbs").Scan(&count)
	fmt.Println(count)

	id := RandomId(count)

	var randomProverb string
	db.Raw("SELECT text FROM public.proverbs where id = $1", id).Scan(&randomProverb)

	fmt.Println(randomProverb)

	// call user struct from models
	type User = models.User

	// initialize slice of emails and usernames
	var emails []string
	var usernames []string

	// retrieve emails and usernames from db
	res := db.Model(&User{}).Pluck("username", &usernames)

	if res.Error != nil {
		fmt.Println(res.Error)
	}

	// retrieve emails from database and append into an array (slice)
	result := db.Model(&User{}).Pluck("email", &emails)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	// Create a userlist struct to store usernames and emails
	type UserList struct {
		Name    string //usernames
		Address string //emails
	}

	// initialize email and username lists
	var EmailList []string
	var UsernameList []string

	// append emails and usernames to lists
	EmailList = append(EmailList, emails...)
	UsernameList = append(UsernameList, usernames...)

	fmt.Println(EmailList, UsernameList)

	// initialize newsletter list to append combine EmailList and UsernameList
	var NewsLetter []UserList

	for i := range EmailList {
		NewsLetter = append(NewsLetter, UserList{
			Name:    UsernameList[i],
			Address: EmailList[i],
		})
	}

	for _, user := range NewsLetter {
		fmt.Printf(user.Name)
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
	for _, r := range NewsLetter {
		m.SetHeader("From", "no-reply@example.com") //set company email here
		m.SetAddressHeader("To", r.Address, r.Name)
		m.SetHeader("Subject", "Proverb of the Day")
		m.SetBody("text/html", fmt.Sprintf("Hello %s!"+r.Name+"Here's your proverb of the day!", proverb))

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Address, err)
		}
		m.Reset()
	}
}
