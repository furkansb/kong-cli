package oauth2

import (
	"github.com/urfave/cli/v2"
)

const (
	consumerUsername = "consumer-username"
	name             = "name"
	clientID         = "client-id"
	clientSecret     = "client-secret"
)

// Oauth2 cmd flags
var (
	addOauth2Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     consumerUsername,
			Usage:    "consumer username or id",
			Required: true,
		},
		&cli.StringFlag{
			Name:  name,
			Usage: "oauth2 name or id",
		},
		&cli.StringFlag{
			Name: clientID,
		},
		&cli.StringFlag{
			Name: clientSecret,
		},
	}
	deleteOauth2Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     consumerUsername,
			Usage:    "consumer username or id",
			Required: true,
		},
		&cli.StringFlag{
			Name:     clientID,
			Usage:    "oauth2 client id or id",
			Required: true,
		},
	}
	getOauth2Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     consumerUsername,
			Usage:    "consumer username or id",
			Required: true,
		},
		&cli.StringFlag{
			Name:     clientID,
			Usage:    "oauth2 client id or id",
			Required: true,
		},
	}
)
