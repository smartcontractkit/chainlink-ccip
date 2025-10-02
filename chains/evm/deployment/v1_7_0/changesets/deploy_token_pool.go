package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var DeployTokenPool = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	tokens.DeployTokenPoolInput,
	evm.Chain,
	tokens.DeployTokenPoolInput,
]{
	Sequence: tokens.DeployTokenPool,
	Describe: func(in tokens.DeployTokenPoolInput, dep evm.Chain) string {
		return fmt.Sprintf("Deploy %s %s for %s to %s", in.TokenPoolType, in.TokenPoolVersion, in.TokenSymbol, dep)
	},
	ResolveInput: func(e cldf_deployment.Environment, cfg tokens.DeployTokenPoolInput) (tokens.DeployTokenPoolInput, error) {
		return cfg, nil
	},
	ResolveDep: changesets.ResolveEVMChainDep[tokens.DeployTokenPoolInput],
})
