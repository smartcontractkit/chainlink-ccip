package pluginconfig

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
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
		TokenDataProcessor        []TokenDataProcessor
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
				TokenDataProcessors:       tt.fields.TokenDataProcessor,
			}
			encoded, err := EncodeExecuteOffchainConfig(e)
			require.NoError(t, err)

			decoded, err := DecodeExecuteOffchainConfig(encoded)
			require.NoError(t, err)

			require.Equal(t, e, decoded)
		})
	}
}

func Test_TokenDataProcessor_Unmarshall(t *testing.T) {
	baseJSON := `{
					  %s
					  "batchGasLimit": 1,
					  "relativeBoostPerWaitHour": 1,
					  "inflightCacheExpiry": "1s",
					  "rootSnoozeTime": "1s",
					  "messageVisibilityInterval": "1s",
					  "batchingStrategyID": 0
				}`

	tests := []struct {
		name    string
		json    string
		want    []TokenDataProcessor
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid config without token data processors",
			json: ``,
			want: nil,
		},
		{
			name: "invalid config with unknown token data processor type",
			json: `"tokenDataProcessors": [
							{
							  "type": "usdc-but-different",
							  "version": "1.0",
							  "param": "500ms"
							}
				  	],`,
			wantErr: true,
			errMsg:  "unknown token data processor type",
		},
		{
			name: "usdc type is set but struct is not initialized properly",
			json: `"tokenDataProcessors": [
							{
							  "type": "usdc-cctp",
							  "version": "1.0",
							  "param": "500ms"
							}
				  	],`,
			wantErr: true,
			errMsg:  "UsdcCctpTokenDataProcessor is empty",
		},
		{
			name: "valid config with UsdcCctpTokenDataProcessor",
			json: `"tokenDataProcessors": [
							{
							  "type": "usdc-cctp",
							  "version": "1.0",
							  "tokens": {
								"1": {
								  "sourceTokenAddress": "0xabc",
								  "sourceMessageTransmitterAddress": "0xefg"
								}
							  },
							  "attestationAPI": "http://localhost:8080",
							  "attestationAPITimeout": "1s",
							  "attestationAPIInterval": "500ms"
							}
				  	],`,
			want: []TokenDataProcessor{
				{
					Type:    "usdc-cctp",
					Version: "1.0",
					UsdcCctpTokenDataProcessor: &UsdcCctpTokenDataProcessor{
						Tokens: map[int]UsdcCctpToken{
							1: {
								SourceTokenAddress:           "0xabc",
								SourceMessageTransmitterAddr: "0xefg",
							},
						},
						AttestationAPI:         "http://localhost:8080",
						AttestationAPITimeout:  commonconfig.MustNewDuration(time.Second),
						AttestationAPIInterval: commonconfig.MustNewDuration(500 * time.Millisecond),
					},
				},
			},
		},
		{
			name: "valid config with multiple tokens per UsdcCctpTokenDataProcessor",
			json: `"tokenDataProcessors": [
							{
							  "type": "usdc-cctp",
							  "version": "1.0",
							  "tokens": {
								"1": {
								  "sourceTokenAddress": "0xabc",
								  "sourceMessageTransmitterAddress": "0xefg"
								},
								"20": {
								  "sourceTokenAddress": "0x123",
								  "sourceMessageTransmitterAddress": "0x456"
								}
							  },
							  "attestationAPI": "http://localhost:8080",
							  "attestationAPITimeout": "1s",
							  "attestationAPIInterval": "500ms"
							}
				  	],`,
			want: []TokenDataProcessor{
				{
					Type:    "usdc-cctp",
					Version: "1.0",
					UsdcCctpTokenDataProcessor: &UsdcCctpTokenDataProcessor{
						Tokens: map[int]UsdcCctpToken{
							1: {
								SourceTokenAddress:           "0xabc",
								SourceMessageTransmitterAddr: "0xefg",
							},
							20: {
								SourceTokenAddress:           "0x123",
								SourceMessageTransmitterAddr: "0x456",
							},
						},
						AttestationAPI:         "http://localhost:8080",
						AttestationAPITimeout:  commonconfig.MustNewDuration(time.Second),
						AttestationAPIInterval: commonconfig.MustNewDuration(500 * time.Millisecond),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			finalJSON := fmt.Sprintf(baseJSON, tt.json)
			e, err := DecodeExecuteOffchainConfig([]byte(finalJSON))

			if tt.wantErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, e.TokenDataProcessors)
			}
		})
	}
}

func Test_TokenDataProcessor_Validation(t *testing.T) {
	withBaseConfig := func(processors ...TokenDataProcessor) ExecuteOffchainConfig {
		return ExecuteOffchainConfig{
			BatchGasLimit:             1,
			RelativeBoostPerWaitHour:  1,
			InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
			RootSnoozeTime:            *commonconfig.MustNewDuration(1),
			MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
			BatchingStrategyID:        0,
			TokenDataProcessors:       processors,
		}
	}

	withUSDCConfig := func() *UsdcCctpTokenDataProcessor {
		return &UsdcCctpTokenDataProcessor{
			Tokens: map[int]UsdcCctpToken{
				1: {
					SourceTokenAddress:           "0xabc",
					SourceMessageTransmitterAddr: "0xefg",
				},
			},
			AttestationAPI:         "http://localhost:8080",
			AttestationAPITimeout:  commonconfig.MustNewDuration(time.Second),
			AttestationAPIInterval: commonconfig.MustNewDuration(500 * time.Millisecond),
		}
	}

	tests := []struct {
		name    string
		config  ExecuteOffchainConfig
		wantErr bool
		errMsg  string
	}{
		{
			name:   "valid config without token data processors",
			config: withBaseConfig(),
		},
		{
			name: "invalid config with unknown token data processor type",
			config: withBaseConfig(
				TokenDataProcessor{
					Type:    "my-fancy-token-processor",
					Version: "1.0",
				}),
			wantErr: true,
			errMsg:  "unknown token data processor type",
		},
		{
			name: "usdc type is set but struct is empty",
			config: withBaseConfig(
				TokenDataProcessor{
					Type:                       "usdc-cctp",
					Version:                    "1.0",
					UsdcCctpTokenDataProcessor: &UsdcCctpTokenDataProcessor{},
				}),
			wantErr: true,
			errMsg:  "AttestationAPI not set",
		},
		{
			name: "usdc type is set but tokens are missing",
			config: withBaseConfig(
				TokenDataProcessor{
					Type:    "usdc-cctp",
					Version: "1.0",
					UsdcCctpTokenDataProcessor: &UsdcCctpTokenDataProcessor{
						AttestationAPI:         "http://localhost:8080",
						AttestationAPITimeout:  commonconfig.MustNewDuration(time.Second),
						AttestationAPIInterval: commonconfig.MustNewDuration(500 * time.Millisecond),
					},
				}),
			wantErr: true,
			errMsg:  "Tokens not set",
		},
		{
			name: "the same processor can't bet set twice",
			config: withBaseConfig(
				TokenDataProcessor{
					Type:                       "usdc-cctp",
					Version:                    "1.0",
					UsdcCctpTokenDataProcessor: withUSDCConfig(),
				},
				TokenDataProcessor{
					Type:                       "usdc-cctp",
					Version:                    "1.0",
					UsdcCctpTokenDataProcessor: withUSDCConfig(),
				}),
			wantErr: true,
			errMsg:  "duplicate token data processor type",
		},
		{
			name: "valid config with multiple the same processors types but different versions",
			config: withBaseConfig(
				TokenDataProcessor{
					Type:                       "usdc-cctp",
					Version:                    "1.0",
					UsdcCctpTokenDataProcessor: withUSDCConfig(),
				},
				TokenDataProcessor{
					Type:                       "usdc-cctp",
					Version:                    "2.0",
					UsdcCctpTokenDataProcessor: withUSDCConfig(),
				},
			),
		},
		{
			name: "valid config with single usdc processor",
			config: withBaseConfig(
				TokenDataProcessor{
					Type:                       "usdc-cctp",
					Version:                    "1.0",
					UsdcCctpTokenDataProcessor: withUSDCConfig(),
				}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
