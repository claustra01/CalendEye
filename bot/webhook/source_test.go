package webhook_test

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/claustra01/calendeye/webhook"
)

func TestUnmarshalSource(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected SourceInterface
		wantErr  bool
	}{
		{
			name: "Unmarshal user source",
			data: []byte(`{"type":"user","userId":"123"}`),
			expected: UserSource{
				Source: Source{
					Type: "user",
				},
				UserId: "123",
			},
			wantErr: false,
		},
		{
			name: "Unmarshal unknown source",
			data: []byte(`{"type":"unknown","someKey":"someValue"}`),
			expected: UnknownSource{
				Type: "unknown",
				Raw: map[string]json.RawMessage{
					"type":    json.RawMessage(`"unknown"`),
					"someKey": json.RawMessage(`"someValue"`),
				},
			},
			wantErr: false,
		},
		{
			name:     "Invalid JSON",
			data:     []byte(`invalid json`),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Missing type field",
			data:     []byte(`{"userId":"123"}`),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalSource(tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("UnmarshalSource() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
