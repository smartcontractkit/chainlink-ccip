package burn_mint_erc677

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = cciputils.BurnMintToken

const burnMintERC677ABI = `[{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"burnAndMinter","type":"address"}],"name":"grantMintAndBurnRoles","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

type burnMintERC677 struct {
	address  common.Address
	contract *bind.BoundContract
}

func newBurnMintERC677(address common.Address, backend bind.ContractBackend) (*burnMintERC677, error) {
	parsed, err := abi.JSON(strings.NewReader(burnMintERC677ABI))
	if err != nil {
		return nil, err
	}

	return &burnMintERC677{
		address:  address,
		contract: bind.NewBoundContract(address, parsed, backend, backend, backend),
	}, nil
}

func (token *burnMintERC677) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := token.contract.Call(opts, &out, "owner")
	if err != nil {
		return common.Address{}, err
	}

	owner := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return owner, nil
}

func (token *burnMintERC677) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return token.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

var GrantMintAndBurnRoles = contract.NewWrite(contract.WriteParams[common.Address, *burnMintERC677]{
	Name:         "burn_mint_erc677:grant-mint-and-burn-roles",
	Version:      cciputils.Version_1_0_0,
	Description:  "Grant mint and burn role on BurnMintERC677 token contract",
	ContractType: ContractType,
	ContractABI:  burnMintERC677ABI,
	NewContract:  newBurnMintERC677,
	IsAllowedCaller: func(token *burnMintERC677, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		owner, err := token.Owner(opts)
		if err != nil {
			return false, err
		}
		return owner == caller, nil
	},
	Validate: func(address common.Address) error { return nil },
	CallContract: func(token *burnMintERC677, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return token.GrantMintAndBurnRoles(opts, input)
	},
})
