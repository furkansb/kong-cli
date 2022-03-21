package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

func (k *KongManager) CreatePlugin(ctx context.Context, plugin *kong.Plugin) *kong.Plugin {
	schema, err := k.client.Schemas.Get(ctx, "plugins")
	if err != nil {
		fmt.Printf("Error getting schema: %s", err)
		return nil
	}
	err = kong.FillEntityDefaults(plugin, schema)
	if err != nil {
		fmt.Printf("Error filling entity defaults: %s", err)
		return nil
	}
	plugin, err = k.client.Plugins.Create(ctx, plugin)
	if err != nil {
		fmt.Printf("Error creating plugin: %s", err)
		return nil
	}
	return plugin
}

func (k *KongManager) DeletePlugin(ctx context.Context, usernameOrID string) {
	err := k.client.Plugins.Delete(ctx, &usernameOrID)
	if err != nil {
		fmt.Printf("Error deleting plugin: %s", err)
	}
}

func (k *KongManager) UpdatePlugin(ctx context.Context, plugin *kong.Plugin) *kong.Plugin {
	plugin, err := k.client.Plugins.Update(ctx, plugin)
	if err != nil {
		fmt.Printf("Error updating plugin: %s", err)
		return nil
	}
	return plugin
}

func (k *KongManager) GetPlugin(ctx context.Context, usernameOrID string) *kong.Plugin {
	plugin, err := k.client.Plugins.Get(ctx, &usernameOrID)
	if err != nil {
		fmt.Printf("Error getting plugin: %s", err)
		return nil
	}
	return plugin
}

func (k *KongManager) ListAllPlugins(ctx context.Context) []*kong.Plugin {
	plugins, err := k.client.Plugins.ListAll(ctx)
	if err != nil {
		fmt.Printf("Error listing plugins: %s", err)
		return nil
	}
	return plugins
}

func (k *KongManager) ListAllPluginsForConsumer(ctx context.Context, consumerIDorName string) []*kong.Plugin {
	plugins, err := k.client.Plugins.ListAllForConsumer(ctx, &consumerIDorName)
	if err != nil {
		fmt.Printf("Error listing plugins: %s", err)
		return nil
	}
	return plugins
}

func (k *KongManager) ListAllPluginsForService(ctx context.Context, serviceIDorName string) []*kong.Plugin {
	plugins, err := k.client.Plugins.ListAllForService(ctx, &serviceIDorName)
	if err != nil {
		fmt.Printf("Error listing plugins for service: %s", err)
		return nil
	}
	return plugins
}

func (k *KongManager) ListAllPluginsForRoute(ctx context.Context, routeIDorName string) []*kong.Plugin {
	plugins, err := k.client.Plugins.ListAllForRoute(ctx, &routeIDorName)
	if err != nil {
		fmt.Printf("Error listing plugins for route: %s", err)
		return nil
	}
	return plugins
}

func PluginString(plugin *kong.Plugin) string {
	return fmt.Sprintf("Plugin ID: %s, Name: %s\n", *plugin.ID, *plugin.Name)
}
