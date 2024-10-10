package execute

import (
	"errors"
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
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
// nolint:gocyclo // todo
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
				// need to go to the next set of executed messages.
				break
			}

			if executed.End() < reportRange.Start() {
				// add report that has non-executed messages.
				reportIdx++
				filtered = append(filtered, reports[i])
				continue
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

func mergeTokenObservations(
	aos []plugincommon.AttributedObservation[exectypes.Observation],
	fChain map[cciptypes.ChainSelector]int,
) (exectypes.TokenDataObservations, error) {
	// Single message can transfer multiple tokens, so we need to find consensus on the token level.
	//nolint:lll
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
			if counts[costlyMessage] >= int(consensus.FPlus1(fChainDest)) {
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
