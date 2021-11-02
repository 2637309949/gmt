package main

import (
	"{{.name}}-service/srv/handler"
	"proto/{{.name}}"
	"context"
	"comm/logger"
	"comm/micro"
)

func main() {
	app, ctx := micro.NewServiceWithName(micro.NameFormat("srv.{{.name}}")), context.Background()
	{{.name}}.Register{{toTitle .name}}Handler(app.Server(), &handler.Handler{})
	if err := app.Run(); err != nil {
		logger.Fatal(ctx, err)
	}
}
