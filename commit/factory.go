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
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
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

	// Estimated maximum number of RMN nodes the system will support.
	estimatedMaxRmnNodesCount = 256

	// Estimated maximum number of priced tokens that the Commit DON supports.
	// This value does not indicate a system limitation but just an estimation to properly tune the OCR parameters.
	// The value can be adjusted as needed.
	estimatedMaxNumberOfPricedTokens = 10_000

	// maxQueryLength is set to twice the maximum size of a theoretical merkle root processor query
	// that assumes estimatedMaxNumberOfSourceChains source chains and
	// estimatedMaxRmnNodesCount (theoretical max) RMN nodes.
	// check factory_test for the calculation
	maxQueryLength = 559_320

	// maxObservationLength is set to the maximum size of an observation
	// check factory_test for the calculation
	maxObservationLength = 1_047_202

	// maxOutcomeLength is set to the maximum size of an outcome
	// check factory_test for the calculation
	maxOutcomeLength = 1_167_765

	// maxReportLength is set to an estimate of a maximum report size
	// check factory_test for the calculation
	maxReportLength = 993_982

	// maxReportCount is set to 1 because the commit plugin only generates one report per round.
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
	return nil, errors.New("this functionality should not be used")
}

func (p PluginFactoryConstructor) NewValidationService(ctx context.Context) (core.ValidationService, error) {
	return nil, errors.New("this functionality should not be used")
}

type PluginFactory struct {
	lggr              logger.Logger
	donID             plugintypes.DonID
	ocrConfig         reader.OCR3ConfigWithMeta
	commitCodec       cciptypes.CommitPluginCodec
	msgHasher         cciptypes.MessageHasher
	homeChainReader   reader.HomeChain
	homeChainSelector cciptypes.ChainSelector
	contractReaders   map[cciptypes.ChainSelector]types.ContractReader
	chainWriters      map[cciptypes.ChainSelector]types.ContractWriter
	rmnPeerClient     rmn.PeerClient
	rmnCrypto         cciptypes.RMNCrypto
}

// NewPluginFactory creates a new PluginFactory instance. For commit plugin, oracle instances are not managed by the
// factory. It is safe to assume that a factory instance will create exactly one plugin instance.
func NewPluginFactory(
	lggr logger.Logger,
	donID plugintypes.DonID,
	ocrConfig reader.OCR3ConfigWithMeta,
	commitCodec cciptypes.CommitPluginCodec,
	msgHasher cciptypes.MessageHasher,
	homeChainReader reader.HomeChain,
	homeChainSelector cciptypes.ChainSelector,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	chainWriters map[cciptypes.ChainSelector]types.ContractWriter,
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
	var rmnHomeReader readerpkg.RMNHome
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
		rmnHomeReader = readerpkg.NewRMNHomePoller(
			rmnCr,
			rmnHomeBoundContract,
			p.lggr,
			5*time.Second,
		)

		if err := rmnHomeReader.Start(ctx); err != nil {
			return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to start RMNHome reader: %w", err)
		}
	}

	// map types to the facade.
	readers := make(map[cciptypes.ChainSelector]contractreader.ContractReaderFacade, len(p.contractReaders))
	for chain, cr := range p.contractReaders {
		readers[chain] = cr
	}

	if err := validateOcrConfig(p.ocrConfig.Config); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("validate ocr config: %w", err)
	}

	ccipReader := readerpkg.NewCCIPChainReader(
		ctx,
		p.lggr,
		readers,
		p.chainWriters,
		p.ocrConfig.Config.ChainSelector,
		p.ocrConfig.Config.OfframpAddress,
	)

	// The node supports the chain that the token prices are on.
	_, ok := readers[offchainConfig.PriceFeedChainSelector]
	if ok {
		// Bind all token aggregate contracts
		var bcs []types.BoundContract
		for _, info := range offchainConfig.TokenInfo {
			bcs = append(bcs, types.BoundContract{
				Address: string(info.AggregatorAddress),
				Name:    consts.ContractNamePriceAggregator,
			})
		}
		if err1 := readers[offchainConfig.PriceFeedChainSelector].Bind(ctx, bcs); err1 != nil {
			return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to bind token price contracts: %w", err1)
		}
	}

	onChainTokenPricesReader := readerpkg.NewPriceReader(
		p.lggr,
		readers,
		offchainConfig.TokenInfo,
		ccipReader,
		offchainConfig.PriceFeedChainSelector,
	)

	return NewPlugin(
			p.donID,
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
				MaxObservationLength: maxObservationLength,
				MaxOutcomeLength:     maxOutcomeLength,
				MaxReportLength:      maxReportLength,
				MaxReportCount:       maxReportCount,
			},
		}, nil
}

func validateOcrConfig(cfg readerpkg.OCR3Config) error {
	if cfg.ChainSelector == 0 {
		return errors.New("chain selector must be set")
	}

	if cfg.OfframpAddress == nil {
		return errors.New("offramp address must be set")
	}

	return nil
}

func (p PluginFactory) Name() string {
	panic("should not be called")
}

func (p PluginFactory) Start(ctx context.Context) error {
	panic("should not be called")
}

func (p PluginFactory) Close() error {
	panic("should not be called")
}

func (p PluginFactory) Ready() error {
	panic("should not be called")
}

func (p PluginFactory) HealthReport() map[string]error {
	panic("should not be called")
}

// Interface compatibility checks.
var _ core.OCR3ReportingPluginClient = &PluginFactoryConstructor{}
var _ core.OCR3ReportingPluginFactory = &PluginFactory{}
