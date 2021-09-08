package main

import (
	"comm/logger"
	"comm/micro"
	"helloworld/api/handler"
	"proto/helloworld"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("api.helloworld"))
	micro.RegisterHandler(app.Server(), &handler.Handler{
		HelloworldEvent:  micro.NewEvent(micro.NameFormat("srv.helloworld"), app.Client()),
		HelloworldClient: helloworld.NewHelloworldService(micro.NameFormat("srv.helloworld"), app.Client()),
	})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
