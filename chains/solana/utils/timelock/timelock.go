package timelock

import (
	"context"
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/mcms"
)

// instruction builder for initializing access controllers
func InitAccessControllersIxs(ctx context.Context, roleAcAccount solana.PublicKey, authority solana.PrivateKey, client *rpc.Client) ([]solana.Instruction, error) {
	ixs := []solana.Instruction{}

	rentExemption, err := client.GetMinimumBalanceForRentExemption(
		ctx,
		config.AccSpace,
		config.DefaultCommitment,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get rent exemption: %w", err)
	}

	createAccIx, err := system.NewCreateAccountInstruction(
		rentExemption,
		config.AccSpace,
		config.AccessControllerProgram,
		authority.PublicKey(),
		roleAcAccount,
	).ValidateAndBuild()
	if err != nil {
		return nil, fmt.Errorf("failed to create account instruction: %w", err)
	}
	ixs = append(ixs, createAccIx)

	initIx, err := access_controller.NewInitializeInstruction(
		roleAcAccount,
		authority.PublicKey(),
	).ValidateAndBuild()
	if err != nil {
		return nil, fmt.Errorf("failed to create initialize instruction: %w", err)
	}
	ixs = append(ixs, initIx)

	return ixs, nil
}

// instructions builder for adding access to a role
func BatchAddAccessIxs(ctx context.Context, roleAcAccount solana.PublicKey, role timelock.Role, addresses []solana.PublicKey, authority solana.PrivateKey, chunkSize int, client *rpc.Client) ([]solana.Instruction, error) {
	var ac access_controller.AccessController
	err := common.GetAccountDataBorshInto(ctx, client, roleAcAccount, config.DefaultCommitment, &ac)
	if err != nil {
		return nil, fmt.Errorf("access controller for role %s is not initialized: %w", role, err)
	}
	ixs := []solana.Instruction{}
	for i := 0; i < len(addresses); i += chunkSize {
		end := i + chunkSize
		if end > len(addresses) {
			end = len(addresses)
		}
		chunk := addresses[i:end]
		ix := timelock.NewBatchAddAccessInstruction(
			role,
			config.TimelockConfigPDA,
			config.AccessControllerProgram,
			roleAcAccount,
			authority.PublicKey(),
		)
		for _, address := range chunk {
			ix.Append(solana.Meta(address))
		}
		vIx, err := ix.ValidateAndBuild()
		if err != nil {
			return nil, fmt.Errorf("failed to build instruction for role %v: %w", role, err)
		}
		ixs = append(ixs, vIx)
	}
	return ixs, nil
}

// instructions builder for preloading instructions to timelock operation
func PreloadOperationIxs(ctx context.Context, op Operation, authority solana.PublicKey, client *rpc.Client) ([]solana.Instruction, error) {
	ixs := []solana.Instruction{}
	initOpIx, ioErr := timelock.NewInitializeOperationInstruction(
		op.OperationID(),
		op.Predecessor,
		op.Salt,
		op.IxsCountU32(),
		op.OperationPDA(),
		config.TimelockConfigPDA,
		authority,
		solana.SystemProgramID,
	).ValidateAndBuild()
	if ioErr != nil {
		return nil, fmt.Errorf("failed to build initialize operation instruction: %w", ioErr)
	}
	ixs = append(ixs, initOpIx)

	for _, instruction := range op.ToInstructionData() {
		appendIxsIx, apErr := timelock.NewAppendInstructionsInstruction(
			op.OperationID(),
			[]timelock.InstructionData{instruction}, // this should be a slice of instruction within 1232 bytes
			op.OperationPDA(),
			config.TimelockConfigPDA,
			authority,
			solana.SystemProgramID, // for reallocation
		).ValidateAndBuild()
		if apErr != nil {
			return nil, fmt.Errorf("failed to build append instructions instruction: %w", apErr)
		}
		ixs = append(ixs, appendIxsIx)
	}

	finOpIx, foErr := timelock.NewFinalizeOperationInstruction(
		op.OperationID(),
		op.OperationPDA(),
		config.TimelockConfigPDA,
		authority,
	).ValidateAndBuild()
	if foErr != nil {
		return nil, fmt.Errorf("failed to build finalize operation instruction: %w", foErr)
	}
	ixs = append(ixs, finOpIx)

	return ixs, nil
}

// mcm + timelock test helpers
type RoleMultisigs struct {
	Multisigs        []mcms.Multisig
	AccessController solana.PrivateKey
}

func (r RoleMultisigs) GetAnyMultisig() mcms.Multisig {
	if len(r.Multisigs) == 0 {
		panic("no multisigs to pick from")
	}
	maxCount := big.NewInt(int64(len(r.Multisigs)))
	n, err := crypto_rand.Int(crypto_rand.Reader, maxCount)
	if err != nil {
		panic(err)
	}
	return r.Multisigs[n.Int64()]
}

func CreateRoleMultisigs(role timelock.Role, numMsigs int) RoleMultisigs {
	msigs := make([]mcms.Multisig, numMsigs)
	for i := 0; i < numMsigs; i++ {
		name, _ := mcms.PadString32(fmt.Sprintf("%s_%d", role.String(), i))
		msig := mcms.NewMcmMultisig(name)
		// Create and set the config for each msig
		//       ┌──────┐
		//       │2-of-2│ root
		//       └──────┘
		//          ▲ ▲
		//   group1 │ │  group2
		//  ┌──────┐│ │┌──────┐
		//  │2-of-3│┘ └│1-of-2│
		//  └──────┘   └──────┘
		//   ▲  ▲  ▲      ▲ ▲
		//   │  │  │      │ │
		// ┌─┐ ┌─┐ ┌─┐  ┌─┐ ┌─┐
		// │A│ │B│ │C│  │D│ │E│ signers
		// └─┘ └─┘ └─┘  └─┘ └─┘
		signerPrivateKeys, _ := eth.GenerateEthPrivateKeys(5)
		signerGroups := []byte{1, 1, 1, 2, 2} // A,B,C in group1; D,E in group2
		groupQuorums := []uint8{2, 2, 1}      // root: 2-of-2, group1: 2-of-3, group2: 1-of-2
		groupParents := []uint8{0, 0, 0}      // both groups under root

		mcmConfig, _ := mcms.NewValidMcmConfig(
			name,
			signerPrivateKeys,
			signerGroups,
			groupQuorums,
			groupParents,
			false,
		)

		msig.RawConfig = *mcmConfig
		signers, _ := eth.GetEvmSigners(signerPrivateKeys)
		msig.Signers = signers

		msigs[i] = msig
	}

	acKey, _ := solana.NewRandomPrivateKey()
	return RoleMultisigs{
		Multisigs:        msigs,
		AccessController: acKey,
	}
}

func WaitForOperationToBeReady(ctx context.Context, client *rpc.Client, opPDA solana.PublicKey, commitment rpc.CommitmentType) error {
	const maxAttempts = 20
	const pollInterval = 500 * time.Millisecond
	const timeBuffer = 2 * time.Second

	var opAccount timelock.Operation
	err := common.GetAccountDataBorshInto(ctx, client, opPDA, commitment, &opAccount)
	if err != nil {
		return fmt.Errorf("failed to get account info: %w", err)
	}

	if opAccount.Timestamp == config.TimelockOpDoneTimestamp {
		return nil
	}

	//nolint:gosec
	scheduledTime := time.Unix(int64(opAccount.Timestamp), 0)

	// add buffer to scheduled time to ensure blockchain has advanced enough
	scheduledTimeWithBuffer := scheduledTime.Add(timeBuffer)

	for attempts := 0; attempts < maxAttempts; attempts++ {
		currentTime, err := common.GetBlockTime(ctx, client, commitment)
		if err != nil {
			return fmt.Errorf("failed to get current block time: %w", err)
		}

		if currentTime.Time().After(scheduledTimeWithBuffer) || currentTime.Time().Equal(scheduledTimeWithBuffer) {
			return nil
		}

		time.Sleep(pollInterval)
	}

	return fmt.Errorf("operation not ready after %d attempts (scheduled for: %v, with buffer: %v)",
		maxAttempts, scheduledTime.UTC(), scheduledTimeWithBuffer.UTC())
}

func GetBlockedFunctionSelectors(
	ctx context.Context,
	client *rpc.Client,
	configPubKey solana.PublicKey,
	commitment rpc.CommitmentType,
) ([][]byte, error) {
	var config timelock.Config
	err := common.GetAccountDataBorshInto(ctx, client, configPubKey, commitment, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch config account data: %w", err)
	}

	blockedCount := config.BlockedSelectors.Len
	if blockedCount == 0 {
		return nil, nil
	}

	// convert to [][]byte for easier comparison
	selectors := make([][]byte, blockedCount)
	for i := uint64(0); i < blockedCount; i++ {
		selectors[i] = config.BlockedSelectors.Xs[i][:] // Convert [8]byte to []byte
	}

	return selectors, nil
}
