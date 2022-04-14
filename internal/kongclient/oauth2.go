package kongclient

import (
	"context"
	"fmt"

	"github.com/kong/go-kong/kong"
)

func (k *KongManager) CreateOauth2Credential(ctx context.Context, consumerUsernameOrID string, credential *kong.Oauth2Credential) (*kong.Oauth2Credential, error) {
	// TODO: check if credential is valid
	credential, err := k.client.Oauth2Credentials.Create(ctx, &consumerUsernameOrID, credential)
	if err != nil {
		return nil, fmt.Errorf("error creating credential: %s", err)
	}
	return credential, nil
}

func (k *KongManager) DeleteOauth2Credential(ctx context.Context, consumerUsernameOrID string, clientIDorID string) error {
	err := k.client.Oauth2Credentials.Delete(ctx, &consumerUsernameOrID, &clientIDorID)
	if err != nil {
		return fmt.Errorf("error deleting credential: %s", err)
	}
	return nil
}

func (k *KongManager) UpdateOauth2Credential(ctx context.Context, consumerUsernameOrID string, credential *kong.Oauth2Credential) (*kong.Oauth2Credential, error) {
	credential, err := k.client.Oauth2Credentials.Update(ctx, &consumerUsernameOrID, credential)
	if err != nil {
		return nil, fmt.Errorf("error updating credential: %s", err)
	}
	return credential, nil
}

func (k *KongManager) GetOauth2Credential(ctx context.Context, consumerUsernameOrID string, clientIDorID string) (*kong.Oauth2Credential, error) {
	credential, err := k.client.Oauth2Credentials.Get(ctx, &consumerUsernameOrID, &clientIDorID)
	if err != nil {
		return nil, fmt.Errorf("error getting credential: %s", err)
	}
	return credential, nil
}

func (k *KongManager) ListAllOauth2Credentials(ctx context.Context) ([]*kong.Oauth2Credential, error) {
	credentials, err := k.client.Oauth2Credentials.ListAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing credentials: %s", err)
	}
	return credentials, nil
}

func (k *KongManager) ListAllOauth2CredentialsForConsumer(ctx context.Context, consumerUsernameOrID string) ([]*kong.Oauth2Credential, error) {
	credentials, _, err := k.client.Oauth2Credentials.ListForConsumer(ctx, &consumerUsernameOrID, &kong.ListOpt{})
	if err != nil {
		return nil, fmt.Errorf("error listing credentials: %s", err)
	}
	return credentials, nil
}

func Oauth2String(oauth2Credential *kong.Oauth2Credential) (string, error) {
	if oauth2Credential == nil {
		return "", nil
	}

	return fmt.Sprintf("Oauth2 ClientID: %s, ConsumerName: %s, ConsumerID: %s\n", strPointerToStr(oauth2Credential.ClientID), strPointerToStr(oauth2Credential.Consumer.Username), strPointerToStr(oauth2Credential.Consumer.ID)), nil
}
