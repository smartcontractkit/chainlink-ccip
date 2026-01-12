package siloed_usdc_token_pool

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/siloed_usdc_token_pool"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "SiloedUSDCTokenPool"
var Version *semver.Version = semver.MustParse("1.6.4")

type ConstructorArgs struct {
	Token              common.Address
	LocalTokenDecimals uint8
	Allowlist          []common.Address
	RMNProxy           common.Address
	Router             common.Address
	LockBox            common.Address
}

type AuthorizedCallerUpdateArgs = siloed_usdc_token_pool.AuthorizedCallersAuthorizedCallerArgs

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "siloed-usdc-token-pool:deploy",
	Version:          Version,
	Description:      "Deploys the SiloedUSDCTokenPool contract",
	ContractMetadata: siloed_usdc_token_pool.SiloedUSDCTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(siloed_usdc_token_pool.SiloedUSDCTokenPoolBin),
		},
	},
	Validate: func(args ConstructorArgs) error {
		if args.Token == (common.Address{}) {
			return errors.New("token address cannot be zero")
		}

		if args.RMNProxy == (common.Address{}) {
			return errors.New("rmnProxy address cannot be zero")
		}

		if args.Router == (common.Address{}) {
			return errors.New("router address cannot be zero")
		}

		if args.LockBox == (common.Address{}) {
			return errors.New("lockBox address cannot be zero")
		}

		if args.LocalTokenDecimals != 6 {
			return fmt.Errorf("local token decimals must be 6")
		}
		return nil
	},
})
