package rmn_remote

import (
	"context"
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
	"github.com/smartcontractkit/mcms/types"
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
		rmn_remote.SetProgramID(input.RMNRemote)
		programData, err := utils.GetSolProgramData(chain, input.RMNRemote)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get program data: %w", err)
		}
		authority := GetAuthority(chain, input.RMNRemote)
		rmnRemoteConfigPDA, _, _ := state.FindRMNRemoteConfigPDA(input.RMNRemote)
		rmnRemoteCursesPDA, _, _ := state.FindRMNRemoteCursesPDA(input.RMNRemote)
		instruction, err := rmn_remote.NewInitializeInstruction(
			rmnRemoteConfigPDA,
			rmnRemoteCursesPDA,
			authority,
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

var TransferOwnership = operations.NewOperation(
	"rmn-remote:transfer-ownership",
	Version,
	"Transfers ownership of the RMNRemote 1.6.0 contract to a new authority",
	func(b operations.Bundle, chain cldf_solana.Chain, input utils.TransferOwnershipParams) (sequences.OnChainOutput, error) {
		rmn_remote.SetProgramID(input.Program)
		authority := GetAuthority(chain, input.Program)
		if authority != input.CurrentOwner {
			return sequences.OnChainOutput{}, fmt.Errorf("current owner %s does not match on-chain authority %s", input.CurrentOwner.String(), authority.String())
		}
		configPDA, _, _ := state.FindConfigPDA(input.Program)
		cursePDA, _, _ := state.FindRMNRemoteCursesPDA(input.Program)
		ixn, err := rmn_remote.NewTransferOwnershipInstruction(
			input.NewOwner,
			configPDA,
			cursePDA,
			authority,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build add dest chain instruction: %w", err)
		}
		if authority != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Program.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add price updater: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var AcceptOwnership = operations.NewOperation(
	"rmn-remote:accept-ownership",
	Version,
	"Accepts ownership of the RMNRemote 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input utils.TransferOwnershipParams) (sequences.OnChainOutput, error) {
		rmn_remote.SetProgramID(input.Program)
		configPDA, _, _ := state.FindConfigPDA(input.Program)
		ixn, err := rmn_remote.NewAcceptOwnershipInstruction(
			configPDA,
			input.NewOwner,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build add dest chain instruction: %w", err)
		}
		if input.NewOwner != chain.DeployerKey.PublicKey() {
			batches, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				[]solana.Instruction{ixn},
				input.Program.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			return sequences.OnChainOutput{BatchOps: []types.BatchOperation{batches}}, nil
		}

		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add price updater: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

func GetAuthority(chain cldf_solana.Chain, program solana.PublicKey) solana.PublicKey {
	programData := rmn_remote.Config{}
	rmnRemoteConfigPDA, _, _ := state.FindRMNRemoteConfigPDA(program)
	err := chain.GetAccountDataBorshInto(context.Background(), rmnRemoteConfigPDA, &programData)
	if err != nil {
		return chain.DeployerKey.PublicKey()
	}
	return programData.Owner
}

type Params struct {
	RMNRemote solana.PublicKey
}
