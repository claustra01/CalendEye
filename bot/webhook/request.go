package webhook

import (
	"encoding/json"
	"fmt"
)

type CallbackRequest struct {
	Destination string           `json:"destination"`
	Events      []EventInterface `json:"events"`
}

func (cr *CallbackRequest) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("JSON parse error in map: %w", err)
	}

	if raw["destination"] != nil {
		err = json.Unmarshal(raw["destination"], &cr.Destination)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(Destination): %w", err)
		}
	}

	if raw["events"] != nil {
		var rawevents []json.RawMessage
		err = json.Unmarshal(raw["events"], &rawevents)
		if err != nil {
			return fmt.Errorf("JSON parse error in events(array): %w", err)
		}
		for _, data := range rawevents {
			e, err := UnmarshalEvent(data)
			if err != nil {
				return fmt.Errorf("JSON parse error in Event(discriminator array): %w", err)
			}
			cr.Events = append(cr.Events, e)
		}
	}

	return nil
}
