package testadapters

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"

	"github.com/Masterminds/semver/v3"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// Based on unfinished implementation at https://github.com/smartcontractkit/chainlink/blob/f0432dc777d33b83a621da2b042657601d5db8b6/integration-tests/smoke/ccip/canonical/types/types.go

type TokenAmount struct {
	// Will be encoded in the source-native format, so EIP-55 for Ethereum,
	// base58 for Solana, etc.
	Token  string
	Amount *big.Int
}

// MessageComponents is a struct that contains the makeup for a general CCIP message
// irrespective of the chain family it originates from.
type MessageComponents struct {
	DestChainSelector uint64
	// Receiver is the receiver on the destination chain.
	// Must be appropriately dest-chain-family encoded, so abi.encode for Ethereum,
	// 32 bytes for Solana, etc.
	Receiver []byte
	// Data is the data to be sent to the destination chain.
	Data []byte
	// Will be encoded in the source-native format, so EIP-55 for Ethereum,
	// base58 for Solana, etc.
	FeeToken string
	// ExtraArgs are the message extra args which tune message semantics and behavior.
	// For example, out of order execution can be specified here.
	ExtraArgs []byte
	// TokenAmounts are the tokens and their respective amounts to be sent to the
	// destination chain.
	// Note that the tokens must be "approved" to the router for the message send to work.
	TokenAmounts []TokenAmount
}

const ExtraArgGasLimit = "gasLimit|computeUnits"
const ExtraArgOOO = "outOfOrderExecutionEnabled"

// ExtraArgOpt is a generic representation of an extra arg that can be applied
// to any kind of ccip message.
// We use this to make it possible to specify extra args in a chain-agnostic way.
type ExtraArgOpt struct {
	Name  string
	Value any
}

func NewOutOfOrderExtraArg(outOfOrder bool) ExtraArgOpt {
	return ExtraArgOpt{
		Name:  ExtraArgOOO,
		Value: outOfOrder,
	}
}

func NewGasLimitExtraArg(gasLimit *big.Int) ExtraArgOpt {
	return ExtraArgOpt{
		Name:  ExtraArgGasLimit,
		Value: gasLimit,
	}
}

// TestAdapter is our interface for interacting with a specific chain, scoped to a family.
// An adapter instance is an instance of a concrete chain.
// So if there are e.g 3 source chains that are EVM and a dest that is Solana,
// we would have 3 EVM adapters and 1 Solana adapter.
type TestAdapter interface {
	// ChainSelector returns the selector of the chain for the given adapter.
	ChainSelector() uint64

	// ChainFamily returns the family of the chain for the given adapter.
	Family() string

	// BuildMessage builds a message from the given components,
	// with the overall message type being ChainFamily2Any, where
	// ChainFamily is the family of the adapter.
	// As a concrete example, for EVM, the message type is router.ClientEVM2AnyMessage,
	// and for Solana, the message type is ccip_router.SVM2AnyMessage.
	BuildMessage(components MessageComponents) (any, error)

	SendMessage(ctx context.Context, destChainSelector uint64, msg any) (uint64, error)

	// // RandomReceiver returns a random receiver for the given chain family.
	// RandomReceiver() []byte

	// // CCIPReceiver returns a CCIP receiver for the given chain family.
	CCIPReceiver() []byte

	// SetReceiverRejectAll configures the receiver to reject all incoming messages.
	// This is used for test cases with a a failing receiver.
	SetReceiverRejectAll(ctx context.Context, rejectAll bool) error

	// NativeFeeToken returns the native fee token for the given chain family.
	NativeFeeToken() string

	// GetExtraArgs returns the default extra args for sending messages to this
	// chain family from the given source family.
	// Therefore the extra args are source-family encoded, so abi.encode for EVM,
	// borsch for Solana, etc.
	GetExtraArgs(receiver []byte, sourceFamily string, opts ...ExtraArgOpt) ([]byte, error)

	// GetInboundNonce returns the inbound nonce for the given sender and source chain selector.
	// For chains that don't have the concept of nonces, this will always return 0.
	GetInboundNonce(ctx context.Context, sender []byte, srcSel uint64) (uint64, error)

	// ValidateCommit validates that the message specified by the given send event was committed.
	ValidateCommit(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNumRange ccipocr3.SeqNumRange)

	// ValidateExec validates that the message specified by the given send event was executed.
	ValidateExec(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNrs []uint64) (execStates map[uint64]int)

	// AllowRouterToWithdrawTokens allows the router to withdraw tokens of the given address and amount from the deployer account.
	AllowRouterToWithdrawTokens(ctx context.Context, tokenAddress string, amount *big.Int) error

	// GetTokenBalance gets the token balance of the given owner address for the given token address.
	GetTokenBalance(ctx context.Context, tokenAddress string, ownerAddress []byte) (*big.Int, error)

	// GetTokenExpansionConfig returns a token expansion deployment config with sensible defaults for testing cross-chain token transfers.
	GetTokenExpansionConfig() tokensapi.TokenExpansionInputPerChain

	// GetRegistryAddress returns the address of the contract on which the token pool must be registered.
	GetRegistryAddress() (string, error)

	// SetAllowlist activates/deactivates the whitelist
	SetAllowlist(ctx context.Context, destChainSelector uint64, enabled bool) (err error, cleanup func() /*err?*/)

	// UpdateSenderAllowlistStatus adds/removes senders to/from the whitelist
	UpdateSenderAllowlistStatus(ctx context.Context, destChainSelector uint64, included bool) (err error, cleanup func() /*err?*/)
}

type TestAdapterFactory = func(env *deployment.Environment, selector uint64) TestAdapter

type testAdapterID string

// TestAdapterRegistry maintains a registry of TestAdapters.
type TestAdapterRegistry struct {
	mu sync.Mutex
	m  map[testAdapterID]TestAdapterFactory
}

// NewTestAdapterRegistry creates a fresh registry.  It is kept unexported
// because callers should obtain the singleton via GetTestAdapterRegistry().
func newTestAdapterRegistry() *TestAdapterRegistry {
	return &TestAdapterRegistry{
		m: make(map[testAdapterID]TestAdapterFactory),
	}
}

// RegisterTestAdapter registers a new adapter; panics if the key already exists.
func (r *TestAdapterRegistry) RegisterTestAdapter(chainFamily string, version *semver.Version, adapter TestAdapterFactory) {
	id := newTestAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; exists {
		panic(fmt.Errorf("TestAdapter '%s %s' already registered", chainFamily, version))
	}
	r.m[id] = adapter
}

// GetTestAdapter looks up an adapter; the second return value tells you if it was found.
func (r *TestAdapterRegistry) GetTestAdapter(chainFamily string, version *semver.Version) (TestAdapterFactory, bool) {
	id := newTestAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	adapter, ok := r.m[id]
	return adapter, ok
}

var (
	singletonRegistry *TestAdapterRegistry
	once              sync.Once
)

// GetTestAdapterRegistry returns the global singleton instance.
// The first call creates the registry; subsequent calls return the same pointer.
func GetTestAdapterRegistry() *TestAdapterRegistry {
	once.Do(func() {
		singletonRegistry = newTestAdapterRegistry()
	})
	return singletonRegistry
}

func newTestAdapterID(chainFamily string, version *semver.Version) testAdapterID {
	return testAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}

// TODO: remove once migration to DataStore is completed and stateview is obsolete
type StateProvider interface {
	GetAddress(datastore.ContractType) (string, error)
}

type DataStoreStateProvider struct {
	Selector uint64
	DS       datastore.DataStore
}

func (p *DataStoreStateProvider) GetAddress(ty datastore.ContractType) (string, error) {
	addr, err := datastore_utils.FindAndFormatRef(p.DS, datastore.AddressRef{
		ChainSelector: p.Selector,
		Type:          ty,
		// TODO: version
	}, p.Selector, datastore_utils.FullRef)
	if err != nil {
		return "", err
	}
	return addr.Address, nil
}

type CommitReportTracker struct {
	seenMessages map[uint64]map[uint64]bool
}

func NewCommitReportTracker(sourceChainSelector uint64, seqNrs ccipocr3.SeqNumRange) CommitReportTracker {
	seenMessages := make(map[uint64]map[uint64]bool)
	seenMessages[sourceChainSelector] = make(map[uint64]bool)

	for i := seqNrs.Start(); i <= seqNrs.End(); i++ {
		seenMessages[sourceChainSelector][uint64(i)] = false
	}
	return CommitReportTracker{seenMessages: seenMessages}
}

func (c *CommitReportTracker) VisitCommitReport(sourceChainSelector uint64, minSeqNr uint64, maxSeqNr uint64) {
	if _, ok := c.seenMessages[sourceChainSelector]; !ok {
		return
	}

	for i := minSeqNr; i <= maxSeqNr; i++ {
		c.seenMessages[sourceChainSelector][i] = true
	}
}

func (c *CommitReportTracker) AllCommitted(sourceChainSelector uint64) bool {
	for _, v := range c.seenMessages[sourceChainSelector] {
		if !v {
			return false
		}
	}
	return true
}
