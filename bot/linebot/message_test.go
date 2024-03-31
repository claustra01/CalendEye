package linebot_test

import (
	"testing"

	. "github.com/claustra01/calendeye/linebot"
)

func TestNewTextMessage(t *testing.T) {
	test := []struct {
		name string
		text string
	}{
		{
			name: "Valid text",
			text: "Hello, world!",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewTextMessage(tt.text)
			if msg.Type != "text" {
				t.Errorf("got %s; want text", msg.Type)
			}
			if msg.Text != tt.text {
				t.Errorf("got %s; want %s", msg.Text, tt.text)
			}
		})
	}
}
