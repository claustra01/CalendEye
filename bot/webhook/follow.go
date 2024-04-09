package webhook

import (
	"encoding/json"
	"fmt"
)

type FollowEvent struct {
	Event
	ReplyToken string `json:"replyToken"`
	Follow     Follow `json:"follow"`
}

type Follow struct {
	IsUnblocked bool `json:"isUnblocked"`
}

func (cr *FollowEvent) UnmarshalJSON(data []byte) error {
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

	if raw["follow"] != nil {
		err = json.Unmarshal(raw["follow"], &cr.DeliveryContext)
		if err != nil {
			return fmt.Errorf("JSON parse error in Follow(Follow): %w", err)
		}
	}

	return nil
}
