package data

import (
	"context"
	"realworld-go-kratos/internal/data/mysql_model"
)

type CommentRepositoryIface interface {
	Create(context.Context, *mysql_model.Comment) (*mysql_model.Comment, error)
	Update(context.Context, *mysql_model.Comment) (*mysql_model.Comment, error)
	FindById(context.Context, int64) (*mysql_model.Comment, error)
	DeleteById(context.Context, int64) (*mysql_model.Comment, error)
	All(context.Context) ([]*mysql_model.Comment, error)
}
