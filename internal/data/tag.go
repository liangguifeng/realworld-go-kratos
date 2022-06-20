package data

import (
	"context"
	"realworld-go-kratos/internal/data/mysql_model"
)

type TagRepositoryIface interface {
	Create(context.Context, *mysql_model.Tag) (*mysql_model.Tag, error)
	Update(context.Context, *mysql_model.Tag) (*mysql_model.Tag, error)
	FindById(context.Context, int64) (*mysql_model.Tag, error)
	DeleteById(context.Context, int64) (*mysql_model.Tag, error)
	All(context.Context) ([]*mysql_model.Tag, error)
}
