package adapters

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// RemoteDomain identifies CCTP-specific parameters for a remote chain.
type RemoteDomain[RemoteContract any] struct {
	// AllowedCallerOnDest is the address allowed to trigger message reception on the remote domain.
	AllowedCallerOnDest RemoteContract
	// AllowedCallerOnSource is the address expected to deposit tokens for burn on the remote chain.
	AllowedCallerOnSource RemoteContract
	// MintRecipientOnDest is the address that will receive tokens on the remote domain.
	// If not set, the tokens will be minted to the receiver of the CCIP message.
	MintRecipientOnDest RemoteContract
	// DomainIdentifier is the identifier of the remote domain.
	DomainIdentifier uint32
}

// RemoteChainConfig configures a CCTP-enabled chain for a remote counterpart.
type RemoteCCTPChainConfig[LocalContract any, RemoteContract any] struct {
	// TokenPoolConfig configures the token pool for the remote chain.
	TokenPoolConfig tokens.RemoteChainConfig[RemoteContract, LocalContract]
	// FeeUSDCCents is the flat fee, in multiples of 0.01 USD cents, charged for verification on the remote chain.
	FeeUSDCents uint16
	// GasForVerification is the gas required to verify the CCTP message on the remote chain.
	GasForVerification uint32
	// PayloadSizeBytes is the size of the CCTP verification payload to be checked on the remote chain.
	PayloadSizeBytes uint32
	// LockOrBurnMechanism specifies the mechanism by which the CCTP message will be handled.
	// Each chain family may interpret this string differently.
	LockOrBurnMechanism string
	// RemoteDomain configures the CCTP-specific parameters for the remote chain.
	RemoteDomain RemoteDomain[RemoteContract]
}

// DeployCCTPInput specifies the input for the DeployCCTPChain sequence.
type DeployCCTPInput[LocalContract any, RemoteContract any] struct {
	// ChainSelector is the selector for the chain being deployed.
	ChainSelector uint64
	// TokenPool is the set of all contracts that comprise the token pool on this chain.
	// i.e. Proxy and various implementations.
	TokenPool []LocalContract
	// CCTPVerifier is set of addresses comprising the CCTPVerifier system.
	CCTPVerifier []LocalContract
	// MessageTransmitterProxy is the address of the MessageTransmitterProxy contract.
	MessageTransmitterProxy LocalContract
	// TokenAdminRegistry is the address of the TokenAdminRegistry contract.
	TokenAdminRegistry LocalContract
	// RMN is the address of the RMN contract.
	RMN LocalContract
	// Router is the address of the Router contract.
	Router LocalContract
	// RemoteChains is the set of remote chains to configure on the CCTPVerifier contract.
	RemoteChains map[uint64]RemoteCCTPChainConfig[LocalContract, RemoteContract]
	// TokenMessenger is the address of the TokenMessenger contract.
	TokenMessenger string
	// USDCToken is the address of the USDCToken contract.
	USDCToken string
	// DeployerContract is a contract that can be used to deploy other contracts.
	// i.e. A CREATE2Factory contract on Ethereum can enable consistent deployments.
	DeployerContract string
	// MinFinalityValue is the minimum finality value required by the token pool.
	MinFinalityValue uint16
	// ThresholdAmountForAdditionalCCVs is the threshold amount above which additional CCVs are required.
	ThresholdAmountForAdditionalCCVs *big.Int
	// FastFinalityBps are the basis points charged for fast finality.
	FastFinalityBps uint16
	// StorageLocations is the set of storage locations for the CCTPVerifier contract.
	StorageLocations []string
	// Allowlist is the set of addresses allowed to transfer the token.
	Allowlist []string
	// RateLimitAdmin is the address allowed to update token pool rate limits.
	RateLimitAdmin string
	// FeeAggregator is the address to which fees are withdrawn.
	FeeAggregator string
	// AllowlistAdmin is address allowed to update the token pool allowlist.
	AllowlistAdmin string
}

// CCTPChain is a configurable CCTP chain.
type CCTPChain interface {
	// DeployCCTPChain deploys the CCTP contracts on the chain.
	DeployCCTPChain() *cldf_ops.Sequence[DeployCCTPInput[string, []byte], sequences.OnChainOutput, cldf_chain.BlockChains]
	// AddressRefToBytes converts an AddressRef to a byte slice representing the address.
	// Each chain family has their own way of serializing addresses from strings and needs to specify this logic.
	AddressRefToBytes(ref datastore.AddressRef) ([]byte, error)
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
