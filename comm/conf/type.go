package conf

import (
	"comm/util/base"
	"encoding/json"
)

type typeConfig struct {
	b []byte
}

// StringList defined TODO
func (t *typeConfig) StringList(defs ...[]string) (src []string) {
	if base.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// String defined TODO
func (t *typeConfig) String(defs ...string) (src string) {
	if base.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// IntList defined TODO
func (t *typeConfig) IntList(defs ...[]int) (src []int) {
	if base.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// Int defined TODO
func (t *typeConfig) Int(defs ...int64) (src int64) {
	if base.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// BoolList defined TODO
func (t *typeConfig) BoolList(defs ...[]bool) (src []bool) {
	if base.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// Bool defined TODO
func (t *typeConfig) Bool(defs ...bool) (src bool) {
	if base.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// FloatList defined TODO
func (t *typeConfig) FloatList(defs ...[]float64) (src []float64) {
	if base.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// Float64 defined TODO
func (t *typeConfig) Float(defs ...float64) (src float64) {
	if base.IsBlank(t.b) && len(defs) > 0 {
		return defs[0]
	}
	json.Unmarshal(t.b, &src)
	return src
}

// Scan defined TODO
func (t *typeConfig) Scan(v interface{}) {
	json.Unmarshal(t.b, v)
}
