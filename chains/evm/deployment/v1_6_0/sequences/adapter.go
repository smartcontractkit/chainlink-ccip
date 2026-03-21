package sequences

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	pingpongdapp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/ping_pong_dapp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

func init() {
	v := semver.MustParse("1.6.0")
	// Use a single EVMAdapter instance so the shared transferOwnershipAdapter state is used by
	// both InitializeTimelockAddress and SequenceTransferOwnershipViaMCMS.
	evmAdapter := &EVMAdapter{transferOwnershipAdapter: &evm1_0_0.EVMTransferOwnershipAdapter{}}
	ccipapi.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilyEVM, v, evmAdapter)
	ccipapi.GetPingPongAdapterRegistry().RegisterPingPongAdapter(chain_selectors.FamilyEVM, v, evmAdapter)
	deployops.GetRegistry().RegisterDeployer(chain_selectors.FamilyEVM, v, evmAdapter)
	deployops.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilyEVM, v, evmAdapter)
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilyEVM, v, evmAdapter)
	// 1.5.1 token pools use the same abstract TokenPool; use the 1.6.0 adapter for config/transfers.
	v151 := semver.MustParse("1.5.1")
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilyEVM, v151, evmAdapter)
}

type EVMAdapter struct {
	// transferOwnershipAdapter is shared so InitializeTimelockAddress populates the same instance
	// used by SequenceTransferOwnershipViaMCMS / SequenceAcceptOwnership.
	transferOwnershipAdapter *evm1_0_0.EVMTransferOwnershipAdapter
}

func (a *EVMAdapter) GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
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

func (a *EVMAdapter) GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
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

func (a *EVMAdapter) GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	// might be multiple fee quoters on a chain, just return the latest one
	refs := ds.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(fee_quoter.ContractType)),
		datastore.AddressRefByChainSelector(chainSelector),
	)
	ref, err := GetFeeQuoterAddress(refs, chainSelector, nil)
	if err != nil {
		return nil, err
	}
	return evm_datastore_utils.ToEVMAddressBytes(ref)
}

func (a *EVMAdapter) GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
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

func (a *EVMAdapter) GetTestRouter(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
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

// GetDefaultTokenPrices returns default fee token prices for EVM chains.
// Returns a map of contract type to USD price (18 decimals).
// The caller resolves contract types to addresses using the datastore.
func (a *EVMAdapter) GetDefaultTokenPrices() map[datastore.ContractType]*big.Int {
	// Default price: $20 per token (20 * 1e18)
	// Realistic price for LINK/WETH in test environments
	// Combined with low gas price (~$1 USD fee), this gives ~0.05 LINK per send
	defaultPrice := new(big.Int).Mul(big.NewInt(20), big.NewInt(1e18))

	return map[datastore.ContractType]*big.Int{
		datastore.ContractType(link.ContractType): defaultPrice,
		datastore.ContractType(weth.ContractType): defaultPrice,
	}
}

func (a *EVMAdapter) GetPingPongDemoAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(pingpongdapp.ContractType),
		Version:       pingpongdapp.Version,
	}, chainSelector, evm_datastore_utils.ToEVMAddressBytes)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// ConfigurePingPong returns a sequence that configures PingPong for a lane.
func (a *EVMAdapter) ConfigurePingPong() *operations.Sequence[ccipapi.PingPongInput, ccipapi.PingPongOutput, cldf_chain.BlockChains] {
	return ConfigurePingPongSequence
}

// ConfigurePingPongSequence is the sequence for configuring PingPong between two EVM chains.
var ConfigurePingPongSequence = operations.NewSequence(
	"ConfigurePingPong",
	semver.MustParse("1.0.0"),
	"Configures PingPong counterpart for a lane between two chains",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input ccipapi.PingPongInput) (ccipapi.PingPongOutput, error) {
		b.Logger.Infof("EVM Configuring PingPong counterpart. src: %d, dest: %d", input.SourceSelector, input.DestSelector)

		chain := chains.EVMChains()[input.SourceSelector]

		// CCIP requires addresses to be 32-byte left-padded for cross-chain messaging
		paddedDestAddr := common.LeftPadBytes(input.DestPingPongAddr, 32)

		_, err := operations.ExecuteOperation(b, pingpongdapp.SetCounterpart, chain, contract.FunctionInput[pingpongdapp.SetCounterpartArgs]{
			ChainSelector: input.SourceSelector,
			Address:       common.BytesToAddress(input.SourcePingPongAddr),
			Args: pingpongdapp.SetCounterpartArgs{
				CounterpartChainSelector: input.DestSelector,
				CounterpartAddress:       paddedDestAddr,
			},
		})
		if err != nil {
			return ccipapi.PingPongOutput{}, err
		}

		b.Logger.Infof("PingPong counterpart set for lane %d -> %d", input.SourceSelector, input.DestSelector)
		return ccipapi.PingPongOutput{}, nil
	},
)

// GetFeeQuoterAddress returns the address of the fee quoter contract for a given chain selector.
// there may be multiple fee quoter addresses for a chain selector, so we return the one with the latest version
// (the next major version on or after 1.6.0).
func GetFeeQuoterAddress(addresses []datastore.AddressRef, chainSelector uint64, tooHighVersion *semver.Version) (datastore.AddressRef, error) {
	var refs []datastore.AddressRef
	for _, ref := range addresses {
		if ref.ChainSelector == chainSelector &&
			ref.Type == datastore.ContractType(fee_quoter.ContractType) {
			refs = append(refs, ref)
		}
	}
	latestVersion := semver.MustParse("1.6.0")
	feeQRef := datastore.AddressRef{}
	for _, ref := range refs {
		v := ref.Version
		if tooHighVersion != nil && v.GreaterThanEqual(tooHighVersion) {
			continue
		}
		if v.GreaterThanEqual(latestVersion) {
			latestVersion = v
			feeQRef = ref
		}
	}
	if feeQRef.Address == "" {
		return datastore.AddressRef{}, fmt.Errorf("no fee quoter address found for chain selector %d", chainSelector)
	}
	return feeQRef, nil
}

func (a *EVMAdapter) GetFeeQuoterDestChainConfig() ccipapi.FeeQuoterDestChainConfig {
	chainHex := utils.GetHexFromString(utils.EVMFamilySelector)
	return ccipapi.FeeQuoterDestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                30_000,
		MaxPerMsgGasLimit:           3_000_000,
		DestGasOverhead:             300_000,
		DestGasPerPayloadByteBase:   16,
		ChainFamilySelector:         binary.BigEndian.Uint32(chainHex[:]),
		DefaultTokenFeeUSDCents:     25,
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		NetworkFeeUSDCents:          10,
		V1Params: &ccipapi.FeeQuoterV1Params{
			MaxNumberOfTokensPerMsg:           10,
			DestGasPerPayloadByteHigh:         40,
			DestGasPerPayloadByteThreshold:    3000,
			DestDataAvailabilityOverheadGas:   100,
			DestGasPerDataAvailabilityByte:    16,
			DestDataAvailabilityMultiplierBps: 1,
			GasMultiplierWeiPerEth:            11e17,
		},
		V2Params: &ccipapi.FeeQuoterV2Params{
			LinkFeeMultiplierPercent: 90,
			USDPerUnitGas:            big.NewInt(1e6),
		},
	}
}

func (a *EVMAdapter) GetDefaultGasPrice() *big.Int {
	return big.NewInt(2e12)
}

// GetFQVersion implements the optional FeeQuoterVersionProvider interface so that
// update_lanes can choose 1.6 vs 2.0 FeeQuoter operations based on the deployed contract version.
func (a *EVMAdapter) GetFQVersion(ds datastore.DataStore, address []byte, chainSelector uint64) (*semver.Version, error) {
	refs := ds.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(fee_quoter.ContractType)),
		datastore.AddressRefByChainSelector(chainSelector),
	)
	ref, err := GetFeeQuoterAddress(refs, chainSelector, nil)
	if err != nil {
		return nil, err
	}
	// Sanity check that the AddressRef we found matches the one we expect.
	if ref.Address != common.BytesToAddress(address).Hex() {
		return nil, fmt.Errorf("fee quoter address mismatch for chain selector %d: expected %s, got %s", chainSelector, common.BytesToAddress(address).Hex(), ref.Address)
	}
	return ref.Version, nil
}
