package commit_offramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ccv_ramp_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/commit_offramp"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CommitOffRamp"

var ProxyType cldf_deployment.ContractType = "CommitOffRampProxy"

type ConstructorArgs struct {
	NonceManager common.Address
}

type SignatureConfigArgs = commit_offramp.SignatureQuorumVerifierSignatureConfigArgs

var Deploy = contract.NewDeploy(
	"commit-offramp:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the SignatureQuorumVerifier contract",
	ContractType,
	commit_offramp.CommitOffRampABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := commit_offramp.DeployCommitOffRamp(opts, backend, args.NonceManager)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var DeployProxy = contract.NewDeploy(
	"commit-on-ramp-proxy:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CommitOnRampProxy contract",
	ProxyType,
	ccv_ramp_proxy.CCVRampProxyABI,
	func(any) error { return nil },
	contract.VMDeployers[any]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args any) (common.Address, *types.Transaction, error) {
			address, tx, _, err := ccv_ramp_proxy.DeployCCVRampProxy(opts, backend)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args any) (common.Address, error)
	},
)

var SetSignatureConfigs = contract.NewWrite(
	"commit-offramp:set-signature-config",
	semver.MustParse("1.7.0"),
	"Sets the signature configuration on the CommitOffRamp",
	ContractType,
	commit_offramp.CommitOffRampABI,
	commit_offramp.NewCommitOffRamp,
	contract.OnlyOwner,
	func(SignatureConfigArgs) error { return nil },
	func(commitOffRamp *commit_offramp.CommitOffRamp, opts *bind.TransactOpts, args SignatureConfigArgs) (*types.Transaction, error) {
		return commitOffRamp.SetSignatureConfig(opts, args)
	},
)
