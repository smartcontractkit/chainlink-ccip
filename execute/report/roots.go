package report

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

// ConstructMerkleTree creates the merkle tree object from the messages in the report.
func ConstructMerkleTree(
	report exectypes.CommitData,
	lggr logger.Logger,
) (*merklemulti.Tree[[32]byte], error) {
	// Ensure we have the expected number of messages
	numMsgs := int(report.SequenceNumberRange.End() - report.SequenceNumberRange.Start() + 1)
	if numMsgs != len(report.Messages) {
		return nil, fmt.Errorf(
			"malformed report %s, unexpected number of messages: expected %d, got %d",
			report.MerkleRoot.String(), numMsgs, len(report.Messages))
	}

	treeLeaves := make([][32]byte, 0)
	for _, msg := range report.Messages {
		if !report.SequenceNumberRange.Contains(msg.Header.SequenceNumber) {
			return nil, fmt.Errorf(
				"malformed report, message %s sequence number %d outside of report range %s",
				report.MerkleRoot.String(), msg.Header.SequenceNumber, report.SequenceNumberRange)
		}
		if report.SourceChain != msg.Header.SourceChainSelector {
			return nil, fmt.Errorf("malformed report, message %s for unexpected source chain: expected %d, got %d",
				report.MerkleRoot.String(), report.SourceChain, msg.Header.SourceChainSelector)
		}
		if msg.Header.MsgHash == [32]byte{} {
			return nil, fmt.Errorf("malformed report, message %s has empty hash", msg)
		}
		leaf := msg.Header.MsgHash
		lggr.Debugw("Hashed message, adding to tree leaves",
			"hash", leaf,
			"msg", msg,
			"merkleRoot", report.MerkleRoot,
			"sourceChain", report.SourceChain)
		treeLeaves = append(treeLeaves, leaf)
	}

	// TODO: Do not hard code the hash function, it should be derived from the message hasher.
	return merklemulti.NewTree(hashutil.NewKeccak(), treeLeaves)
}
