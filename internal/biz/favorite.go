package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/data"
	"realworld-go-kratos/internal/data/mysql_model"
)

type FavoriteUsecase struct {
	FavoriteRepo data.FavoriteRepositoryIface
	log          *log.Helper
}

func NewFavoriteUsecase(repo data.FavoriteRepositoryIface, logger log.Logger) *FavoriteUsecase {
	return &FavoriteUsecase{FavoriteRepo: repo, log: log.NewHelper(logger)}
}

func (uc *FavoriteUsecase) CreateRealWorld(ctx context.Context, g *mysql_model.Favorite) (*mysql_model.Favorite, error) {
	return uc.FavoriteRepo.Create(ctx, g)
}
