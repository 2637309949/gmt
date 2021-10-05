package handler

import (
	"context"

	"proto/push"

	"github.com/netdata/go-orchestrator/logger"
)

// PushEventT1 defined TODO
func (h *Handler) PushEventT1(ctx context.Context, req *push.EventT1, rsp *push.EventT1) error {
	logger.Info("Received srv.push.Handler PushEventT1")
	err := h.Publisher.PushT1(ctx, req)
	if err != nil {
		logger.Infof("error publishing: %v", err)
		return err
	}
	return nil
}
