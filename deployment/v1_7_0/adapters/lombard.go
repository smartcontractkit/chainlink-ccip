package adapters

import (
	"fmt"
	"sync"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type DeployLombardInput[LocalContract any, RemoteContract any] struct {
	// ChainSelector is the selector for the chain being deployed.
	ChainSelector uint64
	// LombardVerifier is set of addresses comprising the LombardVerifier system.
	LombardVerifier []LocalContract
	Token           string
	TokenPool       LocalContract
	// TokenAdminRegistry is the address of the TokenAdminRegistry contract.
	TokenAdminRegistry LocalContract
	// RMN is the address of the RMN contract.
	RMN LocalContract
	// Router is the address of the Router contract.
	Router LocalContract
	// RemoteChains is the set of remote chains to configure on the CCTPVerifier contract.
	RemoteChains map[uint64]RemoteLombardChainConfig[LocalContract, RemoteContract]
	// DeployerContract is a contract that can be used to deploy other contracts.
	// i.e. A CREATE2Factory contract on Ethereum can enable consistent deployments.
	DeployerContract string
	// StorageLocations is the set of storage locations for the LombardVerifier contract.
	StorageLocations []string
	// FeeAggregator is the address to which fees are withdrawn.
	FeeAggregator string
	// Bridge is the address of the Bridge contract.
	Bridge string
	// RateLimitAdmin is the address allowed to update token pool rate limits.
	RateLimitAdmin string
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

// LombardChain is a configurable CCTP chain.
type LombardChain interface {
	// DeployLombardChain deploys the CCTP contracts on the chain.
	DeployLombardChain() *cldf_ops.Sequence[DeployLombardInput[string, []byte], sequences.OnChainOutput, cldf_chain.BlockChains]
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
