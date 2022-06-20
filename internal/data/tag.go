package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/biz"
)

type TagRepository struct {
	data *Data
	log  *log.Helper
}

func NewTagRepository(data *Data, logger log.Logger) biz.TagRepositoryIface {
	return &TagRepository{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 保存
func (r *TagRepository) Create(ctx context.Context, g *biz.Tag) (*biz.Tag, error) {
	return g, nil
}

// Update 更新
func (r *TagRepository) Update(ctx context.Context, g *biz.Tag) (*biz.Tag, error) {
	return g, nil
}

// FindById 通过id查找
func (r *TagRepository) FindById(context.Context, int64) (*biz.Tag, error) {
	return nil, nil
}

// DeleteById 通过id删除
func (r *TagRepository) DeleteById(context.Context, int64) (*biz.Tag, error) {
	return nil, nil
}

// All 获取所有
func (r *TagRepository) All(context.Context) ([]*biz.Tag, error) {
	return nil, nil
}
