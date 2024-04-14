package google

import (
	"context"
	"os"
)

var (
	ClientId     = os.Getenv("GOOGLE_CLIENT_ID")
	ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	Scopes       = []string{"https://www.googleapis.com/auth/calendar"}
	Endpoint     = "https://accounts.google.com/o/oauth2"
	RedirectUri  = os.Getenv("GOOGLE_REDIRECT")
)

type OAuthClientInterface interface {
	GetToken(code string) (string, error)
}

type OAuthClient struct {
	OAuthClientInterface
	Config *Config
	ctx    context.Context
}

type Config struct {
	ClientId     string
	ClientSecret string
	Scopes       []string
	Endpoint     string
	RedirectUri  string
}

func NewOAuthClient(ctx context.Context) *OAuthClient {
	return &OAuthClient{
		Config: &Config{
			ClientId:     ClientId,
			ClientSecret: ClientSecret,
			Scopes:       Scopes,
			Endpoint:     Endpoint,
			RedirectUri:  RedirectUri,
		},
		ctx: ctx,
	}
}
