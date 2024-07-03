package main

import (
	"log"
	"net/http"
	"os"
	"user_auth/models"
	"user_auth/storage"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const userkey = "user"

var secret = []byte("secret")

// Cookie and Session Management
func engine() *gin.Engine {
	r := gin.New()

	//Setup cookie store for session management
	r.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))

	//Login and logout routes
	r.POST("/login", login)
	r.GET("/logout", logout)

	return r
}

// Middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		//Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unathorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

// login handler. Parses a form and checks for specific data
func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

}

func logout(c *gin.Context) {

}

func main() {

	r := engine()
	r.Use(gin.Logger())
	// port := os.Getenv("PORT")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
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

	//GET Request
	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Predictive Care API"})
	})

	r.Run(":" + os.Getenv("PORT"))

}
