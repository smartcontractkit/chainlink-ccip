package exectypes

import (
	"context"
	"fmt"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

// CostlyMessageObserver observes messages that are too costly to execute.
type CostlyMessageObserver interface {
	// Observe takes a set of messages and returns a slice of message IDs that are too costly to execute.
	// Also takes in a map from message ID to the message's timestamp, needed to calculate fee boosting.
	Observe(
		ctx context.Context,
		messages []cciptypes.Message,
		messageTimestamps map[cciptypes.Bytes32]time.Time,
	) ([]cciptypes.Bytes32, error)
}

func NewCostlyMessageObserver() CostlyMessageObserver {
	return &CCIPCostlyMessageObserver{
		// TODO: Implement fee and exec cost calculators
		feeCalculator:      &ZeroMessageFeeUSD18Calculator{},
		execCostCalculator: &ZeroMessageExecCostUSD18Calculator{},
	}
}

type CCIPCostlyMessageObserver struct {
	feeCalculator      MessageFeeE18USDCalculator
	execCostCalculator MessageExecCostUSD18Calculator
}

// Observe returns a slice of message IDs that are too costly to execute.
// It calculates the fee and execution cost of each message. The messages are considered too costly if the fee is less
// than the execution cost.
func (o *CCIPCostlyMessageObserver) Observe(
	ctx context.Context,
	messages []cciptypes.Message,
	messageTimestamps map[cciptypes.Bytes32]time.Time,
) ([]cciptypes.Bytes32, error) {
	messageFees, err := o.feeCalculator.MessageFeeUSD18(ctx, messages, messageTimestamps)
	if err != nil {
		return nil, fmt.Errorf("unable to calculate message fees: %w", err)
	}

	execCosts, err := o.execCostCalculator.MessageExecCostUSD18(ctx, messages)
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
		if fee.Cmp(execCost) < 0 {
			costlyMessages = append(costlyMessages, msg.Header.MessageID)
		}
	}

	return costlyMessages, nil
}

var _ CostlyMessageObserver = &CCIPCostlyMessageObserver{}

// MessageFeeE18USDCalculator Calculates the fees (paid at source) of a set of messages in USD18s.
type MessageFeeE18USDCalculator interface {
	// MessageFeeUSD18 Returns a map from message ID to the message's fee in USD18s.
	MessageFeeUSD18(
		ctx context.Context,
		messages []cciptypes.Message,
		messageTimestamps map[cciptypes.Bytes32]time.Time,
	) (map[cciptypes.Bytes32]plugintypes.USD18, error)
}

// ZeroMessageFeeUSD18Calculator returns a fee of 0 for all messages.
type ZeroMessageFeeUSD18Calculator struct{}

// MessageFeeUSD18 returns a fee of 0 for all messages.
func (n *ZeroMessageFeeUSD18Calculator) MessageFeeUSD18(
	_ context.Context,
	messages []cciptypes.Message,
	_ map[cciptypes.Bytes32]time.Time,
) (map[cciptypes.Bytes32]plugintypes.USD18, error) {
	messageFees := make(map[cciptypes.Bytes32]plugintypes.USD18)
	for _, msg := range messages {
		messageFees[msg.Header.MessageID] = plugintypes.NewUSD18(0)
	}
	return messageFees, nil
}

var _ MessageFeeE18USDCalculator = &ZeroMessageFeeUSD18Calculator{}

// MessageExecCostUSD18Calculator Calculates the execution cost of a set of messages in USD18s.
type MessageExecCostUSD18Calculator interface {
	// MessageExecCostUSD18 Returns a map from message ID to the message's estimated execution cost in USD18s.
	MessageExecCostUSD18(context.Context, []cciptypes.Message) (map[cciptypes.Bytes32]plugintypes.USD18, error)
}

// ZeroMessageExecCostUSD18Calculator returns a cost of 0 for all messages.
type ZeroMessageExecCostUSD18Calculator struct{}

// MessageExecCostUSD18 returns a cost of 0 for all messages.
func (n *ZeroMessageExecCostUSD18Calculator) MessageExecCostUSD18(
	_ context.Context,
	messages []cciptypes.Message,
) (map[cciptypes.Bytes32]plugintypes.USD18, error) {
	messageExecCosts := make(map[cciptypes.Bytes32]plugintypes.USD18)
	for _, msg := range messages {
		messageExecCosts[msg.Header.MessageID] = plugintypes.NewUSD18(0)
	}
	return messageExecCosts, nil
}

var _ MessageExecCostUSD18Calculator = &ZeroMessageExecCostUSD18Calculator{}

// StaticMessageFeeUSD18Calculator returns a static fee for all messages.
type StaticMessageFeeUSD18Calculator struct {
	fees map[cciptypes.Bytes32]plugintypes.USD18
}

func NewStaticMessageFeeUSD18Calculator(
	fees map[cciptypes.Bytes32]plugintypes.USD18,
) *StaticMessageFeeUSD18Calculator {
	return &StaticMessageFeeUSD18Calculator{fees: fees}
}

// MessageFeeUSD18 returns a fee of 0 for all messages.
func (n *StaticMessageFeeUSD18Calculator) MessageFeeUSD18(
	_ context.Context,
	messages []cciptypes.Message,
	_ map[cciptypes.Bytes32]time.Time,
) (map[cciptypes.Bytes32]plugintypes.USD18, error) {
	messageFees := make(map[cciptypes.Bytes32]plugintypes.USD18)

	for _, msg := range messages {
		fee, ok := n.fees[msg.Header.MessageID]
		if !ok {
			return nil, fmt.Errorf("missing fee for message %s", msg.Header.MessageID)
		}
		messageFees[msg.Header.MessageID] = fee
	}

	return messageFees, nil
}

var _ MessageFeeE18USDCalculator = &StaticMessageFeeUSD18Calculator{}

// StaticMessageExecCostUSD18Calculator returns a static cost for all messages.
type StaticMessageExecCostUSD18Calculator struct {
	costs map[cciptypes.Bytes32]plugintypes.USD18
}

func NewStaticMessageExecCostUSD18Calculator(
	costs map[cciptypes.Bytes32]plugintypes.USD18,
) *StaticMessageExecCostUSD18Calculator {
	return &StaticMessageExecCostUSD18Calculator{costs: costs}
}

// MessageExecCostUSD18 returns a cost of 0 for all messages.
func (n *StaticMessageExecCostUSD18Calculator) MessageExecCostUSD18(
	_ context.Context,
	messages []cciptypes.Message,
) (map[cciptypes.Bytes32]plugintypes.USD18, error) {
	messageExecCosts := make(map[cciptypes.Bytes32]plugintypes.USD18)

	for _, msg := range messages {
		cost, ok := n.costs[msg.Header.MessageID]
		if !ok {
			return nil, fmt.Errorf("missing exec cost for message %s", msg.Header.MessageID)
		}
		messageExecCosts[msg.Header.MessageID] = cost
	}

	return messageExecCosts, nil
}

var _ MessageExecCostUSD18Calculator = &StaticMessageExecCostUSD18Calculator{}
