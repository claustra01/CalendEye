package linebot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ReplyMessageRequest struct {
	ReplyToken           string             `json:"replyToken"`
	Messages             []MessageInterface `json:"messages"`
	NotificationDisabled bool               `json:"notificationDisabled"`
}

type ReplyMessageResponse struct {
	SentMessages []SentMessage `json:"sentMessages"`
}

func (client *LineBot) ReplyMessage(replyMessageRequest *ReplyMessageRequest) (*ReplyMessageResponse, error) {
	_, body, err := client.ReplyMessageWithHttpInfo(replyMessageRequest)
	return body, err
}

func (client *LineBot) ReplyMessageWithHttpInfo(replyMessageRequest *ReplyMessageRequest) (*http.Response, *ReplyMessageResponse, error) {
	path := "/v2/bot/message/reply"

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(replyMessageRequest); err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest(http.MethodPost, client.Url(path), &buf)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := client.Do(req)

	if err != nil {
		return res, nil, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, nil, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := ReplyMessageResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return res, &result, nil
}
