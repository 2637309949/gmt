package conf

import (
	"comm/logger"
	"comm/util"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/reader"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

var (
	gtv  map[string]reader.Value
	conf config.Config
)

type Config interface {
	String(defs ...string) (src string)
	StringList(defs ...[]string) (src []string)
	Int(defs ...int64) (src int64)
	Bool(defs ...bool) (src bool)
	Float(defs ...float64) (src float64)
	Scan(v interface{})
}

func init() {
	gtv = map[string]reader.Value{}
	address := flag.String("consul", "", "consul address")
	flag.Parse()

	*address = util.Some(*address, os.Getenv("consul")).(string)
	if registry := os.Getenv("MICRO_REGISTRY"); registry == "consul" {
		*address = util.Some(*address, os.Getenv("MICRO_REGISTRY_ADDRESS")).(string)
	}
	*address = util.Some(*address, "127.0.0.1:8500").(string)

	logger.Infof("configuration center address %v", *address)
	source := consul.NewSource(
		consul.WithAddress(*address),
		consul.WithPrefix("/micro/config"),
		consul.StripPrefix(true),
	)
	var err error
	conf, err = config.NewConfig()
	if err != nil {
		logger.Errorf("NewConfig error %v", err)
	}
	conf.Load(source)
}

// Get defined TODO
func Get(key ...string) reader.Value {
	k := strings.Join(key, ".")
	r, ok := gtv[k]
	if ok {
		return r
	}
	r = conf.Get(key...)

	fmt.Println("----------------66", key)
	fmt.Println("----------------66", string(r.Bytes()))
	// Watch
	go func(key string) {
		w, err := conf.Watch(strings.Split(key, ".")...)
		if err != nil {
			log.Println(err)
		}
		v, err := w.Next()
		if err != nil {
			log.Println(err)
		}
		gtv[key] = v
	}(k)
	gtv[k] = r
	return r
}

// Load defined TODO
func Load(name string, key string) Config {
	vm, b := map[string]interface{}{}, []byte{}
	Get(strings.Split(name, ".")...).Scan(&vm)
	v, ok := vm[key]
	if ok {
		b, _ = json.Marshal(v)
	}
	return &typeConfig{b: b}
}
