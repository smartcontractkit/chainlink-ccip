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
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var _ tokensapi.TokenAdminRegistryReader = (*SolanaAdminRegistryReader)(nil)

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
