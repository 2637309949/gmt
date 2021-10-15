package main

import (
	"{{.name}}/srv/handler"
	"proto/{{.name}}"

	"comm/logger"
	"comm/micro"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("srv.{{.name}}"))
	helloworld.RegisterHelloworldHandler(app.Server(), &handler.Handler{})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
