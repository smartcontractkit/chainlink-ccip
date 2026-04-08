package adapters

import (
	"fmt"
	"math/big"
	"sync"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// CommitteeVerifierSignatureQuorumConfig specifies the quorum required for any given message.
type CommitteeVerifierSignatureQuorumConfig struct {
	Signers   []string
	Threshold uint8
}

// CommitteeVerifierRemoteChainConfig configures the CommitteeVerifier for a remote chain.
type CommitteeVerifierRemoteChainConfig struct {
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []string
	RemovedAllowlistedSenders []string
	FeeUSDCents               uint16
	GasForVerification        uint32
	PayloadSizeBytes          uint32
	SignatureConfig           CommitteeVerifierSignatureQuorumConfig
}

// CommitteeVerifierConfig configures a CommitteeVerifier contract.
type CommitteeVerifierConfig[C any] struct {
	CommitteeVerifier      []C
	RemoteChains           map[uint64]CommitteeVerifierRemoteChainConfig
	AllowedFinalityConfig  finality.Config `json:"allowedFinalityConfig" yaml:"allowedFinalityConfig"`
}

// ExecutorDestChainConfig configures the Executor for a remote chain.
type ExecutorDestChainConfig struct {
	USDCentsFee uint16
	Enabled     bool
}

// FeeQuoterDestChainConfig configures the FeeQuoter for a remote chain.
type FeeQuoterDestChainConfig struct {
	OverrideExistingConfig      bool
	IsEnabled                   bool
	MaxDataBytes                uint32
	MaxPerMsgGasLimit           uint32
	DestGasOverhead             uint32
	DestGasPerPayloadByteBase   uint8
	ChainFamilySelector         [4]byte
	DefaultTokenFeeUSDCents     uint16
	DefaultTokenDestGasOverhead uint32
	DefaultTxGasLimit           uint32
	NetworkFeeUSDCents          uint16
	LinkFeeMultiplierPercent    uint8
	USDPerUnitGas               *big.Int
}

// RemoteChainConfig defines the configuration for a remote chain.
type RemoteChainConfig[RemoteContract any, LocalContract any] struct {
	AllowTrafficFrom          *bool
	OnRamps                   []RemoteContract
	OffRamp                   RemoteContract
	DefaultInboundCCVs        []LocalContract
	LaneMandatedInboundCCVs   []LocalContract
	DefaultOutboundCCVs       []LocalContract
	LaneMandatedOutboundCCVs  []LocalContract
	DefaultExecutor           LocalContract
	FeeQuoterDestChainConfig  FeeQuoterDestChainConfig
	ExecutorDestChainConfig   ExecutorDestChainConfig
	AddressBytesLength        uint8
	BaseExecutionGasCost      uint32
	TokenReceiverAllowed      *bool
	MessageNetworkFeeUSDCents uint16
	TokenNetworkFeeUSDCents   uint16
}

// ConfigureChainForLanesInput is the input for the chain-centric lane configuration sequence.
type ConfigureChainForLanesInput struct {
	ChainSelector       uint64
	AllowOnrampOverride bool
	Router              string
	OnRamp              string
	CommitteeVerifiers  []CommitteeVerifierConfig[datastore.AddressRef]
	FeeQuoter           string
	OffRamp             string
	RemoteChains        map[uint64]RemoteChainConfig[[]byte, string]
	// FamilyExtras holds chain-family-specific configuration passed through
	// from the changeset. Each family adapter's sequence is responsible for
	// interpreting this map. All values must be serializable.
	FamilyExtras map[string]any
}

// ChainFamily is a configurable chain family for chain-centric lane setup.
// It provides both the lane configuration sequence and contract resolution
// methods so that callers don't need to construct datastore.AddressRef manually
// for well-known contract types (OnRamp, OffRamp, FeeQuoter, Router, Executor).
type ChainFamily interface {
	ConfigureChainForLanes() *cldf_ops.Sequence[ConfigureChainForLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	AddressRefToBytes(ref datastore.AddressRef) ([]byte, error)
	GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetTestRouter(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	ResolveExecutor(ds datastore.DataStore, chainSelector uint64, qualifier string) (string, error)
}

// ChainFamilyRegistry maintains a registry of chain families.
type ChainFamilyRegistry struct {
	mu sync.Mutex
	m  map[string]ChainFamily
}

var (
	singletonChainFamilyRegistry *ChainFamilyRegistry
	chainFamilyRegistryOnce      sync.Once
)

func NewChainFamilyRegistry() *ChainFamilyRegistry {
	return &ChainFamilyRegistry{m: make(map[string]ChainFamily)}
}

func GetChainFamilyRegistry() *ChainFamilyRegistry {
	chainFamilyRegistryOnce.Do(func() {
		singletonChainFamilyRegistry = NewChainFamilyRegistry()
	})
	return singletonChainFamilyRegistry
}

func (r *ChainFamilyRegistry) RegisterChainFamily(chainFamily string, adapter ChainFamily) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.m[chainFamily]; exists {
		panic(fmt.Errorf("ChainFamily '%s' already registered", chainFamily))
	}
	r.m[chainFamily] = adapter
}

func (r *ChainFamilyRegistry) GetChainFamily(chainFamily string) (ChainFamily, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	adapter, ok := r.m[chainFamily]
	return adapter, ok
}
