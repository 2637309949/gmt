package main

import (
	"comm/logger"
	"comm/micro"
	"subscriber/srv/subscriber"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("srv.subscriber"))
	micro.RegisterSubscriber(micro.NameFormat("event.t1"), app.Server(), &subscriber.Subscriber{})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
