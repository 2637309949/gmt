package main

import (
	"comm/logger"
	"comm/micro"
	"{{.name}}/api/handler"
	"proto/{{.name}}"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("api.{{.name}}"))
	micro.RegisterHandler(app.Server(), &handler.Handler{
		HelloworldService: helloworld.NewHelloworldService(micro.NameFormat("srv.{{.name}}"), app.Client()),
	})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
