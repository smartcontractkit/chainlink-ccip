package changesets

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type DeployChainContractsCfg struct {
	ChainSel        uint64
	ContractFactory common.Address
	Params          sequences.ContractParams
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
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployChainContractsCfg) (sequences.DeployChainContractsInput, error) {
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))
		return sequences.DeployChainContractsInput{
			ContractFactory:   cfg.ContractFactory,
			ChainSelector:     cfg.ChainSel,
			ExistingAddresses: addresses,
			ContractParams:    cfg.Params,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[DeployChainContractsCfg],
})
