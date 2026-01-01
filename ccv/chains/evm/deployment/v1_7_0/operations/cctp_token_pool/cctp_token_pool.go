package cctp_token_pool

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CCTPTokenPool"

var Version = semver.MustParse("1.7.0")

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
	Version:          Version,
	Description:      "Deploys the CCTPTokenPool contract",
	ContractMetadata: cctp_token_pool.CCTPTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(cctp_token_pool.CCTPTokenPoolBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *cctp_token_pool.CCTPTokenPool]{
	Name:            "cctp-token-pool:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Applies authorized caller updates on the CCTPTokenPool",
	ContractType:    ContractType,
	ContractABI:     cctp_token_pool.CCTPTokenPoolABI,
	NewContract:     cctp_token_pool.NewCCTPTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*cctp_token_pool.CCTPTokenPool, AuthorizedCallerArgs],
	Validate: func(cctpTokenPool *cctp_token_pool.CCTPTokenPool, backend bind.ContractBackend, opts *bind.CallOpts, args AuthorizedCallerArgs) error {
		for _, caller := range args.AddedCallers {
			if caller == (common.Address{}) {
				return errors.New("caller cannot be the zero address")
			}
		}
		return nil
	},
	IsNoop: func(cctpTokenPool *cctp_token_pool.CCTPTokenPool, opts *bind.CallOpts, args AuthorizedCallerArgs) (bool, error) {
		allowedCallers, err := cctpTokenPool.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get all authorized callers: %w", err)
		}
		for _, caller := range args.AddedCallers {
			if !slices.Contains(allowedCallers, caller) {
				return false, nil
			}
		}
		for _, caller := range args.RemovedCallers {
			if slices.Contains(allowedCallers, caller) {
				return false, nil
			}
		}
		return true, nil
	},
	CallContract: func(cctpTokenPool *cctp_token_pool.CCTPTokenPool, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return cctpTokenPool.ApplyAuthorizedCallerUpdates(opts, args)
	},
})
