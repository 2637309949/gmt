package main

import (
	"comm/logger"
	"comm/micro"
	"subscriber/srv/subscriber"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("srv.subscriber"))
	micro.RegisterSubscriber(micro.NameFormat("srv.push.t1"), app.Server(), new(subscriber.Subscriber))
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
