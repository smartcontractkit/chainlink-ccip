package changesets

import (
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
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
	solana.Chain,
	DeployChainContractsCfg,
]{
	Sequence: sequences.DeployChainContracts,
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployChainContractsCfg) (sequences.DeployChainContractsInput, error) {
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))
		return sequences.DeployChainContractsInput{
			ChainSelector:     cfg.ChainSel,
			ExistingAddresses: addresses,
			ContractParams:    cfg.Params,
		}, nil
	},
	ResolveDep: utils.ResolveSolanaChainDep[DeployChainContractsCfg],
})
