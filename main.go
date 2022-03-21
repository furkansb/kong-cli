package main

import (
	"github.com/furkansb/kong-cli/internal/subcommands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = subcommands.AllCommands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
