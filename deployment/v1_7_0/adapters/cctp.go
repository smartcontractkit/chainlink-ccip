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

// RemoteCCTPChainConfig configures a CCTP-enabled chain for a remote counterpart.
type RemoteCCTPChainConfig struct {
	// FeeUSDCCents is the flat fee, in multiples of 0.01 USD cents, charged for verification on the remote chain.
	FeeUSDCents uint16
	// GasForVerification is the gas required to verify the CCTP message on the remote chain.
	GasForVerification uint32
	// PayloadSizeBytes is the size of the CCTP verification payload to be checked on the remote chain.
	PayloadSizeBytes uint32
	// LockOrBurnMechanism specifies the mechanism by which the CCTP message will be handled.
	// Each chain family may interpret this string differently.
	LockOrBurnMechanism string
	// DomainIdentifier is the identifier of the remote domain.
	DomainIdentifier uint32
	// TokenTransferFeeConfig specifies the desired token transfer fee configuration for this remote chain.
	TokenTransferFeeConfig tokens.TokenTransferFeeConfig
}

// ConfigureCCTPChainForLanesInput specifies the input for the ConfigureCCTPChainForLanes sequence.
type ConfigureCCTPChainForLanesInput struct {
	// ChainSelector is the selector for the chain being configured.
	ChainSelector uint64
	// USDCToken is the address of the USDCToken contract.
	USDCToken string
	// RegisteredPoolRef is a reference to the pool that should be set on the registry on this chain.
	RegisteredPoolRef datastore.AddressRef
	// RemoteRegisteredPoolRefs is a map of remote chain selectors to references to the pool that should be set on the registry on the remote chain.
	RemoteRegisteredPoolRefs map[uint64]datastore.AddressRef
	// RemoteChains is the set of remote chains to configure.
	RemoteChains map[uint64]RemoteCCTPChainConfig
}

// DeployCCTPInput specifies the input for the DeployCCTPChain sequence.
type DeployCCTPInput struct {
	// ChainSelector is the selector for the chain being deployed.
	ChainSelector uint64
	// TokenMessengerV1 is the address of the CCTP v1 TokenMessenger contract.
	// Optional. If empty, CCTP V1 pool deployment/configuration is skipped.
	TokenMessengerV1 string
	// TokenMessengerV2 is the address of the CCTP v2 TokenMessenger contract.
	TokenMessengerV2 string
	// USDCToken is the address of the USDCToken contract.
	USDCToken string
	// DeployerContract is a contract that can be used to deploy other contracts.
	// i.e. A CREATE2Factory contract on Ethereum can enable consistent deployments.
	DeployerContract string
	// FastFinalityBps are the basis points charged for fast finality.
	FastFinalityBps uint16
	// StorageLocations is the set of storage locations for the CCTPVerifier contract.
	StorageLocations []string
	// FeeAggregator is the address to which fees are withdrawn.
	FeeAggregator string
}

// DeployCCTPChainDeps are the dependencies for the DeployCCTPChain sequence.
type DeployCCTPChainDeps struct {
	// BlockChains are the chains in the environment.
	BlockChains cldf_chain.BlockChains
	// DataStore defines all addresses in the environment.
	DataStore datastore.DataStore
}

// ConfigureCCTPChainForLanesDeps are the dependencies for the ConfigureCCTPChainForLanes sequence.
type ConfigureCCTPChainForLanesDeps struct {
	// BlockChains are the chains in the environment.
	BlockChains cldf_chain.BlockChains
	// DataStore defines all addresses in the environment.
	DataStore datastore.DataStore
	// RemoteChains are the remote chains in the environment.
	RemoteChains map[uint64]RemoteCCTPChain
}

// RemoteCCTPChain is a connectable remote CCTP chain.
type RemoteCCTPChain interface {
	// PoolAddress returns the address of the token pool on the remote chain in bytes.
	// The ref is used in combination with the chain selector to query the datastore for the registered pool address.
	PoolAddress(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64, registeredPoolRef datastore.AddressRef) ([]byte, error)
	// TokenAddress returns the address of the token on the remote chain in bytes.
	TokenAddress(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error)
	// AllowedCallerOnDest returns the address allowed to trigger message reception on the remote domain.
	AllowedCallerOnDest(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error)
	// AllowedCallerOnSource returns the address allowed to deposit tokens for burn on the remote chain.
	AllowedCallerOnSource(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error)
	// MintRecipientOnDest returns the address that will receive tokens on the remote domain.
	// If not empty, the tokens will be minted to the message receiver.
	MintRecipientOnDest(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error)
}

// CCTPChain is a configurable CCTP chain.
type CCTPChain interface {
	RemoteCCTPChain
	// DeployCCTPChain deploys the CCTP contracts on the chain.
	DeployCCTPChain() *cldf_ops.Sequence[DeployCCTPInput, sequences.OnChainOutput, DeployCCTPChainDeps]
	// ConfigureCCTPChainForLanes configures the CCTP contracts on the chain for lanes.
	ConfigureCCTPChainForLanes() *cldf_ops.Sequence[ConfigureCCTPChainForLanesInput, sequences.OnChainOutput, ConfigureCCTPChainForLanesDeps]
}

// CCTPChainRegistry maintains a registry of CCTP chains.
type CCTPChainRegistry struct {
	mu sync.Mutex
	m  map[string]CCTPChain
}

// NewCCTPChainRegistry creates a new CCTP chain registry.
func NewCCTPChainRegistry() *CCTPChainRegistry {
	return &CCTPChainRegistry{
		m: make(map[string]CCTPChain),
	}
}

// RegisterCCTPChain allows CCTP chains to register their changeset logic.
func (r *CCTPChainRegistry) RegisterCCTPChain(chainFamily string, adapter CCTPChain) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.m[chainFamily]; exists {
		panic(fmt.Errorf("CCTPChain '%s' already registered", chainFamily))
	}
	r.m[chainFamily] = adapter
}

// GetCCTPChain retrieves a registered CCTP chain for the given chain family.
// The boolean return value indicates whether a CCTP chain was found.
func (r *CCTPChainRegistry) GetCCTPChain(chainFamily string) (CCTPChain, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	adapter, ok := r.m[chainFamily]
	return adapter, ok
}
