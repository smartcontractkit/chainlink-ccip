package adapters

import (
	"fmt"
	"sync"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type DeployLombardInput struct {
	// ChainSelector is the selector for the chain being deployed.
	ChainSelector uint64
	// Bridge is the address of the Bridge contract provided by Lombard
	Bridge string
	// Token is the address of the token to be used in the LombardTokenPool.
	Token string
	// DeployerContract is a contract that can be used to deploy other contracts.
	// i.e. A CREATE2Factory contract on Ethereum can enable consistent deployments.
	DeployerContract string
	// StorageLocations is the set of storage locations for the LombardVerifier contract.
	StorageLocations []string
	// FeeAggregator is the address to which fees are withdrawn.
	FeeAggregator string
	// RateLimitAdmin is the address allowed to update token pool rate limits.
	RateLimitAdmin string
}

// DeployLombardChainDeps are the dependencies for the DeployCCTPChain sequence.
type DeployLombardChainDeps struct {
	// BlockChains are the chains in the environment.
	BlockChains cldf_chain.BlockChains
	// DataStore defines all addresses in the environment.
	DataStore datastore.DataStore
}

type RemoteLombardChainConfig[LocalContract any, RemoteContract any] struct {
	// TokenPoolConfig configures the token pool for the remote chain.
	TokenPoolConfig tokens.RemoteChainConfig[RemoteContract, LocalContract]
	RemoteDomain    LombardRemoteDomain[RemoteContract]
}

// LombardRemoteDomain identifies Lombard-specific parameters for a remote chain.
type LombardRemoteDomain[RemoteContract any] struct {
	AllowedCaller RemoteContract
	LChainId      uint32
}

// LombardChain is a configurable Lombard chain.
type LombardChain interface {
	// DeployLombardChain deploys the Lombard contracts on the chain.
	DeployLombardChain() *cldf_ops.Sequence[DeployLombardInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	// AddressRefToBytes converts an AddressRef to a byte slice representing the address.
	// Each chain family has their own way of serializing addresses from strings and needs to specify this logic.
	AddressRefToBytes(ref datastore.AddressRef) ([]byte, error)
}

// LombardChainRegistry maintains a registry of Lombard chains.
type LombardChainRegistry struct {
	mu sync.Mutex
	m  map[string]LombardChain
}

// NewLombardChainRegistry creates a new Lombard chain registry.
func NewLombardChainRegistry() *LombardChainRegistry {
	return &LombardChainRegistry{
		m: make(map[string]LombardChain),
	}
}

// RegisterLombardChain allows Lombard chains to register their changeset logic.
func (r *LombardChainRegistry) RegisterLombardChain(chainFamily string, adapter LombardChain) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.m[chainFamily]; exists {
		panic(fmt.Errorf("CCTPChain '%s' already registered", chainFamily))
	}
	r.m[chainFamily] = adapter
}

// GetLombardChain retrieves a registered Lombard chain for the given chain family.
// The boolean return value indicates whether a Lombard chain was found.
func (r *LombardChainRegistry) GetLombardChain(chainFamily string) (LombardChain, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	adapter, ok := r.m[chainFamily]
	return adapter, ok
}
