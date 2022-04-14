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
					return addRouteCmdFunc(c)
				},
			},
			{
				Name:  "delete",
				Usage: "delete route",
				Flags: deleteRouteFlags,
				Action: func(c *cli.Context) error {
					return deleteRouteCmdFunc(c)
				},
			},
			{
				Name:  "get",
				Usage: "get route",
				Flags: getRouteFlags,
				Action: func(c *cli.Context) error {
					return getRouteCmdFunc(c)
				},
			},
			{
				Name:  "list",
				Usage: "list routes",
				Action: func(c *cli.Context) error {
					return listRoutesCmdFunc(c)
				},
			},
		},
	},
}

func addRouteCmdFunc(c *cli.Context) error {
	hosts := c.StringSlice("hosts")
	paths := c.StringSlice("paths")
	protocols := c.StringSlice("protocols")
	serviceNameOrID := c.String("service-id")
	service, err := kongManager.GetService(c.Context, serviceNameOrID)
	if err != nil {
		return err
	}
	route := k.Route{Hosts: getSliceElementsPointer(hosts), Paths: getSliceElementsPointer(paths), Protocols: getSliceElementsPointer(protocols), Service: service}
	_, err = kongManager.CreateRoute(c.Context, &route)
	if err != nil {
		return err
	}
	return nil
}

func deleteRouteCmdFunc(c *cli.Context) error {
	nameOrID := c.String("name-or-id")
	kongManager.DeleteRoute(c.Context, nameOrID)
	return nil
}

func getRouteCmdFunc(c *cli.Context) error {
	nameOrID := c.String("name-or-id")
	route, err := kongManager.GetRoute(c.Context, nameOrID)
	if err != nil {
		return err
	}
	fmt.Print(kong.RouteString(route))
	return nil
}

func listRoutesCmdFunc(c *cli.Context) error {
	routes, err := kongManager.ListAllRoutes(c.Context)
	if err != nil {
		return err
	}
	for _, r := range routes {
		routeStr, err := kong.RouteString(r)
		if err != nil {
			return err
		}
		fmt.Print(routeStr)
	}
	return nil
}
