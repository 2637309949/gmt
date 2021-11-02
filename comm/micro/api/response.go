package api

import (
	"comm/errors"
	"context"
	"encoding/json"

	go_api "github.com/micro/go-micro/v2/api/proto"
)

// Response defined TODO
type Response struct {
	*go_api.Response
}

// Message defined TODO
type Message interface {
	Reset()
	String() string
	ProtoMessage()
}

func (r *Response) Build(ctx context.Context, data Message, err error) string {
	r.StatusCode = 200
	if err != nil {
		r.Body = errors.Format(err)
		return r.Body
	}
	rspByte, _ := json.Marshal(map[string]interface{}{"code": 200, "data": data})
	r.Body = string(rspByte)
	return r.Body
}

func NewResponse(r *go_api.Response) *Response {
	return &Response{Response: r}
}
