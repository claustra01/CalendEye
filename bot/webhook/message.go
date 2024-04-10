package webhook

import (
	"encoding/json"
	"fmt"
)

type MessageEvent struct {
	Event
	Message    MessageContentInterface `json:"message"`
	ReplyToken string                  `json:"replyToken"`
}

type MessageContentInterface interface {
	GetType() string
}

func (e MessageContent) GetType() string {
	return e.Type
}

type MessageContent struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type UnknownMessageContent struct {
	MessageContentInterface
	Type string
	Raw  map[string]json.RawMessage
}

type TextMessageContent struct {
	MessageContent
	Text string `json:"text"`
}

type ImageMessageContent struct {
	MessageContent
	ContentProvider ImageContentProvider `json:"contentProvider"`
}

type ImageContentProvider struct {
	Type               string `json:"type"`
	OriginalContentUrl string `json:"originalContentUrl"`
	PreviewImageUrl    string `json:"previewImageUrl"`
}

func (cr *MessageEvent) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("JSON parse error in map: %w", err)
	}

	if raw["type"] != nil {
		err = json.Unmarshal(raw["type"], &cr.Type)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(Type): %w", err)
		}
	}

	if raw["source"] != nil {
		if rawsource, ok := raw["source"]; ok && rawsource != nil {
			Source, err := UnmarshalSource(rawsource)
			if err != nil {
				return fmt.Errorf("JSON parse error in Source(discriminator): %w", err)
			}
			cr.Source = Source
		}
	}

	if raw["timestamp"] != nil {
		err = json.Unmarshal(raw["timestamp"], &cr.Timestamp)
		if err != nil {
			return fmt.Errorf("JSON parse error in int64(Timestamp): %w", err)
		}
	}

	if raw["mode"] != nil {
		err = json.Unmarshal(raw["mode"], &cr.Mode)
		if err != nil {
			return fmt.Errorf("JSON parse error in EventMode(Mode): %w", err)
		}
	}

	if raw["webhookEventId"] != nil {
		err = json.Unmarshal(raw["webhookEventId"], &cr.WebhookEventId)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(WebhookEventId): %w", err)
		}
	}

	if raw["deliveryContext"] != nil {
		err = json.Unmarshal(raw["deliveryContext"], &cr.DeliveryContext)
		if err != nil {
			return fmt.Errorf("JSON parse error in DeliveryContext(DeliveryContext): %w", err)
		}
	}

	if raw["replyToken"] != nil {
		err = json.Unmarshal(raw["replyToken"], &cr.ReplyToken)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(ReplyToken): %w", err)
		}
	}

	if raw["message"] != nil {
		if rawmessage, ok := raw["message"]; ok && rawmessage != nil {
			Message, err := UnmarshalMessageContent(rawmessage)
			if err != nil {
				return fmt.Errorf("JSON parse error in MessageContent(discriminator): %w", err)
			}
			cr.Message = Message
		}
	}

	return nil
}

func UnmarshalMessageContent(data []byte) (MessageContentInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalMessageContent: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read type: %w", err)
	}

	switch discriminator {
	case "text":
		var text TextMessageContent
		if err := json.Unmarshal(data, &text); err != nil {
			return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read text: %w", err)
		}
		return text, nil

	case "image":
		var image ImageMessageContent
		if err := json.Unmarshal(data, &image); err != nil {
			return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read image content: %w", err)
		}
		return image, nil

	default:
		var unknown UnknownMessageContent
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
