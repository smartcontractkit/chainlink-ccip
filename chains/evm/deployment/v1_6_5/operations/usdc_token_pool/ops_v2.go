package usdc_token_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var DeployV2 = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:        "usdc-token-pool:deploy",
	Version:     Version,
	Description: "Deploys the USDCTokenPool contract",
	ContractMetadata: &bind.MetaData{
		ABI: USDCTokenPoolABI,
		Bin: USDCTokenPoolBin,
	},
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(USDCTokenPoolBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

func NewReadGetAllAuthorizedCallers(c *USDCTokenPoolContract) *cld_ops.Operation[contract.FunctionInput[struct{}], []common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, []common.Address, *USDCTokenPoolContract]{
		Name:         "usdc-token-pool:get-all-authorized-callers",
		Version:      Version,
		Description:  "Calls getAllAuthorizedCallers on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(c *USDCTokenPoolContract, opts *bind.CallOpts, args struct{}) ([]common.Address, error) {
			return c.GetAllAuthorizedCallers(opts)
		},
	})
}

func NewWriteApplyAuthorizedCallerUpdates(c *USDCTokenPoolContract) *cld_ops.Operation[contract.FunctionInput[AuthorizedCallerArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *USDCTokenPoolContract]{
		Name:            "usdc-token-pool:apply-authorized-caller-updates",
		Version:         Version,
		Description:     "Calls applyAuthorizedCallerUpdates on the contract",
		ContractType:    ContractType,
		ContractABI:     USDCTokenPoolABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*USDCTokenPoolContract, AuthorizedCallerArgs],
		Validate:        func(AuthorizedCallerArgs) error { return nil },
		CallContract: func(
			c *USDCTokenPoolContract,
			opts *bind.TransactOpts,
			args AuthorizedCallerArgs,
		) (*types.Transaction, error) {
			return c.ApplyAuthorizedCallerUpdates(opts, args)
		},
	})
}

func NewWriteAddRemotePool(c *USDCTokenPoolContract) *cld_ops.Operation[contract.FunctionInput[AddRemotePoolArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[AddRemotePoolArgs, *USDCTokenPoolContract]{
		Name:            "usdc-token-pool:add-remote-pool",
		Version:         Version,
		Description:     "Calls addRemotePool on the contract",
		ContractType:    ContractType,
		ContractABI:     USDCTokenPoolABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*USDCTokenPoolContract, AddRemotePoolArgs],
		Validate:        func(AddRemotePoolArgs) error { return nil },
		CallContract: func(
			c *USDCTokenPoolContract,
			opts *bind.TransactOpts,
			args AddRemotePoolArgs,
		) (*types.Transaction, error) {
			return c.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
		},
	})
}

func NewReadGetDomain(c *USDCTokenPoolContract) *cld_ops.Operation[contract.FunctionInput[uint64], Domain, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, Domain, *USDCTokenPoolContract]{
		Name:         "usdc-token-pool:get-domain",
		Version:      Version,
		Description:  "Calls getDomain on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(c *USDCTokenPoolContract, opts *bind.CallOpts, args uint64) (Domain, error) {
			return c.GetDomain(opts, args)
		},
	})
}

func NewWriteSetDomains(c *USDCTokenPoolContract) *cld_ops.Operation[contract.FunctionInput[[]DomainUpdate], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[[]DomainUpdate, *USDCTokenPoolContract]{
		Name:            "usdc-token-pool:set-domains",
		Version:         Version,
		Description:     "Calls setDomains on the contract",
		ContractType:    ContractType,
		ContractABI:     USDCTokenPoolABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*USDCTokenPoolContract, []DomainUpdate],
		Validate:        func([]DomainUpdate) error { return nil },
		CallContract: func(
			c *USDCTokenPoolContract,
			opts *bind.TransactOpts,
			args []DomainUpdate,
		) (*types.Transaction, error) {
			return c.SetDomains(opts, args)
		},
	})
}
