package rmn_remote

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/jsonrpc"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/rmn_remote"
	commonUtil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var ContractType cldf_deployment.ContractType = "RMNRemote"
var ProgramName = "rmn_remote"
var ProgramSize = 3 * 1024 * 1024
var Version *semver.Version = semver.MustParse("1.6.0")

type CurseInput struct {
	Subjects           []api.Subject
	RMNRemoteCursePDA  solana.PublicKey
	RMNRemoteConfigPDA solana.PublicKey
	RMNRemote          solana.PublicKey
}

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
		programData, err := utils.GetSolProgramData(chain.Client, input.RMNRemote)
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
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm ownership transfer: %w", err)
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
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm ownership acceptance: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var Curse = operations.NewOperation(
	"rmn-remote:curse",
	Version,
	"Curses subjects with RMNRemote",
	func(b operations.Bundle, chain cldf_solana.Chain, input CurseInput) (sequences.OnChainOutput, error) {
		rmnRemoteConfigPDA := input.RMNRemoteConfigPDA
		rmnRemoteCursesPDA := input.RMNRemoteCursePDA
		rmn_remote.SetProgramID(input.RMNRemote)
		authority := GetAuthority(chain, input.RMNRemote)
		ins := make([]solana.Instruction, 0)
		for _, subject := range input.Subjects {
			curseSubject := rmn_remote.CurseSubject{
				Value: subject,
			}
			ix, err := rmn_remote.NewCurseInstruction(
				curseSubject,
				rmnRemoteConfigPDA,
				authority,
				rmnRemoteCursesPDA,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build curse instruction: %w", err)
			}
			ins = append(ins, ix)
		}
		batches := make([]types.BatchOperation, 0)
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				ins,
				input.RMNRemote.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			batches = append(batches, b)
		} else {
			for _, ixn := range ins {
				err := chain.Confirm([]solana.Instruction{ixn})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm curse instruction: %w", err)
				}
			}
		}
		return sequences.OnChainOutput{BatchOps: batches}, nil
	},
)

var Uncurse = operations.NewOperation(
	"rmn-remote:uncurse",
	Version,
	"Lifts curses for subjects with RMNRemote",
	func(b operations.Bundle, chain cldf_solana.Chain, input CurseInput) (sequences.OnChainOutput, error) {
		rmnRemoteConfigPDA := input.RMNRemoteConfigPDA
		rmnRemoteCursesPDA := input.RMNRemoteCursePDA
		rmn_remote.SetProgramID(input.RMNRemote)
		authority := GetAuthority(chain, input.RMNRemote)
		ins := make([]solana.Instruction, 0)
		for _, subject := range input.Subjects {
			curseSubject := rmn_remote.CurseSubject{
				Value: subject,
			}
			ix, err := rmn_remote.NewUncurseInstruction(
				curseSubject,
				rmnRemoteConfigPDA,
				authority,
				rmnRemoteCursesPDA,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build uncurse instruction: %w", err)
			}
			ins = append(ins, ix)
		}
		batches := make([]types.BatchOperation, 0)
		if authority != chain.DeployerKey.PublicKey() {
			b, err := utils.BuildMCMSBatchOperation(
				chain.Selector,
				ins,
				input.RMNRemote.String(),
				ContractType.String(),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute or create batch: %w", err)
			}
			batches = append(batches, b)
		} else {
			for _, ixn := range ins {
				err := chain.Confirm([]solana.Instruction{ixn})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm uncurse instruction: %w", err)
				}
			}
		}
		return sequences.OnChainOutput{BatchOps: batches}, nil
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

func IsSubjectCursed(chain cldf_solana.Chain, program solana.PublicKey, subject rmn_remote.CurseSubject) (bool, error) {
	rmnRemoteConfigPDA, _, err := state.FindRMNRemoteConfigPDA(program)
	if err != nil {
		return false, fmt.Errorf("failed to find RMNRemoteConfig PDA: %w", err)
	}
	rmnRemoteCursesPDA, _, err := state.FindRMNRemoteCursesPDA(program)
	if err != nil {
		return false, fmt.Errorf("failed to find RMNRemoteCurses PDA: %w", err)
	}
	ix, err := rmn_remote.NewVerifyNotCursedInstruction(
		subject,
		rmnRemoteCursesPDA,
		rmnRemoteConfigPDA,
	).ValidateAndBuild()
	if err != nil {
		return false, fmt.Errorf("failed to generate instructions: %w", err)
	}
	data, err := ix.Data()
	if err != nil {
		return false, fmt.Errorf("failed to extract data payload from verify not cursed instruction: %w", err)
	}
	// Manually create instruction rather than directly using the ix above
	// Using the ix above requires setting the program ID in the binding directly which panics if called multiple times
	verifyIx := solana.NewInstruction(program, ix.Accounts(), data)
	_, txErr := commonUtil.SendAndConfirmWithLookupTables(context.Background(), chain.Client, []solana.Instruction{verifyIx}, *chain.DeployerKey, rpc.CommitmentConfirmed, nil)
	if txErr == nil {
		// If no error return then it's not cursed
		return false, nil
	}
	// Curse types are returned as errors.
	// ref: https://github.com/smartcontractkit/chainlink-ccip/blob/solana-v1.6.0/chains/solana/contracts/target/idl/rmn_remote.json#L478-L485
	curseType, err := parseSolanaErrorCode(txErr)
	if err != nil {
		return false, fmt.Errorf("failed to parse solana error code: %w", err)
	}
	switch curseType {
	case 9006: // global curse
		globalCurse := rmn_remote.CurseSubject{
			Value: api.GlobalCurseSubject(),
		}
		if subject == globalCurse {
			return true, nil
		}
		return false, fmt.Errorf("unexpected global curse for non-global subject")
	case 9005: // address curse
		return true, nil
	default:
		return false, fmt.Errorf("unexpected error code returned from RMNRemote: %d", curseType)
	}
}

type Params struct {
	RMNRemote solana.PublicKey
}

func parseSolanaErrorCode(err error) (int64, error) {
	var rpcErr *jsonrpc.RPCError
	if !errors.As(err, &rpcErr) {
		return 0, fmt.Errorf("not a jsonrpc.RPCError: %w", err)
	}

	data, ok := rpcErr.Data.(map[string]any)
	if !ok {
		return 0, fmt.Errorf("invalid data format: %w", err)
	}

	errData, ok := data["err"].(map[string]any)
	if !ok {
		return 0, fmt.Errorf("no err field found: %w", err)
	}

	instrErr, ok := errData["InstructionError"].([]any)
	if !ok || len(instrErr) < 2 {
		return 0, fmt.Errorf("invalid InstructionError format: %w", err)
	}

	customErr, ok := instrErr[1].(map[string]any)
	if !ok {
		return 0, fmt.Errorf("invalid custom error format: %w", err)
	}

	custom, ok := customErr["Custom"].(json.Number)
	if !ok {
		return 0, fmt.Errorf("no Custom field found: %w", err)
	}

	errorCode, err := custom.Int64()
	if err != nil {
		return 0, fmt.Errorf("failed to parse custom error number: %w", err)
	}

	return errorCode, nil
}
