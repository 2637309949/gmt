package main

import (
	"helloworld/srv/handler"
	"proto/helloworld"

	"comm/logger"
	"comm/micro"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("srv.helloworld"))
	helloworld.RegisterHelloworldHandler(app.Server(), &handler.Handler{})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
