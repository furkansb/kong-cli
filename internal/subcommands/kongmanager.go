package subcommands

import (
	kong "github.com/furkansb/kong-cli/internal/kongclient"
	"os"
)

var kongManager *kong.KongManager = kong.NewKongManager(getBaseUrl())

func getBaseUrl() string {
	baseUrl := os.Getenv("KONG_ADMIN_URL")
	if baseUrl == "" {
		return "http://localhost:8001"
	}
	return baseUrl
}
