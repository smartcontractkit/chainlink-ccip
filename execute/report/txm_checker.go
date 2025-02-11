package report

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// StatusGetter is a narrowed interface to the contract writer that only exposes the GetTransactionStatus method.
type StatusGetter interface {
	// GetTransactionStatus returns the current status of a transaction in the underlying chain's TXM.
	GetTransactionStatus(ctx context.Context, transactionID string) (types.TransactionStatus, error)
}

type MessageStatusDetails struct {
	numAttempts    uint64
	hasFatalStatus bool
}

// MessageStatusCache is a cache of status metadata needed for each message. It is multi-purpose, intended
// to support report checking and for computing the next txID for the transmitter.
//
// This object needs to be recreated for each round in order to refresh the cache.
type MessageStatusCache struct {
	statusGetter StatusGetter

	cache map[string]MessageStatusDetails
}

// NewMessageStatusCache creates a new cache for message statuses.
func NewMessageStatusCache(statusGetter StatusGetter) *MessageStatusCache {
	return &MessageStatusCache{
		statusGetter: statusGetter,
		cache:        make(map[string]MessageStatusDetails),
	}
}

func toID(msgID string, attempt uint64) string {
	return fmt.Sprintf("%s-%d", msgID, attempt)
}

// NextTransactionID returns the next transaction ID for the given message ID.
func (m *MessageStatusCache) NextTransactionID(msgID string) string {
	// this works with a zero value, so no need to hceck if cache[msgID] exists.
	return toID(msgID, m.cache[msgID].numAttempts)
}

// statuses fetches the statusu details for a given message ID. Results are cached for subsequent calls.
func (m *MessageStatusCache) statuses(ctx context.Context, msgID string) (MessageStatusDetails, error) {
	if details, ok := m.cache[msgID]; ok {
		return details, nil
	}

	var details MessageStatusDetails
	for {
		transactionID := toID(msgID, details.numAttempts)
		status, err := m.statusGetter.GetTransactionStatus(ctx, transactionID)

		if err != nil && status == types.Unknown {
			// TODO: Would be nice to have a typed error here as a more deliberate check.
			// By code inspection we see that if the status is unknown and err not nil,
			// it means the transaction was not found. It will be the next txID.
			break
		} else if err != nil {
			// TODO: Why wasn't this case wasn't handled in 1.5?
			return MessageStatusDetails{}, err
		}
		if status == types.Fatal {
			details.hasFatalStatus = true
		}
		details.numAttempts++
	}

	m.cache[msgID] = details
	return details, nil
}

// NewTXMCheck creates a new check that queries the TXM for the status of a transaction.
//
// Algorithm ported from OCR2 version: chainlink/core/services/relay/evm/statuschecker/txm_status_checker.go#L31
func NewTXMCheck(statuses *MessageStatusCache, maxAttempts uint64) Check {
	return func(
		ctx context.Context,
		lggr logger.Logger,
		msg ccipocr3.Message,
		_ int,
		_ exectypes.CommitData,
	) (messageStatus, error) {
		details, err := statuses.statuses(ctx, msg.Header.MessageID.String())
		if err != nil {
			return None, err
		}

		if details.hasFatalStatus {
			lggr.Infow("Skipping message - found a fatal TXM status",
				"message", msg,
			)
			return TXMFatalStatus, nil
		}

		if details.numAttempts >= maxAttempts {
			lggr.Infow(
				"skipping message due to maximum number of statuses reached, possible infinite loop",
				"message", msg,
				"maxAttempts", maxAttempts,
			)
			return TXMCheckError, nil
		}

		return None, nil
	}
}
