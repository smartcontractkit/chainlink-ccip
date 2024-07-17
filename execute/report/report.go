package report

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/temp"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// buildSingleChainReport converts the on-chain event data stored in cciptypes.ExecutePluginCommitDataWithMessages into
// the final on-chain report format.
//
// The hasher and encoding codec are provided as arguments to allow for chain-specific formats to be used.
//
// The messages argument indicates which messages should be included in the report. If messages is empty
// all messages will be included. This allows the caller to create smaller reports if needed. Executed messages
// are skipped automatically.
func buildSingleChainReport(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	tokenDataReader temp.TokenDataReader,
	encoder cciptypes.ExecutePluginCodec,
	report plugintypes.ExecutePluginCommitDataWithMessages,
	maxMessages int,
) (cciptypes.ExecutePluginReportSingleChain, int, uint64, error) {
	if maxMessages == 0 {
		maxMessages = len(report.Messages)
	}

	lggr.Debugw(
		"constructing merkle tree",
		"sourceChain", report.SourceChain,
		"expectedRoot", report.MerkleRoot.String(),
		"treeLeaves", len(report.Messages))

	tree, err := ConstructMerkleTree(ctx, hasher, report)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, 0, 0,
			fmt.Errorf("unable to construct merkle tree from messages for report (%s): %w", report.MerkleRoot.String(), err)
	}

	// Verify merkle root.
	hash := tree.Root()
	if !bytes.Equal(hash[:], report.MerkleRoot[:]) {
		actualStr := "0x" + hex.EncodeToString(hash[:])
		return cciptypes.ExecutePluginReportSingleChain{}, 0, 0,
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
				return cciptypes.ExecutePluginReportSingleChain{}, 0, 0,
					fmt.Errorf("unable to read token data for message %d: %w", msg.Header.SequenceNumber, err)
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
		return cciptypes.ExecutePluginReportSingleChain{}, 0, 0,
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
		return cciptypes.ExecutePluginReportSingleChain{}, 0, 0, fmt.Errorf("unable to encode report: %w", err)
	}

	return finalReport, len(encoded), 0, nil
}

// buildSingleChainReportMaxSize generates the largest report which fits into maxSizeBytes.
// See buildSingleChainReport for more details about how a report is built.
func buildSingleChainReportMaxSize(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	tokenDataReader temp.TokenDataReader,
	encoder cciptypes.ExecutePluginCodec,
	report plugintypes.ExecutePluginCommitDataWithMessages,
	maxSizeBytes int,
) (cciptypes.ExecutePluginReportSingleChain, int, plugintypes.ExecutePluginCommitDataWithMessages, error) {
	finalReport, encodedSize, _, err :=
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
		finalReport2, encodedSize2, _, err :=
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
			ErrEmptyReport
	}

	report = markNewMessagesExecuted(finalReport, report)
	return finalReport, encodedSize, report, nil
}
