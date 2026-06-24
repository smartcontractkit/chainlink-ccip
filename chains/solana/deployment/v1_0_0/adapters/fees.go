package adapters

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ fees.FeeResolver = (*SolanaFeeResolver)(nil)

type SolanaFeeResolver struct{}

func (a *SolanaFeeResolver) GetOnRampRef(_ cldf_ops.Bundle, chains cldf_chain.BlockChains, ds datastore.DataStore, src uint64, dst uint64) (datastore.AddressRef, error) {
	if _, ok := chains.SolanaChains()[src]; !ok {
		return datastore.AddressRef{}, fmt.Errorf("solana chain not found for selector %d", src)
	}

	// NOTE: for Solana, the Router and OnRamp are the same program.
	routerRef, err := datastore_utils.FindAndFormatRef(
		ds,
		datastore.AddressRef{
			Type:          datastore.ContractType(router.ContractType),
			Version:       router.Version,
			ChainSelector: src,
		},
		src,
		datastore_utils.FullRef,
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to get Router address ref for src %d and dst %d: %w", src, dst, err)
	}

	return routerRef, nil
}
