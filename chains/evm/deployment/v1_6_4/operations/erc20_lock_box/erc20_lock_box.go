package erc20_lock_box

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/erc20_lock_box"

	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "ERC20Lockbox"
var Version *semver.Version = semver.MustParse("1.6.4")

type ConstructorArgs struct {
	TokenAdminRegistry common.Address
}

type AllowedCallerConfigArgs = erc20_lock_box.ERC20LockBoxAllowedCallerConfigArgs

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "erc20-lock-box:deploy",
	Version:          Version,
	Description:      "Deploys the ERC20Lockbox contract",
	ContractMetadata: erc20_lock_box.ERC20LockBoxMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(erc20_lock_box.ERC20LockBoxBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ERC20LockboxConfigureAllowedCallers = contract.NewWrite(contract.WriteParams[[]AllowedCallerConfigArgs, *erc20_lock_box.ERC20LockBox]{
	Name:         "erc20-lock-box:configure-allowed-callers",
	Version:      Version,
	Description:  "Sets the allowed callers on the ERC20LockBox contract",
	ContractType: ContractType,
	ContractABI:  erc20_lock_box.ERC20LockBoxABI,
	NewContract:  erc20_lock_box.NewERC20LockBox,
	IsAllowedCaller: func(erc20Lockbox *erc20_lock_box.ERC20LockBox, opts *bind.CallOpts, caller common.Address, input []AllowedCallerConfigArgs) (bool, error) {
		// TODO: Return to possibly checking for ownership if contract reads are possible
		return true, nil
	},
	Validate: func([]AllowedCallerConfigArgs) error { return nil },
	CallContract: func(erc20Lockbox *erc20_lock_box.ERC20LockBox, opts *bind.TransactOpts, args []AllowedCallerConfigArgs) (*types.Transaction, error) {
		return erc20Lockbox.ConfigureAllowedCallers(opts, args)
	},
})
