package main

import (
	"comm/gmt/cmd"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
