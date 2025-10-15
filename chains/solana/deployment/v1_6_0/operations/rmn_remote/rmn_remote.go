package rmn_remote

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ContractType cldf_deployment.ContractType = "RMNRemote"
var ProgramName = "rmn_remote"
var ProgramSize = 3 * 1024 * 1024
var Version *semver.Version = semver.MustParse("1.6.0")

var Deploy = operations.NewOperation(
	"rmn-remote:deploy",
	Version,
	"Deploys the RMNRemote program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			ContractType,
			Version,
			"",
			ProgramName,
			ProgramSize)
	},
)

var Initialize = operations.NewOperation(
	"rmn-remote:initialize",
	Version,
	"Initializes the RMNRemote 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		programData, err := utils.GetSolProgramData(chain, input.RMNRemote)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get program data: %w", err)
		}
		rmnRemoteConfigPDA, _, _ := state.FindRMNRemoteConfigPDA(input.RMNRemote)
		rmnRemoteCursesPDA, _, _ := state.FindRMNRemoteCursesPDA(input.RMNRemote)
		instruction, err := rmn_remote.NewInitializeInstruction(
			rmnRemoteConfigPDA,
			rmnRemoteCursesPDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
			input.RMNRemote,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm initialization: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

type Params struct {
	RMNRemote solana.PublicKey
}
