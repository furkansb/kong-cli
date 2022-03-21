package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

func (k *KongManager) CreateOauth2Credential(ctx context.Context, consumerUsernameOrID string, credential *kong.Oauth2Credential) *kong.Oauth2Credential {
	schema, err := k.client.Schemas.Get(ctx, "oauth2_credentials")
	if err != nil {
		fmt.Printf("Error getting schema: %s", err)
		return nil
	}
	err = kong.FillEntityDefaults(credential, schema)
	if err != nil {
		fmt.Printf("Error filling entity defaults: %s", err)
		return nil
	}
	credential, err = k.client.Oauth2Credentials.Create(ctx, &consumerUsernameOrID, credential)
	if err != nil {
		fmt.Printf("Error creating credential: %s", err)
		return nil
	}
	return credential
}

func (k *KongManager) DeleteOauth2Credential(ctx context.Context, consumerUsernameOrID string, clientIDorID string) {
	err := k.client.Oauth2Credentials.Delete(ctx, &consumerUsernameOrID, &clientIDorID)
	if err != nil {
		fmt.Printf("Error deleting credential: %s", err)
	}
}

func (k *KongManager) UpdateOauth2Credential(ctx context.Context, consumerUsernameOrID string, credential *kong.Oauth2Credential) *kong.Oauth2Credential {
	credential, err := k.client.Oauth2Credentials.Update(ctx, &consumerUsernameOrID, credential)
	if err != nil {
		fmt.Printf("Error updating credential: %s", err)
		return nil
	}
	return credential
}

func (k *KongManager) GetOauth2Credential(ctx context.Context, consumerUsernameOrID string, clientIDorID string) *kong.Oauth2Credential {
	credential, err := k.client.Oauth2Credentials.Get(ctx, &consumerUsernameOrID, &clientIDorID)
	if err != nil {
		fmt.Printf("Error getting credential: %s", err)
		return nil
	}
	return credential
}

func (k *KongManager) ListAllOauth2Credentials(ctx context.Context, consumerUsernameOrID string) []*kong.Oauth2Credential {
	credentials, err := k.client.Oauth2Credentials.ListAll(ctx)
	if err != nil {
		fmt.Printf("Error listing credentials: %s", err)
		return nil
	}
	return credentials
}

func (k *KongManager) ListAllOauth2CredentialsForConsumer(ctx context.Context, consumerUsernameOrID string) []*kong.Oauth2Credential {
	credentials, _, err := k.client.Oauth2Credentials.ListForConsumer(ctx, &consumerUsernameOrID, &kong.ListOpt{})
	if err != nil {
		fmt.Printf("Error listing credentials: %s", err)
		return nil
	}
	return credentials
}

func Oauth2String(oauth2Credential kong.Oauth2Credential) string {
	return fmt.Sprintf("Oauth2 ClientID: %s, Consumer: %s\n", *oauth2Credential.ClientID, *oauth2Credential.Consumer.Username)
}
