package micro

import (
	"comm/conf"
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

var name string

// GetName defined TODO
func GetName() string {
	return name
}

// SetName defined TODO
func SetName(n string) {
	name = n
}

// Key defined TODO
func Key(key string) conf.Config {
	return conf.Load(GetName(), key)
}

// Name of the service
func Name(name string) micro.Option {
	SetName(strings.ReplaceAll(name, NameFormat(""), ""))
	return micro.Name(name)
}

// NameFormat defined TODO
func NameFormat(n string) string {
	return "go.micro." + n
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
	opts = append(opts, micro.Version("latest"))
	opts = append(opts, micro.Transport(grpc.NewTransport()))
	registryAddress := conf.Load("comm", "registry_address").String("127.0.0.1:8500")
	brokerAddress := conf.Load("comm", "broker_address").String("127.0.0.1:4222")
	logger.Infof("registry_address:%v", registryAddress)
	opts = append(opts, micro.Registry(cache.New(consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{registryAddress}
	}))))
	opts = append(opts, micro.Broker(nats.NewBroker(func(op *broker.Options) {
		op.Addrs = []string{brokerAddress}
	})))
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
	return micro.RegisterSubscriber(NameFormat(topic), s, h, opts...)
}
