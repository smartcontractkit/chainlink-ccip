package deploy

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcmstypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type MCMSDeploymentConfig struct {
	Chains         map[uint64]MCMSDeploymentConfigPerChain `json:"chains"`
	AdapterVersion *semver.Version                         `json:"adapterVersion"`
}

type MCMSDeploymentConfigPerChain struct {
	Canceller        mcmstypes.Config `json:"canceller"`
	Bypasser         mcmstypes.Config `json:"bypasser"`
	Proposer         mcmstypes.Config `json:"proposer"`
	TimelockMinDelay *big.Int         `json:"timelockMinDelay"`
	Label            *string          `json:"label"`
	Qualifier        *string          `json:"qualifier"`
	TimelockAdmin    common.Address   `json:"timelockAdmin"`
}

type MCMSDeploymentConfigPerChainWithAddress struct {
	MCMSDeploymentConfigPerChain
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
}

var (
	// testXXXMCMSSigner is a throwaway private key used for signing MCMS proposals.
	// in tests.
	testXXXMCMSSigner *ecdsa.PrivateKey
)

func init() {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	testXXXMCMSSigner = key
}

func SingleGroupTimelockConfigV2() MCMSDeploymentConfigPerChain {
	return MCMSDeploymentConfigPerChain{
		Canceller:        SingleGroupMCMSV2(),
		Bypasser:         SingleGroupMCMSV2(),
		Proposer:         SingleGroupMCMSV2(),
		TimelockMinDelay: big.NewInt(0),
	}
}

func SingleGroupMCMSV2() mcmstypes.Config {
	publicKey := testXXXMCMSSigner.Public().(*ecdsa.PublicKey)
	// Convert the public key to an Ethereum address
	address := crypto.PubkeyToAddress(*publicKey)
	c, err := mcmstypes.NewConfig(1, []common.Address{address}, []mcmstypes.Config{})
	if err != nil {
		panic(err)
	}
	return c
}

func DeployMCMS(deployerReg *DeployerRegistry) cldf.ChangeSetV2[MCMSDeploymentConfig] {
	return cldf.CreateChangeSet(deployMCMSApply(deployerReg), deployMCMSVerify(deployerReg))
}

func deployMCMSVerify(_ *DeployerRegistry) func(cldf.Environment, MCMSDeploymentConfig) error {
	return func(e cldf.Environment, cfg MCMSDeploymentConfig) error {
		// TODO: implement
		return nil
	}
}

func deployMCMSApply(d *DeployerRegistry) func(cldf.Environment, MCMSDeploymentConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg MCMSDeploymentConfig) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()
		for selector, mcmsCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			deployer, exists := d.GetDeployer(family, cfg.AdapterVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no deployer registered for chain family %s and version %s", family, cfg.AdapterVersion.String())
			}
			// find existing addresses for this chain from the env
			existingAddrs := d.ExistingAddressesForChain(e, selector)
			// create the sequence input
			seqCfg := MCMSDeploymentConfigPerChainWithAddress{
				MCMSDeploymentConfigPerChain: mcmsCfg,
				ExistingAddresses:            existingAddrs,
				ChainSelector:                selector,
			}
			deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, deployer.DeployMCMS(), e.BlockChains,
				seqCfg)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy MCMS on chain with selector %d: %w", selector, err)
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
			Build(mcms.Input{}) // for deployment, we don't need an MCMS proposal
	}
}
