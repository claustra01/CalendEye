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
	model      string
	ctx        context.Context
}

type OpenAIRequest struct {
	Model     string           `json:"model"`
	Messages  []RequestMessage `json:"messages"`
	MaxTokens int              `json:"max_tokens,omitempty"`
}

type RequestMessage struct {
	Role    string                  `json:"role"`
	Content []RequestMessageContent `json:"content"`
}

type RequestMessageContent struct {
	Type  string              `json:"type"`
	Text  string              `json:"text,omitempty"`
	Image RequestMessageImage `json:"image_url,omitempty"`
}

type RequestMessageImage struct {
	Url string `json:"url"`
}

type OpenAIResponse struct {
	Id      string           `json:"id"`
	Object  string           `json:"object"`
	Created int              `json:"created"`
	Model   string           `json:"model"`
	Choices []ResponseChoice `json:"choices"`
	Usage   ResponseUsage    `json:"usage"`
}

type ResponseChoice struct {
	Index   int             `json:"index"`
	Message ResponseMessage `json:"message"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func NewGpt4Vision(ctx context.Context) (*Gpt4Vision, error) {
	c := &Gpt4Vision{
		httpClient: http.DefaultClient,
		apiKey:     os.Getenv("OPENAI_API_KEY"),
		model:      "gpt-4-turbo",
		ctx:        ctx,
	}

	u, err := url.ParseRequestURI("https://api.openai.com/v1/chat/completions")
	if err != nil {
		return nil, err
	}
	c.endpoint = u

	return c, nil
}
