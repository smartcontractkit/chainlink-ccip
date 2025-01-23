package reader

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	rmnMaxSizeCommittee = 256 // bitmap is 256 bits making the max committee size 256
	MaxFailedPolls      = uint(10)
)

type HomeNodeInfo = rmntypes.HomeNodeInfo

type NodeID = rmntypes.NodeID

type RMNHome interface {
	// GetRMNNodesInfo gets the RMNHomeNodeInfo for the given configDigest
	GetRMNNodesInfo(configDigest cciptypes.Bytes32) ([]rmntypes.HomeNodeInfo, error)
	// IsRMNHomeConfigDigestSet checks if the configDigest is set in the RMNHome contract
	IsRMNHomeConfigDigestSet(configDigest cciptypes.Bytes32) bool
	// GetFObserve gets the F value for each source chain in the given configDigest.
	// Maximum number of faulty observers; F+1 observers required to agree on an observation for a source chain.
	GetFObserve(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]int, error)
	// GetOffChainConfig gets the offchain config for the given configDigest
	GetOffChainConfig(configDigest cciptypes.Bytes32) (cciptypes.Bytes, error)
	// GetAllConfigDigests gets the active and candidate RMNHomeConfigs
	GetAllConfigDigests() (activeConfigDigest cciptypes.Bytes32, candidateConfigDigest cciptypes.Bytes32)
	services.Service
}

var (
	instances   = make(map[string]*rmnHomePoller)
	instancesMu sync.Mutex
)

type rmnHomeState struct {
	activeConfigDigest    cciptypes.Bytes32
	candidateConfigDigest cciptypes.Bytes32
	rmnHomeConfig         map[cciptypes.Bytes32]rmntypes.HomeConfig
}

// rmnHomePoller polls the RMNHome contract for the latest RMNHomeConfigs
// It is running in the background with a polling interval of pollingDuration
type rmnHomePoller struct {
	wg                   sync.WaitGroup
	stopCh               services.StopChan
	sync                 services.StateMachine
	mutex                *sync.RWMutex
	contractReader       contractreader.ContractReaderFacade
	rmnHomeBoundContract types.BoundContract
	lggr                 logger.Logger
	rmnHomeState         rmnHomeState
	failedPolls          uint
	pollingDuration      time.Duration // How frequently the poller fetches the chain configs
}

// GetRMNHomePoller returns a rmnHomePoller instance if it already exists, else creates a new one.
// Returned rmnHomePoller is always started, therefore user of that API doesn't need to care about
// concurrency and accidentally starting the same poller multiple times.
//
// In most of the cases, there is going to be only a single RMNHome deployed on a single chain. However,
// OCR3Config allows to pass different RMNHome addresses for different chains. Therefore, we need to
// support that and maintain a separate rmnHomePoller for each RMNHome address.
// Having a singleton here is aimed to reduce the background polling to the underlying RPC node.
func GetRMNHomePoller(
	ctx context.Context,
	lggr logger.Logger,
	rmnHomeChainSelector cciptypes.ChainSelector,
	rmnHomeAddress []byte,
	contractReader contractreader.ContractReaderFacade,
	pollingInterval time.Duration,
) (RMNHome, error) {
	instancesMu.Lock()
	defer instancesMu.Unlock()

	hexEncodedAddr := " 0x" + hex.EncodeToString(rmnHomeAddress)
	key := fmt.Sprintf("%s-%s", rmnHomeChainSelector.String(), hexEncodedAddr)

	instance, ok := instances[key]
	if ok {
		lggr.Infow("RMNHomePoller already exists, reusing instance",
			"chainSelector", rmnHomeChainSelector,
			"address", hexEncodedAddr,
		)
		return instance, nil
	}

	rmnHomeBoundContract := types.BoundContract{
		Address: hexEncodedAddr,
		Name:    consts.ContractNameRMNHome,
	}

	if err := contractReader.Bind(ctx, []types.BoundContract{rmnHomeBoundContract}); err != nil {
		return nil, fmt.Errorf("failed to bind RMNHome contract: %w", err)
	}

	rmnHomeReader := newRMNHomePoller(
		contractReader,
		rmnHomeBoundContract,
		logutil.WithComponent(lggr, "RMNHomePoller"),
		pollingInterval,
	)

	if err := rmnHomeReader.Start(ctx); err != nil {
		return nil, fmt.Errorf("failed to start RMNHome reader: %w", err)
	}

	instances[key] = rmnHomeReader
	return rmnHomeReader, nil
}

func newRMNHomePoller(
	contractReader contractreader.ContractReaderFacade,
	rmnHomeBoundContract types.BoundContract,
	lggr logger.Logger,
	pollingInterval time.Duration,
) *rmnHomePoller {
	return &rmnHomePoller{
		stopCh:               make(chan struct{}),
		contractReader:       contractReader,
		rmnHomeBoundContract: rmnHomeBoundContract,
		rmnHomeState:         rmnHomeState{},
		mutex:                &sync.RWMutex{},
		failedPolls:          0,
		lggr:                 lggr,
		pollingDuration:      pollingInterval,
	}
}

func (r *rmnHomePoller) Start(context.Context) error {
	return r.sync.StartOnce(r.Name(), func() error {
		r.lggr.Infow("Start Polling RMNHome")
		r.wg.Add(1)
		go r.poll()
		return nil
	})
}

func (r *rmnHomePoller) poll() {
	defer r.wg.Done()
	ctx, cancel := r.stopCh.NewCtx()
	defer cancel()
	// Initial fetch once poll is called before any ticks
	if err := r.fetchAndSetRmnHomeConfigs(ctx); err != nil {
		// Just log, don't return error as we want to keep polling
		r.lggr.Errorw("Initial fetch of on-chain configs failed", "err", err)
	}

	ticker := time.NewTicker(r.pollingDuration)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			r.mutex.Lock()
			r.failedPolls = 0
			r.mutex.Unlock()
			return
		case <-ticker.C:
			if err := r.fetchAndSetRmnHomeConfigs(ctx); err != nil {
				r.mutex.Lock()
				r.failedPolls++
				r.mutex.Unlock()
			} else {
				r.mutex.Lock()
				r.failedPolls = 0
				r.mutex.Unlock()
			}
		}
	}
}

func (r *rmnHomePoller) fetchAndSetRmnHomeConfigs(ctx context.Context) error {
	var activeAndCandidateConfigs GetAllConfigsResponse
	err := r.contractReader.GetLatestValue(
		ctx,
		r.rmnHomeBoundContract.ReadIdentifier(consts.MethodNameGetAllConfigs),
		primitives.Unconfirmed,
		map[string]interface{}{},
		&activeAndCandidateConfigs,
	)
	if err != nil {
		return fmt.Errorf("error fetching RMNHomeConfig: %w", err)
	}
	r.lggr.Infow("Fetched RMNHomeConfigs",
		"activeConfig", activeAndCandidateConfigs.ActiveConfig.ConfigDigest,
		"candidateConfig", activeAndCandidateConfigs.CandidateConfig.ConfigDigest,
	)

	if activeAndCandidateConfigs.ActiveConfig.ConfigDigest.IsEmpty() &&
		activeAndCandidateConfigs.CandidateConfig.ConfigDigest.IsEmpty() {
		return fmt.Errorf("both active and candidate config digests are empty")
	}

	r.setRMNHomeState(
		activeAndCandidateConfigs.ActiveConfig.ConfigDigest,
		activeAndCandidateConfigs.CandidateConfig.ConfigDigest,
		convertOnChainConfigToRMNHomeChainConfig(
			r.lggr,
			activeAndCandidateConfigs.ActiveConfig,
			activeAndCandidateConfigs.CandidateConfig,
		),
	)

	return nil
}

func (r *rmnHomePoller) setRMNHomeState(
	activeConfigDigest cciptypes.Bytes32,
	candidateConfigDigest cciptypes.Bytes32,
	rmnHomeConfig map[cciptypes.Bytes32]rmntypes.HomeConfig) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	s := &r.rmnHomeState

	s.activeConfigDigest = activeConfigDigest
	s.candidateConfigDigest = candidateConfigDigest
	s.rmnHomeConfig = rmnHomeConfig
}

func (r *rmnHomePoller) GetRMNNodesInfo(configDigest cciptypes.Bytes32) ([]rmntypes.HomeNodeInfo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	_, ok := r.rmnHomeState.rmnHomeConfig[configDigest]
	if !ok {
		return nil, fmt.Errorf("configDigest %s not found in RMNHomeConfig", configDigest)

	}
	return r.rmnHomeState.rmnHomeConfig[configDigest].Nodes, nil
}

func (r *rmnHomePoller) IsRMNHomeConfigDigestSet(configDigest cciptypes.Bytes32) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	_, ok := r.rmnHomeState.rmnHomeConfig[configDigest]
	return ok
}

func (r *rmnHomePoller) GetFObserve(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	_, ok := r.rmnHomeState.rmnHomeConfig[configDigest]
	if !ok {
		return nil, fmt.Errorf("configDigest %s not found in RMNHomeConfig", configDigest)
	}
	return r.rmnHomeState.rmnHomeConfig[configDigest].SourceChainF, nil
}

func (r *rmnHomePoller) GetOffChainConfig(configDigest cciptypes.Bytes32) (cciptypes.Bytes, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	cfg, ok := r.rmnHomeState.rmnHomeConfig[configDigest]
	if !ok {
		return nil, fmt.Errorf("configDigest %s not found in RMNHomeConfig", configDigest)
	}
	return cfg.OffchainConfig, nil
}

func (r *rmnHomePoller) GetAllConfigDigests() (
	activeConfigDigest cciptypes.Bytes32,
	candidateConfigDigest cciptypes.Bytes32) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.rmnHomeState.activeConfigDigest, r.rmnHomeState.candidateConfigDigest
}

func (r *rmnHomePoller) Close() error {
	err := r.sync.StopOnce(r.Name(), func() error {
		defer r.wg.Wait()
		close(r.stopCh)
		return nil
	})

	if errors.Is(err, services.ErrAlreadyStopped) {
		return nil
	}

	return err
}

func (r *rmnHomePoller) Ready() error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.sync.Ready()
}

func (r *rmnHomePoller) HealthReport() map[string]error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if r.failedPolls >= MaxFailedPolls {
		err := fmt.Errorf("polling failed %d times in a row (maximum allowed: %d)", r.failedPolls, MaxFailedPolls)
		r.sync.SvcErrBuffer.Append(err)
	}
	return map[string]error{r.Name(): r.sync.Healthy()}
}

func (r *rmnHomePoller) Name() string {
	return r.lggr.Name()
}

func validate(config VersionedConfig) error {
	// check if the versionesconfigwithdigests are set (can be empty)
	if config.ConfigDigest.IsEmpty() {
		return fmt.Errorf("configDigest is empty")
	}
	return nil
}

func convertOnChainConfigToRMNHomeChainConfig(
	lggr logger.Logger,
	primaryConfig VersionedConfig,
	secondaryConfig VersionedConfig,
) map[cciptypes.Bytes32]rmntypes.HomeConfig {
	if primaryConfig.ConfigDigest.IsEmpty() && secondaryConfig.ConfigDigest.IsEmpty() {
		lggr.Warnw("no on chain RMNHomeConfigs found, both digests are empty")
		return map[cciptypes.Bytes32]rmntypes.HomeConfig{}
	}

	versionedConfigWithDigests := []VersionedConfig{primaryConfig}
	if !secondaryConfig.ConfigDigest.IsEmpty() {
		versionedConfigWithDigests = append(versionedConfigWithDigests, secondaryConfig)
	}

	rmnHomeConfigs := make(map[cciptypes.Bytes32]rmntypes.HomeConfig)
	for _, versionedConfig := range versionedConfigWithDigests {
		err := validate(versionedConfig)
		if err != nil {
			lggr.Warnw("invalid on chain RMNHomeConfig", "err", err)
			continue
		}

		nodes := make([]rmntypes.HomeNodeInfo, len(versionedConfig.StaticConfig.Nodes))
		for i, node := range versionedConfig.StaticConfig.Nodes {
			pubKey := ed25519.PublicKey(node.OffchainPublicKey[:])

			nodes[i] = rmntypes.HomeNodeInfo{
				ID:                    rmntypes.NodeID(i),
				PeerID:                ragep2ptypes.PeerID(node.PeerID),
				OffchainPublicKey:     &pubKey,
				SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](),
				StreamNamePrefix:      "ccip-rmn/v1_6/", // todo: when contract is updated, this should be fetched from the contract
			}
		}

		homeFMap := make(map[cciptypes.ChainSelector]int)

		for _, chain := range versionedConfig.DynamicConfig.SourceChains {
			homeFMap[chain.ChainSelector] = int(chain.FObserve)
			for j := 0; j < len(nodes); j++ {
				isObserver, err := IsNodeObserver(chain, j, len(nodes))
				if err != nil {
					lggr.Warnw("failed to check if node is observer", "err", err)
					continue
				}
				if isObserver {
					nodes[j].SupportedSourceChains.Add(chain.ChainSelector)
				}
			}
		}

		rmnHomeConfigs[versionedConfig.ConfigDigest] = rmntypes.HomeConfig{
			Nodes:          nodes,
			SourceChainF:   homeFMap,
			ConfigDigest:   versionedConfig.ConfigDigest,
			OffchainConfig: versionedConfig.DynamicConfig.OffchainConfig,
		}
	}
	return rmnHomeConfigs
}

// IsNodeObserver checks if a node is an observer for the given source chain
func IsNodeObserver(sourceChain SourceChain, nodeIndex int, totalNodes int) (bool, error) {
	if totalNodes > rmnMaxSizeCommittee || totalNodes <= 0 {
		return false, fmt.Errorf("invalid total nodes: %d", totalNodes)
	}

	if nodeIndex < 0 || nodeIndex >= totalNodes {
		return false, fmt.Errorf("invalid node index: %d", nodeIndex)
	}

	// Validate the bitmap
	maxValidBitmap := new(big.Int).Lsh(big.NewInt(1), uint(totalNodes))
	maxValidBitmap.Sub(maxValidBitmap, big.NewInt(1))
	if sourceChain.ObserverNodesBitmap.Cmp(maxValidBitmap) > 0 {
		return false, fmt.Errorf("invalid observer nodes bitmap")
	}

	// Create a big.Int with 1 shifted left by nodeIndex
	mask := new(big.Int).Lsh(big.NewInt(1), uint(nodeIndex))

	// Perform the bitwise AND operation
	result := new(big.Int).And(sourceChain.ObserverNodesBitmap, mask)

	// Check if the result equals the mask
	return result.Cmp(mask) == 0, nil
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

var _ RMNHome = (*rmnHomePoller)(nil)
