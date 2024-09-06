package execute

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"

	"github.com/smartcontractkit/libocr/commontypes"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/gas"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
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
	ocrConfig        reader.OCR3ConfigWithMeta
	execCodec        cciptypes.ExecutePluginCodec
	msgHasher        cciptypes.MessageHasher
	homeChainReader  reader.HomeChain
	estimateProvider gas.EstimateProvider
	tokenDataReader  exectypes.TokenDataReader
	contractReaders  map[cciptypes.ChainSelector]types.ContractReader
	chainWriters     map[cciptypes.ChainSelector]types.ChainWriter
}

func NewPluginFactory(
	lggr logger.Logger,
	ocrConfig reader.OCR3ConfigWithMeta,
	execCodec cciptypes.ExecutePluginCodec,
	msgHasher cciptypes.MessageHasher,
	homeChainReader reader.HomeChain,
	tokenDataReader exectypes.TokenDataReader,
	estimateProvider gas.EstimateProvider,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	chainWriters map[cciptypes.ChainSelector]types.ChainWriter,
) *PluginFactory {
	return &PluginFactory{
		lggr:             lggr,
		ocrConfig:        ocrConfig,
		execCodec:        execCodec,
		msgHasher:        msgHasher,
		homeChainReader:  homeChainReader,
		estimateProvider: estimateProvider,
		contractReaders:  contractReaders,
		chainWriters:     chainWriters,
		tokenDataReader:  tokenDataReader,
	}
}

func (p PluginFactory) NewReportingPlugin(
	config ocr3types.ReportingPluginConfig,
) (ocr3types.ReportingPlugin[[]byte], ocr3types.ReportingPluginInfo, error) {
	offchainConfig, err := pluginconfig.DecodeExecuteOffchainConfig(config.OffchainConfig)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to decode exec offchain config: %w", err)
	}

	if err = offchainConfig.Validate(); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to validate exec offchain config: %w", err)
	}

	var oracleIDToP2PID = make(map[commontypes.OracleID]ragep2ptypes.PeerID)
	for oracleID, p2pID := range p.ocrConfig.Config.P2PIds {
		oracleIDToP2PID[commontypes.OracleID(oracleID)] = p2pID
	}

	ccipReader := reader.NewCCIPChainReader(
		p.lggr,
		p.contractReaders,
		p.chainWriters,
		p.ocrConfig.Config.ChainSelector,
	)

	return NewPlugin(
			config,
			pluginconfig.ExecutePluginConfig{
				DestChain:      p.ocrConfig.Config.ChainSelector,
				OffchainConfig: offchainConfig,
			},
			oracleIDToP2PID,
			ccipReader,
			p.execCodec,
			p.msgHasher,
			p.homeChainReader,
			p.tokenDataReader,
			p.estimateProvider,
			p.lggr,
		), ocr3types.ReportingPluginInfo{
			Name: "CCIPRoleExecute",
			Limits: ocr3types.ReportingPluginLimits{
				// No query for this execute implementation.
				MaxQueryLength:       0,
				MaxObservationLength: 20_000,             // 20kB
				MaxOutcomeLength:     20_000,             // 20kB
				MaxReportLength:      maxReportSizeBytes, // 250kB
				MaxReportCount:       10,
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
