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
	ctx := context.Background()

	if h.shouldSkip() {
		h.logger.Info("AddressBook already exists, assuming that home chain is already deployed. skipping")
		return
	}
	// Use the output value
	h.logger.Info("Deploying home chain",
		zap.String("tmp_dir", h.env.TmpDir),
	)

	ccipLogger, _ := v2logger.NewLogger()

	maybeGHAJWTToken := ""
	var err error

	if h.env.CIEnv {
		gap.TestWSConnectionViaGAP(h.env)
		maybeGHAJWTToken, err = gap.FetchJWTTokenForGAP(ctx)
		if err != nil {
			h.logger.Fatal("failed to fetch GA JWT", "error", err)
			os.Exit(1)
		}
	}

	h.logger.Info("debug token", "token", maybeGHAJWTToken)

	envConfig, err := config.GetEnvConfig(h.env, maybeGHAJWTToken)
	if err != nil {
		h.logger.Fatal(err)
		os.Exit(1)
	}
	transmittedChainConfigs := config.GetTransmittedChainConfigs(h.env)

	homeChainID := uint64(1337)
	feedChainID := uint64(2337)
	homeChainSelector := config.ChainSelector(homeChainID)
	feedChainSelector := config.ChainSelector(feedChainID)

	h.logger.Info("Deploying Home Chain Contracts",
		zap.String("jd-grpc-endpoint", envConfig.JDConfig.GRPC),
	)
	capRegConfig, addressBook, err := crib.DeployHomeChainContracts(ctx, ccipLogger, *envConfig, homeChainSelector, feedChainSelector)
	if err != nil {
		h.logger.Fatal("unable to deploy home chain contracts", "err", err)
		os.Exit(1)
	}

	h.envState.SaveNodesTomlOverride(capRegConfig, homeChainID)

	addresses, err := addressBook.Addresses()
	if err != nil {
		h.logger.Fatal("unable to get addresses", "err", err)
		os.Exit(1)
	}

	// Save State to files
	h.envState.SaveAddressBook(addresses)
	h.envState.SaveChainConfigs(transmittedChainConfigs)
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
