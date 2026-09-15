package adapters

import (
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_1/ccip_common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	state "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ tokensapi.TokenAdminRegistryManager = (*SolanaAdminRegistryReader)(nil)

type SolanaAdminRegistryReader struct{}

func (a *SolanaAdminRegistryReader) GetActivePool(e deployment.Environment, chainSelector uint64, tokenRef datastore.AddressRef, overrides ...datastore.AddressRef) ([]byte, error) {
	chain, ok := e.BlockChains.SolanaChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	tokenAddress := tokenRef.Address
	if tokenAddress == "" {
		if ref, err := datastore_utils.FindAndFormatRef(e.DataStore, tokenRef, chainSelector, datastore_utils.FullRef); err != nil {
			return nil, fmt.Errorf("failed to resolve token ref for chain %d: %w", chainSelector, err)
		} else {
			tokenAddress = ref.Address
		}
	}

	mint, err := solana.PublicKeyFromBase58(tokenAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid token mint address %q: %w", tokenAddress, err)
	}

	var router solana.PublicKey
	for _, override := range overrides {
		if !datastore_utils.IsAddressRefEmpty(override) {
			if ref, err := datastore_utils.FindAndFormatRef(e.DataStore, override, chainSelector, datastore_utils.FullRef); err == nil && ref.Address != "" {
				if parsedRouter, parseErr := solana.PublicKeyFromBase58(ref.Address); parseErr != nil {
					return nil, fmt.Errorf("invalid router address %q: %w", ref.Address, parseErr)
				} else {
					router = parsedRouter
					break
				}
			}
		}
	}
	if router.IsZero() {
		if routerRef, err := a.GetTokenAdminRegistryRef(e, chainSelector); err != nil {
			return nil, fmt.Errorf("failed to resolve Router on chain %d: %w", chainSelector, err)
		} else {
			router = solana.MustPublicKeyFromBase58(routerRef.Address)
		}
	}

	tarPDA, _, err := state.FindTokenAdminRegistryPDA(mint, router)
	if err != nil {
		return nil, fmt.Errorf("failed to derive TAR PDA: %w", err)
	}

	var tarAccount ccip_common.TokenAdminRegistry
	if err := chain.GetAccountDataBorshInto(e.OperationsBundle.GetContext(), tarPDA, &tarAccount); err != nil {
		return nil, nil
	}
	if tarAccount.LookupTable.IsZero() {
		return nil, nil
	}

	entries, err := common.GetAddressLookupTable(e.OperationsBundle.GetContext(), chain.Client, tarAccount.LookupTable)
	if errors.Is(err, rpc.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch lookup table %s: %w", tarAccount.LookupTable.String(), err)
	}
	if len(entries) <= 2 {
		return nil, nil
	}

	return entries[2].Bytes(), nil
}

func (a *SolanaAdminRegistryReader) GetTokenAdminRegistryRef(e deployment.Environment, chainSelector uint64) (datastore.AddressRef, error) {
	_, ok := e.BlockChains.SolanaChains()[chainSelector]
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	refs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSelector),
		datastore.AddressRefByType(datastore.ContractType(routerops.ContractType)),
		datastore.AddressRefByVersion(routerops.Version),
	)
	if len(refs) == 0 {
		return datastore.AddressRef{}, fmt.Errorf("Router not found in datastore for chain %d", chainSelector)
	}

	return refs[0], nil
}

// UnregisterToken unregisters a token from the router's TokenAdminRegistry by setting its pool
// lookup table to the zero pubkey. The caller is responsible for the read-check guard (only
// unregister when the token's current active pool is the pool being removed) before executing
// this sequence.
func (a *SolanaAdminRegistryReader) UnregisterToken() *cldf_ops.Sequence[tokensapi.UnregisterTokenSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"solana-admin-registry:unregister-token",
		common_utils.Version_1_0_0,
		"Unregister a token from the router TokenAdminRegistry by setting its pool lookup table to zero",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.UnregisterTokenSequenceInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.SolanaChains()[input.Selector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("solana chain with selector %d not defined", input.Selector)
			}

			routerRef, err := a.GetTokenAdminRegistryRef(deployment.Environment{DataStore: input.ExistingDataStore, BlockChains: chains}, input.Selector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve router on chain %d: %w", input.Selector, err)
			}
			router, err := solana.PublicKeyFromBase58(routerRef.Address)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid router address %q on chain %d: %w", routerRef.Address, input.Selector, err)
			}

			mint, err := solana.PublicKeyFromBase58(input.TokenRef.Address)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid token mint address %q on chain %d: %w", input.TokenRef.Address, input.Selector, err)
			}

			out, err := operations.ExecuteOperation(b, routerops.UnregisterToken, chain, routerops.UnregisterTokenParams{
				Router:    router,
				TokenMint: mint,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to unregister token %s on chain %d: %w", input.TokenRef.Address, input.Selector, err)
			}

			return out.Output, nil
		},
	)
}
