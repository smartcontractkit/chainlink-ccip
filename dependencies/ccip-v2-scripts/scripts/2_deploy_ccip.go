package scripts

import (
	"context"
	"os"

	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	v2logger "github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/gap"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/model"
	"go.uber.org/zap"
)

func DeployCCIPAndAddLanes(logger *zap.SugaredLogger, env config.DevspaceEnv, outputDir string) {
	ctx := context.Background()
	var err error

	if env.CIEnv {
		_, err = gap.FetchJWTTokenForGAP(ctx)
		if err != nil {
			logger.Fatal("failed to fetch GA JWT", "error", err)
			os.Exit(1)
		}
	}

	envConfig, err := config.GetEnvConfig(env)
	if err != nil {
		logger.Fatal("unable to deploy ccip and add lanes", "err", err)
		os.Exit(1)
	}

	alphaChainSel := config.ChainSelector(1337)
	betaChainSel := config.ChainSelector(2337)

	reader := crib.NewOutputReader(outputDir)
	addressBook := reader.ReadAddressBook()

	ccipLogger, _ := v2logger.NewLogger()

	output, err := crib.DeployCCIPAndAddLanes(ctx, ccipLogger, *envConfig, alphaChainSel, betaChainSel, addressBook)
	if err != nil {
		logger.Fatal("deployment failed due to error", "error", err.Error())
		os.Exit(1)
	}

	envState := model.NewEnvState(logger, env)

	addresses, err := output.AddressBook.Addresses()
	if err != nil {
		logger.Fatal("failed to get addresses from address book", "error", err.Error())
		os.Exit(1)
	}
	envState.SaveAddressBook(addresses)
	envState.SaveNodeDetails(crib.NodesDetails{
		NodeIDs: output.NodeIDs,
	})
}
