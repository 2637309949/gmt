package micro

import (
	"comm/conf"
	"fmt"
	"strings"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/cache"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/transport/grpc"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
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

// Version of the service
func Version(v string) micro.Option {
	return micro.Version(v)
}

// Metadata associated with the service
func Metadata(md map[string]string) micro.Option {
	return micro.Metadata(md)
}

// NewServiceWithName defined TODO
func NewServiceWithName(name string) micro.Service {
	return NewService(Name(name))
}

// NewService creates and returns a new Service based on the packages within.
func NewService(opts ...micro.Option) micro.Service {
	registryAddress := conf.Load("comm", "registry_address").String("127.0.0.1:8500")
	brokerAddress := conf.Load("comm", "broker_address").String("127.0.0.1:4222")

	logger.Infof("registry_address:%v", registryAddress)
	logger.Infof("broker_address:%v", brokerAddress)

	opts = append(opts, micro.Version("latest"))
	opts = append(opts, micro.Transport(grpc.NewTransport()))
	opts = append(opts, micro.Registry(cache.New(consul.NewRegistry(func(op *registry.Options) { op.Addrs = []string{registryAddress} }))))
	opts = append(opts, micro.Broker(nats.NewBroker(func(op *broker.Options) { op.Addrs = []string{brokerAddress} })))
	srv := micro.NewService(opts...)
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
