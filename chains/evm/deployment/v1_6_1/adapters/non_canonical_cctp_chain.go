package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	tokens "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ adapters.CCTPChain = &NonCanonicalCCTPChainAdapter{}

// NonCanonicalCCTPChainAdapter is the adapter for non-canonical CCTP chains.
type NonCanonicalCCTPChainAdapter struct{}

// DeployCCTPChain returns the sequence for deploying a CCTP chain.
func (c *NonCanonicalCCTPChainAdapter) DeployCCTPChain() *operations.Sequence[adapters.DeployCCTPInput, seq_core.OnChainOutput, adapters.DeployCCTPChainDeps] {
	return tokens.DeployNonCanonicalCCTPChain
}

// ConfigureCCTPChainForLanes returns the sequence for configuring a CCTP chain for lanes.
func (c *NonCanonicalCCTPChainAdapter) ConfigureCCTPChainForLanes() *operations.Sequence[adapters.ConfigureCCTPChainForLanesInput, seq_core.OnChainOutput, adapters.ConfigureCCTPChainForLanesDeps] {
	return tokens.ConfigureNonCanonicalCCTPChainForLanes
}

// CCTPV1AllowedCallerOnDest is not implemented for non-canonical CCTP chains, as there is no caller of CCTP.
func (c *NonCanonicalCCTPChainAdapter) CCTPV1AllowedCallerOnDest(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil
}

// CCTPV2AllowedCallerOnDest is not implemented for non-canonical CCTP chains, as there is no caller of CCTP.
func (c *NonCanonicalCCTPChainAdapter) CCTPV2AllowedCallerOnDest(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil
}

// AllowedCallerOnSource is not implemented for non-canonical CCTP chains, as USDC was simply locked on source when transferred to this chain.
func (c *NonCanonicalCCTPChainAdapter) AllowedCallerOnSource(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil
}

// MintRecipientOnDest is not implemented for non-canonical CCTP chains, as there is no mint recipient.
func (c *NonCanonicalCCTPChainAdapter) MintRecipientOnDest(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil
}

// PoolAddress returns the address of the token pool on the remote chain in bytes.
func (c *NonCanonicalCCTPChainAdapter) PoolAddress(d datastore.DataStore, b chain.BlockChains, chainSelector uint64, registeredPoolRef datastore.AddressRef) ([]byte, error) {
	registeredPoolAddress, err := datastore_utils.FindAndFormatRef(d, registeredPoolRef, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find registered pool address: %w", err)
	}
	return common.FromHex(registeredPoolAddress.Address), nil
}

// TokenAddress returns the address of the token on the remote chain in bytes.
func (c *NonCanonicalCCTPChainAdapter) TokenAddress(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	poolAddressRef, err := datastore_utils.FindAndFormatRef(d, datastore.AddressRef{
		// We expect at most one BurnMintWithLockReleaseFlagTokenPool deployed on any given non-canonical chain.
		Type: datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType),
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find pool address: %w", err)
	}
	chain, ok := b.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	boundTokenPool, err := token_pool.NewTokenPoolContract(common.HexToAddress(poolAddressRef.Address), chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind token pool: %w", err)
	}
	tokenAddress, err := boundTokenPool.GetToken(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get token address: %w", err)
	}
	return tokenAddress.Bytes(), nil
}
