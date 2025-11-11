package changesets

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/contract_factory"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type DeployContractFactoryCfg struct {
	ChainSel  uint64
	AllowList []common.Address
}

func (c DeployContractFactoryCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var DeployContractFactory = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	contract_factory.DeployContractFactoryInput,
	evm.Chain,
	DeployContractFactoryCfg,
]{
	Sequence: contract_factory.DeployContractFactory,
	ResolveInput: func(e cldf_deployment.Environment, cfg DeployContractFactoryCfg) (contract_factory.DeployContractFactoryInput, error) {
		return contract_factory.DeployContractFactoryInput{
			ChainSelector: cfg.ChainSel,
			AllowList:     cfg.AllowList,
		}, nil
	},
})
