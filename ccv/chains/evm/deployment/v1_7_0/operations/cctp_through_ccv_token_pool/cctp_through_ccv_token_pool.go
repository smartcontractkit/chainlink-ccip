package cctp_through_ccv_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CCTPThroughCCVTokenPool"

var Version = semver.MustParse("1.7.0")

type ConstructorArgs struct {
	Token              common.Address
	LocalTokenDecimals uint8
	RMNProxy           common.Address
	Router             common.Address
	CCTPVerifier       common.Address
	AllowedCallers     []common.Address
}

type AuthorizedCallerArgs = cctp_through_ccv_token_pool.AuthorizedCallersAuthorizedCallerArgs

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "cctp-through-ccv-token-pool:deploy",
	Version:          Version,
	Description:      "Deploys the CCTPThroughCCVTokenPool contract",
	ContractMetadata: cctp_through_ccv_token_pool.CCTPThroughCCVTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(cctp_through_ccv_token_pool.CCTPThroughCCVTokenPoolBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *cctp_through_ccv_token_pool.CCTPThroughCCVTokenPool]{
	Name:            "cctp-through-ccv-token-pool:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Applies authorized caller updates on the CCTPThroughCCVTokenPool",
	ContractType:    ContractType,
	ContractABI:     cctp_through_ccv_token_pool.CCTPThroughCCVTokenPoolABI,
	NewContract:     cctp_through_ccv_token_pool.NewCCTPThroughCCVTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*cctp_through_ccv_token_pool.CCTPThroughCCVTokenPool, AuthorizedCallerArgs],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(cctpThroughCCVTokenPool *cctp_through_ccv_token_pool.CCTPThroughCCVTokenPool, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return cctpThroughCCVTokenPool.ApplyAuthorizedCallerUpdates(opts, args)
	},
})
