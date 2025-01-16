package costlymessages

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/mathslib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	EVMWordBytes              = 32
	MessageFixedBytesPerToken = 32 * ((2 * 3) + 3)
	ConstantMessagePartBytes  = 32 * 14 // A message consists of 14 abi encoded fields 32B each (after encoding)
	daMultiplierBase          = 10_000  // DA multiplier is in multiples of 0.0001, i.e. 1/daMultiplierBase
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
	estimateProvider cciptypes.EstimateProvider,
) Observer {
	return NewObserver(
		lggr,
		enabled,
		NewCCIPMessageFeeUSD18Calculator(
			lggr,
			ccipReader,
			relativeBoostPerWaitHour,
			time.Now,
		),
		&CCIPMessageExecCostUSD18Calculator{
			lggr:             lggr,
			ccipReader:       ccipReader,
			estimateProvider: estimateProvider,
		},
	)
}

// NewObserver allows to specific feeCalculator and execCostCalculator.
// Therefore, it's very convenient for testing.
func NewObserver(
	lggr logger.Logger,
	enabled bool,
	feeCalculator MessageFeeE18USDCalculator,
	execCostCalculator MessageExecCostUSD18Calculator,
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
	feeCalculator      MessageFeeE18USDCalculator
	execCostCalculator MessageExecCostUSD18Calculator
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

	if len(messages) == 0 {
		return make([]cciptypes.Bytes32, 0), nil
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
		if err := validateNonNegative(fee); err != nil {
			return nil, fmt.Errorf("invalid fee for message %s: %w", msg.Header.MessageID, err)
		}
		execCost, ok := execCosts[msg.Header.MessageID]
		if !ok {
			return nil, fmt.Errorf("missing exec cost for message %s", msg.Header.MessageID)
		}
		if err := validateNonNegative(execCost); err != nil {
			return nil, fmt.Errorf("invalid fee for message %s: %w", msg.Header.MessageID, err)
		}
		if fee.Cmp(execCost) < 0 {
			o.lggr.Warnw("Message is too costly to execute", "messageID",
				msg.Header.MessageID.String(), "fee", fee, "execCost", execCost, "seqNum", msg.Header.SequenceNumber)
			costlyMessages = append(costlyMessages, msg.Header.MessageID)
		}
	}

	return costlyMessages, nil
}

func validateNonNegative(fee *big.Int) error {
	if fee == nil {
		return fmt.Errorf("fee is nil")
	}
	if fee.Cmp(big.NewInt(0)) <= 0 {
		return fmt.Errorf("fee is non positive %d", fee)
	}
	return nil
}

var _ Observer = &observer{}

// MessageFeeE18USDCalculator Calculates the fees (paid at source) of a set of messages in USD18s.
type MessageFeeE18USDCalculator interface {
	// MessageFeeUSD18 Returns a map from message ID to the message's fee in USD18s.
	MessageFeeUSD18(
		ctx context.Context,
		messages []cciptypes.Message,
		messageTimestamps map[cciptypes.Bytes32]time.Time,
	) (map[cciptypes.Bytes32]plugintypes.USD18, error)
}

// ConstMessageFeeUSD18Calculator returns a fee of 0 for all messages.
type ConstMessageFeeUSD18Calculator struct {
	fee *big.Int
}

func NewConstMessageFeeUSD18Calculator(fee *big.Int) *ConstMessageFeeUSD18Calculator {
	return &ConstMessageFeeUSD18Calculator{fee: fee}
}

// MessageFeeUSD18 returns a fee of 0 for all messages.
func (n *ConstMessageFeeUSD18Calculator) MessageFeeUSD18(
	_ context.Context,
	messages []cciptypes.Message,
	_ map[cciptypes.Bytes32]time.Time,
) (map[cciptypes.Bytes32]plugintypes.USD18, error) {
	messageFees := make(map[cciptypes.Bytes32]plugintypes.USD18)
	for _, msg := range messages {
		messageFees[msg.Header.MessageID] = n.fee
	}
	return messageFees, nil
}

var _ MessageFeeE18USDCalculator = &ConstMessageFeeUSD18Calculator{}

// MessageExecCostUSD18Calculator Calculates the execution cost of a set of messages in USD18s.
type MessageExecCostUSD18Calculator interface {
	// MessageExecCostUSD18 Returns a map from message ID to the message's estimated execution cost in USD18s.
	MessageExecCostUSD18(context.Context, []cciptypes.Message) (map[cciptypes.Bytes32]plugintypes.USD18, error)
}

// ConstMessageExecCostUSD18Calculator returns a cost of 0 for all messages.
type ConstMessageExecCostUSD18Calculator struct {
	cost *big.Int
}

func NewConstMessageExecCostUSD18Calculator(cost *big.Int) *ConstMessageExecCostUSD18Calculator {
	return &ConstMessageExecCostUSD18Calculator{
		cost: cost,
	}
}

// MessageExecCostUSD18 returns a cost of 0 for all messages.
func (n *ConstMessageExecCostUSD18Calculator) MessageExecCostUSD18(
	_ context.Context,
	messages []cciptypes.Message,
) (map[cciptypes.Bytes32]plugintypes.USD18, error) {
	messageExecCosts := make(map[cciptypes.Bytes32]plugintypes.USD18)
	for _, msg := range messages {
		messageExecCosts[msg.Header.MessageID] = n.cost
	}
	return messageExecCosts, nil
}

var _ MessageExecCostUSD18Calculator = &ConstMessageExecCostUSD18Calculator{}

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

// UpdateCosts updates the costs of the single message. Not thread-safe, meant to be used only for tests.
func (n *StaticMessageExecCostUSD18Calculator) UpdateCosts(msgID cciptypes.Bytes32, cost plugintypes.USD18) {
	n.costs[msgID] = cost
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
		feeUSD18 := mathslib.CalculateUsdPerUnitGas(msg.FeeValueJuels.Int, linkPriceUSD.Int)
		timestamp, ok := messageTimeStamps[msg.Header.MessageID]
		if !ok {
			// If a timestamp is missing we can't do fee boosting, but we still record the fee. In the worst case, the
			// message will not be executed (as it will be considered too costly).
			c.lggr.Warnw("missing timestamp for message", "messageID", msg.Header.MessageID)
		} else {
			// TODO: What's the blockchain timestamp? Should we use now().UTC instead?
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
func waitBoostedFee(
	waitTime time.Duration,
	fee *big.Int,
	relativeBoostPerWaitHour float64) *big.Int {
	k := 1.0 + waitTime.Hours()*relativeBoostPerWaitHour

	boostedFee := big.NewFloat(0).Mul(big.NewFloat(k), new(big.Float).SetInt(fee))
	res, _ := boostedFee.Int(nil)

	return res
}

type CCIPMessageExecCostUSD18Calculator struct {
	lggr             logger.Logger
	ccipReader       readerpkg.CCIPReader
	estimateProvider cciptypes.EstimateProvider
}

// MessageExecCostUSD18 returns a map from message ID to the message's estimated execution cost in USD18s.
func (c *CCIPMessageExecCostUSD18Calculator) MessageExecCostUSD18(
	ctx context.Context,
	messages []cciptypes.Message,
) (map[cciptypes.Bytes32]plugintypes.USD18, error) {
	if len(messages) == 0 {
		return nil, fmt.Errorf("no messages provided")
	}

	messageExecCosts := make(map[cciptypes.Bytes32]plugintypes.USD18)
	feeComponents, err := c.ccipReader.GetDestChainFeeComponents(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get fee components: %w", err)
	}
	if feeComponents.ExecutionFee == nil {
		return nil, fmt.Errorf("missing execution fee")
	}
	if feeComponents.DataAvailabilityFee == nil {
		return nil, fmt.Errorf("missing data availability fee")
	}

	executionFee, daFee, err := c.getFeesUSD18(ctx, feeComponents, messages[0].Header.DestChainSelector)
	if err != nil {
		return nil, fmt.Errorf("unable to convert fee components to USD18: %w", err)
	}

	daConfig, err := c.ccipReader.GetMedianDataAvailabilityGasConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get data availability gas config: %w", err)
	}

	for _, msg := range messages {
		executionCostUSD18 := c.computeExecutionCostUSD18(executionFee, msg)
		dataAvailabilityCostUSD18 := c.computeDataAvailabilityCostUSD18(daFee, daConfig, msg)
		totalCostUSD18 := new(big.Int).Add(executionCostUSD18, dataAvailabilityCostUSD18)
		messageExecCosts[msg.Header.MessageID] = totalCostUSD18
	}

	return messageExecCosts, nil
}

func (c *CCIPMessageExecCostUSD18Calculator) getFeesUSD18(
	ctx context.Context,
	feeComponents types.ChainFeeComponents,
	destChainSelector cciptypes.ChainSelector,
) (plugintypes.USD18, plugintypes.USD18, error) {
	nativeTokenPrices := c.ccipReader.GetWrappedNativeTokenPriceUSD(
		ctx,
		[]cciptypes.ChainSelector{destChainSelector})
	if nativeTokenPrices == nil {
		return nil, nil, fmt.Errorf("unable to get native token prices")
	}
	nativeTokenPrice, ok := nativeTokenPrices[destChainSelector]
	if !ok {
		return nil, nil, fmt.Errorf("missing native token price for chain %s", destChainSelector)
	}

	executionFee := mathslib.CalculateUsdPerUnitGas(feeComponents.ExecutionFee, nativeTokenPrice.Int)
	dataAvailabilityFee := mathslib.CalculateUsdPerUnitGas(feeComponents.DataAvailabilityFee, nativeTokenPrice.Int)

	c.lggr.Debugw("Fee calculation", "nativeTokenPrice", nativeTokenPrice,
		"feeComponents.ExecutionFee", feeComponents.ExecutionFee,
		"feeComponents.DataAvailabilityFee", feeComponents.DataAvailabilityFee,
		"executionFee", executionFee,
		"dataAvailabilityFee", dataAvailabilityFee)

	return executionFee, dataAvailabilityFee, nil
}

// computeExecutionCostUSD18 computes the execution cost of a message in USD18s.
// The cost is:
// messageGas (gas) * executionFee (USD18/gas) = USD18
func (c *CCIPMessageExecCostUSD18Calculator) computeExecutionCostUSD18(
	executionFee *big.Int,
	message cciptypes.Message,
) plugintypes.USD18 {
	messageGas := new(big.Int).SetUint64(c.estimateProvider.CalculateMessageMaxGas(message))
	return new(big.Int).Mul(messageGas, executionFee)
}

// computeDataAvailabilityCostUSD18 computes the data availability cost of a message in USD18s.
func (c *CCIPMessageExecCostUSD18Calculator) computeDataAvailabilityCostUSD18(
	dataAvailabilityFee *big.Int,
	daConfig cciptypes.DataAvailabilityGasConfig,
	message cciptypes.Message,
) plugintypes.USD18 {
	if dataAvailabilityFee == nil || dataAvailabilityFee.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	messageGas := calculateMessageMaxDAGas(message, daConfig)
	return big.NewInt(0).Mul(messageGas, dataAvailabilityFee)
}

// calculateMessageMaxDAGas calculates the total DA gas needed for a CCIP message
func calculateMessageMaxDAGas(
	msg cciptypes.Message,
	daConfig cciptypes.DataAvailabilityGasConfig,
) *big.Int {
	// Calculate token data length
	var totalTokenDataLen int
	for _, tokenAmount := range msg.TokenAmounts {
		totalTokenDataLen += MessageFixedBytesPerToken +
			len(tokenAmount.ExtraData) +
			len(tokenAmount.DestExecData)
	}

	// Calculate total message data length
	dataLen := ConstantMessagePartBytes +
		len(msg.Data) +
		len(msg.Sender) +
		totalTokenDataLen

	// Calculate base gas cost
	dataGas := big.NewInt(int64(dataLen))
	dataGas = new(big.Int).Mul(dataGas, big.NewInt(int64(daConfig.DestGasPerDataAvailabilityByte)))
	dataGas = new(big.Int).Add(dataGas, big.NewInt(int64(daConfig.DestDataAvailabilityOverheadGas)))

	// Then apply the multiplier as: (dataGas * daMultiplier) / multiplierBase
	dataGas = new(big.Int).Mul(dataGas, big.NewInt(int64(daConfig.DestDataAvailabilityMultiplierBps)))
	dataGas = new(big.Int).Div(dataGas, big.NewInt(daMultiplierBase))

	return dataGas
}

var _ MessageExecCostUSD18Calculator = &CCIPMessageExecCostUSD18Calculator{}
