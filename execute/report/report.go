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
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// buildSingleChainReportHelper converts the on-chain event data stored in cciptypes.ExecutePluginCommitData into the
// final on-chain report format. Messages in the report are selected based on the readyMessages argument. If
// readyMessages is empty all messages will be included. This allows the caller to create smaller reports if needed.
//
// Before calling this function all messages should have been checked and processed by the checkMessage function.
//
// The hasher and encoding codec are provided as arguments to allow for chain-specific formats to be used.
func buildSingleChainReportHelper(
	lggr logger.Logger,
	report exectypes.CommitData,
	readyMessages map[int]struct{},
) (ccipocr3.ExecutePluginReportSingleChain, error) {
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
		return ccipocr3.ExecutePluginReportSingleChain{}, nil
	}

	numMsg := len(report.Messages)
	if len(report.MessageTokenData) != numMsg {
		return ccipocr3.ExecutePluginReportSingleChain{},
			fmt.Errorf("token data length mismatch: got %d, expected %d", len(report.MessageTokenData), numMsg)
	}

	lggr.Infow(
		"constructing merkle tree",
		"sourceChain", report.SourceChain,
		"expectedRoot", report.MerkleRoot.String(),
		"treeLeaves", len(report.Messages))

	tree, err := ConstructMerkleTree(report, lggr)
	if err != nil {
		return ccipocr3.ExecutePluginReportSingleChain{},
			fmt.Errorf("unable to construct merkle tree from messages for report (%s): %w", report.MerkleRoot.String(), err)
	}

	// Verify merkle root.
	hash := tree.Root()
	if !bytes.Equal(hash[:], report.MerkleRoot[:]) {
		actualStr := "0x" + hex.EncodeToString(hash[:])
		return ccipocr3.ExecutePluginReportSingleChain{},
			fmt.Errorf("merkle root mismatch: expected %s, got %s", report.MerkleRoot.String(), actualStr)
	}

	lggr.Debugw("merkle root verified",
		"sourceChain", report.SourceChain,
		"commitRoot", report.MerkleRoot.String(),
		"computedRoot", ccipocr3.Bytes32(hash))

	// Iterate sequence range and executed messages to select messages to execute.
	numMsgs := len(report.Messages)
	var toExecute []int
	var offchainTokenData [][][]byte
	var msgInRoot []ccipocr3.Message
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
		return ccipocr3.ExecutePluginReportSingleChain{},
			fmt.Errorf("unable to prove messages for report %s: %w", report.MerkleRoot.String(), err)
	}

	var proofsCast []ccipocr3.Bytes32
	for _, p := range proof.Hashes {
		proofsCast = append(proofsCast, p)
	}

	lggr.Debugw("generated proofs", "sourceChain", report.SourceChain,
		"proofs", proofsCast, "proof", proof)

	finalReport := ccipocr3.ExecutePluginReportSingleChain{
		SourceChainSelector: report.SourceChain,
		Messages:            msgInRoot,
		OffchainTokenData:   offchainTokenData,
		Proofs:              proofsCast,
		ProofFlagBits:       ccipocr3.BigInt{Int: slicelib.BoolsToBitFlags(proof.SourceFlags)},
	}

	lggr.Debugw("final report", "sourceChain", report.SourceChain, "report", finalReport)

	return finalReport, nil
}

type messageStatus string

const (
	None                          messageStatus = ""
	Error                         messageStatus = "error"
	ReadyToExecute                messageStatus = "ready_to_execute"
	AlreadyExecuted               messageStatus = "already_executed"
	TokenDataNotReady             messageStatus = "token_data_not_ready" //nolint:gosec // this is not a password
	PseudoDeleted                 messageStatus = "message_pseudo_deleted"
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

// Check for the messages.
type Check func(lggr logger.Logger, msg ccipocr3.Message, idx int, report exectypes.CommitData) (messageStatus, error)

/*
// Template
func Template() Check {
	return func(lggr logger.Logger, msg ccipocr3.Message, idx int, report exectypes.CommitData) (messageStatus, error) {
		// check

		// else
		return None, nil
	}
}
*/

// CheckIfPseudoDeleted checks if the message has been removed, typically done to reduce observation size.
// This check should happen early because other checks are likely to fail if the message has been deleted.
func CheckIfPseudoDeleted() Check {
	return func(lggr logger.Logger, msg ccipocr3.Message, idx int, report exectypes.CommitData) (messageStatus, error) {
		if msg.IsEmpty() {
			lggr.Errorw("message pseudo deleted", "index", idx)
			return PseudoDeleted, nil
		}
		return None, nil
	}
}

// CheckAlreadyExecuted checks the report executed list to see if the message has been executed.
func CheckAlreadyExecuted() Check {
	return func(lggr logger.Logger, msg ccipocr3.Message, idx int, report exectypes.CommitData) (messageStatus, error) {
		if slices.Contains(report.ExecutedMessages, msg.Header.SequenceNumber) {
			lggr.Infow(
				"message already executed",
				"messageID", msg.Header.MessageID,
				"sourceChain", report.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"messageState", AlreadyExecuted)
			return AlreadyExecuted, nil
		}

		return None, nil
	}
}

// CheckTokenData rejects messages which are missing their token data (attestations, i.e. CCTP).
func CheckTokenData() Check {
	return func(lggr logger.Logger, msg ccipocr3.Message, idx int, report exectypes.CommitData) (messageStatus, error) {
		if idx >= len(report.MessageTokenData) {
			lggr.Errorw("token data index out of range", "index", idx, "messageTokensData", len(report.MessageTokenData))
			return Error, fmt.Errorf("token data index out of range")
		}

		messageTokenData := report.MessageTokenData[idx]
		if !messageTokenData.IsReady() {
			lggr.Infow(
				"unable to read token data - token data not ready",
				"messageID", msg.Header.MessageID,
				"sourceChain", report.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"error", messageTokenData.Error(),
				"messageState", TokenDataNotReady)
			return TokenDataNotReady, nil
		}

		lggr.Infow(
			"read token data",
			"messageID", msg.Header.MessageID,
			"sourceChain", report.SourceChain,
			"seqNum", msg.Header.SequenceNumber,
			"data", messageTokenData.ToByteSlice())

		return None, nil
	}
}

// CheckTooCostly compares the costly list for a given message.
func CheckTooCostly() Check {
	return func(lggr logger.Logger, msg ccipocr3.Message, idx int, report exectypes.CommitData) (messageStatus, error) {
		// 4. Check if the message is too costly to execute.
		if slices.Contains(report.CostlyMessages, msg.Header.MessageID) {
			lggr.Infow(
				"message too costly to execute",
				"messageID", msg.Header.MessageID,
				"sourceChain", report.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"messageState", TooCostly)
			return TooCostly, nil
		}

		return None, nil
	}
}

// CheckNonces ensures that messages are executed in the correct order by
// comparing the expected nonce to the message nonce. The check needs to be
// initialized using a list of sender nonces from the destination chain. In
// order to check multiple messages from the same sender, a copy of the initial
// list is maintained with incremented nonces after each message.
func CheckNonces(sendersNonce map[ccipocr3.ChainSelector]map[string]uint64) Check {
	// temporary map to store state between nonce checks for this round.
	expectedNonce := make(map[ccipocr3.ChainSelector]map[string]uint64)

	return func(lggr logger.Logger, msg ccipocr3.Message, idx int, report exectypes.CommitData) (messageStatus, error) {
		// Setting the Nonce to zero (or omitting it) indicates that the message
		// can be executed out of order. We allow this in the plugin by skipping
		// the nonce check.
		if msg.Header.Nonce == 0 {
			return None, nil
		}

		if _, ok := sendersNonce[report.SourceChain]; !ok {
			lggr.Errorw("Skipping message - nonces not available for chain",
				"messageID", msg.Header.MessageID,
				"sourceChain", report.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"messageState", MissingNoncesForChain)
			return MissingNoncesForChain, nil
		}

		chainNonces := sendersNonce[report.SourceChain]
		sender := typeconv.AddressBytesToString(msg.Sender[:], uint64(msg.Header.SourceChainSelector))
		if _, ok := chainNonces[sender]; !ok {
			lggr.Errorw("Skipping message - missing nonce",
				"messageID", msg.Header.MessageID,
				"sourceChain", report.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"messageState", MissingNonce)
			return MissingNonce, nil
		}

		if _, ok := expectedNonce[report.SourceChain]; !ok {
			// initialize expected nonce if needed.
			expectedNonce[report.SourceChain] = make(map[string]uint64)
		}
		if _, ok := expectedNonce[report.SourceChain][sender]; !ok {
			expectedNonce[report.SourceChain][sender] = chainNonces[sender] + 1
		}

		// Check expected nonce is valid for sequenced messages.
		if msg.Header.Nonce != expectedNonce[report.SourceChain][sender] {
			lggr.Warnw("Skipping message - invalid nonce",
				"messageID", msg.Header.MessageID,
				"sourceChain", report.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"have", msg.Header.Nonce,
				"want", expectedNonce[report.SourceChain][sender],
				"messageState", InvalidNonce)
			return InvalidNonce, nil
		}
		expectedNonce[report.SourceChain][sender] = expectedNonce[report.SourceChain][sender] + 1

		return None, nil
	}
}

// checkMessages to get a set of which are ready to execute.
func (b *execReportBuilder) checkMessages(ctx context.Context, report exectypes.CommitData) (map[int]struct{}, error) {
	readyMessages := make(map[int]struct{})
	for i := 0; i < len(report.Messages); i++ {
		updatedReport, status, err := b.checkMessage(ctx, i, report)
		if err != nil {
			return nil,
				fmt.Errorf("unable to check message: %w", err)
		}
		report = updatedReport
		if status == ReadyToExecute {
			readyMessages[i] = struct{}{}
		}
	}
	return readyMessages, nil
}

// checkMessage for execution readiness.
func (b *execReportBuilder) checkMessage(
	_ context.Context, idx int, execReport exectypes.CommitData,
) (exectypes.CommitData, messageStatus, error) {
	result := execReport

	// OutOfRangeCheck
	if idx >= len(execReport.Messages) {
		b.lggr.Errorw("message index out of range", "index", idx, "numMessages", len(execReport.Messages))
		return execReport, Error, fmt.Errorf("message index out of range")
	}

	msg := execReport.Messages[idx]

	for _, check := range b.checks {
		status, err := check(b.lggr, msg, idx, execReport)
		if err != nil {
			return execReport, Error, err
		}
		if status != None {
			return execReport, status, nil
		}
	}

	// TODO: Check txm for too many failures?

	return result, ReadyToExecute, nil
}

// verifyReport is a final step to ensure the encoded message meets our exec criteria.
func (b *execReportBuilder) verifyReport(
	ctx context.Context,
	execReport ccipocr3.ExecutePluginReportSingleChain,
) (bool, validationMetadata, error) {
	// Compute the size of the encoded report.
	// Note: ExecutePluginReport is a strict array of data, so wrapping the final report
	//       does not add any additional overhead to the size being computed here.
	encoded, err := b.encoder.Encode(
		ctx,
		ccipocr3.ExecutePluginReport{
			ChainReports: []ccipocr3.ExecutePluginReportSingleChain{execReport},
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
// 1. exec report for the builder
// 2. updated commit report after marking new messages from the exec report as executed
//
//nolint:gocyclo // todo
func (b *execReportBuilder) buildSingleChainReport(
	ctx context.Context,
	report exectypes.CommitData,
) (ccipocr3.ExecutePluginReportSingleChain, exectypes.CommitData, error) {
	finalize := func(
		execReport ccipocr3.ExecutePluginReportSingleChain,
		commitReport exectypes.CommitData,
		meta validationMetadata,
	) (ccipocr3.ExecutePluginReportSingleChain, exectypes.CommitData, error) {
		b.accumulated = b.accumulated.accumulate(meta)
		commitReport = markNewMessagesExecuted(execReport, commitReport)
		return execReport, commitReport, nil
	}

	// Check which messages are ready to execute, and update report with any additional metadata needed for execution.
	readyMessages, err := b.checkMessages(ctx, report)
	if err != nil {
		return ccipocr3.ExecutePluginReportSingleChain{},
			exectypes.CommitData{},
			fmt.Errorf("unable to check message: %w", err)
	}
	if len(readyMessages) == 0 {
		return ccipocr3.ExecutePluginReportSingleChain{}, report, ErrEmptyReport
	}

	// Unless there is a message limit, attempt to build a report for executing all ready messages.
	if b.maxMessages == 0 {
		finalReport, err :=
			buildSingleChainReportHelper(b.lggr, report, readyMessages)
		if err != nil {
			return ccipocr3.ExecutePluginReportSingleChain{},
				exectypes.CommitData{},
				fmt.Errorf("unable to build a single chain report (max): %w", err)
		}

		validReport, meta, err := b.verifyReport(ctx, finalReport)
		if err != nil {
			return ccipocr3.ExecutePluginReportSingleChain{},
				exectypes.CommitData{},
				fmt.Errorf("unable to verify report: %w", err)
		} else if validReport {
			return finalize(finalReport, report, meta)
		}
	}

	var finalReport ccipocr3.ExecutePluginReportSingleChain
	var meta validationMetadata
	msgs := make(map[int]struct{})
	for i := range report.Messages {
		if _, ok := readyMessages[i]; !ok {
			continue
		}

		msgs[i] = struct{}{}

		finalReport2, err := buildSingleChainReportHelper(b.lggr, report, msgs)
		if err != nil {
			return ccipocr3.ExecutePluginReportSingleChain{},
				exectypes.CommitData{},
				fmt.Errorf("unable to build a single chain report (messages %d): %w", len(msgs), err)
		}

		validReport, meta2, err := b.verifyReport(ctx, finalReport2)
		if err != nil {
			return ccipocr3.ExecutePluginReportSingleChain{},
				exectypes.CommitData{},
				fmt.Errorf("unable to verify report: %w", err)
		} else if validReport {
			finalReport = finalReport2
			meta = meta2

			// Stop searching if we reach the maximum number of messages.
			if b.maxMessages > 0 && uint64(len(msgs)) >= b.maxMessages {
				break
			}
		} else {
			// this message didn't work, continue to the next one
			delete(msgs, i)
		}
	}

	if len(msgs) == 0 {
		return ccipocr3.ExecutePluginReportSingleChain{}, report, ErrEmptyReport
	}

	return finalize(finalReport, report, meta)
}
