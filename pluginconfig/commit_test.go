package pluginconfig

import (
	"math/big"
	"testing"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"
)

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
		NewMsgScanBatchSize               int
		RemoteGasPriceBatchWriteFrequency commonconfig.Duration
		TokenPriceBatchWriteFrequency     commonconfig.Duration
		PriceSources                      map[types.Account]ArbitrumPriceSource
		TokenPriceChainSelector           uint64
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
				NewMsgScanBatchSize:               256,
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
				PriceSources: map[types.Account]ArbitrumPriceSource{
					remoteTokenAddress: {
						AggregatorAddress: aggregatorAddress,
						DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
					},
				},
				TokenPriceChainSelector: 10,
			},
			false,
		},
		{
			"valid, no token price sources",
			fields{
				NewMsgScanBatchSize:               256,
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(0),
				PriceSources:                      map[types.Account]ArbitrumPriceSource{},
			},
			false,
		},
		{
			"invalid, token price sources with no frequency",
			fields{
				NewMsgScanBatchSize:               256,
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
			"invalid, price sources with no chain selector",
			fields{
				NewMsgScanBatchSize:               256,
				RemoteGasPriceBatchWriteFrequency: *commonconfig.MustNewDuration(1),
				TokenPriceBatchWriteFrequency:     *commonconfig.MustNewDuration(1),
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
			"invalid, new msg scan batch size not provided",
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
				NewMsgScanBatchSize:               tt.fields.NewMsgScanBatchSize,
				RemoteGasPriceBatchWriteFrequency: tt.fields.RemoteGasPriceBatchWriteFrequency,
				TokenPriceBatchWriteFrequency:     tt.fields.TokenPriceBatchWriteFrequency,
				PriceSources:                      tt.fields.PriceSources,
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
		PriceSources                      map[types.Account]ArbitrumPriceSource
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
				PriceSources: map[types.Account]ArbitrumPriceSource{
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
				PriceSources:                      map[types.Account]ArbitrumPriceSource{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommitOffchainConfig{
				RemoteGasPriceBatchWriteFrequency: tt.fields.RemoteGasPriceBatchWriteFrequency,
				TokenPriceBatchWriteFrequency:     tt.fields.TokenPriceBatchWriteFrequency,
				PriceSources:                      tt.fields.PriceSources,
			}

			encoded, err := EncodeCommitOffchainConfig(c)
			require.NoError(t, err)

			decoded, err := DecodeCommitOffchainConfig(encoded)
			require.NoError(t, err)

			require.Equal(t, c, decoded)
		})
	}
}
