package costlymessages

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/costcalculator"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Observer observes messages that are too costly to execute.
type Observer interface {
	// Observe takes a set of messages and returns a slice of message IDs that are too costly to execute.
	// Also takes in a map from message ID to the message's timestamp, needed to calculate fee boosting.
	Observe(
		ctx context.Context,
		messages []cciptypes.Message,
		messageTimestamps map[cciptypes.Bytes32]time.Time,
	) ([]cciptypes.Bytes32, error)
}

// NewObserverWithDefaults creates a new Observer with default calculators.
// The default calculators are:
// - CCIPMessageFeeUSD18Calculator
// - CCIPMessageExecCostUSD18Calculator
func NewObserverWithDefaults(
	lggr logger.Logger,
	enabled bool,
	ccipReader readerpkg.CCIPReader,
	relativeBoostPerWaitHour float64,
	estimateProvider ccipocr3.EstimateProvider,
) Observer {
	return NewObserver(
		lggr,
		enabled,
		costcalculator.NewCCIPMessageFeeUSD18Calculator(
			lggr,
			ccipReader,
			relativeBoostPerWaitHour,
			time.Now,
		),
		costcalculator.NewCCIPMessageExecCostUSD18Calculator(
			lggr,
			ccipReader,
			estimateProvider,
		),
	)
}

// NewObserver allows to specific feeCalculator and execCostCalculator.
// Therefore, it's very convenient for testing.
func NewObserver(
	lggr logger.Logger,
	enabled bool,
	feeCalculator costcalculator.MessageFeeE18USDCalculator,
	execCostCalculator costcalculator.MessageExecCostUSD18Calculator,
) Observer {
	return &observer{
		lggr:               lggr,
		enabled:            enabled,
		feeCalculator:      feeCalculator,
		execCostCalculator: execCostCalculator,
	}
}

type observer struct {
	lggr               logger.Logger
	enabled            bool
	feeCalculator      costcalculator.MessageFeeE18USDCalculator
	execCostCalculator costcalculator.MessageExecCostUSD18Calculator
}

// Observe returns a slice of message IDs that are too costly to execute.
// It calculates the fee and execution cost of each message. The messages are considered too costly if the fee is less
// than the execution cost.
func (o *observer) Observe(
	ctx context.Context,
	messages []cciptypes.Message,
	messageTimestamps map[cciptypes.Bytes32]time.Time,
) ([]cciptypes.Bytes32, error) {
	if !o.enabled {
		o.lggr.Infof("Observer is disabled")
		return nil, nil
	}

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
			o.lggr.Warnw("Message is too costly to execute", "messageID",
				msg.Header.MessageID.String(), "fee", fee, "execCost", execCost, "seqNum", msg.Header.SequenceNumber)
			costlyMessages = append(costlyMessages, msg.Header.MessageID)
		}
	}

	return costlyMessages, nil
}

var _ Observer = &observer{}
