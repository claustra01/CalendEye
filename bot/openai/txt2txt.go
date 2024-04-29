package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Gpt4Vision) Txt2Txt(eventTxt string) (string, error) {
	now := time.Now()
	prompt := fmt.Sprintf(`
	この後に送るテキストは、カレンダーに登録するようなイベントや予定の内容が含まれる、ユーザーが入力した文字列です。
	その画像の内容をGoogleカレンダーへ登録するため、以下のような構造体を用意しました。この構造体に展開できるようなJSON文字列を返してください。
	ただし、時間については特に明示されない限り日本時間（UTC+9）とします。
	また、Typeについては、「event」「work」「reminder」「unknown」のいずれかとし、eventは遊びなどを示すもの、workは仕事や学業に関するもの、reminderは何らかの期限を表すものとします。
	画像がイベントや予定の内容ではない無関係な画像だと判断した場合のみ、unknownを使用してください。
	なお、JSON以外の文字列、すなわち画像自体の説明テキストなどの情報は不要です。必ずJSON文字列のみを返してください。
	改行やコードブロック等に関連する各種記号も不要です。**絶対に**一行のJSON文字列のみを返してください。
	ちなみに、今日は%s年%s月%s日です。
	'''go
	type CalendarContent struct {
		Type     string    'json:"type"'
		Summary  string    'json:"summary"'
		Location string    'json:"location"'
		Start    time.Time 'json:"start"'
		End      time.Time 'json:"end"'
	}
	'''
`, now.Format("2006"), now.Format("01"), now.Format("02"))

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
						Type: "text",
						Text: eventTxt,
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
