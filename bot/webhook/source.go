package webhook

import (
	"encoding/json"
	"fmt"
)

type SourceInterface interface {
	GetType() string
}

func (e Source) GetType() string {
	return e.Type
}

type Source struct {
	Type string `json:"type"`
}

type UserSource struct {
	Source
	UserId string `json:"userId"`
}

type UnknownSource struct {
	SourceInterface
	Type string
	Raw  map[string]json.RawMessage
}

func UnmarshalSource(data []byte) (SourceInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalSource: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalSource: Cannot read type: %w", err)
	}

	switch discriminator {
	case "user":
		var user UserSource
		if err := json.Unmarshal(data, &user); err != nil {
			return nil, fmt.Errorf("UnmarshalSource: Cannot read user: %w", err)
		}
		return user, nil

	default:
		var unknown UnknownSource
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
