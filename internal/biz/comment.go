package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Comment struct {
}

type CommentRepositoryIface interface {
	Create(context.Context, *Comment) (*Comment, error)
	Update(context.Context, *Comment) (*Comment, error)
	FindById(context.Context, int64) (*Comment, error)
	DeleteById(context.Context, int64) (*Comment, error)
	All(context.Context) ([]*Comment, error)
}

type CommentUsecase struct {
	CommentRepo CommentRepositoryIface
	log         *log.Helper
}

func NewCommentUsecase(repo CommentRepositoryIface, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{CommentRepo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentUsecase) CreateRealWorld(ctx context.Context, g *Comment) (*Comment, error) {
	return uc.CommentRepo.Create(ctx, g)
}
