package erc20_lock_box

import (
	"errors"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "ERC20LockBox"
var Version = semver.MustParse("1.7.0")

type ConstructorArgs struct {
	TokenAdminRegistry common.Address
}

type AllowedCallerConfigArgs = erc20_lock_box.ERC20LockBoxAllowedCallerConfigArgs

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "erc20-lock-box:deploy",
	Version:          Version,
	Description:      "Deploys the ERC20LockBox contract",
	ContractMetadata: erc20_lock_box.ERC20LockBoxMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(erc20_lock_box.ERC20LockBoxBin),
		},
	},
	Validate: func(args ConstructorArgs) error {
		if args.TokenAdminRegistry == (common.Address{}) {
			return errors.New("tokenAdminRegistry address cannot be zero")
		}
		return nil
	},
})

var ConfigureAllowedCallers = contract.NewWrite(contract.WriteParams[[]AllowedCallerConfigArgs, *erc20_lock_box.ERC20LockBox]{
	Name:            "erc20-lock-box:configure-allowed-callers",
	Version:         Version,
	Description:     "Configures allowed callers for the ERC20LockBox",
	ContractType:    ContractType,
	ContractABI:     erc20_lock_box.ERC20LockBoxABI,
	NewContract:     erc20_lock_box.NewERC20LockBox,
	IsAllowedCaller: contract.AllCallersAllowed[*erc20_lock_box.ERC20LockBox, []AllowedCallerConfigArgs],
	Validate:        func([]AllowedCallerConfigArgs) error { return nil },
	CallContract: func(lockBox *erc20_lock_box.ERC20LockBox, opts *bind.TransactOpts, args []AllowedCallerConfigArgs) (*types.Transaction, error) {
		return lockBox.ConfigureAllowedCallers(opts, args)
	},
})
