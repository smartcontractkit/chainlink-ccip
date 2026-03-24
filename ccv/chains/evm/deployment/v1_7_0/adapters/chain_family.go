package adapters

import (
	"encoding/binary"
	"math/big"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
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
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
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

func (a *ChainFamilyAdapter) GetChainFamilySelector() [4]byte {
	return utils.GetHexFromString(utils.EVMFamilySelector)
}

func (a *ChainFamilyAdapter) GetFeeQuoterDestChainConfig() lanes.FeeQuoterDestChainConfig {
	sel := a.GetChainFamilySelector()
	return lanes.FeeQuoterDestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                30_000,
		MaxPerMsgGasLimit:           3_000_000,
		DestGasOverhead:             300_000,
		DestGasPerPayloadByteBase:   16,
		ChainFamilySelector:         binary.BigEndian.Uint32(sel[:]),
		DefaultTokenFeeUSDCents:     25,
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		NetworkFeeUSDCents:          10,
		V2Params: &lanes.FeeQuoterV2Params{
			LinkFeeMultiplierPercent: 90,
			USDPerUnitGas:            big.NewInt(1e6),
		},
	}
}

func (a *ChainFamilyAdapter) GetDefaultGasPrice() *big.Int {
	return big.NewInt(2e12)
}
