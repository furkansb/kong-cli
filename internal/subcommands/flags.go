package subcommands

import (
	"github.com/urfave/cli/v2"
)

const (
	serviceName              = "name"
	serviceProtocol          = "protocol"
	serviceHost              = "host"
	servicePort              = "port"
	serviceNameOrID          = "name-or-id"
	serviceRetries           = "retries"
	serviceConnectTimeout    = "connect-timeout"
	serviceWriteTimeout      = "write-timeout"
	serviceReadTimeout       = "read-timeout"
	serviceTags              = "tags"
	serviceClientCertificate = "client-certificate-id"
	serviceTlsVerify         = "tls-verify"
	serviceTlsVerifyDepth    = "tls-verify-depth"
	serviceCaCertificates    = "ca-certificates"
	serviceUrl               = "url"
	servicePath              = "path"

	routeServiceNameOrID = "service-id"
	routeProtocols       = "protocols"
	routeHosts           = "hosts"
	routePaths           = "paths"

	defaultRouteHosts           = "kong-test.com"
	defaultRouteProtocols       = "http"
	defaultRoutePaths           = "/api/v1"
	defaultRouteServiceNameOrID = "kong-test-service"

	consumerUsername = "username"
	consumerCustomId = "custom-id"

	defaultConsumerUsername = "kong-test-username"
)

// Service cmd flags
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
		&cli.StringFlag{
			Name: servicePath,
		},
		&cli.IntFlag{
			Name:     servicePort,
			Value:    8080,
			Required: true,
		},
		&cli.IntFlag{
			Name:  serviceRetries,
			Value: 5,
		},
		&cli.IntFlag{
			Name:  serviceConnectTimeout,
			Value: 60000,
		},
		&cli.IntFlag{
			Name:  serviceWriteTimeout,
			Value: 60000,
		},
		&cli.IntFlag{
			Name:  serviceReadTimeout,
			Value: 60000,
		},
		&cli.StringSliceFlag{
			Name: serviceTags,
		},
		&cli.StringFlag{
			Name: serviceClientCertificate,
		},
		&cli.BoolFlag{
			Name: serviceTlsVerify,
			Value: false,
		},
		&cli.StringFlag{
			Name: serviceTlsVerifyDepth,
		},
		&cli.StringSliceFlag{
			Name: serviceCaCertificates,
		},
		&cli.StringFlag{
			Name: serviceUrl,
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

// Route cmd flags
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
	addConsumerFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     consumerUsername,
			Value:    "kong-test",
			Required: true,
		},
		&cli.StringFlag{
			Name:     consumerCustomId,
			Value:    defaultConsumerUsername,
			Required: true,
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

// Oauth2 cmd flags
var (
	addOauth2Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "consumer-name-or-id",
			Value:    "kong-test",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "name",
			Value: "kong-app",
		},
		&cli.StringFlag{
			Name: "client-id",
		},
		&cli.StringFlag{
			Name: "client-secret",
		},
	}
	deleteOauth2Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "consumer-name-or-id",
			Value:    "kong-test",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "client-id-or-id",
			Value:    "kong-test-user",
			Required: true,
		},
	}
	getOauth2Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "consumer-name-or-id",
			Value:    "kong-test",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "client-id-or-id",
			Value:    "kong-test-user",
			Required: true,
		},
	}
)
