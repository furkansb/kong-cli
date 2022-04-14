package subcommands

import (
	"github.com/urfave/cli/v2"
)

const (
	serviceName     = "name"
	serviceProtocol = "protocol"
	serviceHost     = "host"
	servicePort     = "port"
	serviceNameOrID = "name-or-id"

	routeServiceNameOrID = "service-id"
	routeProtocols       = "protocols"
	routeHosts           = "hosts"
	routePaths           = "paths"

	defaultRouteHosts           = "kong-test.com"
	defaultRouteProtocols       = "http"
	defaultRoutePaths           = "/api/v1"
	defaultRouteServiceNameOrID = "kong-test-service"
)

var (
	addServiceFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  serviceName,
			Value: "kong-test",
		},
		&cli.StringFlag{
			Name:  serviceProtocol,
			Value: "http",
		},
		&cli.StringFlag{
			Name:     serviceHost,
			Value:    "kong-test.com",
			Required: true,
		},
		&cli.IntFlag{
			Name:     servicePort,
			Value:    8080,
			Required: true,
		},
	}
	deleteServiceFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     serviceNameOrID,
			Value:    "kong-test",
			Required: true,
		},
	}
	getServiceFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     serviceNameOrID,
			Value:    "kong-test",
			Required: true,
		},
	}
)

var (
	addRouteFlags = []cli.Flag{
		&cli.StringSliceFlag{
			Name:  routeHosts,
			Value: cli.NewStringSlice(defaultRouteHosts),
		},
		&cli.StringSliceFlag{
			Name:     routeProtocols,
			Value:    cli.NewStringSlice(defaultRouteProtocols),
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     routePaths,
			Value:    cli.NewStringSlice(defaultRoutePaths),
			Required: true,
		},
		&cli.StringFlag{
			Name:     routeServiceNameOrID,
			Value:    defaultRouteServiceNameOrID,
			Required: true,
		},
	}
	deleteRouteFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     serviceNameOrID,
			Value:    "kong-test",
			Required: true,
		},
	}
	getRouteFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     serviceNameOrID,
			Value:    "kong-test",
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

var (
	addOauth2Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "consumer-name-or-id",
			Value: "kong-test",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "name",
			Value: "kong-app",
		},
		&cli.StringFlag{
			Name:  "client-id",
			Value: "kong-test-user",
		},
		&cli.StringFlag{
			Name:  "client-secret",
			Value: "kong-test-secret",
		},
	}
	deleteOauth2Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "consumer-name-or-id",
			Value: "kong-test",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "client-id-or-id",
			Value: "kong-test-user",
			Required: true,
		},
	}
	getOauth2Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "consumer-name-or-id",
			Value: "kong-test",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "client-id-or-id",
			Value: "kong-test-user",
			Required: true,
		},
	}
)
