package route

import (
	"fmt"

	kong "github.com/furkansb/kong-cli/internal/kongclient"
	k "github.com/kong/go-kong/kong"
	"github.com/urfave/cli/v2"
)

var kongManager *kong.KongManager

func init() {
	var err error
	kongManager, err = kong.NewKongManager(kong.GetBaseUrl())
	if err != nil {
		panic(err)
	}
}

func Command() *cli.Command {
	return &cli.Command{
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
				Flags: listRouteFlags,
				Action: func(c *cli.Context) error {
					return listRoutesCmdFunc(c)
				},
			},
		},
	}
}

// TODO: support headers with flags
func addRouteCmdFunc(c *cli.Context) error {
	name := kong.StrToPointer(c.String("name"))
	hosts := kong.PsliceFromStrSlice(c.StringSlice("hosts"))
	protocols := kong.PsliceFromStrSlice(c.StringSlice("protocols"))
	methods := kong.PsliceFromStrSlice(c.StringSlice("methods"))
	paths := kong.PsliceFromStrSlice(c.StringSlice("paths"))
	httpsRedirectStatusCode := c.Int("https-redirect-status-code")
	regexPriority := c.Int("regex-priority")
	stripPath := kong.BoolHandler(c.Bool("strip-path"))
	preserveHost := kong.BoolHandler(c.Bool("preserve-host"))
	requestBuffering := kong.BoolHandler(c.Bool("request-buffering"))
	responseBuffering := kong.BoolHandler(c.Bool("response-buffering"))
	tags := kong.PsliceFromStrSlice(c.StringSlice("tags"))
	serviceID := c.String("service-id")
	service, err := kongManager.GetService(c.Context, serviceID)
	if err != nil {
		return err
	}
	route := k.Route{Name: name, Hosts: hosts, Protocols: protocols, Methods: methods, Paths: paths, HTTPSRedirectStatusCode: &httpsRedirectStatusCode, RegexPriority: &regexPriority, StripPath: stripPath, PreserveHost: preserveHost, RequestBuffering: requestBuffering, ResponseBuffering: responseBuffering, Tags: tags, Service: service}
	_, err = kongManager.CreateRoute(c.Context, &route)
	if err != nil {
		return err
	}
	return nil
}

func deleteRouteCmdFunc(c *cli.Context) error {
	nameOrID := c.String("name")
	kongManager.DeleteRoute(c.Context, nameOrID)
	return nil
}

func getRouteCmdFunc(c *cli.Context) error {
	nameOrID := c.String("name")
	route, err := kongManager.GetRoute(c.Context, nameOrID)
	if err != nil {
		return err
	}
	routeStr, err := kong.RouteStringDetailed(route)
	if err != nil {
		return err
	}
	fmt.Print(routeStr)
	return nil
}

func listRoutesCmdFunc(c *cli.Context) error {
	serviceNameOrID := c.String("service")
	var routes []*k.Route
	if serviceNameOrID == "" {
		var err error
		routes, err = kongManager.ListAllRoutes(c.Context)
		if err != nil {
			return err
		}
	} else {
		var err error
		routes, err = kongManager.ListRoutesForService(c.Context, serviceNameOrID)
		if err != nil {
			return err
		}
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
