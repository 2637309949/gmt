package subscriber

import (
	"comm/logger"
	"context"

	"github.com/micro/go-micro/v2/server"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func subscriberCall(next server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) error {
		logger.Infof(ctx, "================ Subscriber [%v] Start ================", msg.Topic())
		defer logger.Infof(ctx, "================ Subscriber [%v] End ================", msg.Topic())
		return next(ctx, msg)
	}
}

func call(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		otCtx := opentracing.SpanFromContext(ctx)
		if spanCtx, ok := otCtx.Context().(jaeger.SpanContext); ok {
			logger.Infof(ctx, "================ %v Start TraceID [%v] SpanID [%v] ================", req.Endpoint(), spanCtx.TraceID(), spanCtx.SpanID())
			defer logger.Infof(ctx, "================ %v End ================", req.Endpoint())
		}
		return fn(ctx, req, rsp)
	}
}

func NewSubscriberLogWrapper() server.SubscriberWrapper {
	return subscriberCall
}

func NewLogWrapper() server.HandlerWrapper {
	return call
}
