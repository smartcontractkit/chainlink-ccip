package finality

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestRaw_RoundTrips(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		raw    [4]byte
	}{
		{
			name:   "wait for finality (zero value sentinel)",
			config: Config{WaitForFinality: true},
			raw:    [4]byte{0x00, 0x00, 0x00, 0x00},
		},
		{
			name:   "block depth 1",
			config: Config{BlockDepth: 1},
			raw:    [4]byte{0x00, 0x00, 0x00, 0x01},
		},
		{
			name:   "block depth 256",
			config: Config{BlockDepth: 256},
			raw:    [4]byte{0x00, 0x00, 0x01, 0x00},
		},
		{
			name:   "block depth max (65535)",
			config: Config{BlockDepth: 0xFFFF},
			raw:    [4]byte{0x00, 0x00, 0xFF, 0xFF},
		},
		{
			name:   "wait for safe only",
			config: Config{WaitForSafe: true},
			raw:    [4]byte{0x00, 0x01, 0x00, 0x00},
		},
		{
			name:   "wait for safe + block depth (allowed finality combo)",
			config: Config{WaitForSafe: true, BlockDepth: 100},
			raw:    [4]byte{0x00, 0x01, 0x00, 0x64},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.config.Raw()
			assert.Equal(t, tc.raw, got, "Raw() mismatch")
		})
	}
}

func TestFromRaw(t *testing.T) {
	tests := []struct {
		name     string
		raw      [4]byte
		expected Config
	}{
		{
			name:     "zero is WaitForFinality",
			raw:      [4]byte{},
			expected: Config{WaitForFinality: true},
		},
		{
			name:     "block depth 1",
			raw:      [4]byte{0x00, 0x00, 0x00, 0x01},
			expected: Config{BlockDepth: 1},
		},
		{
			name:     "safe flag only",
			raw:      [4]byte{0x00, 0x01, 0x00, 0x00},
			expected: Config{WaitForSafe: true},
		},
		{
			name:     "safe flag + block depth",
			raw:      [4]byte{0x00, 0x01, 0x00, 0x64},
			expected: Config{WaitForSafe: true, BlockDepth: 100},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FromRaw(tc.raw)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestFromRaw_DropsWaitForFinalityWhenOtherModesPresent(t *testing.T) {
	raw := [4]byte{0x00, 0x00, 0x00, 0x05}
	cfg := FromRaw(raw)
	assert.False(t, cfg.WaitForFinality, "WaitForFinality must be false when block depth is set")
	assert.Equal(t, uint16(5), cfg.BlockDepth)
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr string
	}{
		{
			name:    "empty config rejected",
			config:  Config{},
			wantErr: "finality config is empty",
		},
		{
			name:   "wait for finality alone is valid",
			config: Config{WaitForFinality: true},
		},
		{
			name:   "block depth alone is valid",
			config: Config{BlockDepth: 10},
		},
		{
			name:   "wait for safe alone is valid",
			config: Config{WaitForSafe: true},
		},
		{
			name:   "safe + block depth is valid (allowed finality combo)",
			config: Config{WaitForSafe: true, BlockDepth: 50},
		},
		{
			name:    "finality + safe rejected",
			config:  Config{WaitForFinality: true, WaitForSafe: true},
			wantErr: "cannot be combined",
		},
		{
			name:    "finality + block depth rejected",
			config:  Config{WaitForFinality: true, BlockDepth: 5},
			wantErr: "cannot be combined",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.Validate()
			if tc.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIsZero(t *testing.T) {
	assert.True(t, Config{}.IsZero())
	assert.False(t, Config{WaitForFinality: true}.IsZero())
	assert.False(t, Config{BlockDepth: 1}.IsZero())
}

func TestJSON_RoundTrips(t *testing.T) {
	tests := []Config{
		{WaitForFinality: true},
		{BlockDepth: 42},
		{WaitForSafe: true, BlockDepth: 100},
	}
	for _, cfg := range tests {
		data, err := json.Marshal(cfg)
		require.NoError(t, err)
		var got Config
		require.NoError(t, json.Unmarshal(data, &got))
		assert.Equal(t, cfg, got)
	}
}

func TestYAML_RoundTrips(t *testing.T) {
	tests := []Config{
		{WaitForFinality: true},
		{BlockDepth: 42},
		{WaitForSafe: true, BlockDepth: 100},
	}
	for _, cfg := range tests {
		data, err := yaml.Marshal(cfg)
		require.NoError(t, err)
		var got Config
		require.NoError(t, yaml.Unmarshal(data, &got))
		assert.Equal(t, cfg, got)
	}
}
