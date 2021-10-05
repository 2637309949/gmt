package subscriber

import (
	"context"

	"comm/logger"
	"proto/helloworld"
)

type Subscriber struct{}

func (s *Subscriber) Handle(ctx context.Context, msg *helloworld.Message) error {
	logger.Info("Received srv.helloworld.Subscriber request")
	return nil
}