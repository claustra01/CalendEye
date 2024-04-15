package openai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
)

func EncodeImage(img image.Image, format string) (string, error) {
	switch format {
	case "jpeg":
		imgData := new(bytes.Buffer)
		err := jpeg.Encode(imgData, img, nil)
		if err != nil {
			return "", err
		}
		imgEnc := base64.StdEncoding.EncodeToString(imgData.Bytes())
		return string(imgEnc), nil

	case "png":
		imgData := new(bytes.Buffer)
		err := png.Encode(imgData, img)
		if err != nil {
			return "", err
		}
		imgEnc := base64.StdEncoding.EncodeToString(imgData.Bytes())
		return string(imgEnc), nil

	default:
		return "", fmt.Errorf("unsupported image format: %s", format)
	}
}

func (c *Gpt4Vision) Img2Txt(img image.Image, format string) (string, error) {
	base64Image, err := EncodeImage(img, format)
	if err != nil {
		return "", err
	}

	prompt := "日本語でこの画像を完結に説明してください。"

	reqBody := OpenAIRequest{
		Model: c.model,
		Messages: []RequestMessage{
			{
				Role: "user",
				Content: []RequestMessageContent{
					{
						Type: "text",
						Text: prompt,
					},
					{
						Type: "image_url",
						Image: RequestMessageImage{
							Url: fmt.Sprintf("data:image/%s;base64,%s", format, base64Image),
						},
					},
				},
			},
		},
		MaxTokens: 300,
	}

	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, c.endpoint.String(), bytes.NewBuffer(reqJson))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respData OpenAIResponse
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		return "", err
	}
	if len(respData.Choices) != 1 {
		return "", fmt.Errorf("Invalid choices in response")
	}

	return respData.Choices[0].Message.Content, nil
}
