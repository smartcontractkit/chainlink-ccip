package sequences

import (
	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
)

func init() {
	v, err := semver.NewVersion("1.6.0")
	if err != nil {
		panic(err)
	}
	ccipapi.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
}

type EVMAdapter struct{}

func (a *EVMAdapter) GetOnRampAddress(e *cldf.Environment, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(onramp.ContractType),
		Version:       onramp.Version,
	}, chainSelector, evm_datastore_utils.ToPaddedEVMAddress)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *EVMAdapter) GetOffRampAddress(e *cldf.Environment, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(offramp.ContractType),
		Version:       offramp.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *EVMAdapter) GetFQAddress(e *cldf.Environment, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(fee_quoter.ContractType),
		Version:       fee_quoter.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *EVMAdapter) GetRouterAddress(e *cldf.Environment, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, err
	}
	return addr, nil
}
