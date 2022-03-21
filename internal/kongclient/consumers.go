package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

func (k *KongManager) CreateConsumer(ctx context.Context, consumer *kong.Consumer) *kong.Consumer {
	schema, err := k.client.Schemas.Get(ctx, "consumers")
	if err != nil {
		fmt.Printf("Error getting schema: %s", err)
		return nil
	}
	err = kong.FillEntityDefaults(consumer, schema)
	if err != nil {
		fmt.Printf("Error filling entity defaults: %s", err)
		return nil
	}
	consumer, err = k.client.Consumers.Create(ctx, consumer)
	if err != nil {
		fmt.Printf("Error creating consumer: %s", err)
		return nil
	}
	return consumer
}

func (k *KongManager) DeleteConsumer(ctx context.Context, usernameOrID string) {
	err := k.client.Consumers.Delete(ctx, &usernameOrID)
	if err != nil {
		fmt.Printf("Error deleting consumer: %s", err)
	}
}

func (k *KongManager) GetConsumer(ctx context.Context, usernameOrID string) *kong.Consumer {
	consumer, err := k.client.Consumers.Get(ctx, &usernameOrID)
	if err != nil {
		fmt.Printf("Error getting consumer: %s", err)
		return nil
	}
	return consumer
}

func (k *KongManager) UpdateConsumer(ctx context.Context, consumer *kong.Consumer) *kong.Consumer {
	consumer, err := k.client.Consumers.Update(ctx, consumer)
	if err != nil {
		fmt.Printf("Error updating consumer: %s", err)
		return nil
	}
	return consumer
}

func (k *KongManager) ListAllConsumers(ctx context.Context) []*kong.Consumer {
	consumers, err := k.client.Consumers.ListAll(ctx)
	if err != nil {
		fmt.Printf("Error listing consumers: %s", err)
		return nil
	}
	return consumers
}

func ConsumerString(consumer *kong.Consumer) string {
	return fmt.Sprintf("Consumer Name: %s, ID: %s\n", *consumer.Username, *consumer.ID)
}
