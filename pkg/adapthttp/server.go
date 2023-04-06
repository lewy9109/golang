package adapthttp

import (
	userService "edu/pkg/user"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserServerInterface interface {
	CreateUser(context *gin.Context)
	LoginUser(context *gin.Context)
	Authorize(context *gin.Context)
	GetInfo(context *gin.Context)
}

func DefalutUserServer(service userService.UserServiceInterface) UserServerInterface {
	return &userServer{
		service: service,
	}
}

type userServer struct {
	service userService.UserServiceInterface
}

func (u *userServer) CreateUser(context *gin.Context) {
	user := CreateUserRequest{}
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, CreateUserResponse{})
		return
	}

	createUser := userService.User{
		Email:     user.Email,
		FirstName: user.Name,
		LastName:  user.LastName,
		Password:  user.Password,
	}

	err = u.service.CreateUser(createUser)
	if err != nil {
		if errors.Is(err, userService.ErrInternalDBError) || errors.Is(err, userService.ErrInternalServer) {
			context.JSON(http.StatusInternalServerError, CreateUserResponse{})
			return
		}
		context.JSON(http.StatusBadRequest, CreateUserResponse{})
		return
	}

	context.JSON(http.StatusCreated, user)
}
func (u *userServer) LoginUser(context *gin.Context) {
	loginUser := LoginRequest{}
	err := context.BindJSON(&loginUser)
	if err != nil {
		context.JSON(http.StatusUnauthorized, ErrorOccuredModel{Message: err.Error()})
		return
	}

	token, err := u.service.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, LoginResponse{})
	}

	context.JSON(http.StatusOK, LoginResponse{AccessToken: token})
}

func (u *userServer) Authorize(context *gin.Context) {
	token := context.GetHeader("authorized")
	if token == "" {
		context.JSON(http.StatusUnauthorized, nil)
		context.Abort()
		return
	}
	userID, err := u.service.Authorize(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, ErrorOccuredModel{Message: err.Error()})
		context.Abort()
		return
	}
	context.Request.Header.Add("user_id", strconv.Itoa(int(userID)))
	context.Next()
}

func (u *userServer) GetInfo(context *gin.Context) {
	userIdString := context.GetHeader("user_id")
	if userIdString == "" {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	user, err := u.service.GetUserInfo(uint(userId))
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	response := GetUserResponse{
		Name:     user.FirstName,
		LastName: user.LastName,
		Email:    user.Email,
	}

	context.JSON(http.StatusOK, response)
}
