package pluginconfig

import (
	"math/big"
	"testing"
	"time"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestCommitPluginConfigValidate(t *testing.T) {
	testCases := []struct {
		name      string
		input     CommitOffchainConfig
		destChain cciptypes.ChainSelector
		expErr    bool
	}{
		{
			name: "valid cfg",
			input: CommitOffchainConfig{
				NewMsgScanBatchSize:               256,
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
			},
			destChain: cciptypes.ChainSelector(1),
			expErr:    false,
		},
		{
			name: "dest chain is empty",
			input: CommitOffchainConfig{
				NewMsgScanBatchSize: 256,
			},
			expErr: true,
		},
		{
			name: "zero priced tokens",
			input: CommitOffchainConfig{
				NewMsgScanBatchSize: 256,
			},
			destChain: cciptypes.ChainSelector(1),
			expErr:    true,
		},
		{
			name:      "empty batch scan size",
			input:     CommitOffchainConfig{},
			destChain: cciptypes.ChainSelector(1),
			expErr:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.Validate()
			if tc.expErr {
				assert.Error(t, actual)
				return
			}
			assert.NoError(t, actual)
		})
	}
}

func TestArbitrumPriceSource_Validate(t *testing.T) {
	type fields struct {
		AggregatorAddress string
		DeviationPPB      cciptypes.BigInt
		Decimals          uint8
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"valid, eip-55 address",
			fields{
				AggregatorAddress: "0x2e03388D351BF87CF2409EFf18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
				Decimals:          18,
			},
			false,
		},
		{
			"valid, eip-55 address without 0x prefix",
			fields{
				AggregatorAddress: "2e03388D351BF87CF2409EFf18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
				Decimals:          18,
			},
			false,
		},
		{
			"valid, lowercase address",
			fields{
				AggregatorAddress: "0x2e03388d351bf87cf2409eff18c45df59775fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
				Decimals:          18,
			},
			false,
		},
		{
			"invalid, empty address",
			fields{
				AggregatorAddress: "",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
				Decimals:          18,
			},
			true,
		},
		{
			"invalid, invalid address",
			fields{
				AggregatorAddress: "0x2e03388D351BF87CF2409EFf18C45Df59775b",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
				Decimals:          18,
			},
			true,
		},
		{
			"invalid, negative deviation",
			fields{
				AggregatorAddress: "0x2e03388D351BF87CF2409EFf18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(-1)},
				Decimals:          18,
			},
			true,
		},
		{
			"invalid, zero deviation",
			fields{
				AggregatorAddress: "0x2e03388D351BF87CF2409EFf18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(0)},
				Decimals:          18,
			},
			true,
		},
		{
			"invalid, zero decimals",
			fields{
				AggregatorAddress: "0x2e03388D351BF87CF2409EFf18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
				Decimals:          0,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := TokenInfo{
				AggregatorAddress: tt.fields.AggregatorAddress,
				DeviationPPB:      tt.fields.DeviationPPB,
				Decimals:          tt.fields.Decimals,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("TokenInfo.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommitOffchainConfig_Validate(t *testing.T) {
	type fields struct {
		RemoteGasPriceBatchWriteFrequency commonconfig.Duration
		TokenPriceBatchWriteFrequency     commonconfig.Duration
		TokenInfo                         map[types.Account]TokenInfo
		TokenPriceChainSelector           cciptypes.ChainSelector
		TokenDecimals                     map[types.Account]uint8
		NewMsgScanBatchSize               uint32
	}
	//nolint:gosec
	const remoteTokenAddress = "0x260fAB5e97758BaB75C1216873Ec4F88C11E57E3"
	const aggregatorAddress = "0x2e03388D351BF87CF2409EFf18C45Df59775Fbb2"
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"valid, with token price sources",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
				TokenInfo: map[types.Account]TokenInfo{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
						Decimals:          18,
					},
				},
				TokenPriceChainSelector: 10,
				NewMsgScanBatchSize:     256,
			},
			false,
		},
		{
			"valid, no token price sources",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(0),
				TokenInfo:                         map[types.Account]TokenInfo{},
				NewMsgScanBatchSize:               256,
			},
			false,
		},
		{
			"invalid, token price sources with no frequency",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(0),
				TokenInfo: map[types.Account]TokenInfo{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
						Decimals:          18,
					},
				},
				NewMsgScanBatchSize: 256,
			},
			true,
		},
		{
			"invalid, price sources with no chain selector",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
				TokenInfo: map[types.Account]TokenInfo{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
						Decimals:          18,
					},
				},
				NewMsgScanBatchSize: 256,
			},
			true,
		},
		{
			"invalid, no new msg scan batch size",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
				TokenInfo:                         map[types.Account]TokenInfo{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommitOffchainConfig{
				RemoteGasPriceBatchWriteFrequency: tt.fields.RemoteGasPriceBatchWriteFrequency,
				TokenPriceBatchWriteFrequency:     tt.fields.TokenPriceBatchWriteFrequency,
				TokenInfo:                         tt.fields.TokenInfo,
				PriceFeedChainSelector:            tt.fields.TokenPriceChainSelector,
				NewMsgScanBatchSize:               int(tt.fields.NewMsgScanBatchSize),
			}
			err := c.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCommitOffchainConfig_EncodeDecode(t *testing.T) {
	type fields struct {
		RemoteGasPriceBatchWriteFrequency commonconfig.Duration
		TokenPriceBatchWriteFrequency     commonconfig.Duration
		PriceSources                      map[types.Account]TokenInfo
	}
	//nolint:gosec
	const (
		remoteTokenAddress1 = "0x260fAB5e97758BaB75C1216873Ec4F88C11E57E3"
		remoteTokenAddress2 = "0x560fAB5e97758BaB75C1316873Ec4F98C11E57E4"
		aggregatorAddress1  = "0x2e03388D351BF87CF2409EFf18C45Df59775Fbb2"
		aggregatorAddress2  = "0x2e03388D351BF87CF2409EFf18C45Df59775Fbb3"
	)
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"valid, with token price sources",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
				PriceSources: map[types.Account]TokenInfo{
					remoteTokenAddress1: {
						AggregatorAddress: aggregatorAddress1,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
					},
					remoteTokenAddress2: {
						AggregatorAddress: aggregatorAddress2,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(2)},
					},
				},
			},
		},
		{
			"valid, no token price sources",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(0),
				PriceSources:                      map[types.Account]TokenInfo{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommitOffchainConfig{
				RemoteGasPriceBatchWriteFrequency: tt.fields.RemoteGasPriceBatchWriteFrequency,
				TokenPriceBatchWriteFrequency:     tt.fields.TokenPriceBatchWriteFrequency,
				TokenInfo:                         tt.fields.PriceSources,
			}

			encoded, err := EncodeCommitOffchainConfig(c)
			require.NoError(t, err)

			decoded, err := DecodeCommitOffchainConfig(encoded)
			require.NoError(t, err)

			require.Equal(t, c, decoded)
		})
	}
}

func TestCommitOffchainConfig_ApplyDefaults(t *testing.T) {
	tests := []struct {
		name     string
		input    CommitOffchainConfig
		expected CommitOffchainConfig
	}{
		{
			name:  "Empty config",
			input: CommitOffchainConfig{},
			expected: CommitOffchainConfig{
				RMNSignaturesTimeout:               0,
				NewMsgScanBatchSize:                defaultNewMsgScanBatchSize,
				MaxReportTransmissionCheckAttempts: defaultMaxReportTransmissionCheckAttempts,
				MaxMerkleTreeSize:                  EvmDefaultMaxMerkleTreeSize,
			},
		},
		{
			name: "RMN enabled without timeout",
			input: CommitOffchainConfig{
				RMNEnabled: true,
			},
			expected: CommitOffchainConfig{
				RMNEnabled:                         true,
				RMNSignaturesTimeout:               defaultRMNSignaturesTimeout,
				NewMsgScanBatchSize:                defaultNewMsgScanBatchSize,
				MaxReportTransmissionCheckAttempts: defaultMaxReportTransmissionCheckAttempts,
				MaxMerkleTreeSize:                  EvmDefaultMaxMerkleTreeSize,
			},
		},
		{
			name: "Custom values",
			input: CommitOffchainConfig{
				RMNEnabled:                         true,
				RMNSignaturesTimeout:               5 * time.Minute,
				NewMsgScanBatchSize:                500,
				MaxReportTransmissionCheckAttempts: 10,
				MaxMerkleTreeSize:                  1000,
			},
			expected: CommitOffchainConfig{
				RMNEnabled:                         true,
				RMNSignaturesTimeout:               5 * time.Minute,
				NewMsgScanBatchSize:                500,
				MaxReportTransmissionCheckAttempts: 10,
				MaxMerkleTreeSize:                  1000,
			},
		},
		{
			name: "Partial custom values",
			input: CommitOffchainConfig{
				RMNEnabled:          true,
				NewMsgScanBatchSize: 300,
				MaxMerkleTreeSize:   500,
			},
			expected: CommitOffchainConfig{
				RMNEnabled:                         true,
				RMNSignaturesTimeout:               defaultRMNSignaturesTimeout,
				NewMsgScanBatchSize:                300,
				MaxReportTransmissionCheckAttempts: defaultMaxReportTransmissionCheckAttempts,
				MaxMerkleTreeSize:                  500,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := tt.input
			config.ApplyDefaults()
			assert.Equal(t, tt.expected, config)
		})
	}
}

func TestCommitOffchainConfig_ApplyDefaultsAndValidate(t *testing.T) {
	tests := []struct {
		name          string
		input         CommitOffchainConfig
		expectedError string
	}{
		{
			name:  "Empty config applies defaults and validates successfully",
			input: CommitOffchainConfig{},
		},
		{
			name: "Config with some values set applies remaining defaults and validates successfully",
			input: CommitOffchainConfig{
				NewMsgScanBatchSize: 100,
				RMNEnabled:          true,
			},
		},
		{
			name: "Config with all valid values doesn't change and validates successfully",
			input: CommitOffchainConfig{
				RemoteGasPriceBatchWriteFrequency:  *commonconfig.MustNewDuration(2 * time.Minute),
				RMNSignaturesTimeout:               10 * time.Second,
				NewMsgScanBatchSize:                100,
				MaxReportTransmissionCheckAttempts: 10,
				MaxMerkleTreeSize:                  1000,
				RMNEnabled:                         true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := tt.input
			err := config.ApplyDefaultsAndValidate()

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				require.NoError(t, err)

				// Check that defaults were applied
				assert.NotZero(t, config.RemoteGasPriceBatchWriteFrequency.Duration())
				assert.NotZero(t, config.NewMsgScanBatchSize)
				assert.NotZero(t, config.MaxReportTransmissionCheckAttempts)
				assert.NotZero(t, config.MaxMerkleTreeSize)

				// Check specific defaults
				if !tt.input.RMNEnabled {
					assert.Zero(t, config.RMNSignaturesTimeout)
				} else {
					assert.NotZero(t, config.RMNSignaturesTimeout)
				}
			}
		})
	}
}
