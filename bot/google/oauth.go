package google

import (
	"context"
	"net/http"
	"os"
)

type OAuthClientInterface interface {
	GetToken(code string) (string, error)
}

type OAuthClient struct {
	OAuthClientInterface
	httpClient *http.Client
	Config     *Config
	ctx        context.Context
}

type Config struct {
	ClientId     string
	ClientSecret string
	Scopes       []string
	Endpoint     string
	RedirectUri  string
}

func NewOAuthClient(ctx context.Context) *OAuthClient {
	config := &Config{
		ClientId:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     "https://accounts.google.com/o/oauth2",
		RedirectUri:  os.Getenv("GOOGLE_REDIRECT"),
	}

	return &OAuthClient{
		httpClient: http.DefaultClient,
		Config:     config,
		ctx:        ctx,
	}
}
