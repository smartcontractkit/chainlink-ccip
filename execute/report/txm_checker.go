package report

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type StatusGetter interface {
	// GetTransactionStatus returns the current status of a transaction in the underlying chain's TXM.
	GetTransactionStatus(ctx context.Context, transactionID string) (types.TransactionStatus, error)
}

func getStatusGetters(writers map[ccipocr3.ChainSelector]types.ContractWriter) map[ccipocr3.ChainSelector]StatusGetter {
	// cast ContractWriter to StatusGetter
	statusers := make(map[ccipocr3.ChainSelector]StatusGetter)
	for chain, cw := range writers {
		statusers[chain] = cw
	}

	return statusers
}

// NewTXMCheck creates a new check that queries the TXM for the status of a transaction.
//
// Algorithm ported from OCR2 version: chainlink/core/services/relay/evm/statuschecker/txm_status_checker.go#L31
func NewTXMCheck(statusGetter StatusGetter, maxAttempts uint64) Check {
	return func(
		ctx context.Context,
		lggr logger.Logger,
		msg ccipocr3.Message,
		idx int,
		report exectypes.CommitData,
	) (messageStatus, error) {

		allStatuses := make([]types.TransactionStatus, 0)

		var counter uint64
		for {
			transactionID := fmt.Sprintf("%s-%d", msg.Header.MessageID.String(), counter)
			status, err := statusGetter.GetTransactionStatus(ctx, transactionID)
			if err != nil && status == types.Unknown {
				// If the status is unknown and err not nil, it means the transaction was not found
				break
			}
			allStatuses = append(allStatuses, status)
			counter++

			// Break the loop if the cap is reached
			if counter >= maxAttempts {
				lggr.Infow(
					"skipping message due to txm max attempts reached",
					"message",
					msg,
					"maxAttempts",
					maxAttempts,
				)
				return TXMCheckError, nil
			}
		}
		return None, nil
	}
}
