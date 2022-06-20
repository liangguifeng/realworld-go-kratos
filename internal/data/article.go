package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/biz"
)

type ArticleRepository struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepository(data *Data, logger log.Logger) biz.ArticleRepositoryIface {
	return &ArticleRepository{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 保存
func (r *ArticleRepository) Create(ctx context.Context, g *biz.Article) (*biz.Article, error) {
	return g, nil
}

// Update 更新
func (r *ArticleRepository) Update(ctx context.Context, g *biz.Article) (*biz.Article, error) {
	return g, nil
}

// FindById 通过id查找
func (r *ArticleRepository) FindById(context.Context, int64) (*biz.Article, error) {
	return nil, nil
}

// DeleteById 通过id删除
func (r *ArticleRepository) DeleteById(context.Context, int64) (*biz.Article, error) {
	return nil, nil
}

// All 获取所有
func (r *ArticleRepository) All(context.Context) ([]*biz.Article, error) {
	return nil, nil
}
