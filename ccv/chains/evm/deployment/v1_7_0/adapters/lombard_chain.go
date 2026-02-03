package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/lombard"
)

var _ adapters.LombardChain = &LombardChainAdapter{}

// LombardChainAdapter is the adapter for Lombard chains.
type LombardChainAdapter struct{}

// AddressRefToBytes returns the byte representation of an address for this chain family.
func (c *LombardChainAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return common.HexToAddress(ref.Address).Bytes(), nil
}

// DeployLombardChain returns the sequence for deploying a Lombard chain.
func (c *LombardChainAdapter) DeployLombardChain() *operations.Sequence[adapters.DeployLombardInput, seq_core.OnChainOutput, adapters.DeployLombardChainDeps] {
	return lombard.DeployLombardChain
}

func (c *LombardChainAdapter) ConfigureLombardChainForLanes() *operations.Sequence[adapters.ConfigureLombardChainForLanesInput, seq_core.OnChainOutput, adapters.ConfigureLombardChainForLanesDeps] {
	return lombard.ConfigureLombardChainForLanes
}

// AllowedCallerOnDest returns the address allowed to deposit tokens for burn on the remote chain.
// On EVM, the caller of Lombard is LombardVerifier
func (c *LombardChainAdapter) AllowedCallerOnDest(ds datastore.DataStore, chains chain.BlockChains, chainSelector uint64) ([]byte, error) {
	allowedCallerOnSourceAddressRef, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(lombard_verifier.ContractType),
		Version: lombard_verifier.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find allowed caller on source address: %w", err)
	}
	return common.FromHex(allowedCallerOnSourceAddressRef.Address), nil
}

func (c *LombardChainAdapter) TokenPool(ds datastore.DataStore, chains chain.BlockChains, selector uint64) (datastore.AddressRef, error) {
	return datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(lombard_token_pool.ContractType),
		Version: lombard_token_pool.Version,
	}, selector, datastore_utils.FullRef)
}
