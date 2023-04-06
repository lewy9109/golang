package adapthttp

import (
	userService "edu/pkg/user"
	"errors"
	"net/http"

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
		context.JSON(http.StatusBadRequest, nil)
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
			context.JSON(http.StatusInternalServerError, nil)
			return
		}
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	context.JSON(http.StatusOK, CreateUserResponse{})
}
func (u *userServer) LoginUser(context *gin.Context) {

}
func (u *userServer) Authorize(context *gin.Context) {

}
func (u *userServer) GetInfo(context *gin.Context) {

}
