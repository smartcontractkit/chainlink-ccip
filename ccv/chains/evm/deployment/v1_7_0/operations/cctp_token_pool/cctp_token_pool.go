package cctp_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CCTPTokenPool"

type ConstructorArgs struct {
	Token              common.Address
	LocalTokenDecimals uint8
	AdvancedPoolHooks  common.Address
	RMNProxy           common.Address
	Router             common.Address
	AllowedCallers     []common.Address
}

type AuthorizedCallerArgs = cctp_token_pool.AuthorizedCallersAuthorizedCallerArgs

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "cctp-token-pool:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the CCTPTokenPool contract",
	ContractMetadata: cctp_token_pool.CCTPTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(cctp_token_pool.CCTPTokenPoolBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *cctp_token_pool.CCTPTokenPool]{
	Name:            "cctp-token-pool:apply-authorized-caller-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies authorized caller updates on the CCTPTokenPool",
	ContractType:    ContractType,
	ContractABI:     cctp_token_pool.CCTPTokenPoolABI,
	NewContract:     cctp_token_pool.NewCCTPTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*cctp_token_pool.CCTPTokenPool, AuthorizedCallerArgs],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(cctpTokenPool *cctp_token_pool.CCTPTokenPool, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return cctpTokenPool.ApplyAuthorizedCallerUpdates(opts, args)
	},
})
