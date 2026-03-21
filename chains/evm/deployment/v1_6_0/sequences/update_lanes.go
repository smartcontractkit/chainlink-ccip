package sequences

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	fqops2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ConfigureLaneLegAsSource = operations.NewSequence(
	"ConfigureLaneLegAsSource",
	semver.MustParse("1.6.0"),
	"Configures lane leg as source on CCIP 1.6.0",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput
		var err error
		b.Logger.Infof("EVM Configuring lane leg as source. src: %+v, dest: %+v", input.Source, input.Dest)

		isFQ2 := input.Source.FeeQuoterVersion != nil && input.Source.FeeQuoterVersion.Compare(fqops2.Version) >= 0
		fqAddr := common.BytesToAddress(input.Source.FeeQuoter)

		evmChain, ok := chains.EVMChains()[input.Source.Selector]
		if !ok {
			return result, fmt.Errorf("chain with selector %d not defined", input.Source.Selector)
		}

		if isFQ2 {
			result, err = configureFeeQuoterV2(b, input, fqAddr, evmChain, result)
			if err != nil {
				return result, err
			}
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
			b.Logger.Info("Destination configs updated on FeeQuoters")

			if input.Dest.GasPrice != nil || len(input.Source.TokenPrices) > 0 {
				skip, err := feeQuoterPricesAlreadySeeded(b, evmChain, fqAddr, input.Source.Selector, input.Dest.Selector)
				if err != nil {
					return result, err
				}
				if skip {
					b.Logger.Info("Skipping FeeQuoter price updates: prices already seeded for dest chain")
				} else {
					result, err = sequences.RunAndMergeSequence(b, chains, FeeQuoterUpdatePricesSequence, FeeQuoterUpdatePricesSequenceInput{
						Address:       fqAddr,
						ChainSelector: input.Source.Selector,
						UpdatesByChain: fqops.PriceUpdates{
							GasPriceUpdates:   translateGasPrice(input.Dest.Selector, input.Dest.GasPrice),
							TokenPriceUpdates: TranslateTokenPrices(input.Source.TokenPrices),
						},
					}, result)
					if err != nil {
						return result, err
					}
					b.Logger.Info("FeeQuoter prices seeded")
				}
			}
		}

		result, err = sequences.RunAndMergeSequence(b, chains, OnRampApplyDestChainConfigUpdatesSequence, OnRampApplyDestChainConfigUpdatesSequenceInput{
			Address:       common.BytesToAddress(input.Source.OnRamp),
			ChainSelector: input.Source.Selector,
			UpdatesByChain: []onrampops.DestChainConfigArgs{
				{
					Router:            common.BytesToAddress(input.Source.Router),
					DestChainSelector: input.Dest.Selector,
					AllowlistEnabled:  input.Source.AllowListEnabled,
				},
			},
		}, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("OnRamp destination chain configs updated")

		if err := applyOnRampAllowlist(b, evmChain, input, &result); err != nil {
			return result, err
		}

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

// feeQuoterPricesAlreadySeeded checks whether a gas price entry already exists
// for the dest chain. If so, prices have been seeded and should not be overwritten.
func feeQuoterPricesAlreadySeeded(
	b operations.Bundle,
	chain cldf_evm.Chain,
	feeQuoter common.Address,
	sourceChainSelector uint64,
	destChainSelector uint64,
) (bool, error) {
	rep, err := operations.ExecuteOperation(b, fqops.GetDestinationChainGasPrice, chain, contract.FunctionInput[uint64]{
		ChainSelector: sourceChainSelector,
		Address:       feeQuoter,
		Args:          destChainSelector,
	})
	if err != nil {
		return false, fmt.Errorf("read destination gas price from fee quoter: %w", err)
	}
	return rep.Output.Timestamp > 0, nil
}

// feeQuoterV2PricesAlreadySeeded is the FQ 2.0 equivalent of feeQuoterPricesAlreadySeeded.
func feeQuoterV2PricesAlreadySeeded(
	b operations.Bundle,
	chain cldf_evm.Chain,
	feeQuoter common.Address,
	sourceChainSelector uint64,
	destChainSelector uint64,
) (bool, error) {
	rep, err := operations.ExecuteOperation(b, fqops2.GetDestinationChainGasPrice, chain, contract.FunctionInput[uint64]{
		ChainSelector: sourceChainSelector,
		Address:       feeQuoter,
		Args:          destChainSelector,
	})
	if err != nil {
		return false, fmt.Errorf("read destination gas price from fee quoter v2: %w", err)
	}
	return rep.Output.Timestamp > 0, nil
}

func translateGasPrice(destChainSelector uint64, gasPrice *big.Int) []fqops.GasPriceUpdate {
	if gasPrice == nil {
		return nil
	}
	return []fqops.GasPriceUpdate{{DestChainSelector: destChainSelector, UsdPerUnitGas: gasPrice}}
}

// applyOnRampAllowlist writes the full allowlist from the input to the OnRamp.
// The input is treated as the complete desired state.
func applyOnRampAllowlist(b operations.Bundle, chain cldf_evm.Chain, input lanes.UpdateLanesInput, result *sequences.OnChainOutput) error {
	if len(input.Source.AllowList) > 0 && !input.Source.AllowListEnabled {
		return errors.New("OnRamp allowlist: cannot specify AllowList addresses when AllowListEnabled is false")
	}
	desired := make([]common.Address, 0, len(input.Source.AllowList))
	for _, h := range input.Source.AllowList {
		if !common.IsHexAddress(h) {
			return fmt.Errorf("allowlist address %q is not a valid hex address", h)
		}
		desired = append(desired, common.HexToAddress(h))
	}
	writeRep, err := operations.ExecuteOperation(b, onrampops.ApplyAllowlistUpdates, chain, contract.FunctionInput[[]onrampops.AllowlistConfigArgs]{
		ChainSelector: input.Source.Selector,
		Address:       common.BytesToAddress(input.Source.OnRamp),
		Args: []onrampops.AllowlistConfigArgs{{
			DestChainSelector:       input.Dest.Selector,
			AllowlistEnabled:        input.Source.AllowListEnabled,
			AddedAllowlistedSenders: desired,
		}},
	})
	if err != nil {
		return fmt.Errorf("apply onramp allowlist updates: %w", err)
	}
	batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{writeRep.Output})
	if err != nil {
		return err
	}
	result.BatchOps = append(result.BatchOps, batch)
	b.Logger.Info("OnRamp allowlist applied")
	return nil
}

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

func configureFeeQuoterV2(b operations.Bundle, input lanes.UpdateLanesInput, fqAddr common.Address, evmChain cldf_evm.Chain, result sequences.OnChainOutput) (sequences.OnChainOutput, error) {
	destChainCfgReport, err := operations.ExecuteOperation(
		b, fqops2.ApplyDestChainConfigUpdates, evmChain,
		contract.FunctionInput[[]fqops2.DestChainConfigArgs]{
			ChainSelector: evmChain.Selector,
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

	batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{destChainCfgReport.Output})
	if err != nil {
		return result, err
	}
	result.BatchOps = append(result.BatchOps, batch)

	if input.Dest.GasPrice != nil || len(input.Source.TokenPrices) > 0 {
		skip, err := feeQuoterV2PricesAlreadySeeded(b, evmChain, fqAddr, input.Source.Selector, input.Dest.Selector)
		if err != nil {
			return result, err
		}
		if skip {
			b.Logger.Info("Skipping FeeQuoter v2 price updates: prices already seeded for dest chain")
		} else {
			priceReport, err := operations.ExecuteOperation(b, fqops2.UpdatePrices, evmChain, contract.FunctionInput[fqops2.PriceUpdates]{
				ChainSelector: evmChain.Selector,
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
			priceBatch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{priceReport.Output})
			if err != nil {
				return result, err
			}
			result.BatchOps = append(result.BatchOps, priceBatch)
			b.Logger.Info("FeeQuoter v2 prices seeded")
		}
	}

	return result, nil
}
