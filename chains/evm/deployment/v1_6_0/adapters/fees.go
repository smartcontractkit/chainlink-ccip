package adapters

import (
	"fmt"
	"math"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	fqsequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ fees.FeeAdapter = (*FeesAdapter)(nil)

type FeesAdapter struct{}

func NewFeesAdapter() *FeesAdapter {
	return &FeesAdapter{}
}

func (a *FeesAdapter) getFeeQuoterAddress(ds datastore.DataStore, src uint64) (common.Address, error) {
	fqAddrRefDS := datastore.AddressRef{
		Type:    datastore.ContractType(fqops.ContractType),
		Version: fqops.Version,
	}

	fqAddr, err := datastore_utils.FindAndFormatRef(ds, fqAddrRefDS, src, datastore_utils_evm.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
	}

	return fqAddr, nil
}

func (a *FeesAdapter) GetTokenTransferFeeConfigDefaults(src uint64, dst uint64) fees.TokenTransferFeeArgs {
	minFeeUSDCents := uint32(25)

	// NOTE: we validate that src != dst so only one of these if statements will execute
	if src == chain_selectors.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 50
	}
	if dst == chain_selectors.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 150
	}

	return fees.TokenTransferFeeArgs{
		DestBytesOverhead: 32,
		DestGasOverhead:   90_000,
		MinFeeUSDCents:    minFeeUSDCents,
		MaxFeeUSDCents:    math.MaxUint32,
		DeciBps:           0,
		IsEnabled:         true,
	}
}

func (a *FeesAdapter) GetTokenTransferFeeConfig(e cldf.Environment, src uint64, dst uint64, address string) (fees.TokenTransferFeeArgs, error) {
	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("chain with selector %d not defined", src)
	}

	fqAddr, err := a.getFeeQuoterAddress(e.DataStore, src)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
	}

	fq, err := fee_quoter.NewFeeQuoter(fqAddr, chain.Client)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to instantiate FeeQuoter at %s: %w", fqAddr.Hex(), err)
	}

	if !common.IsHexAddress(address) {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("invalid contract address: %s", address)
	}

	// This gets the token transfer fee config for the given token from the FeeQuoter contract
	// https://etherscan.io/address/0x40858070814a57FdF33a613ae84fE0a8b4a874f7#code#F1#L819
	cfg, err := fq.GetTokenTransferFeeConfig(&bind.CallOpts{Context: e.GetContext()}, dst, common.HexToAddress(address))
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get token transfer fee config from FeeQuoter at %s: %w", fqAddr.Hex(), err)
	}

	return fees.TokenTransferFeeArgs{
		DestBytesOverhead: cfg.DestBytesOverhead,
		DestGasOverhead:   cfg.DestGasOverhead,
		MinFeeUSDCents:    cfg.MinFeeUSDCents,
		MaxFeeUSDCents:    cfg.MaxFeeUSDCents,
		IsEnabled:         cfg.IsEnabled,
		DeciBps:           cfg.DeciBps,
	}, nil
}

func (a *FeesAdapter) SetTokenTransferFeeConfig(ds datastore.DataStore, src uint64) *operations.Sequence[fees.SetTokenTransferFeeConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetTokenTransferFee",
		semver.MustParse("1.6.0"),
		"Sets token transfer fee configuration on CCIP 1.6.0 FeeQuoter contracts",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetTokenTransferFeeConfigSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput

			fqAddr, err := a.getFeeQuoterAddress(ds, src)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
			}

			updatesByChain := fqops.ApplyTokenTransferFeeConfigUpdatesInput{}
			for dstSel, dstCfg := range input.Settings {
				tokensToUseDefaultFeeConfigs := []fee_quoter.FeeQuoterTokenTransferFeeConfigRemoveArgs{}
				tokenTransferFeeConfigs := []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{}
				for rawTokenAddress, feeCfg := range dstCfg {
					if !common.IsHexAddress(rawTokenAddress) {
						return sequences.OnChainOutput{}, fmt.Errorf("invalid token address: %s", rawTokenAddress)
					}

					token := common.HexToAddress(rawTokenAddress)
					if feeCfg == nil {
						tokensToUseDefaultFeeConfigs = append(
							tokensToUseDefaultFeeConfigs,
							fee_quoter.FeeQuoterTokenTransferFeeConfigRemoveArgs{
								DestChainSelector: dstSel,
								Token:             token,
							},
						)
					} else {
						tokenTransferFeeConfigs = append(
							tokenTransferFeeConfigs,
							fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
								Token: token,
								TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
									DestBytesOverhead: feeCfg.DestBytesOverhead,
									DestGasOverhead:   feeCfg.DestGasOverhead,
									MinFeeUSDCents:    feeCfg.MinFeeUSDCents,
									MaxFeeUSDCents:    feeCfg.MaxFeeUSDCents,
									IsEnabled:         feeCfg.IsEnabled,
									DeciBps:           feeCfg.DeciBps,
								},
							},
						)
					}
				}

				if len(tokensToUseDefaultFeeConfigs) > 0 {
					updatesByChain.TokensToUseDefaultFeeConfigs = append(updatesByChain.TokensToUseDefaultFeeConfigs, tokensToUseDefaultFeeConfigs...)
				}

				if len(tokenTransferFeeConfigs) > 0 {
					updatesByChain.TokenTransferFeeConfigArgs = append(updatesByChain.TokenTransferFeeConfigArgs, fee_quoter.FeeQuoterTokenTransferFeeConfigArgs{
						TokenTransferFeeConfigs: tokenTransferFeeConfigs,
						DestChainSelector:       dstSel,
					})
				}
			}

			if len(updatesByChain.TokensToUseDefaultFeeConfigs) == 0 && len(updatesByChain.TokenTransferFeeConfigArgs) == 0 {
				return result, nil
			}

			result, err = sequences.RunAndMergeSequence(b, chains,
				fqsequences.FeeQuoterApplyTokenTransferFeeConfigUpdatesSequence,
				fqsequences.FeeQuoterApplyTokenTransferFeeConfigUpdatesSequenceInput{
					UpdatesByChain: updatesByChain,
					ChainSelector:  input.Selector,
					Address:        fqAddr,
				},
				result,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterApplyTokenTransferFeeConfigUpdates operation: %w", err)
			}

			return result, nil
		},
	)
}
