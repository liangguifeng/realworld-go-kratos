package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "realworld-go-kratos/api/realworld/v1"
	"realworld-go-kratos/internal/biz"
)

type ArticleService struct {
	v1.UnimplementedArticleServer
	uc  *biz.ArticleUsecase
	log *log.Helper
}

func NewArticleService(uc *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{uc: uc, log: log.NewHelper(logger)}
}

func (ArticleService) ListArticles(ctx context.Context, r *v1.ListArticlesRequest) (*v1.ListArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}

func (ArticleService) FeedArticles(ctx context.Context, r *v1.FeedArticlesRequest) (*v1.FeedArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}

func (ArticleService) GetArticle(ctx context.Context, r *v1.GetArticleRequest) (*v1.GetArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}

func (ArticleService) CreateArticle(ctx context.Context, r *v1.CreateArticleRequest) (*v1.CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}

func (ArticleService) UpdateArticle(ctx context.Context, r *v1.UpdateArticleRequest) (*v1.UpdateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}

func (ArticleService) DeleteArticle(ctx context.Context, r *v1.DeleteArticleRequest) (*v1.DeleteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
