package main

import (
	"helloworld/srv/handler"
	"helloworld/srv/subscriber"
	"proto/helloworld"

	"comm/go-micro"
	"comm/go-micro/logger"
)

func main() {
	app := micro.NewServiceWithName("srv.helloworld")
	helloworld.RegisterHelloworldHandler(app.Server(), new(handler.Handler))
	micro.RegisterSubscriber("srv.helloworld", app.Server(), new(subscriber.Subscriber))
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
