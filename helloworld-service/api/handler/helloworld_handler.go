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
	logger.Info("Handler.ArticleAdd request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticleAddReq{}, ""

	defer func() {
		logger.Infof("Handler.ArticleAdd res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf("Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof("Handler.ArticleAdd req [%v]", r.String())

	articleAddRsp, err := h.HelloworldService.ArticleAdd(ctx, &r)
	if err != nil {
		logger.Errorf("HelloworldService.ArticleAdd error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, articleAddRsp, err)
	return nil
}

// ArticleDel defined TODO
func (h *Handler) ArticleDel(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Handler.ArticleDel request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticleDelReq{}, ""

	defer func() {
		logger.Infof("Handler.ArticleDel res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf("Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof("Handler.ArticleDel req [%v]", r.String())

	ArticleDelRsp, err := h.HelloworldService.ArticleDel(ctx, &r)
	if err != nil {
		logger.Errorf("HelloworldService.ArticleDel error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, ArticleDelRsp, err)
	return nil
}

// ArticleUpdate defined TODO
func (h *Handler) ArticleUpdate(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Handler.ArticleUpdate request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticleUpdateReq{}, ""

	defer func() {
		logger.Infof("Handler.ArticleUpdate res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf("Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof("Handler.ArticleUpdate req [%v]", r.String())

	ArticleUpdateRsp, err := h.HelloworldService.ArticleUpdate(ctx, &r)
	if err != nil {
		logger.Errorf("HelloworldService.ArticleUpdate error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, ArticleUpdateRsp, err)
	return nil
}

// ArticleOne defined TODO
func (h *Handler) ArticleOne(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Handler.ArticleOne request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticleOneReq{}, ""

	defer func() {
		logger.Infof("Handler.ArticleOne res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf("Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof("Handler.ArticleOne req [%v]", r.String())

	ArticleOneRsp, err := h.HelloworldService.ArticleOne(ctx, &r)
	if err != nil {
		logger.Errorf("HelloworldService.ArticleOne error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, ArticleOneRsp, err)
	return nil
}

// ArticlePage defined TODO
func (h *Handler) ArticlePage(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Handler.ArticlePage request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := helloworld.ArticlePageReq{}, ""

	defer func() {
		logger.Infof("Handler.ArticlePage res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf("Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof("Handler.ArticlePage req [%v]", r.String())

	ArticlePageRsp, err := h.HelloworldService.ArticlePage(ctx, &r)
	if err != nil {
		logger.Errorf("HelloworldService.ArticlePage error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, ArticlePageRsp, err)
	return nil
}
