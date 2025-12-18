package usdc_token_pool_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "USDCTokenPoolProxy"

type ConstructorArgs struct {
	Token        common.Address
	Pools        USDCTokenPoolProxyPoolAddresses
	Router       common.Address
	CCTPVerifier common.Address
}

type USDCTokenPoolProxyPoolAddresses = usdc_token_pool_proxy.USDCTokenPoolProxyPoolAddresses

type UpdateLockOrBurnMechanismsArgs struct {
	RemoteChainSelectors []uint64
	Mechanisms           []uint8
}

type UpdateLockReleasePoolAddressesArgs struct {
	RemoteChainSelectors []uint64
	LockReleasePools     []common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "usdc-token-pool-proxy:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the USDCTokenPoolProxy contract",
	ContractMetadata: usdc_token_pool_proxy.USDCTokenPoolProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(usdc_token_pool_proxy.USDCTokenPoolProxyBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var UpdateLockOrBurnMechanisms = contract.NewWrite(contract.WriteParams[UpdateLockOrBurnMechanismsArgs, *usdc_token_pool_proxy.USDCTokenPoolProxy]{
	Name:            "usdc-token-pool-proxy:update-lock-or-burn-mechanisms",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Updates lock or burn mechanisms on the USDCTokenPoolProxy",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool_proxy.USDCTokenPoolProxyABI,
	NewContract:     usdc_token_pool_proxy.NewUSDCTokenPoolProxy,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool_proxy.USDCTokenPoolProxy, UpdateLockOrBurnMechanismsArgs],
	Validate:        func(UpdateLockOrBurnMechanismsArgs) error { return nil },
	CallContract: func(proxy *usdc_token_pool_proxy.USDCTokenPoolProxy, opts *bind.TransactOpts, args UpdateLockOrBurnMechanismsArgs) (*types.Transaction, error) {
		return proxy.UpdateLockOrBurnMechanisms(opts, args.RemoteChainSelectors, args.Mechanisms)
	},
})

var UpdateLockReleasePoolAddresses = contract.NewWrite(contract.WriteParams[UpdateLockReleasePoolAddressesArgs, *usdc_token_pool_proxy.USDCTokenPoolProxy]{
	Name:            "usdc-token-pool-proxy:update-lock-release-pool-addresses",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Updates lock release pool addresses on the USDCTokenPoolProxy",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool_proxy.USDCTokenPoolProxyABI,
	NewContract:     usdc_token_pool_proxy.NewUSDCTokenPoolProxy,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool_proxy.USDCTokenPoolProxy, UpdateLockReleasePoolAddressesArgs],
	Validate:        func(UpdateLockReleasePoolAddressesArgs) error { return nil },
	CallContract: func(proxy *usdc_token_pool_proxy.USDCTokenPoolProxy, opts *bind.TransactOpts, args UpdateLockReleasePoolAddressesArgs) (*types.Transaction, error) {
		return proxy.UpdateLockReleasePoolAddresses(opts, args.RemoteChainSelectors, args.LockReleasePools)
	},
})
