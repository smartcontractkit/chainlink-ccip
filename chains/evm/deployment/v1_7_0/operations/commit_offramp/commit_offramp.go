package commit_offramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/operations/call"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/operations/deployment"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/commit_offramp"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CommitOffRamp"

type ConstructorArgs struct {
	NonceManager common.Address
}

type SignatureConfigArgs = commit_offramp.SignatureQuorumVerifierSignatureConfigArgs

var Deploy = deployment.New(
	"commit-offramp:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the SignatureQuorumVerifier contract",
	ContractType,
	deployment.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := commit_offramp.DeployCommitOffRamp(opts, backend, args.NonceManager)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var SetSignatureConfig = call.NewWrite(
	"commit-offramp:set-signature-config",
	semver.MustParse("1.7.0"),
	"Sets the signature configuration on the CommitOffRamp",
	ContractType,
	commit_offramp.NewCommitOffRamp,
	func(commitOffRamp *commit_offramp.CommitOffRamp, opts *bind.TransactOpts, args SignatureConfigArgs) (*types.Transaction, error) {
		return commitOffRamp.SetSignatureConfig(opts, args)
	},
)
