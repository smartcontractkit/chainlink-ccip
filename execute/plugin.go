package execute

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/smartcontractkit/libocr/quorumhelper"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/costlymessages"
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/gas"
	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	// transmissionDelayMultiplier is used to calculate the transmission delay for each oracle.
	transmissionDelayMultiplier = 3 * time.Second
)

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
	donID        plugintypes.DonID
	reportingCfg ocr3types.ReportingPluginConfig
	offchainCfg  pluginconfig.ExecuteOffchainConfig
	destChain    cciptypes.ChainSelector

	// providers
	ccipReader   readerpkg.CCIPReader
	reportCodec  cciptypes.ExecutePluginCodec
	msgHasher    cciptypes.MessageHasher
	homeChain    reader.HomeChain
	discovery    *discovery.ContractDiscoveryProcessor
	chainSupport plugincommon.ChainSupport

	oracleIDToP2pID       map[commontypes.OracleID]libocrtypes.PeerID
	tokenDataObserver     tokendata.TokenDataObserver
	costlyMessageObserver costlymessages.Observer
	estimateProvider      gas.EstimateProvider
	lggr                  logger.Logger

	// state
	contractsInitialized bool
}

func NewPlugin(
	donID plugintypes.DonID,
	reportingCfg ocr3types.ReportingPluginConfig,
	offchainCfg pluginconfig.ExecuteOffchainConfig,
	destChain cciptypes.ChainSelector,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	ccipReader readerpkg.CCIPReader,
	reportCodec cciptypes.ExecutePluginCodec,
	msgHasher cciptypes.MessageHasher,
	homeChain reader.HomeChain,
	tokenDataObserver tokendata.TokenDataObserver,
	estimateProvider gas.EstimateProvider,
	lggr logger.Logger,
	costlyMessageObserver costlymessages.Observer,
) *Plugin {
	lggr.Infow("creating new plugin instance", "p2pID", oracleIDToP2pID[reportingCfg.OracleID])

	return &Plugin{
		donID:                 donID,
		reportingCfg:          reportingCfg,
		offchainCfg:           offchainCfg,
		destChain:             destChain,
		oracleIDToP2pID:       oracleIDToP2pID,
		ccipReader:            ccipReader,
		reportCodec:           reportCodec,
		msgHasher:             msgHasher,
		homeChain:             homeChain,
		tokenDataObserver:     tokenDataObserver,
		estimateProvider:      estimateProvider,
		lggr:                  lggr,
		costlyMessageObserver: costlyMessageObserver,
		discovery: discovery.NewContractDiscoveryProcessor(
			lggr,
			&ccipReader,
			homeChain,
			destChain,
			reportingCfg.F,
			oracleIDToP2pID,
		),
		chainSupport: plugincommon.NewChainSupport(
			lggr,
			homeChain,
			oracleIDToP2pID,
			reportingCfg.OracleID,
			destChain,
		),
	}
}

func (p *Plugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

func getPendingExecutedReports(
	ctx context.Context,
	ccipReader readerpkg.CCIPReader,
	dest cciptypes.ChainSelector,
	ts time.Time,
	lggr logger.Logger,
) (exectypes.CommitObservations, error) {
	latestReportTS := time.Time{}
	commitReports, err := ccipReader.CommitReportsGTETimestamp(ctx, dest, ts, 1000)
	if err != nil {
		return nil, err
	}
	lggr.Debugw("commit reports", "commitReports", commitReports, "count", len(commitReports))

	// TODO: this could be more efficient. commitReports is also traversed in 'groupByChainSelector'.
	for _, report := range commitReports {
		if report.Timestamp.After(latestReportTS) {
			latestReportTS = report.Timestamp
		}
	}

	groupedCommits := groupByChainSelector(commitReports)
	lggr.Debugw("grouped commits before removing fully executed reports",
		"groupedCommits", groupedCommits, "count", len(groupedCommits))

	// Remove fully executed reports.
	for selector, reports := range groupedCommits {
		if len(reports) == 0 {
			continue
		}

		ranges, err := computeRanges(reports)
		if err != nil {
			return nil, err
		}

		var executedMessages []cciptypes.SeqNumRange
		for _, seqRange := range ranges {
			executedMessagesForRange, err2 := ccipReader.ExecutedMessageRanges(ctx, selector, dest, seqRange)
			if err2 != nil {
				return nil, err2
			}
			executedMessages = append(executedMessages, executedMessagesForRange...)
		}

		// Remove fully executed reports.
		groupedCommits[selector], err = filterOutExecutedMessages(reports, executedMessages)
		if err != nil {
			return nil, err
		}
	}

	lggr.Debugw("grouped commits after removing fully executed reports",
		"groupedCommits", groupedCommits, "count", len(groupedCommits))

	return groupedCommits, nil
}

func (p *Plugin) ValidateObservation(
	ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation,
) error {
	decodedObservation, err := exectypes.DecodeObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("unable to decode observation: %w", err)
	}

	supportedChains, err := p.supportedChains(ao.Observer)
	if err != nil {
		return fmt.Errorf("error finding supported chains by node: %w", err)
	}

	err = validateObserverReadingEligibility(supportedChains, decodedObservation.Messages)
	if err != nil {
		return fmt.Errorf("validate observer reading eligibility: %w", err)
	}

	if err := validateObservedSequenceNumbers(decodedObservation.CommitReports); err != nil {
		return fmt.Errorf("validate observed sequence numbers: %w", err)
	}

	return nil
}

func (p *Plugin) ObservationQuorum(
	_ context.Context, outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (bool, error) {
	// TODO: should we use f+1 (or less) instead of 2f+1 because it is not needed for security?
	return quorumhelper.ObservationCountReachesObservationQuorum(
		quorumhelper.QuorumFPlusOne, p.reportingCfg.N, p.reportingCfg.F, aos), nil
}

// selectReport takes a list of reports in execution order and selects the first reports that fit within the
// maxReportSizeBytes. Individual messages in a commit report may be skipped for various reasons, for example if an
// out-of-order execution is detected or the message requires additional off-chain metadata which is not yet available.
// If there is not enough space in the final report, it may be partially executed by searching for a subset of messages
// which can fit in the final report.
func selectReport(
	ctx context.Context,
	lggr logger.Logger,
	commitReports []exectypes.CommitData,
	builder report.ExecReportBuilder,
) ([]cciptypes.ExecutePluginReportSingleChain, []exectypes.CommitData, error) {
	// TODO: It may be desirable for this entire function to be an interface so that
	//       different selection algorithms can be used.

	var stillPendingReports []exectypes.CommitData
	for i, report := range commitReports {
		// Reports at the end may not have messages yet.
		if len(report.Messages) == 0 {
			stillPendingReports = append(stillPendingReports, report)
			continue
		}

		var err error
		commitReports[i], err = builder.Add(ctx, report)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to add report to builder: %w", err)
		}
		// If the report has not been fully executed, keep it for the next round.
		if len(commitReports[i].Messages) > len(commitReports[i].ExecutedMessages) {
			stillPendingReports = append(stillPendingReports, commitReports[i])
		}
	}

	execReports, err := builder.Build()

	lggr.Infow(
		"reports have been selected",
		"numReports", len(execReports),
		"numPendingReports", len(stillPendingReports))
	return execReports, stillPendingReports, err
}

func (p *Plugin) Reports(
	ctx context.Context, seqNr uint64, outcome ocr3types.Outcome,
) ([]ocr3types.ReportPlus[[]byte], error) {
	if outcome == nil {
		p.lggr.Warn("no outcome, skipping report generation")
		return nil, nil
	}

	decodedOutcome, err := exectypes.DecodeOutcome(outcome)
	if err != nil {
		return nil, fmt.Errorf("unable to decode outcome: %w", err)
	}

	encoded, err := p.reportCodec.Encode(ctx, decodedOutcome.Report)
	if err != nil {
		return nil, fmt.Errorf("unable to encode report: %w", err)
	}

	transmissionSchedule, err := plugincommon.GetTransmissionSchedule(
		p.chainSupport,
		maps.Keys(p.oracleIDToP2pID),
		transmissionDelayMultiplier,
	)
	if err != nil {
		return nil, fmt.Errorf("get transmission schedule: %w", err)
	}
	p.lggr.Debugw("transmission schedule override",
		"transmissionSchedule", transmissionSchedule, "oracleIDToP2PID", p.oracleIDToP2pID)

	report := []ocr3types.ReportPlus[[]byte]{
		{
			ReportWithInfo: ocr3types.ReportWithInfo[[]byte]{
				Report: encoded,
				Info:   nil,
			},
			TransmissionScheduleOverride: transmissionSchedule,
		},
	}

	return report, nil
}

// validateReport validates various aspects of the report.
// Pure checks are placed earlier in the function on purpose to avoid
// unnecessary network or DB I/O.
// If you're added more checks make sure to follow this pattern.
func (p *Plugin) validateReport(
	ctx context.Context,
	seqNr uint64,
	r ocr3types.ReportWithInfo[[]byte],
) (valid bool, decodedReport cciptypes.ExecutePluginReport, err error) {
	// Just a safety check, should never happen.
	if r.Report == nil {
		p.lggr.Warn("skipping nil report", "seqNr", seqNr)
		return false, cciptypes.ExecutePluginReport{}, nil
	}

	decodedReport, err = p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, cciptypes.ExecutePluginReport{}, fmt.Errorf("decode exec plugin report: %w", err)
	}

	if len(decodedReport.ChainReports) == 0 {
		p.lggr.Info("skipping empty report", "seqNr", seqNr)
		return false, cciptypes.ExecutePluginReport{}, nil
	}

	// check if we support the dest, if not we can't do the checks needed.
	supports, err := p.chainSupport.SupportsDestChain(p.reportingCfg.OracleID)
	if err != nil {
		return false, cciptypes.ExecutePluginReport{}, fmt.Errorf("supports dest chain: %w", err)
	}

	if !supports {
		p.lggr.Warnw("dest chain not supported, can't run report acceptance procedures", "seqNr", seqNr)
		return false, cciptypes.ExecutePluginReport{}, nil
	}

	offRampConfigDigest, err := p.ccipReader.GetOffRampConfigDigest(ctx, consts.PluginTypeExecute)
	if err != nil {
		return false, cciptypes.ExecutePluginReport{}, fmt.Errorf("get offramp config digest: %w", err)
	}

	if !bytes.Equal(offRampConfigDigest[:], p.reportingCfg.ConfigDigest[:]) {
		p.lggr.Warnw("my config digest doesn't match offramp's config digest, not accepting/transmitting report",
			"myConfigDigest", p.reportingCfg.ConfigDigest,
			"offRampConfigDigest", hex.EncodeToString(offRampConfigDigest[:]),
			"seqNr", seqNr,
		)
		return false, cciptypes.ExecutePluginReport{}, nil
	}

	return true, decodedReport, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	valid, decodedReport, err := p.validateReport(ctx, seqNr, r)
	if err != nil {
		return false, fmt.Errorf("validate exec report: %w", err)
	}

	if !valid {
		p.lggr.Warn("report not valid, not accepting", "seqNr", seqNr)
		return false, nil
	}

	// TODO: consider doing this in validateReport,
	// will end up doing it in both ShouldAccept and ShouldTransmit.
	sourceChains := slicelib.Map(decodedReport.ChainReports,
		func(r cciptypes.ExecutePluginReportSingleChain) cciptypes.ChainSelector {
			return r.SourceChainSelector
		})
	isCursed, err := plugincommon.IsReportCursed(ctx, p.lggr, p.ccipReader, p.chainSupport.DestChain(), sourceChains)
	if err != nil {
		p.lggr.Errorw(
			"report not accepted due to curse checking error",
			"err", err,
		)
		return false, err
	}
	if isCursed {
		// Detailed logging is already done by IsReportCursed.
		return false, nil
	}

	p.lggr.Info("ShouldAcceptAttestedReport returns true, report accepted")
	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	valid, decodedReport, err := p.validateReport(ctx, seqNr, r)
	if err != nil {
		return valid, fmt.Errorf("validate exec report: %w", err)
	}

	if !valid {
		p.lggr.Warnw("report not valid, not transmitting", "seqNr", seqNr)
		return false, nil
	}

	p.lggr.Infow("transmitting report", "reports", decodedReport.ChainReports)
	return true, nil
}

func (p *Plugin) Close() error {
	return p.tokenDataObserver.Close()
}

func (p *Plugin) supportedChains(id commontypes.OracleID) (mapset.Set[cciptypes.ChainSelector], error) {
	p2pID, exists := p.oracleIDToP2pID[id]
	if !exists {
		return nil, fmt.Errorf("oracle ID %d not found in oracleIDToP2pID", p.reportingCfg.OracleID)
	}
	supportedChains, err := p.homeChain.GetSupportedChainsForPeer(p2pID)
	if err != nil {
		p.lggr.Warnw("error getting supported chains", "err", err)
		return mapset.NewSet[cciptypes.ChainSelector](), fmt.Errorf("error getting supported chains: %w", err)
	}

	return supportedChains, nil
}

func (p *Plugin) supportsDestChain() (bool, error) {
	chains, err := p.supportedChains(p.reportingCfg.OracleID)
	if err != nil {
		return false, fmt.Errorf("error getting supported chains: %w", err)
	}
	return chains.Contains(p.destChain), nil
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
