package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"go/main.go/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// Importing your models package, adjust based on your actual module path
)

// loadEnvVariable loads environment variables from .env file
func loadEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return os.Getenv(key)
}

func main() {
	var wg sync.WaitGroup
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Predictive Care API"})
	})

	port := ":" + loadEnvVariable("PORT")

	go func() {
		if err := r.Run(port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	if err := models.SetupDatabase(); err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	log.Println("Database setup and migration successful!")
	//keep program running after connecting to db.
	wg.Wait()
}
