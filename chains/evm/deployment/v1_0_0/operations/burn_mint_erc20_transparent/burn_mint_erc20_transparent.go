package burn_mint_erc20_transparent

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/latest/burn_mint_erc20_transparent"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/latest/transparent_upgradeable_proxy"
)

// ContractType is the type used for the proxy address, which is the address
// that token/pool refs resolve to; the implementation address is an internal
// deployment detail and is not tracked as its own ref.
var ContractType cldf_deployment.ContractType = "BurnMintERC20TransparentToken"

// ImplementationContractType types the (logic) implementation deploy, kept
// distinct from ContractType so the impl and proxy deploy operations don't
// collide on the same TypeAndVersion bytecode-lookup key.
var ImplementationContractType cldf_deployment.ContractType = "BurnMintERC20TransparentTokenImplementation"

// ImplConstructorArgs is empty: BurnMintERC20Transparent's constructor takes
// no arguments and only calls _disableInitializers(); all setup happens via
// initialize() through the proxy.
type ImplConstructorArgs struct{}

var DeployImplementation = contract.NewDeploy(contract.DeployParams[ImplConstructorArgs]{
	Name:             "burn_mint_erc20_transparent:deploy-implementation",
	Version:          utils.Version_1_0_0,
	Description:      "Deploys the BurnMintERC20Transparent implementation (logic) contract",
	ContractMetadata: burn_mint_erc20_transparent.BurnMintERC20TransparentMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ImplementationContractType, *utils.Version_1_0_0).String(): {
			EVM: common.FromHex(burn_mint_erc20_transparent.BurnMintERC20TransparentBin),
		},
	},
	Validate: func(ImplConstructorArgs) error { return nil },
})

// ProxyConstructorArgs mirrors TransparentUpgradeableProxy's constructor:
// TransparentUpgradeableProxy(address _logic, address initialOwner, bytes memory _data).
// InitialOwner becomes the owner of the proxy's own auto-deployed ProxyAdmin,
// i.e. the upgrade authority for this token.
type ProxyConstructorArgs struct {
	Logic        common.Address
	InitialOwner common.Address
	Data         []byte
}

var DeployProxy = contract.NewDeploy(contract.DeployParams[ProxyConstructorArgs]{
	Name:             "burn_mint_erc20_transparent:deploy-proxy",
	Version:          utils.Version_1_0_0,
	Description:      "Deploys the TransparentUpgradeableProxy fronting a BurnMintERC20Transparent implementation",
	ContractMetadata: transparent_upgradeable_proxy.TransparentUpgradeableProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *utils.Version_1_0_0).String(): {
			EVM: common.FromHex(transparent_upgradeable_proxy.TransparentUpgradeableProxyBin),
		},
	},
	Validate: func(args ProxyConstructorArgs) error {
		if args.Logic == (common.Address{}) {
			return fmt.Errorf("logic (implementation) address must not be the zero address")
		}
		if args.InitialOwner == (common.Address{}) {
			return fmt.Errorf("initialOwner (proxy admin / upgrade authority) must not be the zero address")
		}
		return nil
	},
})

// InitializeArgs mirrors BurnMintERC20Transparent.initialize's parameters.
// DefaultAdmin receives DEFAULT_ADMIN_ROLE, the default ccipAdmin, and any
// preMint amount, mirroring how BurnMintERC20's constructor implicitly makes
// msg.sender (the deployer) hold all three at deploy time; downstream steps
// in the deploy sequence (Transfer/SetCCIPAdmin) then move preMint/ccipAdmin
// to their intended holders exactly as they do for BurnMintERC20.
type InitializeArgs struct {
	Name         string
	Symbol       string
	Decimals     uint8
	MaxSupply    *big.Int
	PreMint      *big.Int
	DefaultAdmin common.Address
}

var Initialize = contract.NewWrite(contract.WriteParams[InitializeArgs, *burn_mint_erc20_transparent.BurnMintERC20Transparent]{
	Name:         "burn_mint_erc20_transparent:initialize",
	Version:      utils.Version_1_0_0,
	Description:  "Calls initialize on a freshly deployed BurnMintERC20Transparent proxy",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20_transparent.BurnMintERC20TransparentABI,
	NewContract:  burn_mint_erc20_transparent.NewBurnMintERC20Transparent,
	IsAllowedCaller: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.CallOpts, caller common.Address, input InitializeArgs) (bool, error) {
		// initialize() is guarded only by OZ's `initializer` modifier (first-caller-wins), not
		// by any role/ownership check, so any caller - including the deployer key executing
		// this operation immediately after deploy - is allowed to call it.
		return true, nil
	},
	Validate: func(args InitializeArgs) error {
		if args.DefaultAdmin == (common.Address{}) {
			return fmt.Errorf("defaultAdmin must not be the zero address")
		}
		if args.PreMint != nil && args.MaxSupply != nil && args.MaxSupply.Sign() > 0 && args.PreMint.Cmp(args.MaxSupply) > 0 {
			return fmt.Errorf("preMint (%s) exceeds maxSupply (%s)", args.PreMint, args.MaxSupply)
		}
		return nil
	},
	CallContract: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.TransactOpts, input InitializeArgs) (*types.Transaction, error) {
		maxSupply := input.MaxSupply
		if maxSupply == nil {
			maxSupply = big.NewInt(0)
		}
		preMint := input.PreMint
		if preMint == nil {
			preMint = big.NewInt(0)
		}
		return token.Initialize(opts, input.Name, input.Symbol, input.Decimals, maxSupply, preMint, input.DefaultAdmin)
	},
})

var SetCCIPAdmin = contract.NewWrite(contract.WriteParams[string, *burn_mint_erc20_transparent.BurnMintERC20Transparent]{
	Name:         "burn_mint_erc20_transparent:set-ccip-admin",
	Version:      utils.Version_1_0_0,
	Description:  "Set CCIP Admin on a BurnMintERC20Transparent token proxy",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20_transparent.BurnMintERC20TransparentABI,
	NewContract:  burn_mint_erc20_transparent.NewBurnMintERC20Transparent,
	IsAllowedCaller: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.CallOpts, caller common.Address, input string) (bool, error) {
		// setCCIPAdmin is onlyRole(DEFAULT_ADMIN_ROLE) - not owner() - even though owner()
		// (from IERC5313) happens to alias defaultAdmin() on this contract.
		defaultAdminRole, err := token.DEFAULTADMINROLE(opts)
		if err != nil {
			return false, err
		}
		return token.HasRole(opts, defaultAdminRole, caller)
	},
	Validate: func(string) error { return nil },
	CallContract: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.TransactOpts, input string) (*types.Transaction, error) {
		return token.SetCCIPAdmin(opts, common.HexToAddress(input))
	},
})

var GrantMintAndBurnRoles = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc20_transparent.BurnMintERC20Transparent]{
	Name:         "burn_mint_erc20_transparent:grant-mint-and-burn-roles",
	Version:      utils.Version_1_0_0,
	Description:  "Grant mint and burn roles on a BurnMintERC20Transparent token proxy",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20_transparent.BurnMintERC20TransparentABI,
	NewContract:  burn_mint_erc20_transparent.NewBurnMintERC20Transparent,
	IsAllowedCaller: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		defaultAdminRole, err := token.DEFAULTADMINROLE(opts)
		if err != nil {
			return false, err
		}
		return token.HasRole(opts, defaultAdminRole, caller)
	},
	Validate: func(common.Address) error { return nil },
	CallContract: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return token.GrantMintAndBurnRoles(opts, input)
	},
})

var GetCCIPAdmin = contract.NewRead(contract.ReadParams[struct{}, common.Address, *burn_mint_erc20_transparent.BurnMintERC20Transparent]{
	Name:         "burn_mint_erc20_transparent:get-ccip-admin",
	Version:      utils.Version_1_0_0,
	Description:  "Gets the CCIP admin on a BurnMintERC20Transparent token proxy",
	ContractType: ContractType,
	NewContract:  burn_mint_erc20_transparent.NewBurnMintERC20Transparent,
	CallContract: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.CallOpts, input struct{}) (common.Address, error) {
		return token.GetCCIPAdmin(opts)
	},
})

// BeginDefaultAdminTransfer starts the 2-step DEFAULT_ADMIN_ROLE transfer required by
// AccessControlDefaultAdminRulesUpgradeable (grantRole/revokeRole revert unconditionally for
// DEFAULT_ADMIN_ROLE on this contract). It is onlyRole(DEFAULT_ADMIN_ROLE), so the deployer -
// who holds the role right after Initialize - can call it synchronously. Completing the
// transfer still requires a separate AcceptDefaultAdminTransfer call signed by newAdmin.
var BeginDefaultAdminTransfer = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc20_transparent.BurnMintERC20Transparent]{
	Name:         "burn_mint_erc20_transparent:begin-default-admin-transfer",
	Version:      utils.Version_1_0_0,
	Description:  "Begins the 2-step DEFAULT_ADMIN_ROLE transfer on a BurnMintERC20Transparent token proxy",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20_transparent.BurnMintERC20TransparentABI,
	NewContract:  burn_mint_erc20_transparent.NewBurnMintERC20Transparent,
	IsAllowedCaller: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		defaultAdminRole, err := token.DEFAULTADMINROLE(opts)
		if err != nil {
			return false, err
		}
		return token.HasRole(opts, defaultAdminRole, caller)
	},
	Validate: func(newAdmin common.Address) error {
		if newAdmin == (common.Address{}) {
			return fmt.Errorf("newAdmin must not be the zero address")
		}
		return nil
	},
	CallContract: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return token.BeginDefaultAdminTransfer(opts, input)
	},
})

// AcceptDefaultAdminTransfer completes the transfer started by BeginDefaultAdminTransfer.
// acceptDefaultAdminTransfer() has no role gate - it only requires msg.sender == pendingDefaultAdmin.
// The deployer is never pendingDefaultAdmin in our flow, so IsAllowedCaller always returns false
// here when called directly by the deploy sequence: the resulting WriteOutput is an unsigned,
// prepared transaction that folds into the MCMS batch/proposal. When (and only when) newAdmin is
// the chain's timelock, that proposal being executed later makes the timelock itself the caller,
// satisfying the pendingDefaultAdmin check and completing the transfer. For a customer-provided
// newAdmin, this op must not be queued at all - see tokenimpl's Capabilities/GrantAdminRole doc.
var AcceptDefaultAdminTransfer = contract.NewWrite(contract.WriteParams[struct{}, *burn_mint_erc20_transparent.BurnMintERC20Transparent]{
	Name:         "burn_mint_erc20_transparent:accept-default-admin-transfer",
	Version:      utils.Version_1_0_0,
	Description:  "Accepts a pending DEFAULT_ADMIN_ROLE transfer on a BurnMintERC20Transparent token proxy",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20_transparent.BurnMintERC20TransparentABI,
	NewContract:  burn_mint_erc20_transparent.NewBurnMintERC20Transparent,
	IsAllowedCaller: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.CallOpts, caller common.Address, input struct{}) (bool, error) {
		pending, err := token.PendingDefaultAdmin(opts)
		if err != nil {
			return false, err
		}
		return pending.NewAdmin == caller, nil
	},
	Validate: func(struct{}) error { return nil },
	CallContract: func(token *burn_mint_erc20_transparent.BurnMintERC20Transparent, opts *bind.TransactOpts, input struct{}) (*types.Transaction, error) {
		return token.AcceptDefaultAdminTransfer(opts)
	},
})
