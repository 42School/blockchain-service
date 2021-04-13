package api

import (
	"context"
	"github.com/42School/blockchain-service/src/tools"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"os"
)

var FtApi *clientcredentials.Config

func InitApi() error {
	FtApi = &clientcredentials.Config{
		ClientID:     os.Getenv("APP_CLIENT_ID"),
		ClientSecret: os.Getenv("APP_CLIENT_SECRET"),
		TokenURL:     "https://api.intra.42.fr/oauth/token",
		AuthStyle:    oauth2.AuthStyleInParams,
		Scopes:       []string{"public"},
	}
	_, err := FtApi.Token(context.Background())
	if err != nil {
		tools.LogsError(err)
		return err
	}
	return nil
}
