package conf

import (
	"comm/logger"
	"encoding/json"
	"flag"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/reader"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"os"
	"strings"
)

var conf config.Config

type Config interface {
	String(defs ...string) (src string)
	StringList(defs ...[]string) (src []string)
	Int(defs ...int64) (src int64)
	Bool(defs ...bool) (src bool)
	Float(defs ...float64) (src float64)
	Scan(v interface{})
}

func init() {
	address := flag.String("consul", "", "consul address")
	flag.Parse()
	if *address == "" {
		*address = os.Getenv("consul")
	}
	if *address == "" {
		if registry, registry_address := os.Getenv("MICRO_REGISTRY"), os.Getenv("MICRO_REGISTRY_ADDRESS"); registry == "consul" {
			*address = registry_address
		}
	}
	if *address == "" {
		*address = "127.0.0.1:8500"
	}
	logger.Infof("configuration center address %v", *address)
	source := consul.NewSource(
		consul.WithAddress(*address),
		consul.WithPrefix("/micro/config"),
		consul.StripPrefix(true),
	)
	conf, _ = config.NewConfig()
	conf.Load(source)
}

// Get defined TODO
func Get(key ...string) reader.Value {
	return conf.Get(key...)
}

// Load defined TODO
func Load(name string, key string) Config {
	vm, b := map[string]interface{}{}, []byte{}
	Get(strings.Split(name, ".")...).Scan(&vm)
	v, ok := vm[key]
	if ok {
		b, _ = json.Marshal(v)
	}
	return &typeConfig{val: b}
}
