package user

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrEmialIsEmpty             = errors.New("email-is-empty")
	ErrFirstNameIsEmpty         = errors.New("first-name-is-empty")
	ErrLastNameIsEmpty          = errors.New("last-name-is-empty")
	ErrPasswordIsEmpty          = errors.New("password-is-empty")
	ErrPasswordOrEmailIsInvalid = errors.New("password-emial-is-invalid")
	ErrInternalDBError          = errors.New("internal-db-error")
)

type User struct {
	gorm.Model
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	Email       string `gorm:"email"`
	Password    string `gorm:"password"`
	AccessToken string `gorm:"access_token"`
}
