package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	adaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
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
	resolver adaptersV1_0_0.EVMFeeResolver
	evm      *evmseqV1_6_0.EVMAdapter
}

func NewFeesAdapter(evmAdapter *evmseqV1_6_0.EVMAdapter) *FeesAdapter {
	return &FeesAdapter{
		resolver: adaptersV1_0_0.EVMFeeResolver{},
		evm:      evmAdapter,
	}
}

func (a *FeesAdapter) validateFeeRef(feeRef datastore.AddressRef) error {
	if feeRef.Type.String() != fee_quoter.ContractType.String() {
		return fmt.Errorf("unexpected contract type for FeeQuoter address ref: got %s, want %s", feeRef.Type.String(), fee_quoter.ContractType)
	}
	if !utils.StripPatchVersion(feeRef.Version).Equal(utils.Version_1_6_0) {
		return fmt.Errorf("unexpected FeeQuoter contract version: got %s, want %s", feeRef.Version, utils.Version_1_6_0)
	}

	return nil
}

func (a *FeesAdapter) GetFeeContractRef(e cldf.Environment, onRampRef datastore.AddressRef, src uint64, dst uint64) (datastore.AddressRef, error) {
	if onRampRef.Type.String() != onramp.ContractType.String() {
		return datastore.AddressRef{}, fmt.Errorf("unexpected contract type for OnRamp address ref for src %d and dst %d: got %s, want %s", src, dst, onRampRef.Type.String(), onramp.ContractType)
	}
	if !onRampRef.Version.Equal(utils.Version_1_6_0) {
		return datastore.AddressRef{}, fmt.Errorf("unexpected OnRamp contract version for src %d and dst %d: got %s, want %s", src, dst, onRampRef.Version, utils.Version_1_6_0)
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

	// NOTE: version is omitted intentionally here - we could have a v1.6 OnRamp connected to a v2.0 FQ for example
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

func (a *FeesAdapter) GetOnchainTokenTransferFeeConfig(e cldf.Environment, feeRef datastore.AddressRef, src uint64, dst uint64, token string) (fees.TokenTransferFeeArgs, error) {
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
	fq, err := fee_quoter.NewFeeQuoterContract(fqAddr, chain.Client)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to instantiate FeeQuoter contract at address %s on chain selector %d: %w", fqAddr.Hex(), src, err)
	}
	if !common.IsHexAddress(token) {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("invalid token address: %s", token)
	}

	// This gets the token transfer fee config for the given token from the FeeQuoter contract
	// https://etherscan.io/address/0x40858070814a57FdF33a613ae84fE0a8b4a874f7#code#F1#L819
	cfg, err := fq.GetTokenTransferFeeConfig(&bind.CallOpts{Context: e.GetContext()}, dst, common.HexToAddress(token))
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get token transfer fee config from FeeQuoter at %s for src %d, dst %d, token %s: %w", fqAddr.Hex(), src, dst, token, err)
	}

	e.Logger.Infof("Fetched on-chain token transfer fee config for src %d, dst %d, token %s: %+v", src, dst, token, cfg)
	return fees.TokenTransferFeeArgs{
		DestBytesOverhead: cfg.DestBytesOverhead,
		DestGasOverhead:   cfg.DestGasOverhead,
		MinFeeUSDCents:    cfg.MinFeeUSDCents,
		MaxFeeUSDCents:    cfg.MaxFeeUSDCents,
		IsEnabled:         cfg.IsEnabled,
		DeciBps:           cfg.DeciBps,
	}, nil
}

func (a *FeesAdapter) SetTokenTransferFee(e cldf.Environment, feeRef datastore.AddressRef) *operations.Sequence[fees.SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetTokenTransferFee",
		utils.Version_1_6_0,
		"Sets token transfer fee configuration on CCIP 1.6.0 FeeQuoter contracts",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetTokenTransferFeeSequenceInput) (sequences.OnChainOutput, error) {
			src := input.Selector

			err := a.validateFeeRef(feeRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid FeeQuoter address ref: %w", err)
			}

			fqAddr, err := datastore_utils_evm.ToEVMAddress(feeRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert FeeQuoter address ref to EVM address: %w", err)
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
											MinFeeUSDCents:    feeCfg.MinFeeUSDCents,
											MaxFeeUSDCents:    feeCfg.MaxFeeUSDCents,
											IsEnabled:         feeCfg.IsEnabled,
											DeciBps:           feeCfg.DeciBps,
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
					evmseqV1_6_0.FeeQuoterApplyTokenTransferFeeConfigUpdatesSequence,
					evmseqV1_6_0.FeeQuoterApplyTokenTransferFeeConfigUpdatesSequenceInput{
						UpdatesByChain: updatesByChain,
						ChainSelector:  selector,
						Address:        fqAddr,
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
	fq, err := fee_quoter.NewFeeQuoterContract(fqAddr, chain.Client)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to instantiate FeeQuoter contract at address %s on chain selector %d: %w", fqAddr.Hex(), src, err)
	}

	onchain, err := fq.GetDestChainConfig(&bind.CallOpts{Context: e.GetContext()}, dst)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to get dest chain config from FeeQuoter at %s for src %d, dst %d: %w", fqAddr.Hex(), src, dst, err)
	}

	return evmseqV1_6_0.ReverseTranslateFQ(fqops.DestChainConfig{
		IsEnabled:                         onchain.IsEnabled,
		MaxNumberOfTokensPerMsg:           onchain.MaxNumberOfTokensPerMsg,
		MaxDataBytes:                      onchain.MaxDataBytes,
		MaxPerMsgGasLimit:                 onchain.MaxPerMsgGasLimit,
		DestGasOverhead:                   onchain.DestGasOverhead,
		DestGasPerPayloadByteBase:         onchain.DestGasPerPayloadByteBase,
		DestGasPerPayloadByteHigh:         onchain.DestGasPerPayloadByteHigh,
		DestGasPerPayloadByteThreshold:    onchain.DestGasPerPayloadByteThreshold,
		DestDataAvailabilityOverheadGas:   onchain.DestDataAvailabilityOverheadGas,
		DestGasPerDataAvailabilityByte:    onchain.DestGasPerDataAvailabilityByte,
		DestDataAvailabilityMultiplierBps: onchain.DestDataAvailabilityMultiplierBps,
		ChainFamilySelector:               onchain.ChainFamilySelector,
		EnforceOutOfOrder:                 onchain.EnforceOutOfOrder,
		DefaultTokenFeeUSDCents:           onchain.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead:       onchain.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:                 onchain.DefaultTxGasLimit,
		GasMultiplierWeiPerEth:            onchain.GasMultiplierWeiPerEth,
		GasPriceStalenessThreshold:        onchain.GasPriceStalenessThreshold,
		NetworkFeeUSDCents:                onchain.NetworkFeeUSDCents,
	}), nil
}

func (a *FeesAdapter) ApplyDestChainConfigUpdates(e cldf.Environment, feeRef datastore.AddressRef) *operations.Sequence[fees.ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"ApplyDestChainConfigUpdates",
		utils.Version_1_6_0,
		"Applies FeeQuoter 1.6 destination chain config updates",
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
					DestChainConfig:   evmseqV1_6_0.TranslateFQ(cfg),
				})
			}

			var res sequences.OnChainOutput
			for selector, updates := range args {
				res, err = sequences.RunAndMergeSequence(b, chains,
					evmseqV1_6_0.FeeQuoterApplyDestChainConfigUpdatesSequence,
					evmseqV1_6_0.FeeQuoterApplyDestChainConfigUpdatesSequenceInput{
						Address:        fqAddr,
						ChainSelector:  selector,
						UpdatesByChain: updates,
					},
					res,
				)
				if err != nil {
					return res, fmt.Errorf("failed to apply dest chain config updates on FeeQuoter 1.6 for chain %d: %w", src, err)
				}
			}

			return res, nil
		},
	)
}
