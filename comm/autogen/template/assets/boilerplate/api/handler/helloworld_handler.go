package handler

import (
	"comm/logger"
	"comm/micro/api"
	"context"
	"proto/{{.name}}"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = {{.name}}.{{toTitle .entity}}{}
var _ = api.Request{}
var _ = logger.Info

func (h *Handler) {{toTitle .entity}}Add(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .entity}}Add request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .entity}}AddReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .entity}}Add(ctx, &r))
}

func (h *Handler) {{toTitle .entity}}Del(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .entity}}Del request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .entity}}DelReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .entity}}Del(ctx, &r))
}

func (h *Handler) {{toTitle .entity}}Update(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .entity}}Update request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .entity}}UpdateReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .entity}}Update(ctx, &r))
}

func (h *Handler) {{toTitle .entity}}One(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .entity}}One request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .entity}}OneReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .entity}}One(ctx, &r))
}

func (h *Handler) {{toTitle .entity}}Page(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .entity}}Page request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .entity}}PageReq{Page: 1, Size: 20}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .entity}}Page(ctx, &r))
}
