package execute

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
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

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
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
	tokenDataReader TokenDataReader
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

// TokenDataReader is an interface for reading extra token data from an async process.
// TODO: Build a token data reading process.
type TokenDataReader interface {
	ReadTokenData(ctx context.Context, srcChain cciptypes.ChainSelector, num cciptypes.SeqNum) ([][]byte, error)
}

// buildSingleChainReportMaxSize generates the largest report which fits into maxSizeBytes.
// See buildSingleChainReport for more details about how a report is built.
func buildSingleChainReportMaxSize(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	tokenDataReader TokenDataReader,
	encoder cciptypes.ExecutePluginCodec,
	report plugintypes.ExecutePluginCommitDataWithMessages,
	maxSizeBytes int,
) (cciptypes.ExecutePluginReportSingleChain, int, plugintypes.ExecutePluginCommitDataWithMessages, error) {
	finalReport, encodedSize, err :=
		buildSingleChainReport(ctx, lggr, hasher, tokenDataReader, encoder, report, 0)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			0,
			plugintypes.ExecutePluginCommitDataWithMessages{},
			fmt.Errorf("unable to build a single chain report (max): %w", err)
	}

	// return fully executed report
	if encodedSize <= maxSizeBytes {
		report = markNewMessagesExecuted(finalReport, report)
		return finalReport, encodedSize, report, nil
	}

	var searchErr error
	idx := sort.Search(len(report.Messages), func(mid int) bool {
		if searchErr != nil {
			return false
		}
		finalReport2, encodedSize2, err :=
			buildSingleChainReport(ctx, lggr, hasher, tokenDataReader, encoder, report, mid)
		if searchErr != nil {
			searchErr = fmt.Errorf("unable to build a single chain report (messages %d): %w", mid, err)
		}

		if (encodedSize2) <= maxSizeBytes {
			// mid is a valid report size, try something bigger next iteration.
			finalReport = finalReport2
			encodedSize = encodedSize2
			return false // not full
		}
		return true // full
	})
	if searchErr != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, 0, plugintypes.ExecutePluginCommitDataWithMessages{}, searchErr
	}

	// No messages fit into the report.
	if idx <= 0 {
		return cciptypes.ExecutePluginReportSingleChain{},
			0,
			plugintypes.ExecutePluginCommitDataWithMessages{},
			errEmptyReport
	}

	report = markNewMessagesExecuted(finalReport, report)
	return finalReport, encodedSize, report, nil
}

// buildSingleChainReport converts the on-chain event data stored in cciptypes.ExecutePluginCommitDataWithMessages into
// the final on-chain report format.
//
// The hasher and encoding codec are provided as arguments to allow for chain-specific formats to be used.
//
// The maxMessages argument is used to limit the number of messages that are included in the report. If maxMessages is
// set to 0, all messages will be included. This allows the caller to create smaller reports if needed.
func buildSingleChainReport(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	tokenDataReader TokenDataReader,
	encoder cciptypes.ExecutePluginCodec,
	report plugintypes.ExecutePluginCommitDataWithMessages,
	maxMessages int,
) (cciptypes.ExecutePluginReportSingleChain, int, error) {
	// TODO: maxMessages selects messages in FIFO order which may not yield the optimal message size. One message with a
	//       maximum data size could push the report over a size limit even if several smaller messages could have fit.
	if maxMessages == 0 {
		maxMessages = len(report.Messages)
	}

	lggr.Debugw(
		"constructing merkle tree",
		"sourceChain", report.SourceChain,
		"expectedRoot", report.MerkleRoot.String(),
		"treeLeaves", len(report.Messages))

	tree, err := constructMerkleTree(ctx, hasher, report)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, 0,
			fmt.Errorf("unable to construct merkle tree from messages for report (%s): %w", report.MerkleRoot.String(), err)
	}

	// Verify merkle root.
	hash := tree.Root()
	if !bytes.Equal(hash[:], report.MerkleRoot[:]) {
		actualStr := "0x" + hex.EncodeToString(hash[:])
		return cciptypes.ExecutePluginReportSingleChain{}, 0,
			fmt.Errorf("merkle root mismatch: expected %s, got %s", report.MerkleRoot.String(), actualStr)
	}

	// Iterate sequence range and executed messages to select messages to execute.
	numMsgs := len(report.Messages)
	var toExecute []int
	var offchainTokenData [][][]byte
	var msgInRoot []cciptypes.Message
	executedIdx := 0
	for i := 0; i < numMsgs && len(toExecute) <= maxMessages; i++ {
		seqNum := report.SequenceNumberRange.Start() + cciptypes.SeqNum(i)
		// Skip messages which are already executed
		if executedIdx < len(report.ExecutedMessages) && report.ExecutedMessages[executedIdx] == seqNum {
			executedIdx++
		} else {
			msg := report.Messages[i]
			tokenData, err := tokenDataReader.ReadTokenData(context.Background(), report.SourceChain, msg.Header.SequenceNumber)
			if err != nil {
				// TODO: skip message instead of failing the whole thing.
				//       that might mean moving the token data reading out of the loop.
				lggr.Infow(
					"unable to read token data",
					"sourceChain", report.SourceChain,
					"seqNum", msg.Header.SequenceNumber,
					"error", err)
				return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf(
					"unable to read token data for message %d: %w", msg.Header.SequenceNumber, err)
			}

			lggr.Infow(
				"read token data",
				"sourceChain", report.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"data", tokenData)
			offchainTokenData = append(offchainTokenData, tokenData)
			toExecute = append(toExecute, i)
			msgInRoot = append(msgInRoot, msg)
		}
	}

	lggr.Infow(
		"selected messages from commit report for execution",
		"sourceChain", report.SourceChain,
		"commitRoot", report.MerkleRoot.String(),
		"numMessages", numMsgs,
		"toExecute", len(toExecute))
	proof, err := tree.Prove(toExecute)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, 0,
			fmt.Errorf("unable to prove messages for report %s: %w", report.MerkleRoot.String(), err)
	}

	var proofsCast []cciptypes.Bytes32
	for _, p := range proof.Hashes {
		proofsCast = append(proofsCast, p)
	}

	finalReport := cciptypes.ExecutePluginReportSingleChain{
		SourceChainSelector: report.SourceChain,
		Messages:            msgInRoot,
		OffchainTokenData:   offchainTokenData,
		Proofs:              proofsCast,
		ProofFlagBits:       cciptypes.BigInt{Int: slicelib.BoolsToBitFlags(proof.SourceFlags)},
	}

	// Note: ExecutePluginReport is a strict array of data, so wrapping the final report
	//       does not add any additional overhead to the size being computed here.

	// Compute the size of the encoded report.
	encoded, err := encoder.Encode(
		ctx,
		cciptypes.ExecutePluginReport{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{finalReport},
		},
	)
	if err != nil {
		lggr.Errorw("unable to encode report", "err", err, "report", finalReport)
		return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf("unable to encode report: %w", err)
	}

	return finalReport, len(encoded), nil
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
	tokenDataReader TokenDataReader,
	reports []plugintypes.ExecutePluginCommitDataWithMessages,
	maxReportSizeBytes int,
) ([]cciptypes.ExecutePluginReportSingleChain, []plugintypes.ExecutePluginCommitDataWithMessages, error) {
	// TODO: It may be desirable for this entire function to be an interface so that
	//       different selection algorithms can be used.

	// count number of fully executed reports so that they can be removed after iterating the reports.
	fullyExecuted := 0
	accumulatedSize := 0
	var finalReports []cciptypes.ExecutePluginReportSingleChain
	for reportIdx, report := range reports {
		// Reports at the end may not have messages yet.
		if len(report.Messages) == 0 {
			break
		}

		execReport, encodedSize, updatedReport, err :=
			buildSingleChainReportMaxSize(ctx, lggr, hasher, tokenDataReader, encoder,
				report, maxReportSizeBytes-accumulatedSize)
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
	}

	// Remove reports that are about to be executed.
	if fullyExecuted == len(reports) {
		reports = nil
	} else {
		reports = reports[fullyExecuted:]
	}

	lggr.Infow(
		"selected commit reports for execution report",
		"numReports", len(finalReports),
		"size", accumulatedSize,
		"incompleteReports", len(reports),
		"maxSize", maxReportSizeBytes)

	return finalReports, reports, nil
}

// Outcome collects the reports from the two phases and constructs the final outcome. Part of the outcome is a fully
// formed report that will be encoded for final transmission in the reporting phase.
func (p *Plugin) Outcome(
	outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	decodedObservations, err := decodeAttributedObservations(aos)
	if err != nil {
		return ocr3types.Outcome{}, err

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
		return ocr3types.Outcome{}, err
	}

	mergedMessageObservations, err := mergeMessageObservations(decodedObservations, fChain)
	if err != nil {
		return ocr3types.Outcome{}, err
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
	for _, report := range commitReports {
		report.Messages = nil
		for i := report.SequenceNumberRange.Start(); i <= report.SequenceNumberRange.End(); i++ {
			if msg, ok := observation.Messages[report.SourceChain][i]; ok {
				report.Messages = append(report.Messages, msg)
			}
		}
	}

	// TODO: this function should be pure, a context should not be needed.
	outcomeReports, commitReports, err :=
		selectReport(context.Background(), p.lggr, p.msgHasher, p.reportCodec, p.tokenDataReader,
			commitReports, maxReportSizeBytes)
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
		return false, fmt.Errorf("can't know if it's a writer: %w", err)
	}
	if !isWriter {
		p.lggr.Debugw("not a writer, skipping report transmission")
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
	panic("implement me")
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
