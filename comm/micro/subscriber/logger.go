package subscriber

import (
	"context"

	"github.com/micro/go-micro/v2/server"
)

func subscriberCall(next server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) error {
		return next(ctx, msg)
	}
}

func call(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		return fn(ctx, req, rsp)
	}
}

func NewSubscriberLogWrapper() server.SubscriberWrapper {
	return subscriberCall
}

func NewLogWrapper() server.HandlerWrapper {
	return call
}
