package handler

import (
	"comm/logger"
	"comm/micro/api"
	"context"
	"proto/{{.name}}"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = {{.name}}.{{toTitle .proto.Name}}{}
var _ = api.Request{}
var _ = logger.Info

func (h *Handler) {{toTitle .proto.Name}}Add(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .proto.Name}}Add request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .proto.Name}}AddReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .proto.Name}}Add(ctx, &r))
}

func (h *Handler) {{toTitle .proto.Name}}Del(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .proto.Name}}Del request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .proto.Name}}DelReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .proto.Name}}Del(ctx, &r))
}

func (h *Handler) {{toTitle .proto.Name}}Update(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .proto.Name}}Update request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .proto.Name}}UpdateReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .proto.Name}}Update(ctx, &r))
}

func (h *Handler) {{toTitle .proto.Name}}One(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .proto.Name}}One request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .proto.Name}}OneReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .proto.Name}}One(ctx, &r))
}

func (h *Handler) {{toTitle .proto.Name}}Page(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received {{toTitle .proto.Name}}Page request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := {{.name}}.{{toTitle .proto.Name}}PageReq{Page: 1, Size: 20}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.{{toTitle .name}}Service.{{toTitle .proto.Name}}Page(ctx, &r))
}
