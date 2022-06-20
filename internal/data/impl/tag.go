package impl

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/data"
	"realworld-go-kratos/internal/data/mysql_model"
)

type TagRepository struct {
	data *data.Data
	log  *log.Helper
}

func NewTagRepository(data *data.Data, logger log.Logger) data.TagRepositoryIface {
	return &TagRepository{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 保存
func (r *TagRepository) Create(ctx context.Context, g *mysql_model.Tag) (*mysql_model.Tag, error) {
	return g, nil
}

// Update 更新
func (r *TagRepository) Update(ctx context.Context, g *mysql_model.Tag) (*mysql_model.Tag, error) {
	return g, nil
}

// FindById 通过id查找
func (r *TagRepository) FindById(context.Context, int64) (*mysql_model.Tag, error) {
	return nil, nil
}

// DeleteById 通过id删除
func (r *TagRepository) DeleteById(context.Context, int64) (*mysql_model.Tag, error) {
	return nil, nil
}

// All 获取所有
func (r *TagRepository) All(context.Context) ([]*mysql_model.Tag, error) {
	return nil, nil
}
