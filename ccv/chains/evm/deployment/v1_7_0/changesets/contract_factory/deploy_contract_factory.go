package contract_factory

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/contract_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var DeployContractFactory = cldf_deployment.CreateChangeSet(applyDeployContractFactory, verifyDeployContractFactory)

type DeployContractFactoryCfg struct {
	ChainSel  uint64
	AllowList []common.Address
}

func applyDeployContractFactory(e cldf_deployment.Environment, cfg DeployContractFactoryCfg) (cldf_deployment.ChangesetOutput, error) {
	evmChain, ok := e.BlockChains.EVMChains()[cfg.ChainSel]
	if !ok {
		return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain with selector %d not found in environment", cfg.ChainSel)
	}

	deployReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, contract_factory.Deploy, evmChain, contract.DeployInput[contract_factory.ConstructorArgs]{
		ChainSelector:  cfg.ChainSel,
		TypeAndVersion: deployment.NewTypeAndVersion(contract_factory.ContractType, *semver.MustParse("1.7.0")),
		Args: contract_factory.ConstructorArgs{
			AllowList: cfg.AllowList,
		},
	})
	if err != nil {
		return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy contract factory: %w", err)
	}

	ds := datastore.NewMemoryDataStore()
	if err := ds.Addresses().Add(deployReport.Output); err != nil {
		return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to add contract factory address to datastore: %w", err)
	}

	return cldf_deployment.ChangesetOutput{
		DataStore: ds,
		Reports:   []operations.Report[any, any]{deployReport.ToGenericReport()},
	}, nil
}

func verifyDeployContractFactory(e cldf_deployment.Environment, cfg DeployContractFactoryCfg) error {
	return nil
}
