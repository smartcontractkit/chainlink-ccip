package report

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"slices"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// buildSingleChainReportHelper converts the on-chain event data stored in cciptypes.ExecutePluginCommitData into the
// final on-chain report format. Messages in the report are selected based on the readyMessages argument. If
// readyMessages is empty all messages will be included. This allows the caller to create smaller reports if needed.
//
// Before calling this function all messages should have been checked and processed by the checkMessage function.
//
// The hasher and encoding codec are provided as arguments to allow for chain-specific formats to be used.
func buildSingleChainReportHelper(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	report exectypes.CommitData,
	readyMessages map[int]struct{},
) (cciptypes.ExecutePluginReportSingleChain, error) {
	if len(readyMessages) == 0 {
		if readyMessages == nil {
			readyMessages = make(map[int]struct{})
		}
		for i := 0; i < len(report.Messages); i++ {
			readyMessages[i] = struct{}{}
		}
	}

	if len(readyMessages) == 0 {
		lggr.Infow("no messages ready for execution", "sourceChain", report.SourceChain)
		return cciptypes.ExecutePluginReportSingleChain{}, nil
	}

	numMsg := len(report.Messages)
	if len(report.MessageTokenData) != numMsg {
		return cciptypes.ExecutePluginReportSingleChain{},
			fmt.Errorf("token data length mismatch: got %d, expected %d", len(report.MessageTokenData), numMsg)
	}

	lggr.Infow(
		"constructing merkle tree",
		"sourceChain", report.SourceChain,
		"expectedRoot", report.MerkleRoot.String(),
		"treeLeaves", len(report.Messages))

	tree, err := ConstructMerkleTree(report, lggr)
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
	for i, msg := range report.Messages {
		if _, ok := readyMessages[i]; ok {
			offchainTokenData = append(offchainTokenData, report.MessageTokenData[i].ToByteSlice())
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
		"proofs", proofsCast, "proof", proof)

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
	Unknown                       messageStatus = "unknown"
	ReadyToExecute                messageStatus = "ready_to_execute"
	AlreadyExecuted               messageStatus = "already_executed"
	TokenDataNotReady             messageStatus = "token_data_not_ready" //nolint:gosec // this is not a password
	TokenDataFetchError           messageStatus = "token_data_fetch_error"
	InsufficientRemainingBatchGas messageStatus = "insufficient_remaining_batch_gas"
	MissingNoncesForChain         messageStatus = "missing_nonces_for_chain"
	MissingNonce                  messageStatus = "missing_nonce"
	InvalidNonce                  messageStatus = "invalid_nonce"
	TooCostly                     messageStatus = "tooCostly"
	/*
		SenderAlreadySkipped                 messageStatus = "sender_already_skipped"
		MessageMaxGasCalcError               messageStatus = "message_max_gas_calc_error"
		InsufficientRemainingBatchDataLength messageStatus = "insufficient_remaining_batch_data_length"
		AggregateTokenValueComputeError      messageStatus = "aggregate_token_value_compute_error"
		AggregateTokenLimitExceeded          messageStatus = "aggregate_token_limit_exceeded"
		TokenNotInDestTokenPrices            messageStatus = "token_not_in_dest_token_prices"
		TokenNotInSrcTokenPrices             messageStatus = "token_not_in_src_token_prices"
		InsufficientRemainingFee             messageStatus = "insufficient_remaining_fee"
		AddedToBatch                         messageStatus = "added_to_batch"
	*/
)

func (b *execReportBuilder) checkMessage(
	_ context.Context, idx int, execReport exectypes.CommitData,
) (exectypes.CommitData, messageStatus, error) {
	result := execReport

	if idx >= len(execReport.Messages) {
		b.lggr.Errorw("message index out of range", "index", idx, "numMessages", len(execReport.Messages))
		return execReport, Unknown, fmt.Errorf("message index out of range")
	}

	msg := execReport.Messages[idx]

	// 1. Check if the message has already been executed.
	if slices.Contains(execReport.ExecutedMessages, msg.Header.SequenceNumber) {
		b.lggr.Infow(
			"message already executed",
			"messageID", msg.Header.MessageID,
			"sourceChain", execReport.SourceChain,
			"seqNum", msg.Header.SequenceNumber,
			"messageState", AlreadyExecuted)
		return execReport, AlreadyExecuted, nil
	}

	// 2. Check if all tokens are properly initialized with token data
	if idx >= len(execReport.MessageTokenData) {
		b.lggr.Errorw("token data index out of range", "index", idx, "messageTokensData", len(execReport.MessageTokenData))
		return execReport, Unknown, fmt.Errorf("token data index out of range")
	}

	messageTokenData := execReport.MessageTokenData[idx]
	if !messageTokenData.IsReady() {
		b.lggr.Infow(
			"unable to read token data - token data not ready",
			"messageID", msg.Header.MessageID,
			"sourceChain", execReport.SourceChain,
			"seqNum", msg.Header.SequenceNumber,
			"error", messageTokenData.Error(),
			"messageState", TokenDataNotReady)
		return execReport, TokenDataNotReady, nil
	}
	result.MessageTokenData[idx] = messageTokenData
	b.lggr.Infow(
		"read token data",
		"messageID", msg.Header.MessageID,
		"sourceChain", execReport.SourceChain,
		"seqNum", msg.Header.SequenceNumber,
		"data", messageTokenData.ToByteSlice())

	// 3. Check if the message has a valid nonce.
	status := b.checkMessageNonce(msg, execReport)
	if status != "" {
		return execReport, status, nil
	}

	// 4. Check if the message is too costly to execute.
	if slices.Contains(execReport.CostlyMessages, msg.Header.MessageID) {
		b.lggr.Infow(
			"message too costly to execute",
			"messageID", msg.Header.MessageID,
			"sourceChain", execReport.SourceChain,
			"seqNum", msg.Header.SequenceNumber,
			"messageState", TooCostly)
		return execReport, TooCostly, nil
	}

	return result, ReadyToExecute, nil
}

func (b *execReportBuilder) checkMessageNonce(
	msg cciptypes.Message,
	execReport exectypes.CommitData,
) messageStatus {
	if msg.Header.Nonce != 0 {
		// Sequenced messages have non-zero nonces.

		if _, ok := b.sendersNonce[execReport.SourceChain]; !ok {
			b.lggr.Errorw("Skipping message - nonces not available for chain",
				"messageID", msg.Header.MessageID,
				"sourceChain", execReport.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"messageState", MissingNoncesForChain)
			return MissingNoncesForChain
		}

		chainNonces := b.sendersNonce[execReport.SourceChain]
		sender := typeconv.AddressBytesToString(msg.Sender[:], uint64(b.destChainSelector))
		if _, ok := chainNonces[sender]; !ok {
			b.lggr.Errorw("Skipping message - missing nonce",
				"messageID", msg.Header.MessageID,
				"sourceChain", execReport.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"messageState", MissingNonce)
			return MissingNonce
		}

		if b.expectedNonce == nil {
			// initialize expected nonce if needed.
			b.expectedNonce = make(map[cciptypes.ChainSelector]map[string]uint64)
		}
		if _, ok := b.expectedNonce[execReport.SourceChain]; !ok {
			// initialize expected nonce if needed.
			b.expectedNonce[execReport.SourceChain] = make(map[string]uint64)
		}
		if _, ok := b.expectedNonce[execReport.SourceChain][sender]; !ok {
			b.expectedNonce[execReport.SourceChain][sender] = chainNonces[sender] + 1
		}

		// Check expected nonce is valid for sequenced messages.
		if msg.Header.Nonce != b.expectedNonce[execReport.SourceChain][sender] {
			b.lggr.Warnw("Skipping message - invalid nonce",
				"messageID", msg.Header.MessageID,
				"sourceChain", execReport.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"have", msg.Header.Nonce,
				"want", b.expectedNonce[execReport.SourceChain][sender],
				"messageState", InvalidNonce)
			return InvalidNonce
		}
		b.expectedNonce[execReport.SourceChain][sender] = b.expectedNonce[execReport.SourceChain][sender] + 1
	}

	return ""
}

func (b *execReportBuilder) verifyReport(
	ctx context.Context,
	execReport cciptypes.ExecutePluginReportSingleChain,
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
		b.lggr.Infow("invalid report, report size exceeds limit", "size", len(encoded), "maxSize", maxSizeBytes)
		return false, validationMetadata{}, nil
	}

	// Add in accumulated gas
	if b.estimateProvider == nil {
		return false, validationMetadata{}, fmt.Errorf("gas estimator must be initialized")
	}
	gasSum := uint64(0)
	for _, msg := range execReport.Messages {
		gasSum += b.estimateProvider.CalculateMessageMaxGas(msg)
	}
	merkleTreeGas := b.estimateProvider.CalculateMerkleTreeGas(len(execReport.Messages))
	totalGas := gasSum + merkleTreeGas

	maxGas := b.maxGas - b.accumulated.gas
	if totalGas > maxGas {
		b.lggr.Infow("invalid report, report estimated gas usage exceeds limit", "gas", totalGas, "maxGas", maxGas)
		return false, validationMetadata{}, nil
	}

	return true, validationMetadata{
		encodedSizeBytes: uint64(len(encoded)),
		gas:              totalGas,
	}, nil
}

// buildSingleChainReport generates the largest report which fits into maxSizeBytes.
// See buildSingleChainReport for more details about how a report is built.
// returns
// 1. the exec report that's added to builder
// 2. the updated commit report after marking new messages from the exec report as executed
//
//nolint:gocyclo // todo
func (b *execReportBuilder) buildSingleChainReport(
	ctx context.Context,
	report exectypes.CommitData,
) (cciptypes.ExecutePluginReportSingleChain, exectypes.CommitData, error) {
	finalize := func(
		execReport cciptypes.ExecutePluginReportSingleChain,
		commitReport exectypes.CommitData,
		meta validationMetadata,
	) (cciptypes.ExecutePluginReportSingleChain, exectypes.CommitData, error) {
		b.accumulated = b.accumulated.accumulate(meta)
		commitReport = markNewMessagesExecuted(execReport, commitReport)
		return execReport, commitReport, nil
	}

	// Check which messages are ready to execute, and update report with any additional metadata needed for execution.
	readyMessages := make(map[int]struct{})
	for i := 0; i < len(report.Messages); i++ {
		updatedReport, status, err := b.checkMessage(ctx, i, report)
		if err != nil {
			return cciptypes.ExecutePluginReportSingleChain{},
				exectypes.CommitData{},
				fmt.Errorf("unable to check message: %w", err)
		}
		report = updatedReport
		if status == ReadyToExecute {
			readyMessages[i] = struct{}{}
		}
	}

	if len(readyMessages) == 0 {
		return cciptypes.ExecutePluginReportSingleChain{}, report, ErrEmptyReport
	}

	// Attempt to include all messages in the report.
	finalReport, err :=
		buildSingleChainReportHelper(ctx, b.lggr, b.hasher, report, readyMessages)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			exectypes.CommitData{},
			fmt.Errorf("unable to build a single chain report (max): %w", err)
	}

	validReport, meta, err := b.verifyReport(ctx, finalReport)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			exectypes.CommitData{},
			fmt.Errorf("unable to verify report: %w", err)
	} else if validReport {
		return finalize(finalReport, report, meta)
	}

	finalReport = cciptypes.ExecutePluginReportSingleChain{}
	msgs := make(map[int]struct{})
	for i := range report.Messages {
		if _, ok := readyMessages[i]; !ok {
			continue
		}

		msgs[i] = struct{}{}

		finalReport2, err := buildSingleChainReportHelper(ctx, b.lggr, b.hasher, report, msgs)
		if err != nil {
			return cciptypes.ExecutePluginReportSingleChain{},
				exectypes.CommitData{},
				fmt.Errorf("unable to build a single chain report (messages %d): %w", len(msgs), err)
		}

		validReport, meta2, err := b.verifyReport(ctx, finalReport2)
		if err != nil {
			return cciptypes.ExecutePluginReportSingleChain{},
				exectypes.CommitData{},
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
