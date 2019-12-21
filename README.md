kli - Klivo Command Line Interface (CLI)

# Intro

kli is Klivo's official CLI and it uses the Go Library [go-commander](https://github.com/yitsushi/go-commander) to simplify the CLI workflow.

# How to extend it

kli is very simple. All you have to do to add a new command is:

1. Add a file inside the command directory
2. Register your newly created file in the func `registerCommands` inside `main.go`

Run `go build` and your kli should be up to date.
