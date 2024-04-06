package linebot_test

import (
	"testing"

	. "github.com/claustra01/calendeye/linebot"
)

func TestNewBot(t *testing.T) {
	tests := []struct {
		name         string
		channelToken string
		wantErr      error
	}{
		{
			name:         "Valid channel token",
			channelToken: "token",
			wantErr:      nil,
		},
		{
			name:         "Empty channel token",
			channelToken: "",
			wantErr:      ErrNoChannelToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBot(tt.channelToken)

			if err != tt.wantErr {
				t.Errorf("got error %v; want %v", err, tt.wantErr)
			}
		})
	}
}

func TestUrl(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Valid path",
			path: "/v2/bot/message/reply",
			want: "https://api.line.me/v2/bot/message/reply",
		},
		{
			name: "Empty path",
			path: "",
			want: "https://api.line.me",
		},
	}

	bot, _ := NewBot("token")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := bot.Url(tt.path)

			if url != tt.want {
				t.Errorf("got %s; want %s", url, tt.want)
			}
		})
	}
}
