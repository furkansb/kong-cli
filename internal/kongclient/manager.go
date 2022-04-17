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
func NewKongManager(url string) (*KongManager, error) {
	kongClient, err := kong.NewClient(&url, http.DefaultClient)
	if err != nil {
		return nil, fmt.Errorf("error creating kong client: %s", err)
	}
	return &KongManager{client: kongClient}, nil
}
