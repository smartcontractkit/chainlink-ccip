package execute

import (
	"context"
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/gas"
	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

// maxReportSizeBytes that should be returned as an execution report payload.
const maxReportSizeBytes = 250_000

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
	donID        uint32
	reportingCfg ocr3types.ReportingPluginConfig
	cfg          pluginconfig.ExecutePluginConfig

	// providers
	ccipReader   readerpkg.CCIPReader
	readerSyncer *plugincommon.BackgroundReaderSyncer
	reportCodec  cciptypes.ExecutePluginCodec
	msgHasher    cciptypes.MessageHasher
	homeChain    reader.HomeChain
	discovery    *discovery.ContractDiscoveryProcessor

	oracleIDToP2pID   map[commontypes.OracleID]libocrtypes.PeerID
	tokenDataObserver tokendata.TokenDataObserver
	estimateProvider  gas.EstimateProvider
	lggr              logger.Logger

	// state
	contractsInitialized bool
}

func NewPlugin(
	donID uint32,
	reportingCfg ocr3types.ReportingPluginConfig,
	cfg pluginconfig.ExecutePluginConfig,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	ccipReader readerpkg.CCIPReader,
	reportCodec cciptypes.ExecutePluginCodec,
	msgHasher cciptypes.MessageHasher,
	homeChain reader.HomeChain,
	tokenDataObserver tokendata.TokenDataObserver,
	estimateProvider gas.EstimateProvider,
	lggr logger.Logger,
) *Plugin {
	readerSyncer := plugincommon.NewBackgroundReaderSyncer(
		lggr,
		ccipReader,
		syncTimeout(cfg.SyncTimeout),
		syncFrequency(cfg.SyncFrequency),
	)
	if err := readerSyncer.Start(context.Background()); err != nil {
		lggr.Errorw("error starting background reader syncer", "err", err)
	}

	return &Plugin{
		donID:             donID,
		reportingCfg:      reportingCfg,
		cfg:               cfg,
		oracleIDToP2pID:   oracleIDToP2pID,
		ccipReader:        ccipReader,
		readerSyncer:      readerSyncer,
		reportCodec:       reportCodec,
		msgHasher:         msgHasher,
		homeChain:         homeChain,
		tokenDataObserver: tokenDataObserver,
		estimateProvider:  estimateProvider,
		lggr:              lggr,
		discovery: discovery.NewContractDiscoveryProcessor(
			lggr,
			&ccipReader,
			homeChain,
			cfg.DestChain,
			reportingCfg.F,
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

// Observation collects data across two phases which happen in separate rounds.
// These phases happen continuously so that except for the first round, every
// subsequent round can have a new execution report.
//
// Phase 1: Gather commit reports from the destination chain and determine
// which messages are required to build a valid execution report.
//
// Phase 2: Gather messages from the source chains and build the execution
// report.
// nolint:gocyclo // todo
func (p *Plugin) Observation(
	ctx context.Context, outctx ocr3types.OutcomeContext, _ types.Query,
) (types.Observation, error) {
	var err error
	var previousOutcome exectypes.Outcome

	if outctx.PreviousOutcome != nil {
		previousOutcome, err = exectypes.DecodeOutcome(outctx.PreviousOutcome)
		if err != nil {
			return types.Observation{}, fmt.Errorf("unable to decode previous outcome: %w", err)
		}
		p.lggr.Infow("decoded previous outcome", "previousOutcome", previousOutcome)
	}

	var discoveryObs dt.Observation
	// discovery processor disabled by setting it to nil.
	if p.discovery != nil {
		discoveryObs, err = p.discovery.Observation(ctx, dt.Outcome{}, dt.Query{})
		if err != nil {
			p.lggr.Errorw("failed to discover contracts", "err", err)
		}

		if !p.contractsInitialized {
			p.lggr.Infow("contracts not initialized, only making discovery observations")
			return exectypes.Observation{Contracts: discoveryObs}.Encode()
		}
	}

	state := previousOutcome.State.Next()
	p.lggr.Debugw("Execute plugin performing observation", "state", state)
	switch state {
	case exectypes.GetCommitReports:
		fetchFrom := time.Now().Add(-p.cfg.OffchainConfig.MessageVisibilityInterval.Duration()).UTC()

		// Phase 1: Gather commit reports from the destination chain and determine which messages are required to build
		//          a valid execution report.
		supportsDest, err := p.supportsDestChain()
		if err != nil {
			return types.Observation{}, fmt.Errorf("unable to determine if the destination chain is supported: %w", err)
		}
		if supportsDest {
			groupedCommits, err := getPendingExecutedReports(ctx, p.ccipReader, p.cfg.DestChain, fetchFrom, p.lggr)
			if err != nil {
				return types.Observation{}, err
			}

			// TODO: truncate grouped to a maximum observation size?
			return exectypes.NewObservation(groupedCommits, nil, nil, nil, discoveryObs).Encode()
		}

		// No observation for non-dest readers.
		return types.Observation{}, nil
	case exectypes.GetMessages:
		// Phase 2: Gather messages from the source chains and build the execution report.
		messages := make(exectypes.MessageObservations)
		if len(previousOutcome.PendingCommitReports) == 0 {
			p.lggr.Debug("TODO: No reports to execute. This is expected after a cold start.")
			// No reports to execute.
			// This is expected after a cold start.
		} else {
			commitReportCache := make(map[cciptypes.ChainSelector][]exectypes.CommitData)
			for _, report := range previousOutcome.PendingCommitReports {
				commitReportCache[report.SourceChain] = append(commitReportCache[report.SourceChain], report)
			}

			for srcChain, reports := range commitReportCache {
				if len(reports) == 0 {
					continue
				}

				ranges, err := computeRanges(reports)
				if err != nil {
					return types.Observation{}, err
				}

				// Read messages for each range.
				for _, seqRange := range ranges {
					// TODO: check if srcChain is supported.
					msgs, err := p.ccipReader.MsgsBetweenSeqNums(ctx, srcChain, seqRange)
					if err != nil {
						return nil, err
					}
					for _, msg := range msgs {
						if _, ok := messages[srcChain]; !ok {
							messages[srcChain] = make(map[cciptypes.SeqNum]cciptypes.Message)
						}
						messages[srcChain][msg.Header.SequenceNumber] = msg
					}
				}
			}
		}

		// Regroup the commit reports back into the observation format.
		// TODO: use same format for Observation and Outcome.
		groupedCommits := make(exectypes.CommitObservations)
		for _, report := range previousOutcome.PendingCommitReports {
			if _, ok := groupedCommits[report.SourceChain]; !ok {
				groupedCommits[report.SourceChain] = []exectypes.CommitData{}
			}
			groupedCommits[report.SourceChain] = append(groupedCommits[report.SourceChain], report)
		}

		tkData, err1 := p.tokenDataObserver.Observe(ctx, messages)
		if err1 != nil {
			return types.Observation{}, fmt.Errorf("unable to process token data %w", err1)
		}

		return exectypes.NewObservation(groupedCommits, messages, tkData, nil, discoveryObs).Encode()

	case exectypes.Filter:
		// Phase 3: observe nonce for each unique source/sender pair.
		nonceRequestArgs := make(map[cciptypes.ChainSelector]map[string]struct{})

		// Collect unique senders.
		for _, commitReport := range previousOutcome.PendingCommitReports {
			if _, ok := nonceRequestArgs[commitReport.SourceChain]; !ok {
				nonceRequestArgs[commitReport.SourceChain] = make(map[string]struct{})
			}

			for _, msg := range commitReport.Messages {
				sender := typeconv.AddressBytesToString(msg.Sender[:], uint64(p.cfg.DestChain))
				nonceRequestArgs[commitReport.SourceChain][sender] = struct{}{}
			}
		}

		// Read args from chain.
		nonceObservations := make(exectypes.NonceObservations)
		for srcChain, addrSet := range nonceRequestArgs {
			// TODO: check if srcSelector is supported.
			addrs := maps.Keys(addrSet)
			nonces, err := p.ccipReader.Nonces(ctx, srcChain, p.cfg.DestChain, addrs)
			if err != nil {
				return types.Observation{}, fmt.Errorf("unable to get nonces: %w", err)
			}
			nonceObservations[srcChain] = nonces
		}

		return exectypes.NewObservation(nil, nil, nil, nonceObservations, discoveryObs).Encode()
	default:
		return types.Observation{}, fmt.Errorf("unknown state")
	}
}

func (p *Plugin) ValidateObservation(
	outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation,
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
		commitReports[i], err = builder.Add(report)
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

// Outcome collects the reports from the two phases and constructs the final outcome. Part of the outcome is a fully
// formed report that will be encoded for final transmission in the reporting phase.
// nolint:gocyclo // todo
func (p *Plugin) Outcome(
	outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	var previousOutcome exectypes.Outcome
	if outctx.PreviousOutcome != nil {
		var err error
		previousOutcome, err = exectypes.DecodeOutcome(outctx.PreviousOutcome)
		if err != nil {
			return nil, fmt.Errorf("unable to decode previous outcome: %w", err)
		}
	}

	decodedAos, err := decodeAttributedObservations(aos)
	if err != nil {
		return nil, fmt.Errorf("unable to decode observations: %w", err)
	}

	// discovery processor disabled by setting it to nil.
	if p.discovery != nil {
		discoveryAos := make([]plugincommon.AttributedObservation[dt.Observation], len(decodedAos))
		for i := range decodedAos {
			discoveryAos[i] = plugincommon.AttributedObservation[dt.Observation]{
				OracleID:    decodedAos[i].OracleID,
				Observation: decodedAos[i].Observation.Contracts,
			}
		}
		_, err = p.discovery.Outcome(dt.Outcome{}, dt.Query{}, discoveryAos)
		if err != nil {
			return nil, fmt.Errorf("unable to process outcome of discovery processor: %w", err)
		}
		p.contractsInitialized = true
	}

	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to get FChain: %w", err)
	}

	observation, err := getConsensusObservation(
		p.lggr, decodedAos, p.reportingCfg.OracleID, p.cfg.DestChain, p.reportingCfg.F, fChain)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to get consensus observation: %w", err)
	}

	var outcome exectypes.Outcome
	state := previousOutcome.State.Next()
	switch state {
	case exectypes.GetCommitReports:
		// flatten commit reports and sort by timestamp.
		var commitReports []exectypes.CommitData
		for _, report := range observation.CommitReports {
			commitReports = append(commitReports, report...)
		}
		sort.Slice(commitReports, func(i, j int) bool {
			return commitReports[i].Timestamp.Before(commitReports[j].Timestamp)
		})

		outcome = exectypes.NewOutcome(state, commitReports, cciptypes.ExecutePluginReport{})
	case exectypes.GetMessages:
		commitReports := previousOutcome.PendingCommitReports

		// add messages to their commitReports.
		for i, report := range commitReports {
			report.Messages = nil
			for j := report.SequenceNumberRange.Start(); j <= report.SequenceNumberRange.End(); j++ {
				if msg, ok := observation.Messages[report.SourceChain][j]; ok {
					report.Messages = append(report.Messages, msg)
				}

				if tokenData, ok := observation.TokenData[report.SourceChain][j]; ok {
					report.MessageTokenData = append(report.MessageTokenData, tokenData)
				}
			}
			commitReports[i].Messages = report.Messages
			commitReports[i].MessageTokenData = report.MessageTokenData
		}

		outcome = exectypes.NewOutcome(state, commitReports, cciptypes.ExecutePluginReport{})
	case exectypes.Filter:
		commitReports := previousOutcome.PendingCommitReports

		// TODO: this function should be pure, a context should not be needed.
		builder := report.NewBuilder(
			context.Background(),
			p.lggr,
			p.msgHasher,
			p.reportCodec,
			p.estimateProvider,
			observation.Nonces,
			p.cfg.DestChain,
			uint64(maxReportSizeBytes),
			p.cfg.OffchainConfig.BatchGasLimit,
		)
		outcomeReports, commitReports, err := selectReport(
			p.lggr,
			commitReports,
			builder)
		if err != nil {
			return ocr3types.Outcome{}, fmt.Errorf("unable to extract proofs: %w", err)
		}

		execReport := cciptypes.ExecutePluginReport{
			ChainReports: outcomeReports,
		}

		outcome = exectypes.NewOutcome(state, commitReports, execReport)
	default:
		panic("unknown state")
	}

	if outcome.IsEmpty() {
		p.lggr.Warnw(
			fmt.Sprintf("[oracle %d] exec outcome: empty outcome", p.reportingCfg.OracleID),
			"oracle", p.reportingCfg.OracleID,
			"execPluginState", state)
		if p.contractsInitialized {
			return exectypes.Outcome{State: exectypes.Initialized}.Encode()
		}
		return nil, nil
	}

	p.lggr.Infow(
		fmt.Sprintf("[oracle %d] exec outcome: generated outcome", p.reportingCfg.OracleID),
		"oracle", p.reportingCfg.OracleID,
		"execPluginState", state,
		"outcome", outcome)

	return outcome.Encode()
}

func (p *Plugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	if outcome == nil {
		p.lggr.Warn("no outcome, skipping report generation")
		return nil, nil
	}

	decodedOutcome, err := exectypes.DecodeOutcome(outcome)
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
	// Just a safety check, should never happen.
	if r.Report == nil {
		p.lggr.Warn("skipping nil report")
		return false, nil
	}

	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	p.lggr.Infow("Checking if ShouldAcceptAttestedReport", "chainReports", decodedReport.ChainReports)
	if len(decodedReport.ChainReports) == 0 {
		p.lggr.Info("skipping empty report")
		return false, nil
	}

	p.lggr.Info("ShouldAcceptAttestedReport returns true, report accepted")
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
		p.lggr.Debug("not a destination writer, skipping report transmission")
		return false, nil
	}

	// we only transmit reports if we are the "blue" instance.
	// we can check this by reading the OCR conigs home chain.
	isGreen, err := p.isGreenInstance(ctx)
	if err != nil {
		return false, fmt.Errorf("ShouldTransmitAcceptedReport.isGreenInstance: %w", err)
	}

	if isGreen {
		p.lggr.Debugw("not the blue instance, skipping report transmission",
			"myDigest", p.reportingCfg.ConfigDigest.Hex())
		return false, nil
	}

	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	// TODO: Final validation?

	p.lggr.Infow("transmitting report",
		"reports", decodedReport.ChainReports,
	)
	return true, nil
}

func (p *Plugin) isGreenInstance(ctx context.Context) (bool, error) {
	ocrConfigs, err := p.homeChain.GetOCRConfigs(ctx, p.donID, consts.PluginTypeExecute)
	if err != nil {
		return false, fmt.Errorf("failed to get ocr configs from home chain: %w", err)
	}

	return len(ocrConfigs) == 2 && ocrConfigs[1].ConfigDigest == p.reportingCfg.ConfigDigest, nil
}

func (p *Plugin) Close() error {
	timeout := 10 * time.Second // todo: cfg
	ctx, cf := context.WithTimeout(context.Background(), timeout)
	defer cf()

	if err := p.readerSyncer.Close(); err != nil {
		p.lggr.Warnw("error closing reader syncer", "err", err)
	}

	if err := p.ccipReader.Close(ctx); err != nil {
		return fmt.Errorf("close ccip reader: %w", err)
	}

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

func syncFrequency(configuredValue time.Duration) time.Duration {
	if configuredValue.Milliseconds() == 0 {
		return 10 * time.Second
	}
	return configuredValue
}

func syncTimeout(configuredValue time.Duration) time.Duration {
	if configuredValue.Milliseconds() == 0 {
		return 3 * time.Second
	}
	return configuredValue
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
