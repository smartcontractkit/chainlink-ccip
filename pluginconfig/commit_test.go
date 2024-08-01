package pluginconfig

import (
	"math/big"
	"testing"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
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
				DestChain: cciptypes.ChainSelector(1),
				PricedTokens: []types.Account{
					types.Account("0x123"),
					types.Account("0x124"),
				},
				NewMsgScanBatchSize: 256,
				TokenPricesObserver: true,
			},
			expErr: false,
		},
		{
			name: "dest chain is empty",
			input: CommitPluginConfig{
				PricedTokens: []types.Account{
					types.Account("0x123"),
					types.Account("0x124"),
				},
				NewMsgScanBatchSize: 256,
				TokenPricesObserver: true,
			},
			expErr: true,
		},
		{
			name: "zero priced tokens",
			input: CommitPluginConfig{
				DestChain:           cciptypes.ChainSelector(1),
				NewMsgScanBatchSize: 256,
				TokenPricesObserver: true,
			},
			expErr: true,
		},
		{
			name: "empty batch scan size",
			input: CommitPluginConfig{
				DestChain: cciptypes.ChainSelector(1),
				PricedTokens: []types.Account{
					types.Account("0x123"),
					types.Account("0x124"),
				},
				TokenPricesObserver: true,
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
			},
			false,
		},
		{
			"valid, eip-55 address without 0x prefix",
			fields{
				AggregatorAddress: "2e03388D351BF87CF2409EFf18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
			},
			false,
		},
		{
			"valid, lowercase address",
			fields{
				AggregatorAddress: "0x2e03388d351bf87cf2409eff18c45df59775fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
			},
			false,
		},
		{
			"invalid, empty address",
			fields{
				AggregatorAddress: "",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
			},
			true,
		},
		{
			"invalid, invalid address",
			fields{
				AggregatorAddress: "0x2e03388D351BF87CF2409EFf18C45Df59775b",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
			},
			true,
		},
		{
			"invalid, negative deviation",
			fields{
				AggregatorAddress: "0x2e03388D351BF87CF2409EFf18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(-1)},
			},
			true,
		},
		{
			"invalid, zero deviation",
			fields{
				AggregatorAddress: "0x2e03388D351BF87CF2409EFf18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(0)},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ArbitrumPriceSource{
				AggregatorAddress: tt.fields.AggregatorAddress,
				DeviationPPB:      tt.fields.DeviationPPB,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ArbitrumPriceSource.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommitOffchainConfig_Validate(t *testing.T) {
	type fields struct {
		RemoteGasPriceBatchWriteFrequency commonconfig.Duration
		TokenPriceBatchWriteFrequency     commonconfig.Duration
		PriceSources                      map[types.Account]ArbitrumPriceSource
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
				PriceSources: map[types.Account]ArbitrumPriceSource{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
					},
				},
			},
			false,
		},
		{
			"valid, no token price sources",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(0),
				PriceSources:                      map[types.Account]ArbitrumPriceSource{},
			},
			false,
		},
		{
			"invalid, token price sources with no frequency",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(0),
				PriceSources: map[types.Account]ArbitrumPriceSource{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
					},
				},
			},
			true,
		},
		{
			"invalid, token price frequency with no sources",
			fields{
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
				PriceSources:                      map[types.Account]ArbitrumPriceSource{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommitOffchainConfig{
				RemoteGasPriceBatchWriteFrequency: tt.fields.RemoteGasPriceBatchWriteFrequency,
				TokenPriceBatchWriteFrequency:     tt.fields.TokenPriceBatchWriteFrequency,
				PriceSources:                      tt.fields.PriceSources,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("CommitOffchainConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
