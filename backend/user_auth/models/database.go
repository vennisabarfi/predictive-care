package models

import (
	"time"

	"gorm.io/gorm"
)

// User can have many artists to track
// Define User model
type User struct {
	gorm.Model
	ID        uint    `json:"id" gorm:"primary_key"`
	Username  string  `json:"username" gorm:"unique"`
	Email     *string `json:"email" gorm:"unique"`
	Password  string  `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Define Artist Model
type Artist struct {
	gorm.Model
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"` //event_location
	MaxPrice float32 `json:"min_price"`
	MinPrice float32 `json:"max_price"`
	UserID   uint    `gorm: "index"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}

func MigrateArtist(db *gorm.DB) error {
	err := db.AutoMigrate(&Artist{})
	return err
}
