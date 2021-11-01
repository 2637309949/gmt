package main

import (
	"github.com/2637309949/gmt/comm/autogen/cmd"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
