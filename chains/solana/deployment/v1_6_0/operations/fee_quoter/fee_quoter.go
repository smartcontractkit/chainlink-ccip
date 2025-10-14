package fee_quoter

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/fee_quoter"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ContractType cldf_deployment.ContractType = "FeeQuoter"
var ProgramName = "fee_quoter"
var ProgramSize = 5 * 1024 * 1024
var Version *semver.Version = semver.MustParse("1.6.0")

type ConstructorArgs struct {
	StaticConfig                   fee_quoter.FeeQuoterStaticConfig
	PriceUpdaters                  []common.Address
	FeeTokens                      []common.Address
	TokenPriceFeedUpdates          []fee_quoter.FeeQuoterTokenPriceFeedUpdate
	TokenTransferFeeConfigArgs     []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs
	MorePremiumMultiplierWeiPerEth []fee_quoter.FeeQuoterPremiumMultiplierWeiPerEthArgs
	DestChainConfigArgs            []fee_quoter.FeeQuoterDestChainConfigArgs
}

var Deploy = operations.NewOperation(
	"fee-quoter:deploy",
	Version,
	"Deploys the FeeQuoter program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		// BEGIN Validation
		for _, ref := range input {
		if ref.Type == datastore.ContractType(input.TypeAndVersion.Type) {
			return ref, nil
		}
	}

		var (
			addr      common.Address
			tx        *types.Transaction
			deployErr error
		)
		if chain.IsZkSyncVM {
			addr, deployErr = deployZkContract(
				nil,
				bytecode.ZkSyncVM,
				chain.ClientZkSyncVM,
				chain.DeployerKeyZkSyncVM,
				parsedABI,
				args...,
			)
		} else {
			addr, tx, _, deployErr = deployEVMContract(
				chain.DeployerKey,
				*parsedABI,
				bytecode.EVM,
				chain.Client,
				args...,
			)
		}
		if !chain.IsZkSyncVM {
			// Non-ZkSyncVM chains require manual confirmation of the deployment transaction.
			// We attempt to decode any errors with the provided ABI.
			_, confirmErr := deployment.ConfirmIfNoErrorWithABI(chain, tx, params.ContractMetadata.ABI, deployErr)
			if confirmErr != nil {
				return datastore.AddressRef{}, fmt.Errorf("failed to deploy %s to %s with args %+v: %w", input.TypeAndVersion, chain, input.Args, confirmErr)
			}
			b.Logger.Debugw(fmt.Sprintf("Confirmed %s tx on %s", params.Name, chain), "hash", tx.Hash().Hex())
		} else if deployErr != nil {
			// For ZkSyncVM chains, if an error is returned from initial deployment, we return it here.
			return datastore.AddressRef{}, fmt.Errorf("failed to deploy %s to %s with args %+v: %w", input.TypeAndVersion, chain, input.Args, deployErr)
		}
		b.Logger.Debugw(fmt.Sprintf("Deployed %s to %s", input.TypeAndVersion, chain), "args", input.Args)

		return datastore.AddressRef{
			Address:       addr.Hex(),
			ChainSelector: input.ChainSelector,
			Type:          datastore.ContractType(input.TypeAndVersion.Type),
			Version:       &input.TypeAndVersion.Version,
		}, nil
	},
)

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

func DefaultFeeQuoterParams() FeeQuoterParams {
	return FeeQuoterParams{
		MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
		TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
		LinkPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
		WethPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
		TokenPriceFeedUpdates:          []fee_quoter.FeeQuoterTokenPriceFeedUpdate{},
		TokenTransferFeeConfigArgs:     []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs{},
		MorePremiumMultiplierWeiPerEth: []fee_quoter.FeeQuoterPremiumMultiplierWeiPerEthArgs{},
		DestChainConfigArgs:            []fee_quoter.FeeQuoterDestChainConfigArgs{},
	}
}
