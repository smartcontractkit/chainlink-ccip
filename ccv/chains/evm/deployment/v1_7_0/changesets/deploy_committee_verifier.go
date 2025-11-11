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

type DeployCommitteeVerifierCfg struct {
	ChainSel        uint64
	ContractFactory common.Address
	Params          sequences.CommitteeVerifierParams
}

func (c DeployCommitteeVerifierCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var DeployCommitteeVerifier = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.DeployCommitteeVerifierInput,
	evm.Chain,
	DeployCommitteeVerifierCfg,
]{
	Sequence: sequences.DeployCommitteeVerifier,
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployCommitteeVerifierCfg) (sequences.DeployCommitteeVerifierInput, error) {
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))
		return sequences.DeployCommitteeVerifierInput{
			ChainSelector:     cfg.ChainSel,
			ExistingAddresses: addresses,
			Params:            cfg.Params,
			ContractFactory:   cfg.ContractFactory,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[DeployCommitteeVerifierCfg],
})
