package main

import (
	"context"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/mailbox"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	evmConfig "github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/internal/chains/evm/config"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/internal/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/internal/chains/evm/multinode"
	evmTxm "github.com/smartcontractkit/chainlink-modsec/libmodsec/internal/chains/evm/txm"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/internal/runner/db"
)

const (
	dbURL             = "postgresql://postgres:@localhost:5432/postgres?sslmode=disable"
	privateKeyString  = "" // Private key without 0x prefix.
	fromAddressString = "" // From address.
)

// Should be the effective defaults you'll see in a CL node for a single chain.
const configTOML = ``

func main() {
	ctx := context.Background()

	// Init Logger.
	lggrCfg := logger.Config{Level: zapcore.InfoLevel}
	lggr, _ := lggrCfg.New()

	// Open database.
	db, err := db.CreateDB(ctx, lggr, dbURL)
	if err != nil {
		lggr.Fatal("Error creating database", err)
		return
	}

	// Create EVM chain config.
	evmConfig, err := config.CreateNewEVMChainFromTOML(lggr, configTOML)
	if err != nil {
		lggr.Fatal("Error creating EVM chain config", err)
		return
	}

	// Create MultiNode.
	multiNode, err := multinode.NewMultiNode(ctx, lggr, evmConfig)
	if err != nil {
		lggr.Fatal("Error creating MultiNode", err)
		return
	}

	// Start Mailbox Monitor + LogPoller
	mailMon := mailbox.NewMonitor("1", lggr)
	mailMon.Start(ctx)
	lp := logpoller.NewLogPoller(ctx, evmConfig, lggr, db, multiNode, mailMon)
	lp.Start(ctx)

	// Create TXM
	txm, err := createTXM(lggr, multiNode, evmConfig)
	if err != nil {
		lggr.Fatal("Error creating TXM", err)
		return
	}
	txm.Start(ctx)
	lggr.Info("TXM started")

	// Wait 10 seconds, so you can see that LogPoller and HeadTracker are working in the background.
	time.Sleep(10 * time.Second)

	// Get latest block from log poller.
	block, err := lp.LatestBlock(ctx)
	if err != nil {
		lggr.Fatal("Error getting latest block", err)
	}

	lggr.Info("Latest block from LogPoller", "block", block)
}

func createTXM(lggr logger.Logger, multiNode client.Client, evmConfig *evmConfig.ChainScoped) (*txm.Txm, error) {
	// Init Gas Estimator.
	estimator, err := evmTxm.InitGasEstimator(lggr, multiNode, evmConfig.EVM().ChainID(), evmTxm.GetGasConfig())
	if err != nil {
		lggr.Fatal("Error during estimator init", err)
		return nil, err
	}

	// Init Dummy Keystore
	keystore := txm.NewKeystore(evmConfig.EVM().ChainID())
	if err := keystore.Add(privateKeyString); err != nil {
		lggr.Fatal("Error adding key to keystore", err)
		return nil, err
	}

	// Init TXMv2
	txmConfig := txm.Config{
		EIP1559:             true,
		BlockTime:           12 * time.Second,
		RetryBlockThreshold: 3,
		EmptyTxLimitDefault: evmTxm.GetGasConfig().LimitDefaultF,
	}

	txm, err := evmTxm.NewEvmTxmV2(lggr, evmConfig.EVM().ChainID(), multiNode, fromAddressString, evmTxm.GetGasConfig(), estimator, txmConfig, keystore)
	if err != nil {
		lggr.Fatal("Error during txm init", err)
		return nil, err
	}

	return txm, nil
}
