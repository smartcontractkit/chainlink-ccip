package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestRMNRemoteConfig_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		config   cciptypes.RemoteConfig
		expected bool
	}{
		{
			name:     "Completely empty config",
			config:   cciptypes.RemoteConfig{},
			expected: true,
		},
		{
			name: "Config with only ContractAddress",
			config: cciptypes.RemoteConfig{
				ContractAddress: cciptypes.UnknownAddress{1, 2, 3},
			},
			expected: true,
		},
		{
			name: "Config with only ConfigDigest",
			config: cciptypes.RemoteConfig{
				ConfigDigest: cciptypes.Bytes32{1},
			},
			expected: false,
		},
		{
			name: "Config with only Signers",
			config: cciptypes.RemoteConfig{
				Signers: []cciptypes.RemoteSignerInfo{{}, {}},
			},
			expected: false,
		},
		{
			name: "Config with only F",
			config: cciptypes.RemoteConfig{
				FSign: 1,
			},
			expected: false,
		},
		{
			name: "Config with only ConfigVersion",
			config: cciptypes.RemoteConfig{
				ConfigVersion: 1,
			},
			expected: false,
		},
		{
			name: "Config with only RmnReportVersion",
			config: cciptypes.RemoteConfig{
				RmnReportVersion: cciptypes.Bytes32{1},
			},
			expected: false,
		},
		{
			name: "Fully populated config",
			config: cciptypes.RemoteConfig{
				ContractAddress:  cciptypes.UnknownAddress{1, 2, 3},
				ConfigDigest:     cciptypes.Bytes32{1},
				Signers:          []cciptypes.RemoteSignerInfo{{}, {}},
				FSign:            2,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{1},
			},
			expected: false,
		},
		{
			name: "Config with nil ContractAddress",
			config: cciptypes.RemoteConfig{
				ContractAddress:  nil,
				ConfigDigest:     cciptypes.Bytes32{1},
				Signers:          []cciptypes.RemoteSignerInfo{{}, {}},
				FSign:            2,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{1},
			},
			expected: false,
		},
		{
			name: "Config with empty (non-nil) ContractAddress",
			config: cciptypes.RemoteConfig{
				ContractAddress:  cciptypes.UnknownAddress{},
				ConfigDigest:     cciptypes.Bytes32{1},
				Signers:          []cciptypes.RemoteSignerInfo{{}, {}},
				FSign:            2,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{1},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.config.IsEmpty()
			assert.Equal(t, tt.expected, result)
		})
	}
}
