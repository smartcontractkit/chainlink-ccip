package reader

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"math/big"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
)

const (
	rmnMaxSizeCommittee = 256 // bitmap is 256 bits making the max committee size 256
)

type RMNHome interface {
	// GetRMNNodesInfo gets the RMNHomeNodeInfo for the given configDigest
	GetRMNNodesInfo(configDigest cciptypes.Bytes32) ([]rmntypes.RMNHomeNodeInfo, error)
	// IsRMNHomeConfigDigestSet checks if the configDigest is set in the RMNHome contract
	IsRMNHomeConfigDigestSet(configDigest cciptypes.Bytes32) (bool, error)
	// GetMinObservers gets the minimum number of observers required for each chain in the given configDigest
	GetMinObservers(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]uint64, error)
	// GetOffChainConfig gets the offchain config for the given configDigest
	GetOffChainConfig(configDigest cciptypes.Bytes32) (cciptypes.Bytes, error)
	services.Service
}

type RmnHomePoller struct {
	wg                   sync.WaitGroup
	stopCh               services.StopChan
	sync                 services.StateMachine
	contractReader       commontypes.ContractReader
	rmnHomeBoundContract types.BoundContract
	lggr                 logger.Logger
	mutex                *sync.RWMutex
	rmnHomeConfig        map[cciptypes.Bytes32]rmntypes.RMNHomeConfig
	failedPolls          uint
	pollingDuration      time.Duration // How frequently the poller fetches the chain configs
}

func NewRMNHomePoller(
	contractReader types.ContractReader,
	rmnHomeBoundContract types.BoundContract,
	lggr logger.Logger,
	pollingInterval time.Duration,
) RMNHome {
	return &RmnHomePoller{
		stopCh:               make(chan struct{}),
		contractReader:       contractReader,
		rmnHomeBoundContract: rmnHomeBoundContract,
		rmnHomeConfig:        make(map[cciptypes.Bytes32]rmntypes.RMNHomeConfig),
		mutex:                &sync.RWMutex{},
		failedPolls:          0,
		lggr:                 lggr,
		pollingDuration:      pollingInterval,
	}
}

func (r *RmnHomePoller) Start(ctx context.Context) error {
	r.lggr.Infow("Start Polling RMNHome")
	return r.sync.StartOnce(r.Name(), func() error {
		r.wg.Add(1)
		go r.poll()
		return nil
	})
}

func (r *RmnHomePoller) poll() {
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
			}
		}
	}
}

func (r *RmnHomePoller) fetchAndSetRmnHomeConfigs(ctx context.Context) error {
	var versionedConfigWithDigests []rmntypes.VersionedConfigWithDigest
	err := r.contractReader.GetLatestValue(
		ctx,
		r.rmnHomeBoundContract.ReadIdentifier(consts.MethodNameGetVersionedConfigsWithDigests),
		primitives.Unconfirmed,
		map[string]interface{}{
			"offset": 0,
			"limit":  2, // TODO: fetch CONFIG_RING_BUFFER_SIZE
		},
		&versionedConfigWithDigests,
	)
	if err != nil {
		return fmt.Errorf("error fetching RMNHomeConfig: %w", err)
	}

	// TODO: fetch CONFIG_RING_BUFFER_SIZE and compare with len(versionedConfigWithDigests)
	if len(versionedConfigWithDigests) > 2 {
		r.lggr.Errorw("more than 2 RMNHomeConfigs found", "numConfigs", len(versionedConfigWithDigests), "requestedLimit", 2)
		return fmt.Errorf("more than 2 RMNHomeConfigs found")
	}

	r.setRMNHomeState(convertOnChainConfigToRMNHomeChainConfig(r.lggr, versionedConfigWithDigests))

	if len(versionedConfigWithDigests) == 0 {
		// That's a legitimate case if there are no rmn configs on chain yet
		r.lggr.Warnw("no on chain configs found")
		return nil
	}

	return nil
}

func (r *RmnHomePoller) setRMNHomeState(rmnHomeConfig map[cciptypes.Bytes32]rmntypes.RMNHomeConfig) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.rmnHomeConfig = rmnHomeConfig
}

func (r *RmnHomePoller) GetRMNNodesInfo(configDigest cciptypes.Bytes32) ([]rmntypes.RMNHomeNodeInfo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.rmnHomeConfig[configDigest].Nodes, nil
}

func (r *RmnHomePoller) IsRMNHomeConfigDigestSet(configDigest cciptypes.Bytes32) (bool, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	_, ok := r.rmnHomeConfig[configDigest]
	return ok, nil
}

func (r *RmnHomePoller) GetMinObservers(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]uint64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.rmnHomeConfig[configDigest].MinObservers, nil
}

func (r *RmnHomePoller) GetOffChainConfig(configDigest cciptypes.Bytes32) (cciptypes.Bytes, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.rmnHomeConfig[configDigest].OffchainConfig, nil
}

func (r *RmnHomePoller) Close() error {
	return r.sync.StopOnce(r.Name(), func() error {
		defer r.wg.Wait()
		close(r.stopCh)
		return nil
	})
}

func (r *RmnHomePoller) Ready() error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.sync.Ready()
}

func (r *RmnHomePoller) HealthReport() map[string]error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if r.failedPolls >= MaxFailedPolls {
		r.sync.SvcErrBuffer.Append(fmt.Errorf("polling failed %d times in a row", MaxFailedPolls))
	}
	return map[string]error{r.Name(): r.sync.Healthy()}
}

func (r *RmnHomePoller) Name() string {
	return "RmnHomePoller"
}

func convertOnChainConfigToRMNHomeChainConfig(
	lggr logger.Logger,
	versionedConfigWithDigests []rmntypes.VersionedConfigWithDigest,
) map[cciptypes.Bytes32]rmntypes.RMNHomeConfig {
	if len(versionedConfigWithDigests) == 0 {
		lggr.Warnw("no on chain RMNHomeConfigs found")
		return map[cciptypes.Bytes32]rmntypes.RMNHomeConfig{}
	}

	rmnHomeConfigs := make(map[cciptypes.Bytes32]rmntypes.RMNHomeConfig)
	for _, versionedConfigWithDigest := range versionedConfigWithDigests {
		config := versionedConfigWithDigest.VersionedConfig.Config
		nodes := make([]rmntypes.RMNHomeNodeInfo, len(config.Nodes))
		for i, node := range config.Nodes {
			pubKey := ed25519.PublicKey(node.OffchainPublicKey[:])

			nodes[i] = rmntypes.RMNHomeNodeInfo{
				ID:                        rmntypes.NodeID(i),
				PeerID:                    node.PeerID,
				SignObservationsPublicKey: &pubKey,
				SupportedSourceChains:     mapset.NewSet[cciptypes.ChainSelector](),
			}
		}

		minObservers := make(map[cciptypes.ChainSelector]uint64)

		for _, chain := range config.SourceChains {
			minObservers[chain.ChainSelector] = chain.MinObservers
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

		rmnHomeConfigs[versionedConfigWithDigest.ConfigDigest] = rmntypes.RMNHomeConfig{
			Nodes:          nodes,
			MinObservers:   minObservers,
			ConfigDigest:   versionedConfigWithDigest.ConfigDigest,
			OffchainConfig: config.OffchainConfig,
		}
	}
	return rmnHomeConfigs
}

// IsNodeObserver checks if a node is an observer for the given source chain
func IsNodeObserver(sourceChain rmntypes.SourceChain, nodeIndex int, totalNodes int) (bool, error) {
	if totalNodes > rmnMaxSizeCommittee || totalNodes <= 0 {
		return false, fmt.Errorf("invalid total nodes: %d", totalNodes)
	}

	if nodeIndex < 0 || nodeIndex >= totalNodes {
		return false, fmt.Errorf("invalid node index: %d", nodeIndex)
	}

	// Validate the bitmap
	maxValidBitmap := new(big.Int).Lsh(big.NewInt(1), uint(totalNodes))
	maxValidBitmap.Sub(maxValidBitmap, big.NewInt(1))
	if sourceChain.ObserverNodesBitmap.Int.Cmp(maxValidBitmap) > 0 {
		return false, fmt.Errorf("invalid observer nodes bitmap")
	}

	// Create a big.Int with 1 shifted left by nodeIndex
	mask := new(big.Int).Lsh(big.NewInt(1), uint(nodeIndex))

	// Perform the bitwise AND operation
	result := new(big.Int).And(sourceChain.ObserverNodesBitmap.Int, mask)

	// Check if the result equals the mask
	return result.Cmp(mask) == 0, nil
}

var _ RMNHome = (*RmnHomePoller)(nil)
