package commit

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/internal/builder"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/metrics"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
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
	estimatedMaxNumberOfPricedTokens = 14_445

	// maxQueryLength is set to twice the maximum size of a theoretical merkle root processor query
	// that assumes estimatedMaxNumberOfSourceChains source chains and
	// estimatedMaxRmnNodesCount (theoretical max) RMN nodes.
	// check factory_test for the calculation
	maxQueryLength = 242_869

	// maxObservationLength is set to the maximum size of an observation
	// check factory_test for the calculation
	maxObservationLength = 650_307

	// maxOutcomeLength is set to the maximum size of an outcome
	// check factory_test for the calculation
	maxOutcomeLength = 700_620

	// maxReportLength is set to an estimate of a maximum report size
	// check factory_test for the calculation
	maxReportLength = 128_2933

	// maxReportCount is set very high because some chains may require many reports per round.
	maxReportCount = 1000
)

type PluginFactory struct {
	baseLggr          logger.Logger
	donID             plugintypes.DonID
	ocrConfig         reader.OCR3ConfigWithMeta
	commitCodec       cciptypes.CommitPluginCodec
	msgHasher         cciptypes.MessageHasher
	addrCodec         cciptypes.AddressCodec
	homeChainReader   reader.HomeChain
	homeChainSelector cciptypes.ChainSelector
	chainAccessors    map[cciptypes.ChainSelector]cciptypes.ChainAccessor
	extendedReaders   map[cciptypes.ChainSelector]contractreader.Extended
	chainWriters      map[cciptypes.ChainSelector]types.ContractWriter
	rmnPeerClient     rmn.PeerClient
	rmnCrypto         cciptypes.RMNCrypto
}

type CommitPluginFactoryParams struct {
	Lggr              logger.Logger
	DonID             plugintypes.DonID
	OcrConfig         reader.OCR3ConfigWithMeta
	CommitCodec       cciptypes.CommitPluginCodec
	MsgHasher         cciptypes.MessageHasher
	AddrCodec         cciptypes.AddressCodec
	HomeChainReader   reader.HomeChain
	HomeChainSelector cciptypes.ChainSelector
	ChainAccessors    map[cciptypes.ChainSelector]cciptypes.ChainAccessor
	ExtendedReaders   map[cciptypes.ChainSelector]contractreader.Extended
	ContractWriters   map[cciptypes.ChainSelector]types.ContractWriter
	RmnPeerClient     rmn.PeerClient
	RmnCrypto         cciptypes.RMNCrypto
}

// NewCommitPluginFactory creates a new PluginFactory instance. For commit plugin, oracle instances are not managed by
// the factory. It is safe to assume that a factory instance will create exactly one plugin instance.
func NewCommitPluginFactory(params CommitPluginFactoryParams) *PluginFactory {
	return &PluginFactory{
		baseLggr:          params.Lggr,
		donID:             params.DonID,
		ocrConfig:         params.OcrConfig,
		commitCodec:       params.CommitCodec,
		msgHasher:         params.MsgHasher,
		addrCodec:         params.AddrCodec,
		homeChainReader:   params.HomeChainReader,
		homeChainSelector: params.HomeChainSelector,
		chainAccessors:    params.ChainAccessors,
		extendedReaders:   params.ExtendedReaders,
		chainWriters:      params.ContractWriters,
		rmnPeerClient:     params.RmnPeerClient,
		rmnCrypto:         params.RmnCrypto,
	}
}

//nolint:gocyclo
func (p *PluginFactory) NewReportingPlugin(ctx context.Context, config ocr3types.ReportingPluginConfig,
) (ocr3types.ReportingPlugin[[]byte], ocr3types.ReportingPluginInfo, error) {
	lggr := logutil.WithPluginConstants(p.baseLggr, "Commit", p.donID, config.OracleID, config.ConfigDigest)

	offchainConfig, err := pluginconfig.DecodeCommitOffchainConfig(config.OffchainConfig)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to decode commit offchain config: %w", err)
	}

	lggr.Infow("Commit Offchain Config", "offchainConfig", offchainConfig)

	if err = offchainConfig.ApplyDefaultsAndValidate(); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to validate commit offchain config: %w", err)
	}

	var oracleIDToP2PID = make(map[commontypes.OracleID]ragep2ptypes.PeerID)
	for oracleID, node := range p.ocrConfig.Config.Nodes {
		oracleIDToP2PID[commontypes.OracleID(oracleID)] = node.P2pID
	}

	// Validate that the readerFacades were already wrapped in the Extended interface from core.
	readerFacades := make(map[cciptypes.ChainSelector]contractreader.ContractReaderFacade, len(p.extendedReaders))
	for chain, cr := range p.extendedReaders {
		readerFacades[chain] = cr
	}

	// Bind the RMNHome contract
	var rmnHomeReader readerpkg.RMNHome
	if offchainConfig.RMNEnabled {
		rmnHomeAddress := p.ocrConfig.Config.RmnHomeAddress
		rmnCr, ok := p.extendedReaders[p.homeChainSelector]
		if !ok {
			return nil,
				ocr3types.ReportingPluginInfo{},
				fmt.Errorf("failed to find contract reader for home chain %d", p.homeChainSelector)
		}

		rmnHomeReader, err = readerpkg.NewRMNHomeChainReader(
			ctx,
			lggr,
			readerpkg.HomeChainPollingInterval,
			p.homeChainSelector,
			rmnHomeAddress,
			rmnCr,
		)
		if err != nil {
			return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to initialize RMNHome reader: %w", err)
		}

		if err := rmnHomeReader.Start(ctx); err != nil {
			return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to start RMNHome: %w", err)
		}
	}

	if err := validateOcrConfig(p.ocrConfig.Config); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("validate ocr config: %w", err)
	}

	ccipReader, err := readerpkg.NewCCIPChainReader(
		ctx,
		logutil.WithComponent(lggr, "CCIPReader"),
		p.chainAccessors,
		readerFacades,
		p.chainWriters,
		p.ocrConfig.Config.ChainSelector,
		p.ocrConfig.Config.OfframpAddress,
		p.addrCodec,
	)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to create CCIP chain reader: %w", err)
	}

	// The node supports the chain that the token prices are on.
	_, ok := p.chainAccessors[offchainConfig.PriceFeedChainSelector]
	if ok {
		// Bind all token aggregate contracts
		for _, info := range offchainConfig.TokenInfo {
			priceAggAddress, err := cciptypes.NewUnknownAddressFromHex(string(info.AggregatorAddress))
			if err != nil {
				return nil, ocr3types.ReportingPluginInfo{},
					fmt.Errorf("failed to create unknown address from aggregator address %s: %w",
						info.AggregatorAddress, err)
			}
			err = p.chainAccessors[offchainConfig.PriceFeedChainSelector].
				Sync(ctx, consts.ContractNamePriceAggregator, priceAggAddress)
			if err != nil {
				return nil, ocr3types.ReportingPluginInfo{},
					fmt.Errorf("failed to sync price aggregator contract via chainAccessor %s on chain %d: %w",
						consts.ContractNamePriceAggregator, offchainConfig.PriceFeedChainSelector, err)
			}
		}
	}

	onChainTokenPricesReader := readerpkg.NewPriceReader(
		logutil.WithComponent(lggr, "PriceReader"),
		p.chainAccessors,
		offchainConfig.TokenInfo,
		ccipReader,
		offchainConfig.PriceFeedChainSelector,
		p.addrCodec,
	)

	bhClient := beholder.GetClient().ForPackage("ocr3-ccip-commit")

	metricsReporter, err := metrics.NewPromReporter(lggr, p.ocrConfig.Config.ChainSelector, bhClient)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to create metrics reporter: %w", err)
	}

	reportBuilder, err := builder.NewReportBuilder(
		offchainConfig.RMNEnabled,
		offchainConfig.MaxMerkleRootsPerReport,
		offchainConfig.MaxPricesPerReport,
	)
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to create report builder: %w", err)
	}

	return NewPlugin(
			p.donID,
			oracleIDToP2PID,
			offchainConfig,
			p.ocrConfig.Config.ChainSelector,
			ccipReader,
			onChainTokenPricesReader,
			p.commitCodec,
			p.msgHasher,
			lggr,
			p.homeChainReader,
			rmnHomeReader,
			p.rmnCrypto,
			p.rmnPeerClient,
			config,
			metricsReporter,
			p.addrCodec,
			reportBuilder,
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
var _ core.OCR3ReportingPluginFactory = &PluginFactory{}
