package scripts

import (
	"context"
	"os"

	"github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"
	v2logger "github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/model"
	"go.uber.org/zap"
)

func DeployCCIPAndAddLanes(logger *zap.SugaredLogger, env config.DevspaceEnv) {
	logger.Infow("Deploying CCIP Contracts and Adding lanes", "environment", env)
	CallDeployerFn(logger, env, env.TmpDir, crib.DeployCCIPAndAddLanes)
}

func ConfigureOCR(logger *zap.SugaredLogger, env config.DevspaceEnv) {
	logger.Infow("Configuring OCR", "environment", env)
	CallDeployerFn(logger, env, env.TmpDir, crib.ConfigureCCIPOCR)
}

type CCIPDeployerCallerFn func(ctx context.Context, lggr v2logger.Logger, envConfig devenv.EnvironmentConfig, homeChainSel, feedChainSel uint64, ab deployment.AddressBook) (crib.DeployCCIPOutput, error)

func CallDeployerFn(logger *zap.SugaredLogger, env config.DevspaceEnv, stateDirPath string, deployerFn CCIPDeployerCallerFn) {
	ccipLogger, _ := v2logger.NewLogger()

	alphaChainSel := config.ChainSelector(1337)
	betaChainSel := config.ChainSelector(2337)

	reader := crib.NewOutputReader(stateDirPath)
	addressBook := reader.ReadAddressBook()

	envConfig, err := config.GetEnvConfig(env)
	if err != nil {
		logger.Fatal("unable to get envconfig", "err", err)
		os.Exit(1)
	}

	output, err := deployerFn(context.Background(), ccipLogger, *envConfig, alphaChainSel, betaChainSel, addressBook)
	if err != nil {
		logger.Fatal("deployer function failed with error", "err", err)
		os.Exit(1)
	}

	envState := model.NewEnvState(logger, env)
	addresses, err := output.AddressBook.Addresses()
	if err != nil {
		logger.Fatal("unable to get addresses", "err", err)
		os.Exit(1)
	}
	envState.SaveAddressBook(addresses)
	envState.SaveNodeDetails(crib.NodesDetails{
		NodeIDs: output.NodeIDs,
	})
}
