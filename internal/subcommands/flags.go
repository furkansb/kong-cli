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

	routeName                    = "name"
	routeProtocols               = "protocols"
	routeMethods                 = "methods"
	routeHosts                   = "hosts"
	routePaths                   = "paths"
	routeHeaders                 = "headers"
	routeHttpsRedirectStatusCode = "https-redirect-status-code"
	routeRegexPriority           = "regex-priority"
	routeStripPath               = "strip-path"
	routePathHandling            = "path-handling"
	routePreserveHost            = "preserve-host"
	routeRequestBuffering        = "request-buffering"
	routeResponseBuffering       = "response-buffering"
	routeSnis                    = "snis"
	routeSources                 = "sources"
	routeDestinations            = "destinations"
	routeTags                    = "tags"
	routeServiceID               = "service-id"

	defaultRouteHosts     = "kong-test.com"
	defaultRouteProtocols = "http"
	defaultRoutePaths     = "/api/v1"

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
			Name:  serviceTlsVerify,
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
		&cli.StringFlag{
			Name: routeName,
		},
		&cli.StringSliceFlag{
			Name:     routeProtocols,
			Value:    cli.NewStringSlice(defaultRouteProtocols),
			Required: true,
		},
		&cli.StringSliceFlag{
			Name: routeMethods,
		},
		&cli.StringSliceFlag{
			Name:  routeHosts,
			Value: cli.NewStringSlice(defaultRouteHosts),
		},
		&cli.StringSliceFlag{
			Name:     routePaths,
			Value:    cli.NewStringSlice(defaultRoutePaths),
			Required: true,
		},
		&cli.StringSliceFlag{
			Name: routeHeaders,
		},
		&cli.IntFlag{
			Name: routeHttpsRedirectStatusCode,
		},
		&cli.IntFlag{
			Name: routeRegexPriority,
		},
		&cli.BoolFlag{
			Name: routeStripPath,
		},
		&cli.StringFlag{
			Name: routePathHandling,
		},
		&cli.BoolFlag{
			Name: routePreserveHost,
		},
		&cli.BoolFlag{
			Name: routeRequestBuffering,
		},
		&cli.BoolFlag{
			Name: routeResponseBuffering,
		},
		&cli.StringSliceFlag{
			Name: routeTags,
		},
		&cli.StringFlag{
			Name:     routeServiceID,
			Usage:    "The service id or name to which the route belongs to",
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
