package handler

import (
	"comm/logger"
	"comm/micro/api"
	"context"
	"proto/helloworld"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = helloworld.Article{}
var _ = api.Request{}
var _ = logger.Info

func (h *Handler) ArticleAdd(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received ArticleAdd request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := helloworld.ArticleAddReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.HelloworldService.ArticleAdd(ctx, &r))
}

func (h *Handler) ArticleDel(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received ArticleDel request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := helloworld.ArticleDelReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.HelloworldService.ArticleDel(ctx, &r))
}

func (h *Handler) ArticleUpdate(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received ArticleUpdate request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := helloworld.ArticleUpdateReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.HelloworldService.ArticleUpdate(ctx, &r))
}

func (h *Handler) ArticleOne(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received ArticleOne request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := helloworld.ArticleOneReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.HelloworldService.ArticleOne(ctx, &r))
}

func (h *Handler) ArticlePage(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received ArticlePage request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := helloworld.ArticlePageReq{Page: 1, Size: 20}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.HelloworldService.ArticlePage(ctx, &r))
}
