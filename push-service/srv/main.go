package main

import (
	"proto/push"
	"push/srv/handler"
	"push/srv/publisher"

	"comm/logger"
	"comm/micro"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("srv.push"))
	pClient := micro.NewPublisher(micro.NameFormat("srv.push.t1"), app.Client())
	hdl := handler.Handler{Publisher: publisher.Publisher{Client: pClient}}
	push.RegisterPushHandler(app.Server(), &hdl)
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
