package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
}

type ArticleRepositoryIface interface {
	Create(context.Context, *Article) (*Article, error)
	Update(context.Context, *Article) (*Article, error)
	FindById(context.Context, int64) (*Article, error)
	DeleteById(context.Context, int64) (*Article, error)
	All(context.Context) ([]*Article, error)
}

type ArticleUsecase struct {
	ArticleRepo ArticleRepositoryIface
	log         *log.Helper
}

func NewArticleUsecase(repo ArticleRepositoryIface, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{ArticleRepo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUsecase) CreateRealWorld(ctx context.Context, g *Article) (*Article, error) {
	return uc.ArticleRepo.Create(ctx, g)
}
