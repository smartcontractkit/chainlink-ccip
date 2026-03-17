package changesets

import (
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type DeployAdvancedPoolHooksExtractorCfg struct {
	ChainSel uint64
}

func (c DeployAdvancedPoolHooksExtractorCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var DeployAdvancedPoolHooksExtractor = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.DeployAdvancedPoolHooksExtractorInput,
	evm.Chain,
	DeployAdvancedPoolHooksExtractorCfg,
]{
	Sequence: sequences.DeployAdvancedPoolHooksExtractor,
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployAdvancedPoolHooksExtractorCfg) (sequences.DeployAdvancedPoolHooksExtractorInput, error) {
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))
		return sequences.DeployAdvancedPoolHooksExtractorInput{
			ChainSelector:     cfg.ChainSel,
			ExistingAddresses: addresses,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[DeployAdvancedPoolHooksExtractorCfg],
})
