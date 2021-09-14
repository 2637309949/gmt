// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: helloworld/helloworld.handler.proto

package helloworld

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Helloworld service

func NewHelloworldEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Helloworld service

type HelloworldService interface {
	ArticleAdd(ctx context.Context, in *Article, opts ...client.CallOption) (*Article, error)
	ArticleDel(ctx context.Context, in *ArticleFilter, opts ...client.CallOption) (*Article, error)
	ArticleUpdate(ctx context.Context, in *Article, opts ...client.CallOption) (*Article, error)
	ArticleOne(ctx context.Context, in *ArticleFilter, opts ...client.CallOption) (*Article, error)
	ArticlePage(ctx context.Context, in *ArticleFilter, opts ...client.CallOption) (*ArticleList, error)
}

type helloworldService struct {
	c    client.Client
	name string
}

func NewHelloworldService(name string, c client.Client) HelloworldService {
	return &helloworldService{
		c:    c,
		name: name,
	}
}

func (c *helloworldService) ArticleAdd(ctx context.Context, in *Article, opts ...client.CallOption) (*Article, error) {
	req := c.c.NewRequest(c.name, "Helloworld.ArticleAdd", in)
	out := new(Article)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldService) ArticleDel(ctx context.Context, in *ArticleFilter, opts ...client.CallOption) (*Article, error) {
	req := c.c.NewRequest(c.name, "Helloworld.ArticleDel", in)
	out := new(Article)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldService) ArticleUpdate(ctx context.Context, in *Article, opts ...client.CallOption) (*Article, error) {
	req := c.c.NewRequest(c.name, "Helloworld.ArticleUpdate", in)
	out := new(Article)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldService) ArticleOne(ctx context.Context, in *ArticleFilter, opts ...client.CallOption) (*Article, error) {
	req := c.c.NewRequest(c.name, "Helloworld.ArticleOne", in)
	out := new(Article)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldService) ArticlePage(ctx context.Context, in *ArticleFilter, opts ...client.CallOption) (*ArticleList, error) {
	req := c.c.NewRequest(c.name, "Helloworld.ArticlePage", in)
	out := new(ArticleList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Helloworld service

type HelloworldHandler interface {
	ArticleAdd(context.Context, *Article, *Article) error
	ArticleDel(context.Context, *ArticleFilter, *Article) error
	ArticleUpdate(context.Context, *Article, *Article) error
	ArticleOne(context.Context, *ArticleFilter, *Article) error
	ArticlePage(context.Context, *ArticleFilter, *ArticleList) error
}

func RegisterHelloworldHandler(s server.Server, hdlr HelloworldHandler, opts ...server.HandlerOption) error {
	type helloworld interface {
		ArticleAdd(ctx context.Context, in *Article, out *Article) error
		ArticleDel(ctx context.Context, in *ArticleFilter, out *Article) error
		ArticleUpdate(ctx context.Context, in *Article, out *Article) error
		ArticleOne(ctx context.Context, in *ArticleFilter, out *Article) error
		ArticlePage(ctx context.Context, in *ArticleFilter, out *ArticleList) error
	}
	type Helloworld struct {
		helloworld
	}
	h := &helloworldHandler{hdlr}
	return s.Handle(s.NewHandler(&Helloworld{h}, opts...))
}

type helloworldHandler struct {
	HelloworldHandler
}

func (h *helloworldHandler) ArticleAdd(ctx context.Context, in *Article, out *Article) error {
	return h.HelloworldHandler.ArticleAdd(ctx, in, out)
}

func (h *helloworldHandler) ArticleDel(ctx context.Context, in *ArticleFilter, out *Article) error {
	return h.HelloworldHandler.ArticleDel(ctx, in, out)
}

func (h *helloworldHandler) ArticleUpdate(ctx context.Context, in *Article, out *Article) error {
	return h.HelloworldHandler.ArticleUpdate(ctx, in, out)
}

func (h *helloworldHandler) ArticleOne(ctx context.Context, in *ArticleFilter, out *Article) error {
	return h.HelloworldHandler.ArticleOne(ctx, in, out)
}

func (h *helloworldHandler) ArticlePage(ctx context.Context, in *ArticleFilter, out *ArticleList) error {
	return h.HelloworldHandler.ArticlePage(ctx, in, out)
}