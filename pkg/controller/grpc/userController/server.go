package grpc

import (
	context "context"
	"edu/pkg/user"

	grpc "google.golang.org/grpc"
)

type userGRPCServer struct {
	userService user.UserService
}

func DefaultGrpcUserService(userService user.UserService) UserService {
	// return &userGRPCServer{
	// 	userService: userService
	// }
}

func (u *userGRPCServer) CreateUser(ctx context.Context, in *CreateUserRequest, opts grpc.CallOption) (*CreateUserResponse, error) {

	return nil, nil
}

func (u *userGRPCServer) LoginUser(ctx context.Context, in *CreateUserRequest, opts grpc.CallOption) (*CreateUserResponse, error) {
	return nil, nil
}

func (u *userGRPCServer) GetUserInfo(ctx context.Context, in *CreateUserRequest, opts grpc.CallOption) (*CreateUserResponse, error) {
	return nil, nil
}
