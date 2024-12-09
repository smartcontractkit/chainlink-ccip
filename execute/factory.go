package execute

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"

	"github.com/smartcontractkit/chainlink-ccip/execute/costlymessages"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/gas"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	// Estimated maximum number of source chains the system will support.
	// This value should be adjusted as we approach supporting that number of chains.
	// Its primary purpose is to assist in defining the limits below.
	estimatedMaxNumberOfSourceChains = 900

	// maxQueryLength is set to disable queries because they are not used.
	maxQueryLength = 0

	// maxObservationLength is set to the maximum size of an observation
	// check factory_test for the calculation.
	// this is being set to the max maximum observation length due to
	// the observations being so large at the moment, especially when
	// commit reports have many messages.
	// in order to meaningfully decrease this we need to drastically optimise
	// our observation sizes.
	maxObservationLength = ocr3types.MaxMaxObservationLength

	// maxOutcomeLength is set to the maximum size of an outcome
	// check factory_test for the calculation. This is not limited because
	// these are not sent over the network.
	maxOutcomeLength = ocr3types.MaxMaxOutcomeLength

	// maxReportLength is set to an estimate of a maximum report size.
	// This can be tuned over time, it may be more efficient to have
	// smaller reports.

	maxReportLength = ocr3types.MaxMaxReportLength // allowing large reports for now

	// maxReportCount controls how many OCR3 reports can be returned. Note that
	// the actual exec report type (ExecutePluginReport) may contain multiple
	// per-source-chain reports. These are not limited by this value.
	maxReportCount = 1
)

// PluginFactoryConstructor implements common OCR3ReportingPluginClient and is used for initializing a plugin factory
// and a validation service.
type PluginFactoryConstructor struct{}

func NewPluginFactoryConstructor() *PluginFactoryConstructor {
	return &PluginFactoryConstructor{}
}
func (p PluginFactoryConstructor) NewReportingPluginFactory(
	ctx context.Context,
	config core.ReportingPluginServiceConfig,
	grpcProvider grpc.ClientConnInterface,
	pipelineRunner core.PipelineRunnerService,
	telemetry core.TelemetryService,
	errorLog core.ErrorLog,
	capRegistry core.CapabilitiesRegistry,
	keyValueStore core.KeyValueStore,
	relayerSet core.RelayerSet,
) (core.OCR3ReportingPluginFactory, error) {
	return nil, errors.New("unimplemented")
}

func (p PluginFactoryConstructor) NewValidationService(ctx context.Context) (core.ValidationService, error) {
	panic("implement me")
}

// PluginFactory implements common ReportingPluginFactory and is used for (re-)initializing commit plugin instances.
type PluginFactory struct {
	lggr             logger.Logger
	donID            plugintypes.DonID
	ocrConfig        reader.OCR3ConfigWithMeta
	execCodec        cciptypes.ExecutePluginCodec
	msgHasher        cciptypes.MessageHasher
	homeChainReader  reader.HomeChain
	estimateProvider gas.EstimateProvider
	tokenDataEncoder cciptypes.TokenDataEncoder
	contractReaders  map[cciptypes.ChainSelector]types.ContractReader
	chainWriters     map[cciptypes.ChainSelector]types.ContractWriter
}

func NewPluginFactory(
	lggr logger.Logger,
	donID plugintypes.DonID,
	ocrConfig reader.OCR3ConfigWithMeta,
	execCodec cciptypes.ExecutePluginCodec,
	msgHasher cciptypes.MessageHasher,
	homeChainReader reader.HomeChain,
	tokenDataEncoder cciptypes.TokenDataEncoder,
	estimateProvider gas.EstimateProvider,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	chainWriters map[cciptypes.ChainSelector]types.ContractWriter,
) *PluginFactory {
	return &PluginFactory{
		lggr:             lggr,
		donID:            donID,
		ocrConfig:        ocrConfig,
		execCodec:        execCodec,
		msgHasher:        msgHasher,
		homeChainReader:  homeChainReader,
		estimateProvider: estimateProvider,
		contractReaders:  contractReaders,
		chainWriters:     chainWriters,
		tokenDataEncoder: tokenDataEncoder,
	}
}

func (p PluginFactory) NewReportingPlugin(
	ctx context.Context, config ocr3types.ReportingPluginConfig,
) (ocr3types.ReportingPlugin[[]byte], ocr3types.ReportingPluginInfo, error) {
	offchainConfig, err := pluginconfig.DecodeExecuteOffchainConfig(config.OffchainConfig)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to decode exec offchain config: %w", err)
	}

	if err = offchainConfig.Validate(); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to validate exec offchain config: %w", err)
	}

	var oracleIDToP2PID = make(map[commontypes.OracleID]ragep2ptypes.PeerID)
	for oracleID, node := range p.ocrConfig.Config.Nodes {
		oracleIDToP2PID[commontypes.OracleID(oracleID)] = node.P2pID
	}

	// map types to the facade.
	readers := make(map[cciptypes.ChainSelector]contractreader.ContractReaderFacade)
	for chain, cr := range p.contractReaders {
		readers[chain] = cr
	}

	ccipReader := readerpkg.NewCCIPChainReader(
		ctx,
		p.lggr,
		readers,
		p.chainWriters,
		p.ocrConfig.Config.ChainSelector,
		p.ocrConfig.Config.OfframpAddress,
	)

	tokenDataObserver, err := tokendata.NewConfigBasedCompositeObservers(
		ctx,
		logger.Named(p.lggr, "BaseCompositeObserver"),
		p.ocrConfig.Config.ChainSelector,
		offchainConfig.TokenDataObservers,
		p.tokenDataEncoder,
		readers,
	)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to create token data observer: %w", err)
	}

	costlyMessageObserver := costlymessages.NewObserverWithDefaults(
		p.lggr,
		true,
		ccipReader,
		offchainConfig.RelativeBoostPerWaitHour,
		p.estimateProvider,
	)

	return NewPlugin(
			p.donID,
			config,
			offchainConfig,
			p.ocrConfig.Config.ChainSelector,
			oracleIDToP2PID,
			ccipReader,
			p.execCodec,
			p.msgHasher,
			p.homeChainReader,
			tokenDataObserver,
			p.estimateProvider,
			p.lggr,
			costlyMessageObserver,
		), ocr3types.ReportingPluginInfo{
			Name: "CCIPRoleExecute",
			Limits: ocr3types.ReportingPluginLimits{
				// No query for this execute implementation.
				MaxQueryLength:       maxQueryLength,
				MaxObservationLength: maxObservationLength,
				MaxOutcomeLength:     maxOutcomeLength,
				MaxReportLength:      maxReportLength,
				MaxReportCount:       maxReportCount,
			},
		}, nil
}

func (p PluginFactory) Name() string {
	panic("implement me")
}

func (p PluginFactory) Start(ctx context.Context) error {
	panic("implement me")
}

func (p PluginFactory) Close() error {
	panic("implement me")
}

func (p PluginFactory) Ready() error {
	panic("implement me")
}

func (p PluginFactory) HealthReport() map[string]error {
	panic("implement me")
}

// Interface compatibility checks.
var _ core.OCR3ReportingPluginClient = &PluginFactoryConstructor{}
var _ core.OCR3ReportingPluginFactory = &PluginFactory{}
