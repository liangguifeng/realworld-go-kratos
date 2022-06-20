package impl

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/data"
	"realworld-go-kratos/internal/data/mysql_model"
)

type UserRepository struct {
	data *data.Data
	log  *log.Helper
}

func NewUserRepository(data *data.Data, logger log.Logger) data.UserRepositoryIface {
	return &UserRepository{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 保存
func (r *UserRepository) Create(ctx context.Context, g *mysql_model.User) (*mysql_model.User, error) {
	return g, nil
}

// Update 更新
func (r *UserRepository) Update(ctx context.Context, g *mysql_model.User) (*mysql_model.User, error) {
	return g, nil
}

// FindById 通过id查找
func (r *UserRepository) FindById(context.Context, int64) (*mysql_model.User, error) {
	return nil, nil
}

// DeleteById 通过id删除
func (r *UserRepository) DeleteById(context.Context, int64) (*mysql_model.User, error) {
	return nil, nil
}

// All 获取所有
func (r *UserRepository) All(context.Context) ([]*mysql_model.User, error) {
	return nil, nil
}
