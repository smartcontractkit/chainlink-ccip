package commit

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
)

const maxReportTransmissionCheckAttempts = 5
const maxQueryLength = 1024 * 1024 // 1MB

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
	lggr            logger.Logger
	ocrConfig       reader.OCR3ConfigWithMeta
	commitCodec     cciptypes.CommitPluginCodec
	msgHasher       cciptypes.MessageHasher
	homeChainReader reader.HomeChain
	contractReaders map[cciptypes.ChainSelector]types.ContractReader
	chainWriters    map[cciptypes.ChainSelector]types.ChainWriter
}

func NewPluginFactory(
	lggr logger.Logger,
	ocrConfig reader.OCR3ConfigWithMeta,
	commitCodec cciptypes.CommitPluginCodec,
	msgHasher cciptypes.MessageHasher,
	homeChainReader reader.HomeChain,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	chainWriters map[cciptypes.ChainSelector]types.ChainWriter,
) *PluginFactory {
	return &PluginFactory{
		lggr:            lggr,
		ocrConfig:       ocrConfig,
		commitCodec:     commitCodec,
		msgHasher:       msgHasher,
		homeChainReader: homeChainReader,
		contractReaders: contractReaders,
		chainWriters:    chainWriters,
	}
}

func (p *PluginFactory) NewReportingPlugin(config ocr3types.ReportingPluginConfig,
) (ocr3types.ReportingPlugin[[]byte], ocr3types.ReportingPluginInfo, error) {
	offchainConfig, err := pluginconfig.DecodeCommitOffchainConfig(config.OffchainConfig)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to decode commit offchain config: %w", err)
	}

	if err = offchainConfig.Validate(); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to validate commit offchain config: %w", err)
	}

	var oracleIDToP2PID = make(map[commontypes.OracleID]ragep2ptypes.PeerID)
	for oracleID, p2pID := range p.ocrConfig.Config.P2PIds {
		oracleIDToP2PID[commontypes.OracleID(oracleID)] = p2pID
	}

	var onChainTokenPricesReader reader.PriceReader
	// The node supports the chain that the token prices are on.
	tokenPricesCr, ok := p.contractReaders[offchainConfig.PriceFeedChainSelector]
	if ok {
		// Bind all token aggregate contracts
		var bcs []types.BoundContract
		for _, info := range offchainConfig.TokenInfo {
			bcs = append(bcs, types.BoundContract{
				Address: info.AggregatorAddress,
				Name:    consts.ContractNamePriceAggregator,
			})
		}
		if err1 := tokenPricesCr.Bind(context.Background(), bcs); err1 != nil {
			return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to bind token price contracts: %w", err1)
		}
		onChainTokenPricesReader = reader.NewOnchainTokenPricesReader(
			tokenPricesCr,
			offchainConfig.TokenInfo,
		)
	}

	ccipReader := readerpkg.NewCCIPChainReader(
		p.lggr,
		p.contractReaders,
		p.chainWriters,
		p.ocrConfig.Config.ChainSelector,
		p.ocrConfig.Config.OfframpAddress,
	)
	return NewPlugin(
			context.Background(),
			config.OracleID,
			oracleIDToP2PID,
			pluginconfig.CommitPluginConfig{
				DestChain:                          p.ocrConfig.Config.ChainSelector,
				NewMsgScanBatchSize:                merklemulti.MaxNumberTreeLeaves,
				MaxReportTransmissionCheckAttempts: maxReportTransmissionCheckAttempts,
				OffchainConfig:                     offchainConfig,
			},
			ccipReader,
			onChainTokenPricesReader,
			p.commitCodec,
			p.msgHasher,
			p.lggr,
			p.homeChainReader,
			config,
			rmntypes.RMNConfig{}, // todo
		), ocr3types.ReportingPluginInfo{
			Name: "CCIPRoleCommit",
			Limits: ocr3types.ReportingPluginLimits{
				MaxQueryLength:       maxQueryLength,
				MaxObservationLength: 20_000, // 20kB
				MaxOutcomeLength:     10_000, // 10kB
				MaxReportLength:      10_000, // 10kB
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
