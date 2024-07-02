package main

import (
	"log"
	"net/http"

	"go/main.go/models"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv" //import godotenv

	"os"
)

// load environment variables
func env_Variable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	return os.Getenv(key)
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Predictive Care API"})

	})

	port := ":" + env_Variable("PORT")
	//check for errors
	error := r.Run(port)

	if error != nil {
		log.Fatalf("Failed to start server: %v", error)
	}

	// test database.go
	// err := models.ConnectDB()
	// if err != nil {
	// 	log.Fatalf("Error connecting to database", err)
	// 	return
	// }

	err := models.MigrateDatabase()
	if err != nil {
		log.Fatalf("Error migrating database", err)
	}
}
