package userController

type CreateUserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserResponse struct{}

type LoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type GetUserResponse struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

type ErrorOccuredModel struct {
	Message string `json:"message"`
}
