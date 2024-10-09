package pluginconfig

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestArbitrumPriceSource_Validate(t *testing.T) {
	type fields struct {
		AggregatorAddress cciptypes.UnknownEncodedAddress
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
		RemoteGasPriceBatchWriteFrequency  commonconfig.Duration
		TokenPriceBatchWriteFrequency      commonconfig.Duration
		TokenInfo                          map[cciptypes.UnknownEncodedAddress]TokenInfo
		TokenPriceChainSelector            cciptypes.ChainSelector
		TokenDecimals                      map[cciptypes.UnknownEncodedAddress]uint8
		NewMsgScanBatchSize                uint32
		MaxReportTransmissionCheckAttempts uint32
		MaxMerkleTreeSize                  uint32
		SignObservationPrefix              string
	}
	remoteTokenAddress := rand.RandomAddress()
	aggregatorAddress := rand.RandomAddress()
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
				TokenInfo: map[cciptypes.UnknownEncodedAddress]TokenInfo{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
						Decimals:          18,
					},
				},
				TokenPriceChainSelector:            10,
				NewMsgScanBatchSize:                256,
				MaxReportTransmissionCheckAttempts: 10,
				MaxMerkleTreeSize:                  1000,
				SignObservationPrefix:              defaultSignObservationPrefix,
			},
			false,
		},
		{
			"valid, no token price sources",
			fields{
				RemoteGasPriceBatchWriteFrequency:  *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:      *commonconfig.MustNewDuration(0),
				TokenInfo:                          map[cciptypes.UnknownEncodedAddress]TokenInfo{},
				NewMsgScanBatchSize:                256,
				MaxReportTransmissionCheckAttempts: 10,
				MaxMerkleTreeSize:                  1000,
				SignObservationPrefix:              defaultSignObservationPrefix,
			},
			false,
		},
		{
			"invalid, token price sources with no frequency",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(0),
				TokenInfo: map[cciptypes.UnknownEncodedAddress]TokenInfo{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
						Decimals:          18,
					},
				},
				NewMsgScanBatchSize:                256,
				MaxReportTransmissionCheckAttempts: 10,
				MaxMerkleTreeSize:                  1000,
				SignObservationPrefix:              defaultSignObservationPrefix,
			},
			true,
		},
		{
			"invalid, price sources with no chain selector",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
				TokenInfo: map[cciptypes.UnknownEncodedAddress]TokenInfo{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
						Decimals:          18,
					},
				},
				NewMsgScanBatchSize:                256,
				MaxReportTransmissionCheckAttempts: 10,
				MaxMerkleTreeSize:                  1000,
				SignObservationPrefix:              defaultSignObservationPrefix,
			},
			true,
		},
		{
			"invalid, no new msg scan batch size",
			fields{
				RemoteGasPriceBatchWriteFrequency:  *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:      *commonconfig.MustNewDuration(1),
				TokenInfo:                          map[cciptypes.UnknownEncodedAddress]TokenInfo{},
				MaxReportTransmissionCheckAttempts: 10,
				MaxMerkleTreeSize:                  1000,
				SignObservationPrefix:              defaultSignObservationPrefix,
			},
			true,
		},
		{
			"invalid, missing SignObservationPrefix",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
				TokenInfo: map[cciptypes.UnknownEncodedAddress]TokenInfo{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
						Decimals:          18,
					},
				},
				TokenPriceChainSelector:            10,
				NewMsgScanBatchSize:                256,
				MaxReportTransmissionCheckAttempts: 10,
				MaxMerkleTreeSize:                  1000,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommitOffchainConfig{
				RemoteGasPriceBatchWriteFrequency:  tt.fields.RemoteGasPriceBatchWriteFrequency,
				TokenPriceBatchWriteFrequency:      tt.fields.TokenPriceBatchWriteFrequency,
				TokenInfo:                          tt.fields.TokenInfo,
				PriceFeedChainSelector:             tt.fields.TokenPriceChainSelector,
				NewMsgScanBatchSize:                int(tt.fields.NewMsgScanBatchSize),
				MaxReportTransmissionCheckAttempts: uint(tt.fields.MaxReportTransmissionCheckAttempts),
				MaxMerkleTreeSize:                  uint64(tt.fields.MaxMerkleTreeSize),
				SignObservationPrefix:              tt.fields.SignObservationPrefix,
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
		PriceSources                      map[cciptypes.UnknownEncodedAddress]TokenInfo
	}
	remoteTokenAddress1 := rand.RandomAddress()
	remoteTokenAddress2 := rand.RandomAddress()
	aggregatorAddress1 := rand.RandomAddress()
	aggregatorAddress2 := rand.RandomAddress()
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"valid, with token price sources",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
				PriceSources: map[cciptypes.UnknownEncodedAddress]TokenInfo{
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
				PriceSources:                      map[cciptypes.UnknownEncodedAddress]TokenInfo{},
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
				MaxMerkleTreeSize:                  defaultEvmDefaultMaxMerkleTreeSize,
				RemoteGasPriceBatchWriteFrequency:  *commonconfig.MustNewDuration(defaultRemoteGasPriceBatchWriteFrequency),
				SignObservationPrefix:              defaultSignObservationPrefix,
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
				MaxMerkleTreeSize:                  defaultEvmDefaultMaxMerkleTreeSize,
				RemoteGasPriceBatchWriteFrequency:  *commonconfig.MustNewDuration(defaultRemoteGasPriceBatchWriteFrequency),
				SignObservationPrefix:              defaultSignObservationPrefix,
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
				RemoteGasPriceBatchWriteFrequency:  *commonconfig.MustNewDuration(defaultRemoteGasPriceBatchWriteFrequency),
				SignObservationPrefix:              defaultSignObservationPrefix,
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
				RemoteGasPriceBatchWriteFrequency:  *commonconfig.MustNewDuration(defaultRemoteGasPriceBatchWriteFrequency),
				SignObservationPrefix:              defaultSignObservationPrefix,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := tt.input
			config.applyDefaults()
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
				SignObservationPrefix:              defaultSignObservationPrefix,
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
				assert.NotZero(t, config.SignObservationPrefix)

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
