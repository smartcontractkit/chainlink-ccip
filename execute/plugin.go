package execute

import (
	"context"
	"fmt"
	"sort"
	"sync/atomic"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/execute/temp"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// maxReportSizeBytes that should be returned as an execution report payload.
const maxReportSizeBytes = 250_000

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
	reportingCfg ocr3types.ReportingPluginConfig
	cfg          pluginconfig.ExecutePluginConfig

	// providers
	ccipReader  reader.CCIP
	reportCodec cciptypes.ExecutePluginCodec
	msgHasher   cciptypes.MessageHasher
	homeChain   reader.HomeChain

	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID
	tokenDataReader temp.TokenDataReader
	lastReportTS    *atomic.Int64
	lggr            logger.Logger
}

func NewPlugin(
	reportingCfg ocr3types.ReportingPluginConfig,
	cfg pluginconfig.ExecutePluginConfig,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	ccipReader reader.CCIP,
	reportCodec cciptypes.ExecutePluginCodec,
	msgHasher cciptypes.MessageHasher,
	homeChain reader.HomeChain,
	tokenDataReader temp.TokenDataReader,
	lggr logger.Logger,
) *Plugin {
	lastReportTS := &atomic.Int64{}
	lastReportTS.Store(time.Now().Add(-cfg.MessageVisibilityInterval).UnixMilli())

	// TODO: initialize tokenDataReader.

	return &Plugin{
		reportingCfg:    reportingCfg,
		cfg:             cfg,
		oracleIDToP2pID: oracleIDToP2pID,
		ccipReader:      ccipReader,
		reportCodec:     reportCodec,
		msgHasher:       msgHasher,
		homeChain:       homeChain,
		lastReportTS:    lastReportTS,
		tokenDataReader: tokenDataReader,
		lggr:            lggr,
	}
}

func (p *Plugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

func getPendingExecutedReports(
	ctx context.Context, ccipReader reader.CCIP, dest cciptypes.ChainSelector, ts time.Time,
) (plugintypes.ExecutePluginCommitObservations, time.Time, error) {
	latestReportTS := time.Time{}
	commitReports, err := ccipReader.CommitReportsGTETimestamp(ctx, dest, ts, 1000)
	if err != nil {
		return nil, time.Time{}, err
	}

	// TODO: this could be more efficient. commitReports is also traversed in 'groupByChainSelector'.
	for _, report := range commitReports {
		if report.Timestamp.After(latestReportTS) {
			latestReportTS = report.Timestamp
		}
	}

	groupedCommits := groupByChainSelector(commitReports)

	// Remove fully executed reports.
	for selector, reports := range groupedCommits {
		if len(reports) == 0 {
			continue
		}

		ranges, err := computeRanges(reports)
		if err != nil {
			return nil, time.Time{}, err
		}

		var executedMessages []cciptypes.SeqNumRange
		for _, seqRange := range ranges {
			executedMessagesForRange, err2 := ccipReader.ExecutedMessageRanges(ctx, selector, dest, seqRange)
			if err2 != nil {
				return nil, time.Time{}, err2
			}
			executedMessages = append(executedMessages, executedMessagesForRange...)
		}

		// Remove fully executed reports.
		groupedCommits[selector], err = filterOutExecutedMessages(reports, executedMessages)
		if err != nil {
			return nil, time.Time{}, err
		}
	}

	return groupedCommits, latestReportTS, nil
}

// Observation collects data across two phases which happen in separate rounds.
// These phases happen continuously so that except for the first round, every
// subsequent round can have a new execution report.
//
// Phase 1: Gather commit reports from the destination chain and determine
// which messages are required to build a valid execution report.
//
// Phase 2: Gather messages from the source chains and build the execution
// report.
func (p *Plugin) Observation(
	ctx context.Context, outctx ocr3types.OutcomeContext, _ types.Query,
) (types.Observation, error) {
	var err error
	var previousOutcome plugintypes.ExecutePluginOutcome

	if outctx.PreviousOutcome != nil {
		previousOutcome, err = plugintypes.DecodeExecutePluginOutcome(outctx.PreviousOutcome)
		if err != nil {
			return types.Observation{}, fmt.Errorf("unable to decode previous outcome: %w", err)
		}
	}

	// Phase 1: Gather commit reports from the destination chain and determine which messages are required to build a
	//          valid execution report.
	var groupedCommits plugintypes.ExecutePluginCommitObservations
	supportsDest, err := p.supportsDestChain()
	if err != nil {
		return types.Observation{}, fmt.Errorf("unable to determine if the destination chain is supported: %w", err)
	}
	if supportsDest {
		var latestReportTS time.Time
		groupedCommits, latestReportTS, err =
			getPendingExecutedReports(ctx, p.ccipReader, p.cfg.DestChain, time.UnixMilli(p.lastReportTS.Load()))
		if err != nil {
			return types.Observation{}, err
		}
		// Update timestamp to the last report.
		if len(groupedCommits) > 0 {
			p.lastReportTS.Store(latestReportTS.UnixMilli())
		}

		// TODO: truncate grouped commits to a maximum observation size.
		//       Cache everything which is not executed.
	}

	// Phase 2: Gather messages from the source chains and build the execution report.
	messages := make(plugintypes.ExecutePluginMessageObservations)
	if len(previousOutcome.PendingCommitReports) == 0 {
		fmt.Println("TODO: No reports to execute. This is expected after a cold start.")
		// No reports to execute.
		// This is expected after a cold start.
	} else {
		commitReportCache := make(map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages)
		for _, report := range previousOutcome.PendingCommitReports {
			commitReportCache[report.SourceChain] = append(commitReportCache[report.SourceChain], report)
		}

		for selector, reports := range commitReportCache {
			if len(reports) == 0 {
				continue
			}

			ranges, err := computeRanges(reports)
			if err != nil {
				return types.Observation{}, err
			}

			// Read messages for each range.
			for _, seqRange := range ranges {
				msgs, err := p.ccipReader.MsgsBetweenSeqNums(ctx, selector, seqRange)
				if err != nil {
					return nil, err
				}
				for _, msg := range msgs {
					if _, ok := messages[selector]; !ok {
						messages[selector] = make(map[cciptypes.SeqNum]cciptypes.Message)
					}
					messages[selector][msg.Header.SequenceNumber] = msg
				}
			}
		}
	}

	// TODO: Fire off messages for an attestation check service.

	return plugintypes.NewExecutePluginObservation(groupedCommits, messages).Encode()
}

func (p *Plugin) ValidateObservation(
	outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation,
) error {
	decodedObservation, err := plugintypes.DecodeExecutePluginObservation(ao.Observation)
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

func (p *Plugin) ObservationQuorum(outctx ocr3types.OutcomeContext, query types.Query) (ocr3types.Quorum, error) {
	// TODO: should we use f+1 (or less) instead of 2f+1 because it is not needed for security?
	return ocr3types.QuorumFPlusOne, nil
}

// selectReport takes a list of reports in execution order and selects the first reports that fit within the
// maxReportSizeBytes. Individual messages in a commit report may be skipped for various reasons, for example if an
// out-of-order execution is detected or the message requires additional off-chain metadata which is not yet available.
// If there is not enough space in the final report, it may be partially executed by searching for a subset of messages
// which can fit in the final report.
func selectReport(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	encoder cciptypes.ExecutePluginCodec,
	tokenDataReader temp.TokenDataReader,
	commitReports []plugintypes.ExecutePluginCommitDataWithMessages,
	maxReportSizeBytes int,
) ([]cciptypes.ExecutePluginReportSingleChain, []plugintypes.ExecutePluginCommitDataWithMessages, error) {
	// TODO: It may be desirable for this entire function to be an interface so that
	//       different selection algorithms can be used.

	builder := report.NewBuilder(ctx, lggr, hasher, tokenDataReader, encoder, uint64(maxReportSizeBytes), 99)
	// count number of fully executed reports so that they can be removed after iterating the reports.
	//fullyExecuted := 0
	//accumulatedSize := 0
	//accumulatedGas := uint64(0)
	//var finalReports []cciptypes.ExecutePluginReportSingleChain
	var stillPendingReports []plugintypes.ExecutePluginCommitDataWithMessages
	for i, report := range commitReports {
		// Reports at the end may not have messages yet.
		if len(report.Messages) == 0 {
			break
		}

		/*
			execReport, encodedSize, updatedReport, err :=
				buildSingleChainReportMaxSize(ctx, lggr, hasher, tokenDataReader, encoder,
					report,
					maxReportSizeBytes-accumulatedSize)

			// No messages fit into the report, stop adding more reports.
			if errors.Is(err, errEmptyReport) {
				break
			}
			if err != nil {
				return nil, nil, fmt.Errorf("unable to build single chain report: %w", err)
			}
			reports[reportIdx] = updatedReport
			accumulatedSize += encodedSize
			finalReports = append(finalReports, execReport)

			// partially executed report detected, stop adding more reports.
			// TODO: do not break if messages were intentionally skipped.
			if len(updatedReport.Messages) != len(updatedReport.ExecutedMessages) {
				break
			}
			fullyExecuted++
		*/
		var err error
		commitReports[i], err = builder.Add(report)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to add report to builder: %w", err)
		}
		// If the report has not been fully executed, keep it for the next round.
		if len(commitReports[i].Messages) > len(commitReports[i].ExecutedMessages) {
			stillPendingReports = append(stillPendingReports, commitReports[i])
		}
	}

	// Remove reports that are about to be executed.
	/*
		if fullyExecuted == len(commitReports) {
			commitReports = nil
		} else {
			commitReports = reports[fullyExecuted:]
		}
	*/

	// TODO: Move this into build
	/*
		lggr.Infow(
			"selected commit reports for execution report",
			"numReports", len(finalReports),
			"size", accumulatedSize,
			"incompleteReports", len(reports),
			"maxSize", maxReportSizeBytes)
	*/
	execReports, err := builder.Build()
	return execReports, stillPendingReports, err
}

// Outcome collects the reports from the two phases and constructs the final outcome. Part of the outcome is a fully
// formed report that will be encoded for final transmission in the reporting phase.
func (p *Plugin) Outcome(
	outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	decodedObservations, err := decodeAttributedObservations(aos)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to decode observations: %w", err)

	}
	if len(decodedObservations) < p.reportingCfg.F {
		return ocr3types.Outcome{}, fmt.Errorf("below F threshold")
	}

	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to get FChain: %w", err)
	}

	mergedCommitObservations, err := mergeCommitObservations(decodedObservations, fChain)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to merge commit report observations: %w", err)
	}

	mergedMessageObservations, err := mergeMessageObservations(decodedObservations, fChain)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to merge message observations: %w", err)
	}

	observation := plugintypes.NewExecutePluginObservation(
		mergedCommitObservations,
		mergedMessageObservations)

	// flatten commit reports and sort by timestamp.
	var commitReports []plugintypes.ExecutePluginCommitDataWithMessages
	for _, report := range observation.CommitReports {
		commitReports = append(commitReports, report...)
	}
	sort.Slice(commitReports, func(i, j int) bool {
		return commitReports[i].Timestamp.Before(commitReports[j].Timestamp)
	})

	// add messages to their commitReports.
	for i, report := range commitReports {
		report.Messages = nil
		for i := report.SequenceNumberRange.Start(); i <= report.SequenceNumberRange.End(); i++ {
			if msg, ok := observation.Messages[report.SourceChain][i]; ok {
				report.Messages = append(report.Messages, msg)
			}
		}
		commitReports[i].Messages = report.Messages
	}

	// TODO: this function should be pure, a context should not be needed.
	outcomeReports, commitReports, err :=
		selectReport(context.Background(), p.lggr, p.msgHasher, p.reportCodec, p.tokenDataReader,
			commitReports, maxReportSizeBytes) //, p.cfg.BatchGasLimit)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to extract proofs: %w", err)
	}

	execReport := cciptypes.ExecutePluginReport{
		ChainReports: outcomeReports,
	}

	return plugintypes.NewExecutePluginOutcome(commitReports, execReport).Encode()
}

func (p *Plugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	decodedOutcome, err := plugintypes.DecodeExecutePluginOutcome(outcome)
	if err != nil {
		return nil, fmt.Errorf("unable to decode outcome: %w", err)
	}

	// TODO: this function should be pure, a context should not be needed.
	encoded, err := p.reportCodec.Encode(context.Background(), decodedOutcome.Report)
	if err != nil {
		return nil, fmt.Errorf("unable to encode report: %w", err)
	}

	report := []ocr3types.ReportWithInfo[[]byte]{{
		Report: encoded,
		Info:   nil,
	}}

	return report, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	if len(decodedReport.ChainReports) == 0 {
		p.lggr.Infow("skipping empty report")
		return false, nil
	}
	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	isWriter, err := p.supportsDestChain()
	if err != nil {
		return false, fmt.Errorf("unable to determine if the destination chain is supported: %w", err)
	}
	if !isWriter {
		p.lggr.Debugw("not a destination writer, skipping report transmission")
		return false, nil
	}

	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	// TODO: Final validation?

	p.lggr.Debugw("transmitting report",
		"reports", len(decodedReport.ChainReports),
	)
	return true, nil
}

func (p *Plugin) Close() error {
	return nil
}

func (p *Plugin) supportedChains(id commontypes.OracleID) (mapset.Set[cciptypes.ChainSelector], error) {
	p2pID, exists := p.oracleIDToP2pID[id]
	if !exists {
		return nil, fmt.Errorf("oracle ID %d not found in oracleIDToP2pID", p.reportingCfg.OracleID)
	}
	supportedChains, err := p.homeChain.GetSupportedChainsForPeer(p2pID)
	if err != nil {
		p.lggr.Warnw("error getting supported chains", err)
		return mapset.NewSet[cciptypes.ChainSelector](), fmt.Errorf("error getting supported chains: %w", err)
	}

	return supportedChains, nil
}

func (p *Plugin) supportsDestChain() (bool, error) {
	chains, err := p.supportedChains(p.reportingCfg.OracleID)
	if err != nil {
		return false, fmt.Errorf("error getting supported chains: %w", err)
	}
	return chains.Contains(p.cfg.DestChain), nil
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
