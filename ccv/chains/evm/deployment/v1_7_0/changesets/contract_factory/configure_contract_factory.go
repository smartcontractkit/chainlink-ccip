package contract_factory

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/contract_factory"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ConfigureContractFactoryInput[CONTRACT any] struct {
	ChainSel         uint64
	ContractFactory  CONTRACT
	AllowListAdds    []common.Address
	AllowListRemoves []common.Address
}

func (c ConfigureContractFactoryInput[CONTRACT]) ChainSelector() uint64 {
	return c.ChainSel
}

var ConfigureContractFactory = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	ConfigureContractFactoryInput[common.Address],
	evm.Chain,
	ConfigureContractFactoryInput[datastore.AddressRef],
]{
	Sequence: configureContractFactory,
	ResolveInput: func(e cldf_deployment.Environment, cfg ConfigureContractFactoryInput[datastore.AddressRef]) (ConfigureContractFactoryInput[common.Address], error) {
		contractFactoryAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.ContractFactory, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return ConfigureContractFactoryInput[common.Address]{}, fmt.Errorf("failed to find contract factory: %w", err)
		}
		return ConfigureContractFactoryInput[common.Address]{
			ChainSel:         cfg.ChainSel,
			ContractFactory:  contractFactoryAddress,
			AllowListAdds:    cfg.AllowListAdds,
			AllowListRemoves: cfg.AllowListRemoves,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[ConfigureContractFactoryInput[datastore.AddressRef]],
})

var configureContractFactory = cldf_ops.NewSequence(
	"configure-contract-factory",
	semver.MustParse("1.7.0"),
	"Configures the ContractFactory contract",
	func(b operations.Bundle, chain evm.Chain, input ConfigureContractFactoryInput[common.Address]) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)
		configureReport, err := cldf_ops.ExecuteOperation(b, contract_factory.ApplyAllowListUpdates, chain, contract.FunctionInput[contract_factory.ApplyAllowListUpdatesArgs]{
			ChainSelector: chain.Selector,
			Address:       input.ContractFactory,
			Args: contract_factory.ApplyAllowListUpdatesArgs{
				Adds:    input.AllowListAdds,
				Removes: input.AllowListRemoves,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure ContractFactory: %w", err)
		}
		writes = append(writes, configureReport.Output)

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	},
)
