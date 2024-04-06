package webhook_test

import (
	"reflect"
	"testing"

	. "github.com/claustra01/calendeye/webhook"
)

func TestUnmarshalEvent(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected EventInterface
		wantErr  bool
	}{
		{
			name:     "Invalid type field",
			data:     []byte(`{"mode":"unknown","timestamp":123456789}`),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Invalid JSON",
			data:     []byte(`invalid json`),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalEvent(tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("UnmarshalEvent() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
