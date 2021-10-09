package publisher

import (
	"context"
	"proto/publisher"

	"comm/logger"
)

// PushT1 defined TODO
func (p *Publisher) PushT1(ctx context.Context, msg *publisher.EventT1) error {
	logger.Info("Received srv.publisher.Publisher PushT1")
	err := p.Client.Publish(context.Background(), msg)
	if err != nil {
		logger.Infof("error publishing: %v", err)
		return err
	}
	return nil
}
