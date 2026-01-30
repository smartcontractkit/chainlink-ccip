package sequences

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	pingpongdapp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/ping_pong_dapp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

func init() {
	v, err := semver.NewVersion("1.6.0")
	if err != nil {
		panic(err)
	}
	ccipapi.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
	ccipapi.GetPingPongAdapterRegistry().RegisterPingPongAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
	deployops.GetRegistry().RegisterDeployer(chain_selectors.FamilyEVM, v, &EVMAdapter{})
	deployops.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
}

type EVMAdapter struct{}

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
	latestVersion := semver.MustParse("1.6.0")
	tooHighVersion := semver.MustParse("1.7.0")
	var addr []byte
	var err error
	for _, ref := range refs {
		v := ref.Version
		// we want the latest version below 1.7.0
		if v.GreaterThanEqual(latestVersion) &&
			v.LessThan(tooHighVersion) {
			latestVersion = v
			addr, err = evm_datastore_utils.ToEVMAddressBytes(ref)
			if err != nil {
				return nil, err
			}
		}

	}
	return addr, nil
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
