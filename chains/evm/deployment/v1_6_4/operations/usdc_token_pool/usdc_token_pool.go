package usdc_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "USDCTokenPool"
var Version *semver.Version = semver.MustParse("1.6.4")

type DomainUpdate = usdc_token_pool.USDCTokenPoolDomainUpdate

type AuthorizedCallerUpdate = usdc_token_pool.AuthorizedCallersAuthorizedCallerArgs

type ConstructorArgs struct {
	TokenMessenger              common.Address
	CCTPMessageTransmitterProxy common.Address
	Token                       common.Address
	Allowlist                   []common.Address
	RMNProxy                    common.Address
	Router                      common.Address
	SupportedUSDCVersion        uint32
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "usdc-token-pool:deploy",
	Version:          Version,
	Description:      "Deploys the USDCTokenPool contract",
	ContractMetadata: usdc_token_pool.USDCTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(usdc_token_pool.USDCTokenPoolBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var USDCTokenPoolSetDomains = contract.NewWrite(contract.WriteParams[[]DomainUpdate, *usdc_token_pool.USDCTokenPool]{
	Name:            "usdc-token-pool:set-domains",
	Version:         Version,
	Description:     "Sets domains on the USDCTokenPool contract",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool.USDCTokenPoolABI,
	NewContract:     usdc_token_pool.NewUSDCTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool.USDCTokenPool, []DomainUpdate],
	Validate:        func([]DomainUpdate) error { return nil },
	CallContract: func(usdcTokenPool *usdc_token_pool.USDCTokenPool, opts *bind.TransactOpts, args []DomainUpdate) (*types.Transaction, error) {
		return usdcTokenPool.SetDomains(opts, args)
	},
})

var USDCTokenPoolGetDomain = contract.NewRead(contract.ReadParams[uint64, usdc_token_pool.USDCTokenPoolDomain, *usdc_token_pool.USDCTokenPool]{
	Version:      Version,
	Description:  "Gets a domain on the USDCTokenPool contract",
	ContractType: ContractType,
	NewContract:  usdc_token_pool.NewUSDCTokenPool,
	CallContract: func(usdcTokenPool *usdc_token_pool.USDCTokenPool, opts *bind.CallOpts, chainSelector uint64) (usdc_token_pool.USDCTokenPoolDomain, error) {
		return usdcTokenPool.GetDomain(opts, chainSelector)
	},
})

var USDCTokenPoolUpdateAuthorizedCallers = contract.NewWrite(contract.WriteParams[AuthorizedCallerUpdate, *usdc_token_pool.USDCTokenPool]{
	Name:            "usdc-token-pool:update-authorized-callers",
	Version:         Version,
	Description:     "Updates the authorized callers on the USDCTokenPool contract",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool.USDCTokenPoolABI,
	NewContract:     usdc_token_pool.NewUSDCTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool.USDCTokenPool, AuthorizedCallerUpdate],
	Validate:        func(AuthorizedCallerUpdate) error { return nil },
	CallContract: func(usdcTokenPool *usdc_token_pool.USDCTokenPool, opts *bind.TransactOpts, args AuthorizedCallerUpdate) (*types.Transaction, error) {
		return usdcTokenPool.ApplyAuthorizedCallerUpdates(opts, args)
	},
})
