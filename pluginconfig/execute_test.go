package pluginconfig

import (
	"testing"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestExecuteOffchainConfig_Validate(t *testing.T) {
	type fields struct {
		BatchGasLimit             uint64
		RelativeBoostPerWaitHour  float64
		InflightCacheExpiry       commonconfig.Duration
		RootSnoozeTime            commonconfig.Duration
		MessageVisibilityInterval commonconfig.Duration
		BatchingStrategyID        uint32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"valid",
			fields{
				BatchGasLimit:             1,
				RelativeBoostPerWaitHour:  1,
				InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
				RootSnoozeTime:            *commonconfig.MustNewDuration(1),
				MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
				BatchingStrategyID:        0,
			},
			false,
		},
		{
			"invalid, BatchGasLimit not set",
			fields{
				BatchGasLimit:             0,
				RelativeBoostPerWaitHour:  1,
				InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
				RootSnoozeTime:            *commonconfig.MustNewDuration(1),
				MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
				BatchingStrategyID:        0,
			},
			true,
		},
		{
			"invalid, RelativeBoostPerWaitHour not set",
			fields{
				BatchGasLimit:             1,
				RelativeBoostPerWaitHour:  0,
				InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
				RootSnoozeTime:            *commonconfig.MustNewDuration(1),
				MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
				BatchingStrategyID:        0,
			},
			true,
		},
		{
			"invalid, InflightCacheExpiry not set",
			fields{
				BatchGasLimit:             1,
				RelativeBoostPerWaitHour:  1,
				InflightCacheExpiry:       *commonconfig.MustNewDuration(0),
				RootSnoozeTime:            *commonconfig.MustNewDuration(1),
				MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
				BatchingStrategyID:        0,
			},
			true,
		},
		{
			"invalid, RootSnoozeTime not set",
			fields{
				BatchGasLimit:             1,
				RelativeBoostPerWaitHour:  1,
				InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
				RootSnoozeTime:            *commonconfig.MustNewDuration(0),
				MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
				BatchingStrategyID:        0,
			},
			true,
		},
		{
			"invalid, MessageVisibilityInterval not set",
			fields{
				BatchGasLimit:             1,
				RelativeBoostPerWaitHour:  1,
				InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
				RootSnoozeTime:            *commonconfig.MustNewDuration(1),
				MessageVisibilityInterval: *commonconfig.MustNewDuration(0),
				BatchingStrategyID:        0,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ExecuteOffchainConfig{
				BatchGasLimit:             tt.fields.BatchGasLimit,
				RelativeBoostPerWaitHour:  tt.fields.RelativeBoostPerWaitHour,
				InflightCacheExpiry:       tt.fields.InflightCacheExpiry,
				RootSnoozeTime:            tt.fields.RootSnoozeTime,
				MessageVisibilityInterval: tt.fields.MessageVisibilityInterval,
				BatchingStrategyID:        tt.fields.BatchingStrategyID,
			}
			if err := e.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ExecuteOffchainConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExecuteOffchainConfig_EncodeDecode(t *testing.T) {
	type fields struct {
		BatchGasLimit             uint64
		RelativeBoostPerWaitHour  float64
		InflightCacheExpiry       commonconfig.Duration
		RootSnoozeTime            commonconfig.Duration
		MessageVisibilityInterval commonconfig.Duration
		BatchingStrategyID        uint32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"valid",
			fields{
				BatchGasLimit:             1,
				RelativeBoostPerWaitHour:  1,
				InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
				RootSnoozeTime:            *commonconfig.MustNewDuration(1),
				MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
				BatchingStrategyID:        0,
			},
		},
		{
			"valid, boost with decimal places",
			fields{
				BatchGasLimit:             1,
				RelativeBoostPerWaitHour:  1.523,
				InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
				RootSnoozeTime:            *commonconfig.MustNewDuration(1),
				MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
				BatchingStrategyID:        0,
			},
		},
		{
			"valid, large gas limit, zero boost",
			fields{
				BatchGasLimit:             1 << 63,
				RelativeBoostPerWaitHour:  0.0,
				InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
				RootSnoozeTime:            *commonconfig.MustNewDuration(1),
				MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
				BatchingStrategyID:        0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ExecuteOffchainConfig{
				BatchGasLimit:             tt.fields.BatchGasLimit,
				RelativeBoostPerWaitHour:  tt.fields.RelativeBoostPerWaitHour,
				InflightCacheExpiry:       tt.fields.InflightCacheExpiry,
				RootSnoozeTime:            tt.fields.RootSnoozeTime,
				MessageVisibilityInterval: tt.fields.MessageVisibilityInterval,
				BatchingStrategyID:        tt.fields.BatchingStrategyID,
			}
			encoded, err := EncodeExecuteOffchainConfig(e)
			require.NoError(t, err)

			decoded, err := DecodeExecuteOffchainConfig(encoded)
			require.NoError(t, err)

			require.Equal(t, e, decoded)
		})
	}
}
