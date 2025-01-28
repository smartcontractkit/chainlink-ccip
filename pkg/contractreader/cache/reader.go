package cache

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type Reader interface {
	contractreader.Extended
}

type state struct {
}

type maybe[T any] struct {
	value T
	err   error
}

func (m *maybe[T]) set(value any, err error) {
	m.value = value.(T)
	m.err = err
}
func (m *maybe[T]) IsError() bool {
	return m.err != nil
}

type bgReader struct {
	contractreader.Extended

	lggr   logger.Logger
	stopCh services.StopChan
	sync   services.StateMachine
	wg     sync.WaitGroup

	// state
	nativeTokenAddress    maybe[cciptypes.Bytes]
	commitLatestOCRConfig maybe[OCRConfigResponse]
	execLatestOCRConfig   maybe[OCRConfigResponse]
	offrampStaticConfig   maybe[reader.OffRampStaticChainConfig]
	offrampDynamicConfig  maybe[reader.OffRampDynamicChainConfig]
	offrampAllChains      maybe[reader.SelectorsAndConfigs]
	onrampDynamicConfig   maybe[reader.GetOnRampDynamicConfigResponse]
	rmnDigestHeader       maybe[RMNDigestHeader]
	rmnVersionedConfig    maybe[reader.VersionedConfig]
	rmnRemoteAddress      maybe[cciptypes.Bytes32]
	feeQuoterConfig       maybe[reader.FeeQuoterStaticConfig]
	stateMu               sync.RWMutex
}

func (r *bgReader) Start(_ context.Context) error {
	return r.sync.StartOnce("Polling Configuration", func() error {
		r.wg.Add(1)
		go r.poll()
		return nil
	})
}

func (r *bgReader) Close() error {
	err := r.sync.StopOnce("Polling Configuration", func() error {
		defer r.wg.Wait()
		close(r.stopCh)
		return nil
	})

	if errors.Is(err, services.ErrAlreadyStopped) {
		return nil
	}
	return err
}

func (r *bgReader) poll() {
	defer r.wg.Done()
	ctx, cancel := r.stopCh.NewCtx()
	defer cancel()

	// Initial fetch once poll is called before any ticks
	if err := r.batchFetchConfig(ctx); err != nil {
		r.lggr.Errorw("Initial fetch of on-chain configs failed", "err", err)
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := r.batchFetchConfig(ctx); err != nil {
				r.lggr.Errorw("Initial fetch of on-chain configs failed", "err", err)
			}
		}
	}
}

func (r *bgReader) getConfig() {
	r.Lock()
	defer r.Lock()

	if time.Since(r.lastUpdateAt) > 30*time.Second {
		r.batchFetchConfig(ctx)
	}

	return r.config
}

func (r *bgReader) batchFetchConfig(ctx context.Context) error {
	// Router
	var nativeTokenAddress cciptypes.Bytes
	routerRequests := []types.BatchRead{
		{
			ReadName:  consts.MethodNameRouterGetWrappedNative,
			Params:    map[string]any{},
			ReturnVal: &nativeTokenAddress,
		},
	}

	// OnRamp
	var onrampDynamicConfig reader.GetOnRampDynamicConfigResponse
	onRampRequests := []types.BatchRead{
		{
			ReadName:  consts.MethodNameOnRampGetDynamicConfig,
			Params:    map[string]any{},
			ReturnVal: &onrampDynamicConfig,
		},
	}

	// Offramp
	var commitLatestOCRConfig OCRConfigResponse
	var execLatestOCRConfig OCRConfigResponse
	var staticConfig reader.OffRampStaticChainConfig
	var dynamicConfig reader.OffRampDynamicChainConfig
	var selectorsAndConf reader.SelectorsAndConfigs
	offRampRequests := []types.BatchRead{
		{
			ReadName: consts.MethodNameOffRampLatestConfigDetails,
			Params: map[string]any{
				"ocrPluginType": consts.PluginTypeCommit,
			},
			ReturnVal: &commitLatestOCRConfig,
		},
		{
			ReadName: consts.MethodNameOffRampLatestConfigDetails,
			Params: map[string]any{
				"ocrPluginType": consts.PluginTypeExecute,
			},
			ReturnVal: &execLatestOCRConfig,
		},
		{
			ReadName:  consts.MethodNameOffRampGetStaticConfig,
			Params:    map[string]any{},
			ReturnVal: &staticConfig,
		},
		{
			ReadName:  consts.MethodNameOffRampGetDynamicConfig,
			Params:    map[string]any{},
			ReturnVal: &dynamicConfig,
		},
		{
			ReadName:  consts.MethodNameOffRampGetAllSourceChainConfigs,
			Params:    map[string]any{},
			ReturnVal: &selectorsAndConf,
		},
	}

	// RMNRemote
	var rmnDigestHeader RMNDigestHeader
	var rmnVersionConfig reader.VersionedConfig
	rmnRemoteRequests := []types.BatchRead{
		{
			ReadName:  consts.MethodNameGetReportDigestHeader,
			Params:    map[string]any{},
			ReturnVal: &rmnDigestHeader,
		},
		{
			ReadName:  consts.MethodNameGetVersionedConfig,
			Params:    map[string]any{},
			ReturnVal: &rmnVersionConfig,
		},
	}

	// RMNProxy
	var rmnRemoteAddress []byte
	rmnProxyRequests := []types.BatchRead{
		{
			ReadName:  consts.MethodNameGetARM,
			Params:    map[string]any{},
			ReturnVal: &rmnRemoteAddress,
		},
	}

	// FeeQuoter
	var feeQuoterConfig reader.FeeQuoterStaticConfig
	feeQuoterRequests := []types.BatchRead{
		{
			ReadName:  consts.MethodNameFeeQuoterGetStaticConfig,
			Params:    map[string]any{},
			ReturnVal: &feeQuoterConfig,
		},
	}

	requests := contractreader.ExtendedBatchGetLatestValuesRequest{
		consts.ContractNameRouter:    routerRequests,
		consts.ContractNameOnRamp:    onRampRequests,
		consts.ContractNameOffRamp:   offRampRequests,
		consts.ContractNameRMNRemote: rmnRemoteRequests,
		consts.ContractNameRMNProxy:  rmnProxyRequests,
		consts.ContractNameFeeQuoter: feeQuoterRequests,
	}

	batchResult, err := r.Extended.ExtendedBatchGetLatestValues(ctx, requests)
	if err != nil {
		return err
	}

	r.stateMu.Lock()
	defer r.stateMu.Unlock()

	for contract, results := range batchResult {
		switch contract.Name {
		case consts.ContractNameRouter:
			unwrapResult(results, 0, &r.nativeTokenAddress)
		case consts.ContractNameOnRamp:
			unwrapResult(results, 0, &r.onrampDynamicConfig)
		case consts.ContractNameOffRamp:
			unwrapResult(results, 0, &r.commitLatestOCRConfig)
			unwrapResult(results, 1, &r.execLatestOCRConfig)
			unwrapResult(results, 2, &r.offrampStaticConfig)
			unwrapResult(results, 3, &r.offrampDynamicConfig)
			unwrapResult(results, 4, &r.offrampAllChains)
		case consts.ContractNameRMNRemote:
			unwrapResult(results, 0, &r.rmnDigestHeader)
			unwrapResult(results, 1, &r.rmnVersionedConfig)
		case consts.ContractNameRMNProxy:
			unwrapResult(results, 0, &r.rmnRemoteAddress)
		case consts.ContractNameFeeQuoter:
			unwrapResult(results, 0, &r.feeQuoterConfig)
		}
	}
	return nil
}

func unwrapResult[T any](results types.ContractBatchResults, i int, m *maybe[T]) {
	if len(results) <= i {
		m.err = errors.New("no result")
		return
	}

	result, err := results[i].GetResult()
	m.set(result, err)
}

type ConfigInfo struct {
	ConfigDigest                   [32]byte
	F                              uint8
	N                              uint8
	IsSignatureVerificationEnabled bool
}

type OCRConfig struct {
	ConfigInfo   ConfigInfo
	Signers      [][]byte
	Transmitters [][]byte
}

type OCRConfigResponse struct {
	OCRConfig OCRConfig
}

type RMNDigestHeader struct {
	DigestHeader cciptypes.Bytes32
}
