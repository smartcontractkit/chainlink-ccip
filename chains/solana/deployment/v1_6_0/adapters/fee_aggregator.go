package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	solseq "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var _ fees.FeeAggregatorAdapter = (*FeeAggregatorAdapter)(nil)

type FeeAggregatorAdapter struct {
	sol *solseq.SolanaAdapter
}

func NewFeeAggregatorAdapter(solAdapter *solseq.SolanaAdapter) *FeeAggregatorAdapter {
	return &FeeAggregatorAdapter{
		sol: solAdapter,
	}
}

func (a *FeeAggregatorAdapter) GetFeeAggregator(e cldf.Environment, chainSelector uint64) (string, error) {
	chain, ok := e.BlockChains.SolanaChains()[chainSelector]
	if !ok {
		return "", fmt.Errorf("solana chain not found for selector %d", chainSelector)
	}

	routerAddr, err := a.sol.GetRouterAddress(e.DataStore, chainSelector)
	if err != nil {
		return "", fmt.Errorf("failed to get Router address for chain selector %d: %w", chainSelector, err)
	}
	routerPubkey := solana.PublicKeyFromBytes(routerAddr)

	configPDA, _, err := state.FindConfigPDA(routerPubkey)
	if err != nil {
		return "", fmt.Errorf("failed to derive config PDA for router %s on chain %d: %w", routerPubkey, chainSelector, err)
	}

	var cfg ccip_router.Config
	if err := chain.GetAccountDataBorshInto(e.GetContext(), configPDA, &cfg); err != nil {
		return "", fmt.Errorf("failed to read router config on chain %d: %w", chainSelector, err)
	}

	return cfg.FeeAggregator.String(), nil
}

func (a *FeeAggregatorAdapter) resolveRouterPubkey(e cldf.Environment, input fees.SetFeeAggregatorSequenceInput) (solana.PublicKey, error) {
	if len(input.Contracts) > 0 {
		if len(input.Contracts) != 1 {
			return solana.PublicKey{}, fmt.Errorf("Solana 1.6 adapter supports exactly one contract ref, got %d", len(input.Contracts))
		}
		ref := input.Contracts[0]
		if string(ref.Type) != router.ContractType.String() {
			return solana.PublicKey{}, fmt.Errorf("Solana 1.6 adapter only supports contract type %q, got %q", router.ContractType, ref.Type)
		}
		resolved, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return solana.PublicKey{}, fmt.Errorf("failed to resolve Router ref on chain %d: %w", input.ChainSelector, err)
		}
		return solana.MustPublicKeyFromBase58(resolved.Address), nil
	}
	routerAddr, err := a.sol.GetRouterAddress(e.DataStore, input.ChainSelector)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to get Router address for chain selector %d: %w", input.ChainSelector, err)
	}
	return solana.PublicKeyFromBytes(routerAddr), nil
}

func (a *FeeAggregatorAdapter) SetFeeAggregator(e cldf.Environment) *operations.Sequence[fees.SetFeeAggregatorSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetFeeAggregator",
		semver.MustParse("1.6.0"),
		"Sets the fee aggregator address on Solana 1.6.0 Router",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetFeeAggregatorSequenceInput) (sequences.OnChainOutput, error) {
			solChain, ok := chains.SolanaChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("solana chain not found for selector %d", input.ChainSelector)
			}

			feeAggregatorPubkey, err := solana.PublicKeyFromBase58(input.FeeAggregator)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid fee aggregator base58 address %q: %w", input.FeeAggregator, err)
			}

			routerPubkey, err := a.resolveRouterPubkey(e, input)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			ccip_router.SetProgramID(routerPubkey)
			authority := router.GetAuthority(solChain, routerPubkey)

			routerConfigPDA, _, err := state.FindConfigPDA(routerPubkey)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to derive config PDA for router %s: %w", routerPubkey, err)
			}

			ixn, err := ccip_router.NewUpdateFeeAggregatorInstruction(
				feeAggregatorPubkey,
				routerConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build UpdateFeeAggregator instruction: %w", err)
			}

			if authority != solChain.DeployerKey.PublicKey() {
				batch, err := utils.BuildMCMSBatchOperation(
					solChain.Selector,
					[]solana.Instruction{ixn},
					routerPubkey.String(),
					router.ContractType.String(),
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create MCMS batch operation: %w", err)
				}
				return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batch}}, nil
			}

			if err := solChain.Confirm([]solana.Instruction{ixn}); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm UpdateFeeAggregator instruction: %w", err)
			}

			return sequences.OnChainOutput{}, nil
		},
	)
}
