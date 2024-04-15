package linebot

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"path"
)

var (
	ErrNoChannelToken = errors.New("channelToken is required")
)

type LineBot struct {
	httpClient    *http.Client
	endpoint      *url.URL
	channelToken  string
	channelSecret string
	ctx           context.Context
}

func NewBot() (*LineBot, error) {
	c := &LineBot{
		httpClient:    http.DefaultClient,
		channelToken:  os.Getenv("LINE_CHANNEL_TOKEN"),
		channelSecret: os.Getenv("LINE_CHANNEL_SECRET"),
	}

	u, err := url.ParseRequestURI("https://api.line.me")
	if err != nil {
		return nil, err
	}
	c.endpoint = u

	return c, nil
}

func (call *LineBot) WithContext(ctx context.Context) *LineBot {
	call.ctx = ctx
	return call
}

func (client *LineBot) Url(endpointPath string) string {
	newPath := path.Join(client.endpoint.Path, endpointPath)
	u := *client.endpoint
	u.Path = newPath
	return u.String()
}

func (client *LineBot) Do(req *http.Request) (*http.Response, error) {
	if client.channelToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.channelToken)
	}
	if client.ctx != nil {
		req = req.WithContext(client.ctx)
	}
	return client.httpClient.Do(req)
}

func (client *LineBot) GetChannelSecret() string {
	return client.channelSecret
}
