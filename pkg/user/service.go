package user

import "edu/pkg/helper"

type UserService interface {
	CreateUser(user User) error
	Login(email string, password string) (string, error)
	GetUserInfo(id uint) (User, error)
}

type userSercive struct {
	infra     UserInfrastructure
	jwtSecret string
}

func DefalutUserService(userInfrastructure UserInfrastructure, jwtSecret string) UserService {
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

	err := u.infra.CreateUser(user)

	if err != nil {
		return ErrInternalDBError
	}

	return nil
}

func (u *userSercive) Login(email string, password string) (string, error) {

	if email == "" {
		return "", ErrEmialIsEmpty
	}

	if password == "" {
		return "", ErrPasswordIsEmpty
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

	return token, nil
}

func (u *userSercive) GetUserInfo(id uint) (User, error) {
	return User{}, nil
}
