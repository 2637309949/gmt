package handler

import (
	"comm/logger"
	"comm/micro/api"
	"context"
	"proto/helloworld"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

func (h *Handler) Call(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received api.helloworld.call API request")
	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := helloworld.Request{}
	if err := request.Bind(&r); err != nil {
		return err
	}
	return response.Build(h.HelloworldClient.Call(ctx, &r))
}
