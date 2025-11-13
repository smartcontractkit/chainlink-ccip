package create2_factory

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
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

type ConfigureCREATE2FactoryInput[CONTRACT any] struct {
	ChainSel         uint64
	CREATE2Factory  CONTRACT
	AllowListAdds    []common.Address
	AllowListRemoves []common.Address
}

func (c ConfigureCREATE2FactoryInput[CONTRACT]) ChainSelector() uint64 {
	return c.ChainSel
}

var ConfigureCREATE2Factory = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	ConfigureCREATE2FactoryInput[common.Address],
	evm.Chain,
	ConfigureCREATE2FactoryInput[datastore.AddressRef],
]{
	Sequence: configureCREATE2Factory,
	ResolveInput: func(e cldf_deployment.Environment, cfg ConfigureCREATE2FactoryInput[datastore.AddressRef]) (ConfigureCREATE2FactoryInput[common.Address], error) {
		create2FactoryAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.CREATE2Factory, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return ConfigureCREATE2FactoryInput[common.Address]{}, fmt.Errorf("failed to find create2 factory: %w", err)
		}
		return ConfigureCREATE2FactoryInput[common.Address]{
			ChainSel:         cfg.ChainSel,
			CREATE2Factory:  create2FactoryAddress,
			AllowListAdds:    cfg.AllowListAdds,
			AllowListRemoves: cfg.AllowListRemoves,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[ConfigureCREATE2FactoryInput[datastore.AddressRef]],
})

var configureCREATE2Factory = cldf_ops.NewSequence(
	"configure-create2-factory",
	semver.MustParse("1.7.0"),
	"Configures the CREATE2Factory contract",
	func(b operations.Bundle, chain evm.Chain, input ConfigureCREATE2FactoryInput[common.Address]) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)
		configureReport, err := cldf_ops.ExecuteOperation(b, create2_factory.ApplyAllowListUpdates, chain, contract.FunctionInput[create2_factory.ApplyAllowListUpdatesArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CREATE2Factory,
			Args: create2_factory.ApplyAllowListUpdatesArgs{
				Adds:    input.AllowListAdds,
				Removes: input.AllowListRemoves,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure CREATE2Factory: %w", err)
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
