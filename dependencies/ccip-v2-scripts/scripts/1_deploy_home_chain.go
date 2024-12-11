package scripts

import (
	"context"

	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	v2logger "github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/model"
	"go.uber.org/zap"
)

func DeployHomeChain(logger *zap.SugaredLogger, env config.DevspaceEnv) {
	deployer := NewHomeChainDeployer(logger, env)
	deployer.Deploy()
}

type HomeChainDeployer struct {
	logger   *zap.SugaredLogger
	env      config.DevspaceEnv
	envState model.CCIPEnvState
}

func NewHomeChainDeployer(logger *zap.SugaredLogger, env config.DevspaceEnv) *HomeChainDeployer {
	return &HomeChainDeployer{
		logger:   logger,
		env:      env,
		envState: model.NewEnvState(logger, env),
	}
}

func (h HomeChainDeployer) Deploy() {
	if h.shouldSkip() {
		h.logger.Info("AddressBook already exists, assuming that home chain is already deployed. skipping")
		return
	}
	// Use the output value
	h.logger.Info("Deploying home chain",
		zap.String("tmp_dir", h.env.TmpDir),
	)

	ccipLogger, _ := v2logger.NewLogger()
	envConfig := config.NewEnvConfig(h.env)

	homeChainID := uint64(1337)
	feedChainID := uint64(2337)
	homeChainSelector := config.ChainSelector(homeChainID)
	feedChainSelector := config.ChainSelector(feedChainID)

	ctx := context.Background()
	capRegConfig, addressBook, err := crib.DeployHomeChainContracts(ctx, ccipLogger, envConfig, homeChainSelector, feedChainSelector)
	if err != nil {
		panic(err)
	}

	h.envState.SaveNodesTomlOverride(capRegConfig, homeChainID)

	addresses, err := addressBook.Addresses()
	if err != nil {
		panic(err)
	}

	// Save State to files
	h.envState.SaveAddressBook(addresses)
	h.envState.SaveChainConfigs(envConfig.Chains)
}

func (h HomeChainDeployer) shouldSkip() bool {
	if h.envState.AddressBookExists() {
		h.logger.Info("Address book file exists")
		h.logger.Info("Skipping Home Chain Deployment, Home chain is already deployed")
		return true
	}

	h.logger.Info("Address book file does not exist")

	return false
}
