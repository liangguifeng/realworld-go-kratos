package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "realworld-go-kratos/api/realworld/v1"
	"realworld-go-kratos/internal/biz"
)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer
	uc  *biz.RealWorldUsecase
	log *log.Helper
}

func NewRealWorldServiceService(uc *biz.RealWorldUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}

// Login
func (s RealWorldService) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	RealWorld, err := s.uc.CreateRealWorld(ctx, &biz.User{})
	if err != nil {
		return nil, err
	}
	fmt.Println(RealWorld)
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

// Register
func (RealWorldService) Register(ctx context.Context, r *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

// GetCurrentUser
func (RealWorldService) GetCurrentUser(ctx context.Context, r *v1.GetCurrentUserRequest) (*v1.GetCurrentUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentRealWorld not implemented")
}

// UpdateUser
func (RealWorldService) UpdateUser(ctx context.Context, r *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRealWorld not implemented")
}

func (RealWorldService) GetProfile(ctx context.Context, r *v1.GetProfileRequest) (*v1.GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}

func (RealWorldService) FollowUser(ctx context.Context, r *v1.FollowUserRequest) (*v1.FollowUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowRealWorld not implemented")
}

func (RealWorldService) UnfollowUser(ctx context.Context, r *v1.UnfollowUserRequest) (*v1.UnfollowUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowRealWorld not implemented")
}

func (RealWorldService) ListArticles(ctx context.Context, r *v1.ListArticlesRequest) (*v1.ListArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}

func (RealWorldService) FeedArticles(ctx context.Context, r *v1.FeedArticlesRequest) (*v1.FeedArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}

func (RealWorldService) GetArticle(ctx context.Context, r *v1.GetArticleRequest) (*v1.GetArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}

func (RealWorldService) CreateArticle(ctx context.Context, r *v1.CreateArticleRequest) (*v1.CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}

func (RealWorldService) UpdateArticle(ctx context.Context, r *v1.UpdateArticleRequest) (*v1.UpdateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}

func (RealWorldService) DeleteArticle(ctx context.Context, r *v1.DeleteArticleRequest) (*v1.DeleteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}

func (RealWorldService) FavoriteArticle(ctx context.Context, r *v1.FavoriteArticleRequest) (*v1.FavoriteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteArticle not implemented")
}

func (RealWorldService) UnFavoriteArticle(ctx context.Context, r *v1.UnFavoriteArticleRequest) (*v1.UnFavoriteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavoriteArticle not implemented")
}

func (RealWorldService) AddCommentsToAnArticle(ctx context.Context, r *v1.AddCommentsToAnArticleRequest) (*v1.AddCommentsToAnArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCommentsToAnArticle not implemented")
}

func (RealWorldService) GetCommentsFromAnArticle(ctx context.Context, r *v1.GetCommentsFromAnArticleRequest) (*v1.GetCommentsFromAnArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsFromAnArticle not implemented")
}

func (RealWorldService) DeleteComment(ctx context.Context, r *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}

func (RealWorldService) GetTags(ctx context.Context, r *v1.GetTagsRequest) (*v1.GetTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
