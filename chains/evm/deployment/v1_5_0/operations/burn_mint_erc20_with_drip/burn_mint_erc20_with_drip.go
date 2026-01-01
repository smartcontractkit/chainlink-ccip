package burn_mint_erc20_with_drip

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/link_token"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
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
		cldf_deployment.NewTypeAndVersion(link_token.ContractType, *link_token.Version).String(): {
			EVM: common.FromHex(burn_mint_erc20_with_drip.BurnMintERC20WithDripBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var GrantMintAndBurnRoles = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc20_with_drip.BurnMintERC20WithDrip]{
	Name:         "burn-mint-erc20-with-drip:grant-mint-and-burn-roles",
	Version:      Version,
	Description:  "Grants mint and burn roles on the token to an account",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20_with_drip.BurnMintERC20WithDripABI,
	NewContract:  burn_mint_erc20_with_drip.NewBurnMintERC20WithDrip,
	IsAllowedCaller: func(contract *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		return contract.HasRole(opts, DefaultAdminRole, caller)
	},
	Validate: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, backend bind.ContractBackend, opts *bind.CallOpts, account common.Address) error {
		return nil
	},
	IsNoop: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, account common.Address) (bool, error) {
		hasMintRole, err := token.HasRole(opts, [32]byte(MintRole), account)
		if err != nil {
			return false, fmt.Errorf("failed to check if account %s has mint role on token with address %s: %w", account, token.Address(), err)
		}
		hasBurnRole, err := token.HasRole(opts, [32]byte(BurnRole), account)
		if err != nil {
			return false, fmt.Errorf("failed to check if account %s has burn role on token with address %s: %w", account, token.Address(), err)
		}
		return hasMintRole && hasBurnRole, nil
	},
	CallContract: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
		return token.GrantMintAndBurnRoles(opts, account)
	},
})

var RevokeBurnRole = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc20_with_drip.BurnMintERC20WithDrip]{
	Name:         "burn-mint-erc20-with-drip:revoke-burn-role",
	Version:      Version,
	Description:  "Revokes the burn role on the token from an account",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20_with_drip.BurnMintERC20WithDripABI,
	NewContract:  burn_mint_erc20_with_drip.NewBurnMintERC20WithDrip,
	IsAllowedCaller: func(contract *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		return contract.HasRole(opts, DefaultAdminRole, caller)
	},
	Validate: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, backend bind.ContractBackend, opts *bind.CallOpts, account common.Address) error {
		return nil
	},
	IsNoop: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, account common.Address) (bool, error) {
		hasBurnRole, err := token.HasRole(opts, [32]byte(BurnRole), account)
		if err != nil {
			return false, fmt.Errorf("failed to check if account %s has burn role on token with address %s: %w", account, token.Address(), err)
		}
		return !hasBurnRole, nil
	},
	CallContract: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
		return token.RevokeRole(opts, [32]byte(BurnRole), account)
	},
})

var RevokeMintRole = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc20_with_drip.BurnMintERC20WithDrip]{
	Name:         "burn-mint-erc20-with-drip:revoke-mint-role",
	Version:      Version,
	Description:  "Revokes the mint role on the token from an account",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20_with_drip.BurnMintERC20WithDripABI,
	NewContract:  burn_mint_erc20_with_drip.NewBurnMintERC20WithDrip,
	IsAllowedCaller: func(contract *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		return contract.HasRole(opts, DefaultAdminRole, caller)
	},
	Validate: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, backend bind.ContractBackend, opts *bind.CallOpts, account common.Address) error {
		return nil
	},
	IsNoop: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, account common.Address) (bool, error) {
		hasMintRole, err := token.HasRole(opts, [32]byte(MintRole), account)
		if err != nil {
			return false, fmt.Errorf("failed to check if account %s has mint role on token with address %s: %w", account, token.Address(), err)
		}
		return !hasMintRole, nil
	},
	CallContract: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
		return token.RevokeRole(opts, [32]byte(MintRole), account)
	},
})

var Mint = contract.NewWrite(contract.WriteParams[MintArgs, *burn_mint_erc20_with_drip.BurnMintERC20WithDrip]{
	Name:         "burn-mint-erc20-with-drip:mint",
	Version:      Version,
	Description:  "Mint tokens to an account",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20_with_drip.BurnMintERC20WithDripABI,
	NewContract:  burn_mint_erc20_with_drip.NewBurnMintERC20WithDrip,
	IsAllowedCaller: func(contract *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, caller common.Address, input MintArgs) (bool, error) {
		return contract.HasRole(opts, [32]byte(MintRole), caller)
	},
	Validate: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, backend bind.ContractBackend, opts *bind.CallOpts, args MintArgs) error {
		return nil
	},
	IsNoop: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.CallOpts, args MintArgs) (bool, error) {
		return false, nil
	},
	CallContract: func(token *burn_mint_erc20_with_drip.BurnMintERC20WithDrip, opts *bind.TransactOpts, args MintArgs) (*types.Transaction, error) {
		return token.Mint(opts, args.Account, args.Amount)
	},
})
