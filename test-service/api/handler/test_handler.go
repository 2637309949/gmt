package handler

import (
	"comm/logger"
	"comm/micro/api"
	"context"
	"proto/test"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = test.Entity{}
var _ = api.Request{}
var _ = logger.Info

func (h *Handler) EntityAdd(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received EntityAdd request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := test.EntityAddReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.TestService.EntityAdd(ctx, &r))
}

func (h *Handler) EntityDel(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received EntityDel request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := test.EntityDelReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.TestService.EntityDel(ctx, &r))
}

func (h *Handler) EntityUpdate(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received EntityUpdate request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := test.EntityUpdateReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.TestService.EntityUpdate(ctx, &r))
}

func (h *Handler) EntityOne(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received EntityOne request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := test.EntityOneReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.TestService.EntityOne(ctx, &r))
}

func (h *Handler) EntityPage(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received EntityPage request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := test.EntityPageReq{Page: 1, Size: 20}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.TestService.EntityPage(ctx, &r))
}
