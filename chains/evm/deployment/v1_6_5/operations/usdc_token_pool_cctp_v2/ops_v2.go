package usdc_token_pool_cctp_v2

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
	Name:        "usdc-token-pool-cctp-v2:deploy",
	Version:     Version,
	Description: "Deploys the USDCTokenPoolCCTPV2 contract",
	ContractMetadata: &bind.MetaData{
		ABI: USDCTokenPoolCCTPV2ABI,
		Bin: USDCTokenPoolCCTPV2Bin,
	},
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(USDCTokenPoolCCTPV2Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

func NewReadGetAllAuthorizedCallers(c *USDCTokenPoolCCTPV2Contract) *cld_ops.Operation[contract.FunctionInput[struct{}], []common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, []common.Address, *USDCTokenPoolCCTPV2Contract]{
		Name:         "usdc-token-pool-cctp-v2:get-all-authorized-callers",
		Version:      Version,
		Description:  "Calls getAllAuthorizedCallers on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(c *USDCTokenPoolCCTPV2Contract, opts *bind.CallOpts, args struct{}) ([]common.Address, error) {
			return c.GetAllAuthorizedCallers(opts)
		},
	})
}

func NewWriteApplyAuthorizedCallerUpdates(c *USDCTokenPoolCCTPV2Contract) *cld_ops.Operation[contract.FunctionInput[AuthorizedCallerArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *USDCTokenPoolCCTPV2Contract]{
		Name:            "usdc-token-pool-cctp-v2:apply-authorized-caller-updates",
		Version:         Version,
		Description:     "Calls applyAuthorizedCallerUpdates on the contract",
		ContractType:    ContractType,
		ContractABI:     USDCTokenPoolCCTPV2ABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*USDCTokenPoolCCTPV2Contract, AuthorizedCallerArgs],
		Validate:        func(AuthorizedCallerArgs) error { return nil },
		CallContract: func(
			c *USDCTokenPoolCCTPV2Contract,
			opts *bind.TransactOpts,
			args AuthorizedCallerArgs,
		) (*types.Transaction, error) {
			return c.ApplyAuthorizedCallerUpdates(opts, args)
		},
	})
}

func NewReadGetDomain(c *USDCTokenPoolCCTPV2Contract) *cld_ops.Operation[contract.FunctionInput[uint64], Domain, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, Domain, *USDCTokenPoolCCTPV2Contract]{
		Name:         "usdc-token-pool-cctp-v2:get-domain",
		Version:      Version,
		Description:  "Calls getDomain on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(c *USDCTokenPoolCCTPV2Contract, opts *bind.CallOpts, args uint64) (Domain, error) {
			return c.GetDomain(opts, args)
		},
	})
}

func NewWriteSetDomains(c *USDCTokenPoolCCTPV2Contract) *cld_ops.Operation[contract.FunctionInput[[]DomainUpdate], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[[]DomainUpdate, *USDCTokenPoolCCTPV2Contract]{
		Name:            "usdc-token-pool-cctp-v2:set-domains",
		Version:         Version,
		Description:     "Calls setDomains on the contract",
		ContractType:    ContractType,
		ContractABI:     USDCTokenPoolCCTPV2ABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*USDCTokenPoolCCTPV2Contract, []DomainUpdate],
		Validate:        func([]DomainUpdate) error { return nil },
		CallContract: func(
			c *USDCTokenPoolCCTPV2Contract,
			opts *bind.TransactOpts,
			args []DomainUpdate,
		) (*types.Transaction, error) {
			return c.SetDomains(opts, args)
		},
	})
}
