package usdc_token_pool_proxy

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "USDCTokenPoolProxy"

var Version = semver.MustParse("1.7.0")

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

var UpdateLockOrBurnMechanisms = contract.NewWrite(contract.WriteParams[UpdateLockOrBurnMechanismsArgs, *usdc_token_pool_proxy.USDCTokenPoolProxy]{
	Name:            "usdc-token-pool-proxy:update-lock-or-burn-mechanisms",
	Version:         Version,
	Description:     "Updates lock or burn mechanisms on the USDCTokenPoolProxy",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool_proxy.USDCTokenPoolProxyABI,
	NewContract:     usdc_token_pool_proxy.NewUSDCTokenPoolProxy,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool_proxy.USDCTokenPoolProxy, UpdateLockOrBurnMechanismsArgs],
	Validate: func(proxy *usdc_token_pool_proxy.USDCTokenPoolProxy, backend bind.ContractBackend, opts *bind.CallOpts, args UpdateLockOrBurnMechanismsArgs) error {
		if len(args.RemoteChainSelectors) != len(args.Mechanisms) {
			return errors.New("remote chain selectors and mechanisms must have the same length")
		}
		for _, mechanism := range args.Mechanisms {
			if mechanism == 0 || mechanism > 4 {
				return errors.New("invalid mechanism, must be [1-4] - CCTP_V1, CCTP_V2, LOCK_RELEASE, CCTP_V2_WITH_CCV")
			}
		}
		return nil
	},
	IsNoop: func(proxy *usdc_token_pool_proxy.USDCTokenPoolProxy, opts *bind.CallOpts, args UpdateLockOrBurnMechanismsArgs) (bool, error) {
		for i, arg := range args.RemoteChainSelectors {
			actualMechanism, err := proxy.GetLockOrBurnMechanism(opts, arg)
			if err != nil {
				return false, fmt.Errorf("failed to get lock or burn mechanism for remote chain selector %d: %w", arg, err)
			}
			if actualMechanism != args.Mechanisms[i] {
				return false, nil
			}
		}

		return true, nil
	},
	CallContract: func(proxy *usdc_token_pool_proxy.USDCTokenPoolProxy, opts *bind.TransactOpts, args UpdateLockOrBurnMechanismsArgs) (*types.Transaction, error) {
		return proxy.UpdateLockOrBurnMechanisms(opts, args.RemoteChainSelectors, args.Mechanisms)
	},
})
