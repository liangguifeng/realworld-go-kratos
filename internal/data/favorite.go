package data

import (
	"context"
	"realworld-go-kratos/internal/data/mysql_model"
)

type FavoriteRepositoryIface interface {
	Create(context.Context, *mysql_model.Favorite) (*mysql_model.Favorite, error)
	Update(context.Context, *mysql_model.Favorite) (*mysql_model.Favorite, error)
	FindById(context.Context, int64) (*mysql_model.Favorite, error)
	DeleteById(context.Context, int64) (*mysql_model.Favorite, error)
	All(context.Context) ([]*mysql_model.Favorite, error)
}
