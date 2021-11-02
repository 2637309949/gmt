package handler

import (
	"comm/logger"
	"comm/micro/api"
	"context"
	"proto/{{.name}}"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

func (h *Handler) {{toTitle .proto.Name}}Add(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.{{toTitle .proto.Name}}Add request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := {{.name}}.{{toTitle .proto.Name}}AddReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}Add res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}Add req [%v]", r.String())

	{{.proto.Name}}AddRsp, err := h.{{toTitle .name}}Service.{{toTitle .proto.Name}}Add(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "{{toTitle .name}}Service.{{toTitle .proto.Name}}Add error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, {{.proto.Name}}AddRsp, err)
	return nil
}

func (h *Handler) {{toTitle .proto.Name}}Del(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.{{toTitle .proto.Name}}Del request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := {{.name}}.{{toTitle .proto.Name}}DelReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}Del res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}Del req [%v]", r.String())

	{{.proto.Name}}DelRsp, err := h.{{toTitle .name}}Service.{{toTitle .proto.Name}}Del(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "{{toTitle .name}}Service.{{toTitle .proto.Name}}Del error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, {{.proto.Name}}DelRsp, err)
	return nil
}

func (h *Handler) {{toTitle .proto.Name}}Update(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.{{toTitle .proto.Name}}Update request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := {{.name}}.{{toTitle .proto.Name}}UpdateReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}Update res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}Update req [%v]", r.String())

	{{.proto.Name}}UpdateRsp, err := h.{{toTitle .name}}Service.{{toTitle .proto.Name}}Update(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "{{toTitle .name}}Service.{{toTitle .proto.Name}}Update error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, {{.proto.Name}}UpdateRsp, err)
	return nil
}

func (h *Handler) {{toTitle .proto.Name}}One(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.{{toTitle .proto.Name}}One request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := {{.name}}.{{toTitle .proto.Name}}OneReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}One res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}One req [%v]", r.String())

	{{.proto.Name}}OneRsp, err := h.{{toTitle .name}}Service.{{toTitle .proto.Name}}One(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "{{toTitle .name}}Service.{{toTitle .proto.Name}}One error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, {{.proto.Name}}OneRsp, err)
	return nil
}

func (h *Handler) {{toTitle .proto.Name}}Page(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info(ctx, "Handler.{{toTitle .proto.Name}}Page request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r, rspData := {{.name}}.{{toTitle .proto.Name}}PageReq{}, ""

	defer func() {
		logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}Page res [%v]", rspData)
	}()

	err := request.Bind(&r)
	if err != nil {
		logger.Errorf(ctx, "Request.Bind error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}

	logger.Infof(ctx, "Handler.{{toTitle .proto.Name}}Page req [%v]", r.String())

	{{.proto.Name}}PageRsp, err := h.{{toTitle .name}}Service.{{toTitle .proto.Name}}Page(ctx, &r)
	if err != nil {
		logger.Errorf(ctx, "{{toTitle .name}}Service.{{toTitle .proto.Name}}Page error %v", err)
		rspData = response.Build(ctx, nil, err)
		return nil
	}
	rspData = response.Build(ctx, {{.proto.Name}}PageRsp, err)
	return nil
}
