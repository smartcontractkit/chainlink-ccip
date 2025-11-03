package sequences

import (
	"encoding/binary"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ConfigureLaneLegAsSource = operations.NewSequence(
	"ConfigureLaneLegAsSource",
	semver.MustParse("1.0.0"),
	"Configures lane leg as source on CCIP 1.6.0",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput
		b.Logger.Info("EVM Configuring lane leg as source:", input)

		result, err := sequences.RunAndMergeSequence(b, chains, FeeQuoterApplyDestChainConfigUpdatesSequence, FeeQuoterApplyDestChainConfigUpdatesSequenceInput{
			Address: common.BytesToAddress(input.Source.FeeQuoter),
			UpdatesByChain: map[uint64][]fee_quoter.FeeQuoterDestChainConfigArgs{
				input.Source.Selector: {
					{
						DestChainSelector: input.Dest.Selector,
						DestChainConfig:   TranslateFQ(input.Dest.FeeQuoterDestChainConfig),
					},
				},
			},
		}, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Destination configs updated on FeeQuoters")

		result, err = sequences.RunAndMergeSequence(b, chains, FeeQuoterUpdatePricesSequence, FeeQuoterUpdatePricesSequenceInput{
			Address: common.BytesToAddress(input.Source.FeeQuoter),
			UpdatesByChain: map[uint64]fee_quoter.InternalPriceUpdates{
				input.Source.Selector: {
					TokenPriceUpdates: TranslateTokenPrices(input.Source.TokenPrices),
					GasPriceUpdates: []fee_quoter.InternalGasPriceUpdate{
						{
							DestChainSelector: input.Dest.Selector,
							UsdPerUnitGas:     input.Dest.GasPrice,
						},
					},
				},
			},
		}, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Gas prices updated on FeeQuoters")

		result, err = sequences.RunAndMergeSequence(b, chains, OnRampApplyDestChainConfigUpdatesSequence, OnRampApplyDestChainConfigUpdatesSequenceInput{
			Address: common.BytesToAddress(input.Source.OnRamp),
			UpdatesByChain: map[uint64][]onramp.OnRampDestChainConfigArgs{
				input.Source.Selector: {
					{
						Router:            common.BytesToAddress(input.Source.Router),
						DestChainSelector: input.Dest.Selector,
						AllowlistEnabled:  input.Dest.AllowListEnabled,
					},
				},
			},
		}, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Source configs updated on OffRamps")

		onrampUpdate := router.OnRamp{
			DestChainSelector: input.Dest.Selector,
			OnRamp:            common.BytesToAddress(input.Source.OnRamp),
		}
		if input.IsDisabled {
			onrampUpdate.OnRamp = common.HexToAddress("0x0")
		}
		result, err = sequences.RunAndMergeSequence(b, chains, RouterApplyRampUpdatesSequence, RouterApplyRampUpdatesSequenceInput{
			Address: common.BytesToAddress(input.Source.Router),
			UpdatesByChain: map[uint64]router.ApplyRampsUpdatesArgs{
				input.Source.Selector: {
					OnRampUpdates: []router.OnRamp{onrampUpdate},
				},
			},
		}, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Ramps updated on Routers")

		return result, nil
	},
)

var ConfigureLaneLegAsDest = operations.NewSequence(
	"ConfigureLaneLegAsDest",
	semver.MustParse("1.6.0"),
	"Configures lane leg as destination on CCIP 1.6.0",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput
		b.Logger.Info("EVM Configuring lane leg as destination:", input)

		result, err := sequences.RunAndMergeSequence(b, chains, OffRampApplySourceChainConfigUpdatesSequence, OffRampApplySourceChainConfigUpdatesSequenceInput{
			Address: common.BytesToAddress(input.Dest.OffRamp),
			UpdatesByChain: map[uint64][]offramp.OffRampSourceChainConfigArgs{
				input.Dest.Selector: {
					{
						Router:              common.BytesToAddress(input.Dest.Router),
						SourceChainSelector: input.Source.Selector,
						// https://github.com/smartcontractkit/chainlink/blob/f7ca3d51db51258bb3b8ae22a8e1593d03bc040b/deployment/ccip/changeset/v1_6/cs_chain_contracts.go#L1148
						OnRamp:                    common.LeftPadBytes(input.Source.OnRamp, 32),
						IsEnabled:                 !input.IsDisabled,
						IsRMNVerificationDisabled: !input.Source.RMNVerificationEnabled,
					},
				},
			},
		}, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Destination configs updated on OffRamps")

		offrampUpdate := router.OffRamp{
			SourceChainSelector: input.Source.Selector,
			OffRamp:             common.BytesToAddress(input.Dest.OffRamp),
		}
		var offRampAdds []router.OffRamp
		var offRampRemoves []router.OffRamp
		if input.IsDisabled {
			offRampRemoves = []router.OffRamp{offrampUpdate}
		} else {
			offRampAdds = []router.OffRamp{offrampUpdate}
		}
		result, err = sequences.RunAndMergeSequence(b, chains, RouterApplyRampUpdatesSequence, RouterApplyRampUpdatesSequenceInput{
			Address: common.BytesToAddress(input.Dest.Router),
			UpdatesByChain: map[uint64]router.ApplyRampsUpdatesArgs{
				input.Dest.Selector: {
					OffRampAdds:    offRampAdds,
					OffRampRemoves: offRampRemoves,
				},
			},
		}, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Ramps updated on Routers")

		return result, nil
	},
)

func (a *EVMAdapter) ConfigureLaneLegAsSource() *operations.Sequence[lanes.UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return ConfigureLaneLegAsSource
}

func (a *EVMAdapter) ConfigureLaneLegAsDest() *operations.Sequence[lanes.UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return ConfigureLaneLegAsDest
}

func TranslateFQ(fqc lanes.FeeQuoterDestChainConfig) fee_quoter.FeeQuoterDestChainConfig {
	return fee_quoter.FeeQuoterDestChainConfig{
		IsEnabled:                         fqc.IsEnabled,
		MaxNumberOfTokensPerMsg:           fqc.MaxNumberOfTokensPerMsg,
		MaxDataBytes:                      fqc.MaxDataBytes,
		MaxPerMsgGasLimit:                 fqc.MaxPerMsgGasLimit,
		DestGasOverhead:                   fqc.DestGasOverhead,
		DestGasPerPayloadByteBase:         fqc.DestGasPerPayloadByteBase,
		DestGasPerPayloadByteHigh:         fqc.DestGasPerPayloadByteHigh,
		DestGasPerPayloadByteThreshold:    fqc.DestGasPerPayloadByteThreshold,
		DestDataAvailabilityOverheadGas:   fqc.DestDataAvailabilityOverheadGas,
		DestGasPerDataAvailabilityByte:    fqc.DestGasPerDataAvailabilityByte,
		DestDataAvailabilityMultiplierBps: fqc.DestDataAvailabilityMultiplierBps,
		ChainFamilySelector:               [4]byte(binary.BigEndian.AppendUint32(nil, fqc.ChainFamilySelector)),
		EnforceOutOfOrder:                 fqc.EnforceOutOfOrder,
		DefaultTokenFeeUSDCents:           fqc.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead:       fqc.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:                 fqc.DefaultTxGasLimit,
		GasMultiplierWeiPerEth:            fqc.GasMultiplierWeiPerEth,
		GasPriceStalenessThreshold:        fqc.GasPriceStalenessThreshold,
		NetworkFeeUSDCents:                fqc.NetworkFeeUSDCents,
	}
}

func TranslateTokenPrices(prices map[string]*big.Int) []fee_quoter.InternalTokenPriceUpdate {
	var result []fee_quoter.InternalTokenPriceUpdate
	for k, v := range prices {
		result = append(result, fee_quoter.InternalTokenPriceUpdate{
			SourceToken: common.HexToAddress(k),
			UsdPerToken: v,
		})
	}
	return result
}
