package adapters

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	solseq "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fqdests"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var _ fqdests.FQDestsAdapter = (*FQDestsAdapter)(nil)

type FQDestsAdapter struct {
	sol *solseq.SolanaAdapter
}

func NewFQDestsAdapter(solAdapter *solseq.SolanaAdapter) *FQDestsAdapter {
	return &FQDestsAdapter{sol: solAdapter}
}

func (a *FQDestsAdapter) GetFeeContractRef(e cldf.Environment, src, dst uint64) (datastore.AddressRef, error) {
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
	ref, err := datastore_utils.FindAndFormatRef(ds, filter, src, datastore_utils.FullRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to find FeeQuoter address ref for chain selector %d: %w", src, err)
	}
	return ref, nil
}

func (a *FQDestsAdapter) GetDefaultDestChainConfig(src, dst uint64) lanes.FeeQuoterDestChainConfig {
	return a.sol.GetFeeQuoterDestChainConfig()
}

func (a *FQDestsAdapter) GetOnchainDestChainConfig(e cldf.Environment, src, dst uint64) (lanes.FeeQuoterDestChainConfig, error) {
	chain, ok := e.BlockChains.SolanaChains()[src]
	if !ok {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("solana chain not found for selector %d", src)
	}

	fqRef, err := a.GetFeeContractRef(e, src, dst)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, err
	}

	fqAddr := solana.MustPublicKeyFromBase58(fqRef.Address)
	fqRemoteChainPDA, _, err := state.FindFqDestChainPDA(dst, fqAddr)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to find FQ dest chain PDA for src %d, dst %d: %w", src, dst, err)
	}

	var destChainState fee_quoter.DestChain
	err = chain.GetAccountDataBorshInto(e.GetContext(), fqRemoteChainPDA, &destChainState)
	if err != nil {
		if errors.Is(err, rpc.ErrNotFound) {
			return lanes.FeeQuoterDestChainConfig{}, nil
		}
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to read dest chain state for src %d, dst %d: %w", src, dst, err)
	}

	cfg := destChainState.Config
	return lanes.FeeQuoterDestChainConfig{
		IsEnabled:                   cfg.IsEnabled,
		MaxDataBytes:                cfg.MaxDataBytes,
		MaxPerMsgGasLimit:           cfg.MaxPerMsgGasLimit,
		DestGasOverhead:             cfg.DestGasOverhead,
		DestGasPerPayloadByteBase:   uint8(cfg.DestGasPerPayloadByteBase),
		ChainFamilySelector:         bytesToUint32(cfg.ChainFamilySelector),
		DefaultTokenFeeUSDCents:     uint16(cfg.DefaultTokenFeeUsdcents),
		DefaultTokenDestGasOverhead: cfg.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:           cfg.DefaultTxGasLimit,
		NetworkFeeUSDCents:          uint16(cfg.NetworkFeeUsdcents),
		V1Params: &lanes.FeeQuoterV1Params{
			MaxNumberOfTokensPerMsg:           uint16(cfg.MaxNumberOfTokensPerMsg),
			DestGasPerPayloadByteHigh:         uint8(cfg.DestGasPerPayloadByteHigh),
			DestGasPerPayloadByteThreshold:    uint16(cfg.DestGasPerPayloadByteThreshold),
			DestDataAvailabilityOverheadGas:   cfg.DestDataAvailabilityOverheadGas,
			DestGasPerDataAvailabilityByte:    cfg.DestGasPerDataAvailabilityByte,
			DestDataAvailabilityMultiplierBps: cfg.DestDataAvailabilityMultiplierBps,
			EnforceOutOfOrder:                 cfg.EnforceOutOfOrder,
			GasMultiplierWeiPerEth:            cfg.GasMultiplierWeiPerEth,
			GasPriceStalenessThreshold:        cfg.GasPriceStalenessThreshold,
		},
	}, nil
}

func (a *FQDestsAdapter) ApplyDestChainConfigUpdates(e cldf.Environment) *operations.Sequence[
	fqdests.ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains,
] {
	return operations.NewSequence(
		"FQDestsApplyDestChainConfigUpdates",
		semver.MustParse("1.6.0"),
		"Applies FeeQuoter 1.6 Solana destination chain config updates",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fqdests.ApplyDestChainConfigSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput

			solChain, ok := chains.SolanaChains()[input.Selector]
			if !ok {
				return result, fmt.Errorf("solana chain not found for selector %d", input.Selector)
			}

			fqRef, err := a.GetFeeContractRef(e, input.Selector, 0)
			if err != nil {
				return result, fmt.Errorf("failed to get FeeQuoter address for chain %d: %w", input.Selector, err)
			}
			fqAddr := solana.MustPublicKeyFromBase58(fqRef.Address)

			offRampAddr, err := a.sol.GetOffRampAddress(e.DataStore, input.Selector)
			if err != nil {
				return result, fmt.Errorf("failed to get OffRamp address for chain %d: %w", input.Selector, err)
			}

			for dst, cfg := range input.Settings {
				report, err := operations.ExecuteOperation(b, fqops.ConnectChains, solChain, fqops.ConnectChainsParams{
					FeeQuoter:           fqAddr,
					OffRamp:             solana.PublicKeyFromBytes(offRampAddr),
					RemoteChainSelector: dst,
					DestChainConfig:     solseq.TranslateFQ(cfg),
				})
				if err != nil {
					return result, fmt.Errorf("failed to apply dest chain config for dst %d on FeeQuoter %s: %w", dst, fqAddr, err)
				}

				result.Addresses = append(result.Addresses, report.Output.Addresses...)
				result.BatchOps = append(result.BatchOps, report.Output.BatchOps...)
			}

			return result, nil
		},
	)
}

func bytesToUint32(b [4]byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}
