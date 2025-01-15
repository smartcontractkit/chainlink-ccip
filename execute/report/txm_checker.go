package report

import (
	"context"

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

func NewTXMCheck(txm map[ccipocr3.ChainSelector]StatusGetter) Check {
	return func(lggr logger.Logger, msg ccipocr3.Message, idx int, report exectypes.CommitData) (messageStatus, error) {
		var cw types.ContractWriter
		cw.GetTransactionStatus()
		return None, nil
	}
}
