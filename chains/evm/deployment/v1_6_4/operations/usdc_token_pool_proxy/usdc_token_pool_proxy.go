package usdc_token_pool_proxy

import (
	"errors"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool_proxy"

	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "USDCTokenPoolProxy"
var Version *semver.Version = semver.MustParse("1.6.4")

type PoolAddresses = usdc_token_pool_proxy.USDCTokenPoolProxyPoolAddresses

type ConstructorArgs struct {
	Token     common.Address
	USDCToken common.Address
	Router    common.Address
}

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
	Version:          Version,
	Description:      "Deploys the USDCTokenPoolProxy contract",
	ContractMetadata: usdc_token_pool_proxy.USDCTokenPoolProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(usdc_token_pool_proxy.USDCTokenPoolProxyBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var USDCTokenPoolProxyUpdateLockOrBurnMechanisms = contract.NewWrite(contract.WriteParams[UpdateLockOrBurnMechanismsArgs, *usdc_token_pool_proxy.USDCTokenPoolProxy]{
	Name:            "usdc-token-pool-proxy:update-lock-or-burn-mechanisms",
	Version:         Version,
	Description:     "Updates the lock or burn mechanisms on the USDCTokenPoolProxy contract",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool_proxy.USDCTokenPoolProxyABI,
	NewContract:     usdc_token_pool_proxy.NewUSDCTokenPoolProxy,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool_proxy.USDCTokenPoolProxy, UpdateLockOrBurnMechanismsArgs],
	Validate: func(args UpdateLockOrBurnMechanismsArgs) error {
		if len(args.RemoteChainSelectors) != len(args.Mechanisms) {
			return errors.New("remote chain selectors and mechanisms must be the same length")
		}
		return nil
	},
	CallContract: func(usdcTokenPoolProxy *usdc_token_pool_proxy.USDCTokenPoolProxy, opts *bind.TransactOpts, args UpdateLockOrBurnMechanismsArgs) (*types.Transaction, error) {
		return usdcTokenPoolProxy.UpdateLockOrBurnMechanisms(opts, args.RemoteChainSelectors, args.Mechanisms)
	},
})

var USDCTokenPoolProxyUpdatePoolAddresses = contract.NewWrite(contract.WriteParams[PoolAddresses, *usdc_token_pool_proxy.USDCTokenPoolProxy]{
	Name:            "usdc-token-pool-proxy:update-pool-addresses",
	Version:         Version,
	Description:     "Updates the pool addresses on the USDCTokenPoolProxy contract",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool_proxy.USDCTokenPoolProxyABI,
	NewContract:     usdc_token_pool_proxy.NewUSDCTokenPoolProxy,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool_proxy.USDCTokenPoolProxy, PoolAddresses],
	Validate: func(args PoolAddresses) error {
		return nil
	},
	CallContract: func(usdcTokenPoolProxy *usdc_token_pool_proxy.USDCTokenPoolProxy, opts *bind.TransactOpts, args PoolAddresses) (*types.Transaction, error) {
		return usdcTokenPoolProxy.UpdatePoolAddresses(opts, args)
	},
})

var USDCTokenPoolProxyUpdateLockReleasePoolAddresses = contract.NewWrite(contract.WriteParams[UpdateLockReleasePoolAddressesArgs, *usdc_token_pool_proxy.USDCTokenPoolProxy]{
	Name:            "usdc-token-pool-proxy:update-lock-release-pool-addresses",
	Version:         Version,
	Description:     "Updates the lock release pool addresses on the USDCTokenPoolProxy contract",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool_proxy.USDCTokenPoolProxyABI,
	NewContract:     usdc_token_pool_proxy.NewUSDCTokenPoolProxy,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool_proxy.USDCTokenPoolProxy, UpdateLockReleasePoolAddressesArgs],
	Validate: func(args UpdateLockReleasePoolAddressesArgs) error {
		if len(args.RemoteChainSelectors) != len(args.LockReleasePools) {
			return errors.New("remote chain selectors and lock release pools must be the same length")
		}
		return nil
	},
	CallContract: func(usdcTokenPoolProxy *usdc_token_pool_proxy.USDCTokenPoolProxy, opts *bind.TransactOpts, args UpdateLockReleasePoolAddressesArgs) (*types.Transaction, error) {
		return usdcTokenPoolProxy.UpdateLockReleasePoolAddresses(opts, args.RemoteChainSelectors, args.LockReleasePools)
	},
})
