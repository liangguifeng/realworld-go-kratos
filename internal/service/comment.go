package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "realworld-go-kratos/api/realworld/v1"
	"realworld-go-kratos/internal/biz"
)

type CommentService struct {
	v1.UnimplementedCommentServer
	uc  *biz.CommentUsecase
	log *log.Helper
}

func NewCommentService(uc *biz.CommentUsecase, logger log.Logger) *CommentService {
	return &CommentService{uc: uc, log: log.NewHelper(logger)}
}

func (CommentService) AddCommentsToAnArticle(ctx context.Context, r *v1.AddCommentsToAnArticleRequest) (*v1.AddCommentsToAnArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCommentsToAnArticle not implemented")
}

func (CommentService) GetCommentsFromAnArticle(ctx context.Context, r *v1.GetCommentsFromAnArticleRequest) (*v1.GetCommentsFromAnArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsFromAnArticle not implemented")
}

func (CommentService) DeleteComment(ctx context.Context, r *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
