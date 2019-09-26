package main

import (
	"log"
	"os"

	"github.com/digitalonus/venera/commands"
	"github.com/mitchellh/cli"
)

func main() {
	cmd := cli.NewCLI("venera", "mars")
	cmd.Args = os.Args[1:]

	cmd.Commands = map[string]cli.CommandFactory{
		"info": func() (cli.Command, error) {
			return &commands.InfoCommand{}, nil
		},
	}

	rc, err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(rc)
}
