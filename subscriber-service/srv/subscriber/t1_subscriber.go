package subscriber

import (
	"context"

	"comm/logger"
	"proto/subscriber"
)

type Subscriber struct{}

func (s *Subscriber) Handle(ctx context.Context, msg *subscriber.Message) error {
	logger.Info("Received srv.Subscriber request")
	return nil
}
