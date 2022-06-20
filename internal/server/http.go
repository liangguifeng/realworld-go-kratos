package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "realworld-go-kratos/api/realworld/v1"
	"realworld-go-kratos/internal/conf"
	"realworld-go-kratos/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, user *service.UserService, article *service.ArticleService, comment *service.CommentService, favorite *service.FavoriteService, tag *service.TagService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	v1.RegisterUserHTTPServer(srv, user)
	v1.RegisterArticleHTTPServer(srv, article)
	v1.RegisterCommentHTTPServer(srv, comment)
	v1.RegisterFavoriteHTTPServer(srv, favorite)
	v1.RegisterTagHTTPServer(srv, tag)

	return srv
}
