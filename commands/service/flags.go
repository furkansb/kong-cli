package service

import (
	"github.com/urfave/cli/v2"
)

const (
	name              = "name"
	protocol          = "protocol"
	host              = "host"
	port              = "port"
	retries           = "retries"
	connectTimeout    = "connect-timeout"
	writeTimeout      = "write-timeout"
	readTimeout       = "read-timeout"
	tags              = "tags"
	clientCertificateID = "client-certificate-id"
	tlsVerify         = "tls-verify"
	tlsVerifyDepth    = "tls-verify-depth"
	caCertificates    = "ca-certificates"
	url               = "url"
	path              = "path"
)

// Service cmd flags
var (
	addServiceFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  name,
		},
		&cli.StringFlag{
			Name:  protocol,
			Required: true,
		},
		&cli.StringFlag{
			Name:     host,
			Required: true,
		},
		&cli.StringFlag{
			Name: path,
		},
		&cli.IntFlag{
			Name:     port,
			Required: true,
		},
		&cli.IntFlag{
			Name:  retries,
		},
		&cli.IntFlag{
			Name:  connectTimeout,
		},
		&cli.IntFlag{
			Name:  writeTimeout,
		},
		&cli.IntFlag{
			Name:  readTimeout,
		},
		&cli.StringSliceFlag{
			Name: tags,
		},
		&cli.StringFlag{
			Name: clientCertificateID,
		},
		&cli.BoolFlag{
			Name:  tlsVerify,
		},
		&cli.StringFlag{
			Name: tlsVerifyDepth,
		},
		&cli.StringSliceFlag{
			Name: caCertificates,
		},
		&cli.StringFlag{
			Name: url,
		},
	}
	deleteServiceFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     name,
			Required: true,
		},
	}
	getServiceFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     name,
			Required: true,
			Usage: "The service name or id",
		},
	}
)