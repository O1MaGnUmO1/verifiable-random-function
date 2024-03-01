package main

import (
	"vrf/application"

	"github.com/sirupsen/logrus"
)

func main() {
	app, err := application.NewApplication()
	if err != nil {
		logrus.Fatal("failed to load application : reason %v", err)
	}
	app.Run()
}
