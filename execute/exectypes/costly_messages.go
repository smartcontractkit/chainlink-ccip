package exectypes

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

func NewCostlyMessageObserverWithZeroExec(
	lggr logger.Logger,
	enabled bool,
	ccipReader readerpkg.CCIPReader,
	relativeBoostPerWaitHour float64,
) CostlyMessageObserver {
	return NewCostlyMessageObserver(
		lggr,
		enabled,
		&CCIPMessageFeeUSD18Calculator{
			lggr:                     lggr,
			ccipReader:               ccipReader,
			relativeBoostPerWaitHour: relativeBoostPerWaitHour,
			now:                      time.Now,
		},
		&ZeroMessageExecCostUSD18Calculator{},
	)
}

func NewCostlyMessageObserver(
	lggr logger.Logger,
	enabled bool,
	feeCalculator MessageFeeE18USDCalculator,
	execCostCalculator MessageExecCostUSD18Calculator,
) CostlyMessageObserver {
	return &CCIPCostlyMessageObserver{
		lggr:               lggr,
		enabled:            enabled,
		feeCalculator:      feeCalculator,
		execCostCalculator: execCostCalculator,
	}
}

type CCIPCostlyMessageObserver struct {
	lggr               logger.Logger
	enabled            bool
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
	if !o.enabled {
		o.lggr.Infof("CostlyMessageObserver is disabled")
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

func (n *StaticMessageExecCostUSD18Calculator) UpdateCosts(msgID cciptypes.Bytes32, fee plugintypes.USD18) {
	n.costs[msgID] = fee
}

var _ MessageExecCostUSD18Calculator = &StaticMessageExecCostUSD18Calculator{}

// CCIPMessageFeeUSD18Calculator calculates the fees (paid at source) of a set of messages in USD18s.
type CCIPMessageFeeUSD18Calculator struct {
	lggr logger.Logger

	ccipReader readerpkg.CCIPReader

	// RelativeBoostPerWaitHour indicates how much to increase (artificially) the fee paid on the source chain per hour
	// of wait time, such that eventually the fee paid is greater than the execution cost, and we’ll execute it.
	// For example: if set to 0.5, that means the fee paid is increased by 50% every hour the message has been waiting.
	relativeBoostPerWaitHour float64

	now func() time.Time
}

func NewCCIPMessageFeeUSD18Calculator(
	lggr logger.Logger,
	ccipReader readerpkg.CCIPReader,
	relativeBoostPerWaitHour float64,
	now func() time.Time,
) *CCIPMessageFeeUSD18Calculator {
	return &CCIPMessageFeeUSD18Calculator{
		lggr:                     lggr,
		ccipReader:               ccipReader,
		relativeBoostPerWaitHour: relativeBoostPerWaitHour,
		now:                      now,
	}
}

var _ MessageFeeE18USDCalculator = &CCIPMessageFeeUSD18Calculator{}

// MessageFeeUSD18 Returns a map from message ID to the message's fee in USD18s.
func (c *CCIPMessageFeeUSD18Calculator) MessageFeeUSD18(
	ctx context.Context,
	messages []cciptypes.Message,
	messageTimeStamps map[cciptypes.Bytes32]time.Time,
) (map[cciptypes.Bytes32]plugintypes.USD18, error) {
	linkPriceUSD, err := c.ccipReader.LinkPriceUSD(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get LINK price in USD: %w", err)
	}

	messageFees := make(map[cciptypes.Bytes32]plugintypes.USD18)
	for _, msg := range messages {
		feeUSD18 := new(big.Int).Mul(linkPriceUSD.Int, msg.FeeValueJuels.Int)
		timestamp, ok := messageTimeStamps[msg.Header.MessageID]
		if !ok {
			// If a timestamp is missing we can't do fee boosting, but we still record the fee. In the worst case, the
			// message will not be executed (as it will be considered too costly).
			c.lggr.Warnw("missing timestamp for message", "messageID", msg.Header.MessageID)
		} else {
			feeUSD18 = waitBoostedFee(c.now().Sub(timestamp), feeUSD18, c.relativeBoostPerWaitHour)
		}

		messageFees[msg.Header.MessageID] = feeUSD18
	}

	return messageFees, nil
}

// waitBoostedFee boosts the given fee according to the time passed since the msg was sent.
// RelativeBoostPerWaitHour is used to normalize the time diff,
// it makes our loss taking "smooth" and gives us time to react without a hard deadline.
// At the same time, messages that are slightly underpaid will start going through after waiting for a little bit.
//
// wait_boosted_fee(m) = (1 + (now - m.send_time).hours * RELATIVE_BOOST_PER_WAIT_HOUR) * fee(m)
func waitBoostedFee(waitTime time.Duration, fee *big.Int, relativeBoostPerWaitHour float64) *big.Int {
	k := 1.0 + waitTime.Hours()*relativeBoostPerWaitHour

	boostedFee := big.NewFloat(0).Mul(big.NewFloat(k), new(big.Float).SetInt(fee))
	res, _ := boostedFee.Int(nil)

	return res
}
