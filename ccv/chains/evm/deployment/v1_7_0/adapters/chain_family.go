package adapters

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// ChainFamilyAdapter is the adapter for chains of the EVM family.
type ChainFamilyAdapter struct{}

// ConfigureLaneLegAsSource returns the sequence for configuring a chain of the EVM family as a source chain for CCIP lanes.
func (c *ChainFamilyAdapter) ConfigureLaneLegAsSource() *operations.Sequence[lanes.UpdateLanesInput, seq_core.OnChainOutput, chain.BlockChains] {
	return sequences.ConfigureLaneLegAsSource
}

// ConfigureLaneLegAsDest returns the sequence for configuring a chain of the EVM family as a destination chain for CCIP lanes.
func (c *ChainFamilyAdapter) ConfigureLaneLegAsDest() *operations.Sequence[lanes.UpdateLanesInput, seq_core.OnChainOutput, chain.BlockChains] {
	return sequences.ConfigureLaneLegAsDest
}

func (a *ChainFamilyAdapter) GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(onramp.ContractType),
		Version:       onramp.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *ChainFamilyAdapter) GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(offramp.ContractType),
		Version:       offramp.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *ChainFamilyAdapter) GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(fee_quoter.ContractType),
		Version:       fee_quoter.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *ChainFamilyAdapter) GetFQAddressDynamic(ds datastore.DataStore, chainSelector uint64, chains cldf_chain.BlockChains) ([]byte, error) {
	onRampAddr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(onramp.ContractType),
		Version:       onramp.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to find onramp address for chain selector %d: %w", chainSelector, err)
	}

	chain := chains.EVMChains()[chainSelector]

	onrampContract, err := onramp.NewOnRampContract(common.Address(onRampAddr), chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create onramp contract instance for chain selector %d: %w", chainSelector, err)
	}

	dynamicConfig, err := onrampContract.GetDynamicConfig(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to call GetDynamicConfig on onramp contract for chain selector %d: %w", chainSelector, err)
	}

	fqAddress := dynamicConfig.FeeQuoter
	if fqAddress == (common.Address{}) {
		return nil, fmt.Errorf("fee quoter address is zero in onramp dynamic config for chain selector %d", chainSelector)
	}
	return common.Address(fqAddress).Bytes(), nil
}

func (c *ChainFamilyAdapter) DisableRemoteChain() *operations.Sequence[lanes.DisableRemoteChainInput, seq_core.OnChainOutput, chain.BlockChains] {
	return evm_sequences.DisableRemoteChainSequence
}

func (a *ChainFamilyAdapter) GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *ChainFamilyAdapter) GetTestRouter(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(router.TestRouterContractType),
		Version:       router.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *ChainFamilyAdapter) GetFeeQuoterDestChainConfig() lanes.FeeQuoterDestChainConfig {
	return lanes.DefaultFeeQuoterDestChainConfig(false, chainsel.ETHEREUM_MAINNET.Selector)
}

func (a *ChainFamilyAdapter) GetDefaultGasPrice() *big.Int {
	return big.NewInt(2e12)
}
