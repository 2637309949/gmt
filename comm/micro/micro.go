package micro

import (
	"comm/conf"
	"comm/logger"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/cache"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-micro/v2/web"
	brokerNats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	breakerHystrix "github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	monitoringPrometheus "github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	traceOpentracing "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var DefaultServiceNamePrefix = "go.micro."
var name string

// GetServiceName defined TODO
func GetServiceName() string {
	return name
}

// SetServiceName defined TODO
func SetServiceName(n string) {
	if name != "" {
		return
	}
	name = n
}

// Key defined TODO
func Key(key string) conf.Config {
	return conf.Load(GetServiceName(), key)
}

// Name of the service
func Name(name string) micro.Option {
	SetServiceName(strings.ReplaceAll(name, NameFormat(""), ""))
	return micro.Name(name)
}

// NameFormat defined TODO
func NameFormat(n string) string {
	return fmt.Sprintf("%v%v", DefaultServiceNamePrefix, n)
}

// NewServiceWithName defined TODO
func NewServiceWithName(name string) micro.Service {
	return NewService(Name(name))
}

func NewJaegerTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Reporter(reporter),
	)

	return tracer, closer, err
}

// NewService creates and returns a new Service based on the packages within.
func NewService(opts ...micro.Option) micro.Service {
	registryAddress := conf.Load("comm", "registry_address").String("127.0.0.1:8500")
	brokerAddress := conf.Load("comm", "broker_address").String("127.0.0.1:4222")
	opentracingAddress := conf.Load("comm", "opentracing_address").String("127.0.0.1:6831")
	jaegerTracer, _, _ := NewJaegerTracer(NameFormat(GetServiceName()), opentracingAddress)
	hystrix.DefaultSleepWindow = 200
	hystrix.DefaultMaxConcurrent = 3

	logger.Infof("registry_address:%v", registryAddress)
	logger.Infof("broker_address:%v", brokerAddress)
	logger.Infof("opentracing_address:%v", opentracingAddress)

	opts = append(opts, micro.Version("latest"))
	opts = append(opts, micro.Transport(grpc.NewTransport()))
	opts = append(opts, micro.Registry(cache.New(consul.NewRegistry(func(op *registry.Options) { op.Addrs = []string{registryAddress} }))))
	opts = append(opts, micro.Broker(brokerNats.NewBroker(func(op *broker.Options) { op.Addrs = []string{brokerAddress} })))
	opts = append(opts, micro.WrapHandler(traceOpentracing.NewHandlerWrapper(jaegerTracer)))
	opts = append(opts, micro.WrapClient(traceOpentracing.NewClientWrapper(jaegerTracer)))
	opts = append(opts, micro.WrapSubscriber(traceOpentracing.NewSubscriberWrapper(jaegerTracer)))
	opts = append(opts, micro.WrapClient(breakerHystrix.NewClientWrapper()))
	opts = append(opts, micro.WrapHandler(monitoringPrometheus.NewHandlerWrapper()))
	opts = append(opts, micro.Transport(grpc.NewTransport()))

	srv := micro.NewService(opts...)
	defer func() {
		opts := []web.Option{}
		opts = append(opts, web.Name(NameFormat(GetServiceName())+".metrics"))
		opts = append(opts, web.Handler(http.DefaultServeMux))
		opts = append(opts, web.Registry(cache.New(consul.NewRegistry(func(op *registry.Options) { op.Addrs = []string{registryAddress} }))))
		metrics := web.NewService(opts...)
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			if err := metrics.Run(); err != nil {
				logger.Fatalf("failed to serve: %v", err)
			}
		}()
	}()
	defer srv.Init()
	return srv
}

// NewEvent creates a new event publisher
func NewEvent(topic string, c client.Client) micro.Event {
	return micro.NewEvent(topic, c)
}

// RegisterHandler is syntactic sugar for registering a handler
func RegisterHandler(s server.Server, h interface{}, opts ...server.HandlerOption) error {
	return s.Handle(s.NewHandler(h, opts...))
}

// RegisterSubscriber is syntactic sugar for registering a subscriber
func RegisterSubscriber(topic string, s server.Server, h interface{}, opts ...server.SubscriberOption) error {
	return micro.RegisterSubscriber(topic, s, h, opts...)
}
