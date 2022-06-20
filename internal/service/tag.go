package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "realworld-go-kratos/api/realworld/v1"
	"realworld-go-kratos/internal/biz"
)

type TagService struct {
	v1.UnimplementedTagServer
	uc  *biz.TagUsecase
	log *log.Helper
}

func NewTagService(uc *biz.TagUsecase, logger log.Logger) *TagService {
	return &TagService{uc: uc, log: log.NewHelper(logger)}
}

func (TagService) GetTags(ctx context.Context, r *v1.GetTagsRequest) (*v1.GetTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
