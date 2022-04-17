package consumer

import (
	"fmt"
	kong "github.com/furkansb/kong-cli/internal/kongclient"
	k "github.com/kong/go-kong/kong"
	"github.com/urfave/cli/v2"
)

var kongManager *kong.KongManager

func init() {
	var err error
	kongManager, err = kong.NewKongManager(kong.GetBaseUrl())
	if err != nil {
		panic(err)
	}
}

func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "consumer",
			Usage: "consumer related commands",
			Subcommands: []*cli.Command{
				{
					Name:  "add",
					Usage: "add a new consumer",
					Flags: addConsumerFlags,
					Action: func(c *cli.Context) error {
						return addConsumerCmdFunc(c)
					},
				},
				{
					Name:  "delete",
					Usage: "delete consumer",
					Flags: deleteConsumerFlags,
					Action: func(c *cli.Context) error {
						return deleteConsumerCmdFunc(c)
					},
				},
				{
					Name:  "get",
					Usage: "get consumer",
					Flags: getConsumerFlags,
					Action: func(c *cli.Context) error {
						return getConsumerCmdFunc(c)
					},
				},
				{
					Name:  "list",
					Usage: "list consumers",
					Action: func(c *cli.Context) error {
						return listConsumersCmdFunc(c)
					},
				},
			},
		},
	}
}

func addConsumerCmdFunc(c *cli.Context) error {
	username := c.String("username")
	customID := c.String("custom-id")
	tags := kong.GetSliceElementsPointer(c.StringSlice("tags"))
	consumer := k.Consumer{Username: &username, CustomID: &customID, Tags: tags}
	_, err := kongManager.CreateConsumer(c.Context, &consumer)
	if err != nil {
		return err
	}
	return nil
}

func deleteConsumerCmdFunc(c *cli.Context) error {
	username := c.String("username")
	err := kongManager.DeleteConsumer(c.Context, username)
	if err != nil {
		return err
	}
	return nil
}

func getConsumerCmdFunc(c *cli.Context) error {
	username := c.String("username")
	consumer, err := kongManager.GetConsumer(c.Context, username)
	if err != nil {
		return err
	}
	consumerStr, err := kong.ConsumerString(consumer)
	if err != nil {
		return err
	}
	fmt.Print(consumerStr)
	return nil
}

func listConsumersCmdFunc(c *cli.Context) error {
	consumer, err := kongManager.ListAllConsumers(c.Context)
	if err != nil {
		return err
	}
	for _, c := range consumer {
		consumerStr, err := kong.ConsumerString(c)
		if err != nil {
			return err
		}
		fmt.Print(consumerStr)
	}
	return nil
}
