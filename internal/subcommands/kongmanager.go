package subcommands

import (
	kong "github.com/furkansb/kong-cli/internal/kongclient"
	"os"
)

var kongManager *kong.KongManager

func init() {
	var err error
	kongManager, err = kong.NewKongManager(getBaseUrl())
	if err != nil {
		panic(err)
	}
}

func getBaseUrl() string {
	baseUrl := os.Getenv("KONG_ADMIN_URL")
	if baseUrl == "" {
		return "http://localhost:8001"
	}
	return baseUrl
}
