package consumer

import (
	"github.com/urfave/cli/v2"
)

const (
	consumerUsername = "username"
	consumerCustomId = "custom-id"
	consumerTags     = "tags"
)

var (
	addConsumerFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     consumerUsername,
			Value:    "kong-test",
			Required: true,
		},
		&cli.StringFlag{
			Name:     consumerCustomId,
			Value:    "kong-test-username",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name: consumerTags,
		},
	}
	deleteConsumerFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     consumerUsername,
			Value:    "kong-test",
			Required: true,
		},
	}
	getConsumerFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     consumerUsername,
			Value:    "kong-test",
			Required: true,
		},
	}
)
