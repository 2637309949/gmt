package handler

import (
	"context"

	"comm/logger"
	"proto/helloworld"
)

// Call is a single request handler called via client.Call or the generated client code
func (h *Handler) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	logger.Info("Received Helloworld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
