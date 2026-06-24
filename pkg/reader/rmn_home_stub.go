package reader

// TODO(remove-blessing): delete once chainlink drops NewRMNHomeChainReader.

import (
	"context"
	"errors"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

type HomeNodeInfo = rmntypes.HomeNodeInfo

type NodeID = rmntypes.NodeID

// RMNHome is retained for chainlink bootstrap compile compatibility only.
type RMNHome interface {
	GetRMNNodesInfo(configDigest cciptypes.Bytes32) ([]rmntypes.HomeNodeInfo, error)
	IsRMNHomeConfigDigestSet(configDigest cciptypes.Bytes32) bool
	GetRMNEnabledSourceChains(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]bool, error)
	GetFObserve(configDigest cciptypes.Bytes32) (map[cciptypes.ChainSelector]int, error)
	GetOffChainConfig(configDigest cciptypes.Bytes32) (cciptypes.Bytes, error)
	GetAllConfigDigests() (activeConfigDigest cciptypes.Bytes32, candidateConfigDigest cciptypes.Bytes32)
	services.Service
}

type noopRMNHome struct {
	sync services.StateMachine
}

// NewRMNHomeChainReader returns a no-op RMNHome reader.
func NewRMNHomeChainReader(
	_ context.Context,
	_ logger.Logger,
	_ time.Duration,
	_ cciptypes.ChainSelector,
	_ []byte,
	_ contractreader.ContractReaderFacade,
) (RMNHome, error) {
	return &noopRMNHome{}, nil
}

func (r *noopRMNHome) Start(context.Context) error {
	return r.sync.StartOnce(r.Name(), func() error { return nil })
}

func (r *noopRMNHome) Close() error {
	err := r.sync.StopOnce(r.Name(), func() error { return nil })
	if errors.Is(err, services.ErrAlreadyStopped) {
		return nil
	}
	return err
}

func (r *noopRMNHome) Ready() error { return r.sync.Ready() }

func (r *noopRMNHome) HealthReport() map[string]error { return nil }

func (r *noopRMNHome) Name() string { return "RMNHome" }

func (r *noopRMNHome) GetRMNNodesInfo(_ cciptypes.Bytes32) ([]rmntypes.HomeNodeInfo, error) {
	return nil, nil
}

func (r *noopRMNHome) IsRMNHomeConfigDigestSet(_ cciptypes.Bytes32) bool { return false }

func (r *noopRMNHome) GetRMNEnabledSourceChains(_ cciptypes.Bytes32) (map[cciptypes.ChainSelector]bool, error) {
	return map[cciptypes.ChainSelector]bool{}, nil
}

func (r *noopRMNHome) GetFObserve(_ cciptypes.Bytes32) (map[cciptypes.ChainSelector]int, error) {
	return map[cciptypes.ChainSelector]int{}, nil
}

func (r *noopRMNHome) GetOffChainConfig(_ cciptypes.Bytes32) (cciptypes.Bytes, error) {
	return nil, nil
}

func (r *noopRMNHome) GetAllConfigDigests() (cciptypes.Bytes32, cciptypes.Bytes32) {
	return cciptypes.Bytes32{}, cciptypes.Bytes32{}
}

var _ RMNHome = (*noopRMNHome)(nil)
