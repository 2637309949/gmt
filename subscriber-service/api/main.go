package main

import (
	"comm/logger"
	"comm/micro"
)

func main() {
	app := micro.NewServiceWithName(micro.NameFormat("api.subscriber"))
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
