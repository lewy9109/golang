package grpc

import (
	context "context"
	"edu/pkg/user"

	grpc "google.golang.org/grpc"
)

type userGRPCServer struct {
	userService user.UserServiceInterface
}

func DefaultGrpcUserService(service user.UserServiceInterface) UserServiceClient {
	return &userGRPCServer{
		userService: service,
	}
}

func (u *userGRPCServer) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {

	return nil, nil
}

func (u *userGRPCServer) LoginUser(ctx context.Context, in *LoginUserResponse, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	return nil, nil
}

func (u *userGRPCServer) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	return nil, nil
}
