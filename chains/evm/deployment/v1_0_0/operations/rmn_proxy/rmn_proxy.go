package rmn_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"
)

var ContractType cldf_deployment.ContractType = "ARMProxy"

type ConstructorArgs struct {
	RMN common.Address
}

type SetRMNArgs = struct {
	RMN common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "rmn_proxy:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys the RMNProxy contract",
	ContractMetadata: rmn_proxy_contract.RMNProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.0.0")).String(): {
			EVM: common.FromHex(rmn_proxy_contract.RMNProxyBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var SetRMN = contract.NewWrite(contract.WriteParams[SetRMNArgs, *rmn_proxy_contract.RMNProxy]{
	Name:            "rmn_proxy:set-rmn",
	Version:         semver.MustParse("1.0.0"),
	Description:     "Sets the RMN address on the RMNProxy",
	ContractType:    ContractType,
	ContractABI:     rmn_proxy_contract.RMNProxyABI,
	NewContract:     rmn_proxy_contract.NewRMNProxy,
	IsAllowedCaller: contract.OnlyOwner[*rmn_proxy_contract.RMNProxy, SetRMNArgs],
	Validate:        func(SetRMNArgs) error { return nil },
	CallContract: func(rmnProxy *rmn_proxy_contract.RMNProxy, opts *bind.TransactOpts, args SetRMNArgs) (*types.Transaction, error) {
		return rmnProxy.SetARM(opts, args.RMN)
	},
})
