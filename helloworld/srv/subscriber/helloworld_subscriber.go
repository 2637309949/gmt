package subscriber

import (
	"context"

	"comm/logger"
	"proto/helloworld"
)

func (s *Subscriber) Handle(ctx context.Context, msg *helloworld.Message) error {
	logger.Info("Handler Received message: ", msg.Say)
	return nil
}
