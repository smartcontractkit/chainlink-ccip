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
	fqops2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
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
		var err error
		b.Logger.Infof("EVM Configuring lane leg as source. src: %+v, dest: %+v", input.Source, input.Dest)

		isFQ2 := input.Source.FeeQuoterVersion != nil && input.Source.FeeQuoterVersion.Compare(fqops2.Version) >= 0
		fqAddr := common.BytesToAddress(input.Source.FeeQuoter)

		if isFQ2 {
			var err error
			result, err = configureFeeQuoterV2(b, chains, input, fqAddr, result)
			if err != nil {
				return result, err
			}

		} else {
			result, err = sequences.RunAndMergeSequence(b, chains, FeeQuoterApplyDestChainConfigUpdatesSequence, FeeQuoterApplyDestChainConfigUpdatesSequenceInput{
				Address:       common.BytesToAddress(input.Source.FeeQuoter),
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
			b.Logger.Info("Destination configs updated on FeeQuoters")

			result, err = sequences.RunAndMergeSequence(b, chains, FeeQuoterUpdatePricesSequence, FeeQuoterUpdatePricesSequenceInput{
				Address:       common.BytesToAddress(input.Source.FeeQuoter),
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
			b.Logger.Info("Gas prices updated on FeeQuoters")
		}

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

func TranslateFQtoV2(fqc lanes.FeeQuoterDestChainConfig) fqops2.DestChainConfig {
	var v2 lanes.FeeQuoterV2Params
	if fqc.V2Params != nil {
		v2 = *fqc.V2Params
	}
	return fqops2.DestChainConfig{
		IsEnabled:                   fqc.IsEnabled,
		MaxDataBytes:                fqc.MaxDataBytes,
		MaxPerMsgGasLimit:           fqc.MaxPerMsgGasLimit,
		DestGasOverhead:             fqc.DestGasOverhead,
		DestGasPerPayloadByteBase:   fqc.DestGasPerPayloadByteBase,
		ChainFamilySelector:         [4]byte(binary.BigEndian.AppendUint32(nil, fqc.ChainFamilySelector)),
		DefaultTokenFeeUSDCents:     fqc.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead: fqc.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:           fqc.DefaultTxGasLimit,
		NetworkFeeUSDCents:          fqc.NetworkFeeUSDCents,
		LinkFeeMultiplierPercent:    v2.LinkFeeMultiplierPercent,
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

func TranslateTokenPricesV2(prices map[string]*big.Int) []fqops2.TokenPriceUpdate {
	var result []fqops2.TokenPriceUpdate
	for k, v := range prices {
		result = append(result, fqops2.TokenPriceUpdate{
			SourceToken: common.HexToAddress(k),
			UsdPerToken: v,
		})
	}
	return result
}

func configureFeeQuoterV2(b operations.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput, fqAddr common.Address, result sequences.OnChainOutput) (sequences.OnChainOutput, error) {
	chain, ok := chains.EVMChains()[input.Source.Selector]
	if !ok {
		return result, fmt.Errorf("chain with selector %d not defined", input.Source.Selector)
	}
	// FeeQuoter 2.0: apply dest chain config then update prices
	destChainCfgReport, err := operations.ExecuteOperation(
		b, fqops2.ApplyDestChainConfigUpdates, chain,
		contract.FunctionInput[[]fqops2.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       fqAddr,
			Args: []fqops2.DestChainConfigArgs{
				{
					DestChainSelector: input.Dest.Selector,
					DestChainConfig:   TranslateFQtoV2(input.Dest.FeeQuoterDestChainConfig),
				},
			},
		})
	if err != nil {
		return result, err
	}
	b.Logger.Info("Destination configs updated on FeeQuoters")
	// FeeQuoter 2.0: Update prices
	priceReport, err := operations.ExecuteOperation(b, fqops2.UpdatePrices, chain, contract.FunctionInput[fqops2.PriceUpdates]{
		ChainSelector: chain.Selector,
		Address:       fqAddr,
		Args: fqops2.PriceUpdates{
			TokenPriceUpdates: TranslateTokenPricesV2(input.Source.TokenPrices),
			GasPriceUpdates: []fqops2.GasPriceUpdate{
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
	b.Logger.Info("Gas prices updated on FeeQuoters")

	batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{destChainCfgReport.Output, priceReport.Output})
	if err != nil {
		return result, err
	}
	result.BatchOps = append(result.BatchOps, batch)
	return result, nil
}
