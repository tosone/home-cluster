package main

import (
	"home-cluster/cmd/commands"

	"github.com/sirupsen/logrus"
)

func main() {
	err := commands.Execute()
	if err != nil {
		logrus.Error(err)
	}
}
