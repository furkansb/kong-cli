package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

// Create a new *kong.Service
func (k *KongManager) CreateService(ctx context.Context, service *kong.Service) *kong.Service {
	schema, err := k.client.Schemas.Get(ctx, "services")
	if err != nil {
		fmt.Printf("Error getting schema: %s", err)
		return nil
	}
	err = kong.FillEntityDefaults(service, schema)
	if err != nil {
		fmt.Printf("Error filling entity defaults: %s", err)
		return nil
	}
	service, err = k.client.Services.Create(ctx, service)
	if err != nil {
		fmt.Printf("Error creating service: %s", err)
		return nil
	}
	return service
}

func (k *KongManager) DeleteService(ctx context.Context, nameOrId string) {
	err := k.client.Services.Delete(ctx, &nameOrId)
	if err != nil {
		fmt.Printf("Error deleting service: %s", err)
	}
}

func (k *KongManager) GetService(ctx context.Context, nameOrID string) *kong.Service {
	service, err := k.client.Services.Get(ctx, &nameOrID)
	if err != nil {
		fmt.Printf("Error getting service: %s", err)
		return nil
	}
	return service
}

func (k *KongManager) UpdateService(ctx context.Context, service *kong.Service) *kong.Service {
	service, err := k.client.Services.Update(ctx, service)
	if err != nil {
		fmt.Printf("Error updating service: %s", err)
		return nil
	}
	return service
}

func (k *KongManager) ListAllServices(ctx context.Context) []*kong.Service {
	services, err := k.client.Services.ListAll(ctx)
	if err != nil {
		fmt.Printf("Error listing services: %s", err)
		return nil
	}
	return services
}

func ServiceString(service *kong.Service) string {
	serviceID := strPointerToStr(service.ID)
	serviceName := strPointerToStr(service.Name)
	serviceHost := strPointerToStr(service.Host)
	return fmt.Sprintf("Service ID: %s, name: %s, host: %s, port: %d\n", serviceID, serviceName, serviceHost, service.Port)
}
