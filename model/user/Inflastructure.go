package user

type UserInfrastructure interface {
	CreateUser(user User) error
	GetUser(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByToken(accessToken string) (*User, error)
}

type userInfra struct {
}

func DefaultUserInfraStructure() UserInfrastructure {

	return &userInfra{}
}

func (u *userInfra) CreateUser(user User) error {
	return nil
}

func (u *userInfra) GetUser(id uint) (*User, error) {
	return nil, nil
}

func (u *userInfra) GetByEmail(email string) (*User, error) {
	return nil, nil
}

func (u *userInfra) GetByToken(accessToken string) (*User, error) {

	return nil, nil
}
