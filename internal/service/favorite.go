package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "realworld-go-kratos/api/realworld/v1"
	"realworld-go-kratos/internal/biz"
)

type FavoriteService struct {
	v1.UnimplementedFavoriteServer
	uc  *biz.FavoriteUsecase
	log *log.Helper
}

func NewFavoriteService(uc *biz.FavoriteUsecase, logger log.Logger) *FavoriteService {
	return &FavoriteService{uc: uc, log: log.NewHelper(logger)}
}

func (FavoriteService) FavoriteArticle(ctx context.Context, r *v1.FavoriteArticleRequest) (*v1.FavoriteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteArticle not implemented")
}

func (FavoriteService) UnFavoriteArticle(ctx context.Context, r *v1.UnFavoriteArticleRequest) (*v1.UnFavoriteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavoriteArticle not implemented")
}
