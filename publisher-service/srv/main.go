package main

import (
	"proto/publisher"
	"publisher-service/srv/handler"
	srv "publisher-service/srv/publisher"

	"comm/logger"
	"comm/micro"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("srv.publisher"))
	pClient := micro.NewPublisher(micro.NameFormat("event.t1"), app.Client())
	publisher.RegisterPushHandler(app.Server(), &handler.Handler{
		Publisher: srv.Publisher{Client: pClient},
	})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
