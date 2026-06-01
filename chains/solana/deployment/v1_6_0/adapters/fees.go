package adapters

import (
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	datastore_utils_solana "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	adaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_0_0/adapters"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	solseq "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ fees.FeeAdapter = (*FeesAdapter)(nil)

type FeesAdapter struct {
	resolver adaptersV1_0_0.SolanaFeeResolver
	sol      *solseq.SolanaAdapter
}

func NewFeesAdapter(solAdapter *solseq.SolanaAdapter) *FeesAdapter {
	return &FeesAdapter{
		resolver: adaptersV1_0_0.SolanaFeeResolver{},
		sol:      solAdapter,
	}
}

func (a *FeesAdapter) validateFeeRef(feeRef datastore.AddressRef) error {
	if feeRef.Type.String() != fqops.ContractType.String() {
		return fmt.Errorf("unexpected contract type for FeeQuoter address ref: got %s, want %s", feeRef.Type.String(), fqops.ContractType)
	}
	if !feeRef.Version.Equal(utils.Version_1_6_0) {
		return fmt.Errorf("unexpected FeeQuoter contract version: got %s, want %s", feeRef.Version, utils.Version_1_6_0)
	}

	return nil
}

func (a *FeesAdapter) GetFeeContractRef(e cldf.Environment, onRampRef datastore.AddressRef, src uint64, dst uint64) (datastore.AddressRef, error) {
	if onRampRef.Type.String() != routerops.ContractType.String() {
		return datastore.AddressRef{}, fmt.Errorf("unexpected contract type for Router address ref for src %d and dst %d: got %s, want %s", src, dst, onRampRef.Type.String(), routerops.ContractType)
	}
	if !onRampRef.Version.Equal(utils.Version_1_6_0) {
		return datastore.AddressRef{}, fmt.Errorf("unexpected Router contract version for src %d and dst %d: got %s, want %s", src, dst, onRampRef.Version, utils.Version_1_6_0)
	}

	routerPubkey, err := solana.PublicKeyFromBase58(onRampRef.Address)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to parse Router address %q for chain selector %d: %w", onRampRef.Address, src, err)
	} else {
		ccip_router.SetProgramID(routerPubkey)
	}

	chain, ok := e.BlockChains.SolanaChains()[src]
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("solana chain not found for selector %d", src)
	}

	routerConfigPDA, _, err := state.FindConfigPDA(routerPubkey)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to derive Router config PDA for router %s on chain selector %d: %w", routerPubkey, src, err)
	}

	var routerConfig ccip_router.Config
	if err := chain.GetAccountDataBorshInto(e.GetContext(), routerConfigPDA, &routerConfig); err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to read Router config for router %s on chain selector %d: %w", routerPubkey, src, err)
	}

	fqRef, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			Type:          datastore.ContractType(fqops.ContractType),
			Address:       routerConfig.FeeQuoter.String(),
			ChainSelector: src,
		},
		src,
		datastore_utils.FullRef,
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to find FeeQuoter address ref for chain selector %d: %w", src, err)
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

	chain, ok := e.BlockChains.SolanaChains()[src]
	if !ok {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("solana chain not found for selector %d", src)
	}
	fqAddr, err := datastore_utils_solana.ToAddress(feeRef)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to convert FeeQuoter address ref to solana.PublicKey for chain selector %d: %w", src, err)
	}
	token, err := solana.PublicKeyFromBase58(address)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("invalid base58 token address: %s", address)
	}

	remoteBillingPDA, _, err := state.FindFqPerChainPerTokenConfigPDA(dst, token, fqAddr)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to find remote billing token config pda (src = %d, dst = %d, token = %s): %w", src, dst, token, err)
	}

	// NOTE: ErrNotFound is expected if no config has been set on-chain yet - we return a zeroed config in that case
	var cfg fee_quoter.PerChainPerTokenConfig
	err = chain.GetAccountDataBorshInto(e.GetContext(), remoteBillingPDA, &cfg)
	if err != nil && !errors.Is(err, rpc.ErrNotFound) {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to deserialize PerChainPerTokenConfig (src = %d, dst = %d, token = %s, pda = %s): %w", src, dst, token, remoteBillingPDA, err)
	}

	e.Logger.Infof("Fetched on-chain token transfer fee config for src %d, dst %d, token %s: %+v", src, dst, token, cfg)
	return fees.TokenTransferFeeArgs{
		DestBytesOverhead: cfg.TokenTransferConfig.DestBytesOverhead,
		DestGasOverhead:   cfg.TokenTransferConfig.DestGasOverhead,
		MinFeeUSDCents:    cfg.TokenTransferConfig.MinFeeUsdcents,
		MaxFeeUSDCents:    cfg.TokenTransferConfig.MaxFeeUsdcents,
		IsEnabled:         cfg.TokenTransferConfig.IsEnabled,
		DeciBps:           cfg.TokenTransferConfig.DeciBps,
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

			fqAddr, err := solana.PublicKeyFromBase58(feeRef.Address)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to parse FeeQuoter address %q for chain selector %d: %w", feeRef.Address, src, err)
			}

			remoteChainConfigs := map[uint64]map[solana.PublicKey]fee_quoter.TokenTransferFeeConfig{}
			for dst, dstCfg := range input.Settings {
				remoteChainConfigs[dst] = map[solana.PublicKey]fee_quoter.TokenTransferFeeConfig{}
				for rawTokenAddress, feeCfg := range dstCfg {
					token, err := solana.PublicKeyFromBase58(rawTokenAddress)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("invalid base58 token address for src %d and dst %d: %s", src, dst, rawTokenAddress)
					}
					if feeCfg == nil || !feeCfg.IsEnabled {
						// NOTE: the Solana FeeQuoter will always perform validation checks on the input
						// config even if we are trying to disable it. As a result, we need to provide a
						// proper set of values even though none of them will actually be used.
						defaults := a.GetDefaultTokenTransferFeeConfig(src, dst)
						remoteChainConfigs[dst][token] = fee_quoter.TokenTransferFeeConfig{
							DestBytesOverhead: defaults.DestBytesOverhead,
							DestGasOverhead:   defaults.DestGasOverhead,
							MinFeeUsdcents:    defaults.MinFeeUSDCents,
							MaxFeeUsdcents:    defaults.MaxFeeUSDCents,
							IsEnabled:         false, // disable the fee config
							DeciBps:           defaults.DeciBps,
						}
					} else {
						remoteChainConfigs[dst][token] = fee_quoter.TokenTransferFeeConfig{
							DestBytesOverhead: feeCfg.DestBytesOverhead,
							DestGasOverhead:   feeCfg.DestGasOverhead,
							MinFeeUsdcents:    feeCfg.MinFeeUSDCents,
							MaxFeeUsdcents:    feeCfg.MaxFeeUSDCents,
							IsEnabled:         feeCfg.IsEnabled,
							DeciBps:           feeCfg.DeciBps,
						}
					}
				}
			}

			if len(remoteChainConfigs) == 0 {
				return sequences.OnChainOutput{}, nil
			}

			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(b, chains,
				solseq.SetTokenTransferFeeConfig,
				solseq.FeeQuoterSetTokenTransferFeeConfigSequenceInput{
					RemoteChainConfigs: remoteChainConfigs,
					FeeQuoter:          fqAddr,
					Selector:           input.Selector,
				},
				result,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set token transfer fee config on FeeQuoter %s for chain selector %d: %w", fqAddr, src, err)
			}

			return result, nil
		},
	)
}

func (a *FeesAdapter) GetDefaultDestChainConfig(src, dst uint64) lanes.FeeQuoterDestChainConfig {
	return a.sol.GetFeeQuoterDestChainConfig()
}

func (a *FeesAdapter) GetOnchainDestChainConfig(e cldf.Environment, feeRef datastore.AddressRef, src uint64, dst uint64) (lanes.FeeQuoterDestChainConfig, error) {
	err := a.validateFeeRef(feeRef)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("invalid FeeQuoter address ref for src %d and dst %d: %w", src, dst, err)
	}

	chain, ok := e.BlockChains.SolanaChains()[src]
	if !ok {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("solana chain not found for selector %d", src)
	}
	fqAddr, err := solana.PublicKeyFromBase58(feeRef.Address)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to parse FeeQuoter address %q for chain selector %d: %w", feeRef.Address, src, err)
	}

	destChainPDA, _, err := state.FindFqDestChainPDA(dst, fqAddr)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to derive DestChain PDA for dst %d on chain %d: %w", dst, src, err)
	}

	var destChainAccount fee_quoter.DestChain
	err = chain.GetAccountDataBorshInto(e.GetContext(), destChainPDA, &destChainAccount)
	if err != nil {
		if errors.Is(err, rpc.ErrNotFound) {
			return lanes.FeeQuoterDestChainConfig{}, nil
		}
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to read DestChain account for dst %d on chain %d: %w", dst, src, err)
	}

	return solseq.ReverseTranslateFQ(destChainAccount.Config), nil
}

func (a *FeesAdapter) ApplyDestChainConfigUpdates(e cldf.Environment, feeRef datastore.AddressRef) *operations.Sequence[fees.ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"ApplyDestChainConfigUpdatesSolana",
		utils.Version_1_6_0,
		"Applies FeeQuoter destination chain config updates on Solana",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.ApplyDestChainConfigSequenceInput) (sequences.OnChainOutput, error) {
			src := input.Selector

			chain, ok := chains.SolanaChains()[src]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("solana chain with selector %d not defined", src)
			}

			err := a.validateFeeRef(feeRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid FeeQuoter address ref for src %d and dst %d: %w", src, 0, err)
			}

			fqAddr, err := solana.PublicKeyFromBase58(feeRef.Address)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to parse FeeQuoter address %q for chain selector %d: %w", feeRef.Address, src, err)
			}

			offRampAddr, err := a.sol.GetOffRampAddress(e.DataStore, src)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get OffRamp address for chain %d: %w", src, err)
			}

			var res sequences.OnChainOutput
			for dst, cfg := range input.Settings {
				report, err := operations.ExecuteOperation(b, fqops.ConnectChains, chain, fqops.ConnectChainsParams{
					FeeQuoter:           fqAddr,
					OffRamp:             solana.PublicKeyFromBytes(offRampAddr),
					RemoteChainSelector: dst,
					DestChainConfig:     solseq.TranslateFQ(cfg),
				})
				if err != nil {
					return res, fmt.Errorf("failed to apply dest chain config for dst %d on Solana chain %d: %w", dst, input.Selector, err)
				}
				res.Addresses = append(res.Addresses, report.Output.Addresses...)
				res.BatchOps = append(res.BatchOps, report.Output.BatchOps...)
			}

			return res, nil
		},
	)
}
