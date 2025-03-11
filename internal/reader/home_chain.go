package reader

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	defaultConfigPageSize = uint64(100)
)

type HomeChain interface {
	GetChainConfig(chainSelector cciptypes.ChainSelector) (ChainConfig, error)
	GetAllChainConfigs() (map[cciptypes.ChainSelector]ChainConfig, error)
	// GetSupportedChainsForPeer Gets all chain selectors that the peerID can read/write from/to
	GetSupportedChainsForPeer(id libocrtypes.PeerID) (mapset.Set[cciptypes.ChainSelector], error)
	// GetKnownCCIPChains Gets all chain selectors that are known to CCIP
	GetKnownCCIPChains() (mapset.Set[cciptypes.ChainSelector], error)
	// GetFChain Gets the FChain value for each chain
	GetFChain() (map[cciptypes.ChainSelector]int, error)
	// GetOCRConfigs Gets the OCR3Configs for a given donID and pluginType
	GetOCRConfigs(ctx context.Context, donID uint32, pluginType uint8) (ActiveAndCandidate, error)
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
	wg                      sync.WaitGroup
	stopCh                  services.StopChan
	sync                    services.StateMachine
	homeChainReader         contractreader.ContractReaderFacade
	lggr                    logger.Logger
	mutex                   *sync.RWMutex
	state                   state
	failedPolls             atomic.Uint32
	ccipConfigBoundContract types.BoundContract
	// How frequently the poller fetches the chain configs
	pollingDuration time.Duration
}

const MaxFailedPolls = 10

func NewHomeChainConfigPoller(
	homeChainReader contractreader.ContractReaderFacade,
	lggr logger.Logger,
	pollingInterval time.Duration,
	ccipConfigBoundContract types.BoundContract,
) HomeChain {
	return &homeChainPoller{
		stopCh:                  make(chan struct{}),
		homeChainReader:         homeChainReader,
		state:                   state{},
		mutex:                   &sync.RWMutex{},
		failedPolls:             atomic.Uint32{},
		lggr:                    lggr,
		pollingDuration:         pollingInterval,
		ccipConfigBoundContract: ccipConfigBoundContract,
	}
}

func (r *homeChainPoller) Start(_ context.Context) error {
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
		r.failedPolls.Add(1)
	} else {
		r.failedPolls.Store(0)
	}

	ticker := time.NewTicker(r.pollingDuration)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			r.failedPolls.Store(0)
			return
		case <-ticker.C:
			if err := r.fetchAndSetConfigs(ctx); err != nil {
				r.failedPolls.Add(1)
			} else {
				r.failedPolls.Store(0)
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

		validCfgInfos := getValidChainConfigInfos(r.lggr, chainConfigInfos)

		if len(validCfgInfos) == 0 {
			continue
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
) (ActiveAndCandidate, error) {
	var (
		activeAndCandidate ActiveAndCandidate
	)

	err := r.homeChainReader.GetLatestValue(
		ctx,
		r.ccipConfigBoundContract.ReadIdentifier(consts.MethodNameGetOCRConfig),
		primitives.Unconfirmed,
		map[string]any{
			"donId":      donID,
			"pluginType": pluginType,
		}, &activeAndCandidate)
	if err != nil {
		return ActiveAndCandidate{}, fmt.Errorf("error fetching OCR configs: %w", err)
	}

	r.lggr.Infow(
		"GetOCRConfigs",
		"activeConfig", activeAndCandidate.ActiveConfig,
		"candidateConfig", activeAndCandidate.CandidateConfig,
	)

	return activeAndCandidate, nil
}

func (r *homeChainPoller) Close() error {
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

func (r *homeChainPoller) Ready() error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.sync.Ready()
}

func (r *homeChainPoller) HealthReport() map[string]error {
	f := r.failedPolls.Load()
	if f >= MaxFailedPolls {
		r.sync.SvcErrBuffer.Append(fmt.Errorf("polling failed %d times in a row", MaxFailedPolls))
	}
	return map[string]error{r.Name(): r.sync.Healthy()}
}

func (r *homeChainPoller) Name() string {
	return r.lggr.Name()
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

func validateChainConfigInfos(chainConfigInfo ChainConfigInfo) error {
	if chainConfigInfo.ChainSelector == 0 {
		return errors.New("chain selector is 0")
	}
	if chainConfigInfo.ChainConfig.FChain == 0 {
		return errors.New("fChain is 0")
	}

	if len(chainConfigInfo.ChainConfig.Readers) == 0 {
		return errors.New("readers is empty")
	}

	if len(chainConfigInfo.ChainConfig.Config) == 0 {
		return errors.New("config is empty")
	}

	return nil
}

func getValidChainConfigInfos(lggr logger.Logger, chainConfigInfos []ChainConfigInfo) []ChainConfigInfo {
	validChainConfigInfos := make([]ChainConfigInfo, 0)
	for _, chainConfigInfo := range chainConfigInfos {
		if err := validateChainConfigInfos(chainConfigInfo); err != nil {
			lggr.Warnw("invalid chain config info", "err", err)
			continue
		}
		validChainConfigInfos = append(validChainConfigInfos, chainConfigInfo)
	}
	return validChainConfigInfos
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
	//nolint:lll // don't split up the long url
	// Calling function https://github.com/smartcontractkit/ccip/blob/330c5e98f624cfb10108c92fe1e00ced6d345a99/contracts/src/v0.8/ccip/capability/CCIPConfig.sol#L140
	ChainSelector cciptypes.ChainSelector `json:"chainSelector"`
	ChainConfig   HomeChainConfigMapper   `json:"chainConfig"`
}

// ChainConfig will live on the home chain and will be used to update chain configuration
// like FRoleDon value and supported nodes dynamically.
type ChainConfig struct {
	// FChain defines the FChain value for the chain. FChain is used while forming consensus based on the observations.
	FChain int `json:"fChain"`
	// SupportedNodes is a map of PeerIDs to SupportedChains.
	SupportedNodes mapset.Set[libocrtypes.PeerID] `json:"supportedNodes"`
	// Config is the chain specific configuration.
	Config chainconfig.ChainConfig `json:"config"`
}

// See https://github.com/smartcontractkit/chainlink/blob/e964798a974f3246ee1da011feffe33509b358df/contracts/src/v0.8/ccip/capability/CCIPHome.sol#L105-L131
//
//nolint:lll
type OCR3Node struct {
	P2pID          [32]byte `json:"p2pId"`
	SignerKey      []byte   `json:"signerKey"`
	TransmitterKey []byte   `json:"transmitterKey"`
}

// See https://github.com/smartcontractkit/chainlink/blob/e964798a974f3246ee1da011feffe33509b358df/contracts/src/v0.8/ccip/capability/CCIPHome.sol#L105-L131
//
//nolint:lll
type OCR3Config struct {
	PluginType            uint8                   `json:"pluginType"`
	ChainSelector         cciptypes.ChainSelector `json:"chainSelector"`
	FRoleDON              uint8                   `json:"fRoleDON"`
	OffchainConfigVersion uint64                  `json:"offchainConfigVersion"`
	OfframpAddress        []byte                  `json:"offrampAddress"`
	RmnHomeAddress        []byte                  `json:"rmnHomeAddress"`
	Nodes                 []OCR3Node              `json:"nodes"`
	OffchainConfig        []byte                  `json:"offchainConfig"`
}

// OCR3ConfigWithMeta
// https://github.com/smartcontractkit/chainlink/blob/e964798a974f3246ee1da011feffe33509b358df/contracts/src/v0.8/ccip/capability/CCIPHome.sol#L105-L131
// TODO: we might need to change it from OCR3ConfigWithMeta to VersionedConfig
// If so, we'll create a new package so that we don't have conflict naming with RMNHome
//
//nolint:lll
type OCR3ConfigWithMeta struct {
	Version      uint32     `json:"version"`
	ConfigDigest [32]byte   `json:"configDigest"`
	Config       OCR3Config `json:"config"`
}

type ActiveAndCandidate struct {
	ActiveConfig    OCR3ConfigWithMeta `json:"activeConfig"`
	CandidateConfig OCR3ConfigWithMeta `json:"candidateConfig"`
}

var _ HomeChain = (*homeChainPoller)(nil)
