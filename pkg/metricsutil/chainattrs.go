// Package metricsutil holds chain-identity attribute helpers shared by every CCIP metrics
// reporter.
// Every reporter needs to turn a chain selector into the same chainID/chainFamily/chainName/
// chainSelector attribute set. Centralizing that resolution here means there's exactly one
// place to fix if the fallback behavior or label-key naming ever needs to change, instead of N
// copies silently drifting apart.
package metricsutil

import (
	"strconv"

	"go.opentelemetry.io/otel/attribute"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs"
)

// UnknownLabelValue is used for any chain-identity label that couldn't be resolved (e.g. an
// out-of-date chain-selectors lib not recognizing a selector).
// Pulled into a constant so that the linter doesn't complain.
const UnknownLabelValue = "unknown"

// ChainAttrs resolves chainSelector into an unprefixed chainID/chainFamily/chainName/
// chainSelector attribute set. Use this when a metric only ever describes one chain at a time.
//
// If a single metric can carry two different chains at once (e.g. a per-lane metric with both a
// source and a destination, or a reader/config-poller instance that's scoped to one destination
// chain but reads about others), use DestChainAttrs/SourceChainAttrs for the other side instead
// of calling this twice -- otherwise both chains attach identically-named labels and collide.
func ChainAttrs(chainSelector ccipocr3.ChainSelector) []attribute.KeyValue {
	return chainAttrsWithPrefix("", chainSelector)
}

// DestChainAttrs resolves chainSelector into destChainID/destChainFamily/destChainName/
// destChainSelector attributes.
func DestChainAttrs(chainSelector ccipocr3.ChainSelector) []attribute.KeyValue {
	return chainAttrsWithPrefix("dest", chainSelector)
}

// SourceChainAttrs resolves chainSelector into sourceChainID/sourceChainFamily/sourceChainName/
// sourceChainSelector attributes.
func SourceChainAttrs(chainSelector ccipocr3.ChainSelector) []attribute.KeyValue {
	return chainAttrsWithPrefix("source", chainSelector)
}

func chainAttrsWithPrefix(prefix string, chainSelector ccipocr3.ChainSelector) []attribute.KeyValue {
	// Could happen due to an out-of-date chain-selectors lib.
	chainFamily, chainID, ok := libs.GetChainInfoFromSelector(chainSelector)
	if !ok {
		// graceful fallback - we could even alert on such a thing.
		chainFamily = UnknownLabelValue
		chainID = UnknownLabelValue
	}

	chainName, err := libs.GetNameFromIDAndFamily(chainID, chainFamily)
	if err != nil {
		chainName = UnknownLabelValue
	}

	return []attribute.KeyValue{
		attribute.String(labelKey(prefix, "ID"), chainID),
		attribute.String(labelKey(prefix, "Family"), chainFamily),
		attribute.String(labelKey(prefix, "Name"), chainName),
		attribute.String(labelKey(prefix, "Selector"), strconv.FormatUint(uint64(chainSelector), 10)),
	}
}

// labelKey builds "chainID" (no prefix) or "destChainID"/"sourceChainID" (with a prefix).
func labelKey(prefix, suffix string) string {
	if prefix == "" {
		return "chain" + suffix
	}
	return prefix + "Chain" + suffix
}
