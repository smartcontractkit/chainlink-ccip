package sequences

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	fq2ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var feeQuoterV2 = semver.MustParse("2.0.0")

var ConfigureLaneLegAsSource = operations.NewSequence(
	"ConfigureLaneLegAsSource",
	semver.MustParse("1.0.0"),
	"Configures lane leg as source on CCIP 1.6.0",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput
		b.Logger.Infof("EVM Configuring lane leg as source. src: %+v, dest: %+v", input.Source, input.Dest)

		fqAddr := common.BytesToAddress(input.Source.FeeQuoter)
		useFQV2 := input.Source.FeeQuoterVersion != nil && !input.Source.FeeQuoterVersion.LessThan(feeQuoterV2)

		var err error
		if useFQV2 {
			chain, ok := chains.EVMChains()[input.Source.Selector]
			if !ok {
				return result, fmt.Errorf("chain with selector %d not defined", input.Source.Selector)
			}
			// FeeQuoter 2.0: apply dest chain config then update prices
			destCfgReport, err := operations.ExecuteOperation(b, fq2ops.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]fq2ops.DestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       fqAddr,
				Args: []fq2ops.DestChainConfigArgs{
					{
						DestChainSelector: input.Dest.Selector,
						DestChainConfig:   TranslateFQToV2(input.Dest.FeeQuoterDestChainConfig),
					},
				},
			})
			if err != nil {
				return result, err
			}
			priceReport, err := operations.ExecuteOperation(b, fq2ops.UpdatePrices, chain, contract.FunctionInput[fq2ops.PriceUpdates]{
				ChainSelector: chain.Selector,
				Address:       fqAddr,
				Args: fq2ops.PriceUpdates{
					TokenPriceUpdates: TranslateTokenPricesToV2(input.Source.TokenPrices),
					GasPriceUpdates: []fq2ops.GasPriceUpdate{
						{
							DestChainSelector: input.Dest.Selector,
							UsdPerUnitGas:     input.Dest.GasPrice,
						},
					},
				},
			})
			if err != nil {
				return result, err
			}
			batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{destCfgReport.Output, priceReport.Output})
			if err != nil {
				return result, err
			}
			result.BatchOps = append(result.BatchOps, batch)
		} else {
			result, err = sequences.RunAndMergeSequence(b, chains, FeeQuoterApplyDestChainConfigUpdatesSequence, FeeQuoterApplyDestChainConfigUpdatesSequenceInput{
				Address:       fqAddr,
				ChainSelector: input.Source.Selector,
				UpdatesByChain: []fqops.DestChainConfigArgs{
					{
						DestChainSelector: input.Dest.Selector,
						DestChainConfig:   TranslateFQ(input.Dest.FeeQuoterDestChainConfig),
					},
				},
			}, result)
			if err != nil {
				return result, err
			}

			result, err = sequences.RunAndMergeSequence(b, chains, FeeQuoterUpdatePricesSequence, FeeQuoterUpdatePricesSequenceInput{
				Address:       fqAddr,
				ChainSelector: input.Source.Selector,
				UpdatesByChain: fqops.PriceUpdates{
					TokenPriceUpdates: TranslateTokenPrices(input.Source.TokenPrices),
					GasPriceUpdates: []fqops.GasPriceUpdate{
						{
							DestChainSelector: input.Dest.Selector,
							UsdPerUnitGas:     input.Dest.GasPrice,
						},
					},
				},
			}, result)
			if err != nil {
				return result, err
			}
		}
		b.Logger.Info("Destination configs updated on FeeQuoters")
		b.Logger.Info("Gas prices updated on FeeQuoters")

	result, err = sequences.RunAndMergeSequence(b, chains, OnRampApplyDestChainConfigUpdatesSequence, OnRampApplyDestChainConfigUpdatesSequenceInput{
		Address:       common.BytesToAddress(input.Source.OnRamp),
		ChainSelector: input.Source.Selector,
		UpdatesByChain: []onrampops.DestChainConfigArgs{
			{
				Router:            common.BytesToAddress(input.Source.Router),
				DestChainSelector: input.Dest.Selector,
				AllowlistEnabled:  input.Dest.AllowListEnabled,
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
			Address:       common.BytesToAddress(input.Source.Router),
			ChainSelector: input.Source.Selector,
			UpdatesByChain: router.ApplyRampsUpdatesArgs{
				OnRampUpdates: []router.OnRamp{onrampUpdate},
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
		b.Logger.Infof("EVM Configuring lane leg as destination. src: %+v, dest: %+v", input.Source, input.Dest)

	result, err := sequences.RunAndMergeSequence(b, chains, OffRampApplySourceChainConfigUpdatesSequence, OffRampApplySourceChainConfigUpdatesSequenceInput{
		Address:       common.BytesToAddress(input.Dest.OffRamp),
		ChainSelector: input.Dest.Selector,
		UpdatesByChain: []offrampops.SourceChainConfigArgs{
			{
				Router:              common.BytesToAddress(input.Dest.Router),
				SourceChainSelector: input.Source.Selector,
				// https://github.com/smartcontractkit/chainlink/blob/f7ca3d51db51258bb3b8ae22a8e1593d03bc040b/deployment/ccip/changeset/v1_6/cs_chain_contracts.go#L1148
				OnRamp:                    common.LeftPadBytes(input.Source.OnRamp, 32),
				IsEnabled:                 !input.IsDisabled,
					IsRMNVerificationDisabled: !input.Source.RMNVerificationEnabled,
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
			Address:       common.BytesToAddress(input.Dest.Router),
			ChainSelector: input.Dest.Selector,
			UpdatesByChain: router.ApplyRampsUpdatesArgs{
				OffRampAdds:    offRampAdds,
				OffRampRemoves: offRampRemoves,
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

func TranslateFQ(fqc lanes.FeeQuoterDestChainConfig) fqops.DestChainConfig {
	var v1 lanes.FeeQuoterV1Params
	if fqc.V1Params != nil {
		v1 = *fqc.V1Params
	}
	return fqops.DestChainConfig{
		IsEnabled:                         fqc.IsEnabled,
		MaxNumberOfTokensPerMsg:           v1.MaxNumberOfTokensPerMsg,
		MaxDataBytes:                      fqc.MaxDataBytes,
		MaxPerMsgGasLimit:                 fqc.MaxPerMsgGasLimit,
		DestGasOverhead:                   fqc.DestGasOverhead,
		DestGasPerPayloadByteBase:         fqc.DestGasPerPayloadByteBase,
		DestGasPerPayloadByteHigh:         v1.DestGasPerPayloadByteHigh,
		DestGasPerPayloadByteThreshold:    v1.DestGasPerPayloadByteThreshold,
		DestDataAvailabilityOverheadGas:   v1.DestDataAvailabilityOverheadGas,
		DestGasPerDataAvailabilityByte:    v1.DestGasPerDataAvailabilityByte,
		DestDataAvailabilityMultiplierBps: v1.DestDataAvailabilityMultiplierBps,
		ChainFamilySelector:               [4]byte(binary.BigEndian.AppendUint32(nil, fqc.ChainFamilySelector)),
		EnforceOutOfOrder:                 v1.EnforceOutOfOrder,
		DefaultTokenFeeUSDCents:           fqc.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead:       fqc.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:                 fqc.DefaultTxGasLimit,
		GasMultiplierWeiPerEth:            v1.GasMultiplierWeiPerEth,
		GasPriceStalenessThreshold:        v1.GasPriceStalenessThreshold,
		NetworkFeeUSDCents:                uint32(fqc.NetworkFeeUSDCents),
	}
}

func TranslateTokenPrices(prices map[string]*big.Int) []fqops.TokenPriceUpdate {
	var result []fqops.TokenPriceUpdate
	for k, v := range prices {
		result = append(result, fqops.TokenPriceUpdate{
			SourceToken: common.HexToAddress(k),
			UsdPerToken: v,
		})
	}
	return result
}

// TranslateFQToV2 maps lanes.FeeQuoterDestChainConfig to FeeQuoter 2.0 DestChainConfig.
// v2.0 has a smaller struct (no MaxNumberOfTokensPerMsg, payload byte high/threshold, data availability, etc.).
func TranslateFQToV2(fqc lanes.FeeQuoterDestChainConfig) fq2ops.DestChainConfig {
	networkFeeUSDCents := uint16(fqc.NetworkFeeUSDCents)
	if fqc.NetworkFeeUSDCents > 0xffff {
		networkFeeUSDCents = 0xffff
	}
	return fq2ops.DestChainConfig{
		IsEnabled:                   fqc.IsEnabled,
		MaxDataBytes:                fqc.MaxDataBytes,
		MaxPerMsgGasLimit:           fqc.MaxPerMsgGasLimit,
		DestGasOverhead:             fqc.DestGasOverhead,
		DestGasPerPayloadByteBase:   fqc.DestGasPerPayloadByteBase,
		ChainFamilySelector:         [4]byte(binary.BigEndian.AppendUint32(nil, fqc.ChainFamilySelector)),
		DefaultTokenFeeUSDCents:     fqc.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead: fqc.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:           fqc.DefaultTxGasLimit,
		NetworkFeeUSDCents:          networkFeeUSDCents,
		LinkFeeMultiplierPercent:    100, // 100% = no extra LINK multiplier
	}
}

func TranslateTokenPricesToV2(prices map[string]*big.Int) []fq2ops.TokenPriceUpdate {
	var result []fq2ops.TokenPriceUpdate
	for k, v := range prices {
		result = append(result, fq2ops.TokenPriceUpdate{
			SourceToken: common.HexToAddress(k),
			UsdPerToken: v,
		})
	}
	return result
}
