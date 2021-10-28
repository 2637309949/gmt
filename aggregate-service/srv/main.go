package main

import (
	"aggregate-service/srv/handler"
	"comm/logger"
	"comm/micro"
	"proto/aggregate"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("srv.aggregate"))
	aggregate.RegisterAggregateHandler(app.Server(), &handler.Handler{})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
