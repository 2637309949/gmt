package handler

import (
	"comm/logger"
	"comm/micro/api"
	"context"
	"proto/helloworld"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

// ArticleAdd defined TODO
func (h *Handler) ArticleAdd(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.ArticleAdd request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticleAddReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.ArticleAdd res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.ArticleAdd req [%v]", r.String())

	articleAddRsp, err := h.HelloworldService.ArticleAdd(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "HelloworldService.ArticleAdd error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, articleAddRsp, err)
	return nil
}

// ArticleDel defined TODO
func (h *Handler) ArticleDel(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.ArticleDel request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticleDelReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.ArticleDel res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.ArticleDel req [%v]", r.String())

	ArticleDelRsp, err := h.HelloworldService.ArticleDel(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "HelloworldService.ArticleDel error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, ArticleDelRsp, err)
	return nil
}

// ArticleUpdate defined TODO
func (h *Handler) ArticleUpdate(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.ArticleUpdate request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticleUpdateReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.ArticleUpdate res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.ArticleUpdate req [%v]", r.String())

	ArticleUpdateRsp, err := h.HelloworldService.ArticleUpdate(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "HelloworldService.ArticleUpdate error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, ArticleUpdateRsp, err)
	return nil
}

// ArticleOne defined TODO
func (h *Handler) ArticleOne(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.ArticleOne request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticleOneReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.ArticleOne res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.ArticleOne req [%v]", r.String())

	ArticleOneRsp, err := h.HelloworldService.ArticleOne(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "HelloworldService.ArticleOne error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, ArticleOneRsp, err)
	return nil
}

// ArticlePage defined TODO
func (h *Handler) ArticlePage(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.ArticlePage request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticlePageReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.ArticlePage res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.ArticlePage req [%v]", r.String())

	ArticlePageRsp, err := h.HelloworldService.ArticlePage(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "HelloworldService.ArticlePage error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, ArticlePageRsp, err)
	return nil
}
