package contractreader

import (
	"context"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"math/big"
)

// ContractWriterFacade wraps the public functions of ContractWriter in chainlink-common so that we can mock it.
//
//nolint:lll // don't read this interface.
type ContractWriterFacade interface {
	services.Service

	// SubmitTransaction packs and broadcasts a transaction to the underlying chain.
	//
	// - `args` should be any object which maps a set of method param into the contract and method specific method params.
	// - `transactionID` will be used by the underlying TXM as an idempotency key, and unique reference to track transaction attempts.
	SubmitTransaction(ctx context.Context, contractName, method string, args any, transactionID string, toAddress string, meta *types.TxMeta, value *big.Int) error

	// GetTransactionStatus returns the current status of a transaction in the underlying chain's TXM.
	GetTransactionStatus(ctx context.Context, transactionID string) (types.TransactionStatus, error)

	// GetFeeComponents retrieves the associated gas costs for executing a transaction.
	GetFeeComponents(ctx context.Context) (*types.ChainFeeComponents, error)
	//mustEmbedUnimplementedContractWriterServer()
}
