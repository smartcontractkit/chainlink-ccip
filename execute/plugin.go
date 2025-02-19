package execute

import (
	"bytes"
	"cmp"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"slices"
	"sort"
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
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/cache"
	"github.com/smartcontractkit/chainlink-ccip/execute/metrics"
	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

var (
	errAlreadyExecuted = errors.New("messages already executed")
)

type ContractDiscoveryInterface plugincommon.PluginProcessor[dt.Query, dt.Observation, dt.Outcome]

type inflightMessageCache interface {
	IsInflight(src cciptypes.ChainSelector, msgID cciptypes.Bytes32) bool
	MarkInflight(src cciptypes.ChainSelector, msgID cciptypes.Bytes32)
	Delete(src cciptypes.ChainSelector, msgID cciptypes.Bytes32)
}

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
	discovery    ContractDiscoveryInterface
	chainSupport plugincommon.ChainSupport
	observer     metrics.Reporter

	oracleIDToP2pID       map[commontypes.OracleID]libocrtypes.PeerID
	tokenDataObserver     tokendata.TokenDataObserver
	costlyMessageObserver costlymessages.Observer
	estimateProvider      cciptypes.EstimateProvider
	lggr                  logger.Logger
	ocrTypeCodec          ocrtypecodec.ExecCodec

	// state

	contractsInitialized bool
	// commitRootsCache remembers commit root details to optimize DB lookups.
	commitRootsCache cache.CommitsRootsCache
	// inflightMessageCache prevents duplicate reports from being sent for the same message.
	inflightMessageCache inflightMessageCache
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
	estimateProvider cciptypes.EstimateProvider,
	lggr logger.Logger,
	costlyMessageObserver costlymessages.Observer,
	metricsReporter metrics.Reporter,
) ocr3types.ReportingPlugin[[]byte] {
	lggr.Infow("creating new plugin instance", "p2pID", oracleIDToP2pID[reportingCfg.OracleID])

	p := Plugin{
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
		lggr:                  logutil.WithComponent(lggr, "ExecutePlugin"),
		costlyMessageObserver: costlyMessageObserver,
		discovery: discovery.NewContractDiscoveryProcessor(
			logutil.WithComponent(lggr, "Discovery"),
			&ccipReader,
			homeChain,
			destChain,
			reportingCfg.F,
			oracleIDToP2pID,
		),
		chainSupport: plugincommon.NewChainSupport(
			logutil.WithComponent(lggr, "ChainSupport"),
			homeChain,
			oracleIDToP2pID,
			reportingCfg.OracleID,
			destChain,
		),
		observer: metricsReporter,
		commitRootsCache: cache.NewCommitRootsCache(
			logutil.WithComponent(lggr, "CommitRootsCache"),
			offchainCfg.MessageVisibilityInterval.Duration(),
			time.Minute*5),
		inflightMessageCache: cache.NewInflightMessageCache(offchainCfg.InflightCacheExpiry.Duration()),
		ocrTypeCodec:         ocrtypecodec.NewExecCodecJSON(),
	}
	return NewTrackedPlugin(p, metricsReporter)
}

func (p *Plugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

type CanExecuteHandle = func(sel cciptypes.ChainSelector, merkleRoot cciptypes.Bytes32) bool

// getPendingExecutedReports is used to find commit reports which need to be executed.
// It considers all commit reports as of the given timestamp. Of the reports found, the
// provided canExecute function is used to filter out reports which the caller knows to
// be ineligible (i.e. already executed, or snoozed for some reason). The final step
// is to check their execution state to see if the messages for each report are already
// executed. Any fully executed reports are returned separately for the caller to remember.
func getPendingExecutedReports(
	ctx context.Context,
	ccipReader readerpkg.CCIPReader,
	canExecute CanExecuteHandle,
	ts time.Time,
	lggr logger.Logger,
) (exectypes.CommitObservations, []exectypes.CommitData /* fully executed roots */, error) {
	var fullyExecuted []exectypes.CommitData
	commitReports, err := ccipReader.CommitReportsGTETimestamp(ctx, ts, 1000) // todo: configurable limit
	if err != nil {
		return nil, nil, err
	}
	lggr.Debugw("commit reports", "commitReports", commitReports, "count", len(commitReports))

	groupedCommits := groupByChainSelector(commitReports)
	lggr.Debugw("grouped commits before removing fully executed reports",
		"groupedCommits", groupedCommits, "count", len(groupedCommits))

	// Remove fully executed reports.
	for selector, reports := range groupedCommits {
		if len(reports) == 0 {
			continue
		}

		// Filter out reports that cannot be executed (executed or snoozed).
		var filtered []exectypes.CommitData
		{
			var skippedCommitRoots []string
			for _, commitReport := range reports {
				if !canExecute(commitReport.SourceChain, commitReport.MerkleRoot) {
					skippedCommitRoots = append(skippedCommitRoots, commitReport.MerkleRoot.String())
					continue
				}
				filtered = append(filtered, commitReport)
			}
			lggr.Infow(
				"skipping reports marked as executed or snoozed",
				"selector", selector,
				"skippedCommitRoots", skippedCommitRoots,
			)
		}
		reports = filtered

		lggr.Debugw("grouped reports", "selector", selector, "reports", reports, "count", len(reports))
		sort.Slice(reports, func(i, j int) bool {
			return reports[i].SequenceNumberRange.Start() < reports[j].SequenceNumberRange.Start()
		})
		// todo: remove this logs after investigating whether the sorting above can be safely removed
		lggr.Debugw("sorted reports", "selector", selector, "reports", reports, "count", len(reports))

		ranges, err := computeRanges(reports)
		if err != nil {
			return nil, nil, fmt.Errorf("compute report ranges: %w", err)
		}

		executedMessageSet := mapset.NewSet[cciptypes.SeqNum]()
		for _, seqRange := range ranges {
			executedMessagesForRange, err2 := ccipReader.ExecutedMessages(ctx, selector, seqRange)
			if err2 != nil {
				return nil, nil, fmt.Errorf("get %d executed messages in range %v: %w", selector, seqRange, err2)
			}
			executedMessageSet = executedMessageSet.Union(mapset.NewSet(executedMessagesForRange...))
		}

		executedMessages := executedMessageSet.ToSlice()
		sort.Slice(executedMessages, func(i, j int) bool { return executedMessages[i] < executedMessages[j] })

		// Remove fully executed reports.
		// Populate executed messages on the reports.
		var executedCommits []exectypes.CommitData
		groupedCommits[selector], executedCommits = combineReportsAndMessages(reports, executedMessages)
		fullyExecuted = append(fullyExecuted, executedCommits...)
	}

	lggr.Debugw("grouped commits after removing fully executed reports",
		"groupedCommits", groupedCommits, "count", len(groupedCommits))

	return groupedCommits, fullyExecuted, nil
}

func (p *Plugin) ValidateObservation(
	_ context.Context, outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation,
) error {
	decodedObservation, err := p.ocrTypeCodec.DecodeObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("unable to decode observation: %w", err)
	}

	var previousOutcome exectypes.Outcome

	previousOutcome, err = p.ocrTypeCodec.DecodeOutcome(outctx.PreviousOutcome)
	if err != nil {
		return fmt.Errorf("unable to decode previous outcome: %w", err)
	}

	supportedChains, err := p.supportedChains(ao.Observer)
	if err != nil {
		return fmt.Errorf("error finding supported chains by node: %w", err)
	}

	state := previousOutcome.State.Next()
	if state == exectypes.Initialized || state == exectypes.GetCommitReports {
		err = validateNoMessageRelatedObservations(
			decodedObservation.Messages,
			decodedObservation.TokenData,
			decodedObservation.Hashes,
			decodedObservation.CostlyMessages,
		)
		if err != nil {
			return err
		}
	}

	if err = validateCommonStateObservations(p, ao.Observer, decodedObservation, supportedChains); err != nil {
		return err
	}

	// check message related validations when states can contain messages
	if state == exectypes.GetMessages || state == exectypes.Filter {
		if err = validateMsgsReadingEligibility(supportedChains, decodedObservation.Messages); err != nil {
			return fmt.Errorf("validate observer reading eligibility: %w", err)
		}

		err = validateMessagesRelatedObservations(
			decodedObservation.CommitReports,
			decodedObservation.Messages,
			decodedObservation.TokenData,
			decodedObservation.Hashes,
			decodedObservation.CostlyMessages,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateCommonStateObservations(
	p *Plugin,
	oracleID commontypes.OracleID,
	decodedObservation exectypes.Observation,
	supportedChains mapset.Set[cciptypes.ChainSelector],
) error {
	if err := plugincommon.ValidateFChain(decodedObservation.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	// These checks are common to all states.
	if err := validateCommitReportsReadingEligibility(supportedChains, decodedObservation.CommitReports); err != nil {
		return fmt.Errorf("validate commit reports reading eligibility: %w", err)
	}

	if err := validateObservedSequenceNumbers(decodedObservation.CommitReports); err != nil {
		return fmt.Errorf("validate observed sequence numbers: %w", err)
	}

	if p.discovery != nil {
		aos := plugincommon.AttributedObservation[dt.Observation]{
			OracleID:    oracleID,
			Observation: decodedObservation.Contracts,
		}

		if err := p.discovery.ValidateObservation(dt.Outcome{}, dt.Query{}, aos); err != nil {
			return fmt.Errorf("process contracts: %w", err)
		}
	}

	return nil
}

func validateNoMessageRelatedObservations(
	messages exectypes.MessageObservations,
	tokenData exectypes.TokenDataObservations,
	hashes exectypes.MessageHashes,
	costlyMessages []cciptypes.Bytes32,
) error {
	if len(messages) > 0 {
		return fmt.Errorf("messages are not expected in initial or GetCommitRerports states")
	}
	if len(tokenData) > 0 {
		return fmt.Errorf("token data is not expected in initial or GetCommitRerports states")
	}
	if len(costlyMessages) > 0 {
		return fmt.Errorf("costly messages are not expected in initial or GetCommitRerports states")
	}
	if len(hashes) > 0 {
		return fmt.Errorf("hashes are not expected in initial or GetCommitRerports states")
	}

	return nil
}

func validateMessagesRelatedObservations(
	commitReports exectypes.CommitObservations,
	messages exectypes.MessageObservations,
	tokenData exectypes.TokenDataObservations,
	hashes exectypes.MessageHashes,
	costlyMessages []cciptypes.Bytes32,
) error {

	if err := validateMessagesConformToCommitReports(commitReports, messages); err != nil {
		return fmt.Errorf("validate messages conform to commit reports: %w", err)
	}
	if err := validateHashesExist(messages, hashes); err != nil {
		return fmt.Errorf("validate hashes exist: %w", err)
	}
	if err := validateTokenDataObservations(messages, tokenData); err != nil {
		return fmt.Errorf("validate token data observations: %w", err)
	}
	if err := validateCostlyMessagesObservations(messages, costlyMessages); err != nil {
		return fmt.Errorf("validate costly messages: %w", err)
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

	pendingReports := 0
	for i, commitReport := range commitReports {
		// handle incomplete observations.
		if len(commitReport.Messages) == 0 {
			pendingReports++
			continue
		}

		var err error
		// The builder may attach metadata to the commit report.
		commitReports[i], err = builder.Add(ctx, commitReport)
		if err != nil {
			pendingReports++
			lggr.Errorw("unable to add report to builder", "err", err)
			continue
		}

		// If the report has not been fully executed, keep it for the next round.
		// Detect a report was not fully executed
		if len(commitReports[i].Messages) > len(commitReports[i].ExecutedMessages) {
			pendingReports++
		}
	}

	execReports, selectedReports, err := builder.Build()

	lggr.Infow(
		"reports have been selected",
		"numReports", len(execReports),
		"numPendingReports", pendingReports)
	return execReports, selectedReports, err
}

func extractReportInfo(report exectypes.Outcome) cciptypes.ExecuteReportInfo {
	merkleRoots := []cciptypes.MerkleRootChain{}

	for _, commitReport := range report.CommitReports {
		merkleRoots = append(merkleRoots, cciptypes.MerkleRootChain{
			ChainSel:      commitReport.SourceChain,
			OnRampAddress: commitReport.OnRampAddress,
			SeqNumsRange:  commitReport.SequenceNumberRange,
			MerkleRoot:    commitReport.MerkleRoot,
		})
	}

	return cciptypes.ExecuteReportInfo{
		AbstractReports: report.Report.ChainReports,
		MerkleRoots:     merkleRoots,
	}
}

func (p *Plugin) Reports(
	ctx context.Context, seqNr uint64, outcome ocr3types.Outcome,
) ([]ocr3types.ReportPlus[[]byte], error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseReports)

	if outcome == nil {
		lggr.Warn("no outcome, skipping report generation")
		return nil, nil
	}

	decodedOutcome, err := p.ocrTypeCodec.DecodeOutcome(outcome)
	if err != nil {
		return nil, fmt.Errorf("unable to decode outcome: %w", err)
	}

	encodedReport, err := p.reportCodec.Encode(ctx, decodedOutcome.Report)
	if err != nil {
		return nil, fmt.Errorf("unable to encode report: %w", err)
	}

	reportInfo := extractReportInfo(decodedOutcome)
	encodedInfo, err := reportInfo.Encode()
	if err != nil {
		return nil, err
	}

	transmissionSchedule, err := plugincommon.GetTransmissionSchedule(
		p.chainSupport,
		maps.Keys(p.oracleIDToP2pID),
		p.offchainCfg.TransmissionDelayMultiplier,
	)
	if err != nil {
		return nil, fmt.Errorf("get transmission schedule: %w", err)
	}
	lggr.Debugw("transmission schedule override",
		"transmissionSchedule", transmissionSchedule, "oracleIDToP2PID", p.oracleIDToP2pID)

	r := []ocr3types.ReportPlus[[]byte]{
		{
			ReportWithInfo: ocr3types.ReportWithInfo[[]byte]{
				Report: encodedReport,
				Info:   encodedInfo,
			},
			TransmissionScheduleOverride: transmissionSchedule,
		},
	}

	return r, nil
}

// validateReport validates various aspects of the report.
// Pure checks are placed earlier in the function on purpose to avoid
// unnecessary network or DB I/O.
// If you're added more checks make sure to follow this pattern.
func (p *Plugin) validateReport(
	ctx context.Context,
	lggr logger.Logger,
	r ocr3types.ReportWithInfo[[]byte],
) (valid bool, decodedReport cciptypes.ExecutePluginReport, err error) {
	// Just a safety check, should never happen.
	if r.Report == nil {
		lggr.Warn("skipping nil report")
		return false, cciptypes.ExecutePluginReport{}, nil
	}

	decodedReport, err = p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, cciptypes.ExecutePluginReport{}, fmt.Errorf("decode exec plugin report: %w", err)
	}

	if len(decodedReport.ChainReports) == 0 {
		lggr.Infow("skipping empty report")
		return false, cciptypes.ExecutePluginReport{}, nil
	}

	// check if we support the dest, if not we can't do the checks needed.
	supports, err := p.chainSupport.SupportsDestChain(p.reportingCfg.OracleID)
	if err != nil {
		return false, cciptypes.ExecutePluginReport{}, fmt.Errorf("supports dest chain: %w", err)
	}

	if !supports {
		lggr.Warnw("dest chain not supported, can't run report acceptance procedures")
		return false, cciptypes.ExecutePluginReport{}, nil
	}

	offRampConfigDigest, err := p.ccipReader.GetOffRampConfigDigest(ctx, consts.PluginTypeExecute)
	if err != nil {
		return false, cciptypes.ExecutePluginReport{}, fmt.Errorf("get offramp config digest: %w", err)
	}

	if !bytes.Equal(offRampConfigDigest[:], p.reportingCfg.ConfigDigest[:]) {
		lggr.Warnw("my config digest doesn't match offramp's config digest, not accepting/transmitting report",
			"myConfigDigest", p.reportingCfg.ConfigDigest,
			"offRampConfigDigest", hex.EncodeToString(offRampConfigDigest[:]),
		)
		return false, cciptypes.ExecutePluginReport{}, nil
	}

	// check that the messages in the report are not already executed onchain.
	// note that this involves a set of DB queries, hence why its last in the checks.
	seqNrRangesBySource := getSnRangeSetPairsBySource(decodedReport.ChainReports)
	err = p.checkAlreadyExecuted(ctx, lggr, seqNrRangesBySource)
	if errors.Is(err, errAlreadyExecuted) {
		// Some messages in the report have already
		// been executed, so we don't want to re-execute them.
		// This gives the exec plugin a chance to remedy the situation
		// by selecting a different set of messages.
		return false, decodedReport, nil
	}
	if err != nil {
		// TODO: should we return true here if we couldn't check for already executed messages?
		return false, decodedReport, fmt.Errorf("checking for already executed messages failed: %w", err)
	}

	return true, decodedReport, nil
}

// checkAlreadyExecuted checks if the messages in the report have already been executed
// on the destination chain. It queries the DB for executed messages in the given sequence
// number range for each source chain in the report.
func (p *Plugin) checkAlreadyExecuted(
	ctx context.Context,
	lggr logger.Logger,
	seqNrRangesBySource map[cciptypes.ChainSelector]snRangeSetPair,
) error {
	// TODO: batch these queries? these are all DB reads.
	// maybe some alternative queries exist.
	for sourceChainSelector, seqNrRange := range seqNrRangesBySource {
		executed, err := p.ccipReader.ExecutedMessages(ctx, sourceChainSelector, seqNrRange.snRange)
		if err != nil {
			return fmt.Errorf("couldn't check if messages already executed: %w", err)
		}

		if intersection := mapset.NewSet(executed...).Intersect(seqNrRange.set); !intersection.IsEmpty() {
			// Some messages in the report have been executed, don't accept it.
			reportSeqNrsSlice := seqNrRange.set.ToSlice()
			lggr.Warnw("some messages in report already executed",
				"alreadyExecuted", executed,
				"reportSeqNrs", reportSeqNrsSlice,
			)
			return fmt.Errorf("%w: already executed messages %+v report seq nrs %+v",
				errAlreadyExecuted, executed, reportSeqNrsSlice)
		}
	}

	return nil
}

// snRangeSetPair is an internal data structure used to store a sequence number range
// and a set of sequence numbers simultaneously.
type snRangeSetPair struct {
	// snRange is a range of [min, max] sequence numbers of the messages in the report for a particular source chain.
	// it is used to query the CCIPReader for executed messages.
	snRange cciptypes.SeqNumRange
	// set is the sequence numbers of the messages in the report for a particular source chain.
	// it is used to check whether the returned array from CCIPReader has a non-empty intersection
	// with the set of sequence numbers in the report.
	set mapset.Set[cciptypes.SeqNum]
}

// getSnRangeSetPairsBySource returns a map of source chain selector to
// the sequence number range of the messages in the report.
func getSnRangeSetPairsBySource(
	chainReports []cciptypes.ExecutePluginReportSingleChain,
) map[cciptypes.ChainSelector]snRangeSetPair {
	seqNrRangesBySource := make(map[cciptypes.ChainSelector]snRangeSetPair)
	for _, chainReport := range chainReports {
		// This should never happen, indicates a bug in the report building and accepting process.
		// But we sanity check since slices.Min/Max will panic on empty slices.
		if len(chainReport.Messages) == 0 {
			continue
		}

		cmpr := func(a, b cciptypes.Message) int {
			return cmp.Compare(a.Header.SequenceNumber, b.Header.SequenceNumber)
		}
		minMsg := slices.MinFunc(chainReport.Messages, cmpr)
		maxMsg := slices.MaxFunc(chainReport.Messages, cmpr)
		seqNrSet := mapset.NewSet(
			slicelib.Map(
				chainReport.Messages,
				func(msg cciptypes.Message) cciptypes.SeqNum {
					return msg.Header.SequenceNumber
				},
			)...,
		)
		seqNrRangesBySource[chainReport.SourceChainSelector] = snRangeSetPair{
			snRange: cciptypes.NewSeqNumRange(
				minMsg.Header.SequenceNumber,
				maxMsg.Header.SequenceNumber),
			set: seqNrSet,
		}
	}
	return seqNrRangesBySource
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseShouldAccept)

	valid, decodedReport, err := p.validateReport(ctx, lggr, r)
	if err != nil {
		return false, fmt.Errorf("validate exec report: %w", err)
	}

	if !valid {
		lggr.Infow("report is not accepted", "seqNr", seqNr)
		return false, nil
	}

	// TODO: consider doing this in validateReport,
	// will end up doing it in both ShouldAccept and ShouldTransmit.
	sourceChains := slicelib.Map(decodedReport.ChainReports,
		func(r cciptypes.ExecutePluginReportSingleChain) cciptypes.ChainSelector {
			return r.SourceChainSelector
		})
	isCursed, err := plugincommon.IsReportCursed(ctx, lggr, p.ccipReader, p.chainSupport.DestChain(), sourceChains)
	if err != nil {
		lggr.Errorw(
			"report not accepted due to curse checking error",
			"err", err,
		)
		return false, err
	}
	if isCursed {
		// Detailed logging is already done by IsReportCursed.
		return false, nil
	}

	lggr.Infow("ShouldAcceptAttestedReport returns true, report accepted",
		"seqNr", seqNr,
		"reports", decodedReport.ChainReports,
	)
	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseShouldTransmit)

	valid, decodedReport, err := p.validateReport(ctx, lggr, r)
	if err != nil {
		return valid, fmt.Errorf("validate exec report: %w", err)
	}

	if !valid {
		lggr.Infow("report not accepted for transmit")
		return false, nil
	}

	lggr.Infow("ShouldTransmitAttestedReport returns true, report accepted",
		"reports", decodedReport.ChainReports,
	)
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
