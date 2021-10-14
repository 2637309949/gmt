package handler

import (
	"comm/logger"
	"comm/micro/api"
	"context"
	"proto/aggregate"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

func (h *Handler) UploadDoc(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received UploadDoc request")

	request, response := api.NewRequest(req), api.NewResponse(rsp)
	r := aggregate.UploadDocReq{}
	if err := request.Bind(&r); err != nil {
		logger.Infof("Request.Bind error %v", err)
		return err
	}
	return response.Build(h.AggregateService.UploadDoc(ctx, &r))
}
