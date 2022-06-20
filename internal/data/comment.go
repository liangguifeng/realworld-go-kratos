package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/biz"
)

type CommentRepository struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepository(data *Data, logger log.Logger) biz.CommentRepositoryIface {
	return &CommentRepository{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 保存
func (r *CommentRepository) Create(ctx context.Context, g *biz.Comment) (*biz.Comment, error) {
	return g, nil
}

// Update 更新
func (r *CommentRepository) Update(ctx context.Context, g *biz.Comment) (*biz.Comment, error) {
	return g, nil
}

// FindById 通过id查找
func (r *CommentRepository) FindById(context.Context, int64) (*biz.Comment, error) {
	return nil, nil
}

// DeleteById 通过id删除
func (r *CommentRepository) DeleteById(context.Context, int64) (*biz.Comment, error) {
	return nil, nil
}

// All 获取所有
func (r *CommentRepository) All(context.Context) ([]*biz.Comment, error) {
	return nil, nil
}
