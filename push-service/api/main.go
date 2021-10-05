package main

import (
	"comm/logger"
	"comm/micro"
	"push/api/handler"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("api.push"))
	micro.RegisterHandler(app.Server(), &handler.Handler{})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
