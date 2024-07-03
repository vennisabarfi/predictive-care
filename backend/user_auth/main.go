package main

import (
	"log"
	"os"
	"user_auth/models"
	"user_auth/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	port := ":" + os.Getenv("PORT")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	err = models.MigrateUser(db)
	if err != nil {
		log.Fatal("User Database could not be migrated", err)
	}

	err = r.Run(port)

	if err != nil {
		log.Fatal("Server is down!", err)
	}
}
