package burn_mint_erc20_with_drip

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"
)

var (
	ContractType cldf_deployment.ContractType = "BurnMintERC20WithDrip"
	Version                                   = semver.MustParse("1.5.0")

	DefaultAdminRole = [32]byte{}
	BurnRole         = [32]byte(crypto.Keccak256([]byte("BURNER_ROLE")))
	MintRole         = [32]byte(crypto.Keccak256([]byte("MINTER_ROLE")))
)

type ConstructorArgs struct {
	Name   string
	Symbol string
}

type MintArgs struct {
	Account common.Address
	Amount  *big.Int
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "burn-mint-erc20-with-drip:deploy",
	Version:          Version,
	Description:      "Deploys the BurnMintERC20WithDrip contract",
	ContractMetadata: burn_mint_erc20_with_drip.BurnMintERC20WithDripMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(burn_mint_erc20_with_drip.BurnMintERC20WithDripBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

func NewWriteGrantMintAndBurnRoles(c *burn_mint_erc20_with_drip.BurnMintERC20WithDrip) *cld_ops.Operation[contract.FunctionInput[common.Address], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc20_with_drip.BurnMintERC20WithDrip]{
		Name:         "burn-mint-erc20-with-drip:grant-mint-and-burn-roles",
		Version:      Version,
		Description:  "Grants mint and burn roles on the token to an account",
		ContractType: ContractType,
		ContractABI:  burn_mint_erc20_with_drip.BurnMintERC20WithDripABI,
		Contract:     c,
		IsAllowedCaller: func(c *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
			return c.HasRole(opts, DefaultAdminRole, caller)
		},
		Validate: func(common.Address) error { return nil },
		CallContract: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
			return token.GrantMintAndBurnRoles(opts, account)
		},
	})
}

func NewWriteRevokeBurnRole(c *burn_mint_erc20_with_drip.BurnMintERC20WithDrip) *cld_ops.Operation[contract.FunctionInput[common.Address], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc20_with_drip.BurnMintERC20WithDrip]{
		Name:         "burn-mint-erc20-with-drip:revoke-burn-role",
		Version:      Version,
		Description:  "Revokes the burn role on the token from an account",
		ContractType: ContractType,
		ContractABI:  burn_mint_erc20_with_drip.BurnMintERC20WithDripABI,
		Contract:     c,
		IsAllowedCaller: func(c *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
			return c.HasRole(opts, DefaultAdminRole, caller)
		},
		Validate: func(common.Address) error { return nil },
		CallContract: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
			return token.RevokeRole(opts, [32]byte(BurnRole), account)
		},
	})
}

func NewWriteRevokeMintRole(c *burn_mint_erc20_with_drip.BurnMintERC20WithDrip) *cld_ops.Operation[contract.FunctionInput[common.Address], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc20_with_drip.BurnMintERC20WithDrip]{
		Name:         "burn-mint-erc20-with-drip:revoke-mint-role",
		Version:      Version,
		Description:  "Revokes the mint role on the token from an account",
		ContractType: ContractType,
		ContractABI:  burn_mint_erc20_with_drip.BurnMintERC20WithDripABI,
		Contract:     c,
		IsAllowedCaller: func(c *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
			return c.HasRole(opts, DefaultAdminRole, caller)
		},
		Validate: func(common.Address) error { return nil },
		CallContract: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
			return token.RevokeRole(opts, [32]byte(MintRole), account)
		},
	})
}

func NewWriteMint(c *burn_mint_erc20_with_drip.BurnMintERC20WithDrip) *cld_ops.Operation[contract.FunctionInput[MintArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[MintArgs, *burn_mint_erc20_with_drip.BurnMintERC20WithDrip]{
		Name:         "burn-mint-erc20-with-drip:mint",
		Version:      Version,
		Description:  "Mint tokens to an account",
		ContractType: ContractType,
		ContractABI:  burn_mint_erc20_with_drip.BurnMintERC20WithDripABI,
		Contract:     c,
		IsAllowedCaller: func(c *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, caller common.Address, input MintArgs) (bool, error) {
			return c.HasRole(opts, [32]byte(MintRole), caller)
		},
		Validate: func(MintArgs) error { return nil },
		CallContract: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.TransactOpts, args MintArgs) (*types.Transaction, error) {
			return token.Mint(opts, args.Account, args.Amount)
		},
	})
}
