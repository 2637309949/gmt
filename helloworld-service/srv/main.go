package main

import (
	"context"
	"helloworld-service/srv/handler"
	"proto/helloworld"

	"comm/logger"
	"comm/micro"
)

func main() {
	ctx := context.Background()
	app := micro.NewServiceWithName(micro.NameFormat("srv.helloworld"))
	helloworld.RegisterHelloworldHandler(app.Server(), &handler.Handler{})
	if err := app.Run(); err != nil {
		logger.Fatal(ctx, err)
	}
}
