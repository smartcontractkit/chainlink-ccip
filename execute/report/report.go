package report

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"slices"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// buildSingleChainReportHelper converts the on-chain event data stored in
// cciptypes.ExecutePluginCommitDataWithMessages into the final on-chain report format.
//
// The hasher and encoding codec are provided as arguments to allow for chain-specific formats to be used.
//
// The messages argument indicates which messages should be included in the report. If messages is empty
// all messages will be included. This allows the caller to create smaller reports if needed. Executed messages
// are skipped automatically.
func buildSingleChainReportHelper(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	tokenDataReader types.TokenDataReader,
	report plugintypes.ExecutePluginCommitData,
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

	lggr.Infow(
		"constructing merkle tree",
		"sourceChain", report.SourceChain,
		"expectedRoot", report.MerkleRoot.String(),
		"treeLeaves", len(report.Messages))

	tree, err := ConstructMerkleTree(ctx, hasher, report, lggr)
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

	lggr.Debugw("merkle root verified",
		"sourceChain", report.SourceChain,
		"commitRoot", report.MerkleRoot.String(),
		"computedRoot", cciptypes.Bytes32(hash))

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
			var tokenData [][]byte
			var err error
			if tokenDataReader != nil {
				tokenData, err = tokenDataReader.ReadTokenData(context.Background(), report.SourceChain, msg.Header.SequenceNumber)
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
			}
			offchainTokenData = append(offchainTokenData, tokenData)
			toExecute = append(toExecute, i)
			msgInRoot = append(msgInRoot, msg)
		}
	}

	lggr.Infow(
		"selected messages from commit report for execution, generating merkle proofs",
		"sourceChain", report.SourceChain,
		"commitRoot", report.MerkleRoot.String(),
		"numMessages", numMsgs,
		"toExecute", toExecute)
	proof, err := tree.Prove(toExecute)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			fmt.Errorf("unable to prove messages for report %s: %w", report.MerkleRoot.String(), err)
	}

	var proofsCast []cciptypes.Bytes32
	for _, p := range proof.Hashes {
		proofsCast = append(proofsCast, p)
	}

	lggr.Debugw("generated proofs", "sourceChain", report.SourceChain,
		"proofsCast", proofsCast, "proof", proof)

	finalReport := cciptypes.ExecutePluginReportSingleChain{
		SourceChainSelector: report.SourceChain,
		Messages:            msgInRoot,
		OffchainTokenData:   offchainTokenData,
		Proofs:              proofsCast,
		ProofFlagBits:       cciptypes.BigInt{Int: slicelib.BoolsToBitFlags(proof.SourceFlags)},
	}

	lggr.Debugw("final report", "sourceChain", report.SourceChain, "report", finalReport)

	return finalReport, nil
}

type messageStatus string

const (
	ReadyToExecute  messageStatus = "ready_to_execute"
	AlreadyExecuted messageStatus = "already_executed"
	/*
		SenderAlreadySkipped                 messageStatus = "sender_already_skipped"
		MessageMaxGasCalcError               messageStatus = "message_max_gas_calc_error"
		InsufficientRemainingBatchDataLength messageStatus = "insufficient_remaining_batch_data_length"
		InsufficientRemainingBatchGas        messageStatus = "insufficient_remaining_batch_gas"
		MissingNonce                         messageStatus = "missing_nonce"
		InvalidNonce                         messageStatus = "invalid_nonce"
		AggregateTokenValueComputeError      messageStatus = "aggregate_token_value_compute_error"
		AggregateTokenLimitExceeded          messageStatus = "aggregate_token_limit_exceeded"
		TokenDataNotReady                    messageStatus = "token_data_not_ready"
		TokenDataFetchError                  messageStatus = "token_data_fetch_error"
		TokenNotInDestTokenPrices            messageStatus = "token_not_in_dest_token_prices"
		TokenNotInSrcTokenPrices             messageStatus = "token_not_in_src_token_prices"
		InsufficientRemainingFee             messageStatus = "insufficient_remaining_fee"
		AddedToBatch                         messageStatus = "added_to_batch"
	*/
)

func (b *execReportBuilder) checkMessage(
	_ context.Context, idx int, execReport plugintypes.ExecutePluginCommitData,
	// TODO: get rid of the nolint when the error is used
) (messageStatus, error) { // nolint this will use the error eventually
	if slices.Contains(execReport.ExecutedMessages, execReport.Messages[idx].Header.SequenceNumber) {
		return AlreadyExecuted, nil
	}

	return ReadyToExecute, nil
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

// buildSingleChainReport generates the largest report which fits into maxSizeBytes.
// See buildSingleChainReport for more details about how a report is built.
func (b *execReportBuilder) buildSingleChainReport(
	ctx context.Context,
	report plugintypes.ExecutePluginCommitData,
) (cciptypes.ExecutePluginReportSingleChain, plugintypes.ExecutePluginCommitData, error) {
	finalize := func(
		execReport cciptypes.ExecutePluginReportSingleChain,
		commitReport plugintypes.ExecutePluginCommitData,
		meta validationMetadata,
	) (cciptypes.ExecutePluginReportSingleChain, plugintypes.ExecutePluginCommitData, error) {
		b.accumulated.encodedSizeBytes += meta.encodedSizeBytes
		commitReport = markNewMessagesExecuted(execReport, commitReport)
		return execReport, commitReport, nil
	}
	// Attempt to include all messages in the report.
	finalReport, err :=
		buildSingleChainReportHelper(b.ctx, b.lggr, b.hasher, b.tokenDataReader, report, nil)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			plugintypes.ExecutePluginCommitData{},
			fmt.Errorf("unable to build a single chain report (max): %w", err)
	}

	validReport, meta, err := b.verifyReport(ctx, finalReport)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			plugintypes.ExecutePluginCommitData{},
			fmt.Errorf("unable to verify report: %w", err)
	} else if validReport {
		return finalize(finalReport, report, meta)
	}

	finalReport = cciptypes.ExecutePluginReportSingleChain{}
	msgs := make(map[int]struct{})
	for i := range report.Messages {
		status, err := b.checkMessage(ctx, i, report)
		if err != nil {
			return cciptypes.ExecutePluginReportSingleChain{},
				plugintypes.ExecutePluginCommitData{},
				fmt.Errorf("unable to check message: %w", err)
		}
		if status != ReadyToExecute {
			continue
		}

		msgs[i] = struct{}{}

		finalReport2, err :=
			buildSingleChainReportHelper(b.ctx, b.lggr, b.hasher, b.tokenDataReader, report, msgs)
		if err != nil {
			return cciptypes.ExecutePluginReportSingleChain{},
				plugintypes.ExecutePluginCommitData{},
				fmt.Errorf("unable to build a single chain report (messages %d): %w", len(msgs), err)
		}

		validReport, meta2, err := b.verifyReport(ctx, finalReport2)
		if err != nil {
			return cciptypes.ExecutePluginReportSingleChain{},
				plugintypes.ExecutePluginCommitData{},
				fmt.Errorf("unable to verify report: %w", err)
		} else if validReport {
			finalReport = finalReport2
			meta = meta2
		} else {
			// this message didn't work, continue to the next one
			delete(msgs, i)
		}
	}

	if len(msgs) == 0 {
		return cciptypes.ExecutePluginReportSingleChain{}, report, ErrEmptyReport
	}

	return finalize(finalReport, report, meta)
}
