package webhook_test

import (
	"reflect"
	"testing"

	. "github.com/claustra01/calendeye/webhook"
)

func TestFollowEvent_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected FollowEvent
		wantErr  bool
	}{
		{
			name: "Valid follow event",
			data: []byte(`{"type":"follow","mode":"test_mode","timestamp":123456789,"source":{"type":"user","userId":"123"},"webhookEventId":"event_id","deliveryContext":{"isRedelivery":true},"follow":{"isUnblocked":false},"replyToken":"reply_token"}`),
			expected: FollowEvent{
				Event: Event{
					Type:            "follow",
					Mode:            "test_mode",
					Timestamp:       123456789,
					Source:          UserSource{Source: Source{Type: "user"}, UserId: "123"},
					WebhookEventId:  "event_id",
					DeliveryContext: DeliveryContext{IsRedelivery: true},
				},
				ReplyToken: "reply_token",
				Follow:     Follow{IsUnblocked: false},
			},
			wantErr: false,
		},
		{
			name:     "Invalid JSON",
			data:     []byte(`invalid json`),
			expected: FollowEvent{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cr FollowEvent
			err := cr.UnmarshalJSON(tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("FollowEvent.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(cr, tt.expected) {
				t.Errorf("FollowEvent.UnmarshalJSON() got = %v, want %v", cr, tt.expected)
			}
		})
	}
}
