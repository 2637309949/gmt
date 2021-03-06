package conf

import (
	"comm/util/base"
	"encoding/json"
	"flag"
	"os"
	"strings"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/reader"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/config/source/consul/v2"
	lgs "github.com/micro/go-plugins/logger/logrus/v2"
	"github.com/sirupsen/logrus"
)

var (
	gtv  map[string]reader.Value
	conf config.Config
	l    logger.Logger
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
	logrusLogger := logrus.New()
	l = logger.NewHelper(logger.NewLogger(lgs.WithLogger(logrusLogger)))
	gtv = map[string]reader.Value{}
	address := *flag.String("consul", "", "consul address")
	flag.Parse()

	address = base.Some(address, os.Getenv("consul")).(string)
	if registry := os.Getenv("MICRO_REGISTRY"); registry == "consul" {
		address = base.Some(address, os.Getenv("MICRO_REGISTRY_ADDRESS")).(string)
	}
	address = base.Some(address, "127.0.0.1:8500").(string)

	l.Logf(logger.InfoLevel, "configuration center address %v", address)
	source := consul.NewSource(
		consul.WithAddress(address),
		consul.WithPrefix("/micro/config"),
		consul.StripPrefix(true),
	)
	var err error
	conf, err = config.NewConfig()
	if err != nil {
		l.Logf(logger.ErrorLevel, "NewConfig error %v", err)
	}
	conf.Load(source)
}

// Set defined TODO
func Set(val interface{}, path ...string) {
	conf.Set(val, path...)
}

// Del defined TODO
func Del(path ...string) {
	conf.Del(path...)
}

// Get defined TODO
func Get(key ...string) reader.Value {
	k := strings.Join(key, ".")
	r, ok := gtv[k]
	if ok {
		return r
	}
	r = conf.Get(key...)
	gtv[k] = r

	// Watch
	go func(key string) {
		w, err := conf.Watch(strings.Split(key, ".")...)
		if err != nil {
			l.Log(logger.ErrorLevel, err)
		}
		v, err := w.Next()
		if err != nil {
			l.Log(logger.ErrorLevel, err)
		}
		gtv[key] = v
	}(k)
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
