package exectypes

import (
	"context"
	"fmt"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// CostlyMessageObserver observes messages that are too costly to execute.
type CostlyMessageObserver interface {
	// Observe takes a set of messages and returns a slice of message IDs that are too costly to execute.
	// Also takes in a map from message ID to the message's timestamp, needed to calculate fee boosting.
	Observe(
		ctx context.Context,
		obs MessageObservations,
		messageTimestamps map[cciptypes.Bytes32]time.Time,
	) ([]cciptypes.Bytes32, error)
}

func NewCostlyMessageObserver() CostlyMessageObserver {
	return &CcipCostlyMessageObserver{
		// TODO: Implement fee and exec cost calculators
		feeCalculator:      &NoOpMessageFeeE18USDCalculator{},
		execCostCalculator: &NoOpMessageExecCostE18USDCalculator{},
	}
}

type CcipCostlyMessageObserver struct {
	feeCalculator      MessageFeeE18USDCalculator
	execCostCalculator MessageExecCostE18USDCalculator
}

// Observe returns a slice of message IDs that are too costly to execute.
// It calculates the fee and execution cost of each message. The messages are considered too costly if the fee is less
// than the execution cost.
func (o *CcipCostlyMessageObserver) Observe(
	ctx context.Context,
	obs MessageObservations,
	messageTimestamps map[cciptypes.Bytes32]time.Time,
) ([]cciptypes.Bytes32, error) {
	messages := make([]cciptypes.Message, 0)
	for _, seqNumToMsg := range obs {
		for _, msg := range seqNumToMsg {
			messages = append(messages, msg)
		}
	}

	messageFees, err := o.feeCalculator.MessageFeeE18USD(ctx, messages, messageTimestamps)
	if err != nil {
		return nil, fmt.Errorf("unable to calculate message fees: %w", err)
	}

	execCosts, err := o.execCostCalculator.MessageExecCostE18USD(ctx, messages)
	if err != nil {
		return nil, fmt.Errorf("unable to calculate message execution costs: %w", err)
	}

	costlyMessages := make([]cciptypes.Bytes32, 0)
	for _, msg := range messages {
		fee, ok := messageFees[msg.Header.MessageID]
		if !ok {
			return nil, fmt.Errorf("missing fee for message %s", msg.Header.MessageID)
		}
		execCost, ok := execCosts[msg.Header.MessageID]
		if !ok {
			return nil, fmt.Errorf("missing exec cost for message %s", msg.Header.MessageID)
		}
		if fee.Cmp(execCost.Int) < 0 {
			costlyMessages = append(costlyMessages, msg.Header.MessageID)
		}
	}

	return costlyMessages, nil
}

var _ CostlyMessageObserver = &CcipCostlyMessageObserver{}

// MessageFeeE18USDCalculator Calculates the fees (paid at source) of a set of messages in e-18 USDs.
type MessageFeeE18USDCalculator interface {
	// MessageFeeE18USD Returns a map from message ID to the message's fee in e-18 USDs. For example, if the message's
	// fee is 12USD, this function return this message's fee as 12 * 1e18. You can think of this function returning the
	// fee not in USD, but in a small denomination of USD, analogous to returning the cost in wei instead of ETH
	// (1 wei = 1e-18 ETH).
	MessageFeeE18USD(
		ctx context.Context,
		messages []cciptypes.Message,
		messageTimestamps map[cciptypes.Bytes32]time.Time,
	) (map[cciptypes.Bytes32]cciptypes.BigInt, error)
}

// NoOpMessageFeeE18USDCalculator returns a fee of 0 for all messages.
type NoOpMessageFeeE18USDCalculator struct{}

// MessageFeeE18USD returns a fee of 0 for all messages.
func (n *NoOpMessageFeeE18USDCalculator) MessageFeeE18USD(
	_ context.Context,
	messages []cciptypes.Message,
	_ map[cciptypes.Bytes32]time.Time,
) (map[cciptypes.Bytes32]cciptypes.BigInt, error) {
	messageFees := make(map[cciptypes.Bytes32]cciptypes.BigInt)
	for _, msg := range messages {
		messageFees[msg.Header.MessageID] = cciptypes.NewBigIntFromInt64(0)
	}
	return messageFees, nil
}

var _ MessageFeeE18USDCalculator = &NoOpMessageFeeE18USDCalculator{}

// MessageExecCostE18USDCalculator Calculates the execution cost of a set of messages in 1e-18 USDs.
type MessageExecCostE18USDCalculator interface {
	// MessageExecCostE18USD Returns a map from message ID to the message's estimated execution cost in 1e-18 USDs.
	// For example, if the cost of executing a message is 12USD, this function return this message's cost as 12 * 1e18.
	// You can think of this function returning the cost not in USD, but in a small denomination of USD, analogous to
	// returning the cost in wei instead of ETH (1 wei = 1e-18 ETH).
	MessageExecCostE18USD(context.Context, []cciptypes.Message) (map[cciptypes.Bytes32]cciptypes.BigInt, error)
}

// NoOpMessageExecCostE18USDCalculator returns a cost of 0 for all messages.
type NoOpMessageExecCostE18USDCalculator struct{}

// MessageExecCostE18USD returns a cost of 0 for all messages.
func (n *NoOpMessageExecCostE18USDCalculator) MessageExecCostE18USD(
	_ context.Context,
	messages []cciptypes.Message,
) (map[cciptypes.Bytes32]cciptypes.BigInt, error) {
	messageExecCosts := make(map[cciptypes.Bytes32]cciptypes.BigInt)
	for _, msg := range messages {
		messageExecCosts[msg.Header.MessageID] = cciptypes.NewBigIntFromInt64(0)
	}
	return messageExecCosts, nil
}

var _ MessageExecCostE18USDCalculator = &NoOpMessageExecCostE18USDCalculator{}
