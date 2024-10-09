package exectypes

import (
	"context"
	"fmt"
	"math/big"
	"time"

	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
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

func NewCostlyMessageObserver(
	lggr logger.Logger,
	ccipReader readerpkg.CCIPReader,
	RelativeBoostPerWaitHour float64,
) CostlyMessageObserver {
	return &CcipCostlyMessageObserver{
		lggr: lggr,
		feeCalculator: &CcipMessageFeeE18USDCalculator{
			lggr:                     lggr,
			ccipReader:               ccipReader,
			RelativeBoostPerWaitHour: RelativeBoostPerWaitHour,
		},
		// TODO: Implement exec cost calculator
		execCostCalculator: &NoOpMessageExecCostE18USDCalculator{},
	}
}

type CcipCostlyMessageObserver struct {
	lggr               logger.Logger
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
			o.lggr.Infow("message is too costly to execute", "messageID", msg.Header.MessageID)
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

// CcipMessageFeeE18USDCalculator calculates the fees (paid at source) of a set of messages in e-18 USDs.
type CcipMessageFeeE18USDCalculator struct {
	lggr logger.Logger

	ccipReader readerpkg.CCIPReader

	// RelativeBoostPerWaitHour indicates how much to increase (artificially) the fee paid on the source chain per hour
	// of wait time, such that eventually the fee paid is greater than the execution cost, and we’ll execute it.
	// For example: if set to 0.5, that means the fee paid is increased by 50% every hour the message has been waiting.
	RelativeBoostPerWaitHour float64 `json:"relativeBoostPerWaitHour"`
}

var _ MessageFeeE18USDCalculator = &CcipMessageFeeE18USDCalculator{}

// MessageFeeE18USD Returns a map from message ID to the message's fee in e-18 USDs. For example, if the message's
// fee is 12USD, this function return this message's fee as 12 * 1e18. You can think of this function returning the
// fee not in USD, but in a small denomination of USD, analogous to returning the cost in wei instead of ETH
// (1 wei = 1e-18 ETH).
func (c *CcipMessageFeeE18USDCalculator) MessageFeeE18USD(
	ctx context.Context,
	messages []cciptypes.Message,
	messageTimeStamps map[cciptypes.Bytes32]time.Time,
) (map[cciptypes.Bytes32]cciptypes.BigInt, error) {
	linkPriceUSD, err := c.ccipReader.LinkPriceUSD(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get LINK price in USD: %w", err)
	}

	messageFees := make(map[cciptypes.Bytes32]cciptypes.BigInt)
	for _, msg := range messages {
		feeE18USD := big.NewInt(0).Mul(linkPriceUSD.Int, msg.FeeValueJuels.Int)
		timestamp, ok := messageTimeStamps[msg.Header.MessageID]
		if !ok {
			// If a timestamp is missing we can't do fee boosting, but we still record the fee. In the worst case, the
			// message will not be executed (as it will be considered too costly).
			c.lggr.Warnw("missing timestamp for message", "messageID", msg.Header.MessageID)
		} else {
			feeE18USD = waitBoostedFee(time.Since(timestamp), feeE18USD, c.RelativeBoostPerWaitHour)
		}

		messageFees[msg.Header.MessageID] = cciptypes.NewBigInt(feeE18USD)
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
