package pluginconfig

import (
	"fmt"
	"testing"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/require"
)

func Test_TokenDataObserver_Unmarshall(t *testing.T) {
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
		want    []TokenDataObserverConfig
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid config without token data observers",
			json: ``,
			want: nil,
		},
		{
			name: "invalid config with unknown token data observer type",
			json: `"tokenDataObservers": [
							{
							  "type": "usdc-but-different",
							  "version": "1.0",
							  "param": "500ms"
							}
				  	],`,
			wantErr: true,
			errMsg:  "unknown token data observer type",
		},
		{
			name: "usdc type is set but struct is not initialized properly",
			json: `"tokenDataObservers": [
							{
							  "type": "usdc-cctp",
							  "version": "1.0",
							  "param": "500ms"
							}
				  	],`,
			wantErr: true,
			errMsg:  "USDCCCTPObserverConfig is empty",
		},
		{
			name: "valid config with USDCCCTPObserverConfig",
			json: `"tokenDataObservers": [
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
			want: []TokenDataObserverConfig{
				{
					Type:    "usdc-cctp",
					Version: "1.0",
					USDCCCTPObserverConfig: &USDCCCTPObserverConfig{
						Tokens: map[cciptypes.ChainSelector]USDCCCTPTokenConfig{
							1: {
								SourcePoolAddress:            "0xabc",
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
			name: "valid config with multiple tokens per USDCCCTPObserverConfig",
			json: `"tokenDataObservers": [
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
			want: []TokenDataObserverConfig{
				{
					Type:    "usdc-cctp",
					Version: "1.0",
					USDCCCTPObserverConfig: &USDCCCTPObserverConfig{
						Tokens: map[cciptypes.ChainSelector]USDCCCTPTokenConfig{
							1: {
								SourcePoolAddress:            "0xabc",
								SourceMessageTransmitterAddr: "0xefg",
							},
							20: {
								SourcePoolAddress:            "0x123",
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
				require.Equal(t, tt.want, e.TokenDataObservers)
			}
		})
	}
}

func Test_TokenDataObserver_Validation(t *testing.T) {
	withBaseConfig := func(configs ...TokenDataObserverConfig) ExecuteOffchainConfig {
		return ExecuteOffchainConfig{
			BatchGasLimit:             1,
			RelativeBoostPerWaitHour:  1,
			InflightCacheExpiry:       *commonconfig.MustNewDuration(1),
			RootSnoozeTime:            *commonconfig.MustNewDuration(1),
			MessageVisibilityInterval: *commonconfig.MustNewDuration(1),
			BatchingStrategyID:        0,
			TokenDataObservers:        configs,
		}
	}

	withUSDCConfig := func() *USDCCCTPObserverConfig {
		return &USDCCCTPObserverConfig{
			Tokens: map[cciptypes.ChainSelector]USDCCCTPTokenConfig{
				1: {
					SourcePoolAddress:            "0xabc",
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
			name:   "valid config without token data observers",
			config: withBaseConfig(),
		},
		{
			name: "invalid config with unknown token data observer type",
			config: withBaseConfig(
				TokenDataObserverConfig{
					Type:    "my-fancy-token",
					Version: "1.0",
				}),
			wantErr: true,
			errMsg:  "unknown token data observer type",
		},
		{
			name: "usdc type is set but struct is empty",
			config: withBaseConfig(
				TokenDataObserverConfig{
					Type:                   "usdc-cctp",
					Version:                "1.0",
					USDCCCTPObserverConfig: &USDCCCTPObserverConfig{},
				}),
			wantErr: true,
			errMsg:  "AttestationAPI not set",
		},
		{
			name: "usdc type is set but tokens are missing",
			config: withBaseConfig(
				TokenDataObserverConfig{
					Type:    "usdc-cctp",
					Version: "1.0",
					USDCCCTPObserverConfig: &USDCCCTPObserverConfig{
						AttestationAPI:         "http://localhost:8080",
						AttestationAPITimeout:  commonconfig.MustNewDuration(time.Second),
						AttestationAPIInterval: commonconfig.MustNewDuration(500 * time.Millisecond),
					},
				}),
			wantErr: true,
			errMsg:  "Tokens not set",
		},
		{
			name: "the same observer can't bet set twice",
			config: withBaseConfig(
				TokenDataObserverConfig{
					Type:                   "usdc-cctp",
					Version:                "1.0",
					USDCCCTPObserverConfig: withUSDCConfig(),
				},
				TokenDataObserverConfig{
					Type:                   "usdc-cctp",
					Version:                "1.0",
					USDCCCTPObserverConfig: withUSDCConfig(),
				}),
			wantErr: true,
			errMsg:  "duplicate token data observer type",
		},
		{
			name: "valid config with multiple the same observers types but different versions",
			config: withBaseConfig(
				TokenDataObserverConfig{
					Type:                   "usdc-cctp",
					Version:                "1.0",
					USDCCCTPObserverConfig: withUSDCConfig(),
				},
				TokenDataObserverConfig{
					Type:                   "usdc-cctp",
					Version:                "2.0",
					USDCCCTPObserverConfig: withUSDCConfig(),
				},
			),
		},
		{
			name: "valid config with single usdc observer",
			config: withBaseConfig(
				TokenDataObserverConfig{
					Type:                   "usdc-cctp",
					Version:                "1.0",
					USDCCCTPObserverConfig: withUSDCConfig(),
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
