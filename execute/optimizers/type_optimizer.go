package optimizers

import (
	"sort"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"slices"
)

type ObservationOptimizer struct {
	maxEncodedSize    int
	emptyEncodedSizes EmptyEncodeSizes
	lggr              logger.Logger
	ocrTypeCodec      ocrtypecodec.ExecCodec
}

func NewObservationOptimizer(
	lggr logger.Logger, maxEncodedSize int, ocrTypeCodec ocrtypecodec.ExecCodec,
) ObservationOptimizer {
	return ObservationOptimizer{
		lggr:              logutil.WithComponent(lggr, "ObservationOptimizer"),
		maxEncodedSize:    maxEncodedSize,
		emptyEncodedSizes: NewEmptyEncodeSizes(),
		ocrTypeCodec:      ocrTypeCodec,
	}
}

type EmptyEncodeSizes struct {
	MessageAndTokenData int
	CommitData          int
	SeqNumMap           int
}

func NewEmptyEncodeSizes() EmptyEncodeSizes {
	emptyMsg := cciptypes.Message{}
	emptyTokenData := exectypes.MessageTokenData{}
	emptyCommitData := exectypes.CommitData{}
	emptySeqNrSize := internal.EncodedSize(make(map[cciptypes.SeqNum]cciptypes.Message))

	return EmptyEncodeSizes{
		MessageAndTokenData: internal.EncodedSize(emptyMsg) + internal.EncodedSize(emptyTokenData),
		CommitData:          internal.EncodedSize(emptyCommitData), // 305
		SeqNumMap:           emptySeqNrSize,                        // 2
	}
}

// TruncateObservation truncates the observation to fit within the given op.maxEncodedSize after encoding.
// It removes data from the observation in the following order:
// For each chain, pick last report and start removing messages one at a time.
// If removed all messages from the report, remove the report.
// If removed last report in the chain, remove the chain.
// After removing full report from a chain, move to the next chain and repeat. This ensures that we don't
// exclude messages from one chain only.
// Keep repeating this process until the encoded observation fits within the op.maxEncodedSize
// Important Note: We can't delete messages completely from single reports as we need them to create merkle proofs.
//
// Errors are returned if the observation cannot be encoded.
//
//nolint:gocyclo
func (op ObservationOptimizer) TruncateObservation(observation exectypes.Observation) (exectypes.Observation, error) {
	obs := observation
	encodedObs, err := op.ocrTypeCodec.EncodeObservation(obs)
	if err != nil {
		return exectypes.Observation{}, err
	}
	encodedObsSize := len(encodedObs)
	if encodedObsSize <= op.maxEncodedSize {
		return obs, nil
	}

	chains := maps.Keys(obs.CommitReports)
	slices.Sort(chains)

	messageAndTokenDataEncodedSizes := exectypes.GetEncodedMsgAndTokenDataSizes(obs.Messages, obs.TokenData)
	// While the encoded obs is too large, continue filtering data.
	for encodedObsSize > op.maxEncodedSize {
		// go through each chain and truncate observations for the final commit report.
		for _, chain := range chains {
			commits := obs.CommitReports[chain]
			if len(commits) == 0 {
				continue
			}
			lastCommit := &commits[len(commits)-1]
			seqNum := lastCommit.SequenceNumberRange.Start()
			// Remove messages one by one starting from the last message of the last commit report.
			for seqNum <= lastCommit.SequenceNumberRange.End() {
				if _, ok := obs.Messages[chain][seqNum]; !ok {
					op.lggr.Errorw("missing message", "seqNum", seqNum, "chain", chain)
					continue
				}
				op.lggr.Debugw("truncating message", "seqNum", seqNum, "chain", chain)
				obs.Messages[chain][seqNum] = cciptypes.Message{}
				obs.TokenData[chain][seqNum] = exectypes.NewMessageTokenData()
				// Subtract the removed message and token size
				encodedObsSize -= messageAndTokenDataEncodedSizes[chain][seqNum]
				// Add empty message and token encoded size
				encodedObsSize += op.emptyEncodedSizes.MessageAndTokenData
				seqNum++
				// Once we assert the estimation is less than the max size we double-check with actual encoding size.
				// Otherwise, we short circuit after checking the estimation only
				if encodedObsSize <= op.maxEncodedSize && fitsWithinSize(op.ocrTypeCodec, obs, op.maxEncodedSize) {
					return obs, nil
				}
			}

			op.lggr.Debugw("truncating last commit report", "chain", chain)
			var bytesTruncated int
			// Reaching here means that all messages in the report are truncated, truncate the last commit
			obs, bytesTruncated = op.truncateLastCommit(obs, chain)

			encodedObsSize -= bytesTruncated

			if len(obs.CommitReports[chain]) == 0 {
				op.lggr.Debugw("truncating chain", "chain", chain)
				// If the last commit report was truncated, truncate the chain
				obs = op.truncateChain(obs, chain)
			}

			// Once we assert the estimation is less than the max size we double-check with actual encoding size.
			// Otherwise, we short circuit after checking the estimation only
			if encodedObsSize <= op.maxEncodedSize && fitsWithinSize(op.ocrTypeCodec, obs, op.maxEncodedSize) {
				return obs, nil
			}
		}
		// Truncated all chains. Return obs as is. (it has other data like contract discovery)
		if len(obs.CommitReports) == 0 {
			return obs, nil
		}
		// Encoding again after doing a full iteration on all chains and removing messages/commits.
		// That is because using encoded sizes is not 100% accurate and there are some missing bytes in the calculation.
		encodedObs, err = op.ocrTypeCodec.EncodeObservation(obs)
		if err != nil {
			return exectypes.Observation{}, err
		}
		encodedObsSize = len(encodedObs)
	}

	return obs, nil
}

func fitsWithinSize(codec ocrtypecodec.ExecCodec, obs exectypes.Observation, maxEncodedSize int) bool {
	encodedObs, err := codec.EncodeObservation(obs)
	if err != nil {
		return false
	}
	return len(encodedObs) <= maxEncodedSize
}

// truncateLastCommit removes the last commit from the observation.
// returns observation and the number of bytes truncated.
func (op ObservationOptimizer) truncateLastCommit(
	obs exectypes.Observation,
	chain cciptypes.ChainSelector,
) (exectypes.Observation, int) {
	observation := obs
	bytesTruncated := 0
	commits := observation.CommitReports[chain]
	if len(commits) == 0 {
		return observation, bytesTruncated
	}
	lastCommit := commits[len(commits)-1]
	// Remove the last commit from the list.
	commits = commits[:len(commits)-1]
	observation.CommitReports[chain] = commits
	// Remove from the encoded size.
	bytesTruncated = bytesTruncated + op.emptyEncodedSizes.CommitData + 4 // brackets, and commas
	for seqNum := range observation.Messages[chain] {
		if lastCommit.SequenceNumberRange.Contains(seqNum) {
			// Remove the message from the observation.
			delete(observation.Messages[chain], seqNum)
			// Remove the token data from the observation.
			delete(observation.TokenData[chain], seqNum)
			//delete(observation.Hashes[chain], seqNum)
			// Remove the encoded size of the message and token data.
			bytesTruncated += op.emptyEncodedSizes.MessageAndTokenData
			bytesTruncated = bytesTruncated + 2*op.emptyEncodedSizes.SeqNumMap
			bytesTruncated += 4 // for brackets and commas
			// Leaving Nonces untouched
		}
	}

	return observation, bytesTruncated
}

// truncateChain removes all data related to the given chain from the observation.
// returns true if the chain was found and truncated, false otherwise.
func (op ObservationOptimizer) truncateChain(
	obs exectypes.Observation,
	chain cciptypes.ChainSelector,
) exectypes.Observation {
	observation := obs
	if _, ok := observation.CommitReports[chain]; !ok {
		return observation
	}

	delete(observation.CommitReports, chain)
	delete(observation.Messages, chain)
	delete(observation.TokenData, chain)
	delete(observation.Nonces, chain)

	return observation
}
