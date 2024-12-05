package scripts

import (
	"fmt"
	"log/slog"

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
		envState: model.NewEnvState(env),
	}
}

func (h HomeChainDeployer) Deploy() {
	if h.shouldSkip() {
		h.logger.Info("AddressBook already exists, assuming that home chain is already deployed. skipping")
		return
	}
	// Use the output value
	h.logger.Info("Deploying home chain",
		slog.String("tmp_dir", h.env.TmpDir),
	)

	ccipLogger, _ := v2logger.NewLogger()
	envConfig := config.NewEnvConfig(h.env)

	homeChainID := uint64(1337)
	homeChainSelector := config.ChainSelector(homeChainID)

	capRegConfig, addressBook, err := crib.DeployHomeChainContracts(ccipLogger, envConfig, homeChainSelector)
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
		fmt.Printf("Address book file exists\n")
		slog.Info("Skipping Home Chain Deployment, Home chain is already deployed")
		return true
	}

	fmt.Printf("Address book file does not exist\n")

	return false
}
