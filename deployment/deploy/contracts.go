package deploy

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type ContractDeploymentConfig struct {
	Chains map[uint64]ContractDeploymentConfigPerChain
	// MCMS configures the resulting proposal if we're transferring to MCMS post-deployment.
	MCMS mcms.Input
}

type ContractDeploymentConfigPerChain struct {
	Version *semver.Version
	// LINK TOKEN CONFIG
	// token private key used to deploy the LINK token. Solana: base58 encoded private key
	TokenPrivKey string
	// token decimals used to deploy the LINK token
	TokenDecimals uint8
	// FEE QUOTER CONFIG
	MaxFeeJuelsPerMsg            *big.Int
	TokenPriceStalenessThreshold uint32
	// Expressed in Wei per Eth on EVM chains
	LinkPremiumMultiplier uint64
	// WETH on EVM chains
	// Expressed in Wei per Eth on EVM chains
	NativeTokenPremiumMultiplier uint64
	// OFFRAMP CONFIG
	// Manual execution can be performed after this threshold (in seconds)
	PermissionLessExecutionThresholdSeconds uint32
	// EVM only.
	GasForCallExactCheck uint16
	// EVM only. Validates incoming messages to offramp
	MessageInterceptor string
	// RMN REMOTE CONFIG
	LegacyRMN string
}

type ContractDeploymentConfigPerChainWithAddress struct {
	ContractDeploymentConfigPerChain
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
}

func DeployContracts(deployerReg *DeployerRegistry) cldf.ChangeSetV2[ContractDeploymentConfig] {
	return cldf.CreateChangeSet(deployContractsApply(deployerReg), deployContractsVerify(deployerReg))
}

func deployContractsVerify(_ *DeployerRegistry) func(cldf.Environment, ContractDeploymentConfig) error {
	return func(e cldf.Environment, cfg ContractDeploymentConfig) error {
		for selector, config := range cfg.Chains {
			_, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return fmt.Errorf("no selector %d found in environment: %w", selector, err)
			}
			if config.Version == nil {
				return fmt.Errorf("no version specified for chain with selector %d", selector)
			}
		}
		// TODO: implement
		return nil
	}
}

func deployContractsApply(d *DeployerRegistry) func(cldf.Environment, ContractDeploymentConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ContractDeploymentConfig) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()
		for selector, contractCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			deployer, exists := d.GetDeployer(family, contractCfg.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no deployer registered for chain family %s and version %s", family, contractCfg.Version.String())
			}
			// find existing addresses for this chain from the env
			existingAddrs := d.ExistingAddressesForChain(e, selector)
			// create the sequence input
			seqCfg := ContractDeploymentConfigPerChainWithAddress{
				ContractDeploymentConfigPerChain: contractCfg,
				ExistingAddresses:                existingAddrs,
				ChainSelector:                    selector,
			}
			deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, deployer.DeployChainContracts(), e.BlockChains,
				seqCfg)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy Contract on chain with selector %d: %w", selector, err)
			}
			for _, r := range deployReport.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}
			reports = append(reports, deployReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithDataStore(ds).
			Build(cfg.MCMS)
	}
}
