package contracts

import (
	"strconv"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/accesscontroller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	timelockutil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestTimelockBypasserExecute(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	timelock.SetProgramID(config.TimelockProgram)
	access_controller.SetProgramID(config.AccessControllerProgram)

	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	roles, roleMap := timelockutil.TestRoleAccounts(config.NumAccountsPerRole)
	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)

	allowance := struct {
		timelockAuthority uint64
		recipient         uint64
	}{
		timelockAuthority: 5 * solana.LAMPORTS_PER_SOL,
		recipient:         1 * solana.LAMPORTS_PER_SOL,
	}

	tokenProgram := solana.TokenProgramID
	wsol := solana.WrappedSol
	wsolDecimal := uint8(9)

	adminATA, _, err := solana.FindAssociatedTokenAddress(admin.PublicKey(), wsol)
	require.NoError(t, err)

	recipient, kerr := solana.NewRandomPrivateKey()
	require.NoError(t, kerr)

	recipientATA, _, err := solana.FindAssociatedTokenAddress(recipient.PublicKey(), wsol)
	require.NoError(t, err)

	t.Run("setup:funding", func(t *testing.T) {
		all := []solana.PrivateKey{}
		all = append(all, admin)
		// utils.FundAccounts(ctx, []solana.PrivateKey{admin}, client, t)
		for _, role := range roles {
			all = append(all, role.Accounts...)
		}
		testutils.FundAccounts(ctx, all, solanaGoClient, t)
	})

	t.Run("setup:init access controllers", func(t *testing.T) {
		for _, data := range roleMap {
			initAccIxs, initAccIxsErr := timelockutil.GetInitAccessControllersIxs(ctx, data.AccessController.PublicKey(), admin, solanaGoClient)
			require.NoError(t, initAccIxsErr)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, initAccIxs, admin, config.DefaultCommitment, common.AddSigners(data.AccessController))

			var ac access_controller.AccessController
			acAccErr := common.GetAccountDataBorshInto(ctx, solanaGoClient, data.AccessController.PublicKey(), config.DefaultCommitment, &ac)
			if acAccErr != nil {
				require.NoError(t, acAccErr, "failed to get account info")
			}
		}
	})

	t.Run("setup:initialize timelock instance", func(t *testing.T) {
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

		initTimelockIx, initErr := timelock.NewInitializeInstruction(
			config.TestTimelockID,
			config.MinDelay,
			timelockutil.GetConfigPDA(config.TestTimelockID),
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
		require.NoError(t, initErr)

		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initTimelockIx}, admin, config.DefaultCommitment)

		var configAccount timelock.Config
		cfgErr := common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment, &configAccount)
		if cfgErr != nil {
			require.NoError(t, cfgErr, "failed to get account info")
		}

		require.Equal(t, admin.PublicKey(), configAccount.Owner, "Owner doesn't match")
		require.Equal(t, config.MinDelay, configAccount.MinDelay, "MinDelay doesn't match")
		require.Equal(t, roleMap[timelock.Proposer_Role].AccessController.PublicKey(), configAccount.ProposerRoleAccessController, "ProposerRoleAccessController doesn't match")
		require.Equal(t, roleMap[timelock.Executor_Role].AccessController.PublicKey(), configAccount.ExecutorRoleAccessController, "ExecutorRoleAccessController doesn't match")
		require.Equal(t, roleMap[timelock.Canceller_Role].AccessController.PublicKey(), configAccount.CancellerRoleAccessController, "CancellerRoleAccessController doesn't match")
		require.Equal(t, roleMap[timelock.Bypasser_Role].AccessController.PublicKey(), configAccount.BypasserRoleAccessController, "BypasserRoleAccessController doesn't match")
	})

	t.Run("setup:register access list & verify", func(t *testing.T) {
		for role, data := range roleMap {
			addresses := []solana.PublicKey{}
			for _, account := range data.Accounts {
				addresses = append(addresses, account.PublicKey())
			}
			batchAddAccessIxs, batchAddAccessIxsErr := timelockutil.GetBatchAddAccessIxs(ctx, config.TestTimelockID, data.AccessController.PublicKey(), role, addresses, admin, config.BatchAddAccessChunkSize, solanaGoClient)
			require.NoError(t, batchAddAccessIxsErr)

			for _, ix := range batchAddAccessIxs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			for _, account := range data.Accounts {
				found, ferr := accesscontroller.HasAccess(ctx, solanaGoClient, data.AccessController.PublicKey(), account.PublicKey(), config.DefaultCommitment)
				require.NoError(t, ferr)
				require.True(t, found, "Account %s not found in %s AccessList", account.PublicKey(), data.Role)
			}
		}
	})

	t.Run("schedule_execute: multiple wsol transfer operation", func(t *testing.T) {
		t.Run("setup: wsol transfer operation", func(t *testing.T) {
			requiredAmount := allowance.recipient

			fundPDAIx := system.NewTransferInstruction(allowance.timelockAuthority, admin.PublicKey(), timelockutil.GetSignerPDA(config.TestTimelockID)).Build()

			createAdminATAIx, _, caErr := tokens.CreateAssociatedTokenAccount(tokenProgram, wsol, admin.PublicKey(), admin.PublicKey())
			require.NoError(t, caErr)

			wrapSolIx, wsErr := system.NewTransferInstruction(
				requiredAmount,
				admin.PublicKey(),
				adminATA,
			).ValidateAndBuild()
			require.NoError(t, wsErr)

			syncNativeIx, snErr := tokens.SyncNative(
				tokenProgram,
				adminATA, // token account
			)
			require.NoError(t, snErr)

			// approve can't be deligated to timelock authority(CPI Guard)
			approveIx, aiErr := tokens.TokenApproveChecked(
				requiredAmount,
				wsolDecimal,
				tokenProgram,
				adminATA,
				wsol,
				timelockutil.GetSignerPDA(config.TestTimelockID),
				admin.PublicKey(),
				nil,
			)
			require.NoError(t, aiErr)

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient,
				[]solana.Instruction{createAdminATAIx, wrapSolIx, syncNativeIx, fundPDAIx, approveIx},
				admin, config.DefaultCommitment)
			require.NotNil(t, result)

			// check results
			timelockAuthorityBalance, tlBalanceErr := solanaGoClient.GetBalance(
				ctx,
				timelockutil.GetSignerPDA(config.TestTimelockID),
				config.DefaultCommitment,
			)
			require.NoError(t, tlBalanceErr)
			require.Equal(t, allowance.timelockAuthority, timelockAuthorityBalance.Value)

			adminWsolBalance, adminATABalanceErr := solanaGoClient.GetTokenAccountBalance(
				ctx,
				adminATA,
				config.DefaultCommitment,
			)
			require.NoError(t, adminATABalanceErr)
			require.Equal(t, strconv.Itoa(int(requiredAmount)), adminWsolBalance.Value.Amount)
		})

		t.Run("success: preload and execute batch instructions via bypasser_execute_batch", func(t *testing.T) {
			salt, err := timelockutil.SimpleSalt()
			require.NoError(t, err)
			op := timelockutil.Operation{
				TimelockID:   config.TestTimelockID,
				Predecessor:  config.TimelockEmptyOpID,
				Salt:         salt,
				Delay:        uint64(1000),
				IsBypasserOp: true,
			}

			cIx, _, ciErr := tokens.CreateAssociatedTokenAccount(
				tokenProgram,
				wsol,
				recipient.PublicKey(),
				timelockutil.GetSignerPDA(config.TestTimelockID),
			)
			require.NoError(t, ciErr)
			op.AddInstruction(
				cIx,
				[]solana.PublicKey{tokenProgram, solana.SPLAssociatedTokenAccountProgramID},
			)

			tIx, tiErr := tokens.TokenTransferChecked(
				allowance.recipient,
				wsolDecimal,
				tokenProgram,
				adminATA,
				wsol,
				recipientATA,
				timelockutil.GetSignerPDA(config.TestTimelockID),
				nil,
			)
			require.NoError(t, tiErr)
			op.AddInstruction(tIx, []solana.PublicKey{tokenProgram})

			id := op.OperationID()
			operationPDA := op.OperationPDA()
			signer := roleMap[timelock.Bypasser_Role].RandomPick()
			ac := roleMap[timelock.Bypasser_Role].AccessController

			ixs, err := timelockutil.GetPreloadBypasserOperationIxs(config.TestTimelockID, op, signer.PublicKey(), ac.PublicKey())
			require.NoError(t, err)
			for _, ix := range ixs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment)
			}

			var opAccount timelock.Operation
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, operationPDA, config.DefaultCommitment, &opAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}

			require.Equal(t,
				timelock.Finalized_OperationState,
				opAccount.State,
				"operation is not finalized",
			)

			t.Run("success: bypass_execute_batch", func(t *testing.T) {
				ix := timelock.NewBypasserExecuteBatchInstruction(
					config.TestTimelockID,
					id,
					operationPDA,
					timelockutil.GetConfigPDA(config.TestTimelockID),
					timelockutil.GetSignerPDA(config.TestTimelockID),
					ac.PublicKey(),
					signer.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op.RemainingAccounts()...)

				vIx, err := ix.ValidateAndBuild()
				require.NoError(t, err)

				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, signer, config.DefaultCommitment)
				require.NotNil(t, tx)

				parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
					[]common.EventMapping{
						common.EventMappingFor[timelockutil.BypasserCallExecuted]("BypasserCallExecuted"),
					},
				)

				for i, ixx := range op.ToInstructionData() {
					event := parsedLogs[0].EventData[i].Data.(*timelockutil.BypasserCallExecuted)
					require.Equal(t, uint64(i), event.Index)
					require.Equal(t, ixx.ProgramId, event.Target)
					require.Equal(t, ixx.Data, common.NormalizeData(event.Data))
				}

				// check if operation pda is closed
				testutils.AssertClosedAccount(ctx, t, solanaGoClient, operationPDA, config.DefaultCommitment)

				recipientWsolBalance, err := solanaGoClient.GetTokenAccountBalance(
					ctx,
					recipientATA,
					config.DefaultCommitment,
				)
				require.NoError(t, err)
				require.Equal(t,
					strconv.Itoa(int(allowance.recipient)),
					recipientWsolBalance.Value.Amount,
					"Recipient balance mismatch",
				)
			})
		})
	})
}
