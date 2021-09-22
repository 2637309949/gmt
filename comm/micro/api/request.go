package api

import (
	"comm/micro/bind"
	"reflect"
	"sync"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

// Request defined TODO
type Request struct {
	*go_api.Request
	Keys map[string]interface{}
	mu   sync.RWMutex
}

// Bind defined TODO
func (r *Request) Bind(i interface{}) error {
	err := bind.Bind(i, r.Request)
	if err != nil {
		return err
	}
	iv := reflect.ValueOf(i)
	iv = reflect.Indirect(iv)
	xi := iv.FieldByName("XUser")
	if xi.CanSet() && xi.IsValid() && xi.Type().Kind() == reflect.String {
		xUser, _ := r.Get("X-User")
		xi.Set(reflect.ValueOf(xUser))
	}
	for k, v := range r.Keys {
		if k == "X-User" {
			continue
		}
		xi := iv.FieldByName(k)
		if xi.CanSet() && xi.IsValid() && xi.IsZero() && xi.Type().Kind() == reflect.TypeOf(v).Kind() {
			xi.Set(reflect.ValueOf(v))
		}
	}
	return nil
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (r *Request) Get(key string) (value interface{}, exists bool) {
	r.mu.RLock()
	value, exists = r.Keys[key]
	r.mu.RUnlock()
	return
}

// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
func (r *Request) Set(key string, value interface{}) {
	r.mu.Lock()
	if r.Keys == nil {
		r.Keys = make(map[string]interface{})
	}
	r.Keys[key] = value
	r.mu.Unlock()
}

// NewRequest defined TODO
func NewRequest(r *go_api.Request) *Request {
	rr := Request{Request: r}
	xUser := r.GetHeader()["X-User"]
	if len(xUser.GetValues()) > 0 {
		rr.Set("X-User", xUser.GetValues()[0])
	}
	return &rr
}
