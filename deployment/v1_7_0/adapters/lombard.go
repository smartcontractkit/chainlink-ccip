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
	// TokenQualifier is the qualifier for matching token with the tokenPool during deployment
	TokenQualifier string
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

// DeployLombardChainDeps are the dependencies for the DeployLombardChain sequence.
type DeployLombardChainDeps struct {
	// BlockChains are the chains in the environment.
	BlockChains cldf_chain.BlockChains
	// DataStore defines all addresses in the environment.
	DataStore datastore.DataStore
}

// ConfigureLombardChainForLanesInput specifies the input for the ConfigureLOmbardChainForLanes sequence.
type ConfigureLombardChainForLanesInput struct {
	// ChainSelector is the selector for the chain being configured.
	ChainSelector uint64
	// Token is the address of the Token contract.
	Token string
	// TokenQualifier is the qualifier for matching token with the tokenPool during deployment
	TokenQualifier string
	// RemoteChains is the set of remote chains to configure.
	RemoteChains map[uint64]RemoteLombardChainConfig
}

// RemoteLombardChainConfig configures a Lombard-enabled chain for a remote counterpart.
type RemoteLombardChainConfig struct {
	TokenTransferFeeConfig tokens.TokenTransferFeeConfig
	LombardChainId         uint32
	FeeUSDCents            uint16
	GasForVerification     uint32
	PayloadSizeBytes       uint32
}

// ConfigureLombardChainForLanesDeps are the dependencies for the ConfigureLombardChainForLanes sequence.
type ConfigureLombardChainForLanesDeps struct {
	// BlockChains are the chains in the environment.
	BlockChains cldf_chain.BlockChains
	// DataStore defines all addresses in the environment.
	DataStore    datastore.DataStore
	RemoteChains map[uint64]RemoteLombardChain
}

// LombardChain is a configurable Lombard chain.
type LombardChain interface {
	RemoteLombardChain
	// DeployLombardChain deploys the Lombard contracts on the chain.
	DeployLombardChain() *cldf_ops.Sequence[DeployLombardInput, sequences.OnChainOutput, DeployLombardChainDeps]
	ConfigureLombardChainForLanes() *cldf_ops.Sequence[ConfigureLombardChainForLanesInput, sequences.OnChainOutput, ConfigureLombardChainForLanesDeps]
	// AddressRefToBytes converts an AddressRef to a byte slice representing the address.
	// Each chain family has their own way of serializing addresses from strings and needs to specify this logic.
	AddressRefToBytes(ref datastore.AddressRef) ([]byte, error)
}

// RemoteLombardChain is a connectable remote Lombard chain.
type RemoteLombardChain interface {
	// AllowedCallerOnDest returns the address allowed to deposit tokens for burn on the remote chain.
	AllowedCallerOnDest(ds datastore.DataStore, chains cldf_chain.BlockChains, selector uint64) ([]byte, error)

	RemoteTokenAddress(bundle cldf_ops.Bundle, ds datastore.DataStore, chains cldf_chain.BlockChains, selector uint64, tokenQualifier string) ([]byte, error)

	RemoteTokenPoolAddress(ds datastore.DataStore, chains cldf_chain.BlockChains, selector uint64, tokenQualifier string) ([]byte, error)
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
		panic(fmt.Errorf("LombardChain '%s' already registered", chainFamily))
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
