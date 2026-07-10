package create2_factory

import (
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	create2_factory_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
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

	tv := deployment.NewTypeAndVersion(create2_factory.ContractType, *create2_factory.Version)

	// The CREATE2Factory is expected to be the very first contract deployed by the
	// deployer key, so it is deployed at the address derived from the deployer key
	// and nonce 0. This makes the factory address deterministic across chains.
	nonce, err := evmChain.Client.PendingNonceAt(e.GetContext(), evmChain.DeployerKey.From)
	if err != nil {
		return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to get nonce of deployer key: %w", err)
	}

	// If the deployer key has already sent transactions, the factory may already have
	// been deployed by a previous run of this changeset. Rather than failing, look up
	// the contract that was deployed at nonce 0 and, if it is the CREATE2Factory,
	// record its address in the datastore instead of deploying a new one.
	if nonce != 0 {
		factoryAddr := crypto.CreateAddress(evmChain.DeployerKey.From, 0)

		code, err := evmChain.Client.CodeAt(e.GetContext(), factoryAddr, nil)
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to get code at expected create2 factory address %s: %w", factoryAddr, err)
		}
		if len(code) == 0 {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("deployer key nonce is %d but no contract is deployed at the expected create2 factory address %s", nonce, factoryAddr)
		}

		factory, err := create2_factory_bindings.NewCREATE2Factory(factoryAddr, evmChain.Client)
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to bind create2 factory at %s: %w", factoryAddr, err)
		}
		gotTV, err := factory.TypeAndVersion(nil)
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to get type and version of contract at %s: %w", factoryAddr, err)
		}
		if gotTV != tv.String() {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("contract at expected create2 factory address %s is not a %s, got %q", factoryAddr, tv, gotTV)
		}

		ds := datastore.NewMemoryDataStore()
		if err := ds.Addresses().Add(datastore.AddressRef{
			Address:       factoryAddr.Hex(),
			ChainSelector: cfg.ChainSel,
			Type:          datastore.ContractType(tv.Type),
			Version:       &tv.Version,
		}); err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to add existing create2 factory address to datastore: %w", err)
		}

		return cldf_deployment.ChangesetOutput{
			DataStore: ds,
		}, nil
	}

	deployReport, err := evmops.ExecuteDeploy(e.OperationsBundle, create2_factory.Deploy, evmChain, contract.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: tv,
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
