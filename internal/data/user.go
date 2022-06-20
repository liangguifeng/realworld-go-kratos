package data

import (
	"context"
	"realworld-go-kratos/internal/data/mysql_model"
)

type UserRepositoryIface interface {
	Create(context.Context, *mysql_model.User) (*mysql_model.User, error)
	Update(context.Context, *mysql_model.User) (*mysql_model.User, error)
	FindById(context.Context, int64) (*mysql_model.User, error)
	DeleteById(context.Context, int64) (*mysql_model.User, error)
	All(context.Context) ([]*mysql_model.User, error)
}
