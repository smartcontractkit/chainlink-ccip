package execute

import (
	"errors"
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// validateCommitReportsReadingEligibility validates that all commit reports' source chains are supported by observer
func validateCommitReportsReadingEligibility(
	supportedChains mapset.Set[cciptypes.ChainSelector],
	observedData exectypes.CommitObservations,
) error {
	for chainSel := range observedData {
		if !supportedChains.Contains(chainSel) {
			return fmt.Errorf("observer not allowed to read from chain %d", chainSel)
		}
		for _, data := range observedData[chainSel] {
			if data.SourceChain != chainSel {
				return fmt.Errorf("observer not allowed to read from chain %d", data.SourceChain)
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

// validateCostlyMessagesObservations validates that all costly messages belong to already observed messages
func validateCostlyMessagesObservations(
	observedMsgs exectypes.MessageObservations,
	costlyMessages []cciptypes.Bytes32,
) error {
	msgs := observedMsgs.Flatten()
	msgsIDMap := make(map[cciptypes.Bytes32]struct{})
	for _, msg := range msgs {
		msgsIDMap[msg.Header.MessageID] = struct{}{}
	}
	for _, id := range costlyMessages {
		if _, ok := msgsIDMap[id]; !ok {
			return fmt.Errorf("costly message %s not found in observed messages", id)
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
		_, ok := hashes[chain]
		if !ok {
			return fmt.Errorf("hash not found for chain %d", chain)
		}

		for seq, msg := range msgs {
			if _, ok := hashes[chain][seq]; !ok {
				return fmt.Errorf("hash not found for message %s", msg)
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
	msgsCount := 0
	for chain, report := range observedData {
		for _, data := range report {
			msgsMap, ok := observedMsgs[chain]
			if !ok {
				return fmt.Errorf("no messages observed for chain %d, while report was observed", chain)
			}

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
	observedData map[cciptypes.ChainSelector][]exectypes.CommitData,
) error {
	for _, commitData := range observedData {
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

// groupByChainSelector groups the reports by their chain selector and remove disabled chains from the result.
func groupByChainSelector(
	reports []plugintypes2.CommitPluginReportWithMeta) exectypes.CommitObservations {
	commitReportCache := make(map[cciptypes.ChainSelector][]exectypes.CommitData)
	for _, report := range reports {
		for _, singleReport := range report.Report.MerkleRoots {
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
	return commitReportCache
}

// filterOutExecutedMessages returns a new reports slice with fully executed messages removed.
// Reports that have all of their messages executed are not included in the result.
// The provided reports must be sorted by sequence number range starting sequence number.
func filterOutExecutedMessages(
	reports []exectypes.CommitData, executedMessages []cciptypes.SeqNum) []exectypes.CommitData {
	if len(executedMessages) == 0 {
		return reports
	}

	// filtered contains the reports with fully executed messages removed
	// and the executed messages appended to the report sorted by sequence number.
	var filtered []exectypes.CommitData

	for i, report := range reports {
		reportRange := report.SequenceNumberRange

		executedMsgsInReportRange := reportRange.FilterSlice(executedMessages)
		if len(executedMsgsInReportRange) == reportRange.Length() { // skip fully executed report.
			continue
		}

		sort.Slice(executedMsgsInReportRange, func(i, j int) bool {
			return executedMsgsInReportRange[i] < executedMsgsInReportRange[j]
		})
		report.ExecutedMessages = append(reports[i].ExecutedMessages, executedMsgsInReportRange...)
		filtered = append(filtered, report)
	}

	return filtered
}

func decodeAttributedObservations(
	aos []types.AttributedObservation,
) ([]plugincommon.AttributedObservation[exectypes.Observation], error) {
	decoded := make([]plugincommon.AttributedObservation[exectypes.Observation], len(aos))
	for i, ao := range aos {
		observation, err := exectypes.DecodeObservation(ao.Observation)
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
	validators := make(map[cciptypes.ChainSelector]consensus.MinObservation[cciptypes.Message])
	for selector, f := range fChain {
		validators[selector] = consensus.NewMinObservation[cciptypes.Message](consensus.FPlus1(f), nil)
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
				validator.Add(msg)
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
	validators := make(map[cciptypes.ChainSelector]consensus.MinObservation[exectypes.CommitData])
	for selector, f := range fChain {
		validators[selector] =
			consensus.NewMinObservation[exectypes.CommitData](consensus.FPlus1(f), nil)
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
				validator.Add(commitReport)
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
	validators := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]consensus.MinObservation[cciptypes.Bytes32])
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
				validators[selector] = make(map[cciptypes.SeqNum]consensus.MinObservation[cciptypes.Bytes32])
			}

			for seqNr, hash := range seqMap {
				if _, ok := validators[selector][seqNr]; !ok {
					validators[selector][seqNr] =
						consensus.NewMinObservation[cciptypes.Bytes32](consensus.FPlus1(f), nil)
				}
				validators[selector][seqNr].Add(hash)
			}

		}
	}

	for selector, seqNrs := range validators {
		for seqNum, validator := range seqNrs {
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
	validators := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]consensus.MinObservation[exectypes.TokenData])
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
				validators[selector] = make(map[reader.MessageTokenID]consensus.MinObservation[exectypes.TokenData])
			}

			initResultsAndValidators(selector, f, seqMap, results, validators)
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
	validators map[cciptypes.ChainSelector]map[reader.MessageTokenID]consensus.MinObservation[exectypes.TokenData],
) {
	for seqNr, msgTokenData := range seqMap {
		if _, ok := results[selector][seqNr]; !ok {
			results[selector][seqNr] = exectypes.NewMessageTokenData()
		}

		for tokenIndex, tokenData := range msgTokenData.TokenData {
			messageTokenID := reader.NewMessageTokenID(seqNr, tokenIndex)
			if _, ok := validators[selector][messageTokenID]; !ok {
				validators[selector][messageTokenID] =
					consensus.NewMinObservation[exectypes.TokenData](consensus.FPlus1(f), exectypes.TokenDataHash)
			}
			validators[selector][messageTokenID].Add(tokenData)
		}
	}
}

// mergeNonceObservations merges all observations which reach the fChain threshold into a single result.
// Any observations, or subsets of observations, which do not reach the threshold are ignored.
func mergeNonceObservations(
	daos []plugincommon.AttributedObservation[exectypes.Observation],
	fChainDest int,
) exectypes.NonceObservations {
	// Nonces store context in a map key, so a different container type is needed for the observation filter.
	type NonceTriplet struct {
		source cciptypes.ChainSelector
		sender []byte
		nonce  uint64
	}

	// Create one validator because nonces are only observed from the destination chain.
	validator := consensus.NewMinObservation[NonceTriplet](consensus.FPlus1(fChainDest), nil)

	// Add reports to the validator for each chain selector.
	for _, ao := range daos {
		if len(ao.Observation.Nonces) == 0 {
			continue
		}

		for sourceSelector, nonces := range ao.Observation.Nonces {
			for sender, nonce := range nonces {
				validator.Add(NonceTriplet{
					source: sourceSelector,
					sender: []byte(sender),
					nonce:  nonce,
				})
			}
		}
	}

	// Convert back to the observation format.
	results := make(exectypes.NonceObservations)
	nonces := validator.GetValid()
	for _, nonce := range nonces {
		if _, ok := results[nonce.source]; !ok {
			results[nonce.source] = make(map[string]uint64)
		}
		results[nonce.source][string(nonce.sender)] = nonce.nonce
	}

	if len(results) == 0 {
		return nil
	}

	return results
}

// mergeCostlyMessages merges all costly message observations. A message is considered costly if it is observed by more
// than `fChainDest` observers.
func mergeCostlyMessages(
	aos []plugincommon.AttributedObservation[exectypes.Observation],
	fChainDest int,
) []cciptypes.Bytes32 {
	costlyMessages := mapset.NewSet[cciptypes.Bytes32]()
	counts := make(map[cciptypes.Bytes32]int)
	for _, ao := range aos {
		for _, costlyMessage := range ao.Observation.CostlyMessages {
			counts[costlyMessage]++
			if consensus.GteFPlusOne(fChainDest, counts[costlyMessage]) {
				costlyMessages.Add(costlyMessage)
			}
		}
	}

	if costlyMessages.Cardinality() == 0 {
		return nil
	}

	return costlyMessages.ToSlice()
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
	if len(aos) < F {
		return exectypes.Observation{}, fmt.Errorf("below F threshold")
	}

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

	mergedCostlyMessages := mergeCostlyMessages(aos, fChain[destChainSelector])
	lggr.Debugw("merged costly messages", "mergedCostlyMessages", mergedCostlyMessages)

	mergedNonceObservations :=
		mergeNonceObservations(aos, fChain[destChainSelector])
	lggr.Debugw("merged nonce observations", "mergedNonceObservations", mergedNonceObservations)

	observation := exectypes.NewObservation(
		mergedCommitObservations,
		mergedMessageObservations,
		mergedCostlyMessages,
		mergedTokenObservations,
		mergedNonceObservations,
		dt.Observation{},
		mergedHashes,
	)

	return observation, nil
}
