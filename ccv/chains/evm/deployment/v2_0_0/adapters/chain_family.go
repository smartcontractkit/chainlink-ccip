package adapters

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/sequences"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// ChainFamilyAdapter is the adapter for chains of the EVM family.
type ChainFamilyAdapter struct{}

var _ ccvadapters.ChainFamily = (*ChainFamilyAdapter)(nil)

// ConfigureLaneLegAsSource returns the sequence for configuring a chain of the EVM family as a source chain for CCIP lanes.
func (a *ChainFamilyAdapter) ConfigureLaneLegAsSource() *operations.Sequence[lanes.UpdateLanesInput, seq_core.OnChainOutput, cldf_chain.BlockChains] {
	return sequences.ConfigureLaneLegAsSource
}

// ConfigureLaneLegAsDest returns the sequence for configuring a chain of the EVM family as a destination chain for CCIP lanes.
func (a *ChainFamilyAdapter) ConfigureLaneLegAsDest() *operations.Sequence[lanes.UpdateLanesInput, seq_core.OnChainOutput, cldf_chain.BlockChains] {
	return sequences.ConfigureLaneLegAsDest
}

// ConfigureChainForLanes returns the sequence for configuring an EVM chain for multiple remote lanes.
func (a *ChainFamilyAdapter) ConfigureChainForLanes() *operations.Sequence[ccvadapters.ConfigureChainForLanesInput, seq_core.OnChainOutput, cldf_chain.BlockChains] {
	return sequences.ConfigureChainForLanes
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

	onrampContract, err := onramp.NewOnRampContract(common.BytesToAddress(onRampAddr), chain.Client)
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

func (a *ChainFamilyAdapter) DisableRemoteChain() *operations.Sequence[lanes.DisableRemoteChainInput, seq_core.OnChainOutput, cldf_chain.BlockChains] {
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

func (a *ChainFamilyAdapter) ResolveExecutor(ds datastore.DataStore, chainSelector uint64, qualifier string) (string, error) {
	toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(sequences.ExecutorProxyType),
		Qualifier: qualifier,
		Version:   executor.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return "", fmt.Errorf("failed to resolve executor proxy (qualifier %q) on chain %d: %w", qualifier, chainSelector, err)
	}
	if !common.IsHexAddress(addr) {
		return "", fmt.Errorf("resolved executor proxy address %q is not a valid hex address (qualifier %q, chain %d)", addr, qualifier, chainSelector)
	}
	if common.HexToAddress(addr) == (common.Address{}) {
		return "", fmt.Errorf("resolved executor proxy address is zero (qualifier %q, chain %d)", qualifier, chainSelector)
	}
	return addr, nil
}

// AddressRefToBytes returns the byte representation of an EVM address ref.
// It validates the hex string and rejects zero addresses.
func (a *ChainFamilyAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	if !common.IsHexAddress(ref.Address) {
		return nil, fmt.Errorf("invalid EVM address %q in ref %s/%s", ref.Address, ref.Type, ref.Version)
	}
	addr := common.HexToAddress(ref.Address)
	if addr == (common.Address{}) {
		return nil, fmt.Errorf("zero address in ref %s/%s", ref.Type, ref.Version)
	}
	return addr.Bytes(), nil
}

// evmFamilySelector is bytes4(keccak256("CCIP ChainFamilySelector EVM")) = 0x2812d52c.
var evmFamilySelector = [4]byte{0x28, 0x12, 0xd5, 0x2c}

func (a *ChainFamilyAdapter) GetChainFamilySelector() [4]byte {
	return evmFamilySelector
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
