package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "realworld-go-kratos/api/realworld/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type User struct {
}

type UserRepositoryIface interface {
	Create(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindById(context.Context, int64) (*User, error)
	DeleteById(context.Context, int64) (*User, error)
	All(context.Context) ([]*User, error)
}

type RealWorldUsecase struct {
	userRepo UserRepositoryIface
	log  *log.Helper
}

func NewRealWorldUsecase(repo UserRepositoryIface, logger log.Logger) *RealWorldUsecase {
	return &RealWorldUsecase{userRepo: repo, log: log.NewHelper(logger)}
}

func (uc *RealWorldUsecase) CreateRealWorld(ctx context.Context, g *User) (*User, error) {
	return uc.userRepo.Create(ctx, g)
}
