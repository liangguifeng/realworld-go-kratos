package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/biz"
)

type RealWorldRepository struct {
	data *Data
	log  *log.Helper
}

func NewRealWorldRepository(data *Data, logger log.Logger) biz.UserRepositoryIface {
	return &RealWorldRepository{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 保存
func (r *RealWorldRepository) Create(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

// Update 更新
func (r *RealWorldRepository) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

// FindById 通过id查找
func (r *RealWorldRepository) FindById(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

// DeleteById 通过id删除
func (r *RealWorldRepository) DeleteById(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

// All 获取所有
func (r *RealWorldRepository) All(context.Context) ([]*biz.User, error) {
	return nil, nil
}
