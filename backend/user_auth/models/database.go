package models

import (
	"time"

	"gorm.io/gorm"
)

// Define User model
type User struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	Username  string  `json:"username" gorm:"unique"`
	Email     *string `json:"email" gorm:"unique"`
	Password  string  `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
