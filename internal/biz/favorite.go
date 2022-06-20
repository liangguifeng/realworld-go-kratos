package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Favorite struct {
}

type FavoriteRepositoryIface interface {
	Create(context.Context, *Favorite) (*Favorite, error)
	Update(context.Context, *Favorite) (*Favorite, error)
	FindById(context.Context, int64) (*Favorite, error)
	DeleteById(context.Context, int64) (*Favorite, error)
	All(context.Context) ([]*Favorite, error)
}

type FavoriteUsecase struct {
	FavoriteRepo FavoriteRepositoryIface
	log          *log.Helper
}

func NewFavoriteUsecase(repo FavoriteRepositoryIface, logger log.Logger) *FavoriteUsecase {
	return &FavoriteUsecase{FavoriteRepo: repo, log: log.NewHelper(logger)}
}

func (uc *FavoriteUsecase) CreateRealWorld(ctx context.Context, g *Favorite) (*Favorite, error) {
	return uc.FavoriteRepo.Create(ctx, g)
}
