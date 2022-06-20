package impl

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/data"
	"realworld-go-kratos/internal/data/mysql_model"
)

type FavoriteRepository struct {
	data *data.Data
	log  *log.Helper
}

func NewFavoriteRepository(data *data.Data, logger log.Logger) data.FavoriteRepositoryIface {
	return &FavoriteRepository{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 保存
func (r *FavoriteRepository) Create(ctx context.Context, g *mysql_model.Favorite) (*mysql_model.Favorite, error) {
	return g, nil
}

// Update 更新
func (r *FavoriteRepository) Update(ctx context.Context, g *mysql_model.Favorite) (*mysql_model.Favorite, error) {
	return g, nil
}

// FindById 通过id查找
func (r *FavoriteRepository) FindById(context.Context, int64) (*mysql_model.Favorite, error) {
	return nil, nil
}

// DeleteById 通过id删除
func (r *FavoriteRepository) DeleteById(context.Context, int64) (*mysql_model.Favorite, error) {
	return nil, nil
}

// All 获取所有
func (r *FavoriteRepository) All(context.Context) ([]*mysql_model.Favorite, error) {
	return nil, nil
}
