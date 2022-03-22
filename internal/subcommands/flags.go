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

	consumerUsername = "username"
	consumerCustomId = "custom_id"

	defaultConsumerUsername = "kong-test-username"
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
			Name:  serviceHost,
			Value: "kong-test.com",
		},
		&cli.IntFlag{
			Name:  servicePort,
			Value: 8080,
		},
	}
	deleteServiceFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  serviceNameOrID,
			Value: "kong-test",
		},
	}
	getServiceFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  serviceNameOrID,
			Value: "kong-test",
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
			Name:  routeProtocols,
			Value: cli.NewStringSlice(defaultRouteProtocols),
		},
		&cli.StringSliceFlag{
			Name:  routePaths,
			Value: cli.NewStringSlice(defaultRoutePaths),
		},
		&cli.StringFlag{
			Name:  routeServiceNameOrID,
			Value: defaultRouteServiceNameOrID,
		},
	}
	deleteRouteFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  serviceNameOrID,
			Value: "kong-test",
		},
	}
	getRouteFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  serviceNameOrID,
			Value: "kong-test",
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
	addConsumerFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  consumerUsername,
			Value: "kong-test",
		},
		&cli.StringFlag{
			Name:  consumerCustomId,
			Value: defaultConsumerUsername,
		},
	}
	deleteConsumerFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  consumerUsername,
			Value: "kong-test",
		},
	}
	getConsumerFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  consumerUsername,
			Value: "kong-test",
		},
	}
)
