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
	ErrInternalServer           = errors.New("internal-server-error")
	ErrTokenCreate              = errors.New("token-create-error")
	ErrUserEmailIsExist         = errors.New("email-exist-system")
	ErrUserUnAuthorized         = errors.New("user-unauthorized")
)

type User struct {
	gorm.Model
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	Email       string `gorm:"email"`
	Password    string `gorm:"password"`
	AccessToken string `gorm:"access_token"`
}

type UserInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
