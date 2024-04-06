package webhook_test

import (
	"reflect"
	"testing"

	. "github.com/claustra01/calendeye/webhook"
)

func TestMessageEvent_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected MessageEvent
		wantErr  bool
	}{
		{
			name: "Valid text message",
			data: []byte(`{"type":"message","mode":"test_mode","timestamp":123456789,"source":{"type":"user","userId":"123"},"webhookEventId":"event_id","deliveryContext":{"isRedelivery":true},"message":{"type":"text","id":"message_id"},"replyToken":"reply_token"}`),
			expected: MessageEvent{
				Event: Event{
					Type:            "message",
					Mode:            "test_mode",
					Timestamp:       123456789,
					Source:          UserSource{Source: Source{Type: "user"}, UserId: "123"},
					WebhookEventId:  "event_id",
					DeliveryContext: DeliveryContext{IsRedelivery: true},
				},
				ReplyToken: "reply_token",
				Message:    TextMessageContent{MessageContent: MessageContent{Type: "text", Id: "message_id"}},
			},
			wantErr: false,
		},
		{
			name:     "Invalid JSON",
			data:     []byte(`invalid json`),
			expected: MessageEvent{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cr MessageEvent
			err := cr.UnmarshalJSON(tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("MessageEvent.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(cr, tt.expected) {
				t.Errorf("MessageEvent.UnmarshalJSON() got = %v, want %v", cr, tt.expected)
			}
		})
	}
}
