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
					name := c.String("name")
					port := c.Int("port")
					host := c.String("host")
					protocol := c.String("protocol")
					service := k.Service{Name: &name, Port: &port, Host: &host, Protocol: &protocol}
					kongManager.CreateService(c.Context, &service)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "delete service",
				Flags: deleteServiceFlags,
				Action: func(c *cli.Context) error {
					nameOrID := c.String("name-or-id")
					kongManager.DeleteService(c.Context, nameOrID)
					return nil
				},
			},
			{
				Name:  "get",
				Usage: "get service",
				Flags: getServiceFlags,
				Action: func(c *cli.Context) error {
					nameOrID := c.String("name-or-id")
					service := kongManager.GetService(c.Context, nameOrID)
					fmt.Print(kong.ServiceString(service))
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "list services",
				Action: func(c *cli.Context) error {
					service := kongManager.ListAllServices(c.Context)
					for _, s := range service {
						fmt.Print(kong.ServiceString(s))
					}
					return nil
				},
			},
		},
	},
}
