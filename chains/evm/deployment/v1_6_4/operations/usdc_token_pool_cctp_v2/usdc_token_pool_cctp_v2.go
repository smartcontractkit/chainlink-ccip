package usdc_token_pool_cctp_v2_cctp_v2

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool_cctp_v2"
)

var ContractType cldf_deployment.ContractType = "USDCTokenPoolCCTPV2"
var Version = semver.MustParse("1.6.4")

type DomainUpdate = usdc_token_pool_cctp_v2.USDCTokenPoolDomainUpdate
type AuthorizedCallerUpdate = usdc_token_pool_cctp_v2.AuthorizedCallersAuthorizedCallerArgs


type ConstructorArgs struct {
	TokenMessenger              common.Address
	CCTPMessageTransmitterProxy common.Address
	Token                       common.Address
	Allowlist                   []common.Address
	RMNProxy                    common.Address
	Router                      common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "usdc-token-pool-cctp-v2:deploy",
	Version:          Version,
	Description:      "Deploys the USDCTokenPoolCCTPV2 contract",
	ContractMetadata: usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2Bin),
		},
	},
	Validate: func(args ConstructorArgs) error {
		// Ensure none of the critical addresses or allowlist are zeroed.
		if args.TokenMessenger == (common.Address{}) {
			return errors.New("tokenMessenger address cannot be zero")
		}
		if args.CCTPMessageTransmitterProxy == (common.Address{}) {
			return errors.New("cctpMessageTransmitterProxy address cannot be zero")
		}
		if args.Token == (common.Address{}) {
			return errors.New("token address cannot be zero")
		}
		if args.RMNProxy == (common.Address{}) {
			return errors.New("rmnProxy address cannot be zero")
		}
		if args.Router == (common.Address{}) {
			return errors.New("router address cannot be zero")
		}
		for i, addr := range args.Allowlist {
			if addr == (common.Address{}) {
				return fmt.Errorf("allowlist address at index %d cannot be zero", i)
			}
		}
		return nil
	},
})

var USDCTokenPoolSetDomains = contract.NewWrite(contract.WriteParams[[]DomainUpdate, *usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2]{
	Name:            "usdc-token-pool:set-domains",
	Version:         Version,
	Description:     "Sets domains on the USDCTokenPool contract",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2ABI,
	NewContract:     usdc_token_pool_cctp_v2.NewUSDCTokenPoolCCTPV2,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2, []DomainUpdate],
	Validate:        func([]DomainUpdate) error { return nil },
	CallContract: func(usdcTokenPool *usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2, opts *bind.TransactOpts, args []DomainUpdate) (*types.Transaction, error) {
		return usdcTokenPool.SetDomains(opts, args)
	},
})

var USDCTokenPoolGetDomain = contract.NewRead(contract.ReadParams[uint64, usdc_token_pool_cctp_v2.USDCTokenPoolDomain, *usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2]{
	Version:      Version,
	Description:  "Gets a domain on the USDCTokenPool contract",
	ContractType: ContractType,
	NewContract:  usdc_token_pool_cctp_v2.NewUSDCTokenPoolCCTPV2,
	CallContract: func(usdcTokenPool *usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2, opts *bind.CallOpts, chainSelector uint64) (usdc_token_pool_cctp_v2.USDCTokenPoolDomain, error) {
		return usdcTokenPool.GetDomain(opts, chainSelector)
	},
})

var USDCTokenPoolUpdateAuthorizedCallers = contract.NewWrite(contract.WriteParams[AuthorizedCallerUpdate, *usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2]{
	Name:            "usdc-token-pool:update-authorized-callers",
	Version:         Version,
	Description:     "Updates the authorized callers on the USDCTokenPool contract",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2ABI,
	NewContract:     usdc_token_pool_cctp_v2.NewUSDCTokenPoolCCTPV2,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2, AuthorizedCallerUpdate],
	Validate:        func(AuthorizedCallerUpdate) error { return nil },
	CallContract: func(usdcTokenPool *usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2, opts *bind.TransactOpts, args AuthorizedCallerUpdate) (*types.Transaction, error) {
		return usdcTokenPool.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var USDCTokenPoolTransferOwnership = contract.NewWrite(contract.WriteParams[common.Address, *usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2]{
	Name:            "usdc-token-pool:transfer-ownership",
	Version:         Version,
	Description:     "Transfers ownership of the USDCTokenPool contract",
	ContractType:    ContractType,
	ContractABI:     usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2ABI,
	NewContract:     usdc_token_pool_cctp_v2.NewUSDCTokenPoolCCTPV2,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2, common.Address],
	Validate: func(newOwner common.Address) error {
		if newOwner == (common.Address{}) {
			return fmt.Errorf("new owner cannot be the zero address")
		}
		return nil
	},
	CallContract: func(usdcTokenPool *usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return usdcTokenPool.TransferOwnership(opts, args)
	},
})
