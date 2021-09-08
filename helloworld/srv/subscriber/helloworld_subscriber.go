package subscriber

import (
	"context"

	"comm/go-micro/logger"
	"proto/helloworld"
)

func (s *Subscriber) Handle(ctx context.Context, msg *helloworld.Message) error {
	logger.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *helloworld.Message) error {
	logger.Info("Function Received message: ", msg.Say)
	return nil
}
