package google

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type CalendarContentInterface interface {
	GetType() string
}

func (c CalendarContent) GetType() string {
	return c.Type
}

type CalendarContent struct {
	Type     string    `json:"type"`
	Summary  string    `json:"summary"`
	Location string    `json:"location"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}

type CalenderEvent struct {
	Summary  string        `json:"summary"`
	Location string        `json:"location"`
	Start    EventDateTime `json:"start"`
	End      EventDateTime `json:"end"`
}

type EventDateTime struct {
	DateTime time.Time `json:"dateTime"`
	TimeZone string    `json:"timeZone"`
}

func ParseCalendarContent(jsonStr string) (CalendarContent, error) {
	var c CalendarContent
	err := json.Unmarshal([]byte(jsonStr), &c)
	if err != nil {
		return c, err
	}

	if c.Type == "unknown" {
		return c, errors.New("Calendar content type is unknown")
	}

	return c, nil
}

func (c *OAuthClient) RegisterCalenderEvent(content CalendarContent, accessToken string) error {
	event := CalenderEvent{
		Summary:  content.Summary,
		Location: content.Location,
		Start: EventDateTime{
			DateTime: content.Start,
			TimeZone: content.Start.Location().String(),
		},
		End: EventDateTime{
			DateTime: content.End,
			TimeZone: content.End.Location().String(),
		},
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, c.Config.CalenderEndpoint+"/calendars/primary/events", bytes.NewBuffer(eventJson))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register event: %s", resp.Status)
	}

	return nil
}
