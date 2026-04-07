package sequences

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/stretchr/testify/assert"
)

func TestReverseTranslateFQ_RoundTrip(t *testing.T) {
	original := lanes.FeeQuoterDestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                30_000,
		MaxPerMsgGasLimit:           4_000_000,
		DestGasOverhead:             350_000,
		DestGasPerPayloadByteBase:   16,
		ChainFamilySelector:         1,
		DefaultTokenFeeUSDCents:     100,
		DefaultTokenDestGasOverhead: 34_000,
		DefaultTxGasLimit:           200_000,
		NetworkFeeUSDCents:          50,
		V1Params: &lanes.FeeQuoterV1Params{
			MaxNumberOfTokensPerMsg:           5,
			DestGasPerPayloadByteHigh:         32,
			DestGasPerPayloadByteThreshold:    1024,
			DestDataAvailabilityOverheadGas:   6000,
			DestGasPerDataAvailabilityByte:    16,
			DestDataAvailabilityMultiplierBps: 10000,
			EnforceOutOfOrder:                 true,
			GasMultiplierWeiPerEth:            12e17,
			GasPriceStalenessThreshold:        86400,
		},
	}

	translated := TranslateFQ(original)
	roundTripped := ReverseTranslateFQ(translated)

	assert.Equal(t, original.IsEnabled, roundTripped.IsEnabled)
	assert.Equal(t, original.MaxDataBytes, roundTripped.MaxDataBytes)
	assert.Equal(t, original.MaxPerMsgGasLimit, roundTripped.MaxPerMsgGasLimit)
	assert.Equal(t, original.DestGasOverhead, roundTripped.DestGasOverhead)
	assert.Equal(t, original.DestGasPerPayloadByteBase, roundTripped.DestGasPerPayloadByteBase)
	assert.Equal(t, original.ChainFamilySelector, roundTripped.ChainFamilySelector)
	assert.Equal(t, original.DefaultTokenFeeUSDCents, roundTripped.DefaultTokenFeeUSDCents)
	assert.Equal(t, original.DefaultTokenDestGasOverhead, roundTripped.DefaultTokenDestGasOverhead)
	assert.Equal(t, original.DefaultTxGasLimit, roundTripped.DefaultTxGasLimit)
	assert.Equal(t, original.NetworkFeeUSDCents, roundTripped.NetworkFeeUSDCents)
	assert.Equal(t, original.V1Params, roundTripped.V1Params)
	assert.Nil(t, roundTripped.V2Params)
}

func TestReverseTranslateFQV2_RoundTrip(t *testing.T) {
	original := lanes.FeeQuoterDestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                50_000,
		MaxPerMsgGasLimit:           8_000_000,
		DestGasOverhead:             450_000,
		DestGasPerPayloadByteBase:   20,
		ChainFamilySelector:         2,
		DefaultTokenFeeUSDCents:     200,
		DefaultTokenDestGasOverhead: 50_000,
		DefaultTxGasLimit:           300_000,
		NetworkFeeUSDCents:          75,
		V2Params: &lanes.FeeQuoterV2Params{
			LinkFeeMultiplierPercent: 120,
		},
	}

	translated := TranslateFQtoV2(original)
	roundTripped := ReverseTranslateFQV2(translated)

	assert.Equal(t, original.IsEnabled, roundTripped.IsEnabled)
	assert.Equal(t, original.MaxDataBytes, roundTripped.MaxDataBytes)
	assert.Equal(t, original.MaxPerMsgGasLimit, roundTripped.MaxPerMsgGasLimit)
	assert.Equal(t, original.DestGasOverhead, roundTripped.DestGasOverhead)
	assert.Equal(t, original.DestGasPerPayloadByteBase, roundTripped.DestGasPerPayloadByteBase)
	assert.Equal(t, original.ChainFamilySelector, roundTripped.ChainFamilySelector)
	assert.Equal(t, original.DefaultTokenFeeUSDCents, roundTripped.DefaultTokenFeeUSDCents)
	assert.Equal(t, original.DefaultTokenDestGasOverhead, roundTripped.DefaultTokenDestGasOverhead)
	assert.Equal(t, original.DefaultTxGasLimit, roundTripped.DefaultTxGasLimit)
	assert.Equal(t, original.NetworkFeeUSDCents, roundTripped.NetworkFeeUSDCents)
	assert.Equal(t, original.V2Params, roundTripped.V2Params)
	assert.Nil(t, roundTripped.V1Params)
}
