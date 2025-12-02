package main

import (
	"os"

	"github.com/RadishXZ/cloudtest/cmd/fg-apiserver/app"
	_ "go.uber.org/automaxprocs"
)

func main() {
	command := app.NewCloudTestGOCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}