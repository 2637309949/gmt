package main

import (
	"comm/logger"
	"comm/micro"
	"context"
	"helloworld-service/api/handler"
	"proto/helloworld"
)

func main() {
	app, ctx := micro.NewServiceWithName(micro.NameFormat("api.helloworld")), context.Background()
	micro.RegisterHandler(app.Server(), &handler.Handler{
		HelloworldService: helloworld.NewHelloworldService(micro.NameFormat("srv.helloworld"), app.Client()),
	})
	if err := app.Run(); err != nil {
		logger.Fatal(ctx, err)
	}
}
