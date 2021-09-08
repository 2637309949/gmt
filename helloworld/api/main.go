package main

import (
	"comm/go-micro"
	"comm/go-micro/logger"
	"helloworld/api/handler"
	"proto/helloworld"
)

func main() {
	app := micro.NewServiceWithName("api.helloworld")
	micro.RegisterHandler(app.Server(), &handler.Handler{
		HelloworldClient: helloworld.NewHelloworldService(micro.NameFormat("srv.helloworld"), app.Client()),
	})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
