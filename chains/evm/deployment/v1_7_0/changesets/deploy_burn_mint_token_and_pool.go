package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var DeployBurnMintTokenAndPool = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	tokens.DeployBurnMintTokenAndPoolInput,
	evm.Chain,
	tokens.DeployBurnMintTokenAndPoolInput,
]{
	Sequence: tokens.DeployBurnMintTokenAndPool,
	Describe: func(in tokens.DeployBurnMintTokenAndPoolInput, dep evm.Chain) string {
		return fmt.Sprintf("Deploy token and %s %s for %s to %s", in.DeployTokenPoolInput.TokenPoolType, in.DeployTokenPoolInput.TokenPoolVersion, in.DeployTokenPoolInput.TokenSymbol, dep)
	},
	ResolveInput: func(e cldf_deployment.Environment, cfg tokens.DeployBurnMintTokenAndPoolInput) (tokens.DeployBurnMintTokenAndPoolInput, error) {
		return cfg, nil
	},
	ResolveDep: changesets.ResolveEVMChainDep[tokens.DeployBurnMintTokenAndPoolInput],
})
