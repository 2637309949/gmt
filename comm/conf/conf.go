package conf

import (
	"comm/util/reflect"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/micro/go-micro/config/reader"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"os"
	"strconv"
	"strings"
)

var conf config.Config

type typeConfig struct {
	val interface{}
}

type Config interface {
	String(defs ...string) (src string)
	Int(defs ...int) (src int)
	Int64(defs ...int64) (src int64)
	Bool(defs ...bool) (src bool)
	Float64(defs ...float64) (src float64)
	Float32(defs ...float32) (src float32)
	Scan(v interface{})
}

// String defined TODO
func (t *typeConfig) String(defs ...string) (src string) {
	if reflect.IsBlank(t.val) && len(defs) > 0 {
		return defs[0]
	}
	return fmt.Sprintf("%v", t.val)
}

// Int defined TODO
func (t *typeConfig) Int(defs ...int) (src int) {
	if reflect.IsBlank(t.val) && len(defs) > 0 {
		return defs[0]
	}
	i, _ := strconv.ParseInt(fmt.Sprintf("%v", t.val), 10, 64)
	return int(i)
}

// Int defined TODO
func (t *typeConfig) Int64(defs ...int64) (src int64) {
	if reflect.IsBlank(t.val) && len(defs) > 0 {
		return defs[0]
	}
	i, _ := strconv.ParseInt(fmt.Sprintf("%v", t.val), 10, 64)
	return i
}

// Bool defined TODO
func (t *typeConfig) Bool(defs ...bool) (src bool) {
	if reflect.IsBlank(t.val) && len(defs) > 0 {
		return defs[0]
	}
	i, _ := strconv.ParseBool(fmt.Sprintf("%v", t.val))
	return i
}

// Float32 defined TODO
func (t *typeConfig) Float32(defs ...float32) (src float32) {
	if reflect.IsBlank(t.val) && len(defs) > 0 {
		return defs[0]
	}
	i, _ := strconv.ParseFloat(fmt.Sprintf("%v", t.val), 64)
	return float32(i)
}

// Float64 defined TODO
func (t *typeConfig) Float64(defs ...float64) (src float64) {
	if reflect.IsBlank(t.val) && len(defs) > 0 {
		return defs[0]
	}
	i, _ := strconv.ParseFloat(fmt.Sprintf("%v", t.val), 64)
	return float64(i)
}

// Scan defined TODO
func (t *typeConfig) Scan(v interface{}) {
	b, _ := json.Marshal(&t.val)
	json.Unmarshal(b, v)
}

func init() {
	address := flag.String("consul", "127.0.0.1:8500", "consul address")
	flag.Parse()
	if es := os.Getenv("consul"); *address == "" && es != "" {
		*address = es
	}
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
	vm := map[string]interface{}{}
	Get(strings.Split(name, ".")...).Scan(&vm)
	return &typeConfig{val: vm[key]}
}
