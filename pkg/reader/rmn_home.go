package reader

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	rmnMaxSizeCommittee = 256 // bitmap is 256 bits making the max committee size 256
	maxFailedPolls      = uint32(10)
)

type HomeNodeInfo = rmntypes.HomeNodeInfo

type NodeID = rmntypes.NodeID

type RMNHome interface {
	// GetRMNNodesInfo gets the RMNHomeNodeInfo for the given configDigest
	GetRMNNodesInfo(configDigest cciptypes.Bytes32) ([]rmntypes.HomeNodeInfo, error)
	// IsRMNHomeConfigDigestSet checks if the configDigest is set in the RMNHome contract
	IsRMNHomeConfigDigestSet(configDigest cciptypes.Bytes32) bool
	// GetRMNEnabledSourceChains gets the RMN-enabled source chains for the given configDigest.
	// If a chain is not RMN-enabled it means that we don't need to do RMN signature related operations for that chain.
	GetRMNEnabledSourceChains(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]bool, error)
	// GetFObserve gets the F value for each source chain in the given configDigest.
	// Maximum number of faulty observers; F+1 observers required to agree on an observation for a source chain.
	GetFObserve(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]int, error)
	// GetOffChainConfig gets the offchain config for the given configDigest
	GetOffChainConfig(configDigest cciptypes.Bytes32) (cciptypes.Bytes, error)
	// GetAllConfigDigests gets the active and candidate RMNHomeConfigs
	GetAllConfigDigests() (activeConfigDigest cciptypes.Bytes32, candidateConfigDigest cciptypes.Bytes32)
	services.Service
}

type rmnHome struct {
	sync     services.StateMachine
	bgPoller *rmnHomePoller
}

// NewRMNHomeChainReader creates a new RMNHome. RMNHome is very lightweight layer
// on top of RMNHomePoller. Every consumer should follow the `Service` pattern in which they
// are responsible for properly starting and closing the service when done. (using Start()/Close() methods)
//
// RMNHome is a smart component that behind the scenes share the same RMNHomePoller instance.
// Whenever all RMNHome instances are closed, the underlying RMNHomePoller instance is also closed.
// As long as RMNHome remembers to Close() upon termination there should not be any orphaned poller instances
// working in the background
func NewRMNHomeChainReader(
	ctx context.Context,
	lggr logger.Logger,
	pollingInterval time.Duration,
	rmnHomeChainSelector cciptypes.ChainSelector,
	rmnHomeAddress []byte,
	contractReader contractreader.ContractReaderFacade,
) (RMNHome, error) {
	bgPoller, err := getRMNHomePoller(
		ctx,
		lggr,
		rmnHomeChainSelector,
		rmnHomeAddress,
		contractReader,
		pollingInterval,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RMNHomePoller: %w", err)
	}

	return &rmnHome{bgPoller: bgPoller}, nil
}

func (r *rmnHome) Start(ctx context.Context) error {
	return r.sync.StartOnce(r.Name(), func() error {
		return r.bgPoller.Start(ctx, r)
	})
}

func (r *rmnHome) Close() error {
	err := r.sync.StopOnce(r.Name(), func() error {
		return r.bgPoller.Close(r)
	})

	if errors.Is(err, services.ErrAlreadyStopped) {
		return nil
	}

	return err
}

func (r *rmnHome) GetRMNNodesInfo(configDigest cciptypes.Bytes32) ([]rmntypes.HomeNodeInfo, error) {
	state := r.bgPoller.getRMNHomeState()
	_, ok := state.rmnHomeConfig[configDigest]
	if !ok {
		return nil, fmt.Errorf("configDigest %s not found in RMNHomeConfig", configDigest)

	}
	return state.rmnHomeConfig[configDigest].Nodes, nil
}

func (r *rmnHome) IsRMNHomeConfigDigestSet(configDigest cciptypes.Bytes32) bool {
	_, ok := r.bgPoller.getRMNHomeState().rmnHomeConfig[configDigest]
	return ok
}

func (r *rmnHome) GetFObserve(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]int, error) {
	state := r.bgPoller.getRMNHomeState()
	_, ok := state.rmnHomeConfig[configDigest]
	if !ok {
		return nil, fmt.Errorf("configDigest %s not found in RMNHomeConfig", configDigest)
	}
	return state.rmnHomeConfig[configDigest].SourceChainF, nil
}

// GetRMNEnabledSourceChains returns the source chains that are RMN-enabled. A chain is considered RMN-enabled if
// F is present in RMNHome config.
func (r *rmnHome) GetRMNEnabledSourceChains(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]bool, error) {
	state := r.bgPoller.getRMNHomeState()
	homeCfg, ok := state.rmnHomeConfig[configDigest]
	if !ok {
		return map[cciptypes.ChainSelector]bool{},
			fmt.Errorf("configDigest %s not found in RMNHomeConfig", configDigest)
	}

	enabledChains := make(map[cciptypes.ChainSelector]bool, len(homeCfg.SourceChainF))
	for chain := range homeCfg.SourceChainF {
		enabledChains[chain] = true
	}

	return enabledChains, nil
}

func (r *rmnHome) GetOffChainConfig(configDigest cciptypes.Bytes32) (cciptypes.Bytes, error) {
	state := r.bgPoller.getRMNHomeState()
	cfg, ok := state.rmnHomeConfig[configDigest]
	if !ok {
		return nil, fmt.Errorf("configDigest %s not found in RMNHomeConfig", configDigest)
	}
	return cfg.OffchainConfig, nil
}

func (r *rmnHome) GetAllConfigDigests() (
	activeConfigDigest cciptypes.Bytes32,
	candidateConfigDigest cciptypes.Bytes32,
) {
	state := r.bgPoller.getRMNHomeState()
	return state.activeConfigDigest, state.candidateConfigDigest
}

func (r *rmnHome) Ready() error {
	return r.sync.Ready()
}

func (r *rmnHome) HealthReport() map[string]error {
	return r.bgPoller.HealthReport()
}

func (r *rmnHome) Name() string {
	return "RMNHome"
}

// GetAllConfigsResponse mirrors RMNHome.sol's getAllConfigs() return value.
type GetAllConfigsResponse struct {
	ActiveConfig    VersionedConfig `json:"activeConfig"`
	CandidateConfig VersionedConfig `json:"candidateConfig"`
}

// VersionedConfig mirrors RMNHome.sol's VersionedConfig struct
type VersionedConfig struct {
	Version       uint32            `json:"version"`
	ConfigDigest  cciptypes.Bytes32 `json:"configDigest"`
	StaticConfig  StaticConfig      `json:"staticConfig"`
	DynamicConfig DynamicConfig     `json:"dynamicConfig"`
}

// StaticConfig mirrors RMNHome.sol's StaticConfig struct
type StaticConfig struct {
	Nodes          []Node          `json:"nodes"`
	OffchainConfig cciptypes.Bytes `json:"offchainConfig"`
}

type DynamicConfig struct {
	SourceChains   []SourceChain   `json:"sourceChains"`
	OffchainConfig cciptypes.Bytes `json:"offchainConfig"`
}

// Node mirrors RMNHome.sol's Node struct
type Node struct {
	PeerID            cciptypes.Bytes32 `json:"peerId"`
	OffchainPublicKey cciptypes.Bytes32 `json:"offchainPublicKey"`
}

// SourceChain mirrors RMNHome.sol's SourceChain struct
type SourceChain struct {
	ChainSelector       cciptypes.ChainSelector `json:"chainSelector"`
	FObserve            uint64                  `json:"fObserve"` // previously: MinObservers / F
	ObserverNodesBitmap *big.Int                `json:"observerNodesBitmap"`
}

var _ RMNHome = (*rmnHome)(nil)
