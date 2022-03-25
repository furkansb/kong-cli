package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

// Create a new *kong.Service
func (k *KongManager) CreateService(ctx context.Context, service *kong.Service) (*kong.Service, error) {
	schema, err := k.client.Schemas.Get(ctx, "services")
	if err != nil {
		return nil, fmt.Errorf("error getting schema: %s", err)
	}
	err = kong.FillEntityDefaults(service, schema)
	if err != nil {
		return nil, fmt.Errorf("error filling entity defaults: %s", err)
	}
	service, err = k.client.Services.Create(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("error creating service: %s", err)
	}
	return service, nil
}

func (k *KongManager) DeleteService(ctx context.Context, nameOrId string) error {
	err := k.client.Services.Delete(ctx, &nameOrId)
	if err != nil {
		return fmt.Errorf("error deleting service: %s", err)
	}
	return nil
}

func (k *KongManager) GetService(ctx context.Context, nameOrID string) (*kong.Service, error) {
	service, err := k.client.Services.Get(ctx, &nameOrID)
	if err != nil {
		return nil, fmt.Errorf("error getting service: %s", err)
	}
	return service, nil
}

func (k *KongManager) UpdateService(ctx context.Context, service *kong.Service) (*kong.Service, error) {
	service, err := k.client.Services.Update(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("error updating service: %s", err)
	}
	return service, nil
}

func (k *KongManager) ListAllServices(ctx context.Context) ([]*kong.Service, error) {
	services, err := k.client.Services.ListAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing services: %s", err)
	}
	return services, nil
}

func ServiceString(service *kong.Service) string {
	serviceID := strPointerToStr(service.ID)
	serviceName := strPointerToStr(service.Name)
	serviceHost := strPointerToStr(service.Host)
	return fmt.Sprintf("Service ID: %s, name: %s, host: %s, port: %d\n", serviceID, serviceName, serviceHost, service.Port)
}