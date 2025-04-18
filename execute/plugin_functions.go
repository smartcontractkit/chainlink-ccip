package execute

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// validateCommitReportsReadingEligibility validates that all commit reports' source chains are supported by observer
func validateCommitReportsReadingEligibility(
	supportedChains mapset.Set[cciptypes.ChainSelector],
	observedData exectypes.CommitObservations,
) error {
	for chainSel, observedDataOfChain := range observedData {
		if !supportedChains.Contains(chainSel) {
			return fmt.Errorf("observer not allowed to read from chain %d", chainSel)
		}
		for _, data := range observedDataOfChain {
			if data.SourceChain != chainSel {
				return fmt.Errorf("invalid observed data, key=%d but data chain=%d",
					chainSel, data.SourceChain)
			}
		}
	}

	return nil
}

// validateMsgsReadingEligibility checks all observed messages are from supported chains
func validateMsgsReadingEligibility(
	supportedChains mapset.Set[cciptypes.ChainSelector],
	observedMsgs exectypes.MessageObservations,
) error {
	for chainSel, msgs := range observedMsgs {
		if len(msgs) == 0 {
			continue
		}

		if !supportedChains.Contains(chainSel) {
			return fmt.Errorf("observer not allowed to read from chain %d", chainSel)
		}
	}

	return nil
}

// validateTokenDataObservations validates that all token data observations belong to already observed messages
func validateTokenDataObservations(
	observedMsgs exectypes.MessageObservations,
	tokenData exectypes.TokenDataObservations,
) error {

	if len(observedMsgs) != len(tokenData) {
		return fmt.Errorf("unexpected number of token data observations: expected %d, got %d",
			len(observedMsgs), len(tokenData))
	}

	for chain, msgs := range observedMsgs {
		chainTd, ok := tokenData[chain]
		if !ok {
			return fmt.Errorf("token data not found for chain %d", chain)
		}
		if len(msgs) != len(chainTd) {
			return fmt.Errorf("unexpected number of token data observations for chain %d: expected %d, got %d",
				chain, len(msgs), len(chainTd))
		}
		for seq, msg := range msgs {
			if _, ok := tokenData[chain][seq]; !ok {
				return fmt.Errorf("token data not found for message %s", msg)
			}
		}
	}

	return nil
}

// validateHashesExist checks if the hashes exist for all the messages in the observation.
func validateHashesExist(
	observedMsgs exectypes.MessageObservations,
	hashes exectypes.MessageHashes,
) error {
	if len(observedMsgs) != len(hashes) {
		return fmt.Errorf("malformed observation, unexpected number of message hashes: expected %d, got %d",
			len(observedMsgs), len(hashes))
	}

	for chain, msgs := range observedMsgs {
		hashesForChain, ok := hashes[chain]
		if !ok {
			return fmt.Errorf("hash not found for chain %d", chain)
		}

		if len(msgs) != len(hashesForChain) {
			return fmt.Errorf("unexpected number of message hashes for chain %d: expected %d, got %d",
				chain, len(msgs), len(hashesForChain))
		}

		for seq, msg := range msgs {
			h, exists := hashes[chain][seq]
			if !exists {
				return fmt.Errorf("hash not found for message %s", msg)
			}
			if h.IsEmpty() {
				return fmt.Errorf("hash is empty for message %s", msg)
			}
		}
	}

	return nil
}

// validateMessagesConformToCommitReports cross-checks messages and reports
// 1. checks if the messages observed are exactly the same as the messages in the commit reports. No more and no less.
// 2. checks all reports have their messages observed.
func validateMessagesConformToCommitReports(
	observedData exectypes.CommitObservations,
	observedMsgs exectypes.MessageObservations,
) error {
	if len(observedData) != len(observedMsgs) {
		return fmt.Errorf("count of observed data=%d and observed msgs=%d do not match",
			len(observedData), len(observedMsgs))
	}

	msgsCount := 0
	for chain, report := range observedData {
		msgsMap, ok := observedMsgs[chain]
		if !ok {
			return fmt.Errorf("no messages observed for chain %d, while report was observed", chain)
		}

		for _, data := range report {
			for seqNum := data.SequenceNumberRange.Start(); seqNum <= data.SequenceNumberRange.End(); seqNum++ {
				_, ok = msgsMap[seqNum]
				if !ok {
					return fmt.Errorf("no message observed for sequence number %d, "+
						"while report's range sholud include it for chain %d", seqNum, chain)
				}
				msgsCount++
			}
		}
	}
	allMsgs := observedMsgs.Flatten()
	// need to make sure that only messages that are in the commit reports are observed
	if msgsCount != len(allMsgs) {
		return fmt.Errorf("messages observed %d do not match the messages in the commit reports %d",
			len(allMsgs), msgsCount)
	}

	return nil
}

// validateObservedSequenceNumbers checks if the sequence numbers of the provided messages are unique for each chain
// and that they match the observed max sequence numbers.
func validateObservedSequenceNumbers(
	supportedChains mapset.Set[cciptypes.ChainSelector],
	observedData map[cciptypes.ChainSelector][]exectypes.CommitData,
) error {
	for chainSel, commitData := range observedData {
		if !supportedChains.Contains(chainSel) {
			return fmt.Errorf("observed a non-supported chain %d", chainSel)
		}

		// observed commitData must not contain duplicates

		observedMerkleRoots := mapset.NewSet[string]()
		observedRanges := make([]cciptypes.SeqNumRange, 0)

		for _, data := range commitData {
			rootStr := data.MerkleRoot.String()
			if observedMerkleRoots.Contains(rootStr) {
				return fmt.Errorf("duplicate merkle root %s observed", rootStr)
			}
			observedMerkleRoots.Add(rootStr)

			for _, rng := range observedRanges {
				if rng.Overlaps(data.SequenceNumberRange) {
					return fmt.Errorf("sequence number range %v overlaps with %v", data.SequenceNumberRange, rng)
				}
			}
			observedRanges = append(observedRanges, data.SequenceNumberRange)

			// Executed sequence numbers should belong in the observed range.
			for _, seqNum := range data.ExecutedMessages {
				if !data.SequenceNumberRange.Contains(seqNum) {
					return fmt.Errorf("executed message %d not in observed range %v", seqNum, data.SequenceNumberRange)
				}
			}
		}
	}

	return nil
}

// msgsConformToSeqRange returns true if all sequence numbers in the range are observed in the messages.
// messages should map exactly one message to each sequence number in the range with no missing sequence numbers.
func msgsConformToSeqRange(msgs []cciptypes.Message, numberRange cciptypes.SeqNumRange) bool {
	msgMap := make(map[cciptypes.SeqNum]struct{})
	for _, msg := range msgs {
		msgMap[msg.Header.SequenceNumber] = struct{}{}
	}
	if len(msgMap) != numberRange.Length() {
		return false
	}
	// Check for missing sequence numbers in observed messages.
	for i := numberRange.Start(); i <= numberRange.End(); i++ {
		if _, ok := msgMap[i]; !ok {
			return false
		}
	}
	return true
}

var errOverlappingRanges = errors.New("overlapping sequence numbers in reports")

// computeRanges takes a slice of reports and computes the smallest number of contiguous ranges
// that cover all the sequence numbers in the reports.
// Example: computeRanges([10, 12], [13, 15], [20, 22]) = [10,15], [20,22]
// Note: reports need all messages to create a proof even if some are already executed.
// Note: the provided reports must be sorted by sequence number range starting sequence number.
func computeRanges(reports []exectypes.CommitData) ([]cciptypes.SeqNumRange, error) {
	var ranges []cciptypes.SeqNumRange

	if len(reports) == 0 {
		return nil, nil
	}

	var seqRange cciptypes.SeqNumRange
	for i, report := range reports {
		if i == 0 {
			// initialize
			seqRange = cciptypes.NewSeqNumRange(report.SequenceNumberRange.Start(), report.SequenceNumberRange.End())
		} else if seqRange.End()+1 == report.SequenceNumberRange.Start() {
			// extend the contiguous range
			seqRange.SetEnd(report.SequenceNumberRange.End())
		} else if report.SequenceNumberRange.Overlaps(seqRange) {
			return nil, errOverlappingRanges
		} else {
			ranges = append(ranges, seqRange)

			// Reset the range.
			seqRange = cciptypes.NewSeqNumRange(report.SequenceNumberRange.Start(), report.SequenceNumberRange.End())
		}
	}
	// add final range
	ranges = append(ranges, seqRange)

	return ranges, nil
}

// groupByChainSelector groups the reports by their chain selector.
// This skips cursed chains and reports with empty roots.
func groupByChainSelectorWithFilter(
	lggr logger.Logger,
	reports []cciptypes.CommitPluginReportWithMeta,
	cursedSourceChains map[cciptypes.ChainSelector]bool,
) exectypes.CommitObservations {
	commitReportCache := make(map[cciptypes.ChainSelector][]exectypes.CommitData)
	var filteredRoots int
	filteredByChain := make(map[cciptypes.ChainSelector]int)

	for _, report := range reports {
		merkleRoots := append(report.Report.BlessedMerkleRoots, report.Report.UnblessedMerkleRoots...)

		for _, singleReport := range merkleRoots {
			// Skip cursed chains
			if cursedSourceChains != nil && cursedSourceChains[singleReport.ChainSel] {
				filteredRoots++
				filteredByChain[singleReport.ChainSel]++
				lggr.Debugw("filtered merkle root from cursed chain",
					"chainSelector", singleReport.ChainSel,
					"merkleRoot", singleReport.MerkleRoot.String())
				continue
			}

			commitReportCache[singleReport.ChainSel] = append(commitReportCache[singleReport.ChainSel],
				exectypes.CommitData{
					SourceChain:         singleReport.ChainSel,
					OnRampAddress:       singleReport.OnRampAddress,
					Timestamp:           report.Timestamp,
					BlockNum:            report.BlockNum,
					MerkleRoot:          singleReport.MerkleRoot,
					SequenceNumberRange: singleReport.SeqNumsRange,
				})
		}
	}

	if filteredRoots > 0 {
		// Extract chains for more readable logging
		var filteredChains []string
		for chainSel, count := range filteredByChain {
			filteredChains = append(filteredChains,
				fmt.Sprintf("%s(%d)", chainSel.String(), count))
		}

		lggr.Infow("filtered cursed chain merkle roots during grouping",
			"filteredRoots", filteredRoots,
			"filteredChains", filteredChains)
	}

	return commitReportCache
}

// combineReportsAndMessages returns a new reports slice with fully executed messages removed.
// Reports that have all of their messages executed are not included in the result.
// The provided reports must be sorted by sequence number range starting sequence number.
func combineReportsAndMessages(
	reports []exectypes.CommitData, executedMessages []cciptypes.SeqNum,
) (pending []exectypes.CommitData, fullyExecuted []exectypes.CommitData) {
	if len(executedMessages) == 0 {
		return reports, nil
	}

	// filtered contains the reports with fully executed messages removed
	// and the executed messages appended to the report sorted by sequence number.
	for i, report := range reports {
		reportRange := report.SequenceNumberRange

		executedMsgsInReportRange := reportRange.FilterSlice(executedMessages)
		if len(executedMsgsInReportRange) == reportRange.Length() { // skip fully executed report.
			fullyExecuted = append(fullyExecuted, report)
			continue
		}

		sort.Slice(executedMsgsInReportRange, func(i, j int) bool {
			return executedMsgsInReportRange[i] < executedMsgsInReportRange[j]
		})
		report.ExecutedMessages = append(reports[i].ExecutedMessages, executedMsgsInReportRange...)
		pending = append(pending, report)
	}

	return pending, fullyExecuted
}

func decodeAttributedObservations(
	aos []types.AttributedObservation,
	ocrTypeCodec ocrtypecodec.ExecCodec,
) ([]plugincommon.AttributedObservation[exectypes.Observation], error) {
	decoded := make([]plugincommon.AttributedObservation[exectypes.Observation], len(aos))
	for i, ao := range aos {
		observation, err := ocrTypeCodec.DecodeObservation(ao.Observation)
		if err != nil {
			return nil, err
		}
		decoded[i] = plugincommon.AttributedObservation[exectypes.Observation]{
			Observation: observation,
			OracleID:    ao.Observer,
		}
	}
	return decoded, nil
}

func computeMessageObservationsConsensus(
	lggr logger.Logger,
	aos []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) exectypes.MessageObservations {
	validators := prepareValidatorsForComputeMessageObservationsConsensus(lggr, aos, fChain)

	messagesReachingConsensus := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]cciptypes.Message)
	for selector, validator := range validators {
		if msgs := validator.GetValid(); len(msgs) > 0 {
			if _, ok := messagesReachingConsensus[selector]; !ok {
				messagesReachingConsensus[selector] = make(map[cciptypes.SeqNum][]cciptypes.Message)
			}
			for _, msg := range msgs {
				if _, ok := messagesReachingConsensus[selector][msg.Header.SequenceNumber]; !ok {
					messagesReachingConsensus[selector][msg.Header.SequenceNumber] = make([]cciptypes.Message, 0)
				}
				messagesReachingConsensus[selector][msg.Header.SequenceNumber] = append(
					messagesReachingConsensus[selector][msg.Header.SequenceNumber],
					msg,
				)
			}
		}
	}

	results := make(exectypes.MessageObservations)
	for chain, seqNumToMsgs := range messagesReachingConsensus {
		for seqNum, msgsWithConsensus := range seqNumToMsgs {
			switch len(msgsWithConsensus) {
			case 0:
				lggr.Debugw("no message reached consensus for sequence number, skipping it",
					"chain", chain, "seqNum", seqNum)
			case 1:
				if _, ok := results[chain]; !ok {
					results[chain] = make(map[cciptypes.SeqNum]cciptypes.Message)
				}
				results[chain][seqNum] = msgsWithConsensus[0]
			default:
				lggr.Warnw("more than one message reached consensus for a sequence number, skipping it",
					"chain", chain, "seqNum", seqNum, "msgs", msgsWithConsensus)
			}
		}
	}

	if len(results) == 0 {
		return nil
	}
	return results
}

func prepareValidatorsForComputeMessageObservationsConsensus(
	lggr logger.Logger,
	observations []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]consensus.OracleMinObservation[cciptypes.Message] {
	// Create a validator for each chain
	validators := make(map[cciptypes.ChainSelector]consensus.OracleMinObservation[cciptypes.Message])
	for selector, f := range fChain {
		validators[selector] = consensus.NewOracleMinObservation[cciptypes.Message](consensus.FPlus1(f), nil)
	}

	// Add messages to the validator for each chain selector.
	for _, ao := range observations {
		for selector, messages := range ao.Observation.Messages {
			validator, ok := validators[selector]
			if !ok {
				lggr.Warnw("no F defined for chain", "chain", selector)
				continue
			}
			// Add reports
			for _, msg := range messages {
				validator.Add(msg, ao.OracleID)
			}
		}
	}

	return validators
}

// computeCommitObservationsConsensus come to consensus over observations of CommitData.
// Selects a merkle root and the relevant data if it has at least f+1 observations.
// If a merkle root is agreed twice but with different relevant data (e.g. OnRampAddress) then we don't have consensus.
// Executed messages within the root are agreed when at least f+1 nodes observed the execution.
func computeCommitObservationsConsensus(
	lggr logger.Logger,
	observations []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) exectypes.CommitObservations {
	merkleRootsVotes, executedMsgVotes := aggregateMerkleRootObservations(observations)

	validRoots := make([]merkleRootData, 0, len(merkleRootsVotes))
	for mr, votes := range merkleRootsVotes {
		if consensus.GteFPlusOne(fChain[mr.SourceChain], votes) {
			validRoots = append(validRoots, mr)
		} else {
			lggr.Debugw("merkle root with less than f+1 votes was found, skipping it", "mr", mr, "votes", votes)
		}
	}

	seenCount := make(map[cciptypes.Bytes32]int)
	for _, mr := range validRoots {
		seenCount[mr.MerkleRoot]++
	}
	validRoots = slicelib.Filter(validRoots, func(mr merkleRootData) bool { return seenCount[mr.MerkleRoot] == 1 })

	result := make(exectypes.CommitObservations)
	for _, mr := range validRoots {
		f, ok := fChain[mr.SourceChain]
		if !ok {
			lggr.Warnw("no fChain defined for chain", "chain", mr.SourceChain, "fChain", fChain)
			continue
		}

		executedMessages := make([]cciptypes.SeqNum, 0)
		for seqNum, count := range executedMsgVotes[mr] {
			if consensus.LtFPlusOne(f, count) {
				lggr.Debugw("skipping executed msg, less than f+1 votes", "mr", mr, "votes", count, "seqNum", seqNum)
				continue
			}
			executedMessages = append(executedMessages, seqNum)
		}
		sort.Slice(executedMessages, func(i, j int) bool { return executedMessages[i] < executedMessages[j] })

		if _, ok := result[mr.SourceChain]; !ok {
			result[mr.SourceChain] = make([]exectypes.CommitData, 0)
		}
		onRampAddress, err := base64.StdEncoding.DecodeString(mr.OnRampAddressBase64)
		if err != nil {
			lggr.Errorw("error decoding base64 encoded onRampAddress", "err", err, "addr", mr.OnRampAddressBase64)
			continue
		}
		result[mr.SourceChain] = append(result[mr.SourceChain], exectypes.CommitData{
			SourceChain:         mr.SourceChain,
			OnRampAddress:       onRampAddress,
			Timestamp:           mr.Timestamp,
			BlockNum:            mr.BlockNum,
			MerkleRoot:          mr.MerkleRoot,
			SequenceNumberRange: mr.SequenceNumberRange,
			ExecutedMessages:    executedMessages,
		})
	}

	for _, mr := range result {
		sort.Slice(mr, func(i, j int) bool { return mr[i].SequenceNumberRange.Start() < mr[j].SequenceNumberRange.Start() })
	}

	if len(result) == 0 {
		return nil
	}
	return result
}

// aggregateMerkleRootObservations groups the observations by merkle root and counts the votes for each.
// It returns two maps:
// 1. merkleRootsVotes: a map of merkle root data to the number of votes it received.
// 2. executedMsgVotes: a map of merkle root data to a map of executed messages and the number of votes they received.
func aggregateMerkleRootObservations(
	observations []plugincommon.AttributedObservation[exectypes.Observation],
) (map[merkleRootData]int, map[merkleRootData]map[cciptypes.SeqNum]int) {
	merkleRootsVotes := make(map[merkleRootData]int)
	executedMsgVotes := make(map[merkleRootData]map[cciptypes.SeqNum]int)

	for _, observation := range observations {
		for sourceChain, merkleRoots := range observation.Observation.CommitReports {
			for _, merkleRoot := range merkleRoots {
				data := merkleRootData{
					SourceChain:         sourceChain,
					OnRampAddressBase64: base64.StdEncoding.EncodeToString(merkleRoot.OnRampAddress),
					Timestamp:           merkleRoot.Timestamp,
					BlockNum:            merkleRoot.BlockNum,
					MerkleRoot:          merkleRoot.MerkleRoot,
					SequenceNumberRange: merkleRoot.SequenceNumberRange,
				}

				merkleRootsVotes[data]++

				if _, ok := executedMsgVotes[data]; !ok {
					executedMsgVotes[data] = make(map[cciptypes.SeqNum]int)
				}
				unqExecutedSeqNums := mapset.NewSet(merkleRoot.ExecutedMessages...).ToSlice()
				for _, executedSeqNum := range unqExecutedSeqNums {
					executedMsgVotes[data][executedSeqNum]++
				}
			}
		}
	}

	return merkleRootsVotes, executedMsgVotes
}

// merkleRootData is a helper comparable data structure for counting merkle root votes.
type merkleRootData struct {
	SourceChain         cciptypes.ChainSelector
	OnRampAddressBase64 string
	Timestamp           time.Time
	BlockNum            uint64
	MerkleRoot          cciptypes.Bytes32
	SequenceNumberRange cciptypes.SeqNumRange
}

// computeMessageHashesConsensus will iterate over observations of message hashes and come to consensus on them.
// We select a hash if it has at least f+1 votes. If more than one hashes exist with at
// least f+1 votes (reaching consensus threshold) then we log an error and skip that specific message.
func computeMessageHashesConsensus(
	lggr logger.Logger,
	observations []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) exectypes.MessageHashes {
	lggr = logger.With(lggr, "function", "computeMessageHashesConsensus", "fChain", fChain)

	// for each (chain, seqNum) pair keep a count of the votes for each hash
	type chainSeqNumPair struct {
		chain  cciptypes.ChainSelector
		seqNum cciptypes.SeqNum
	}

	hashVotes := make(map[chainSeqNumPair]map[cciptypes.Bytes32]int)
	for _, observation := range observations {
		for chain, hashes := range observation.Observation.Hashes {
			for seqNum, hash := range hashes {
				pair := chainSeqNumPair{chain: chain, seqNum: seqNum}
				if _, ok := hashVotes[pair]; !ok {
					hashVotes[pair] = make(map[cciptypes.Bytes32]int)
				}
				hashVotes[pair][hash]++
			}
		}
	}

	results := make(exectypes.MessageHashes)

	for chainSeqNum, hashesAndVotes := range hashVotes {
		f, ok := fChain[chainSeqNum.chain]
		if !ok {
			lggr.Warnw("no fChain defined, chain skipped from message hashes consensus", "chain", chainSeqNum.chain)
			continue
		}

		for hash, numVotes := range hashesAndVotes {
			if consensus.LtFPlusOne(f, numVotes) {
				lggr.Debugw("hash with less than f+1 votes was found, skipping it",
					"chain", chainSeqNum.chain, "seqNum", chainSeqNum.seqNum, "hash", hash, "votes", numVotes)
				continue
			}

			if _, chainResultsInitialized := results[chainSeqNum.chain]; !chainResultsInitialized {
				results[chainSeqNum.chain] = make(map[cciptypes.SeqNum]cciptypes.Bytes32)
			}

			existingHash, hashExists := results[chainSeqNum.chain][chainSeqNum.seqNum]
			results[chainSeqNum.chain][chainSeqNum.seqNum] = hash

			if hashExists {
				lggr.Errorw("more than one hash reached consensus for a message, message skipped",
					"chain", chainSeqNum.chain, "seqNum", chainSeqNum.seqNum, "hash1", existingHash, "hash2", hash)
				delete(results[chainSeqNum.chain], chainSeqNum.seqNum)
				if len(results[chainSeqNum.chain]) == 0 {
					delete(results, chainSeqNum.chain)
				}
				break
			}
		}
	}

	if len(results) == 0 {
		return nil
	}
	return results
}

func computeTokenDataObservationsConsensus(
	lggr logger.Logger,
	aos []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) exectypes.TokenDataObservations {
	// Single message can transfer multiple tokens, so we need to find consensus on the token level.
	validators :=
		make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]consensus.OracleMinObservation[exectypes.TokenData])
	results := make(exectypes.TokenDataObservations)

	for _, ao := range aos {
		for selector, seqMap := range ao.Observation.TokenData {
			f, ok := fChain[selector]
			if !ok {
				lggr.Warnw("no F defined for chain", "chain", selector)
				continue
			}

			if _, ok1 := results[selector]; !ok1 {
				results[selector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
			}

			if _, ok1 := validators[selector]; !ok1 {
				validators[selector] = make(map[reader.MessageTokenID]consensus.OracleMinObservation[exectypes.TokenData])
			}

			initResultsAndValidators(selector, f, seqMap, results, validators, ao.OracleID)
		}
	}

	for selector, seqNrs := range validators {
		for tokenID, validator := range seqNrs {
			mtd := results[selector][tokenID.SeqNr]
			if values := validator.GetValid(); len(values) == 1 {
				mtd = mtd.Append(tokenID.Index, values[0])
			} else {
				// Can't reach consensus for this particular token, marking it's as not ready
				mtd = mtd.Append(tokenID.Index, exectypes.NotReadyToken())
			}
			results[selector][tokenID.SeqNr] = mtd
		}
	}

	if len(results) == 0 {
		return nil
	}

	return results
}

func initResultsAndValidators(
	selector cciptypes.ChainSelector,
	f int,
	seqMap map[cciptypes.SeqNum]exectypes.MessageTokenData,
	results exectypes.TokenDataObservations,
	validators map[cciptypes.ChainSelector]map[reader.MessageTokenID]consensus.OracleMinObservation[exectypes.TokenData],
	oracleID commontypes.OracleID,
) {
	for seqNr, msgTokenData := range seqMap {
		if _, ok := results[selector][seqNr]; !ok {
			results[selector][seqNr] = exectypes.NewMessageTokenData()
		}

		for tokenIndex, tokenData := range msgTokenData.TokenData {
			messageTokenID := reader.NewMessageTokenID(seqNr, tokenIndex)
			if _, ok := validators[selector][messageTokenID]; !ok {
				validators[selector][messageTokenID] =
					consensus.NewOracleMinObservation(consensus.FPlus1(f), exectypes.TokenDataHash)
			}
			validators[selector][messageTokenID].Add(tokenData, oracleID)
		}
	}
}

// computeNoncesConsensus computes the consensus on the observed nonces.
// For each (chain, sender) pair we sort the observed nonces ascending and select the f-th observation.
func computeNoncesConsensus(
	lggr logger.Logger,
	observations []plugincommon.AttributedObservation[exectypes.Observation],
	fChainDest int,
) exectypes.NonceObservations {
	type chainSenderPair struct {
		chain  cciptypes.ChainSelector
		sender string
	}

	observedNonces := make(map[chainSenderPair][]uint64)
	for _, obs := range observations {
		for chain, nonces := range obs.Observation.Nonces {
			for sender, nonce := range nonces {
				pair := chainSenderPair{chain: chain, sender: sender}
				observedNonces[pair] = append(observedNonces[pair], nonce)
			}
		}
	}

	lggr.Debugw("computing nonces consensus",
		"observedNonces", observedNonces, "fChainDest", fChainDest)

	consensusNonces := make(exectypes.NonceObservations, len(observedNonces))
	for pair, nonces := range observedNonces {
		if len(nonces) == 0 || fChainDest >= len(nonces) {
			lggr.Debugw("no consensus on chain/sender pair",
				"chain", pair.chain, "sender", pair.sender, "nonces", nonces)
			continue
		}

		sort.Slice(nonces, func(i, j int) bool { return nonces[i] < nonces[j] })
		consensusNonce := nonces[fChainDest]

		if _, ok := consensusNonces[pair.chain]; !ok {
			consensusNonces[pair.chain] = make(map[string]uint64)
		}
		consensusNonces[pair.chain][pair.sender] = consensusNonce
	}

	return consensusNonces
}

// computeConsensusObservation aggregates multiple attributed observations to produce a single consensus observation.
// The provided f is required for computing the consensus on fChain prior to computing the observation consensus.
func computeConsensusObservation(
	lggr logger.Logger,
	observations []plugincommon.AttributedObservation[exectypes.Observation],
	destChain cciptypes.ChainSelector,
	f int,
) (exectypes.Observation, error) {
	fChain := getConsensusFChain(lggr, observations, f)

	destFChain, ok := fChain[destChain]
	if !ok {
		return exectypes.Observation{}, fmt.Errorf("destination chain %d is not in FChain %v", destChain, fChain)
	}

	consensusObservation := exectypes.NewObservation(
		computeCommitObservationsConsensus(lggr, observations, fChain),
		computeMessageObservationsConsensus(lggr, observations, fChain),
		computeTokenDataObservationsConsensus(lggr, observations, fChain),
		computeNoncesConsensus(lggr, observations, destFChain),
		dt.Observation{},
		computeMessageHashesConsensus(lggr, observations, fChain),
	)

	lggr.Debugw("computeConsensusObservation has finished computing the consensus observation",
		"fChain", fChain,
		"observations", observations,
		"destChain", destChain,
		"f", f,
		"consensusObservation", consensusObservation,
	)

	return consensusObservation, nil
}

// getConsensusFChain computes and the consensus (using 2F+1 threshold) on the observed fChain values for each chain.
func getConsensusFChain(
	lggr logger.Logger,
	observations []plugincommon.AttributedObservation[exectypes.Observation],
	f int,
) map[cciptypes.ChainSelector]int {
	observedFChains := make(map[cciptypes.ChainSelector][]int)
	for _, ao := range observations {
		for chain, f := range ao.Observation.FChain {
			observedFChains[chain] = append(observedFChains[chain], f)
		}
	}

	// consensus on the fChain map uses the role DON F value (all nodes can observe the home chain)
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(f))
	fChain := consensus.GetConsensusMap(lggr, "fChain", observedFChains, donThresh)

	return fChain
}
