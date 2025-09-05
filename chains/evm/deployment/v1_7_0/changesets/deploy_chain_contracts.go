package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type DeployChainContractsCfg struct {
	ChainSel uint64
	Params   sequences.ContractParams
}

func (c DeployChainContractsCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var DeployChainContracts = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.DeployChainContractsInput,
	evm.Chain,
	DeployChainContractsCfg,
]{
	Sequence: sequences.DeployChainContracts,
	Describe: func(in sequences.DeployChainContractsInput, dep evm.Chain) string {
		return fmt.Sprintf("Deploy & configure 1.7.0 contracts on %s", dep)
	},
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployChainContractsCfg) (sequences.DeployChainContractsInput, error) {
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))
		return sequences.DeployChainContractsInput{
			ChainSelector:     cfg.ChainSel,
			ExistingAddresses: addresses,
			ContractParams:    cfg.Params,
		}, nil
	},
	ResolveDep: changesets.ResolveEVMChainDep[DeployChainContractsCfg],
})
