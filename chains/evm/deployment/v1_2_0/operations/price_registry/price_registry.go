package price_registry

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/price_registry"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var (
	ContractType = cldf.ContractType("PriceRegistry")
	Version      = semver.MustParse("1.2.0")
)

var PriceRegistryGetFeeToken = contract.NewRead(contract.ReadParams[any, []common.Address, *price_registry.PriceRegistry]{
	Name:         "price_registry:getfeetokens",
	Version:      Version,
	Description:  "gets fee token from price registry 1.2",
	ContractType: ContractType,
	NewContract:  price_registry.NewPriceRegistry,
	CallContract: func(pr *price_registry.PriceRegistry, opts *bind.CallOpts, args any) ([]common.Address, error) {
		return pr.GetFeeTokens(opts)
	},
})

var PriceRegistryGetTokenPrices = contract.NewRead(contract.ReadParams[[]common.Address, []price_registry.InternalTimestampedPackedUint224, *price_registry.PriceRegistry]{
	Name:         "price_registry:get-token-prices",
	Version:      Version,
	Description:  "Calls getTokenPrices on the contract",
	ContractType: ContractType,
	NewContract:  price_registry.NewPriceRegistry,
	CallContract: func(pr *price_registry.PriceRegistry, opts *bind.CallOpts, args []common.Address) ([]price_registry.InternalTimestampedPackedUint224, error) {
		return pr.GetTokenPrices(opts, args)
	},
})

var PriceRegistryGetDestinationChainGasPrice = contract.NewRead(contract.ReadParams[uint64, price_registry.InternalTimestampedPackedUint224, *price_registry.PriceRegistry]{
	Name:         "price_registry:get-destination-chain-gas-price",
	Version:      Version,
	Description:  "Calls getDestinationChainGasPrice on the contract",
	ContractType: ContractType,
	NewContract:  price_registry.NewPriceRegistry,
	CallContract: func(pr *price_registry.PriceRegistry, opts *bind.CallOpts, args uint64) (price_registry.InternalTimestampedPackedUint224, error) {
		return pr.GetDestinationChainGasPrice(opts, args)
	},
})
