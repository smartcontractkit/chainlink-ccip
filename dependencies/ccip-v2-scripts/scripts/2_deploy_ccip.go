package scripts

import (
	"context"

	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	v2logger "github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/model"
	"github.com/smartcontractkit/crib/sdk/ccip"
	"go.uber.org/zap"
)

func DeployCCIPAndAddLanes(logger *zap.SugaredLogger, env config.DevspaceEnv, outputDir string) {
	envConfig := config.NewEnvConfig(env)

	alphaChainSel := config.ChainSelector(1337)
	betaChainSel := config.ChainSelector(2337)

	reader := ccip.NewOutputReader(outputDir)
	addressBook := reader.ReadAddressBook()

	ccipLogger, _ := v2logger.NewLogger()
	ctx := context.Background()

	output, err := crib.DeployCCIPAndAddLanes(ctx, ccipLogger, envConfig, alphaChainSel, betaChainSel, addressBook)
	if err != nil {
		logger.Error("deployment failed due to error", "error", err.Error())
		panic(err)
	}

	envState := model.NewEnvState(logger, env)

	addresses, err := output.AddressBook.Addresses()
	if err != nil {
		logger.Error("failed to get addresses from address book", "error", err.Error())
		panic(err)
	}
	envState.SaveAddressBook(addresses)
	envState.SaveNodeDetails(ccip.NodesDetails{
		NodeIDs: output.NodeIDs,
	})
}
