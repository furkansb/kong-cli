package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

func (k *KongManager) CreatePlugin(ctx context.Context, plugin *kong.Plugin) (*kong.Plugin, error) {
	schema, err := k.client.Schemas.Get(ctx, "plugins")
	if err != nil {
		return nil, fmt.Errorf("error getting schema: %s", err)
	}
	err = kong.FillEntityDefaults(plugin, schema)
	if err != nil {
		return nil, fmt.Errorf("error filling entity defaults: %s", err)
	}
	plugin, err = k.client.Plugins.Create(ctx, plugin)
	if err != nil {
		return nil, fmt.Errorf("error creating plugin: %s", err)
	}
	return plugin, nil
}

func (k *KongManager) DeletePlugin(ctx context.Context, usernameOrID string) error {
	err := k.client.Plugins.Delete(ctx, &usernameOrID)
	if err != nil {
		return fmt.Errorf("error deleting plugin: %s", err)
	}
	return nil
}

func (k *KongManager) UpdatePlugin(ctx context.Context, plugin *kong.Plugin) (*kong.Plugin, error) {
	plugin, err := k.client.Plugins.Update(ctx, plugin)
	if err != nil {
		return nil, fmt.Errorf("error updating plugin: %s", err)
	}
	return plugin, nil
}

func (k *KongManager) GetPlugin(ctx context.Context, usernameOrID string) (*kong.Plugin, error) {
	plugin, err := k.client.Plugins.Get(ctx, &usernameOrID)
	if err != nil {
		return nil, fmt.Errorf("error getting plugin: %s", err)
	}
	return plugin, nil
}

func (k *KongManager) ListAllPlugins(ctx context.Context) ([]*kong.Plugin, error) {
	plugins, err := k.client.Plugins.ListAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing plugins: %s", err)
	}
	return plugins, nil
}

func (k *KongManager) ListAllPluginsForConsumer(ctx context.Context, consumerIDorName string) ([]*kong.Plugin, error) {
	plugins, err := k.client.Plugins.ListAllForConsumer(ctx, &consumerIDorName)
	if err != nil {
		return nil, fmt.Errorf("error listing plugins: %s", err)
	}
	return plugins, nil
}

func (k *KongManager) ListAllPluginsForService(ctx context.Context, serviceIDorName string) ([]*kong.Plugin, error) {
	plugins, err := k.client.Plugins.ListAllForService(ctx, &serviceIDorName)
	if err != nil {
		return nil, fmt.Errorf("error listing plugins for service: %s", err)
	}
	return plugins, nil
}

func (k *KongManager) ListAllPluginsForRoute(ctx context.Context, routeIDorName string) ([]*kong.Plugin, error) {
	plugins, err := k.client.Plugins.ListAllForRoute(ctx, &routeIDorName)
	if err != nil {
		return nil, fmt.Errorf("error listing plugins for route: %s", err)
	}
	return plugins, nil
}
