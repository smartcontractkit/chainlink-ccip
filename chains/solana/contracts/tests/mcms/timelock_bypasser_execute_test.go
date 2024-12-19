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

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/accesscontroller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils"
	mcmsUtils "github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"
)

func TestTimelockBypasserExecute(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	timelock.SetProgramID(config.TimelockProgram)
	access_controller.SetProgramID(config.AccessControllerProgram)

	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	roles, roleMap := mcmsUtils.TestRoleAccounts(t, config.NumAccountsPerRole)
	solanaGoClient := utils.DeployAllPrograms(t, utils.PathToAnchorConfig, admin)

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
		utils.FundAccounts(ctx, all, solanaGoClient, t)
	})

	t.Run("setup:init access controllers", func(t *testing.T) {
		for _, data := range roleMap {
			initAccIxs, initAccIxsErr := InitAccessControllersIxs(ctx, data.AccessController.PublicKey(), admin, solanaGoClient)
			require.NoError(t, initAccIxsErr)

			utils.SendAndConfirm(ctx, t, solanaGoClient, initAccIxs, admin, config.DefaultCommitment, utils.AddSigners(data.AccessController))

			var ac access_controller.AccessController
			acAccErr := utils.GetAccountDataBorshInto(ctx, solanaGoClient, data.AccessController.PublicKey(), config.DefaultCommitment, &ac)
			if acAccErr != nil {
				require.NoError(t, acAccErr, "failed to get account info")
			}
		}
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

		initTimelockIx, initErr := timelock.NewInitializeInstruction(
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
		require.NoError(t, initErr)

		utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initTimelockIx}, admin, config.DefaultCommitment)

		var configAccount timelock.Config
		cfgErr := utils.GetAccountDataBorshInto(ctx, solanaGoClient, config.TimelockConfigPDA, config.DefaultCommitment, &configAccount)
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
			batchAddAccessIxs, batchAddAccessIxsErr := TimelockBatchAddAccessIxs(ctx, data.AccessController.PublicKey(), role, addresses, admin, config.BatchAddAccessChunkSize, solanaGoClient)
			require.NoError(t, batchAddAccessIxsErr)

			for _, ix := range batchAddAccessIxs {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
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

			fundPDAIx := system.NewTransferInstruction(allowance.timelockAuthority, admin.PublicKey(), config.TimelockSignerPDA).Build()

			createAdminATAIx, _, caErr := utils.CreateAssociatedTokenAccount(tokenProgram, wsol, admin.PublicKey(), admin.PublicKey())
			require.NoError(t, caErr)

			wrapSolIx, wsErr := system.NewTransferInstruction(
				requiredAmount,
				admin.PublicKey(),
				adminATA,
			).ValidateAndBuild()
			require.NoError(t, wsErr)

			syncNativeIx, snErr := utils.SyncNative(
				tokenProgram,
				adminATA, // token account
			)
			require.NoError(t, snErr)

			// approve can't be deligated to timelock authority(CPI Guard)
			approveIx, aiErr := utils.TokenApproveChecked(
				requiredAmount,
				wsolDecimal,
				tokenProgram,
				adminATA,
				wsol,
				config.TimelockSignerPDA,
				admin.PublicKey(),
				nil,
			)
			require.NoError(t, aiErr)

			result := utils.SendAndConfirm(ctx, t, solanaGoClient,
				[]solana.Instruction{createAdminATAIx, wrapSolIx, syncNativeIx, fundPDAIx, approveIx},
				admin, config.DefaultCommitment)
			require.NotNil(t, result)

			// check results
			timelockAuthorityBalance, tlBalanceErr := solanaGoClient.GetBalance(
				ctx,
				config.TimelockSignerPDA,
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

		t.Run("success: schedule and execute batch instructions", func(t *testing.T) {
			salt, err := mcmsUtils.SimpleSalt()
			require.NoError(t, err)
			op := TimelockOperation{
				Predecessor: config.TimelockEmptyOpID,
				Salt:        salt,
				Delay:       uint64(1000),
			}

			cIx, _, ciErr := utils.CreateAssociatedTokenAccount(
				tokenProgram,
				wsol,
				recipient.PublicKey(),
				config.TimelockSignerPDA,
			)
			require.NoError(t, ciErr)
			op.AddInstruction(
				cIx,
				[]solana.PublicKey{tokenProgram, solana.SPLAssociatedTokenAccountProgramID},
			)

			tIx, tiErr := utils.TokenTransferChecked(
				allowance.recipient,
				wsolDecimal,
				tokenProgram,
				adminATA,
				wsol,
				recipientATA,
				config.TimelockSignerPDA,
				nil,
			)
			require.NoError(t, tiErr)
			op.AddInstruction(tIx, []solana.PublicKey{tokenProgram})

			id := op.OperationID()
			operationPDA := op.OperationPDA()
			signer := roleMap[timelock.Proposer_Role].RandomPick()

			ixs, err := TimelockPreloadOperationIxs(ctx, op, signer.PublicKey(), solanaGoClient)
			require.NoError(t, err)
			for _, ix := range ixs {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment)
			}

			var opAccount timelock.Operation
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, operationPDA, config.DefaultCommitment, &opAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}

			require.Equal(t,
				true,
				opAccount.IsFinalized,
				"operation is not finalized",
			)

			t.Run("success: bypass_execute_batch", func(t *testing.T) {
				signer := roleMap[timelock.Bypasser_Role].RandomPick()
				ac := roleMap[timelock.Bypasser_Role].AccessController

				ix := timelock.NewBypasserExecuteBatchInstruction(
					id,
					operationPDA,
					config.TimelockConfigPDA,
					config.TimelockSignerPDA,
					ac.PublicKey(),
					signer.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op.RemainingAccounts()...)

				vIx, err := ix.ValidateAndBuild()
				require.NoError(t, err)

				tx := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, signer, config.DefaultCommitment)
				require.NotNil(t, tx)

				parsedLogs := utils.ParseLogMessages(tx.Meta.LogMessages,
					[]utils.EventMapping{
						utils.EventMappingFor[BypasserCallExecuted]("BypasserCallExecuted"),
					},
				)

				for i, ixx := range op.ToInstructionData() {
					event := parsedLogs[0].EventData[i].Data.(*BypasserCallExecuted)
					require.Equal(t, uint64(i), event.Index)
					require.Equal(t, ixx.ProgramId, event.Target)
					require.Equal(t, ixx.Data, utils.NormalizeData(event.Data))
				}

				var opAccount timelock.Operation
				err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, operationPDA, config.DefaultCommitment, &opAccount)
				if err != nil {
					require.NoError(t, err, "failed to get account info")
				}

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
