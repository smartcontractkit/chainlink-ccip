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

// Mechanism specifies the mechanism by which the CCTP message will be handled.
type Mechanism string

const (
	CCTPV1Mechanism        Mechanism = "CCTP_V1"
	CCTPV2Mechanism        Mechanism = "CCTP_V2"
	LockReleaseMechanism   Mechanism = "LOCK_RELEASE"
	CCTPV2WithCCVMechanism Mechanism = "CCTP_V2_WITH_CCV"
)

// IsValid checks if the mechanism is valid.
func (m Mechanism) IsValid() bool {
	switch m {
	case CCTPV1Mechanism, CCTPV2Mechanism, LockReleaseMechanism, CCTPV2WithCCVMechanism:
		return true
	default:
		return false
	}
}

// RemoteDomain identifies CCTP-specific parameters for a remote chain.
type RemoteDomain[Contract any] struct {
	// AllowedCallerOnDest is the address allowed to trigger message reception on the remote domain.
	AllowedCallerOnDest Contract
	// AllowedCallerOnSource is the address expected to deposit tokens for burn on the remote chain.
	AllowedCallerOnSource Contract
	// MintRecipientOnDest is the address that will receive tokens on the remote domain.
	// If not set, the tokens will be minted to the receiver of the CCIP message.
	MintRecipientOnDest Contract
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
	// i.e. CCTP V1, CCTP V2, lock-release, or CCTP V2 with CCVs.
	LockOrBurnMechanism Mechanism
	// LockRelease pool is the address of the lock-release pool used for funds sent to / received from the remote chain.
	// Only required if the lock-release mechanism is used.
	LockReleasePool LocalContract
	// RemoteDomain configures the CCTP-specific parameters for the remote chain.
	RemoteDomain RemoteDomain[RemoteContract]
}

// TokenPools specifies all possible CCTP token pools that may exist on a given chain.
type TokenPools[Contract any] struct {
	// LegacyCCTPV1Pool is the address of the legacy CCTP V1 token pool.
	LegacyCCTPV1Pool Contract
	// CCTPV1Pool is the address of the CCTP V1 token pool.
	CCTPV1Pool Contract
	// CCTPV2Pool is the address of the CCTP V2 token pool.
	CCTPV2Pool Contract
	// CCTPV2PoolWithCCV is the address of the CCTP V2 token pool that uses CCVs.
	CCTPV2PoolWithCCV Contract
}

// DeployCCTPInput specifies the input for the DeployCCTPChain sequence.
type DeployCCTPInput[LocalContract any, RemoteContract any] struct {
	// ChainSelector is the selector for the chain being deployed.
	ChainSelector uint64
	// TokenPools specifies all possible CCTP token pools that may exist on a given chain.
	TokenPools TokenPools[LocalContract]
	// USDCTokenPoolProxy is the address of the USDCTokenPoolProxy contract.
	USDCTokenPoolProxy LocalContract
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
