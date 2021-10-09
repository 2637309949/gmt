package main

import (
	"comm/logger"
	"comm/micro"
	"publisher/api/handler"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("api.publisher"))
	micro.RegisterHandler(app.Server(), &handler.Handler{})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
