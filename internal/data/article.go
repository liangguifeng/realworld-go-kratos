package data

import (
	"context"
	"realworld-go-kratos/internal/data/mysql_model"
)

type ArticleRepositoryIface interface {
	Create(context.Context, *mysql_model.Article) (*mysql_model.Article, error)
	Update(context.Context, *mysql_model.Article) (*mysql_model.Article, error)
	FindById(context.Context, int64) (*mysql_model.Article, error)
	DeleteById(context.Context, int64) (*mysql_model.Article, error)
	All(context.Context) ([]*mysql_model.Article, error)
}
