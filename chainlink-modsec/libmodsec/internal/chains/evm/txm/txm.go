package txm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/storage"
)

func NewEvmTxmV2(lggr logger.Logger, chainID *big.Int, client client.Client, fromAddressString string, gasConfig GasEstimator, estimator gas.EvmFeeEstimator, txmConfig txm.Config, keystore *txm.DummyKeystore) (*txm.Txm, error) {
	// AttemptBuilder creates new attempts using Gas Estimator's estimates and signs them using the Keystore
	ab := txm.NewAttemptBuilder(gasConfig.PriceMaxKey, estimator, keystore)

	// Init InMemory storage instead of a Database.
	store := storage.NewInMemoryStoreManager(lggr, chainID)
	fromAddress := common.HexToAddress(fromAddressString)
	if err := store.Add(fromAddress); err != nil {
		lggr.Fatal("Error adding address to InMemory store", err)
		return nil, err
	}

	txmClient := clientwrappers.NewChainClient(client)
	return txm.NewTxm(lggr, chainID, txmClient, ab, store, nil, txmConfig, keystore), nil
}
