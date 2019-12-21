package command

import (
	"fmt"
	"github.com/yitsushi/go-commander"
)

type HelloWorld struct {
}

func (c *HelloWorld) Execute(opts *commander.CommandHelper) {
	fmt.Println("Hello World", opts.Arg(0))
}

func NewHelloWorld(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &HelloWorld{},
		Help: &commander.CommandDescriptor{
			Name:             "helloworld",
			ShortDescription: "Say Hello World",
			Arguments:        "[name]",
			Examples: []string{
				"",
				"name",
			},
		},
	}
}
