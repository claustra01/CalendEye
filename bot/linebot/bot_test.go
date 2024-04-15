package linebot_test

import (
	"testing"

	. "github.com/claustra01/calendeye/linebot"
)

func TestUrl(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Valid path",
			path:     "/v2/bot/message/reply",
			expected: "https://api.line.me/v2/bot/message/reply",
		},
		{
			name:     "Empty path",
			path:     "",
			expected: "https://api.line.me",
		},
	}

	bot, _ := NewBot()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := bot.Url(tt.path)

			if url != tt.expected {
				t.Errorf("got %s; want %s", url, tt.expected)
			}
		})
	}
}
