package subcommands

import (
	"fmt"
	kong "github.com/furkansb/kong-cli/internal/kongclient"
	k "github.com/kong/go-kong/kong"
	"github.com/urfave/cli/v2"
)

var ConsumerCommands []*cli.Command = []*cli.Command{
	{
		Name:  "consumer",
		Usage: "consumer related commands",
		Subcommands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a new consumer",
				Flags: addConsumerFlags,
				Action: func(c *cli.Context) error {
					username := c.String("username")
					customID := c.String("custom_id")
					consumer := k.Consumer{Username: &username, CustomID: &customID}
					kongManager.CreateConsumer(c.Context, &consumer)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "delete consumer",
				Flags: deleteConsumerFlags,
				Action: func(c *cli.Context) error {
					username := c.String("username")
					kongManager.DeleteConsumer(c.Context, username)
					return nil
				},
			},
			{
				Name:  "get",
				Usage: "get consumer",
				Flags: getConsumerFlags,
				Action: func(c *cli.Context) error {
					username := c.String("username")
					consumer := kongManager.GetConsumer(c.Context, username)
					fmt.Print(kong.ConsumerString(consumer))
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "list consumers",
				Action: func(c *cli.Context) error {
					consumer := kongManager.ListAllConsumers(c.Context)
					for _, s := range consumer {
						fmt.Print(kong.ConsumerString(s))
					}
					return nil
				},
			},
		},
	},
}
