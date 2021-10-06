package main

import (
	"aggregate/api/handler"
	"comm/logger"
	"comm/micro"
	"proto/aggregate"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("api.aggregate"))
	micro.RegisterHandler(app.Server(), &handler.Handler{
		AggregateService: aggregate.NewAggregateService(micro.NameFormat("srv.aggregate"), app.Client()),
	})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
