package adapters

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	solseq "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ fees.FeeAdapter = (*FeesAdapter)(nil)

type FeesAdapter struct {
	sol *solseq.SolanaAdapter
}

func NewFeesAdapter(solAdapter *solseq.SolanaAdapter) *FeesAdapter {
	return &FeesAdapter{
		sol: solAdapter,
	}
}

func (a *FeesAdapter) GetFeeContractRef(e cldf.Environment, src uint64, dst uint64) (datastore.AddressRef, error) {
	ds := e.DataStore
	fqAddr, err := a.sol.GetFQAddress(ds, src)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
	}

	filter := datastore.AddressRef{
		Type:          datastore.ContractType(fqops.ContractType),
		Address:       solana.PublicKeyFromBytes(fqAddr).String(),
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
	chain, ok := e.BlockChains.SolanaChains()[src]
	if !ok {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("solana chain not found for selector %d", src)
	}

	fqRef, err := a.GetFeeContractRef(e, src, dst)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
	}

	fqAddr := solana.MustPublicKeyFromBase58(fqRef.Address)

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

func (a *FeesAdapter) SetTokenTransferFee(e cldf.Environment) *operations.Sequence[fees.SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetTokenTransferFee",
		semver.MustParse("1.6.0"),
		"Sets token transfer fee configuration on CCIP 1.6.0 FeeQuoter contracts",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetTokenTransferFeeSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			src := input.Selector

			fqRef, err := a.GetFeeContractRef(e, src, 0)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
			}

			fqAddr := solana.MustPublicKeyFromBase58(fqRef.Address)

			remoteChainConfigs := map[uint64]map[solana.PublicKey]fee_quoter.TokenTransferFeeConfig{}
			for dst, dstCfg := range input.Settings {
				remoteChainConfigs[dst] = map[solana.PublicKey]fee_quoter.TokenTransferFeeConfig{}
				for rawTokenAddress, feeCfg := range dstCfg {
					token, err := solana.PublicKeyFromBase58(rawTokenAddress)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("invalid base58 token address for src %d and dst %d: %s", src, dst, rawTokenAddress)
					}

					if feeCfg == nil {
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
				return result, nil
			}

			result, err = sequences.RunAndMergeSequence(b, chains,
				solseq.SetTokenTransferFeeConfig,
				solseq.FeeQuoterSetTokenTransferFeeConfigSequenceInput{
					RemoteChainConfigs: remoteChainConfigs,
					DataStore:          e.DataStore,
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

func (a *FeesAdapter) GetOnchainDestChainConfig(e cldf.Environment, src uint64, dst uint64) (lanes.FeeQuoterDestChainConfig, error) {
	chain, ok := e.BlockChains.SolanaChains()[src]
	if !ok {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("solana chain not found for selector %d", src)
	}

	fqAddr, err := a.sol.GetFQAddress(e.DataStore, src)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
	}
	fqPubkey := solana.PublicKeyFromBytes(fqAddr)

	destChainPDA, _, err := state.FindFqDestChainPDA(dst, fqPubkey)
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

func (a *FeesAdapter) ApplyDestChainConfigUpdates(e cldf.Environment) *operations.Sequence[fees.ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"ApplyDestChainConfigUpdatesSolana",
		semver.MustParse("1.6.0"),
		"Applies FeeQuoter destination chain config updates on Solana",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.ApplyDestChainConfigSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput

			solChain, ok := chains.SolanaChains()[input.Selector]
			if !ok {
				return result, fmt.Errorf("solana chain with selector %d not defined", input.Selector)
			}

			fqAddr, err := a.sol.GetFQAddress(e.DataStore, input.Selector)
			if err != nil {
				return result, fmt.Errorf("failed to get FeeQuoter address for chain %d: %w", input.Selector, err)
			}
			fqPubkey := solana.PublicKeyFromBytes(fqAddr)

			offRampAddr, err := a.sol.GetOffRampAddress(e.DataStore, input.Selector)
			if err != nil {
				return result, fmt.Errorf("failed to get OffRamp address for chain %d: %w", input.Selector, err)
			}
			offRampPubkey := solana.PublicKeyFromBytes(offRampAddr)

			for dst, cfg := range input.Settings {
				report, err := operations.ExecuteOperation(b, fqops.ConnectChains, solChain, fqops.ConnectChainsParams{
					FeeQuoter:           fqPubkey,
					OffRamp:             offRampPubkey,
					RemoteChainSelector: dst,
					DestChainConfig:     solseq.TranslateFQ(cfg),
				})
				if err != nil {
					return result, fmt.Errorf("failed to apply dest chain config for dst %d on Solana chain %d: %w", dst, input.Selector, err)
				}
				result.Addresses = append(result.Addresses, report.Output.Addresses...)
				result.BatchOps = append(result.BatchOps, report.Output.BatchOps...)
			}

			return result, nil
		},
	)
}
