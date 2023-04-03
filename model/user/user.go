package user

import (
// "gorm.io/gorm"
)

type User struct {
	// gorm.Model
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	Email       string `gorm:"email"`
	Password    string `gorm:"password"`
	AccessToken string `gorm:"access_token"`
}
