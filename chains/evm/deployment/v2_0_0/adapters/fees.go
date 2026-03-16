package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	evmseq16 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ fees.FeeAdapter = (*FeesAdapter)(nil)

type FeesAdapter struct {
	evm *evmseq16.EVMAdapter
}

func NewFeesAdapter(evmAdapter *evmseq16.EVMAdapter) *FeesAdapter {
	return &FeesAdapter{
		evm: evmAdapter,
	}
}

func (a *FeesAdapter) GetFeeContractRef(e cldf.Environment, src uint64, dst uint64) (datastore.AddressRef, error) {
	ds := e.DataStore
	fqAddr, err := a.evm.GetFQAddress(ds, src)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
	}

	filter := datastore.AddressRef{
		Type:          datastore.ContractType(fqops.ContractType),
		Address:       common.BytesToAddress(fqAddr).Hex(),
		ChainSelector: src,
	}

	feeContractRef, err := datastore_utils.FindAndFormatRef(
		ds,
		filter,
		src,
		datastore_utils.FullRef,
	)

	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to find FeeQuoter address ref for chain selector %d: %w", src, err)

	}

	return feeContractRef, nil
}

func (a *FeesAdapter) GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) fees.TokenTransferFeeArgs {
	return fees.GetDefaultChainAgnosticTokenTransferFeeConfig(src, dst)
}

func (a *FeesAdapter) GetOnchainTokenTransferFeeConfig(e cldf.Environment, src uint64, dst uint64, address string) (fees.TokenTransferFeeArgs, error) {
	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("chain with selector %d not defined", src)
	}

	fqRef, err := a.GetFeeContractRef(e, src, dst)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
	}

	fqAddr := common.HexToAddress(fqRef.Address)
	fq, err := fqops.NewFeeQuoterContract(fqAddr, chain.Client)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to instantiate FeeQuoter contract at address %s on chain selector %d: %w", fqAddr.Hex(), src, err)
	}

	if !common.IsHexAddress(address) {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("invalid token address: %s", address)
	}

	// This gets the token transfer fee config for the given token from the FeeQuoter contract
	// https://etherscan.io/address/0x40858070814a57FdF33a613ae84fE0a8b4a874f7#code#F1#L819 --TODO: update comment with link to correct line in 2.0.0 version of FeeQuoter contract once available
	cfg, err := fq.GetTokenTransferFeeConfig(&bind.CallOpts{Context: e.GetContext()}, dst, common.HexToAddress(address))
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get token transfer fee config from FeeQuoter at %s for src %d, dst %d, token %s: %w", fqAddr.Hex(), src, dst, address, err)
	}

	e.Logger.Infof("Fetched on-chain token transfer fee config for src %d, dst %d, token %s: %+v", src, dst, address, cfg)
	return fees.TokenTransferFeeArgs{
		DestBytesOverhead: cfg.DestBytesOverhead,
		DestGasOverhead:   cfg.DestGasOverhead,
		IsEnabled:         cfg.IsEnabled,
		MaxFeeUSDCents:    0,               // Max fee is not defined in 2.0.0 version of FeeQuoter contract, so we set it to 0
		MinFeeUSDCents:    cfg.FeeUSDCents, // In 2.0.0 version of FeeQuoter contract, there is only a single fee parameter (FeeUSDCents) https://github.com/smartcontractkit/chainlink-ccip/blob/73fcb2020b9335c965a7d2bb5d932c0fa05c7948/chains/evm/contracts/FeeQuoter.sol#L97
	}, nil
}

func (a *FeesAdapter) SetTokenTransferFee(e cldf.Environment) *operations.Sequence[fees.SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetTokenTransferFee",
		semver.MustParse("2.0.0"),
		"Sets token transfer fee configuration on CCIP 2.0.0 FeeQuoter contracts",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetTokenTransferFeeSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			src := input.Selector

			updatesByChain := fqops.ApplyTokenTransferFeeConfigUpdatesArgs{}

			fqRef, errFq := a.GetFeeContractRef(e, src, 0) // dst is not needed to get the fee quoter address in 1.6.0+
			if errFq != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, errFq)
			}

			for dst, dstCfg := range input.Settings {

				var tokensToUseDefaultFeeConfigs []fqops.TokenTransferFeeConfigRemoveArgs
				var tokenTransferFeeConfigs []fqops.TokenTransferFeeConfigSingleTokenArgs
				for rawTokenAddress, feeCfg := range dstCfg {
					if !common.IsHexAddress(rawTokenAddress) {
						return sequences.OnChainOutput{}, fmt.Errorf("invalid token address for src %d and dst %d: %s", src, dst, rawTokenAddress)
					}

					token := common.HexToAddress(rawTokenAddress)
					if feeCfg == nil {
						tokensToUseDefaultFeeConfigs = append(
							tokensToUseDefaultFeeConfigs,
							fqops.TokenTransferFeeConfigRemoveArgs{
								DestChainSelector: dst,
								Token:             token,
							},
						)
					} else {
						tokenTransferFeeConfigs = append(
							tokenTransferFeeConfigs,
							fqops.TokenTransferFeeConfigSingleTokenArgs{
								Token: token,
								TokenTransferFeeConfig: fqops.TokenTransferFeeConfig{
									DestBytesOverhead: feeCfg.DestBytesOverhead,
									DestGasOverhead:   feeCfg.DestGasOverhead,
									FeeUSDCents:       feeCfg.MinFeeUSDCents,
									IsEnabled:         feeCfg.IsEnabled,
								},
							},
						)
					}
				}

				if len(tokensToUseDefaultFeeConfigs) > 0 {
					updatesByChain.TokensToUseDefaultFeeConfigs = append(updatesByChain.TokensToUseDefaultFeeConfigs, tokensToUseDefaultFeeConfigs...)
				}

				if len(tokenTransferFeeConfigs) > 0 {
					updatesByChain.TokenTransferFeeConfigArgs = append(updatesByChain.TokenTransferFeeConfigArgs, fqops.TokenTransferFeeConfigArgs{
						TokenTransferFeeConfigs: tokenTransferFeeConfigs,
						DestChainSelector:       dst,
					})
				}
			}

			if len(updatesByChain.TokensToUseDefaultFeeConfigs) == 0 && len(updatesByChain.TokenTransferFeeConfigArgs) == 0 {
				return result, nil
			}

			result, err := sequences.RunAndMergeSequence(b, chains,
				evmseq.SequenceFeeQuoterUpdate,
				evmseq.FeeQuoterUpdate{
					ChainSelector:                 input.Selector,
					ExistingAddresses:             []datastore.AddressRef{fqRef},
					TokenTransferFeeConfigUpdates: updatesByChain,
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
