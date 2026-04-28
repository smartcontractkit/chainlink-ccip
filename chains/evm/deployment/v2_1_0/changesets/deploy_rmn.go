package changesets

import (
	"github.com/ethereum/go-ethereum/common"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// DeployRMNCfg is configuration for the DeployRMN changeset (deploy RMN on one EVM chain).
type DeployRMNCfg struct {
	ChainSel    uint64
	CurseAdmins []common.Address
}

func (c DeployRMNCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var DeployRMN = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.DeployRMNInput,
	evm.Chain,
	DeployRMNCfg,
]{
	Sequence: sequences.DeployRMN,
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployRMNCfg) (sequences.DeployRMNInput, error) {
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))
		return sequences.DeployRMNInput{
			ChainSelector:     cfg.ChainSel,
			ExistingAddresses: addresses,
			Args: rmn.ConstructorArgs{
				CurseAdmins: cfg.CurseAdmins,
			},
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[DeployRMNCfg],
})
