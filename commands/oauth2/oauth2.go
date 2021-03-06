package oauth2

import (
	"fmt"
	kong "github.com/furkansb/kong-cli/internal/kongclient"
	k "github.com/kong/go-kong/kong"
	"github.com/urfave/cli/v2"
)

var kongManager *kong.KongManager

func init() {
	var err error
	kongManager, err = kong.NewKongManager(kong.GetBaseUrl())
	if err != nil {
		panic(err)
	}
}

func Command() *cli.Command {
	return &cli.Command{
		Name:  "oauth2",
		Usage: "oauth2 related commands",
		Subcommands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a new oauth2 credential",
				Flags: addOauth2Flags,
				Action: func(c *cli.Context) error {
					return addOauth2CredCmdFunc(c)
				},
			},
			{
				Name:  "delete",
				Usage: "delete oauth2 credential",
				Flags: deleteOauth2Flags,
				Action: func(c *cli.Context) error {
					return deleteOauth2CredCmdFunc(c)
				},
			},
			{
				Name:  "get",
				Usage: "get oauth2 credential",
				Flags: getOauth2Flags,
				Action: func(c *cli.Context) error {
					return getOauth2CredCmdFunc(c)
				},
			},
			{
				Name:  "list",
				Usage: "list oauth2 credentials",
				Flags: listOauth2Flags,
				Action: func(c *cli.Context) error {
					return listOauth2CredsCmdFunc(c)
				},
			},
		},
	}
}

func addOauth2CredCmdFunc(c *cli.Context) error {
	name := c.String("name")
	clientIDStr := c.String("client-id")
	clientID := kong.StrToPointer(clientIDStr)
	clientSecretStr := c.String("client-secret")
	clientSecret := kong.StrToPointer(clientSecretStr)
	consumerUsernameOrID := c.String("consumer-username")
	consumer, err := kongManager.GetConsumer(c.Context, consumerUsernameOrID)
	if err != nil {
		return err
	}
	oauth2 := k.Oauth2Credential{Name: &name, ClientID: clientID, ClientSecret: clientSecret, Consumer: consumer}
	_, err = kongManager.CreateOauth2Credential(c.Context, consumerUsernameOrID, &oauth2)
	if err != nil {
		return err
	}
	return nil
}

func deleteOauth2CredCmdFunc(c *cli.Context) error {
	consumerUsernameOrID := c.String("consumer-name")
	clientIDorID := c.String("client-id")
	err := kongManager.DeleteOauth2Credential(c.Context, consumerUsernameOrID, clientIDorID)
	if err != nil {
		return err
	}
	return nil
}

func getOauth2CredCmdFunc(c *cli.Context) error {
	consumerUsernameOrID := c.String("consumer-username")
	clientIDorID := c.String("client-id")
	cred, err := kongManager.GetOauth2Credential(c.Context, consumerUsernameOrID, clientIDorID)
	if err != nil {
		return err
	}
	oauth2Str, err := kong.Oauth2StringWithClientSecret(cred)
	if err != nil {
		return err
	}
	fmt.Print(oauth2Str)
	return nil
}

func listOauth2CredsCmdFunc(c *cli.Context) error {
	consumerUsername := c.String("consumer-username")
	var creds []*k.Oauth2Credential
	if consumerUsername == "" {
		var err error
		creds, err = kongManager.ListAllOauth2Credentials(c.Context)
		if err != nil {
			return err
		}
	} else {
		var err error
		creds, err = kongManager.ListAllOauth2CredentialsForConsumer(c.Context, consumerUsername)
		if err != nil {
			return err
		}
	}
	for _, cred := range creds {
		oauth2Str, err := kong.Oauth2String(cred)
		if err != nil {
			return err
		}
		fmt.Print(oauth2Str)
	}
	return nil
}
