package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/data"
	"realworld-go-kratos/internal/data/mysql_model"
)

type ArticleUsecase struct {
	ArticleRepo data.ArticleRepositoryIface
	log         *log.Helper
}

func NewArticleUsecase(repo data.ArticleRepositoryIface, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{ArticleRepo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUsecase) CreateRealWorld(ctx context.Context, g *mysql_model.Article) (*mysql_model.Article, error) {
	return uc.ArticleRepo.Create(ctx, g)
}
