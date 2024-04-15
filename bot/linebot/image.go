package linebot

import (
	"context"
	"fmt"
	"image"
	"net/http"
)

func (c *LineBot) FetchLineImage(ctx context.Context, id string) (image.Image, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api-data.line.me/v2/bot/message/"+id+"/content", nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.channelToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("failed to fetch image: %s", resp.Status)
	}
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	img, format, err := image.Decode(resp.Body)
	return img, format, nil
}

func (c *LineBot) FetchExternalImage(ctx context.Context, url string) (image.Image, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("failed to fetch image: %s", resp.Status)
	}
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	img, format, err := image.Decode(resp.Body)
	return img, format, nil
}
