package main

import (
	"github.com/klivohealth/kli/command"
	"github.com/yitsushi/go-commander"
)

func registerCommands(registry *commander.CommandRegistry) {
	registry.Register(command.NewHelloWorld)
}

func main() {
	registry := commander.NewCommandRegistry()

	registerCommands(registry)

	registry.Execute()
}
