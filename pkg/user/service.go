package user

import "log"

type UserService interface {
	CreateUser(user User) error
	Login(email string, password string) (string, error)
	GetUserInfo(id uint) (User, error)
}

type userSercive struct {
	infra UserInfrastructure
}

func DefalutUserService(userInfrastructure UserInfrastructure) UserService {
	return &userSercive{
		infra: userInfrastructure,
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
	return "", nil
}

func (u *userSercive) GetUserInfo(id uint) (User, error) {
	return User{}, nil
}
