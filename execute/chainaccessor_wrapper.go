package execute

import (
	"context"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

// ChainAccessorWrapper wraps a ChainAccessor to control TxHash population
type ChainAccessorWrapper struct {
	inner                 cciptypes.ChainAccessor
	populateTxHashEnabled bool
}

// NewChainAccessorWrapper creates a new wrapper that can control TxHash population
func NewChainAccessorWrapper(inner cciptypes.ChainAccessor, populateTxHashEnabled bool) cciptypes.ChainAccessor {
	return &ChainAccessorWrapper{
		inner:                 inner,
		populateTxHashEnabled: populateTxHashEnabled,
	}
}

// MsgsBetweenSeqNums reads the provided chains, finds and returns ccip messages
// submitted between the provided sequence numbers. If populateTxHashEnabled is false,
// it will clear the TxHash field from the returned messages.
func (w *ChainAccessorWrapper) MsgsBetweenSeqNums(
	ctx context.Context,
	chain cciptypes.ChainSelector,
	seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	messages, err := w.inner.MsgsBetweenSeqNums(ctx, chain, seqNumRange)
	if err != nil {
		return nil, err
	}

	// If TxHash population is disabled, clear the TxHash field
	if !w.populateTxHashEnabled {
		for i := range messages {
			messages[i].Header.TxHash = ""
		}
	}

	return messages, nil
}

// All other methods are delegated to the inner ChainAccessor

func (w *ChainAccessorWrapper) GetContractAddress(contractName string) ([]byte, error) {
	return w.inner.GetContractAddress(contractName)
}

func (w *ChainAccessorWrapper) GetAllConfigsLegacy(
	ctx context.Context,
	destChainSelector cciptypes.ChainSelector,
	sourceChainSelectors []cciptypes.ChainSelector,
) (cciptypes.ChainConfigSnapshot, map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error) {
	return w.inner.GetAllConfigsLegacy(ctx, destChainSelector, sourceChainSelectors)
}

func (w *ChainAccessorWrapper) GetChainFeeComponents(ctx context.Context) (cciptypes.ChainFeeComponents, error) {
	return w.inner.GetChainFeeComponents(ctx)
}

func (w *ChainAccessorWrapper) Sync(
	ctx context.Context,
	contractName string,
	contractAddress cciptypes.UnknownAddress,
) error {
	return w.inner.Sync(ctx, contractName, contractAddress)
}

func (w *ChainAccessorWrapper) CommitReportsGTETimestamp(
	ctx context.Context,
	ts time.Time,
	confidence primitives.ConfidenceLevel,
	limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	return w.inner.CommitReportsGTETimestamp(ctx, ts, confidence, limit)
}

func (w *ChainAccessorWrapper) ExecutedMessages(
	ctx context.Context,
	rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	confidence primitives.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	return w.inner.ExecutedMessages(ctx, rangesPerChain, confidence)
}

func (w *ChainAccessorWrapper) NextSeqNum(ctx context.Context, sources []cciptypes.ChainSelector) (
	map[cciptypes.ChainSelector]cciptypes.SeqNum, error) {
	return w.inner.NextSeqNum(ctx, sources)
}

func (w *ChainAccessorWrapper) Nonces(
	ctx context.Context,
	addresses map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	return w.inner.Nonces(ctx, addresses)
}

func (w *ChainAccessorWrapper) GetChainFeePriceUpdate(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.TimestampedUnixBig, error) {
	return w.inner.GetChainFeePriceUpdate(ctx, selectors)
}

func (w *ChainAccessorWrapper) GetLatestPriceSeqNr(ctx context.Context) (cciptypes.SeqNum, error) {
	return w.inner.GetLatestPriceSeqNr(ctx)
}

// SourceAccessor methods
func (w *ChainAccessorWrapper) LatestMessageTo(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	return w.inner.LatestMessageTo(ctx, dest)
}

func (w *ChainAccessorWrapper) GetExpectedNextSequenceNumber(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	return w.inner.GetExpectedNextSequenceNumber(ctx, dest)
}

func (w *ChainAccessorWrapper) GetTokenPriceUSD(
	ctx context.Context,
	address cciptypes.UnknownAddress,
) (cciptypes.TimestampedUnixBig, error) {
	return w.inner.GetTokenPriceUSD(ctx, address)
}

func (w *ChainAccessorWrapper) GetFeeQuoterDestChainConfig(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.FeeQuoterDestChainConfig, error) {
	return w.inner.GetFeeQuoterDestChainConfig(ctx, dest)
}

// USDCMessageReader methods
func (w *ChainAccessorWrapper) MessagesByTokenID(
	ctx context.Context,
	source, dest cciptypes.ChainSelector,
	tokens map[cciptypes.MessageTokenID]cciptypes.RampTokenAmount,
) (map[cciptypes.MessageTokenID]cciptypes.Bytes, error) {
	return w.inner.MessagesByTokenID(ctx, source, dest, tokens)
}

// PriceReader methods
func (w *ChainAccessorWrapper) GetFeedPricesUSD(
	ctx context.Context,
	tokens []cciptypes.UnknownEncodedAddress,
	tokenInfo map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo,
) (cciptypes.TokenPriceMap, error) {
	return w.inner.GetFeedPricesUSD(ctx, tokens, tokenInfo)
}

func (w *ChainAccessorWrapper) GetFeeQuoterTokenUpdates(
	ctx context.Context,
	tokensBytes []cciptypes.UnknownAddress,
) (map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedUnixBig, error) {
	return w.inner.GetFeeQuoterTokenUpdates(ctx, tokensBytes)
}
