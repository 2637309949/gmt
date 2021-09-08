package main

import (
	"helloworld/srv/handler"
	"helloworld/srv/subscriber"
	"proto/helloworld"

	"comm/logger"
	"comm/micro"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("srv.helloworld"))
	helloworld.RegisterHelloworldHandler(app.Server(), new(handler.Handler))
	micro.RegisterSubscriber(micro.NameFormat("srv.helloworld"), app.Server(), new(subscriber.Subscriber))
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
