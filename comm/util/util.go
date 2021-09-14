package util

import (
	"comm/util/reflect"
	rft "reflect"
)

// Some defined TODO
func Some(src ...interface{}) interface{} {
	for i := range src {
		s := src[i]
		if !reflect.IsBlank(s) {
			return s
		}
	}
	return rft.Zero(rft.TypeOf(src[0])).Interface()
}
