package publisher

import (
	"context"

	"comm/logger"
	"proto/push"
)

// PushT1 defined TODO
func (p *Publisher) PushT1(ctx context.Context, msg *push.EventT1) error {
	logger.Info("Received srv.push.Publisher PushT1")
	err := p.Client.Publish(context.Background(), msg)
	if err != nil {
		logger.Infof("error publishing: %v", err)
		return err
	}
	return nil
}
