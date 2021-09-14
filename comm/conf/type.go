package conf

import (
	"comm/util/reflect"
	"encoding/json"
)

type typeConfig struct {
	b []byte
}

// StringList defined TODO
func (t *typeConfig) StringList(defs ...[]string) (src []string) {
	if reflect.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// String defined TODO
func (t *typeConfig) String(defs ...string) (src string) {
	if reflect.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// Int defined TODO
func (t *typeConfig) Int(defs ...int64) (src int64) {
	if reflect.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// Bool defined TODO
func (t *typeConfig) Bool(defs ...bool) (src bool) {
	if reflect.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// Float64 defined TODO
func (t *typeConfig) Float(defs ...float64) (src float64) {
	if reflect.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// Scan defined TODO
func (t *typeConfig) Scan(v interface{}) {
	json.Unmarshal(t.b, v)
}
