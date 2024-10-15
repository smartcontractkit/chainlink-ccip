package commit

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

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
	lggr              logger.Logger
	donID             plugintypes.DonID
	ocrConfig         reader.OCR3ConfigWithMeta
	commitCodec       cciptypes.CommitPluginCodec
	msgHasher         cciptypes.MessageHasher
	homeChainReader   reader.HomeChain
	homeChainSelector cciptypes.ChainSelector
	contractReaders   map[cciptypes.ChainSelector]types.ContractReader
	chainWriters      map[cciptypes.ChainSelector]types.ChainWriter
	rmnPeerClient     rmn.PeerClient
	rmnCrypto         cciptypes.RMNCrypto
}

func NewPluginFactory(
	lggr logger.Logger,
	donID plugintypes.DonID,
	ocrConfig reader.OCR3ConfigWithMeta,
	commitCodec cciptypes.CommitPluginCodec,
	msgHasher cciptypes.MessageHasher,
	homeChainReader reader.HomeChain,
	homeChainSelector cciptypes.ChainSelector,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	chainWriters map[cciptypes.ChainSelector]types.ChainWriter,
	rmnPeerClient rmn.PeerClient,
	rmnCrypto cciptypes.RMNCrypto,
) *PluginFactory {
	return &PluginFactory{
		lggr:              lggr,
		donID:             donID,
		ocrConfig:         ocrConfig,
		commitCodec:       commitCodec,
		msgHasher:         msgHasher,
		homeChainReader:   homeChainReader,
		homeChainSelector: homeChainSelector,
		contractReaders:   contractReaders,
		chainWriters:      chainWriters,
		rmnPeerClient:     rmnPeerClient,
		rmnCrypto:         rmnCrypto,
	}
}

func (p *PluginFactory) NewReportingPlugin(ctx context.Context, config ocr3types.ReportingPluginConfig,
) (ocr3types.ReportingPlugin[[]byte], ocr3types.ReportingPluginInfo, error) {
	offchainConfig, err := pluginconfig.DecodeCommitOffchainConfig(config.OffchainConfig)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to decode commit offchain config: %w", err)
	}

	if err = offchainConfig.ApplyDefaultsAndValidate(); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to validate commit offchain config: %w", err)
	}

	var oracleIDToP2PID = make(map[commontypes.OracleID]ragep2ptypes.PeerID)
	for oracleID, node := range p.ocrConfig.Config.Nodes {
		oracleIDToP2PID[commontypes.OracleID(oracleID)] = node.P2pID
	}

	// Bind the RMNHome contract
	var rmnHomeReader reader.RMNHome
	if offchainConfig.RMNEnabled {
		rmnHomeAddress := p.ocrConfig.Config.RmnHomeAddress
		rmnCr, ok := p.contractReaders[p.homeChainSelector]
		if !ok {
			return nil,
				ocr3types.ReportingPluginInfo{},
				fmt.Errorf("failed to find contract reader for home chain %d", p.homeChainSelector)
		}
		rmnHomeBoundContract := types.BoundContract{
			Address: "0x" + hex.EncodeToString(rmnHomeAddress),
			Name:    consts.ContractNameRMNHome,
		}

		if err1 := rmnCr.Bind(ctx, []types.BoundContract{rmnHomeBoundContract}); err1 != nil {
			return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to bind RMNHome contract: %w", err1)
		}
		rmnHomeReader = reader.NewRMNHomePoller(
			rmnCr,
			rmnHomeBoundContract,
			p.lggr,
			5*time.Second,
		)

		if err := rmnHomeReader.Start(ctx); err != nil {
			return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to start RMNHome reader: %w", err)
		}
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
		if err1 := tokenPricesCr.Bind(ctx, bcs); err1 != nil {
			return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to bind token price contracts: %w", err1)
		}
		onChainTokenPricesReader = reader.NewOnchainTokenPricesReader(
			tokenPricesCr,
			offchainConfig.TokenInfo,
		)
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
	return NewPlugin(
			p.donID,
			config.OracleID,
			oracleIDToP2PID,
			offchainConfig,
			p.ocrConfig.Config.ChainSelector,
			ccipReader,
			onChainTokenPricesReader,
			p.commitCodec,
			p.msgHasher,
			p.lggr,
			p.homeChainReader,
			rmnHomeReader,
			p.rmnCrypto,
			p.rmnPeerClient,
			config,
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
