package main

import (
	"github.com/furkansb/kong-cli/commands/service"
	"github.com/furkansb/kong-cli/commands/route"
	"github.com/furkansb/kong-cli/commands/oauth2"
	"github.com/furkansb/kong-cli/commands/consumer"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = allCommands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mergeCommands(commands ...[]*cli.Command) []*cli.Command {
	var result []*cli.Command
	for _, f := range commands {
		result = append(result, f...)
	}
	return result
}

func allCommands() []*cli.Command {
	return mergeCommands(service.Commands(), route.Commands(), oauth2.Commands(), consumer.Commands())
}