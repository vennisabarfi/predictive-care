package handlers

import (
	"user_auth/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// App struct to hold dependencies
type App struct {
	DB *gorm.DB
}

// view artist in database
func viewEvent(c *gin.Context) {
	var artist []Models.Artist
	err := DB.Where("name = ?", artist).Find(&models.Artist)

	if err != nil {

	}
}
