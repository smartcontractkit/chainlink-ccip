package execute

import (
	"context"
	"fmt"

	sel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"

	"github.com/smartcontractkit/chainlink-ccip/execute/costlymessages"
	"github.com/smartcontractkit/chainlink-ccip/execute/metrics"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
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

// PluginFactory implements common ReportingPluginFactory and is used for (re-)initializing commit plugin instances.
type PluginFactory struct {
	baseLggr         logger.Logger
	donID            plugintypes.DonID
	ocrConfig        reader.OCR3ConfigWithMeta
	execCodec        cciptypes.ExecutePluginCodec
	msgHasher        cciptypes.MessageHasher
	extraDataCodec   cciptypes.ExtraDataCodec
	homeChainReader  reader.HomeChain
	estimateProvider cciptypes.EstimateProvider
	tokenDataEncoder cciptypes.TokenDataEncoder
	contractReaders  map[cciptypes.ChainSelector]types.ContractReader
	chainWriters     map[cciptypes.ChainSelector]types.ContractWriter
}

type PluginFactoryParams struct {
	Lggr             logger.Logger
	DonID            plugintypes.DonID
	OcrConfig        reader.OCR3ConfigWithMeta
	ExecCodec        cciptypes.ExecutePluginCodec
	MsgHasher        cciptypes.MessageHasher
	ExtraDataCodec   cciptypes.ExtraDataCodec
	HomeChainReader  reader.HomeChain
	TokenDataEncoder cciptypes.TokenDataEncoder
	EstimateProvider cciptypes.EstimateProvider
	ContractReaders  map[cciptypes.ChainSelector]types.ContractReader
	ContractWriters  map[cciptypes.ChainSelector]types.ContractWriter
}

// NewExecutePluginFactory creates a new PluginFactory instance. For execute plugin, oracle instances are not managed by
// the factory. It is safe to assume that a factory instance will create exactly one plugin instance.
func NewExecutePluginFactory(params PluginFactoryParams) *PluginFactory {
	return &PluginFactory{
		baseLggr:         params.Lggr,
		donID:            params.DonID,
		ocrConfig:        params.OcrConfig,
		execCodec:        params.ExecCodec,
		msgHasher:        params.MsgHasher,
		extraDataCodec:   params.ExtraDataCodec,
		homeChainReader:  params.HomeChainReader,
		estimateProvider: params.EstimateProvider,
		tokenDataEncoder: params.TokenDataEncoder,
		contractReaders:  params.ContractReaders,
		chainWriters:     params.ContractWriters,
	}
}

// deprectated: use NewExececutePluginFactory
func NewPluginFactory(
	lggr logger.Logger,
	donID plugintypes.DonID,
	ocrConfig reader.OCR3ConfigWithMeta,
	execCodec cciptypes.ExecutePluginCodec,
	msgHasher cciptypes.MessageHasher,
	extraDataCodec cciptypes.ExtraDataCodec,
	homeChainReader reader.HomeChain,
	tokenDataEncoder cciptypes.TokenDataEncoder,
	estimateProvider cciptypes.EstimateProvider,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	chainWriters map[cciptypes.ChainSelector]types.ContractWriter,
) *PluginFactory {
	return &PluginFactory{
		baseLggr:         lggr,
		donID:            donID,
		ocrConfig:        ocrConfig,
		execCodec:        execCodec,
		msgHasher:        msgHasher,
		extraDataCodec:   extraDataCodec,
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
	lggr := logutil.WithPluginConstants(p.baseLggr, "Execute", p.donID, config.OracleID, config.ConfigDigest)

	offchainConfig, err := pluginconfig.DecodeExecuteOffchainConfig(config.OffchainConfig)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to decode exec offchain config: %w", err)
	}

	if err = offchainConfig.ApplyDefaultsAndValidate(); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to validate exec offchain config: %w", err)
	}

	var oracleIDToP2PID = make(map[commontypes.OracleID]ragep2ptypes.PeerID)
	for oracleID, node := range p.ocrConfig.Config.Nodes {
		oracleIDToP2PID[commontypes.OracleID(oracleID)] = node.P2pID
	}

	// Map contract readers to ContractReaderFacade:
	// - Extended reader adds finality violation and contract binding management.
	// - Observed reader adds metric reporting.
	readers := make(map[cciptypes.ChainSelector]contractreader.ContractReaderFacade)
	for chain, cr := range p.contractReaders {
		chainID, err1 := sel.GetChainIDFromSelector(uint64(chain))
		if err1 != nil {
			return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to get chain id from selector: %w", err1)
		}
		readers[chain] = contractreader.NewExtendedContractReader(
			contractreader.NewObserverReader(cr, lggr, chainID))
	}

	ccipReader := readerpkg.NewCCIPChainReader(
		ctx,
		logutil.WithComponent(lggr, "CCIPReader"),
		readers,
		p.chainWriters,
		p.ocrConfig.Config.ChainSelector,
		p.ocrConfig.Config.OfframpAddress,
		p.extraDataCodec,
	)

	tokenDataObserver, err := tokendata.NewConfigBasedCompositeObservers(
		ctx,
		logutil.WithComponent(lggr, "TokenDataObserver"),
		p.ocrConfig.Config.ChainSelector,
		offchainConfig.TokenDataObservers,
		p.tokenDataEncoder,
		readers,
	)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to create token data observer: %w", err)
	}

	costlyMessageObserver := costlymessages.NewObserverWithDefaults(
		logutil.WithComponent(lggr, "CostlyMessages"),
		true,
		ccipReader,
		offchainConfig.RelativeBoostPerWaitHour,
		p.estimateProvider,
	)

	metricsReporter, err := metrics.NewPromReporter(lggr, p.ocrConfig.Config.ChainSelector)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to create metrics reporter: %w", err)
	}

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
			lggr,
			costlyMessageObserver,
			metricsReporter,
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
var _ core.OCR3ReportingPluginFactory = &PluginFactory{}
