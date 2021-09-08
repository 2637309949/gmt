package api

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	go_api "github.com/micro/go-micro/v2/api/proto"
)

// Response defined TODO
type Response struct {
	*go_api.Response
}

func (r *Response) Build(data proto.Message, err error) error {
	code := 200
	r.StatusCode = 200
	if err != nil {
		code = 500
	}
	x, err := json.Marshal(map[string]interface{}{
		"code": code,
		"data": data,
	})
	r.Body = string(x)
	return err
}

func NewResponse(r *go_api.Response) *Response {
	return &Response{Response: r}
}
