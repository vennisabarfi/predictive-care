package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"
	"user_auth/models"
	"user_auth/storage"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// const userkey = "user"

// var secret = []byte("secret")

// // Cookie and Session Management
// func engine() *gin.Engine {
// 	r := gin.New()

// 	//Setup cookie store for session management
// 	r.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))

// 	//Login and logout routes
// 	r.POST("/login", login)
// 	r.GET("/logout", logout)

// 	return r
// }

// // Middleware to check the session
// func AuthRequired(c *gin.Context) {
// 	session := sessions.Default(c)
// 	user := session.Get(userkey)
// 	if user == nil {
// 		//Abort the request with the appropriate error code
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unathorized"})
// 		return
// 	}
// 	// Continue down the chain to handler etc
// 	c.Next()
// }

// login handler. Parses a form and checks for specific data
func login(c *gin.Context) {
	// parse email and password from body
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	//Get user from database
	var db *gorm.DB
	var user models.User
	result := db.Where("email =?", body.Email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid email or password",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server error",
			})
		}
		return
	}
	//Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create token",
		})
		return
	}

	//Set Cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)
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
