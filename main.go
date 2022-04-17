package main

import (
	"github.com/furkansb/kong-cli/commands/consumer"
	"github.com/furkansb/kong-cli/commands/oauth2"
	"github.com/furkansb/kong-cli/commands/route"
	"github.com/furkansb/kong-cli/commands/service"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func commands() []*cli.Command {
	return []*cli.Command{
		service.Command(),
		route.Command(),
		oauth2.Command(),
		consumer.Command(),
	}
}
