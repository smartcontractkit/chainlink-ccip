package create2_factory

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var DeployCREATE2Factory = cldf_deployment.CreateChangeSet(applyDeployCREATE2Factory, verifyDeployCREATE2Factory)

type DeployCREATE2FactoryCfg struct {
	ChainSel  uint64
	AllowList []common.Address
}

func applyDeployCREATE2Factory(e cldf_deployment.Environment, cfg DeployCREATE2FactoryCfg) (cldf_deployment.ChangesetOutput, error) {
	evmChain, ok := e.BlockChains.EVMChains()[cfg.ChainSel]
	if !ok {
		return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain with selector %d not found in environment", cfg.ChainSel)
	}

	deployReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, create2_factory.Deploy, evmChain, contract.DeployInput[create2_factory.ConstructorArgs]{
		ChainSelector:  cfg.ChainSel,
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
		Args: create2_factory.ConstructorArgs{
			AllowList: cfg.AllowList,
		},
	})
	if err != nil {
		return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy create2 factory: %w", err)
	}

	ds := datastore.NewMemoryDataStore()
	if err := ds.Addresses().Add(deployReport.Output); err != nil {
		return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to add create2 factory address to datastore: %w", err)
	}

	return cldf_deployment.ChangesetOutput{
		DataStore: ds,
		Reports:   []operations.Report[any, any]{deployReport.ToGenericReport()},
	}, nil
}

func verifyDeployCREATE2Factory(e cldf_deployment.Environment, cfg DeployCREATE2FactoryCfg) error {
	return nil
}
