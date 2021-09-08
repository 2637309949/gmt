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
	var x []byte
	if err != nil {
		x, _ = json.Marshal(map[string]interface{}{
			"code": 500,
			"msg":  err.Error(),
		})
	} else {
		x, _ = json.Marshal(map[string]interface{}{
			"code": 200,
			"data": data,
		})
	}
	r.StatusCode = 200
	r.Body = string(x)
	return err
}

func NewResponse(r *go_api.Response) *Response {
	return &Response{Response: r}
}
