package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

func (k *KongManager) CreateRoute(ctx context.Context, route *kong.Route) *kong.Route {
	schema, err := k.client.Schemas.Get(ctx, "routes")
	if err != nil {
		fmt.Printf("Error getting schema: %s", err)
		return nil
	}
	err = kong.FillEntityDefaults(route, schema)
	if err != nil {
		fmt.Printf("Error filling entity defaults: %s", err)
		return nil
	}
	route, err = k.client.Routes.Create(ctx, route)
	if err != nil {
		fmt.Printf("Error creating route: %s", err)
		return nil
	}
	return route
}

func (k *KongManager) DeleteRoute(ctx context.Context, nameOrId string) {
	err := k.client.Routes.Delete(ctx, &nameOrId)
	if err != nil {
		fmt.Printf("Error deleting route: %s", err)
	}
}

func (k *KongManager) UpdateRoute(ctx context.Context, route *kong.Route) *kong.Route {
	route, err := k.client.Routes.Update(ctx, route)
	if err != nil {
		fmt.Printf("Error updating route: %s", err)
		return nil
	}
	return route
}

func (k *KongManager) GetRoute(ctx context.Context, nameOrID string) *kong.Route {
	route, err := k.client.Routes.Get(ctx, &nameOrID)
	if err != nil {
		fmt.Printf("Error getting route: %s", err)
		return nil
	}
	return route
}

func (k *KongManager) ListAllRoutes(ctx context.Context) []*kong.Route {
	routes, err := k.client.Routes.ListAll(ctx)
	if err != nil {
		fmt.Printf("Error listing routes: %s", err)
		return nil
	}
	return routes
}

func (k *KongManager) ListRoutesForService(ctx context.Context, serviceNameOrID string) []*kong.Route {
	routes, _, err := k.client.Routes.ListForService(ctx, &serviceNameOrID, &kong.ListOpt{})
	if err != nil {
		fmt.Printf("Error listing routes for service: %s", err)
		return nil
	}
	return routes
}

// TODO: Add more descriptive output
func RouteString(route *kong.Route) string {
	id := strPointerToStr(route.ID)
	firstPath := strPointerToStr(route.Paths[0])
	return fmt.Sprintf("Route ID: %s, Path: %s\n", id, firstPath)
}
