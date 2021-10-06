package main

import (
	"proto/publisher"
	"publisher/srv/handler"
	srv "publisher/srv/publisher"

	"comm/logger"
	"comm/micro"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("srv.publisher"))
	pClient := micro.NewPublisher(micro.NameFormat("srv.publisher.t1"), app.Client())
	hdl := handler.Handler{Publisher: srv.Publisher{Client: pClient}}
	publisher.RegisterPushHandler(app.Server(), &hdl)
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
