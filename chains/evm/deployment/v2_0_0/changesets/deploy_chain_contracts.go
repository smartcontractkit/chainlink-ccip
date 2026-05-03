package changesets

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	seq_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type DeployChainContractsCfg struct {
	ChainSel         uint64
	CREATE2Factory   common.Address
	Params           sequences.ContractParams
	DeployTestRouter bool
	DeployerKeyOwned bool
}

func (c DeployChainContractsCfg) ChainSelector() uint64 {
	return c.ChainSel
}

// wrappedDeployChainContracts adapts the sequence output to OnChainOutput so the
// generic NewFromOnChainSequence helper can continue to be used.
var wrappedDeployChainContracts = cldf_ops.NewSequence(
	"wrapped-deploy-chain-contracts",
	semver.MustParse("2.0.0"),
	"Wraps DeployChainContracts and returns the base on-chain output",
	func(
		b cldf_ops.Bundle,
		chain evm.Chain,
		input sequences.DeployChainContractsInput,
	) (seq_utils.OnChainOutput, error) {
		report, err := cldf_ops.ExecuteSequence(b, sequences.DeployChainContracts, chain, input)
		if err != nil {
			return seq_utils.OnChainOutput{}, err
		}
		return report.Output.OnChainOutput, nil
	},
)

var DeployChainContracts = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.DeployChainContractsInput,
	evm.Chain,
	DeployChainContractsCfg,
]{
	Sequence: wrappedDeployChainContracts,
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployChainContractsCfg) (sequences.DeployChainContractsInput, error) {
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))
		return sequences.DeployChainContractsInput{
			CREATE2Factory:    cfg.CREATE2Factory,
			ChainSelector:     cfg.ChainSel,
			ExistingAddresses: addresses,
			ContractParams:    cfg.Params,
			DeployTestRouter:  cfg.DeployTestRouter,
			DeployerKeyOwned:  cfg.DeployerKeyOwned,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[DeployChainContractsCfg],
})
