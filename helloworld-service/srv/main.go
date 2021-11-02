package main

import (
	"context"
	"helloworld-service/srv/handler"
	"proto/helloworld"

	"comm/logger"
	"comm/micro"
)

func main() {
	app, ctx := micro.NewServiceWithName(micro.NameFormat("srv.helloworld")), context.Background()
	helloworld.RegisterHelloworldHandler(app.Server(), &handler.Handler{})
	if err := app.Run(); err != nil {
		logger.Fatal(ctx, err)
	}
}
