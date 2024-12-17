package contracts

import (
	"testing"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils"
	mcmsUtils "github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/mcms"
)

func TestTimelockRBAC(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	timelock.SetProgramID(config.TimelockProgram)
	access_controller.SetProgramID(config.AccessControllerProgram)

	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	anotherAdmin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	user, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	roles, roleMap := mcmsUtils.TestRoleAccounts(t, config.NumAccountsPerRole)
	solanaGoClient := utils.DeployAllPrograms(t, utils.PathToAnchorConfig, admin)

	t.Run("setup:funding", func(t *testing.T) {
		all := []solana.PrivateKey{}
		all = append(all, admin, anotherAdmin, user)
		// utils.FundAccounts(ctx, []solana.PrivateKey{admin}, client, t)
		for _, role := range roles {
			all = append(all, role.Accounts...)
		}
		utils.FundAccounts(ctx, all, solanaGoClient, t)
	})

	t.Run("setup:init access controllers", func(t *testing.T) {
		for _, data := range roleMap {
			initAccIxs, err := InitAccessControllersIxs(ctx, data.AccessController.PublicKey(), admin, solanaGoClient)
			require.NoError(t, err)

			utils.SendAndConfirm(ctx, t, solanaGoClient, initAccIxs, admin, config.DefaultCommitment, utils.AddSigners(data.AccessController))

			var ac access_controller.AccessController
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, data.AccessController.PublicKey(), config.DefaultCommitment, &ac)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
		}
	})

	t.Run("fail: NOT able to init program from non-deployer user", func(t *testing.T) {
		// get program data account
		data, accErr := solanaGoClient.GetAccountInfoWithOpts(ctx, config.TimelockProgram, &rpc.GetAccountInfoOpts{
			Commitment: config.DefaultCommitment,
		})
		require.NoError(t, accErr)

		// decode program data
		var programData struct {
			DataType uint32
			Address  solana.PublicKey
		}
		require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

		initTimelockIx, err := timelock.NewInitializeInstruction(
			config.MinDelay,
			config.TimelockConfigPDA,
			anotherAdmin.PublicKey(),
			solana.SystemProgramID,
			config.TimelockProgram,
			programData.Address,
			config.AccessControllerProgram,
			roleMap[timelock.Proposer_Role].AccessController.PublicKey(),
			roleMap[timelock.Executor_Role].AccessController.PublicKey(),
			roleMap[timelock.Canceller_Role].AccessController.PublicKey(),
			roleMap[timelock.Bypasser_Role].AccessController.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)

		result := utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{initTimelockIx}, anotherAdmin, config.DefaultCommitment, []string{"Error Code: " + timelock.Unauthorized_TimelockError.String()})
		require.NotNil(t, result)
	})

	t.Run("setup:init timelock program", func(t *testing.T) {
		// get program data account
		data, accErr := solanaGoClient.GetAccountInfoWithOpts(ctx, config.TimelockProgram, &rpc.GetAccountInfoOpts{
			Commitment: config.DefaultCommitment,
		})
		require.NoError(t, accErr)

		// decode program data
		var programData struct {
			DataType uint32
			Address  solana.PublicKey
		}
		require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

		initTimelockIx, err := timelock.NewInitializeInstruction(
			config.MinDelay,
			config.TimelockConfigPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
			config.TimelockProgram,
			programData.Address,
			config.AccessControllerProgram,
			roleMap[timelock.Proposer_Role].AccessController.PublicKey(),
			roleMap[timelock.Executor_Role].AccessController.PublicKey(),
			roleMap[timelock.Canceller_Role].AccessController.PublicKey(),
			roleMap[timelock.Bypasser_Role].AccessController.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)

		utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initTimelockIx}, admin, config.DefaultCommitment)

		var configAccount timelock.Config
		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, config.TimelockConfigPDA, config.DefaultCommitment, &configAccount)
		if err != nil {
			require.NoError(t, err, "failed to get account info")
		}

		require.Equal(t, admin.PublicKey(), configAccount.Owner, "Owner doesn't match")
		require.Equal(t, config.MinDelay, configAccount.MinDelay, "MinDelay doesn't match")
		require.Equal(t, roleMap[timelock.Proposer_Role].AccessController.PublicKey(), configAccount.ProposerRoleAccessController, "ProposerRoleAccessController doesn't match")
		require.Equal(t, roleMap[timelock.Executor_Role].AccessController.PublicKey(), configAccount.ExecutorRoleAccessController, "ExecutorRoleAccessController doesn't match")
		require.Equal(t, roleMap[timelock.Canceller_Role].AccessController.PublicKey(), configAccount.CancellerRoleAccessController, "CancellerRoleAccessController doesn't match")
		require.Equal(t, roleMap[timelock.Bypasser_Role].AccessController.PublicKey(), configAccount.BypasserRoleAccessController, "BypasserRoleAccessController doesn't match")
	})

	t.Run("timelock:ownership", func(t *testing.T) {
		// Fail to transfer ownership when not owner
		instruction, err := timelock.NewTransferOwnershipInstruction(
			anotherAdmin.PublicKey(),
			config.TimelockConfigPDA,
			user.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result := utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + timelock.Unauthorized_TimelockError.String()})
		require.NotNil(t, result)

		// successfully transfer ownership
		instruction, err = timelock.NewTransferOwnershipInstruction(
			anotherAdmin.PublicKey(),
			config.TimelockConfigPDA,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment)
		require.NotNil(t, result)

		// Fail to accept ownership when not proposed_owner
		instruction, err = timelock.NewAcceptOwnershipInstruction(
			config.TimelockConfigPDA,
			user.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + timelock.Unauthorized_TimelockError.String()})
		require.NotNil(t, result)

		// Successfully accept ownership
		// anotherAdmin becomes owner for remaining tests
		instruction, err = timelock.NewAcceptOwnershipInstruction(
			config.TimelockConfigPDA,
			anotherAdmin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment)
		require.NotNil(t, result)

		// Current owner cannot propose self
		instruction, err = timelock.NewTransferOwnershipInstruction(
			anotherAdmin.PublicKey(),
			config.TimelockConfigPDA,
			anotherAdmin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment, []string{"Error Code: " + timelock.InvalidInput_TimelockError.String()})
		require.NotNil(t, result)

		// Validate proposed set to 0-address after accepting ownership
		var configAccount timelock.Config
		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, config.TimelockConfigPDA, config.DefaultCommitment, &configAccount)
		if err != nil {
			require.NoError(t, err, "failed to get account info")
		}
		require.Equal(t, anotherAdmin.PublicKey(), configAccount.Owner)
		require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)

		// get it back
		instruction, err = timelock.NewTransferOwnershipInstruction(
			admin.PublicKey(),
			config.TimelockConfigPDA,
			anotherAdmin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment)
		require.NotNil(t, result)

		instruction, err = timelock.NewAcceptOwnershipInstruction(
			config.TimelockConfigPDA,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment)
		require.NotNil(t, result)

		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, config.TimelockConfigPDA, config.DefaultCommitment, &configAccount)
		if err != nil {
			require.NoError(t, err, "failed to get account info")
		}

		require.Equal(t, admin.PublicKey(), configAccount.Owner)
		require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
	})

	t.Run("setup:register access list & verify", func(t *testing.T) {
		for role, data := range roleMap {
			addresses := []solana.PublicKey{}
			for _, account := range data.Accounts {
				addresses = append(addresses, account.PublicKey())
			}
			batchAddAccessIxs, err := TimelockBatchAddAccessIxs(ctx, data.AccessController.PublicKey(), role, addresses, admin, config.BatchAddAccessChunkSize, solanaGoClient)
			require.NoError(t, err)

			for _, ix := range batchAddAccessIxs {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			var ac access_controller.AccessController
			err = utils.GetAccountDataBorshInto(
				ctx,
				solanaGoClient,
				data.AccessController.PublicKey(),
				config.DefaultCommitment,
				&ac,
			)
			require.NoError(t, err)

			require.Equal(t, uint64(len(data.Accounts)), ac.AccessList.Len,
				"AccessList length mismatch for %s", data.Role)

			for _, account := range data.Accounts {
				targetPubKey := account.PublicKey()
				_, found := mcmsUtils.FindInSortedList(ac.AccessList.Xs[:ac.AccessList.Len], targetPubKey)
				require.True(t, found, "Account %s not found in %s AccessList",
					targetPubKey, data.Role)
			}
		}
	})

	t.Run("rbac: schedule and cancel a timelock operation", func(t *testing.T) {
		salt, err := mcmsUtils.SimpleSalt()
		require.NoError(t, err)
		nonExecutableOp := TimelockOperation{
			Predecessor: config.TimelockEmptyOpID,
			Salt:        salt,
			Delay:       uint64(1),
		}

		ix := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), config.TimelockSignerPDA).Build()
		nonExecutableOp.AddInstruction(ix, []solana.PublicKey{})

		id := nonExecutableOp.OperationID()
		operationPDA := nonExecutableOp.OperationPDA()

		t.Run("rbac: Should able to schedule tx with proposer role", func(t *testing.T) {
			signer := roleMap[timelock.Proposer_Role].RandomPick()
			ac := roleMap[timelock.Proposer_Role].AccessController

			ixs := make([]solana.Instruction, 0)
			initOpIx, err := timelock.NewInitializeOperationInstruction(
				nonExecutableOp.OperationID(),
				nonExecutableOp.Predecessor,
				nonExecutableOp.Salt,
				uint32(len(nonExecutableOp.instructions)),
				config.TimelockConfigPDA,
				operationPDA,
				signer.PublicKey(),
				signer.PublicKey(), // proposer - direct schedule batch here
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			ixs = append(ixs, initOpIx)

			appendIxIx, err := timelock.NewAppendInstructionsInstruction(
				nonExecutableOp.OperationID(),
				nonExecutableOp.ToInstructionData(),
				operationPDA,
				signer.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			ixs = append(ixs, appendIxIx)

			finIxIx, err := timelock.NewFinalizeOperationInstruction(
				nonExecutableOp.OperationID(),
				operationPDA,
				signer.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			ixs = append(ixs, finIxIx)

			utils.SendAndConfirm(ctx, t, solanaGoClient, ixs, signer, config.DefaultCommitment)

			ix, err := timelock.NewScheduleBatchInstruction(
				nonExecutableOp.OperationID(),
				nonExecutableOp.Delay,
				config.TimelockConfigPDA,
				operationPDA,
				ac.PublicKey(),
				signer.PublicKey(),
			).ValidateAndBuild()

			require.NoError(t, err)

			result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment)
			require.NotNil(t, result)

			var opAccount timelock.Operation
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, operationPDA, config.DefaultCommitment, &opAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}

			require.Equal(t,
				result.BlockTime.Time().Add(time.Duration(nonExecutableOp.Delay)*time.Second).Unix(),
				int64(opAccount.Timestamp),
				"Scheduled Times don't match",
			)
			require.Equal(t,
				id,
				opAccount.Id,
				"Ids don't match",
			)
		})

		t.Run("rbac: cancel scheduled tx", func(t *testing.T) {
			t.Run("fail: should feed the right role access controller", func(t *testing.T) {
				signer := roleMap[timelock.Canceller_Role].RandomPick()
				ac := roleMap[timelock.Proposer_Role].AccessController

				ix, err := timelock.NewCancelInstruction(
					id,
					config.TimelockConfigPDA,
					operationPDA,
					ac.PublicKey(),
					signer.PublicKey(),
				).ValidateAndBuild()

				require.NoError(t, err)

				result := utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment, []string{"Error Code: " + "InvalidAccessController."})
				require.NotNil(t, result)
			})

			t.Run("fail: unauthorized on cancel attempt from non-canceller(proposer)", func(t *testing.T) {
				signer := roleMap[timelock.Proposer_Role].RandomPick()
				ac := roleMap[timelock.Canceller_Role].AccessController

				ix, err := timelock.NewCancelInstruction(
					id,
					config.TimelockConfigPDA,
					operationPDA,
					ac.PublicKey(),
					signer.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				result := utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment, []string{"Error Code: " + timelock.Unauthorized_TimelockError.String()})
				require.NotNil(t, result)
			})

			t.Run("success: Should able to cancel scheduled tx: PDA closed", func(t *testing.T) {
				signer := roleMap[timelock.Canceller_Role].RandomPick()
				ac := roleMap[timelock.Canceller_Role].AccessController

				ix, err := timelock.NewCancelInstruction(
					id,

					config.TimelockConfigPDA,
					operationPDA,

					ac.PublicKey(),
					signer.PublicKey(),
				).ValidateAndBuild()

				require.NoError(t, err)

				result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment)
				require.NotNil(t, result)

				utils.AssertClosedAccount(ctx, t, solanaGoClient, operationPDA, config.DefaultCommitment)
			})
		})
	})

	t.Run("rbac: only_admin authentication ", func(t *testing.T) {
		newMinDelay := uint64(14000)

		t.Run("fail: only admin can call functions with only_admin macro", func(t *testing.T) {
			signer := roleMap[timelock.Proposer_Role].RandomPick()

			ix, err := timelock.NewUpdateDelayInstruction(
				newMinDelay,
				config.TimelockConfigPDA,
				signer.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			result := utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment, []string{"Error Code: " + timelock.Unauthorized_TimelockError.String()})
			require.NotNil(t, result)
		})

		t.Run("success: only admin can call functions with only_admin macro", func(t *testing.T) {
			signer := admin

			ix, err := timelock.NewUpdateDelayInstruction(
				newMinDelay,
				config.TimelockConfigPDA,
				signer.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment)
			require.NotNil(t, result)

			var configAccount timelock.Config
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, config.TimelockConfigPDA, config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
			require.Equal(t, newMinDelay, configAccount.MinDelay, "MinDelay is not updated")
		})
	})
}
