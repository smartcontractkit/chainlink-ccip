package adapters

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var _ fees.FeeResolver = (*SolanaFeeResolver)(nil)

type SolanaFeeResolver struct{}

func (a *SolanaFeeResolver) GetOnRampRef(e cldf.Environment, src uint64, dst uint64) (datastore.AddressRef, error) {
	if _, ok := e.BlockChains.SolanaChains()[src]; !ok {
		return datastore.AddressRef{}, fmt.Errorf("solana chain not found for selector %d", src)
	}

	// NOTE: for Solana, the Router and OnRamp are the same program.
	routerRef, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
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
