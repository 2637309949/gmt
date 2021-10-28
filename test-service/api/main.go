package main

import (
	"comm/logger"
	"comm/micro"
	"test-service/api/handler"
	"proto/test"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("api.test"))
	micro.RegisterHandler(app.Server(), &handler.Handler{
		TestService: test.NewTestService(micro.NameFormat("srv.test"), app.Client()),
	})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
