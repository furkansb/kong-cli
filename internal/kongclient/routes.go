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

// TODO: Add more descriptive output
func RouteString(route *kong.Route) (string, error) {
	if route == nil {
		return "", fmt.Errorf("route is nil")
	}
	id := strPointerToStr(route.ID)
	firstPath := strPointerToStr(route.Paths[0])
	return fmt.Sprintf("Route ID: %s, Path: %s\n", id, firstPath), nil
}
