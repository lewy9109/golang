package user

import (
	"edu/pkg/constraints"
	"edu/pkg/helper"
)

type UserServiceInterface interface {
	CreateUser(user User) error
	Login(email string, password string) (string, error)
	GetUserInfo(id uint) (UserInfo, error)
	Authorize(accessToken string) (uint, error)
}

type userSercive struct {
	infra     UserInfrastructure
	jwtSecret string
}

func DefalutUserService(userInfrastructure UserInfrastructure, jwtSecret string) UserServiceInterface {
	return &userSercive{
		infra:     userInfrastructure,
		jwtSecret: jwtSecret,
	}
}

func (u *userSercive) CreateUser(user User) error {

	if user.Email == "" {
		return ErrEmialIsEmpty
	}

	if user.FirstName == "" {
		return ErrFirstNameIsEmpty
	}

	if user.LastName == "" {
		return ErrLastNameIsEmpty
	}

	if user.Password == "" {
		return ErrPasswordIsEmpty
	}

	if u.checkIsEmailExist(user.Email) != nil {
		return ErrUserEmailIsExist
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = u.infra.CreateUser(user)

	if err != nil {
		return ErrInternalDBError
	}

	return nil
}

func (u *userSercive) Login(email string, password string) (string, error) {

	err := validateLoginCredentials(email, password)
	if err != nil {
		return "", err
	}

	user, err := u.infra.GetByEmail(email)
	if err != nil {
		return "", ErrPasswordOrEmailIsInvalid
	}

	if !helper.ComaparePasswords(user.Password, password) {
		return "", ErrPasswordOrEmailIsInvalid
	}

	token, err := helper.CreateJWTToken(user.ID, u.jwtSecret)

	if err != nil {
		return "", ErrTokenCreate
	}

	err = u.infra.UpdateUserAccessToken(user.ID, token)
	if err != nil {
		return "", ErrInternalDBError
	}

	return token, nil
}

func (u *userSercive) GetUserInfo(id uint) (UserInfo, error) {
	user, err := u.infra.FindUser(id)
	if err != nil {
		return UserInfo{}, ErrInternalServer

	}
	user.Password = ""

	userInfo := UserInfo{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return userInfo, nil
}

func validateLoginCredentials(email string, password string) error {
	constraints := constraints.DefaultValidationStruct()

	if constraints.IsBlank(email) {
		return ErrEmialIsEmpty
	}

	if constraints.IsBlank(password) {
		return ErrPasswordIsEmpty
	}
	return nil
}

func (u *userSercive) checkIsEmailExist(email string) error {
	user, _ := u.infra.GetByEmail(email)
	if user != nil && user.Email == email {
		return ErrUserEmailIsExist
	}
	return nil
}
func (u *userSercive) Authorize(accessToken string) (uint, error) {
	_, err := helper.ValidateJWTToken(accessToken, u.jwtSecret)
	if err != nil {
		return 0, ErrUserUnAuthorized
	}

	user, err := u.infra.GetByToken(accessToken)
	if err != nil {
		return 0, ErrInternalDBError
	}

	return user.ID, nil
}
