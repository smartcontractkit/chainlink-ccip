package scripts

import (
	"context"
	"os"

	"github.com/smartcontractkit/chainlink/deployment/ccip/changeset"
	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"
	v2logger "github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/model"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/rmn"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/utils"
)

type RMNConfigurer struct {
	ccipLogger     v2logger.Logger
	devspaceEnv    config.DevspaceEnv
	nodeCount      int
	envState       model.CCIPEnvState
	envStateReader *crib.OutputReader
	envConfig      *devenv.EnvironmentConfig
}

func NewRMNConfigurer(devspaceEnv config.DevspaceEnv, nodeCount int) RMNConfigurer {
	ccipLogger, _ := v2logger.NewLogger()
	envState := model.NewEnvState(logger, devspaceEnv)
	stateDirPath := devspaceEnv.TmpDir
	reader := crib.NewOutputReader(stateDirPath)

	envConfig, err := config.GetEnvConfig(devspaceEnv)
	if err != nil {
		logger.Fatalw("unable to get envconfig", "err", err)
		os.Exit(1)
	}

	return RMNConfigurer{
		ccipLogger:     ccipLogger,
		devspaceEnv:    devspaceEnv,
		nodeCount:      nodeCount,
		envConfig:      envConfig,
		envState:       envState,
		envStateReader: reader,
	}
}

func (r RMNConfigurer) SetupRMNOnChain() {
	rmnNodes := r.GenerateNodeIdentities()

	ctx := context.Background()
	ccipLogger, _ := v2logger.NewLogger()

	alphaChainSel := config.ChainSelector(1337)
	betaChainSel := config.ChainSelector(2337)

	addressBook, _ := r.envStateReader.ReadAddressBook()

	logger.Infow("Setting up RMN On Chains")
	output, err := crib.SetupRMNNodeOnAllChains(ctx, ccipLogger, *r.envConfig, alphaChainSel, betaChainSel, addressBook, rmnNodes)
	if err != nil {
		logger.Fatalw("failed to setup rmn node on all chains", "err", err)
		os.Exit(1)
	}
	logger.Infow("Finished setting up RMN On Chains")

	addresses, err := output.AddressBook.Addresses()
	if err != nil {
		logger.Fatalw("unable to get addresses", "err", err)
		os.Exit(1)
	}
	r.envState.SaveAddressBook(addresses)
}

func (r RMNConfigurer) GenerateTOMLConfigs() {
	envState := model.NewEnvState(logger, r.devspaceEnv)
	if envState.RMNTomlConfigsExists() {
		logger.Info("RMN TOML Configs already exists, reusing existing configs")
		return
	}

	devEnv, _, err := devenv.NewEnvironment(func() context.Context {
		ctx := context.Background()
		return ctx
	}, r.ccipLogger, *r.envConfig)
	if err != nil {
		logger.Fatalw("unable to create devenv", "err", err)
		os.Exit(1)
	}
	addressBook, _ := r.envStateReader.ReadAddressBook()
	devEnv.ExistingAddresses = addressBook
	state, err := changeset.LoadOnchainState(*devEnv)
	if err != nil {
		logger.Fatalw("failed to load chain state: %w", err)
		os.Exit(1)
	}

	logger.Debugw("generating shared-toml config")

	nodesDetails, _ := r.envStateReader.ReadNodesDetails()

	err = rmn.GenerateSharedConfigTOML(state, r.envConfig, r.envState, nodesDetails.BootstrapNode)
	if err != nil {
		logger.Fatalw("failed to generate shared config", "err", err)
		os.Exit(1)
	}

	logger.Debugw("generating local-toml config")
	err = rmn.GenerateLocalToml(r.devspaceEnv, r.envConfig.Chains, r.envState)
	if err != nil {
		logger.Fatalw("failed to generate local TOML", "err", err)
		os.Exit(1)
	}
}

func (r RMNConfigurer) GenerateNodeIdentities() []crib.RMNNodeConfig {
	envState := model.NewEnvState(logger, r.devspaceEnv)
	if envState.RMNIdentitiesExists() {
		logger.Info("RMN identities already exists, reusing existing identities")
		reader := crib.NewOutputReader(r.devspaceEnv.TmpDir)
		rmnConfig, _ := reader.ReadRMNNodeConfigs()
		return rmnConfig
	}

	logger.Info("Generating RMN config")

	rageProxyImageURI := "804282218731.dkr.ecr.us-west-2.amazonaws.com/rageproxy"
	rageProxyImageTag := "c0eebdc"
	afn2ProxyImageURI := "804282218731.dkr.ecr.us-west-2.amazonaws.com/afn2proxy"
	afn2ProxyImageTag := "master-f1878a1"

	imagePlatform := "linux/amd64"

	bounds, err := utils.SafeConvertWithBounds(r.nodeCount)
	utils.ExitOnError(err, logger)

	identities, err := crib.GenerateRMNNodeIdentities(
		bounds, rageProxyImageURI, rageProxyImageTag, afn2ProxyImageURI, afn2ProxyImageTag, imagePlatform,
	)
	utils.ExitOnError(err, logger)

	envState.SaveRMNNodeConfigs(identities)
	logger.Infow("RMN config generated", "identities", identities)
	return identities
}
