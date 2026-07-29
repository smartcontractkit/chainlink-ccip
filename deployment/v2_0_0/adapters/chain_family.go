package adapters

import (
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
	PayloadSizeBytes          uint16
	SignatureConfig           CommitteeVerifierSignatureQuorumConfig
}

// CommitteeVerifierConfig configures a CommitteeVerifier contract.
type CommitteeVerifierConfig[C any] struct {
	CommitteeVerifier     []C
	RemoteChains          map[uint64]CommitteeVerifierRemoteChainConfig
	AllowedFinalityConfig finality.Config `json:"allowedFinalityConfig" yaml:"allowedFinalityConfig"`
}

// ExecutorDestChainConfig configures the Executor for a remote chain.
type ExecutorDestChainConfig struct {
	USDCentsFee uint16
	Enabled     bool
}

// CommitteeVerifierRemoteChainDefaults provides sensible defaults for CommitteeVerifier
// remote chain configuration fields. Each ChainFamily adapter returns its own defaults;
// callers override individual fields as needed via pointer overrides.
type CommitteeVerifierRemoteChainDefaults struct {
	AllowlistEnabled   bool
	FeeUSDCents        uint16
	GasForVerification uint32
	PayloadSizeBytes   uint16
}

// RemoteChainDefaults provides sensible defaults for remote chain configuration
// fields that are chain-family-specific. Each ChainFamily adapter returns its own
// defaults for a source→remote lane pair; callers override individual fields as needed.
type RemoteChainDefaults struct {
	AllowTrafficFrom          bool
	ExecutorDestChainConfig   ExecutorDestChainConfig
	BaseExecutionGasCost      uint32
	TokenReceiverAllowed      bool
	MessageNetworkFeeUSDCents uint16
	TokenNetworkFeeUSDCents   uint16
	// DefaultExecutor is an optional address override returned by an adapter. When
	// non-empty, it is used directly as the OnRamp default executor instead of
	// resolving a datastore executor ref via ResolveExecutor. This lets adapters
	// supply family-specific executor addresses (e.g. a no-exec executor for Canton).
	DefaultExecutor string
}

// RemoteChainConfig defines the configuration for a remote chain.
type RemoteChainConfig[RemoteContract any, LocalContract any] struct {
	AllowTrafficFrom *bool
	// OnRamps are the remote chain's OnRamp addresses as returned by that chain's
	// GetOnRampAddress: the encoding it writes into the messages it sends.
	OnRamps []RemoteContract
	// OffRamp is the remote chain's OffRamp address in that chain's native
	// encoding: destination-side addresses travel unpadded, so this is the
	// 20-byte address for an EVM remote. GetOffRampAddress returns it directly.
	OffRamp                   RemoteContract
	DefaultInboundCCVs        []LocalContract
	LaneMandatedInboundCCVs   []LocalContract
	DefaultOutboundCCVs       []LocalContract
	LaneMandatedOutboundCCVs  []LocalContract
	DefaultExecutor           LocalContract
	FeeQuoterDestChainConfig  FeeQuoterDestChainConfigOverrides
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
	Router              []byte
	// OnRamp is the local OnRamp as returned by this chain's GetOnRampAddress, so it
	// carries the family's message encoding rather than its plain contract address.
	// On EVM that means 32 abi-encoded bytes; the sequence decodes the address from it.
	OnRamp             []byte
	CommitteeVerifiers []CommitteeVerifierConfig[datastore.AddressRef]
	FeeQuoter          []byte
	OffRamp            []byte
	RemoteChains       map[uint64]RemoteChainConfig[[]byte, string]
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
	// GetOnRampAddress returns the OnRamp address of chainSelector in the encoding that
	// chain writes into the onRampAddress field of the messages it sends. A destination
	// chain whitelists that exact byte string, because the OffRamp identifies the source
	// onramp by hashing the bytes carried in the message. For most families this is the
	// native address; EVM abi-encodes it, so it is 20 bytes left-padded to 32.
	GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	// GetOffRampAddress returns the OffRamp address of chainSelector in that chain's
	// native encoding. Destination-side addresses are never padded, so an EVM OffRamp
	// is 20 bytes.
	GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetTestRouter(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	ResolveExecutor(ds datastore.DataStore, chainSelector uint64, qualifier string) (string, error)
	GetAddressBytesLength() uint8
	GetChainFamilySelector() [4]byte
	GetDefaultFeeQuoterDestChainConfig(chainSelector, remoteChainSelector uint64, chainFamilySelector [4]byte) FeeQuoterDestChainConfigOverrides
	GetDefaultRemoteChainConfig(sourceChainSelector, remoteChainSelector uint64) RemoteChainDefaults
	GetDefaultCommitteeVerifierRemoteChainConfig() CommitteeVerifierRemoteChainDefaults
	GetDefaultFinalityConfig() finality.Config
	ValidateNOPsTopology(chainSelector string, nopCount int) error
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
	if _, exists := r.m[chainFamily]; !exists {
		r.m[chainFamily] = adapter
	}
}

func (r *ChainFamilyRegistry) GetChainFamily(chainFamily string) (ChainFamily, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	adapter, ok := r.m[chainFamily]
	return adapter, ok
}

// FeeQuoterDestChainConfigOverrides is the unified type used both for chain-family adapter
// defaults (returned by GetDefaultFeeQuoterDestChainConfig) and for lane-pair-specific
// user-provided overrides. Nil pointer fields are ignored during merge; non-nil fields
// (including those explicitly set to zero) replace the corresponding value. This ensures
// user-supplied zero values are honored rather than silently dropped.
type FeeQuoterDestChainConfigOverrides struct {
	OverrideExistingConfig      bool
	IsEnabled                   *bool
	MaxDataBytes                *uint32
	MaxPerMsgGasLimit           *uint32
	DestGasOverhead             *uint32
	DestGasPerPayloadByteBase   *uint8
	ChainFamilySelector         [4]byte
	DefaultTokenFeeUSDCents     *uint16
	DefaultTokenDestGasOverhead *uint32
	DefaultTxGasLimit           *uint32
	NetworkFeeUSDCents          *uint16
	LinkFeeMultiplierPercent    *uint8
	USDPerUnitGas               *big.Int
}
