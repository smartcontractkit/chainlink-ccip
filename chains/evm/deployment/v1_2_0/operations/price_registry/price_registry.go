package price_registry

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/price_registry"
)

var (
	ContractType = cldf.ContractType("PriceRegistry")
	Version      = semver.MustParse("1.2.0")
)

type ApplyFeeTokensInput struct {
	FeeTokensToAdd    []common.Address
	FeeTokensToRemove []common.Address
}

func NewReadGetFeeToken(c gobindings.PriceRegistryInterface) *cld_ops.Operation[contract.FunctionInput[struct{}], []common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, []common.Address, gobindings.PriceRegistryInterface]{
		Name:         "price_registry:getfeetokens",
		Version:      Version,
		Description:  "gets fee token from price registry 1.2",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(pr gobindings.PriceRegistryInterface, opts *bind.CallOpts, args struct{}) ([]common.Address, error) {
			return pr.GetFeeTokens(opts)
		},
	})
}

func NewReadGetTokenPrices(c gobindings.PriceRegistryInterface) *cld_ops.Operation[contract.FunctionInput[[]common.Address], []gobindings.InternalTimestampedPackedUint224, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[[]common.Address, []gobindings.InternalTimestampedPackedUint224, gobindings.PriceRegistryInterface]{
		Name:         "price_registry:get-token-prices",
		Version:      Version,
		Description:  "Calls getTokenPrices on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(pr gobindings.PriceRegistryInterface, opts *bind.CallOpts, args []common.Address) ([]gobindings.InternalTimestampedPackedUint224, error) {
			return pr.GetTokenPrices(opts, args)
		},
	})
}

func NewReadGetDestinationChainGasPrice(c gobindings.PriceRegistryInterface) *cld_ops.Operation[contract.FunctionInput[uint64], gobindings.InternalTimestampedPackedUint224, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, gobindings.InternalTimestampedPackedUint224, gobindings.PriceRegistryInterface]{
		Name:         "price_registry:get-destination-chain-gas-price",
		Version:      Version,
		Description:  "Calls getDestinationChainGasPrice on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(pr gobindings.PriceRegistryInterface, opts *bind.CallOpts, args uint64) (gobindings.InternalTimestampedPackedUint224, error) {
			return pr.GetDestinationChainGasPrice(opts, args)
		},
	})
}

func NewWriteApplyFeeTokenUpdates(c gobindings.PriceRegistryInterface) *cld_ops.Operation[contract.FunctionInput[ApplyFeeTokensInput], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[ApplyFeeTokensInput, gobindings.PriceRegistryInterface]{
		Name:            "price_registry:apply-feetoken-updates",
		Version:         Version,
		Description:     "Calls applyFeeTokenUpdates on the contract",
		ContractType:    ContractType,
		ContractABI:     gobindings.PriceRegistryABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[gobindings.PriceRegistryInterface, ApplyFeeTokensInput],
		Validate:        func(ApplyFeeTokensInput) error { return nil },
		CallContract: func(pr gobindings.PriceRegistryInterface, opts *bind.TransactOpts, args ApplyFeeTokensInput) (*types.Transaction, error) {
			return pr.ApplyFeeTokensUpdates(opts, args.FeeTokensToAdd, args.FeeTokensToRemove)
		},
	})
}
