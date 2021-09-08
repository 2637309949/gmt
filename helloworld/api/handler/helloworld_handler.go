package handler

import (
	"comm/go-micro"
	"comm/go-micro/logger"
	"context"
	"encoding/json"
	"proto/helloworld"
	"strings"

	go_api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
)

func (h *Handler) Call(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	logger.Info("Received api.helloworld.call API request")
	logger.Info("name = ", micro.Key("name").String())
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("api.helloworld.call", "Name cannot be blank")
	}

	response, err := h.HelloworldClient.Call(ctx, &helloworld.Request{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}
