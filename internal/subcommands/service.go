package subcommands

import (
	"fmt"

	kong "github.com/furkansb/kong-cli/internal/kongclient"
	k "github.com/kong/go-kong/kong"
	"github.com/urfave/cli/v2"
)

var ServiceCommands []*cli.Command = []*cli.Command{
	{
		Name:  "service",
		Usage: "service related commands",
		Subcommands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a new service",
				Flags: addServiceFlags,
				Action: func(c *cli.Context) error {
					return addServiceCmdFunc(c)
				},
			},
			{
				Name:  "delete",
				Usage: "delete service",
				Flags: deleteServiceFlags,
				Action: func(c *cli.Context) error {
					return deleteServiceCmdFunc(c)
				},
			},
			{
				Name:  "get",
				Usage: "get service",
				Flags: getServiceFlags,
				Action: func(c *cli.Context) error {
					return getServiceCmdFunc(c)
				},
			},
			{
				Name:  "list",
				Usage: "list services",
				Action: func(c *cli.Context) error {
					return listServicesCmdFunc(c)
				},
			},
		},
	},
}

func addServiceCmdFunc(c *cli.Context) error {
	name := c.String("name")
	port := c.Int("port")
	host := c.String("host")
	protocol := c.String("protocol")
	path := kong.StrToPointer(c.String("path"))
	connectTimeout := c.Int("connect-timeout")
	writeTimeout := c.Int("write-timeout")
	readTimeout := c.Int("read-timeout")
	tags := kong.GetSliceElementsPointer(c.StringSlice("tags"))
	clientCertificateID := c.String("client-certificate-id")
	tlsVerify := kong.BoolHandler(c.Bool("tls-verify"))
	tlsVerifyDepth := tlsIntHandler(c.Int("tls-verify-depth"))
	caCertificates := kong.GetSliceElementsPointer(c.StringSlice("ca-certificates"))
	url := kong.StrToPointer(c.String("url"))
	clientCertificate := clientCertFromID(clientCertificateID)
	service := k.Service{Name: &name, Port: &port, Host: &host, Protocol: &protocol, Path: path, ConnectTimeout: &connectTimeout, WriteTimeout: &writeTimeout, ReadTimeout: &readTimeout, Tags: tags, ClientCertificate: clientCertificate, TLSVerify: tlsVerify, TLSVerifyDepth: tlsVerifyDepth, CACertificates: caCertificates, URL: url}
	_, err := kongManager.CreateService(c.Context, &service)
	if err != nil {
		return err
	}
	return nil
}

func deleteServiceCmdFunc(c *cli.Context) error {
	nameOrID := c.String("name-or-id")
	kongManager.DeleteService(c.Context, nameOrID)
	return nil
}

func getServiceCmdFunc(c *cli.Context) error {
	nameOrID := c.String("name-or-id")
	service, err := kongManager.GetService(c.Context, nameOrID)
	if err != nil {
		return err
	}
	fmt.Print(kong.ServiceString(service))
	return nil
}

func listServicesCmdFunc(c *cli.Context) error {
	service, err := kongManager.ListAllServices(c.Context)
	if err != nil {
		return err
	}
	for _, s := range service {
		serviceStr, err := kong.ServiceString(s)
		if err != nil {
			return err
		}
		fmt.Print(serviceStr)
	}
	return nil
}

func tlsIntHandler(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func clientCertFromID(clientCertID string) *k.Certificate {
	var clientCertificate *k.Certificate
	if clientCertID != "" {
		clientCertificate = &k.Certificate{ID: &clientCertID}
	} else {
		clientCertificate = nil
	}
	return clientCertificate
}
