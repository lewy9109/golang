package user

import (
	"errors"
	"gorm.io/gorm"
)

type UserInfrastructure interface {
	CreateUser(user User) error
	FindUser(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByToken(accessToken string) (*User, error)
	UpdateUserAccessToken(id uint, access_token string) error
}

type userInfra struct {
	db *gorm.DB
}

func DefaultUserInfraStructure(db *gorm.DB) UserInfrastructure {

	return &userInfra{
		db,
	}
}

func (u *userInfra) CreateUser(user User) error {

	result := u.db.Create(&user)

	if result.Error != nil {
		return result.Error

	}

	return nil
}

func (u *userInfra) FindUser(id uint) (*User, error) {

	user := User{}

	result := u.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error

	}
	return &user, nil
}

func (u *userInfra) GetByEmail(email string) (*User, error) {
	user := User{}

	result := u.db.First(&user, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error

	}
	return &user, nil
}

func (u *userInfra) GetByToken(accessToken string) (*User, error) {
	user := User{}

	result := u.db.First(&user, "access_token = ?", accessToken)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error

	}
	return &user, nil
}

func (u *userInfra) UpdateUserAccessToken(id uint, access_token string) error {

	result := u.db.Model(&User{}).Where("id = ?", id).Update("access_token", access_token)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error

	}
	return nil
}
