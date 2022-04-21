package route

import (
	"github.com/urfave/cli/v2"
)

const (
	name                    = "name"
	protocols               = "protocols"
	methods                 = "methods"
	hosts                   = "hosts"
	paths                   = "paths"
	headers                 = "headers"
	httpsRedirectStatusCode = "https-redirect-status-code"
	regexPriority           = "regex-priority"
	stripPath               = "strip-path"
	pathHandling            = "path-handling"
	preserveHost            = "preserve-host"
	requestBuffering        = "request-buffering"
	responseBuffering       = "response-buffering"
	snis                    = "snis"
	sources                 = "sources"
	destinations            = "destinations"
	tags                    = "tags"
	serviceID               = "service-id"

	defaultRouteHosts     = "kong-test.com"
	defaultRouteProtocols = "http"
	defaultRoutePaths     = "/api/v1"
)

// Route cmd flags
var (
	addRouteFlags = []cli.Flag{
		&cli.StringFlag{
			Name: name,
		},
		&cli.StringSliceFlag{
			Name:     protocols,
			Required: true,
		},
		&cli.StringSliceFlag{
			Name: methods,
		},
		&cli.StringSliceFlag{
			Name: hosts,
		},
		&cli.StringSliceFlag{
			Name:     paths,
			Required: true,
		},
		&cli.StringSliceFlag{
			Name: headers,
		},
		&cli.IntFlag{
			Name: httpsRedirectStatusCode,
		},
		&cli.IntFlag{
			Name: regexPriority,
		},
		&cli.BoolFlag{
			Name: stripPath,
		},
		&cli.StringFlag{
			Name: pathHandling,
		},
		&cli.BoolFlag{
			Name: preserveHost,
		},
		&cli.BoolFlag{
			Name: requestBuffering,
		},
		&cli.BoolFlag{
			Name: responseBuffering,
		},
		&cli.StringSliceFlag{
			Name: tags,
		},
		&cli.StringFlag{
			Name:     serviceID,
			Usage:    "The service id or name to which the route belongs to",
			Required: true,
		},
	}
	deleteRouteFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     name,
			Usage:    "The ID or name of the route to delete",
			Required: true,
		},
	}
	getRouteFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     name,
			Usage:    "The ID or name of the route to delete",
			Required: true,
		},
	}
	// listRouteForServiceFlags = []cli.Flag{
	// 	&cli.StringFlag{
	// 		Name:  serviceNameOrID,
	// 		Value: "kong-test",
	// 	},
	// }
)
