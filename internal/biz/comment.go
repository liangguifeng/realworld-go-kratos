package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld-go-kratos/internal/data"
	"realworld-go-kratos/internal/data/mysql_model"
)

type CommentUsecase struct {
	CommentRepo data.CommentRepositoryIface
	log         *log.Helper
}

func NewCommentUsecase(repo data.CommentRepositoryIface, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{CommentRepo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentUsecase) CreateRealWorld(ctx context.Context, g *mysql_model.Comment) (*mysql_model.Comment, error) {
	return uc.CommentRepo.Create(ctx, g)
}
