package openai

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
)

var (
	ErrNoApiKey = errors.New("apiKey is required")
)

type Gpt4Vision struct {
	httpClient *http.Client
	endpoint   *url.URL
	apiKey     string
	ctx        context.Context
}

func NewGpt4Vision(ctx context.Context) (*Gpt4Vision, error) {
	c := &Gpt4Vision{
		httpClient: http.DefaultClient,
		apiKey:     os.Getenv("OPENAI_API_KEY"),
		ctx:        ctx,
	}

	u, err := url.ParseRequestURI("https://api.openai.com/v1/chat/completions")
	if err != nil {
		return nil, err
	}
	c.endpoint = u

	return c, nil
}
