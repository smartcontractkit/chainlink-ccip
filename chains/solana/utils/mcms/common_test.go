package mcms

import (
	"fmt"
	"testing"
)

func TestPadString32(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		shouldFail     bool
		expectedOutput string
	}{
		{
			name:           "basic case - test-mcm",
			input:          "test-mcm",
			shouldFail:     false,
			expectedOutput: "000000000000000000000000000000000000000000000000746573742d6d636d",
		},
		{
			name:           "empty string",
			input:          "",
			shouldFail:     false,
			expectedOutput: "0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			name:           "single character",
			input:          "a",
			shouldFail:     false,
			expectedOutput: "0000000000000000000000000000000000000000000000000000000000000061",
		},
		{
			name:           "exactly 32 bytes",
			input:          "12345678901234567890123456789012",
			shouldFail:     false,
			expectedOutput: "3132333435363738393031323334353637383930313233343536373839303132",
		},
		{
			name:       "too long - over 32 bytes",
			input:      "123456789012345678901234567890123",
			shouldFail: true,
		},
		{
			name:           "special characters",
			input:          "test@#$%",
			shouldFail:     false,
			expectedOutput: "0000000000000000000000000000000000000000000000007465737440232425",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := PadString32(tt.input)

			if (err != nil) != tt.shouldFail {
				t.Errorf("PadString32() error = %v, wantError %v", err, tt.shouldFail)
				return
			}

			if tt.shouldFail {
				return
			}

			if len(result) != 32 {
				t.Errorf("PadString32() returned buffer length = %d, want 32", len(result))
			}

			if tt.expectedOutput != "" {
				gotHex := fmt.Sprintf("%x", result)
				if gotHex != tt.expectedOutput {
					t.Errorf("PadString32() = %v, want %v", gotHex, tt.expectedOutput)
				}
			}
		})
	}
}
