package api

import (
	"comm/micro/bind"
	go_api "github.com/micro/go-micro/v2/api/proto"
)

// Request defined TODO
type Request struct {
	*go_api.Request
}

func (r *Request) Bind(i interface{}) error {
	return bind.Bind(i, r.Request)
}

func NewRequest(r *go_api.Request) *Request {
	return &Request{Request: r}
}
