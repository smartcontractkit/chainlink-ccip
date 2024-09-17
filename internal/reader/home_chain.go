package reader

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"math/big"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
)

const (
	defaultConfigPageSize = uint64(100)
	rmnMaxSizeCommittee   = 256 // bitmap is 256 bits making the max committee size 256
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
}

type CCIPHome interface {
	// GetChainConfig gets the ChainConfig for the given chainSelector
	GetChainConfig(chainSelector cciptypes.ChainSelector) (ChainConfig, error)
	GetAllChainConfigs() (map[cciptypes.ChainSelector]ChainConfig, error)
	// GetSupportedChainsForPeer Gets all chain selectors that the peerID can read/write from/to
	GetSupportedChainsForPeer(id libocrtypes.PeerID) (mapset.Set[cciptypes.ChainSelector], error)
	// GetKnownCCIPChains Gets all chain selectors that are known to CCIP
	GetKnownCCIPChains() (mapset.Set[cciptypes.ChainSelector], error)
	// GetFChain Gets the FChain value for each chain
	GetFChain() (map[cciptypes.ChainSelector]int, error)
	// GetOCRConfigs Gets the OCR3Configs for a given donID and pluginType
	GetOCRConfigs(ctx context.Context, donID uint32, pluginType uint8) ([]OCR3ConfigWithMeta, error)
}

type HomeChain interface {
	CCIPHome
	RMNHome
	services.Service
}

type state struct {
	// gets updated by the polling loop
	chainConfigs map[cciptypes.ChainSelector]ChainConfig
	// mapping between each node's peerID and the chains it supports. derived from chainConfigs
	nodeSupportedChains map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector]
	// set of chains that are known to CCIP, derived from chainConfigs
	knownSourceChains mapset.Set[cciptypes.ChainSelector]
	// map of chain to FChain value, derived from chainConfigs
	fChain map[cciptypes.ChainSelector]int
}

type homeChainPoller struct {
	wg              sync.WaitGroup
	stopCh          services.StopChan
	sync            services.StateMachine
	homeChainReader types.ContractReader
	lggr            logger.Logger
	mutex           *sync.RWMutex
	state           state
	rmnHomeConfig   map[cciptypes.Bytes32]rmntypes.RMNHomeConfig
	failedPolls     uint
	// TODO: currently unused but will be passed into GetLatestValue
	// once the chainlink-common breaking change comes in
	// (https://github.com/smartcontractkit/chainlink-common/pull/603).
	ccipConfigBoundContract types.BoundContract
	rmnHomeBoundContract    types.BoundContract
	// How frequently the poller fetches the chain configs
	pollingDuration time.Duration
}

const MaxFailedPolls = 10

func NewHomeChainConfigPoller(
	homeChainReader types.ContractReader,
	lggr logger.Logger,
	pollingInterval time.Duration,
	ccipConfigBoundContract types.BoundContract,
	rmnHomeBoundContract types.BoundContract,
) HomeChain {
	return &homeChainPoller{
		stopCh:                  make(chan struct{}),
		homeChainReader:         homeChainReader,
		state:                   state{},
		rmnHomeConfig:           make(map[cciptypes.Bytes32]rmntypes.RMNHomeConfig),
		mutex:                   &sync.RWMutex{},
		failedPolls:             0,
		lggr:                    lggr,
		pollingDuration:         pollingInterval,
		ccipConfigBoundContract: ccipConfigBoundContract,
		rmnHomeBoundContract:    rmnHomeBoundContract,
	}
}

func (r *homeChainPoller) Start(ctx context.Context) error {
	r.lggr.Infow("Start Polling ChainConfig")
	return r.sync.StartOnce(r.Name(), func() error {
		r.wg.Add(1)
		go r.poll()
		return nil
	})
}

func (r *homeChainPoller) poll() {
	defer r.wg.Done()
	ctx, cancel := r.stopCh.NewCtx()
	defer cancel()
	// Initial fetch once poll is called before any ticks
	if err := r.fetchAndSetConfigs(ctx); err != nil {
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
			// TODO: fetch concurrently using a waitgroup
			if err := r.fetchAndSetConfigs(ctx); err != nil {
				r.mutex.Lock()
				r.failedPolls++
				r.mutex.Unlock()
			}
			if err := r.fetchAndSetRMNConfig(ctx); err != nil {
				r.mutex.Lock()
				r.failedPolls++
				r.mutex.Unlock()
			}
		}
	}
}

func (r *homeChainPoller) fetchAndSetConfigs(ctx context.Context) error {
	var allChainConfigInfos []ChainConfigInfo
	pageIndex := uint64(0)

	for {
		var chainConfigInfos []ChainConfigInfo
		err := r.homeChainReader.GetLatestValue(
			ctx,
			r.ccipConfigBoundContract.ReadIdentifier(consts.MethodNameGetAllChainConfigs),
			primitives.Unconfirmed,
			map[string]interface{}{
				"pageIndex": pageIndex,
				"pageSize":  defaultConfigPageSize,
			},
			&chainConfigInfos,
		)
		if err != nil {
			return fmt.Errorf("get config index:%d pagesize:%d: %w", pageIndex, defaultConfigPageSize, err)
		}

		allChainConfigInfos = append(allChainConfigInfos, chainConfigInfos...)

		if uint64(len(chainConfigInfos)) < defaultConfigPageSize {
			break
		}

		pageIndex++
	}

	r.setState(convertOnChainConfigToHomeChainConfig(r.lggr, allChainConfigInfos))

	if len(allChainConfigInfos) == 0 {
		// That's a legitimate case if there are no chain configs on chain yet
		r.lggr.Warnw("no on chain configs found")
		return nil
	}

	return nil
}

func (r *homeChainPoller) setState(chainConfigs map[cciptypes.ChainSelector]ChainConfig) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	s := &r.state
	s.chainConfigs = chainConfigs
	s.nodeSupportedChains = createNodesSupportedChains(chainConfigs)
	s.knownSourceChains = createKnownChains(chainConfigs)
	s.fChain = createFChain(chainConfigs)
}

func (r *homeChainPoller) fetchAndSetRMNConfig(ctx context.Context) error {
	var versionedConfigWithDigests []rmntypes.VersionedConfigWithDigest
	err := r.homeChainReader.GetLatestValue(
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

func (r *homeChainPoller) setRMNHomeState(rmnHomeConfig map[cciptypes.Bytes32]rmntypes.RMNHomeConfig) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.rmnHomeConfig = rmnHomeConfig
}

func (r *homeChainPoller) GetChainConfig(chainSelector cciptypes.ChainSelector) (ChainConfig, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	s := r.state
	if chainConfig, ok := s.chainConfigs[chainSelector]; ok {
		return chainConfig, nil
	}
	return ChainConfig{}, fmt.Errorf("chain config not found for chain %v", chainSelector)
}

func (r *homeChainPoller) GetAllChainConfigs() (map[cciptypes.ChainSelector]ChainConfig, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.state.chainConfigs, nil
}

func (r *homeChainPoller) GetSupportedChainsForPeer(
	id libocrtypes.PeerID,
) (mapset.Set[cciptypes.ChainSelector], error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	s := r.state
	if _, ok := s.nodeSupportedChains[id]; !ok {
		// empty set to denote no chains supported
		return mapset.NewSet[cciptypes.ChainSelector](), nil
	}
	return s.nodeSupportedChains[id], nil
}

func (r *homeChainPoller) GetKnownCCIPChains() (mapset.Set[cciptypes.ChainSelector], error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	knownSourceChains := mapset.NewSet[cciptypes.ChainSelector]()
	for chain := range r.state.chainConfigs {
		knownSourceChains.Add(chain)
	}

	return knownSourceChains, nil
}

func (r *homeChainPoller) GetFChain() (map[cciptypes.ChainSelector]int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.state.fChain, nil
}

func (r *homeChainPoller) GetOCRConfigs(
	ctx context.Context, donID uint32, pluginType uint8,
) ([]OCR3ConfigWithMeta, error) {
	var ocrConfigs []OCR3ConfigWithMeta
	err := r.homeChainReader.GetLatestValue(
		ctx,
		r.ccipConfigBoundContract.ReadIdentifier(consts.MethodNameGetOCRConfig),
		primitives.Unconfirmed,
		map[string]any{
			"donId":      donID,
			"pluginType": pluginType,
		}, &ocrConfigs)
	if err != nil {
		return nil, fmt.Errorf("error fetching OCR configs: %w", err)
	}

	return ocrConfigs, nil
}

func (r *homeChainPoller) GetRMNNodesInfo(configDigest cciptypes.Bytes32) ([]rmntypes.RMNHomeNodeInfo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.rmnHomeConfig[configDigest].Nodes, nil
}

func (r *homeChainPoller) IsRMNHomeConfigDigestSet(configDigest cciptypes.Bytes32) (bool, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	_, ok := r.rmnHomeConfig[configDigest]
	return ok, nil
}

func (r *homeChainPoller) GetMinObservers(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]uint64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.rmnHomeConfig[configDigest].MinObservers, nil
}

func (r *homeChainPoller) GetOffChainConfig(configDigest cciptypes.Bytes32) (cciptypes.Bytes, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.rmnHomeConfig[configDigest].OffchainConfig, nil
}

func (r *homeChainPoller) Close() error {
	return r.sync.StopOnce(r.Name(), func() error {
		defer r.wg.Wait()
		close(r.stopCh)
		return nil
	})
}

func (r *homeChainPoller) Ready() error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.sync.Ready()
}

func (r *homeChainPoller) HealthReport() map[string]error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if r.failedPolls >= MaxFailedPolls {
		r.sync.SvcErrBuffer.Append(fmt.Errorf("polling failed %d times in a row", MaxFailedPolls))
	}
	return map[string]error{r.Name(): r.sync.Healthy()}
}

func (r *homeChainPoller) Name() string {
	return "homeChainPoller"
}

func createFChain(chainConfigs map[cciptypes.ChainSelector]ChainConfig) map[cciptypes.ChainSelector]int {
	fChain := map[cciptypes.ChainSelector]int{}
	for chain, config := range chainConfigs {
		fChain[chain] = config.FChain
	}
	return fChain
}

func createKnownChains(chainConfigs map[cciptypes.ChainSelector]ChainConfig) mapset.Set[cciptypes.ChainSelector] {
	knownChains := mapset.NewSet[cciptypes.ChainSelector]()
	for chain := range chainConfigs {
		knownChains.Add(chain)
	}
	return knownChains
}

func createNodesSupportedChains(
	chainConfigs map[cciptypes.ChainSelector]ChainConfig,
) map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector] {
	nodeSupportedChains := map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector]{}
	for chainSelector, config := range chainConfigs {
		for _, p2pID := range config.SupportedNodes.ToSlice() {
			if _, ok := nodeSupportedChains[p2pID]; !ok {
				nodeSupportedChains[p2pID] = mapset.NewSet[cciptypes.ChainSelector]()
			}
			// add chain to SupportedChains
			nodeSupportedChains[p2pID].Add(chainSelector)
		}
	}
	return nodeSupportedChains
}

func convertOnChainConfigToHomeChainConfig(
	lggr logger.Logger,
	chainConfigInfos []ChainConfigInfo,
) map[cciptypes.ChainSelector]ChainConfig {
	chainConfigs := make(map[cciptypes.ChainSelector]ChainConfig)
	for _, chainConfigInfo := range chainConfigInfos {
		chainSelector := chainConfigInfo.ChainSelector
		chainConfig := chainConfigInfo.ChainConfig
		decoded, err := chainconfig.DecodeChainConfig(chainConfig.Config)
		if err != nil {
			lggr.Warnw(fmt.Sprintf("failed to decode opaque chain config of chain selector %d", chainSelector), "err", err)
			continue
		}

		chainConfigs[chainSelector] = ChainConfig{
			FChain:         int(chainConfig.FChain),
			SupportedNodes: mapset.NewSet(chainConfig.Readers...),
			Config:         decoded,
		}
	}
	return chainConfigs
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
			for j := 0; j < rmnMaxSizeCommittee; j++ {
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

// HomeChainConfigMapper This is a 1-1 mapping between the config that we get from the contract to make
// se/deserializing easier
type HomeChainConfigMapper struct {
	Readers []libocrtypes.PeerID `json:"readers"`
	FChain  uint8                `json:"fChain"`
	Config  []byte               `json:"config"`
}

// ChainConfigInfo This is a 1-1 mapping between the config that we get from the contract to make
// se/deserializing easier
type ChainConfigInfo struct {
	// nolint:lll // don't split up the long url
	// Calling function https://github.com/smartcontractkit/ccip/blob/330c5e98f624cfb10108c92fe1e00ced6d345a99/contracts/src/v0.8/ccip/capability/CCIPConfig.sol#L140
	ChainSelector cciptypes.ChainSelector `json:"chainSelector"`
	ChainConfig   HomeChainConfigMapper   `json:"chainConfig"`
}

// ChainConfig will live on the home chain and will be used to update chain configuration like F value and supported
// nodes dynamically.
type ChainConfig struct {
	// FChain defines the FChain value for the chain. FChain is used while forming consensus based on the observations.
	FChain int `json:"fChain"`
	// SupportedNodes is a map of PeerIDs to SupportedChains.
	SupportedNodes mapset.Set[libocrtypes.PeerID] `json:"supportedNodes"`
	// Config is the chain specific configuration.
	Config chainconfig.ChainConfig `json:"config"`
}

// OCR3Config mirrors CCIPConfig.sol's OCR3Config struct
type OCR3Config struct {
	PluginType            uint8                   `json:"pluginType"`
	ChainSelector         cciptypes.ChainSelector `json:"chainSelector"`
	F                     uint8                   `json:"F"`
	OffchainConfigVersion uint64                  `json:"offchainConfigVersion"`
	OfframpAddress        []byte                  `json:"offrampAddress"`
	P2PIds                [][32]byte              `json:"p2pIds"`
	Signers               [][]byte                `json:"signers"`
	Transmitters          [][]byte                `json:"transmitters"`
	OffchainConfig        []byte                  `json:"offchainConfig"`
}

// OCR3ConfigWithmeta mirrors CCIPConfig.sol's OCR3ConfigWithMeta struct
type OCR3ConfigWithMeta struct {
	Config       OCR3Config `json:"config"`
	ConfigCount  uint64     `json:"configCount"`
	ConfigDigest [32]byte   `json:"configDigest"`
}

var _ HomeChain = (*homeChainPoller)(nil)
