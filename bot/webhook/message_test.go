package webhook_test

import (
	"reflect"
	"testing"

	. "github.com/claustra01/calendeye/webhook"
)

func TestMessageEvent_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		expect  MessageEvent
		wantErr bool
	}{
		{
			name: "Valid text message",
			data: []byte(`{"type":"message","mode":"test_mode","timestamp":123456789,"source":{"type":"user","userId":"123"},"webhookEventId":"event_id","deliveryContext":{"isRedelivery":true},"message":{"type":"text","id":"message_id","text":"Hello, world!"},"replyToken":"reply_token"}`),
			expect: MessageEvent{
				Event: Event{
					Type:            "message",
					Mode:            "test_mode",
					Timestamp:       123456789,
					Source:          UserSource{Source: Source{Type: "user"}, UserId: "123"},
					WebhookEventId:  "event_id",
					DeliveryContext: DeliveryContext{IsRedelivery: true},
				},
				ReplyToken: "reply_token",
				Message: TextMessageContent{
					MessageContent: MessageContent{Type: "text", Id: "message_id"},
					Text:           "Hello, world!",
				},
			},
			wantErr: false,
		},
		{
			name: "Valid line image message",
			data: []byte(`{"type":"message","mode":"test_mode","timestamp":123456789,"source":{"type":"user","userId":"123"},"webhookEventId":"event_id","deliveryContext":{"isRedelivery":true},"message":{"type":"image","id":"message_id","contentProvider":{"type":"line"}},"replyToken":"reply_token"}`),
			expect: MessageEvent{
				Event: Event{
					Type:            "message",
					Mode:            "test_mode",
					Timestamp:       123456789,
					Source:          UserSource{Source: Source{Type: "user"}, UserId: "123"},
					WebhookEventId:  "event_id",
					DeliveryContext: DeliveryContext{IsRedelivery: true},
				},
				ReplyToken: "reply_token",
				Message: ImageMessageContent{
					MessageContent:  MessageContent{Type: "image", Id: "message_id"},
					ContentProvider: ImageContentProvider{Type: "line", OriginalContentUrl: "", PreviewImageUrl: ""},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid external image message",
			data: []byte(`{"type":"message","mode":"test_mode","timestamp":123456789,"source":{"type":"user","userId":"123"},"webhookEventId":"event_id","deliveryContext":{"isRedelivery":true},"message":{"type":"image","id":"message_id","contentProvider":{"type":"external","originalContentUrl":"original_url","previewImageUrl":"preview_url"}},"replyToken":"reply_token"}`),
			expect: MessageEvent{
				Event: Event{
					Type:            "message",
					Mode:            "test_mode",
					Timestamp:       123456789,
					Source:          UserSource{Source: Source{Type: "user"}, UserId: "123"},
					WebhookEventId:  "event_id",
					DeliveryContext: DeliveryContext{IsRedelivery: true},
				},
				ReplyToken: "reply_token",
				Message: ImageMessageContent{
					MessageContent:  MessageContent{Type: "image", Id: "message_id"},
					ContentProvider: ImageContentProvider{Type: "external", OriginalContentUrl: "original_url", PreviewImageUrl: "preview_url"},
				},
			},
			wantErr: false,
		},
		{
			name:    "Invalid JSON",
			data:    []byte(`invalid json`),
			expect:  MessageEvent{},
			wantErr: true,
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

			if !reflect.DeepEqual(cr, tt.expect) {
				t.Errorf("MessageEvent.UnmarshalJSON() got = %v, want %v", cr, tt.expect)
			}
		})
	}
}
