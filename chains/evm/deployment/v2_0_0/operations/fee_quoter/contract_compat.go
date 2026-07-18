package fee_quoter

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"
)

// FeeQuoterContract preserves the pre-ops-v2 binding wrapper used by chainlink deployment.
type FeeQuoterContract struct {
	*gobindings.FeeQuoter
	address common.Address
}

func NewFeeQuoterContract(address common.Address, backend bind.ContractBackend) (*FeeQuoterContract, error) {
	fq, err := gobindings.NewFeeQuoter(address, backend)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterContract{FeeQuoter: fq, address: address}, nil
}

func (c *FeeQuoterContract) Address() common.Address {
	return c.address
}

func (c *FeeQuoterContract) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
	return c.FeeQuoter.ApplyDestChainConfigUpdates(opts, args)
}

func (c *FeeQuoterContract) UpdatePrices(opts *bind.TransactOpts, args PriceUpdates) (*types.Transaction, error) {
	return c.FeeQuoter.UpdatePrices(opts, args)
}

func (c *FeeQuoterContract) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
	return c.FeeQuoter.ApplyAuthorizedCallerUpdates(opts, args)
}

func (c *FeeQuoterContract) ApplyTokenTransferFeeConfigUpdates(
	opts *bind.TransactOpts,
	tokenTransferFeeConfigArgs []TokenTransferFeeConfigArgs,
	tokensToUseDefaultFeeConfigs []TokenTransferFeeConfigRemoveArgs,
) (*types.Transaction, error) {
	return c.FeeQuoter.ApplyTokenTransferFeeConfigUpdates(opts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}
