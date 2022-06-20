package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/data"
	"realworld-go-kratos/internal/data/mysql_model"
)

type TagUsecase struct {
	TagRepo data.TagRepositoryIface
	log     *log.Helper
}

func NewTagUsecase(repo data.TagRepositoryIface, logger log.Logger) *TagUsecase {
	return &TagUsecase{TagRepo: repo, log: log.NewHelper(logger)}
}

func (uc *TagUsecase) CreateRealWorld(ctx context.Context, g *mysql_model.Tag) (*mysql_model.Tag, error) {
	return uc.TagRepo.Create(ctx, g)
}
