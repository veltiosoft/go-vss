package vss

import (
	"log"

	"github.com/mitchellh/cli"
	"github.com/vssio/go-vss/internal/command"
)

const version = "0.1.0"

func Version() string {
	return version
}

func Run(args []string) int {
	metaPtr := new(command.Meta)
	c := &cli.CLI{
		Name:         "vss",
		Version:      version,
		Args:         args,
		Autocomplete: true,
		Commands:     initCommands(metaPtr),
	}
	exitCode, err := c.Run()
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	return exitCode
}
