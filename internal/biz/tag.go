package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Tag struct {
}

type TagRepositoryIface interface {
	Create(context.Context, *Tag) (*Tag, error)
	Update(context.Context, *Tag) (*Tag, error)
	FindById(context.Context, int64) (*Tag, error)
	DeleteById(context.Context, int64) (*Tag, error)
	All(context.Context) ([]*Tag, error)
}

type TagUsecase struct {
	TagRepo TagRepositoryIface
	log     *log.Helper
}

func NewTagUsecase(repo TagRepositoryIface, logger log.Logger) *TagUsecase {
	return &TagUsecase{TagRepo: repo, log: log.NewHelper(logger)}
}

func (uc *TagUsecase) CreateRealWorld(ctx context.Context, g *Tag) (*Tag, error) {
	return uc.TagRepo.Create(ctx, g)
}
