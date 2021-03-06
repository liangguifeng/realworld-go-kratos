// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationArticleCreateArticle = "/realworld.v1.Article/CreateArticle"
const OperationArticleDeleteArticle = "/realworld.v1.Article/DeleteArticle"
const OperationArticleFeedArticles = "/realworld.v1.Article/FeedArticles"
const OperationArticleGetArticle = "/realworld.v1.Article/GetArticle"
const OperationArticleListArticles = "/realworld.v1.Article/ListArticles"
const OperationArticleUpdateArticle = "/realworld.v1.Article/UpdateArticle"

type ArticleHTTPServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*FeedArticlesResponse, error)
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleResponse, error)
	ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesResponse, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error)
}

func RegisterArticleHTTPServer(s *http.Server, srv ArticleHTTPServer) {
	r := s.Route("/")
	r.GET("/api/articles", _Article_ListArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/feed", _Article_FeedArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/{slug}", _Article_GetArticle0_HTTP_Handler(srv))
	r.POST("/api/articles", _Article_CreateArticle0_HTTP_Handler(srv))
	r.PUT("/api/articles/{slug}", _Article_UpdateArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}", _Article_DeleteArticle0_HTTP_Handler(srv))
}

func _Article_ListArticles0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleListArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticles(ctx, req.(*ListArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListArticlesResponse)
		return ctx.Result(200, reply)
	}
}

func _Article_FeedArticles0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeedArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleFeedArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeedArticles(ctx, req.(*FeedArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FeedArticlesResponse)
		return ctx.Result(200, reply)
	}
}

func _Article_GetArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleGetArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticle(ctx, req.(*GetArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetArticleResponse)
		return ctx.Result(200, reply)
	}
}

func _Article_CreateArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleCreateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*CreateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateArticleResponse)
		return ctx.Result(200, reply)
	}
}

func _Article_UpdateArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleUpdateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*UpdateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateArticleResponse)
		return ctx.Result(200, reply)
	}
}

func _Article_DeleteArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleDeleteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*DeleteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteArticleResponse)
		return ctx.Result(200, reply)
	}
}

type ArticleHTTPClient interface {
	CreateArticle(ctx context.Context, req *CreateArticleRequest, opts ...http.CallOption) (rsp *CreateArticleResponse, err error)
	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, opts ...http.CallOption) (rsp *DeleteArticleResponse, err error)
	FeedArticles(ctx context.Context, req *FeedArticlesRequest, opts ...http.CallOption) (rsp *FeedArticlesResponse, err error)
	GetArticle(ctx context.Context, req *GetArticleRequest, opts ...http.CallOption) (rsp *GetArticleResponse, err error)
	ListArticles(ctx context.Context, req *ListArticlesRequest, opts ...http.CallOption) (rsp *ListArticlesResponse, err error)
	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, opts ...http.CallOption) (rsp *UpdateArticleResponse, err error)
}

type ArticleHTTPClientImpl struct {
	cc *http.Client
}

func NewArticleHTTPClient(client *http.Client) ArticleHTTPClient {
	return &ArticleHTTPClientImpl{client}
}

func (c *ArticleHTTPClientImpl) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...http.CallOption) (*CreateArticleResponse, error) {
	var out CreateArticleResponse
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationArticleCreateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...http.CallOption) (*DeleteArticleResponse, error) {
	var out DeleteArticleResponse
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleDeleteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...http.CallOption) (*FeedArticlesResponse, error) {
	var out FeedArticlesResponse
	pattern := "/api/articles/feed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleFeedArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...http.CallOption) (*GetArticleResponse, error) {
	var out GetArticleResponse
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleGetArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...http.CallOption) (*ListArticlesResponse, error) {
	var out ListArticlesResponse
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleListArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...http.CallOption) (*UpdateArticleResponse, error) {
	var out UpdateArticleResponse
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationArticleUpdateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
