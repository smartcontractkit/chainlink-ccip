package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	v1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	onramp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ fees.FeeAdapter = (*FeesAdapter)(nil)

type FeesAdapter struct {
	resolver v1_0_0.EVMFeeResolver
}

func NewFeesAdapter() *FeesAdapter {
	return &FeesAdapter{
		resolver: v1_0_0.EVMFeeResolver{},
	}
}

func (a *FeesAdapter) validateFeeRef(feeRef datastore.AddressRef) error {
	if feeRef.Type.String() != onramp.ContractType.String() {
		return fmt.Errorf("unexpected contract type for OnRamp address ref: got %s, want %s", feeRef.Type.String(), onramp.ContractType)
	}
	if !feeRef.Version.Equal(utils.Version_1_5_0) {
		return fmt.Errorf("unexpected OnRamp contract version: got %s, want %s", feeRef.Version, utils.Version_1_5_0)
	}

	return nil
}

func (a *FeesAdapter) GetFeeContractRef(e cldf.Environment, onRampRef datastore.AddressRef, src uint64, dst uint64) (datastore.AddressRef, error) {
	if err := a.validateFeeRef(onRampRef); err != nil {
		return datastore.AddressRef{}, fmt.Errorf("invalid OnRamp address ref for src %d and dst %d: %w", src, dst, err)
	}

	return onRampRef, nil
}

func (a *FeesAdapter) GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) fees.TokenTransferFeeArgs {
	return fees.GetDefaultChainAgnosticTokenTransferFeeConfig(src, dst)
}

func (a *FeesAdapter) GetOnchainTokenTransferFeeConfig(e cldf.Environment, feeRef datastore.AddressRef, src uint64, dst uint64, token string) (fees.TokenTransferFeeArgs, error) {
	err := a.validateFeeRef(feeRef)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("invalid OnRamp address ref for src %d and dst %d: %w", src, dst, err)
	}

	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("chain with selector %d not defined", src)
	}
	onRampAddr, err := datastore_utils_evm.ToEVMAddress(feeRef)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to convert OnRamp address ref to EVM address for src %d and dst %d: %w", src, dst, err)
	}
	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddr, chain.Client)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to instantiate OnRamp contract at address %s on chain selector %d: %w", onRampAddr.Hex(), src, err)
	}
	if !common.IsHexAddress(token) {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("invalid token address: %s", token)
	}

	// This gets the token transfer fee config for the given token from the EVM2EVMOnRamp contract
	// https://sepolia.etherscan.io/address/0xf9765c80F6448e6d4d02BeF4a6b4152131A2F513#code#F1#L719
	cfg, err := onRamp.GetTokenTransferFeeConfig(&bind.CallOpts{Context: e.GetContext()}, common.HexToAddress(token))
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get token transfer fee config from OnRamp at %s for src %d, dst %d, token %s: %w", onRampAddr.Hex(), src, dst, token, err)
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
		utils.Version_1_5_0,
		"Set token transfer fee config on the OnRamp 1.5.0 contract across multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetTokenTransferFeeSequenceInput) (sequences.OnChainOutput, error) {
			err := a.validateFeeRef(feeRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid OnRamp address ref: %w", err)
			}

			onRampAddr, err := datastore_utils_evm.ToEVMAddress(feeRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert OnRamp address ref to EVM address: %w", err)
			}

			args := map[uint64]onramp.SetTokenTransferFeeConfigInput{}
			for dst, dstCfg := range input.Settings {
				src := input.Selector
				val := args[src]
				for rawTokenAddress, feeCfg := range dstCfg {
					var token common.Address
					if !common.IsHexAddress(rawTokenAddress) {
						return sequences.OnChainOutput{}, fmt.Errorf("invalid token address for src %d and dst %d: %s", src, dst, rawTokenAddress)
					} else {
						token = common.HexToAddress(rawTokenAddress)
					}

					if feeCfg == nil {
						val.TokensToUseDefaultFeeConfigs = append(val.TokensToUseDefaultFeeConfigs, token)
					} else {
						val.TokenTransferFeeConfigArgs = append(
							val.TokenTransferFeeConfigArgs,
							evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
								DestBytesOverhead: feeCfg.DestBytesOverhead,
								DestGasOverhead:   feeCfg.DestGasOverhead,
								MinFeeUSDCents:    feeCfg.MinFeeUSDCents,
								MaxFeeUSDCents:    feeCfg.MaxFeeUSDCents,
								DeciBps:           feeCfg.DeciBps,
								Token:             token,

								// NOTE: Aggregate rate limit should be false by default, we do
								// not do lane based rate limits anymore, we limit on pools now
								AggregateRateLimitEnabled: false,
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
					evmseq.OnRampSetTokenTransferFeeConfigSequence,
					evmseq.OnRampSetTokenTransferFeeConfigSequenceInput{
						UpdatesByChain: updatesByChain,
						ChainSelector:  selector,
						Address:        onRampAddr,
					},
					res,
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set token transfer fee config on OnRamp contract at %s: %w", onRampAddr.Hex(), err)
				}
			}

			return res, nil
		},
	)
}

// GetDefaultDestChainConfig is not supported for v1.5 (no FeeQuoter contract).
func (a *FeesAdapter) GetDefaultDestChainConfig(src, dst uint64) lanes.FeeQuoterDestChainConfig {
	return lanes.FeeQuoterDestChainConfig{}
}

// GetOnchainDestChainConfig is not supported for v1.5 (no FeeQuoter contract).
func (a *FeesAdapter) GetOnchainDestChainConfig(_ cldf.Environment, _ datastore.AddressRef, _, _ uint64) (lanes.FeeQuoterDestChainConfig, error) {
	return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("FeeQuoter dest chain config reads are not supported for v1.5 (no FeeQuoter contract)")
}

// ApplyDestChainConfigUpdates is not supported for v1.5 (no FeeQuoter contract).
func (a *FeesAdapter) ApplyDestChainConfigUpdates(e cldf.Environment, feeRef datastore.AddressRef) *operations.Sequence[fees.ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"ApplyDestChainConfigUpdatesV1_5",
		utils.Version_1_5_0,
		"Not supported for v1.5 — no FeeQuoter contract exists",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.ApplyDestChainConfigSequenceInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, fmt.Errorf("FeeQuoter dest chain config updates are not supported for v1.5 (no FeeQuoter contract)")
		},
	)
}
