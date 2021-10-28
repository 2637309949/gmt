package main

import (
	"comm/logger"
	"comm/micro"
	"{{.name}}-service/api/handler"
	"proto/{{.name}}"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("api.{{.name}}"))
	micro.RegisterHandler(app.Server(), &handler.Handler{
		{{toTitle .name}}Service: {{.name}}.New{{toTitle .name}}Service(micro.NameFormat("srv.{{.name}}"), app.Client()),
	})
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
