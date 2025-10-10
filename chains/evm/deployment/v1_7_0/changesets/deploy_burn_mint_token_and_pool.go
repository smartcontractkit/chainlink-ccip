package changesets

import (
	evm_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var DeployBurnMintTokenAndPool = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	tokens.DeployBurnMintTokenAndPoolInput,
	evm.Chain,
	tokens.DeployBurnMintTokenAndPoolInput,
]{
	Sequence: tokens.DeployBurnMintTokenAndPool,
	ResolveInput: func(e cldf_deployment.Environment, cfg tokens.DeployBurnMintTokenAndPoolInput) (tokens.DeployBurnMintTokenAndPoolInput, error) {
		return cfg, nil
	},
	ResolveDep: evm_seq.ResolveEVMChainDep[tokens.DeployBurnMintTokenAndPoolInput],
})
