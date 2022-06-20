package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/data"
	"realworld-go-kratos/internal/data/mysql_model"
)

type UserUsecase struct {
	userRepo data.UserRepositoryIface
	log      *log.Helper
}

func NewUserUsecase(repo data.UserRepositoryIface, logger log.Logger) *UserUsecase {
	return &UserUsecase{userRepo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateRealWorld(ctx context.Context, g *mysql_model.User) (*mysql_model.User, error) {
	return uc.userRepo.Create(ctx, g)
}
