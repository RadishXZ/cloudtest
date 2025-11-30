package main

import (
	"os"

	"github.com/RadishXZ/cloudtest/cmd/fg-apiserver/app"
	_ "go.ubser.org/automaxprocs"
)

func main() {
	command := app.NewCloudTestGoCommand()
	
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}


