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
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// buildSingleChainReportHelper converts the on-chain event data stored in ccipocr3.ExecutePluginCommitData into the
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

	lggr.Debugw("in-progress report built",
		"sourceChain", report.SourceChain,
		"report", finalReport)

	return finalReport, nil
}

type messageStatus string

const (
	None                          messageStatus = ""
	Error                         messageStatus = "error"
	ReadyToExecute                messageStatus = "ready_to_execute"
	AlreadyExecuted               messageStatus = "already_executed"
	AlreadyInflight               messageStatus = "already_inflight"
	TokenDataNotReady             messageStatus = "token_data_not_ready" //nolint:gosec // this is not a password
	PseudoDeleted                 messageStatus = "message_pseudo_deleted"
	TokenDataFetchError           messageStatus = "token_data_fetch_error"
	InsufficientRemainingBatchGas messageStatus = "insufficient_remaining_batch_gas"
	MissingNoncesForChain         messageStatus = "missing_nonces_for_chain"
	MissingNonce                  messageStatus = "missing_nonce"
	InvalidNonce                  messageStatus = "invalid_nonce"
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
		if msg.IsPseudoDeleted() {
			lggr.Infow("message pseudo deleted", "index", idx, "messageID", msg.Header.MessageID)
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

// CheckNonces ensures that messages are executed in the correct order by
// comparing the expected nonce to the message nonce. The check needs to be
// initialized using a list of sender nonces from the destination chain. In
// order to check multiple messages from the same sender, a copy of the initial
// list is maintained with incremented nonces after each message.
//
// NOTE: this check is currently only applied prior to the report generation process.
// Its function is _not_ to check that nonces are skipped, but to make sure that we
// skip messages that have an invalid nonce - i.e the nonce of the message is _not_
// the expected nonce, which starts as the onchain nonce + 1.
// As it stands now, this is a well-formedness check of the commit report w/ data
// and also a filterer of messages w/ invalid nonces from the getgo, relative to onchain nonces.
// for example: if the onchain nonce is 10, and the report has the messages:
// * 8, 9, 10, 11, 12
// then the messages 8, 9, 10 will be skipped because their nonces are <= the onchain nonce.
//
// TODO: CCIP-5374. There is some duplication w/ verifyReportNonceContinuity below.
func CheckNonces(sendersNonce map[ccipocr3.ChainSelector]map[string]uint64, addressCodec ccipocr3.AddressCodec) Check {
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
		sender, err := addressCodec.AddressBytesToString(msg.Sender[:], msg.Header.SourceChainSelector)
		if err != nil {
			return Error, fmt.Errorf("unable to convert sender address to string: %w, sender address: %v", err, msg.Sender[:])
		}

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

type IsInflight func(src ccipocr3.ChainSelector, msgID ccipocr3.Bytes32) bool

func CheckIfInflight(inflight IsInflight) Check {
	return func(lggr logger.Logger, msg ccipocr3.Message, idx int, report exectypes.CommitData) (messageStatus, error) {
		if inflight(report.SourceChain, msg.Header.MessageID) {
			lggr.Infow(
				"message already in flight",
				"messageID", msg.Header.MessageID,
				"sourceChain", report.SourceChain,
				"seqNum", msg.Header.SequenceNumber,
				"messageState", "inflight")
			return AlreadyInflight, nil
		}
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

// verifyReportNonceContinuity returns an error if the provided report contains a single skipped
// nonce from a single sender, rendering the report invalid.
//
// we cannot have "skipped" nonces since these messages will be skipped onchain.
// they will also end up in the inflight cache and "snoozed" until they
// expire from the inflight cache TTL.
//
// TODO: CCIP-5374. There is some duplication with CheckNonces above.
func verifyReportNonceContinuity(
	addressCodec ccipocr3.AddressCodec,
	execReport ccipocr3.ExecutePluginReportSingleChain,
) error {
	// note that the messages in the report _must_ be ordered correctly in nonce order,
	// otherwise the report is invalid. this is because we loop through the messages in order
	// onchain and try to execute each in turn.
	// for that reason, we can use this map approach by checking the latest nonce for each sender,
	// instead of accumulating the nonces in an array and searching for gaps.
	// so this check catches incorrectly ordered messages as well as skipped nonces.
	latestNoncePerSender := make(map[string]uint64)
	for _, msg := range execReport.Messages {
		if msg.Header.Nonce == 0 {
			// nonce is zero means an out of order execution message, it doesn't affect nonce ordering.
			// we can safely skip this check for this message.
			continue
		}

		sender, err := addressCodec.AddressBytesToString(msg.Sender, msg.Header.SourceChainSelector)
		if err != nil {
			return fmt.Errorf("unable to convert sender address to string: %w, sender address: %v", err, msg.Sender)
		}

		latestNonce, ok := latestNoncePerSender[sender]
		if !ok {
			// first time we ever see this sender, populate the nonce.
			latestNoncePerSender[sender] = msg.Header.Nonce
			continue
		} else if latestNonce == 0 {
			// this is a bug in this code, should never happen since we check if msg.Header.Nonce
			// is zero before this check.
			return fmt.Errorf("assumption violation: bug in the code, latestNonce is zero when it should never be")
		} else if msg.Header.Nonce != latestNonce+1 {
			// we've seen this sender before and the msg.Header.Nonce is not the expected value,
			// it should be latestNonce + 1.
			return fmt.Errorf(
				"skipped nonce detected for sender %s (source chain %d): nonce in report %d != last nonce seen %d",
				sender, msg.Header.SourceChainSelector, msg.Header.Nonce, latestNonce)
		}

		// otherwise, the nonce is as expected, continue to the next message.
		// update the map to reflect the latest seen nonce for this sender.
		latestNoncePerSender[sender] = msg.Header.Nonce
	}

	return nil
}

// verifyReport is a final step to ensure the encoded message meets our exec criteria.
func (b *execReportBuilder) verifyReport(
	ctx context.Context,
	execReport ccipocr3.ExecutePluginReportSingleChain,
) (bool, validationMetadata, error) {
	err := verifyReportNonceContinuity(b.addressCodec, execReport)
	if err != nil {
		b.lggr.Infow("invalid report, skipped nonce detected",
			"err", err,
			"sourceChain", execReport.SourceChainSelector)
		return false, validationMetadata{}, nil
	}

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
	commitData exectypes.CommitData,
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
	readyMessages, err := b.checkMessages(ctx, commitData)
	if err != nil {
		return ccipocr3.ExecutePluginReportSingleChain{},
			exectypes.CommitData{},
			fmt.Errorf("unable to check message: %w", err)
	}
	if len(readyMessages) == 0 {
		return ccipocr3.ExecutePluginReportSingleChain{}, commitData, ErrEmptyReport
	}

	// Unless there is a message limit, attempt to build a report for executing all ready messages.
	// It is possible that the report produced here is invalid for some reason, such as
	// report size or gas usage.
	// In that case, we will execute the loop below to iteratively build a report
	// with fewer messages until we find a valid report.
	if b.maxMessages == 0 {
		finalReport, err :=
			buildSingleChainReportHelper(b.lggr, commitData, readyMessages)
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
			b.lggr.Infow("messages added to report",
				"messageIDs", slicelib.Map(finalReport.Messages, func(m ccipocr3.Message) ccipocr3.Bytes32 {
					return m.Header.MessageID
				}),
				"seqNums", slicelib.Map(finalReport.Messages, func(m ccipocr3.Message) ccipocr3.SeqNum {
					return m.Header.SequenceNumber
				}),
				"senders", slicelib.Map(finalReport.Messages, func(m ccipocr3.Message) ccipocr3.UnknownAddress {
					return m.Sender
				}),
				"nonces", slicelib.Map(finalReport.Messages, func(m ccipocr3.Message) uint64 {
					return m.Header.Nonce
				}),
				"sourceChain", commitData.SourceChain,
				"reportSizeBytes", meta.encodedSizeBytes,
				"reportGas", meta.gas)
			return finalize(finalReport, commitData, meta)
		}
	}

	var finalReport ccipocr3.ExecutePluginReportSingleChain
	var meta validationMetadata
	msgs := make(map[int]struct{})
	for i := range commitData.Messages {
		if _, ok := readyMessages[i]; !ok {
			continue
		}

		msgs[i] = struct{}{}

		finalReport2, err := buildSingleChainReportHelper(b.lggr, commitData, msgs)
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

			b.lggr.Infow(
				"message added to report",
				"sourceChain", commitData.Messages[i].Header.SourceChainSelector,
				"seqNum", commitData.Messages[i].Header.SequenceNumber,
				"messageID", commitData.Messages[i].Header.MessageID,
				"nonce", commitData.Messages[i].Header.Nonce,
				"sender", commitData.Messages[i].Sender.String(),
				"reportSizeBytes", meta.encodedSizeBytes,
				"reportGas", meta.gas,
			)
			// Stop searching if we reach the maximum number of messages.
			if b.maxMessages > 0 && uint64(len(msgs)) >= b.maxMessages {
				b.lggr.Infow(
					"reached report builder's max messages, breaking",
					"maxMessages", b.maxMessages,
					"numMessages", len(msgs),
				)
				break
			}
		} else {
			// this message didn't work, continue to the next one
			b.lggr.Debugw("message did not fit in report, deleting from in-progress report built",
				"sourceChain", commitData.Messages[i].Header.SourceChainSelector,
				"messageID", commitData.Messages[i].Header.MessageID,
				"seqNum", commitData.Messages[i].Header.SequenceNumber,
			)
			delete(msgs, i)
		}
	}

	if len(msgs) == 0 {
		return ccipocr3.ExecutePluginReportSingleChain{}, commitData, ErrEmptyReport
	}

	return finalize(finalReport, commitData, meta)
}
