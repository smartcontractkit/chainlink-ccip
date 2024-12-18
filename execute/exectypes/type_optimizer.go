package exectypes

import (
	"encoding/json"
	"fmt"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"golang.org/x/exp/maps"
	"sort"
)

type ObservationOptimizer struct {
	maxEncodedSize    int
	emptyEncodedSizes EmptyEncodeSizes
}

type EmptyEncodeSizes struct {
	MessageAndTokenData int
	CommitData          int
	SeqNumMap           int
}

func NewEmptyEncodeSizes() EmptyEncodeSizes {
	emptyMsg := cciptypes.Message{}
	emptyTokenData := MessageTokenData{}
	emptyCommitData := CommitData{}
	emptySeqNr := make(map[cciptypes.SeqNum]cciptypes.Message)
	emptySeqNrSize := 0

	enc, err := json.Marshal(emptySeqNr)
	if err == nil {
		emptySeqNrSize = len(enc)
	}

	return EmptyEncodeSizes{
		MessageAndTokenData: emptyMsg.EncodedSize() + emptyTokenData.EncodedSize(), // 397 + 18 = 415
		CommitData:          emptyCommitData.EncodedSize(),                         // 305
		SeqNumMap:           emptySeqNrSize,                                        // 2
	}
}

// truncateObservation truncates the observation to fit within the given op.maxEncodedSize after encoding.
// It removes data from the observation in the following order:
// For each chain, pick last report and start removing messages one at a time.
// If removed all messages from the report, remove the report.
// If removed last report in the chain, remove the chain.
// After removing full report from a chain, move to the next chain and repeat. This ensures that we don't
// exclude messages from one chain only.
// Keep repeating this process until the encoded observation fits within the op.maxEncodedSize
// Important Note: We can't delete messages completely from single reports as we need them to create merkle proofs.
func (op ObservationOptimizer) truncateObservation(observation Observation) (Observation, error) {
	obs := observation
	encodedObs, err := obs.Encode()
	if err != nil {
		return Observation{}, err
	}
	encodedObsSize := len(encodedObs)
	if encodedObsSize <= op.maxEncodedSize {
		return obs, nil
	}

	chains := maps.Keys(obs.CommitReports)
	sort.Slice(chains, func(i, j int) bool {
		return chains[i] < chains[j]
	})

	// If the encoded obs is too large, start filtering data.
	for encodedObsSize > op.maxEncodedSize {
		for _, chain := range chains {
			commits := obs.CommitReports[chain]
			if len(commits) == 0 {
				continue
			}
			lastCommit := &commits[len(commits)-1]
			seqNum := lastCommit.SequenceNumberRange.Start()

			for seqNum <= lastCommit.SequenceNumberRange.End() {
				if _, ok := obs.Messages[chain][seqNum]; !ok {
					return Observation{}, fmt.Errorf("missing message with seqNr %d from chain %d", seqNum, chain)
				}
				obs.Messages[chain][seqNum] = cciptypes.Message{}
				obs.TokenData[chain][seqNum] = NewMessageTokenData()

				encodedObsSize = encodedObsSize -
					obs.MessageAndTokenDataEncodedSizes[chain][seqNum] + // Subtract the removed message and token size
					op.emptyEncodedSizes.MessageAndTokenData // Add empty message and token encoded size
				seqNum++
				if encodedObsSize <= op.maxEncodedSize {
					return obs, nil
				}
			}

			// Reaching here means that all messages in the report are truncated, truncate the last commit
			obs, encodedObsSize = op.truncateLastCommit(obs, chain, encodedObsSize)

			if len(obs.CommitReports[chain]) == 0 {
				// If the last commit report was truncated, truncate the chain
				obs = op.truncateChain(obs, chain)
			}

			if encodedObsSize <= op.maxEncodedSize {
				return obs, nil
			}
			chains = maps.Keys(obs.CommitReports)
		}
		// Truncated all chains. Return obs as is. (it has other data like contract discovery)
		if len(obs.CommitReports) == 0 {
			return obs, nil
		}
		// Encoding again after doing a full iteration on all chains and removing messages/commits.
		// That is because using encoded sizes is not 100% accurate and there are some missing bytes in the calculation.
		encodedObs, err = obs.Encode()
		if err != nil {
			return Observation{}, nil
		}
		encodedObsSize = len(encodedObs)
	}

	return obs, nil
}

// truncateLastCommit removes the last commit from the observation.
// errors if there are no commits to truncate.
func (op ObservationOptimizer) truncateLastCommit(
	obs Observation,
	chain cciptypes.ChainSelector,
	currentObsEncSize int,
) (Observation, int) {
	observation := obs
	newSize := currentObsEncSize
	commits := observation.CommitReports[chain]
	if len(commits) == 0 {
		return observation, currentObsEncSize
	}
	lastCommit := commits[len(commits)-1]
	// Remove the last commit from the list.
	commits = commits[:len(commits)-1]
	observation.CommitReports[chain] = commits
	// Remove from the encoded size.
	newSize = newSize - op.emptyEncodedSizes.CommitData - 4 // brackets, and commas
	for seqNum, msg := range observation.Messages[chain] {
		if lastCommit.SequenceNumberRange.Contains(seqNum) {
			// Remove the message from the observation.
			delete(observation.Messages[chain], seqNum)
			// Remove the token data from the observation.
			delete(observation.TokenData[chain], seqNum)
			//delete(observation.Hashes[chain], seqNum)
			// Remove the encoded size of the message and token data.
			newSize = newSize - op.emptyEncodedSizes.MessageAndTokenData - 2*op.emptyEncodedSizes.SeqNumMap -
				4 // for brackets and commas
			// Remove costly messages
			for i, costlyMessage := range observation.CostlyMessages {
				if costlyMessage == msg.Header.MessageID {
					observation.CostlyMessages = append(observation.CostlyMessages[:i], observation.CostlyMessages[i+1:]...)
				}
			}
			// Leaving Nonces untouched
		}
	}

	return observation, newSize
}

// truncateChain removes all data related to the given chain from the observation.
// returns true if the chain was found and truncated, false otherwise.
func (op ObservationOptimizer) truncateChain(
	obs Observation,
	chain cciptypes.ChainSelector,
) Observation {
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
