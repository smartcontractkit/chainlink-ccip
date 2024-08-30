package execute

import (
	"errors"
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/validation"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
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
	reports []plugintypes.CommitPluginReportWithMeta) exectypes.CommitObservations {
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

type decodedAttributedObservation struct {
	Observation exectypes.Observation
	Observer    commontypes.OracleID
}

func decodeAttributedObservations(aos []types.AttributedObservation) ([]decodedAttributedObservation, error) {
	decoded := make([]decodedAttributedObservation, len(aos))
	for i, ao := range aos {
		observation, err := exectypes.DecodeObservation(ao.Observation)
		if err != nil {
			return nil, err
		}
		decoded[i] = decodedAttributedObservation{
			Observation: observation,
			Observer:    ao.Observer,
		}
	}
	return decoded, nil
}

func mergeMessageObservations(
	aos []decodedAttributedObservation, fChain map[cciptypes.ChainSelector]int,
) (exectypes.MessageObservations, error) {
	// Create a validator for each chain
	validators := make(map[cciptypes.ChainSelector]validation.MinObservationFilter[cciptypes.Message])
	for selector, f := range fChain {
		validators[selector] = validation.NewMinObservationValidator[cciptypes.Message](f+1, nil)
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
	aos []decodedAttributedObservation, fChain map[cciptypes.ChainSelector]int,
) (exectypes.CommitObservations, error) {
	// Create a validator for each chain
	validators := make(map[cciptypes.ChainSelector]validation.MinObservationFilter[exectypes.CommitData])
	for selector, f := range fChain {
		validators[selector] =
			validation.NewMinObservationValidator[exectypes.CommitData](f+1, nil)
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

// mergeNonceObservations merges all observations which reach the fChain threshold into a single result.
// Any observations, or subsets of observations, which do not reach the threshold are ignored.
func mergeNonceObservations(
	daos []decodedAttributedObservation, dest cciptypes.ChainSelector, fChainDest int,
) (exectypes.NonceObservations, error) {
	// Nonces store context in a map key, so a different container type is needed for the observation filter.
	type NonceTriplet struct {
		source cciptypes.ChainSelector
		sender []byte
		nonce  uint64
	}

	// Create one validator because nonces are only observed from the destination chain.
	validator := validation.NewMinObservationValidator[NonceTriplet](fChainDest, nil)

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
		return nil, nil
	}

	return results, nil
}

// getConsensusObservation merges all attributed observations into a single observation based on which values have
// consensus among the observers.
func getConsensusObservation(
	lggr logger.Logger,
	aos []types.AttributedObservation,
	oracleID commontypes.OracleID,
	destChainSelector cciptypes.ChainSelector,
	F int,
	fChain map[cciptypes.ChainSelector]int,
) (exectypes.Observation, error) {
	decodedObservations, err := decodeAttributedObservations(aos)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to decode observations: %w", err)
	}
	if len(decodedObservations) < F {
		return exectypes.Observation{}, fmt.Errorf("below F threshold")
	}

	lggr.Debugw(
		fmt.Sprintf("[oracle %d] exec outcome: decoded observations", oracleID),
		"oracle", oracleID,
		"decodedObservations", decodedObservations)

	mergedCommitObservations, err := mergeCommitObservations(decodedObservations, fChain)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to merge commit report observations: %w", err)
	}
	lggr.Debugw(
		fmt.Sprintf("[oracle %d] exec outcome: merged commit observations", oracleID),
		"oracle", oracleID,
		"mergedCommitObservations", mergedCommitObservations)

	mergedMessageObservations, err := mergeMessageObservations(decodedObservations, fChain)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to merge message observations: %w", err)
	}
	lggr.Debugw(
		fmt.Sprintf("[oracle %d] exec outcome: merged message observations", oracleID),
		"oracle", oracleID,
		"mergedMessageObservations", mergedMessageObservations)

	mergedNonceObservations, err := mergeNonceObservations(decodedObservations, destChainSelector, fChain[destChainSelector])
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to merge nonce observations: %w", err)
	}
	lggr.Debugw(
		fmt.Sprintf("[oracle %d] exec outcome: merged nonce observations", oracleID),
		"oracle", oracleID,
		"mergedNonceObservations", mergedNonceObservations)

	observation := exectypes.NewObservation(
		mergedCommitObservations,
		mergedMessageObservations,
		mergedNonceObservations)

	return observation, nil
}
