package webhook

import (
	"encoding/json"
	"fmt"
)

type EventInterface interface {
	GetType() string
}

func (e Event) GetType() string {
	return e.Type
}

type Event struct {
	Type            string          `json:"type"`
	Mode            string          `json:"mode"`
	Timestamp       int64           `json:"timestamp"`
	Source          SourceInterface `json:"source,omitempty"`
	WebhookEventId  string          `json:"webhookEventId"`
	DeliveryContext DeliveryContext `json:"deliveryContext"`
}

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

type UnknownEvent struct {
	EventInterface
	Type string
	Raw  map[string]json.RawMessage
}

func UnmarshalEvent(data []byte) (EventInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalEvent: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalEvent: Cannot read type: %w", err)
	}

	switch discriminator {
	case "message":
		var message MessageEvent
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, fmt.Errorf("UnmarshalEvent: Cannot read message: %w", err)
		}
		return message, nil

	default:
		var unknown UnknownEvent
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
