package main

import (
	"github.com/2637309949/gmt/comm/gmt/cmd"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
