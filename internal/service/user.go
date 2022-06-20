package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "realworld-go-kratos/api/realworld/v1"
	"realworld-go-kratos/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

// Login
func (s UserService) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

// Register
func (UserService) Register(ctx context.Context, r *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

// GetCurrentUser
func (UserService) GetCurrentUser(ctx context.Context, r *v1.GetCurrentUserRequest) (*v1.GetCurrentUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}

// UpdateUser
func (UserService) UpdateUser(ctx context.Context, r *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func (UserService) GetProfile(ctx context.Context, r *v1.GetProfileRequest) (*v1.GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}

func (UserService) FollowUser(ctx context.Context, r *v1.FollowUserRequest) (*v1.FollowUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}

func (UserService) UnfollowUser(ctx context.Context, r *v1.UnfollowUserRequest) (*v1.UnfollowUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUser not implemented")
}
