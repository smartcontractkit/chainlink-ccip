package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/contract_factory"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type ConfigureContractFactoryCfg struct {
	ChainSel         uint64
	AllowListAdds    []common.Address
	AllowListRemoves []common.Address
	ContractFactory  datastore.AddressRef
}

func (c ConfigureContractFactoryCfg) ChainSelector() uint64 {
	return c.ChainSel
}

var ConfigureContractFactory = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	contract_factory.ConfigureContractFactoryInput,
	evm.Chain,
	ConfigureContractFactoryCfg,
]{
	Sequence: contract_factory.ConfigureContractFactory,
	ResolveInput: func(e cldf_deployment.Environment, cfg ConfigureContractFactoryCfg) (contract_factory.ConfigureContractFactoryInput, error) {
		contractFactoryAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.ContractFactory, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return contract_factory.ConfigureContractFactoryInput{}, fmt.Errorf("failed to find contract factory address: %w", err)
		}
		return contract_factory.ConfigureContractFactoryInput{
			ChainSelector:          cfg.ChainSel,
			ContractFactoryAddress: contractFactoryAddress,
			AllowListAdds:          cfg.AllowListAdds,
			AllowListRemoves:       cfg.AllowListRemoves,
		}, nil
	},
})
