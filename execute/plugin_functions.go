package execute

import (
	"errors"
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/commontypes"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
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
func groupByChainSelectorWithFilter(
	lggr logger.Logger,
	reports []plugintypes2.CommitPluginReportWithMeta,
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

func mergeMessageObservations(
	lggr logger.Logger,
	aos []plugincommon.AttributedObservation[exectypes.Observation], fChain map[cciptypes.ChainSelector]int,
) exectypes.MessageObservations {
	// Create a validator for each chain
	validators := make(map[cciptypes.ChainSelector]consensus.OracleMinObservation[cciptypes.Message])
	for selector, f := range fChain {
		validators[selector] = consensus.NewOracleMinObservation[cciptypes.Message](consensus.TwoFPlus1(f), nil)
	}

	// Add messages to the validator for each chain selector.
	for _, ao := range aos {
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

	results := make(exectypes.MessageObservations)
	for selector, validator := range validators {
		if msgs := validator.GetValid(); len(msgs) > 0 {
			if _, ok := results[selector]; !ok {
				results[selector] = make(map[cciptypes.SeqNum]cciptypes.Message)
			}
			for _, msg := range msgs {
				results[selector][msg.Header.SequenceNumber] = msg
			}
		}
	}

	if len(results) == 0 {
		return nil
	}

	return results
}

// mergeCommitObservations merges all observations which reach the fChain threshold into a single result.
// Any observations, or subsets of observations, which do not reach the threshold are ignored.
func mergeCommitObservations(
	lggr logger.Logger,
	aos []plugincommon.AttributedObservation[exectypes.Observation], fChain map[cciptypes.ChainSelector]int,
) exectypes.CommitObservations {
	// Create a validator for each chain
	validators := make(map[cciptypes.ChainSelector]consensus.OracleMinObservation[exectypes.CommitData])
	for selector, f := range fChain {
		validators[selector] =
			consensus.NewOracleMinObservation[exectypes.CommitData](consensus.TwoFPlus1(f), nil)
	}

	// Add reports to the validator for each chain selector.
	for _, ao := range aos {
		for selector, commitReports := range ao.Observation.CommitReports {
			validator, ok := validators[selector]
			if !ok {
				lggr.Warnw("no F defined for chain", "chain", selector)
				continue
			}
			// Add reports
			for _, commitReport := range commitReports {
				validator.Add(commitReport, ao.OracleID)
			}
		}
	}

	results := make(exectypes.CommitObservations)
	for selector, validator := range validators {
		if values := validator.GetValid(); len(values) > 0 {
			results[selector] = validator.GetValid()
		}
	}

	if len(results) == 0 {
		return nil
	}

	return results
}

func mergeMessageHashes(
	lggr logger.Logger,
	aos []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) exectypes.MessageHashes {
	// Single message can transfer multiple tokens, so we need to find consensus on the token level.
	validators := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]consensus.OracleMinObservation[cciptypes.Bytes32])
	results := make(exectypes.MessageHashes)

	for _, ao := range aos {
		for selector, seqMap := range ao.Observation.Hashes {
			f, ok := fChain[selector]
			if !ok {
				lggr.Warnw("no F defined for chain", "chain", selector)
				continue
			}

			if _, ok1 := results[selector]; !ok1 {
				results[selector] = make(map[cciptypes.SeqNum]cciptypes.Bytes32)
			}

			if _, ok1 := validators[selector]; !ok1 {
				validators[selector] = make(map[cciptypes.SeqNum]consensus.OracleMinObservation[cciptypes.Bytes32])
			}

			for seqNr, hash := range seqMap {
				if _, ok := validators[selector][seqNr]; !ok {
					validators[selector][seqNr] =
						consensus.NewOracleMinObservation[cciptypes.Bytes32](consensus.TwoFPlus1(f), nil)
				}
				validators[selector][seqNr].Add(hash, ao.OracleID)
			}

		}
	}

	for selector, seqNumValidator := range validators {
		for seqNum, validator := range seqNumValidator {
			if hashes := validator.GetValid(); len(hashes) == 1 {
				results[selector][seqNum] = hashes[0]
			}
		}
	}

	if len(results) == 0 {
		return nil
	}

	return results
}

func mergeTokenObservations(
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
					consensus.NewOracleMinObservation(consensus.TwoFPlus1(f), exectypes.TokenDataHash)
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

// getConsensusObservation merges all attributed observations into a single observation based on which values have
// consensus among the observers.
func getConsensusObservation(
	lggr logger.Logger,
	aos []plugincommon.AttributedObservation[exectypes.Observation],
	destChainSelector cciptypes.ChainSelector,
	F int,
	destChain cciptypes.ChainSelector,
) (exectypes.Observation, error) {
	observedFChains := make(map[cciptypes.ChainSelector][]int)
	for _, ao := range aos {
		obs := ao.Observation
		for chain, f := range obs.FChain {
			observedFChains[chain] = append(observedFChains[chain], f)
		}
	}
	// consensus on the fChain map uses the role DON F value
	// because all nodes can observe the home chain.
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(F))
	fChain := consensus.GetConsensusMap(lggr, "fChain", observedFChains, donThresh)
	_, ok := fChain[destChain] // check if the destination chain is in the FChain.
	if !ok {
		return exectypes.Observation{}, fmt.Errorf("destination chain %d is not in FChain", destChain)
	}

	lggr.Debugw("getConsensusObservation decoded observations", "aos", aos)

	mergedCommitObservations := mergeCommitObservations(lggr, aos, fChain)

	lggr.Debugw("merged commit observations", "mergedCommitObservations", mergedCommitObservations)

	mergedMessageObservations := mergeMessageObservations(lggr, aos, fChain)
	lggr.Debugw("merged message observations", "mergedMessageObservations", mergedMessageObservations)

	mergedTokenObservations := mergeTokenObservations(lggr, aos, fChain)

	lggr.Debugw("merged token data observations", "mergedTokenObservations", mergedTokenObservations)

	mergedHashes := mergeMessageHashes(lggr, aos, fChain)
	lggr.Debugw("merged message hashes", "mergedHashes", mergedHashes)

	agreedNoncesAmongOracles := computeNoncesConsensus(lggr, aos, fChain[destChainSelector])
	lggr.Debugw("agreed nonces among oracles", "agreedNoncesAmongOracles", agreedNoncesAmongOracles)

	observation := exectypes.NewObservation(
		mergedCommitObservations,
		mergedMessageObservations,
		mergedTokenObservations,
		agreedNoncesAmongOracles,
		dt.Observation{},
		mergedHashes,
	)

	return observation, nil
}
