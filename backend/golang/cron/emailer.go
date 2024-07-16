package cron

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
	"user_auth/models"
	"user_auth/storage"

	"github.com/go-gomail/gomail"
)

var mu sync.Mutex

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

// choose random proverb from database
func RandomProverb() string {

	db, err := storage.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	var count int64
	var randomProverb string

	// find the total number of rows in database
	db.Table("proverbs").Count(&count)
	// db.Raw("SELECT COUNT(*) FROM public.proverbs").Scan(&count)
	fmt.Println(count)

	id := RandomId(count)

	db.Raw("SELECT text FROM public.proverbs where id = $1", id).Scan(&randomProverb)
	print("Random proverb selected")

	return randomProverb
}

func SendMail() {
	//using mutex to avoid race conditions
	// only one goroutine can access shared resource
	mu.Lock()
	defer mu.Unlock()

	db, err := storage.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// call user struct from models
	type User = models.User

	// initialize slice of emails and usernames
	var emails []string
	var usernames []string

	// retrieve usernames from db
	res := db.Model(&User{}).Pluck("username", &usernames)

	if res.Error != nil {
		fmt.Println(res.Error)
	}

	// retrieve emails from database
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

	// initialize newsletter list to combine EmailList and UsernameList
	var NewsLetter []UserList

	// append usernamelist and emaillist to newsletter
	for i := range EmailList {
		NewsLetter = append(NewsLetter, UserList{
			Name:    UsernameList[i],
			Address: EmailList[i],
		})
	}

	// print out usernames from newsletter(to test)
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

	// work on refactoring and including text.html
	m := gomail.NewMessage()
	for _, r := range NewsLetter {
		m.SetHeader("From", "no-reply@example.com") //set sender email here
		m.SetAddressHeader("To", r.Address, r.Name)
		m.SetHeader("Subject", "Proverb of the Day")
		m.SetBody("text/html", fmt.Sprintf("Hello %s!"+r.Name+"Here's your proverb of the day!", RandomProverb()))

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Address, err)
		}
		m.Reset()
	}
}
