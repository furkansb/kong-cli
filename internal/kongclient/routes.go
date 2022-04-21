package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

func (k *KongManager) CreateRoute(ctx context.Context, route *kong.Route) (*kong.Route, error) {
	schema, err := k.client.Schemas.Get(ctx, "routes")
	if err != nil {
		return nil, fmt.Errorf("error getting schema: %s", err)
	}
	err = kong.FillEntityDefaults(route, schema)
	if err != nil {
		return nil, fmt.Errorf("error filling entity defaults: %s", err)
	}
	route, err = k.client.Routes.Create(ctx, route)
	if err != nil {
		return nil, fmt.Errorf("error creating route: %s", err)
	}
	return route, nil
}

func (k *KongManager) DeleteRoute(ctx context.Context, nameOrId string) error {
	err := k.client.Routes.Delete(ctx, &nameOrId)
	if err != nil {
		return fmt.Errorf("error deleting route: %s", err)
	}
	return nil
}

func (k *KongManager) UpdateRoute(ctx context.Context, route *kong.Route) (*kong.Route, error) {
	route, err := k.client.Routes.Update(ctx, route)
	if err != nil {
		return nil, fmt.Errorf("error updating route: %s", err)
	}
	return route, nil
}

func (k *KongManager) GetRoute(ctx context.Context, nameOrID string) (*kong.Route, error) {
	route, err := k.client.Routes.Get(ctx, &nameOrID)
	if err != nil {
		return nil, fmt.Errorf("error getting route: %s", err)
	}
	return route, nil
}

func (k *KongManager) ListAllRoutes(ctx context.Context) ([]*kong.Route, error) {
	routes, err := k.client.Routes.ListAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing routes: %s", err)
	}
	return routes, nil
}

func (k *KongManager) ListRoutesForService(ctx context.Context, serviceNameOrID string) ([]*kong.Route, error) {
	routes, _, err := k.client.Routes.ListForService(ctx, &serviceNameOrID, &kong.ListOpt{})
	if err != nil {
		return nil, fmt.Errorf("error listing routes for service: %s", err)
	}
	return routes, nil
}

func RouteString(route *kong.Route) (string, error) {
	if route == nil {
		return "", fmt.Errorf("route is nil")
	}
	id := strPointerToStr(route.ID)
	hosts := StrSliceFromPSlice(route.Hosts)
	name := *route.Name
	paths := StrSliceFromPSlice(route.Paths)
	serviceID := *route.Service.ID
	tags := StrSliceFromPSlice(route.Tags)
	return fmt.Sprintf("Route ID: %s, Name, %s, Hosts: %s, Paths: %s, ServiceID: %s, tags: %s\n", id, name, hosts, paths, serviceID, tags), nil
}

func RouteStringDetailed(route *kong.Route) (string, error) {
	if route == nil {
		return "", fmt.Errorf("route is nil")
	}
	id := strPointerToStr(route.ID)
	hosts := StrSliceFromPSlice(route.Hosts)
	name := *route.Name
	methods := StrSliceFromPSlice(route.Methods)
	paths := StrSliceFromPSlice(route.Paths)
	preserveHost := *route.PreserveHost
	protocols := StrSliceFromPSlice(route.Protocols)
	serviceID := *route.Service.ID
	stripPath := *route.StripPath
	tags := StrSliceFromPSlice(route.Tags)
	return fmt.Sprintf("Route ID: %s, Name, %s, Hosts: %s, Methods: %s, Paths: %s, PreserveHost: %t, Protocols: %s, ServiceID: %s, StripPath: %t, tags: %s\n", id, name, hosts, methods, paths, preserveHost, protocols, serviceID, stripPath, tags), nil
}
