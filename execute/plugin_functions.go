package execute

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"golang.org/x/exp/maps"

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

// validateObserverReadingEligibility checks if the observer is eligible to observe the messages it observed.
func validateObserverReadingEligibility(
	supportedChains mapset.Set[cciptypes.ChainSelector],
	observedMsgs exectypes.MessageObservations,
) error {
	// TODO: validate that CommitReports and Nonces are only observed if the destChain is supported.

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

func validateTokenDataObservations(observedMsgs exectypes.MessageObservations, tokenData exectypes.TokenDataObservations) error {

	if len(observedMsgs) != len(tokenData) {
		return fmt.Errorf("unexpected number of token data observations: expected %d, got %d",
			len(observedMsgs), len(tokenData))
	}

	for chain, msgs := range observedMsgs {
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
		for seq, msg := range msgs {
			if _, ok := hashes[chain][seq]; !ok {
				return fmt.Errorf("hash not found for message %s", msg)
			}
		}
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

var errOverlappingRanges = errors.New("overlapping sequence numbers in reports")

// computeRanges takes a slice of reports and computes the smallest number of contiguous ranges
// that cover all the sequence numbers in the reports.
// Note: reports need all messages to create a proof even if some are already executed.
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
		} else if report.SequenceNumberRange.Start() < seqRange.End() {
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

func groupByChainSelector(
	reports []plugintypes2.CommitPluginReportWithMeta) exectypes.CommitObservations {
	commitReportCache := make(map[cciptypes.ChainSelector][]exectypes.CommitData)
	for _, report := range reports {
		for _, singleReport := range report.Report.MerkleRoots {
			commitReportCache[singleReport.ChainSel] = append(commitReportCache[singleReport.ChainSel],
				exectypes.CommitData{
					SourceChain:         singleReport.ChainSel,
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
// Unordered inputs are supported.
func filterOutExecutedMessages(
	reports []exectypes.CommitData, executedMessages []cciptypes.SeqNumRange,
) ([]exectypes.CommitData, error) {
	sort.Slice(reports, func(i, j int) bool {
		return reports[i].SequenceNumberRange.Start() < reports[j].SequenceNumberRange.Start()
	})

	// If none are executed, return the (sorted) input.
	if len(executedMessages) == 0 {
		return reports, nil
	}

	sort.Slice(executedMessages, func(i, j int) bool {
		return executedMessages[i].Start() < executedMessages[j].Start()
	})

	// Make sure they do not overlap
	previousMax := cciptypes.SeqNum(0)
	for _, seqRange := range executedMessages {
		if seqRange.Start() < previousMax {
			return nil, errOverlappingRanges
		}
		previousMax = seqRange.End()
	}

	var filtered []exectypes.CommitData

	reportIdx := 0
	for _, executed := range executedMessages {
		for i := reportIdx; i < len(reports); i++ {
			reportRange := reports[i].SequenceNumberRange
			if executed.End() < reportRange.Start() {
				// No messages in current executed range are in reports[i]
				// need to go to the next set of executed range.
				break
			}

			if reportRange.Start() >= executed.Start() && reportRange.End() <= executed.End() {
				// skip fully executed report.
				reportIdx++
				continue
			}

			s := executed.Start()
			if reportRange.Start() > executed.Start() {
				s = reportRange.Start()
			}
			for ; s <= executed.End(); s++ {
				// This range runs into the next report.
				if s > reports[i].SequenceNumberRange.End() {
					reportIdx++
					filtered = append(filtered, reports[i])
					break
				}
				reports[i].ExecutedMessages = append(reports[i].ExecutedMessages, s)
			}
		}
	}

	// Add any remaining reports that were not fully executed.
	for i := reportIdx; i < len(reports); i++ {
		filtered = append(filtered, reports[i])
	}

	return filtered, nil
}

// truncateObservation truncates the observation to fit within the given maxSize after encoding.
// It removes data from the observation in the following order:
// For each chain, remove commit reports one by one, if the encoded observation is still too large,
// remove the entire chain.
// Keep repeating this process until the encoded observation fits within the maxSize or there's only
// one chain with one report left
// Note: This function doesn't split one report into multiple parts.
func truncateObservation(
	obs exectypes.Observation,
	maxSize int,
) (exectypes.Observation, error) {
	observation := obs
	encodedObs, err := observation.Encode()
	if err != nil {
		return exectypes.Observation{}, err
	}

	chains := maps.Keys(observation.CommitReports)
	sort.Slice(chains, func(i, j int) bool {
		return chains[i] < chains[j]
	})

	// If the encoded observation is too large, start filtering data.
	for len(encodedObs) > maxSize {
		for _, chain := range chains {
			commits := observation.CommitReports[chain]
			if len(commits) == 0 {
				continue
			}
			lastCommit := &commits[len(commits)-1]
			seqNum := lastCommit.SequenceNumberRange.Start()

			observation.PseudoDeleted[chain] = make(map[cciptypes.SeqNum]bool)
			for seqNum <= lastCommit.SequenceNumberRange.End() {
				if _, ok := observation.Messages[chain][seqNum]; !ok {
					return exectypes.Observation{}, fmt.Errorf("missing message with seqNr %d from chain %d", seqNum, chain)
				}
				observation.Messages[chain][seqNum] = PseudoDeleteMessage(observation.Messages[chain][seqNum])
				observation.PseudoDeleted[chain][seqNum] = true

				if _, ok := observation.TokenData[chain][seqNum]; !ok {
					return exectypes.Observation{}, fmt.Errorf(
						"missing tokenData for message with seqNr %d from chain %d", seqNum, chain,
					)
				}
				observation.TokenData[chain][seqNum] = PseudoDeleteTokenData(observation.TokenData[chain][seqNum])

				seqNum++
				if observationFitsSize(observation, maxSize) {
					return observation, nil
				}
			}

			// Reaching here means that all messages in the report are truncated, truncate the last commit
			observation = truncateLastCommit(observation, chain)

			if len(observation.CommitReports[chain]) == 0 {
				// If the last commit report was truncated, truncate the chain
				observation = truncateChain(observation, chain)
			}
			chains = maps.Keys(observation.CommitReports)
		}
		// Truncated all chains.
		if len(observation.CommitReports) == 0 {
			return exectypes.Observation{}, fmt.Errorf("no more data to truncate")
		}
		encodedObs, err = observation.Encode()
		if err != nil {
			return exectypes.Observation{}, nil
		}
	}

	return observation, nil
}

func observationFitsSize(obs exectypes.Observation, maxSize int) bool {
	encodedObs, err := obs.Encode()
	if err != nil {
		return false
	}
	return len(encodedObs) <= maxSize
}
func PseudoDeleteMessage(msg cciptypes.Message) cciptypes.Message {
	pseudoDeletedMsg := msg

	pseudoDeletedMsg.Data = nil
	pseudoDeletedMsg.ExtraArgs = nil
	pseudoDeletedMsg.TokenAmounts = nil

	return pseudoDeletedMsg
}

func PseudoDeleteTokenData(tokenData exectypes.MessageTokenData) exectypes.MessageTokenData {
	msgTokenData := tokenData
	for _, td := range msgTokenData.TokenData {
		td.Data = nil
		td.Ready = false
	}
	return msgTokenData
}

// truncateLastCommit removes the last commit from the observation.
// errors if there are no commits to truncate.
func truncateLastCommit(
	obs exectypes.Observation,
	chain cciptypes.ChainSelector,
) exectypes.Observation {
	observation := obs
	commits := observation.CommitReports[chain]
	if len(commits) == 0 {
		return observation
	}
	lastCommit := commits[len(commits)-1]
	// Remove the last commit from the list.
	commits = commits[:len(commits)-1]
	observation.CommitReports[chain] = commits
	for seqNum, msg := range observation.Messages[chain] {
		if lastCommit.SequenceNumberRange.Contains(seqNum) {
			// Remove the message from the observation.
			delete(observation.Messages[chain], seqNum)
			// Remove the token data from the observation.
			delete(observation.TokenData[chain], seqNum)
			// Remove costly messages
			for i, costlyMessage := range observation.CostlyMessages {
				if costlyMessage == msg.Header.MessageID {
					observation.CostlyMessages = append(observation.CostlyMessages[:i], observation.CostlyMessages[i+1:]...)
				}
			}
			// Leaving Nonces untouched
		}
	}

	return observation
}

// truncateChain removes all data related to the given chain from the observation.
// returns true if the chain was found and truncated, false otherwise.
func truncateChain(
	obs exectypes.Observation,
	chain cciptypes.ChainSelector,
) exectypes.Observation {
	observation := obs
	if _, ok := observation.CommitReports[chain]; !ok {
		return observation
	}
	messageIDs := make(map[cciptypes.Bytes32]struct{})
	// To remove costly message IDs we need to iterate over all messages and find the ones that belong to the chain.
	for _, seqNumMap := range observation.Messages {
		for _, message := range seqNumMap {
			messageIDs[message.Header.MessageID] = struct{}{}
		}
	}

	deleteCostlyMessages := func() {
		for i, costlyMessage := range observation.CostlyMessages {
			if _, ok := messageIDs[costlyMessage]; ok {
				observation.CostlyMessages = append(observation.CostlyMessages[:i], observation.CostlyMessages[i+1:]...)
			}
		}
	}

	delete(observation.CommitReports, chain)
	delete(observation.Messages, chain)
	delete(observation.TokenData, chain)
	delete(observation.Nonces, chain)
	deleteCostlyMessages()

	return observation
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
	aos []plugincommon.AttributedObservation[exectypes.Observation], fChain map[cciptypes.ChainSelector]int,
) (exectypes.MessageObservations, error) {
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
				return exectypes.MessageObservations{}, fmt.Errorf("no validator for chain %d", selector)
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
		return nil, nil
	}

	return results, nil
}

// mergeCommitObservations merges all observations which reach the fChain threshold into a single result.
// Any observations, or subsets of observations, which do not reach the threshold are ignored.
func mergeCommitObservations(
	aos []plugincommon.AttributedObservation[exectypes.Observation], fChain map[cciptypes.ChainSelector]int,
) (exectypes.CommitObservations, error) {
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
				return exectypes.CommitObservations{}, fmt.Errorf("no validator for chain %d", selector)
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
		return nil, nil
	}

	return results, nil
}

func mergePseudoDeletedMessages(
	aos []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) (exectypes.PseudoDeletedMessages, error) {
	// Single message can transfer multiple tokens, so we need to find consensus on the token level.
	validators := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]consensus.MinObservation[bool])
	results := make(exectypes.PseudoDeletedMessages)

	for _, ao := range aos {
		for selector, seqMap := range ao.Observation.PseudoDeleted {
			f, ok := fChain[selector]
			if !ok {
				return exectypes.PseudoDeletedMessages{}, fmt.Errorf("no F defined for chain %d", selector)
			}

			if _, ok1 := results[selector]; !ok1 {
				results[selector] = make(map[cciptypes.SeqNum]bool)
			}

			if _, ok1 := validators[selector]; !ok1 {
				validators[selector] = make(map[cciptypes.SeqNum]consensus.MinObservation[bool])
			}

			for seqNr, deleted := range seqMap {
				if _, ok := validators[selector][seqNr]; !ok {
					validators[selector][seqNr] =
						consensus.NewMinObservation[bool](consensus.FPlus1(f), nil)
				}
				validators[selector][seqNr].Add(deleted)
			}

		}
	}

	for selector, seqNrs := range validators {
		for seqNum, validator := range seqNrs {
			if deleted := validator.GetValid(); len(deleted) == 1 {
				results[selector][seqNum] = deleted[0]
			}
		}
	}

	if len(results) == 0 {
		return nil, nil
	}

	return results, nil
}

func mergeMessageHashes(
	aos []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) (exectypes.MessageHashes, error) {
	// Single message can transfer multiple tokens, so we need to find consensus on the token level.
	validators := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]consensus.MinObservation[cciptypes.Bytes32])
	results := make(exectypes.MessageHashes)

	for _, ao := range aos {
		for selector, seqMap := range ao.Observation.Hashes {
			f, ok := fChain[selector]
			if !ok {
				return exectypes.MessageHashes{}, fmt.Errorf("no F defined for chain %d", selector)
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
		return nil, nil
	}

	return results, nil
}

func mergeTokenObservations(
	aos []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) (exectypes.TokenDataObservations, error) {
	// Single message can transfer multiple tokens, so we need to find consensus on the token level.
	validators := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]consensus.MinObservation[exectypes.TokenData])
	results := make(exectypes.TokenDataObservations)

	for _, ao := range aos {
		for selector, seqMap := range ao.Observation.TokenData {
			f, ok := fChain[selector]
			if !ok {
				return exectypes.TokenDataObservations{}, fmt.Errorf("no F defined for chain %d", selector)
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
		return nil, nil
	}

	return results, nil
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
	fChain map[cciptypes.ChainSelector]int,
) (exectypes.Observation, error) {
	if len(aos) < F {
		return exectypes.Observation{}, fmt.Errorf("below F threshold")
	}

	lggr.Debugw("getConsensusObservation decoded observations", "aos", aos)

	mergedCommitObservations, err := mergeCommitObservations(aos, fChain)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to merge commit report observations: %w", err)
	}
	lggr.Debugw("merged commit observations", "mergedCommitObservations", mergedCommitObservations)

	mergedMessageObservations, err := mergeMessageObservations(aos, fChain)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to merge message observations: %w", err)
	}
	lggr.Debugw("merged message observations", "mergedMessageObservations", mergedMessageObservations)

	mergedTokenObservations, err := mergeTokenObservations(aos, fChain)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to merge token data observations: %w", err)
	}
	lggr.Debugw("merged token data observations", "mergedTokenObservations", mergedTokenObservations)

	mergedHashes, err := mergeMessageHashes(aos, fChain)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to merge message hashes: %w", err)
	}
	lggr.Debugw("merged message hashes", "mergedHashes", mergedHashes)

	mergedPseudoDeletedMessages, err := mergePseudoDeletedMessages(aos, fChain)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to merge pseudo deleted messages: %w", err)
	}
	lggr.Debugw("merged pseudo deleted messages", "mergedPseudoDeletedMessages", mergedPseudoDeletedMessages)

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
		mergedPseudoDeletedMessages,
	)

	return observation, nil
}

// getMessageTimestampMap returns a map of message IDs to their timestamps.
// cciptypes.Message does not contain a timestamp, so we need to derive the timestamp from the commit data.
func getMessageTimestampMap(
	commitReportCache map[cciptypes.ChainSelector][]exectypes.CommitData,
	messages exectypes.MessageObservations,
) (map[cciptypes.Bytes32]time.Time, error) {
	messageTimestamps := make(map[cciptypes.Bytes32]time.Time)

	for chainSel, SeqNumToMsg := range messages {
		commitData, ok := commitReportCache[chainSel]
		if !ok {
			return nil, fmt.Errorf("missing commit data for chain %s", chainSel)
		}

		for seqNum, msg := range SeqNumToMsg {
			for _, commit := range commitData {
				if commit.SequenceNumberRange.Contains(seqNum) {
					messageTimestamps[msg.Header.MessageID] = commit.Timestamp
				}
			}
		}
	}

	return messageTimestamps, nil
}
