package rmn_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"
)

var ContractType cldf_deployment.ContractType = "ARMProxy"

var Version *semver.Version = semver.MustParse("1.0.0")

type ConstructorArgs struct {
	RMN common.Address
}

type SetRMNArgs = struct {
	RMN common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "rmn_proxy:deploy",
	Version:          Version,
	Description:      "Deploys the RMNProxy contract",
	ContractMetadata: rmn_proxy_contract.RMNProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(rmn_proxy_contract.RMNProxyBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

func NewWriteSetRMN(c *rmn_proxy_contract.RMNProxy) *cld_ops.Operation[contract.FunctionInput[SetRMNArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[SetRMNArgs, *rmn_proxy_contract.RMNProxy]{
		Name:            "rmn_proxy:set-rmn",
		Version:         Version,
		Description:     "Sets the RMN address on the RMNProxy",
		ContractType:    ContractType,
		ContractABI:     rmn_proxy_contract.RMNProxyABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*rmn_proxy_contract.RMNProxy, SetRMNArgs],
		Validate:        func(SetRMNArgs) error { return nil },
		CallContract: func(rmnProxy *rmn_proxy_contract.RMNProxy, opts *bind.TransactOpts, args SetRMNArgs) (*types.Transaction, error) {
			return rmnProxy.SetARM(opts, args.RMN)
		},
	})
}

func NewReadGetRMN(c *rmn_proxy_contract.RMNProxy) *cld_ops.Operation[contract.FunctionInput[struct{}], common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, common.Address, *rmn_proxy_contract.RMNProxy]{
		Name:         "rmn_proxy:get-rmn",
		Version:      semver.MustParse("1.0.0"),
		Description:  "Gets the RMN address set on the RMNProxy",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(rmnProxy *rmn_proxy_contract.RMNProxy, opts *bind.CallOpts, args struct{}) (common.Address, error) {
			return rmnProxy.GetARM(opts)
		},
	})
}
