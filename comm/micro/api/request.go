package api

import (
	"comm/micro/bind"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

// Request defined TODO
type Request struct {
	*go_api.Request
	User string
}

func (r *Request) Bind(i interface{}) error {
	return bind.Bind(i, r.Request)
}

func NewRequest(r *go_api.Request) *Request {
	rr := Request{Request: r}
	xUser := r.GetHeader()["x-user"]
	if len(xUser.GetValues()) > 0 {
		rr.User = xUser.GetValues()[0]
	}
	return &rr
}
