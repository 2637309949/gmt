package base

import (
	"bytes"
	"reflect"
)

// IsBlank defined TODO
func IsBlank(v interface{}) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// Some defined TODO
func Some(arr ...interface{}) interface{} {
	if len(arr) == 0 {
		return nil
	}
	for i := range arr {
		s := arr[i]
		if !IsBlank(s) {
			return s
		}
	}
	return reflect.Zero(sliceElem(reflect.TypeOf(arr[0]))).Interface()
}

// sliceElem defined TODO
func sliceElem(rtype reflect.Type) reflect.Type {
	for {
		if rtype.Kind() != reflect.Slice && rtype.Kind() != reflect.Array {
			return rtype
		}

		rtype = rtype.Elem()
	}
}

// IsEqual defined TODO
func IsEqual(expected interface{}, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	if exp, ok := expected.([]byte); ok {
		act, ok := actual.([]byte)
		if !ok {
			return false
		}
		if exp == nil || act == nil {
			return true
		}
		return bytes.Equal(exp, act)
	}
	return reflect.DeepEqual(expected, actual)
}

// IsType defined TODO
func IsType(expected interface{}, actual interface{}) bool {
	return IsEqual(reflect.TypeOf(expected), reflect.TypeOf(actual))
}

// Equal defined TODO
func Equal(expected interface{}, actual interface{}) bool {
	return IsEqual(expected, actual)
}

// NotEqual defined TODO
func NotEqual(expected interface{}, actual interface{}) bool {
	return !IsEqual(expected, actual)
}
