package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	v1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ fees.FeeAdapter = (*FeesAdapter)(nil)

type FeesAdapter struct {
	resolver v1_0_0.EVMFeeResolver
	evm      *evmseqV1_6_0.EVMAdapter
}

func NewFeesAdapter(evmAdapter *evmseqV1_6_0.EVMAdapter) *FeesAdapter {
	return &FeesAdapter{
		resolver: v1_0_0.EVMFeeResolver{},
		evm:      evmAdapter,
	}
}

func (a *FeesAdapter) validateFeeRef(feeRef datastore.AddressRef) error {
	if feeRef.Type.String() != fqops.ContractType.String() {
		return fmt.Errorf("unexpected contract type for FeeQuoter address ref: got %s, want %s", feeRef.Type.String(), fqops.ContractType)
	}
	if !feeRef.Version.Equal(utils.Version_2_0_0) {
		return fmt.Errorf("unexpected FeeQuoter contract version: got %s, want %s", feeRef.Version, utils.Version_2_0_0)
	}

	return nil
}

func (a *FeesAdapter) GetFeeContractRef(e cldf.Environment, onRampRef datastore.AddressRef, src uint64, dst uint64) (datastore.AddressRef, error) {
	if onRampRef.Type.String() != onramp.ContractType.String() {
		return datastore.AddressRef{}, fmt.Errorf("unexpected contract type for OnRamp address ref for src %d and dst %d: got %s, want %s", src, dst, onRampRef.Type.String(), onramp.ContractType)
	}
	if !onRampRef.Version.Equal(utils.Version_2_0_0) {
		return datastore.AddressRef{}, fmt.Errorf("unexpected OnRamp contract version for src %d and dst %d: got %s, want %s", src, dst, onRampRef.Version, utils.Version_2_0_0)
	}

	onRampAddr, err := datastore_utils_evm.ToEVMAddress(onRampRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to convert OnRamp address ref to EVM address for src %d and dst %d: %w", src, dst, err)
	}

	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("chain with selector %d not defined", src)
	}

	report, err := operations.ExecuteOperation(
		e.OperationsBundle,
		onramp.GetDynamicConfig,
		chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: src,
			Address:       onRampAddr,
			Args:          struct{}{},
		},
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to execute GetDynamicConfig operation for OnRamp at %s on chain selector %d with dst %d: %w", onRampAddr.Hex(), src, dst, err)
	}

	// NOTE: version is omitted intentionally here - we could have a v2.0 OnRamp connected to a v1.6 FQ for example
	fqRef, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			ChainSelector: src,
			Address:       report.Output.FeeQuoter.Hex(),
			Type:          datastore.ContractType(fqops.ContractType),
		},
		src,
		datastore_utils.FullRef,
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to find FeeQuoter address ref in datastore for address %s on chain selector %d: %w", report.Output.FeeQuoter.Hex(), src, err)
	}

	return fqRef, nil
}

func (a *FeesAdapter) GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) fees.TokenTransferFeeArgs {
	return fees.GetDefaultChainAgnosticTokenTransferFeeConfig(src, dst)
}

func (a *FeesAdapter) GetOnchainTokenTransferFeeConfig(e cldf.Environment, feeRef datastore.AddressRef, src uint64, dst uint64, address string) (fees.TokenTransferFeeArgs, error) {
	err := a.validateFeeRef(feeRef)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("invalid FeeQuoter address ref for src %d and dst %d: %w", src, dst, err)
	}

	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("chain with selector %d not defined", src)
	}
	fqAddr, err := datastore_utils_evm.ToEVMAddress(feeRef)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to convert FeeQuoter address ref to EVM address for src %d and dst %d: %w", src, dst, err)
	}
	fq, err := fqops.NewFeeQuoterContract(fqAddr, chain.Client)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to instantiate FeeQuoter contract at address %s on chain selector %d: %w", fqAddr.Hex(), src, err)
	}
	if !common.IsHexAddress(address) {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("invalid token address: %s", address)
	}

	// This gets the token transfer fee config for the given token from the FeeQuoter contract
	// https://testnet-explorer.plume.org/address/0x66bc24445e94FF302710E66Ce127E3174F723BD4?tab=contract
	cfg, err := fq.GetTokenTransferFeeConfig(&bind.CallOpts{Context: e.GetContext()}, dst, common.HexToAddress(address))
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get token transfer fee config from FeeQuoter at %s for src %d, dst %d, token %s: %w", fqAddr.Hex(), src, dst, address, err)
	}

	// Max fee is not defined in 2.0.0 version of FeeQuoter contract, so we set it to 0
	// In 2.0.0 version of FeeQuoter contract, there is only a single fee parameter (FeeUSDCents): https://github.com/smartcontractkit/chainlink-ccip/blob/73fcb2020b9335c965a7d2bb5d932c0fa05c7948/chains/evm/contracts/FeeQuoter.sol#L97
	e.Logger.Infof("Fetched on-chain token transfer fee config for src %d, dst %d, token %s: %+v", src, dst, address, cfg)
	return fees.TokenTransferFeeArgs{
		DestBytesOverhead: cfg.DestBytesOverhead,
		DestGasOverhead:   cfg.DestGasOverhead,
		IsEnabled:         cfg.IsEnabled,
		MaxFeeUSDCents:    0,
		MinFeeUSDCents:    cfg.FeeUSDCents,
	}, nil
}

func (a *FeesAdapter) SetTokenTransferFee(e cldf.Environment, feeRef datastore.AddressRef) *operations.Sequence[fees.SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetTokenTransferFee",
		utils.Version_2_0_0,
		"Sets token transfer fee configuration on CCIP 2.0.0 FeeQuoter contracts",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetTokenTransferFeeSequenceInput) (sequences.OnChainOutput, error) {
			src := input.Selector

			err := a.validateFeeRef(feeRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid FeeQuoter address ref: %w", err)
			}

			args := map[uint64]fqops.ApplyTokenTransferFeeConfigUpdatesArgs{}
			for dst, dstCfg := range input.Settings {
				val := args[src]
				for rawTokenAddress, feeCfg := range dstCfg {
					var token common.Address
					if !common.IsHexAddress(rawTokenAddress) {
						return sequences.OnChainOutput{}, fmt.Errorf("invalid token address for src %d and dst %d: %s", src, dst, rawTokenAddress)
					} else {
						token = common.HexToAddress(rawTokenAddress)
					}

					if feeCfg == nil || !feeCfg.IsEnabled {
						val.TokensToUseDefaultFeeConfigs = append(
							val.TokensToUseDefaultFeeConfigs,
							fqops.TokenTransferFeeConfigRemoveArgs{
								DestChainSelector: dst,
								Token:             token,
							},
						)
					} else {
						val.TokenTransferFeeConfigArgs = append(
							val.TokenTransferFeeConfigArgs,
							fqops.TokenTransferFeeConfigArgs{
								DestChainSelector: dst,
								TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{
									{
										Token: token,
										TokenTransferFeeConfig: fqops.TokenTransferFeeConfig{
											DestBytesOverhead: feeCfg.DestBytesOverhead,
											DestGasOverhead:   feeCfg.DestGasOverhead,
											FeeUSDCents:       feeCfg.MinFeeUSDCents,
											IsEnabled:         feeCfg.IsEnabled,
										},
									},
								},
							},
						)
					}
				}

				if len(val.TokensToUseDefaultFeeConfigs) != 0 || len(val.TokenTransferFeeConfigArgs) != 0 {
					args[src] = val
				}
			}

			var res sequences.OnChainOutput
			for selector, updatesByChain := range args {
				res, err = sequences.RunAndMergeSequence(b, chains,
					evmseq.SequenceFeeQuoterUpdate,
					evmseq.FeeQuoterUpdate{
						ExistingAddresses:             []datastore.AddressRef{feeRef},
						ChainSelector:                 selector,
						TokenTransferFeeConfigUpdates: updatesByChain,
					},
					res,
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterApplyTokenTransferFeeConfigUpdates operation: %w", err)
				}
			}

			return res, nil
		},
	)
}

func (a *FeesAdapter) GetDefaultDestChainConfig(src, dst uint64) lanes.FeeQuoterDestChainConfig {
	return a.evm.GetFeeQuoterDestChainConfig()
}

func (a *FeesAdapter) GetOnchainDestChainConfig(e cldf.Environment, feeRef datastore.AddressRef, src uint64, dst uint64) (lanes.FeeQuoterDestChainConfig, error) {
	err := a.validateFeeRef(feeRef)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("invalid FeeQuoter address ref for src %d and dst %d: %w", src, dst, err)
	}

	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("chain with selector %d not defined", src)
	}
	fqAddr, err := datastore_utils_evm.ToEVMAddress(feeRef)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to convert FeeQuoter address ref to EVM address for src %d and dst %d: %w", src, dst, err)
	}
	fq, err := fqops.NewFeeQuoterContract(fqAddr, chain.Client)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to instantiate FeeQuoter contract at address %s on chain selector %d: %w", fqAddr.Hex(), src, err)
	}

	onchain, err := fq.GetDestChainConfig(&bind.CallOpts{Context: e.GetContext()}, dst)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to get dest chain config from FeeQuoter 2.0 at %s for src %d, dst %d: %w", fqAddr.Hex(), src, dst, err)
	}

	return evmseqV1_6_0.ReverseTranslateFQV2(onchain), nil
}

func (a *FeesAdapter) ApplyDestChainConfigUpdates(e cldf.Environment, feeRef datastore.AddressRef) *operations.Sequence[fees.ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"ApplyDestChainConfigUpdatesV2",
		utils.Version_2_0_0,
		"Applies FeeQuoter 2.0 destination chain config updates",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.ApplyDestChainConfigSequenceInput) (sequences.OnChainOutput, error) {
			src := input.Selector

			err := a.validateFeeRef(feeRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid FeeQuoter address ref for src %d and dst %d: %w", src, 0, err)
			}

			fqAddr, err := datastore_utils_evm.ToEVMAddress(feeRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert FeeQuoter address ref to EVM address: %w", err)
			}

			args := map[uint64][]fqops.DestChainConfigArgs{}
			for dst, cfg := range input.Settings {
				args[src] = append(args[src], fqops.DestChainConfigArgs{
					DestChainSelector: dst,
					DestChainConfig:   evmseqV1_6_0.TranslateFQtoV2(cfg),
				})
			}

			evmChain, ok := chains.EVMChains()[src]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", src)
			}

			writes := make([]contract.WriteOutput, 0, len(args))
			for selector, updates := range args {
				report, err := operations.ExecuteOperation(
					b, fqops.ApplyDestChainConfigUpdates, evmChain,
					contract.FunctionInput[[]fqops.DestChainConfigArgs]{
						ChainSelector: selector,
						Address:       fqAddr,
						Args:          updates,
					},
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates on FeeQuoter 2.0 for chain %d: %w", src, err)
				}
				writes = append(writes, report.Output)
			}

			var result sequences.OnChainOutput
			if len(writes) != 0 {
				if batch, err := contract.NewBatchOperationFromWrites(writes); err != nil {
					return result, fmt.Errorf("failed to create batch operation from writes for chain %d: %w", src, err)
				} else {
					result.BatchOps = append(result.BatchOps, batch)
				}
			}

			return result, nil
		},
	)
}
