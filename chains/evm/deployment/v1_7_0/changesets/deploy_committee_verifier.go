package changesets

import (
	"fmt"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type DeployCommitteeVerifierCfg struct {
	ChainSel  uint64
	Params    sequences.CommitteeVerifierParams
	FeeQuoter datastore.AddressRef
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
		feeQuoter, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.FeeQuoter, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.DeployCommitteeVerifierInput{}, fmt.Errorf("failed to resolve fee quoter ref: %w", err)
		}
		return sequences.DeployCommitteeVerifierInput{
			ChainSelector:     cfg.ChainSel,
			ExistingAddresses: addresses,
			Params:            cfg.Params,
			FeeQuoter:         feeQuoter,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[DeployCommitteeVerifierCfg],
})
