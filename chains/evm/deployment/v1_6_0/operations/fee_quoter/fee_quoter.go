package fee_quoter

import (
	"errors"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "FeeQuoter"
var Version *semver.Version = semver.MustParse("1.6.3")

type ConstructorArgs struct {
	StaticConfig                   fee_quoter.FeeQuoterStaticConfig
	PriceUpdaters                  []common.Address
	FeeTokens                      []common.Address
	TokenPriceFeedUpdates          []fee_quoter.FeeQuoterTokenPriceFeedUpdate
	TokenTransferFeeConfigArgs     []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs
	MorePremiumMultiplierWeiPerEth []fee_quoter.FeeQuoterPremiumMultiplierWeiPerEthArgs
	DestChainConfigArgs            []fee_quoter.FeeQuoterDestChainConfigArgs
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "fee-quoter:deploy",
	Version:          Version,
	Description:      "Deploys the FeeQuoter contract",
	ContractMetadata: fee_quoter.FeeQuoterMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(fee_quoter.FeeQuoterBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var FeeQuoterApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]fee_quoter.FeeQuoterDestChainConfigArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:apply-dest-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to destination chain configs on the FeeQuoter 1.6.0 contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, []fee_quoter.FeeQuoterDestChainConfigArgs],
	Validate:        func([]fee_quoter.FeeQuoterDestChainConfigArgs) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args []fee_quoter.FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyDestChainConfigUpdates(opts, args)
	},
})

var FeeQuoterUpdatePrices = contract.NewWrite(contract.WriteParams[fee_quoter.InternalPriceUpdates, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:update-prices",
	Version:         Version,
	Description:     "Updates prices on the FeeQuoter 1.6.0 contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, fee_quoter.InternalPriceUpdates],
	Validate:        func(fee_quoter.InternalPriceUpdates) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args fee_quoter.InternalPriceUpdates) (*types.Transaction, error) {
		return feeQuoter.UpdatePrices(opts, args)
	},
})

type FeeQuoterParams struct {
	MaxFeeJuelsPerMsg              *big.Int
	TokenPriceStalenessThreshold   uint32
	LinkPremiumMultiplierWeiPerEth uint64
	WethPremiumMultiplierWeiPerEth uint64
	MorePremiumMultiplierWeiPerEth []fee_quoter.FeeQuoterPremiumMultiplierWeiPerEthArgs
	TokenPriceFeedUpdates          []fee_quoter.FeeQuoterTokenPriceFeedUpdate
	TokenTransferFeeConfigArgs     []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs
	DestChainConfigArgs            []fee_quoter.FeeQuoterDestChainConfigArgs
}

func (c FeeQuoterParams) Validate() error {
	if c.MaxFeeJuelsPerMsg == nil {
		return errors.New("MaxFeeJuelsPerMsg is nil")
	}
	if c.MaxFeeJuelsPerMsg.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("MaxFeeJuelsPerMsg must be positive")
	}
	if c.TokenPriceStalenessThreshold == 0 {
		return errors.New("TokenPriceStalenessThreshold can't be 0")
	}
	return nil
}
