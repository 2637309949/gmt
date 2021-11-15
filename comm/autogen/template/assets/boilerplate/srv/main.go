package main

import (
	"{{.name}}-service/srv/handler"
	"proto/{{.name}}"
	"context"
	"comm/logger"
	"comm/micro"
)

func main() {
	ctx := context.Background()
	app := micro.NewServiceWithName(micro.NameFormat("srv.{{.name}}"))
	{{.name}}.Register{{toTitle .name}}Handler(app.Server(), &handler.Handler{})
	if err := app.Run(); err != nil {
		logger.Fatal(ctx, err)
	}
}
