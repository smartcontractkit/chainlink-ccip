package pluginconfig

import (
	"math/big"
	"testing"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommitPluginConfigValidate(t *testing.T) {
	testCases := []struct {
		name   string
		input  CommitPluginConfig
		expErr bool
	}{
		{
			name: "valid cfg",
			input: CommitPluginConfig{
				DestChain:           cciptypes.ChainSelector(1),
				NewMsgScanBatchSize: 256,
				OffchainConfig: CommitOffchainConfig{
					RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				},
			},
			expErr: false,
		},
		{
			name: "dest chain is empty",
			input: CommitPluginConfig{
				NewMsgScanBatchSize: 256,
			},
			expErr: true,
		},
		{
			name: "zero priced tokens",
			input: CommitPluginConfig{
				DestChain:           cciptypes.ChainSelector(1),
				NewMsgScanBatchSize: 256,
			},
			expErr: true,
		},
		{
			name: "empty batch scan size",
			input: CommitPluginConfig{
				DestChain: cciptypes.ChainSelector(1),
			},
			expErr: true,
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
		TokenPriceChainSelector           uint64
		TokenDecimals                     map[types.Account]uint8
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
			},
			false,
		},
		{
			"valid, no token price sources",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(0),
				TokenInfo:                         map[types.Account]TokenInfo{},
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
				TokenPriceChainSelector:           tt.fields.TokenPriceChainSelector,
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
