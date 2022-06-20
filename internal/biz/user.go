package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

type UserUsecase struct {
	userRepo UserRepositoryIface
	log      *log.Helper
}

func NewUserUsecase(repo UserRepositoryIface, logger log.Logger) *UserUsecase {
	return &UserUsecase{userRepo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateRealWorld(ctx context.Context, g *User) (*User, error) {
	return uc.userRepo.Create(ctx, g)
}
