package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv" //import godotenv

	"os"
)

// load environment variables
func env_Variable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		port := env_Variable("PORT")
		//check for errors
		error := r.Run(port)

		if error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve tasks.",
			})
			return
		}
	})

}
