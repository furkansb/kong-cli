package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

func (k *KongManager) CreateConsumer(ctx context.Context, consumer *kong.Consumer) (*kong.Consumer, error) {
	consumer, err := k.client.Consumers.Create(ctx, consumer)
	if err != nil {
		return nil, fmt.Errorf("error creating consumer: %s", err)
	}
	return consumer, nil
}

func (k *KongManager) DeleteConsumer(ctx context.Context, usernameOrID string) error {
	err := k.client.Consumers.Delete(ctx, &usernameOrID)
	if err != nil {
		return fmt.Errorf("error deleting consumer: %s", err)
	}
	return nil
}

func (k *KongManager) GetConsumer(ctx context.Context, usernameOrID string) (*kong.Consumer, error) {
	consumer, err := k.client.Consumers.Get(ctx, &usernameOrID)
	if err != nil {
		return nil, fmt.Errorf("error getting consumer: %s", err)
	}
	return consumer, nil
}

func (k *KongManager) UpdateConsumer(ctx context.Context, consumer *kong.Consumer) (*kong.Consumer, error) {
	consumer, err := k.client.Consumers.Update(ctx, consumer)
	if err != nil {
		return nil, fmt.Errorf("error updating consumer: %s", err)
	}
	return consumer, nil
}

func (k *KongManager) ListAllConsumers(ctx context.Context) ([]*kong.Consumer, error) {
	consumers, err := k.client.Consumers.ListAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing consumers: %s", err)
	}
	return consumers, nil
}

func ConsumerString(consumer *kong.Consumer) (string, error) {
	if consumer == nil {
		return "", fmt.Errorf("consumer is nil")
	}
	ID := strPointerToStr(consumer.ID)
	username := strPointerToStr(consumer.Username)
	customID := strPointerToStr(consumer.CustomID)
	return fmt.Sprintf("ID: %s, Consumer Name: %s, ID: %s\n", ID, username, customID), nil
}
