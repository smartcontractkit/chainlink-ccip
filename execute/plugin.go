package execute

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
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

	"github.com/smartcontractkit/chainlink-common/pkg/services"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/cache"
	"github.com/smartcontractkit/chainlink-ccip/execute/metrics"
	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/observer"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
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

	oracleIDToP2pID   map[commontypes.OracleID]libocrtypes.PeerID
	tokenDataObserver observer.TokenDataObserver
	estimateProvider  cciptypes.EstimateProvider
	lggr              logger.Logger
	ocrTypeCodec      ocrtypecodec.ExecCodec
	addrCodec         cciptypes.AddressCodec

	// state
	contractsInitialized bool
	commitRootsCache     cache.CommitsRootsCache
	commitReportCache    cache.CommitReportCache
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
	tokenDataObserver observer.TokenDataObserver,
	estimateProvider cciptypes.EstimateProvider,
	lggr logger.Logger,
	metricsReporter metrics.Reporter,
	addrCodec cciptypes.AddressCodec,
) ocr3types.ReportingPlugin[[]byte] {
	lggr.Infow("creating new plugin instance", "p2pID", oracleIDToP2pID[reportingCfg.OracleID])

	ocrTypCodec := ocrtypecodec.DefaultExecCodec

	// Initialize CommitReportCacheConfig
	commitReportCacheCfg := cache.CommitReportCacheConfig{
		MessageVisibilityInterval: offchainCfg.MessageVisibilityInterval.Duration(),
		EvictionGracePeriod:       cache.EvictionGracePeriod,
		CleanupInterval:           cache.CleanupInterval,
	}

	p := &Plugin{
		donID:             donID,
		reportingCfg:      reportingCfg,
		offchainCfg:       offchainCfg,
		destChain:         destChain,
		oracleIDToP2pID:   oracleIDToP2pID,
		ccipReader:        ccipReader,
		reportCodec:       reportCodec,
		msgHasher:         msgHasher,
		homeChain:         homeChain,
		tokenDataObserver: tokenDataObserver,
		estimateProvider:  estimateProvider,
		lggr:              logutil.WithComponent(lggr, "ExecutePlugin"),
		discovery: discovery.NewContractDiscoveryProcessor(
			logutil.WithComponent(lggr, "Discovery"),
			&ccipReader,
			homeChain,
			destChain,
			reportingCfg.F,
			oracleIDToP2pID,
			reportingCfg.OracleID,
			metricsReporter,
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
			offchainCfg.RootSnoozeTime.Duration(),
		),
		commitReportCache: cache.NewCommitReportCache(
			logutil.WithComponent(lggr, "CommitReportCache"),
			commitReportCacheCfg,
			&cache.RealTimeProvider{},
			ccipReader,
		),
		inflightMessageCache: cache.NewInflightMessageCache(offchainCfg.InflightCacheExpiry.Duration()),
		ocrTypeCodec:         ocrTypCodec,
		addrCodec:            addrCodec,
	}
	return NewTrackedPlugin(p, lggr, metricsReporter, ocrTypCodec)
}

func (p *Plugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

type CanExecuteHandle = func(sel cciptypes.ChainSelector, merkleRoot cciptypes.Bytes32) bool

// getSortedExecutableReports returns reports that can be executed based on the canExecute function
func getSortedExecutableReports(lggr logger.Logger,
	selector cciptypes.ChainSelector,
	reports []exectypes.CommitData,
	canExecute CanExecuteHandle) []exectypes.CommitData {
	// Filter out reports that cannot be executed (executed or snoozed).
	var executableReports []exectypes.CommitData
	{
		var skippedCommitRoots []string
		for _, commitReport := range reports {
			if !canExecute(commitReport.SourceChain, commitReport.MerkleRoot) {
				skippedCommitRoots = append(skippedCommitRoots, commitReport.MerkleRoot.String())
				continue
			}
			executableReports = append(executableReports, commitReport)
		}
		lggr.Infow(
			"skipping reports marked as executed or snoozed",
			"selector", selector,
			"skippedCommitRoots", skippedCommitRoots,
		)
	}
	sort.Slice(executableReports, func(i, j int) bool {
		return exectypes.LessThan(executableReports[i], executableReports[j])
	})

	return executableReports
}

// getExecutableReportRanges returns the ranges of reports that can be executed based on the canExecute function
func getExecutableReportRanges(lggr logger.Logger,
	groupedCommits exectypes.CommitObservations,
	canExecute CanExecuteHandle,
) (
	map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	map[cciptypes.ChainSelector][]exectypes.CommitData,
	error) {
	rangesBySelector := make(map[cciptypes.ChainSelector][]cciptypes.SeqNumRange)
	executableReports := make(map[cciptypes.ChainSelector][]exectypes.CommitData)

	for selector, reports := range groupedCommits {
		if len(reports) == 0 {
			continue
		}

		executableReports[selector] = getSortedExecutableReports(lggr, selector, reports, canExecute)

		lggr.Debugw("grouped and sorted",
			"selector", selector,
			"reports", executableReports[selector],
			"count", len(executableReports[selector]))

		ranges, err := computeRanges(executableReports[selector])
		if err != nil {
			return rangesBySelector, executableReports, fmt.Errorf("compute report ranges: %w", err)
		}
		rangesBySelector[selector] = ranges
	}
	return rangesBySelector, executableReports, nil
}

// removeUnconfirmedAndFinalizedMessages removes messages that have been executed with finality or are unconfirmed.
// It returns the remaining reports, fully executed finalized reports, and fully executed unfinalized reports
func removeUnconfirmedAndFinalizedMessages(
	executableReports map[cciptypes.ChainSelector][]exectypes.CommitData,
	finalizedMessages, unconfirmedMessages map[cciptypes.ChainSelector][]cciptypes.SeqNum,
) (
	remainingReportsBySelector map[cciptypes.ChainSelector][]exectypes.CommitData,
	fullyExecutedFinalized []exectypes.CommitData,
	fullyExecutedUnfinalized []exectypes.CommitData,
) {
	remainingReportsBySelector = make(map[cciptypes.ChainSelector][]exectypes.CommitData)
	for selector, reports := range executableReports {
		unconfirmedMsgSet := mapset.NewSet(unconfirmedMessages[selector]...)
		finalizedMsgSet := mapset.NewSet(finalizedMessages[selector]...)
		// Get unfinalized messages by taking the difference
		unfinalizedMessages := slicelib.ToSortedSlice(unconfirmedMsgSet.Difference(finalizedMsgSet))

		sortedFinalizedMessages := slicelib.ToSortedSlice(finalizedMsgSet)

		// Fully finalized roots are removed from the reports and set in groupedCommits
		remainingReports, executedCommitsFinalized := combineReportsAndMessages(reports, sortedFinalizedMessages)
		fullyExecutedFinalized = append(fullyExecutedFinalized, executedCommitsFinalized...)

		// Process unfinalized messages
		finalRemainingReports, executedCommitsUnfinalized := combineReportsAndMessages(remainingReports, unfinalizedMessages)
		fullyExecutedUnfinalized = append(fullyExecutedUnfinalized, executedCommitsUnfinalized...)

		// Update groupedCommits with the remaining reports
		remainingReportsBySelector[selector] = finalRemainingReports
	}
	return
}

// getPendingReportsForExecution is used to find commit reports which need to be executed.
//
// The function leverages an optimized caching strategy for commit reports:
// 1. Uses CommitReportCache to determine the optimal timestamp to query from
// 2. Retrieves previously cached reports (filtered to only contain those with Merkle roots)
// 3. Fetches incremental reports from the chain reader
// 4. Deduplicates combined reports to prevent redundant processing
//
// Execution status is checked at two levels:
// 1. Gets all executed messages (both finalized and unfinalized) via primitives.Unconfirmed
// 2. Gets only finalized executed messages via primitives.Finalized
//
// Reports are then classified into three categories:
// - fullyExecutedFinalized: All messages executed with finality (permanently marked as executed)
// - fullyExecutedUnfinalized: All messages executed but not finalized (temporarily snoozed)
// - groupedCommits: Reports with unexecuted messages (available for execution)
func getPendingReportsForExecution(
	ctx context.Context,
	ccipReader readerpkg.CCIPReader,
	commitReportCache cache.CommitReportCache,
	canExecute CanExecuteHandle,
	fetchFrom time.Time,
	cursedSourceChains map[cciptypes.ChainSelector]bool,
	limit int,
	lggr logger.Logger,
) (
	groupedCommits exectypes.CommitObservations,
	fullyExecutedFinalized []exectypes.CommitData,
	fullyExecutedUnfinalized []exectypes.CommitData,
	err error,
) {
	// get the reports from the cache
	reportsFromCache := commitReportCache.GetCachedReports(fetchFrom)
	lggr.Infow("Querying commit reports from cache", "fetchFrom", fetchFrom)

	// get the timestamp to fetch from cache
	queryFromTimestampForNewReports := commitReportCache.GetReportsToQueryFromTimestamp()

	// Assuming each report can have minimum one message, max reports shouldn't exceed the max messages
	unfinalizedIncrementReports, err := ccipReader.CommitReportsGTETimestamp(
		ctx, queryFromTimestampForNewReports, primitives.Unconfirmed, limit,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	lggr.Infow("Querying commit reports from chain", "fetchFrom", queryFromTimestampForNewReports)

	// merge the reports from the cache and the reports from the reader
	candidateReports := cache.DeduplicateReports(lggr, append(reportsFromCache, unfinalizedIncrementReports...))

	lggr.Debugw("commit reports", "candidateReports", candidateReports, "count", len(candidateReports))

	groupedCommits = groupByChainSelectorWithFilter(lggr, candidateReports, cursedSourceChains)
	lggr.Debugw("grouped commits before removing fully executed reports",
		"groupedCommits", groupedCommits, "count", len(groupedCommits))

	rangesBySelector, executableReports, err := getExecutableReportRanges(lggr, groupedCommits, canExecute)
	if err != nil {
		return nil, nil, nil, err
	}

	// Get all executed messages
	unconfirmedMessages, err := ccipReader.ExecutedMessages(ctx, rangesBySelector, primitives.Unconfirmed)
	if err != nil {
		return nil, nil, nil,
			fmt.Errorf("get executed messages in range %v: %w", rangesBySelector, err)
	}
	// Get finalized messages
	finalizedMessages, err := ccipReader.ExecutedMessages(ctx, rangesBySelector, primitives.Finalized)
	if err != nil {
		return nil, nil, nil,
			fmt.Errorf("get finalized executed messages in range %v: %w", rangesBySelector, err)
	}

	remainingReportsBySelector, fullyExecutedFinalized, fullyExecutedUnfinalized :=
		removeUnconfirmedAndFinalizedMessages(executableReports, finalizedMessages, unconfirmedMessages)
	lggr.Debugw("grouped commits after removing fully executed reports",
		"remainingReportsBySelector", remainingReportsBySelector,
		"countFinalized", len(fullyExecutedFinalized),
		"countUnfinalized", len(fullyExecutedUnfinalized))

	return remainingReportsBySelector,
		fullyExecutedFinalized,
		fullyExecutedUnfinalized,
		nil
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

	nextState := previousOutcome.State.Next()
	if nextState == exectypes.GetCommitReports {
		err = validateNoMessageRelatedObservations(
			decodedObservation.Messages,
			decodedObservation.TokenData,
			decodedObservation.Hashes,
		)
		if err != nil {
			return err
		}
	}

	state := previousOutcome.State.Next()

	if err = validateCommonStateObservations(p, ao.Observer, decodedObservation, supportedChains, state); err != nil {
		return err
	}

	// check message related validations when states can contain messages
	if nextState == exectypes.GetMessages || nextState == exectypes.Filter {
		if err = validateMsgsReadingEligibility(supportedChains, decodedObservation.Messages); err != nil {
			return fmt.Errorf("validate observer reading eligibility: %w", err)
		}

		err = validateMessagesRelatedObservations(
			decodedObservation.CommitReports,
			decodedObservation.Messages,
			decodedObservation.TokenData,
			decodedObservation.Hashes,
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
	state exectypes.PluginState,
) error {
	if err := plugincommon.ValidateFChain(decodedObservation.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	// These checks are common to all states.
	if err := validateCommitReportsReadingEligibility(
		state, supportedChains, p.destChain, decodedObservation.CommitReports); err != nil {
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
) error {
	if len(messages) > 0 {
		return fmt.Errorf("messages are not expected in initial or GetCommitRerports states")
	}
	if len(tokenData) > 0 {
		return fmt.Errorf("token data is not expected in initial or GetCommitRerports states")
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

	return nil
}

func (p *Plugin) ObservationQuorum(
	_ context.Context, outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (bool, error) {
	// 2f+1 is required for fChain consensus.
	return quorumhelper.ObservationCountReachesObservationQuorum(
		quorumhelper.QuorumTwoFPlusOne, p.reportingCfg.N, p.reportingCfg.F, aos), nil
}

// selectReports takes a list of reports in execution order and selects the first reports that fit within the
// maxReportSizeBytes. Individual messages in a commit report may be skipped for various reasons, for example if an
// out-of-order execution is detected or the message requires additional off-chain metadata which is not yet available.
// If there is not enough space in the final report, it may be partially executed by searching for a subset of messages
// which can fit in the final report.
//
// The output of this function are canonical data structures which can be directly included in the final outcome.
func selectReports(
	ctx context.Context,
	lggr logger.Logger,
	commitReports []exectypes.CommitData,
	builder report.ExecReportBuilder,
) ([]cciptypes.ExecutePluginReport, [][]exectypes.CommitData, error) {
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

		// If the commit report has not been fully executed, keep it for the next round.
		// Detect a report was not fully executed
		if len(commitReports[i].Messages) > len(commitReports[i].ExecutedMessages) {
			pendingReports++
		}
	}

	execReports, selectedReports, err := builder.Build()

	// TODO: count pending reports during Build() when we select which reports can be returned.

	lggr.Debugw("selected report to be executed", "reports", selectedReports)
	lggr.Infow(
		"reports have been selected",
		"numReports", len(execReports),
		"numPendingReports", pendingReports)
	return execReports, selectedReports, err
}

func extractReportInfo(report exectypes.Outcome) ([]cciptypes.ExecuteReportInfo, error) {
	var ri []cciptypes.ExecuteReportInfo
	commitReportIdx := 0

	for _, execReport := range report.Reports {
		var merkleRoots []cciptypes.MerkleRootChain

		for i := 0; i < len(execReport.ChainReports); i++ {
			if commitReportIdx >= len(report.CommitReports) {
				return nil, fmt.Errorf("not enough commit reports for the "+
					"number of chain reports in exec report %d", i)
			}
			cr := report.CommitReports[commitReportIdx]
			merkleRoots = append(merkleRoots, cciptypes.MerkleRootChain{
				ChainSel:      cr.SourceChain,
				OnRampAddress: cr.OnRampAddress,
				SeqNumsRange:  cr.SequenceNumberRange,
				MerkleRoot:    cr.MerkleRoot,
			})
			commitReportIdx++
		}

		ri = append(ri, cciptypes.ExecuteReportInfo{
			AbstractReports: execReport.ChainReports,
			MerkleRoots:     merkleRoots,
		})
	}

	return ri, nil
}

func (p *Plugin) Reports(
	ctx context.Context, seqNr uint64, outcome ocr3types.Outcome,
) ([]ocr3types.ReportPlus[[]byte], error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseReports)

	if outcome == nil {
		lggr.Warn("no outcome, skipping report generation")
		return nil, nil
	}

	lggr.Debug("Decoding outcome in Reports function")
	decodedOutcome, err := p.ocrTypeCodec.DecodeOutcome(outcome)
	if err != nil {
		lggr.Errorw("unable to decode outcome", "err", err)
		return nil, fmt.Errorf("unable to decode outcome: %w", err)
	}

	if len(decodedOutcome.Reports) == 0 {
		lggr.Warnw("empty report", "outcome", decodedOutcome)
		return nil, nil
	}

	lggr.Debug("extracting report info")
	reportInfos, err := extractReportInfo(decodedOutcome)
	lggr.Debugw("report infos", "reportInfos", reportInfos)
	if err != nil {
		return nil, fmt.Errorf("extract report info: %w", err)
	}
	if len(reportInfos) != len(decodedOutcome.Reports) {
		lggr.Errorw("report infos length mismatch", "expected", len(decodedOutcome.Reports),
			"got", len(reportInfos))
		return nil, fmt.Errorf("expected %d report infos that equals number of reports"+
			", got %d", len(decodedOutcome.Reports), len(reportInfos))
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

	lggr.Debug("Encoding reports and adding them to final report")
	// Handle all reports in the outcome
	var reports []ocr3types.ReportPlus[[]byte]
	for i, execReport := range decodedOutcome.Reports {
		lggr.Debugw("Encoding Report: ", "report", execReport)
		encodedReport, err := p.reportCodec.Encode(ctx, execReport)
		if err != nil {
			return nil, fmt.Errorf("unable to encode report %w", err)
		}

		encodedInfo, err := reportInfos[i].Encode()
		if err != nil {
			lggr.Errorw("unable to encode report info", "err", err)
			return nil, fmt.Errorf("unable to encode report info %w", err)
		}

		reports = append(reports, ocr3types.ReportPlus[[]byte]{
			ReportWithInfo: ocr3types.ReportWithInfo[[]byte]{
				Report: encodedReport,
				Info:   encodedInfo,
			},
			TransmissionScheduleOverride: transmissionSchedule,
		})
	}

	return reports, nil
}

// validateReport validates various aspects of the report.
// Pure checks are placed earlier in the function on purpose to avoid
// unnecessary network or DB I/O.
// If you're added more checks make sure to follow this pattern.
func (p *Plugin) validateReport(
	ctx context.Context,
	lggr logger.Logger,
	r ocr3types.ReportWithInfo[[]byte],
) (decodedReport cciptypes.ExecutePluginReport, err error) {
	// Just a safety check, should never happen.
	if r.Report == nil {
		lggr.Warn("skipping nil report")
		return cciptypes.ExecutePluginReport{}, plugincommon.NewErrInvalidReport("nil report")
	}

	decodedReport, err = p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return cciptypes.ExecutePluginReport{},
			plugincommon.NewErrValidatingReport(fmt.Errorf("decode exec plugin report: %w", err))
	}

	if len(decodedReport.ChainReports) == 0 {
		lggr.Infow("skipping empty report")
		return cciptypes.ExecutePluginReport{},
			plugincommon.NewErrInvalidReport("empty report")
	}

	// check if we support the dest, if not we can't do the checks needed.
	supports, err := p.chainSupport.SupportsDestChain(p.reportingCfg.OracleID)
	if err != nil {
		return cciptypes.ExecutePluginReport{},
			plugincommon.NewErrValidatingReport(fmt.Errorf("supports dest chain: %w", err))
	}

	if !supports {
		lggr.Errorw("dest chain not supported by this oracle, can't run report acceptance procedures, " +
			"transmission schedule is wrong - check CCIPHome chainConfigs and ensure that the right oracles are " +
			"assigned as readers of the destination chain, or if " +
			"this oracle should support the destination chain but isn't!")
		return cciptypes.ExecutePluginReport{},
			plugincommon.NewErrInvalidReport("dest chain not supported")
	}
	if err := p.checkConfigDigest(); err != nil {
		if errors.Is(err, errOffRampConfigMismatch) {
			return cciptypes.ExecutePluginReport{},
				plugincommon.NewErrInvalidReport(err.Error())
		}
		return cciptypes.ExecutePluginReport{},
			plugincommon.NewErrValidatingReport(err)
	}

	// check that the messages in the report are not already executed onchain.
	// note that this involves a set of DB queries, hence why its last in the checks.
	err = p.checkAlreadyExecuted(ctx, lggr, decodedReport.ChainReports)
	if errors.Is(err, errAlreadyExecuted) {
		// Some messages in the report have already
		// been executed, so we don't want to re-execute them.
		// This gives the exec plugin a chance to remedy the situation
		// by selecting a different set of messages.
		return decodedReport, plugincommon.NewErrInvalidReport(err.Error())
	}
	if err != nil {
		// TODO: should we return true here if we couldn't check for already executed messages?
		err := fmt.Errorf("checking for already executed messages failed: %w", err)
		return decodedReport, plugincommon.NewErrValidatingReport(err)
	}

	return decodedReport, nil
}

// checkAlreadyExecuted checks if all the messages from all source chains in a single report have already been executed
// on the destination chain. It queries the DB for executed messages in the given sequence
// number range for each source chain in the report.
//
// The reasoning behind the "all or nothing" approach is that if some messages have already been executed
// and some have not, its possible that messages with skipped nonces have been executed. This could happen
// due to out of order transmissions. Giving the transmission another shot would mean we won't have
// extended delays due to skipped nonces.
//
// However, if all messages in the report are already executed, which is the usual case, we can safely skip
// transmitting the report.
func (p *Plugin) checkAlreadyExecuted(
	ctx context.Context,
	lggr logger.Logger,
	reports []cciptypes.ExecutePluginReportSingleChain,
) error {
	seqNrRangesBySource := getSeqNrRangesBySource(reports)

	executed, err := p.ccipReader.ExecutedMessages(
		ctx,
		seqNrRangesBySource,
		primitives.Unconfirmed)
	if err != nil {
		return fmt.Errorf("couldn't check if messages already executed: %w", err)
	}

	for sourceChainSelector, reportRanges := range seqNrRangesBySource {
		if _, exists := executed[sourceChainSelector]; !exists {
			lggr.Debugw("no messages from source chain executed yet",
				"sourceChain", sourceChainSelector)
			return nil
		}

		snSet := mapset.NewSet[cciptypes.SeqNum]()
		for _, snRange := range reportRanges {
			snSet.Append(snRange.ToSlice()...)
		}

		executedSet := mapset.NewSet(executed[sourceChainSelector]...)
		intersection := executedSet.Intersect(snSet)
		if intersection.Cardinality() != snSet.Cardinality() {
			// Some messages have not been executed, return early to accept/transmit the report.
			notYetExecuted := snSet.Difference(executedSet)
			lggr.Debugw("some messages from source not yet executed",
				"sourceChain", sourceChainSelector,
				"seqNrRange", seqNrRangesBySource[sourceChainSelector],
				"executed", executed,
				"notYetExecuted", notYetExecuted,
			)
			return nil
		}

		// All messages from this source have been executed, check the next source chain.
		lggr.Debugw("all messages from source already executed, checking next source",
			"sourceChain", sourceChainSelector,
			"seqNrRange", seqNrRangesBySource[sourceChainSelector],
			"executed", executed,
		)
	}

	return fmt.Errorf(
		"%w: all messages from all sources in provided reports already executed, not accepting/transmitting",
		errAlreadyExecuted)
}

// getSeqNrRangesBySource returns a map of source chain selector to
// the sequence number range of the messages in the report.
func getSeqNrRangesBySource(
	chainReports []cciptypes.ExecutePluginReportSingleChain,
) map[cciptypes.ChainSelector][]cciptypes.SeqNumRange {
	seqNrRangesBySource := make(map[cciptypes.ChainSelector][]cciptypes.SeqNumRange)
	for _, chainReport := range chainReports {
		// This should never happen, indicates a bug in the report building and accepting process.
		// But we sanity check since slices.Min/Max will panic on empty slices.
		if len(chainReport.Messages) == 0 {
			continue
		}
		if _, exists := seqNrRangesBySource[chainReport.SourceChainSelector]; !exists {
			seqNrRangesBySource[chainReport.SourceChainSelector] = []cciptypes.SeqNumRange{}
		}

		cmpr := func(a, b cciptypes.Message) int {
			return cmp.Compare(a.Header.SequenceNumber, b.Header.SequenceNumber)
		}
		minMsg := slices.MinFunc(chainReport.Messages, cmpr)
		maxMsg := slices.MaxFunc(chainReport.Messages, cmpr)
		seqNrRangesBySource[chainReport.SourceChainSelector] = []cciptypes.SeqNumRange{cciptypes.NewSeqNumRange(
			minMsg.Header.SequenceNumber,
			maxMsg.Header.SequenceNumber,
		)}
	}
	return seqNrRangesBySource
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseShouldAccept)

	supportsDest, err := p.supportsDestChain()
	if err != nil {
		lggr.Errorw("error checking if destination chain is supported", "err", err)
		return false, fmt.Errorf("checking if destination chain is supported: %w", err)
	}
	if !supportsDest {
		lggr.Errorw("dest chain not supported by this oracle, can't run report acceptance procedures, " +
			"transmission schedule is wrong - check CCIPHome chainConfigs and ensure that the right oracles are " +
			"assigned as readers of the destination chain, or if " +
			"this oracle should support the destination chain but isn't!")
		return false, plugincommon.NewErrInvalidReport("destination chain not supported")
	}

	decodedReport, err := p.validateReport(ctx, lggr, r)
	if errors.Is(err, plugincommon.ErrInvalidReport) {
		lggr.Infow("report not valid, not accepting", "err", err)
		return false, nil
	}
	if err != nil {
		lggr.Infow("validation error", "err", err)
		return false, fmt.Errorf("validating report: %w", err)
	}

	// TODO: consider doing this in validateReport,
	// will end up doing it in both ShouldAccept and ShouldTransmit.
	sourceChains := slicelib.Map(decodedReport.ChainReports,
		func(r cciptypes.ExecutePluginReportSingleChain) cciptypes.ChainSelector {
			return r.SourceChainSelector
		})
	isCursed, err := plugincommon.IsReportCursed(ctx, lggr, p.ccipReader, sourceChains)
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

	decodedReport, err := p.validateReport(ctx, lggr, r)
	if errors.Is(err, plugincommon.ErrInvalidReport) {
		lggr.Infow("report not valid, not transmitting", "err", err)
		return false, nil
	}
	if err != nil {
		lggr.Infow("validation error", "err", err)
		return false, fmt.Errorf("validating report: %w", err)
	}

	lggr.Infow("ShouldTransmitAttestedReport returns true, report accepted",
		"reports", decodedReport.ChainReports,
	)
	return true, nil
}

func (p *Plugin) Close() error {
	p.lggr.Infow("closing exec plugin")

	closeable := []io.Closer{
		p.tokenDataObserver,
		p.ccipReader,
	}

	return services.CloseAll(closeable...)
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
	return p.supportsChain(p.destChain)
}

func (p *Plugin) supportsChain(chainSel cciptypes.ChainSelector) (bool, error) {
	chains, err := p.supportedChains(p.reportingCfg.OracleID)
	if err != nil {
		return false, fmt.Errorf("error getting supported chains: %w", err)
	}
	return chains.Contains(chainSel), nil
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
