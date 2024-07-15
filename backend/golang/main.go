package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"
	"user_auth/cron"
	"user_auth/handlers"
	"user_auth/models"
	"user_auth/storage"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// App struct to hold dependencies

var DB *gorm.DB

// View all proverbs in database
// "/viewproverbs"
func ViewProverbs(c *gin.Context) {

	var proverb []handlers.Proverb

	err := DB.Find(&proverb)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		println(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello World",
		"result":  proverb,
	})
}

// add tests for when no more proverbs
// View Specific Proverb by ID ("/viewproverbs/:id")
func ViewProverb(c *gin.Context) {
	var proverb []handlers.Proverb
	id := c.Param("id")

	err := DB.First(&proverb, id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		println(err)
	}
	//update to just return text section of data
	c.IndentedJSON(http.StatusOK, gin.H{
		"result": proverb,
	})
}

// Sign up handler
func register(c *gin.Context) {
	//Get email and password from req body

	var body struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if (c.ShouldBindJSON(&body)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
		})
		return
	}

	//Hash password
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read input data",
		})
		return
	}

	user := models.User{Username: body.Username, Email: &body.Email, Password: string(hashedpassword)}

	// Log the received user data
	log.Printf("Received user data: %+v", user)

	// Check if the db variable is not nil
	if DB == nil {
		log.Println("Database connection is not initialized")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}
	result := DB.Create(&user)

	// Log successful user creation
	log.Printf("User created: %+v", user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create the user",
		})
		return
	}

	// send a response with user name
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user.Username,
	})

}

/*
login handler. Parses a form and checks for specific data
*/
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

	var user models.User
	result := DB.Where("email =?", body.Email).First(&user)

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

	/*
		Token created with 30-day expiration date and signed with secret key.
		JWT token set as cookie with secure flag and sent over HTTPS
	*/
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

	//Set Cookie ()
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)

	//send it as a response
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"token":   tokenString,
	})

}

// logout handler
func logout(c *gin.Context) {
	//Clear cookie
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "successfully logged user out",
	})

}

// func (app *App) FindUsers(c *gin.Context) {
// 	// Retrieve users from database
// 	var users []models.User

// 	result := app.DB.Find(&users)

// 	c.IndentedJSON(http.StatusOK, gin.H{
// 		"result": result,
// 	})

// }

func main() {

	r := gin.New()
	// port := os.Getenv("PORT")

	// set up security headers
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	} else {
		log.Println(".env file loaded successfully.")
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Error connecting to database", err)
	} else {
		log.Println(("Database connection established"))
	}

	err = models.MigrateUser(db)
	if err != nil {
		log.Fatal("User Database could not be migrated", err)
	} else {
		log.Println("User Database migrated successfully.")
	}

	handlers.InsertProverb()

	cron.SendMail()

	// cron.FirstCron()

	// app := &App{DB: db}

	//test

	//login user
	r.POST("/login", login)
	//register user
	r.POST("/register", register)
	//logout user
	r.GET("/logout", logout)
	//view all proverbs
	r.GET("/viewproverbs", ViewProverbs)
	//view specific proverb
	r.GET("/viewproverbs/:id", ViewProverb)

	r.Run(":" + os.Getenv("PORT"))

}
