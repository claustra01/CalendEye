package openai

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"strings"
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

	body := fmt.Sprintf(`{
		"model": "gpt-4-turbo",
		"messages": [
		  {
			"role": "user",
			"content": [
			  {
				"type": "text",
				"text": "What's in this image?"
			  },
			  {
				"type": "image_url",
				"image_url": {
				  "url": "data:image/%s;base64,%s"
				}
			  }
			]
		  }
		],
		"max_tokens": 300
	}`, format, base64Image)

	req, err := http.NewRequest(http.MethodPost, c.endpoint.String(), strings.NewReader(body))
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

	return string(respBody), nil
}
