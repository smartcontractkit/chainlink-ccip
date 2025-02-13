package reader

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"sync"
	"sync/atomic"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

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
// It is running in the background with a polling interval of pollingDuration.
// In order to reduce a pressure on the underlying RPC node and redundancy there always should be a single
// instance for particular chain/address pair. Please see getRMNHomePoller for more details.
type rmnHomePoller struct {
	lggr                 logger.Logger
	contractReader       contractreader.ContractReaderFacade
	rmnHomeBoundContract types.BoundContract
	pollingDuration      time.Duration

	// Background polling
	sync   services.StateMachine
	stopCh services.StopChan
	wg     sync.WaitGroup

	// State
	observers      []RMNHome
	observersMu    *sync.Mutex
	rmnHomeState   rmnHomeState
	rmnHomeStateMu *sync.RWMutex
	failedPolls    atomic.Uint32
}

// getRMNHomePoller returns a rmnHomePoller instance if it already exists, else creates a new one.
// Returned rmnHomePoller is always started, therefore user of that API doesn't need to care about
// concurrency and accidentally starting the same poller multiple times.
//
// In most of the cases, there is going to be only a single RMNHome deployed on a single chain. However,
// OCR3Config allows to pass different RMNHome addresses for different chains. Therefore, we need to
// support that and maintain a separate rmnHomePoller for each RMNHome address.
// Having a singleton here is aimed to reduce the background polling to the underlying RPC node.
func getRMNHomePoller(
	ctx context.Context,
	lggr logger.Logger,
	rmnHomeChainSelector cciptypes.ChainSelector,
	rmnHomeAddress []byte,
	contractReader contractreader.ContractReaderFacade,
	pollingInterval time.Duration,
) (*rmnHomePoller, error) {
	instancesMu.Lock()
	defer instancesMu.Unlock()

	hexEncodedAddr := "0x" + hex.EncodeToString(rmnHomeAddress)
	key := fmt.Sprintf("%s-%s", rmnHomeChainSelector.String(), hexEncodedAddr)

	instance, ok := getInstance(key)
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
		logutil.WithComponent(lggr, "RMNHomePoller"),
		contractReader,
		rmnHomeBoundContract,
		pollingInterval,
	)

	instances[key] = rmnHomeReader
	return rmnHomeReader, nil
}

func getInstance(key string) (*rmnHomePoller, bool) {
	instance, ok := instances[key]
	if !ok {
		return nil, false
	}

	if notStopped := instance.sync.IfNotStopped(func() {}); !notStopped {
		return nil, false
	}

	return instance, ok
}

// newRMNHomePoller creates a new rmnHomePoller instance. Only meant to be used by getRMNHomePoller
// or in tests. Never use that directly in the production code otherwise it will create to exhaustive polling.
func newRMNHomePoller(
	lggr logger.Logger,
	contractReader contractreader.ContractReaderFacade,
	rmnHomeBoundContract types.BoundContract,
	pollingInterval time.Duration,
) *rmnHomePoller {
	return &rmnHomePoller{
		stopCh:               make(chan struct{}),
		contractReader:       contractReader,
		rmnHomeBoundContract: rmnHomeBoundContract,
		rmnHomeState:         rmnHomeState{},
		observersMu:          &sync.Mutex{},
		rmnHomeStateMu:       &sync.RWMutex{},
		failedPolls:          atomic.Uint32{},
		lggr:                 lggr,
		pollingDuration:      pollingInterval,
	}
}

func (r *rmnHomePoller) Start(_ context.Context, caller RMNHome) error {
	r.observersMu.Lock()
	defer r.observersMu.Unlock()

	r.observers = append(r.observers, caller)

	if len(r.observers) > 1 {
		return nil
	}

	return r.sync.StartOnce(r.Name(), func() error {
		r.lggr.Infow("Start Polling RMNHome")
		r.wg.Add(1)
		go r.poll()
		return nil
	})
}

func (r *rmnHomePoller) Close(caller RMNHome) error {
	r.observersMu.Lock()
	defer r.observersMu.Unlock()

	r.observers = slices.DeleteFunc(r.observers, func(r RMNHome) bool {
		return r == caller
	})

	if len(r.observers) > 0 {
		return nil
	}

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
	return r.sync.Ready()
}

func (r *rmnHomePoller) HealthReport() map[string]error {
	f := r.failedPolls.Load()
	if f >= maxFailedPolls {
		err := fmt.Errorf("polling failed %d times in a row (maximum allowed: %d)", f, maxFailedPolls)
		r.sync.SvcErrBuffer.Append(err)
	}
	return map[string]error{r.Name(): r.sync.Healthy()}
}

func (r *rmnHomePoller) Name() string {
	return r.lggr.Name()
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
			r.failedPolls.Store(0)
			return
		case <-ticker.C:
			if err := r.fetchAndSetRmnHomeConfigs(ctx); err != nil {
				r.failedPolls.Add(1)
			} else {
				r.failedPolls.Store(0)
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

func (r *rmnHomePoller) getRMNHomeState() rmnHomeState {
	r.rmnHomeStateMu.RLock()
	defer r.rmnHomeStateMu.RUnlock()

	return r.rmnHomeState
}

func (r *rmnHomePoller) setRMNHomeState(
	activeConfigDigest cciptypes.Bytes32,
	candidateConfigDigest cciptypes.Bytes32,
	rmnHomeConfig map[cciptypes.Bytes32]rmntypes.HomeConfig,
) {
	r.rmnHomeStateMu.Lock()
	defer r.rmnHomeStateMu.Unlock()

	s := &r.rmnHomeState
	s.activeConfigDigest = activeConfigDigest
	s.candidateConfigDigest = candidateConfigDigest
	s.rmnHomeConfig = rmnHomeConfig
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

	versionedConfigWithDigests := make([]VersionedConfig, 0, 2)

	if !primaryConfig.ConfigDigest.IsEmpty() {
		versionedConfigWithDigests = append(versionedConfigWithDigests, primaryConfig)
	}

	if !secondaryConfig.ConfigDigest.IsEmpty() {
		versionedConfigWithDigests = append(versionedConfigWithDigests, secondaryConfig)
	}

	rmnHomeConfigs := make(map[cciptypes.Bytes32]rmntypes.HomeConfig)
	for _, cfg := range versionedConfigWithDigests {
		err := validate(cfg)
		if err != nil {
			lggr.Warnw("invalid on chain RMNHomeConfig", "err", err)
			continue
		}

		nodes := make([]rmntypes.HomeNodeInfo, len(cfg.StaticConfig.Nodes))
		for i, node := range cfg.StaticConfig.Nodes {
			pubKey := ed25519.PublicKey(node.OffchainPublicKey[:])

			nodes[i] = rmntypes.HomeNodeInfo{
				ID:                    rmntypes.NodeID(i),
				PeerID:                ragep2ptypes.PeerID(node.PeerID),
				OffchainPublicKey:     &pubKey,
				SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](),
				StreamNamePrefix:      "ccip-rmn/v1_6/", // todo: when contract is updated, this should be fetched from the contract
			}
		}

		homeFMap := make(map[cciptypes.ChainSelector]int, len(cfg.DynamicConfig.SourceChains))

		for _, chain := range cfg.DynamicConfig.SourceChains {
			homeFMap[chain.ChainSelector] = int(chain.FObserve)
			for j := 0; j < len(nodes); j++ {
				isObserver, err := isNodeObserver(chain, j, len(nodes))
				if err != nil {
					lggr.Warnw("failed to check if node is observer", "err", err)
					continue
				}
				if isObserver {
					nodes[j].SupportedSourceChains.Add(chain.ChainSelector)
				}
			}
		}

		rmnHomeConfigs[cfg.ConfigDigest] = rmntypes.HomeConfig{
			Nodes:          nodes,
			SourceChainF:   homeFMap,
			ConfigDigest:   cfg.ConfigDigest,
			OffchainConfig: cfg.DynamicConfig.OffchainConfig,
		}
	}
	return rmnHomeConfigs
}

// isNodeObserver checks if a node is an observer for the given source chain
func isNodeObserver(sourceChain SourceChain, nodeIndex int, totalNodes int) (bool, error) {
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
