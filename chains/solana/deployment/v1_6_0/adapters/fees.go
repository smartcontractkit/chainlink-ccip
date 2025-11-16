package adapters

import (
	"errors"
	"fmt"
	"math"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	solseq "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
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

func (a *FeesAdapter) getFeeQuoterAddress(ds datastore.DataStore, src uint64) (solana.PublicKey, error) {
	fqAddr, err := a.sol.GetFQAddress(ds, src)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
	}

	return solana.PublicKeyFromBytes(fqAddr), nil
}

func (a *FeesAdapter) GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) fees.TokenTransferFeeArgs {
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

func (a *FeesAdapter) GetOnchainTokenTransferFeeConfig(e cldf.Environment, src uint64, dst uint64, address string) (fees.TokenTransferFeeArgs, error) {
	chain, ok := e.BlockChains.SolanaChains()[src]
	if !ok {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("solana chain not found for selector %d", src)
	}

	fqAddr, err := a.getFeeQuoterAddress(e.DataStore, src)
	if err != nil {
		return fees.TokenTransferFeeArgs{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
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

func (a *FeesAdapter) SetTokenTransferFee(e cldf.Environment) *operations.Sequence[fees.SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetTokenTransferFee",
		semver.MustParse("1.6.0"),
		"Sets token transfer fee configuration on CCIP 1.6.0 FeeQuoter contracts",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetTokenTransferFeeSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			src := input.Selector

			fqAddr, err := a.getFeeQuoterAddress(e.DataStore, src)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
			}

			remoteChainConfigs := map[uint64]map[solana.PublicKey]fee_quoter.TokenTransferFeeConfig{}
			for dst, dstCfg := range input.Settings {
				remoteChainConfigs[dst] = map[solana.PublicKey]fee_quoter.TokenTransferFeeConfig{}
				for rawTokenAddress, feeCfg := range dstCfg {
					token, err := solana.PublicKeyFromBase58(rawTokenAddress)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("invalid base58 token address for src %d and dst %d: %s", src, dst, rawTokenAddress)
					}

					if feeCfg == nil {
						remoteChainConfigs[dst][token] = fee_quoter.TokenTransferFeeConfig{
							DestBytesOverhead: 0,
							DestGasOverhead:   0,
							MinFeeUsdcents:    0,
							MaxFeeUsdcents:    0,
							DeciBps:           0,
							IsEnabled:         false,
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
