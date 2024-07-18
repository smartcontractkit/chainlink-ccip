package report

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/types"
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
	tokenDataReader types.TokenDataReader,
	report plugintypes.ExecutePluginCommitDataWithMessages,
	messages map[int]struct{},
) (cciptypes.ExecutePluginReportSingleChain, error) {
	if len(messages) == 0 {
		if messages == nil {
			messages = make(map[int]struct{})
		}
		for i := 0; i < len(report.Messages); i++ {
			messages[i] = struct{}{}
		}
	}

	lggr.Debugw(
		"constructing merkle tree",
		"sourceChain", report.SourceChain,
		"expectedRoot", report.MerkleRoot.String(),
		"treeLeaves", len(report.Messages))

	tree, err := ConstructMerkleTree(ctx, hasher, report)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			fmt.Errorf("unable to construct merkle tree from messages for report (%s): %w", report.MerkleRoot.String(), err)
	}

	// Verify merkle root.
	hash := tree.Root()
	if !bytes.Equal(hash[:], report.MerkleRoot[:]) {
		actualStr := "0x" + hex.EncodeToString(hash[:])
		return cciptypes.ExecutePluginReportSingleChain{},
			fmt.Errorf("merkle root mismatch: expected %s, got %s", report.MerkleRoot.String(), actualStr)
	}

	// Iterate sequence range and executed messages to select messages to execute.
	numMsgs := len(report.Messages)
	var toExecute []int
	var offchainTokenData [][][]byte
	var msgInRoot []cciptypes.Message
	executedIdx := 0
	for i, msg := range report.Messages {
		seqNum := report.SequenceNumberRange.Start() + cciptypes.SeqNum(i)
		// Skip messages which are already executed
		if executedIdx < len(report.ExecutedMessages) && report.ExecutedMessages[executedIdx] == seqNum {
			executedIdx++
		} else if _, ok := messages[i]; ok {
			tokenData, err := tokenDataReader.ReadTokenData(context.Background(), report.SourceChain, msg.Header.SequenceNumber)
			if err != nil {
				// TODO: skip message instead of failing the whole thing.
				//       that might mean moving the token data reading out of the loop.
				lggr.Infow(
					"unable to read token data",
					"sourceChain", report.SourceChain,
					"seqNum", msg.Header.SequenceNumber,
					"error", err)
				return cciptypes.ExecutePluginReportSingleChain{},
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
		return cciptypes.ExecutePluginReportSingleChain{},
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

	return finalReport, nil
}

func (b *execReportBuilder) verifyReport(
	ctx context.Context, execReport cciptypes.ExecutePluginReportSingleChain,
) (bool, validationMetadata, error) {
	// Compute the size of the encoded report.
	// Note: ExecutePluginReport is a strict array of data, so wrapping the final report
	//       does not add any additional overhead to the size being computed here.
	encoded, err := b.encoder.Encode(
		ctx,
		cciptypes.ExecutePluginReport{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{execReport},
		},
	)
	if err != nil {
		b.lggr.Errorw("unable to encode report", "err", err, "report", execReport)
		return false, validationMetadata{}, fmt.Errorf("unable to encode report: %w", err)
	}

	maxSizeBytes := int(b.maxReportSizeBytes - b.accumulated.encodedSizeBytes)
	if len(encoded) > maxSizeBytes {
		b.lggr.Debugw("invalid report, report size exceeds limit", "size", len(encoded), "maxSize", maxSizeBytes)
		return false, validationMetadata{}, nil
	}

	return true, validationMetadata{
		encodedSizeBytes: uint64(len(encoded)),
	}, nil
}

// buildSingleChainReportMaxSize generates the largest report which fits into maxSizeBytes.
// See buildSingleChainReport for more details about how a report is built.
func (b *execReportBuilder) buildSingleChainReportMaxSize(
	ctx context.Context,
	report plugintypes.ExecutePluginCommitDataWithMessages,
) (cciptypes.ExecutePluginReportSingleChain, plugintypes.ExecutePluginCommitDataWithMessages, error) {
	finalize := func(
		execReport cciptypes.ExecutePluginReportSingleChain,
		commitReport plugintypes.ExecutePluginCommitDataWithMessages,
		meta validationMetadata,
	) (cciptypes.ExecutePluginReportSingleChain, plugintypes.ExecutePluginCommitDataWithMessages, error) {
		b.accumulated.encodedSizeBytes += meta.encodedSizeBytes
		commitReport = markNewMessagesExecuted(execReport, commitReport)
		return execReport, commitReport, nil
	}
	// Attempt to include all messages in the report.
	finalReport, err :=
		buildSingleChainReport(b.ctx, b.lggr, b.hasher, b.tokenDataReader, report, nil)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			plugintypes.ExecutePluginCommitDataWithMessages{},
			fmt.Errorf("unable to build a single chain report (max): %w", err)
	}

	validReport, meta, err := b.verifyReport(ctx, finalReport)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			plugintypes.ExecutePluginCommitDataWithMessages{},
			fmt.Errorf("unable to verify report: %w", err)
	} else if validReport {
		return finalize(finalReport, report, meta)
	}

	var searchErr error
	idx := sort.Search(len(report.Messages), func(mid int) bool {
		if searchErr != nil {
			return false
		}
		msgs := make(map[int]struct{})
		for i := 0; i < mid; i++ {
			msgs[i] = struct{}{}
		}
		finalReport2, err :=
			buildSingleChainReport(b.ctx, b.lggr, b.hasher, b.tokenDataReader, report, msgs)
		if searchErr != nil {
			searchErr = fmt.Errorf("unable to build a single chain report (messages %d): %w", mid, err)
			return false
		}

		validReport, meta2, err := b.verifyReport(ctx, finalReport2)
		if err != nil {
			searchErr = fmt.Errorf("unable to verify report: %w", err)
			return false
		} else if validReport {
			finalReport = finalReport2
			meta = meta2
		}

		return !validReport // full
	})
	if searchErr != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, plugintypes.ExecutePluginCommitDataWithMessages{}, searchErr
	}

	// No messages fit into the report.
	if idx <= 0 {
		return cciptypes.ExecutePluginReportSingleChain{},
			plugintypes.ExecutePluginCommitDataWithMessages{},
			ErrEmptyReport
	}

	return finalize(finalReport, report, meta)
}
