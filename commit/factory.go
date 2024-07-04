package commit

import (
	"context"
	"math/big"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"

	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
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
	return nil, nil
}

func (p PluginFactoryConstructor) NewValidationService(ctx context.Context) (core.ValidationService, error) {
	panic("implement me")
}

// PluginFactory implements common ReportingPluginFactory and is used for (re-)initializing commit plugin instances.
type PluginFactory struct {
	lggr            logger.Logger
	ocrConfig       reader.OCR3ConfigWithMeta
	commitCodec     cciptypes.CommitPluginCodec
	msgHasher       cciptypes.MessageHasher
	homeChainReader reader.HomeChain
	contractReaders map[cciptypes.ChainSelector]types.ContractReader
}

func NewPluginFactory(
	lggr logger.Logger,
	ocrConfig reader.OCR3ConfigWithMeta,
	commitCodec cciptypes.CommitPluginCodec,
	msgHasher cciptypes.MessageHasher,
	homeChainReader reader.HomeChain,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
) *PluginFactory {
	return &PluginFactory{
		lggr:            lggr,
		ocrConfig:       ocrConfig,
		commitCodec:     commitCodec,
		msgHasher:       msgHasher,
		homeChainReader: homeChainReader,
		contractReaders: contractReaders,
	}
}

func (p PluginFactory) NewReportingPlugin(config ocr3types.ReportingPluginConfig,
) (ocr3types.ReportingPlugin[[]byte], ocr3types.ReportingPluginInfo, error) {
	// TODO: Get this from ocr config, it's the mapping of the oracleId index in the DON
	var oracleIDToP2pID = make(map[commontypes.OracleID]libocrtypes.PeerID)
	for i, p2pID := range p.ocrConfig.Config.P2PIds {
		oracleIDToP2pID[commontypes.OracleID(i)] = p2pID
	}

	onChainTokenPricesReader := reader.NewOnchainTokenPricesReader(
		reader.TokenPriceConfig{ // TODO: Inject config
			StaticPrices: map[ocr2types.Account]big.Int{},
		},
		nil, // TODO: Inject this
	)
	ccipReader := reader.NewCCIPChainReader(
		p.contractReaders,
		map[cciptypes.ChainSelector]types.ChainWriter{}, // TODO: pass in chain writers
		p.ocrConfig.Config.ChainSelector,
	)
	return NewPlugin(
		context.Background(),
		config.OracleID,
		oracleIDToP2pID,
		cciptypes.CommitPluginConfig{},
		ccipReader,
		onChainTokenPricesReader,
		p.commitCodec,
		p.msgHasher,
		p.lggr,
		p.homeChainReader,
	), ocr3types.ReportingPluginInfo{}, nil
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
