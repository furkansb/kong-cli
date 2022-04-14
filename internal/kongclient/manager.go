package kongclient

import (
	"fmt"
	"github.com/kong/go-kong/kong"
	"net/http"
)

type KongManager struct {
	client *kong.Client
}

// Create a new KongManager
func NewKongManager(url string) *KongManager {
	kongClient, err := kong.NewClient(&url, http.DefaultClient)
	if err != nil {
		fmt.Errorf("Error creating kong client: %s", err)
		return nil
	}
	return &KongManager{client: kongClient}
}
