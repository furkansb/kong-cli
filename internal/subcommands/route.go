package subcommands

import (
	"fmt"

	kong "github.com/furkansb/kong-cli/internal/kongclient"
	k "github.com/kong/go-kong/kong"
	"github.com/urfave/cli/v2"
)

var RouteCommands []*cli.Command = []*cli.Command{
	{
		Name:  "route",
		Usage: "route related commands",
		Subcommands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a new route",
				Flags: addRouteFlags,
				Action: func(c *cli.Context) error {
					hosts := c.StringSlice("hosts")
					paths := c.StringSlice("paths")
					protocols := c.StringSlice("protocols")
					serviceNameOrID := c.String("service-id")
					service := kongManager.GetService(c.Context, serviceNameOrID)
					route := k.Route{Hosts: getSliceElementsPointer(hosts), Paths: getSliceElementsPointer(paths), Protocols: getSliceElementsPointer(protocols), Service: service}
					kongManager.CreateRoute(c.Context, &route)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "delete route",
				Flags: deleteRouteFlags,
				Action: func(c *cli.Context) error {
					nameOrID := c.String("name-or-id")
					kongManager.DeleteRoute(c.Context, nameOrID)
					return nil
				},
			},
			{
				Name:  "get",
				Usage: "get route",
				Flags: getRouteFlags,
				Action: func(c *cli.Context) error {
					nameOrID := c.String("name-or-id")
					route := kongManager.GetRoute(c.Context, nameOrID)
					fmt.Print(kong.RouteString(route))
					return c.Err()
				},
			},
			{
				Name:  "list",
				Usage: "list routes",
				Action: func(c *cli.Context) error {
					routes := kongManager.ListAllRoutes(c.Context)
					for _, r := range routes {
						fmt.Print(kong.RouteString(r))
					}
					return nil
				},
			},
		},
	},
}
