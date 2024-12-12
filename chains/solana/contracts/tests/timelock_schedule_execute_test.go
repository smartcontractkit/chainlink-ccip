package contracts

import (
	"strconv"
	"testing"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/generated/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/generated/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils"
	mcmsUtils "github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/mcms"
)

func TestTimelockScheduleAndExecute(t *testing.T) {
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

			var ac access_controller.AccessController
			acAccErr := utils.GetAccountDataBorshInto(
				ctx,
				solanaGoClient,
				data.AccessController.PublicKey(),
				config.DefaultCommitment,
				&ac,
			)
			require.NoError(t, acAccErr)

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

	t.Run("schedule_execute: multiple wsol transfer operation", func(t *testing.T) {
		t.Run("setup: update delay for testing", func(t *testing.T) {
			newMinDelay := uint64(1)

			ix, updateDelayIxErr := timelock.NewUpdateDelayInstruction(
				newMinDelay,
				config.TimelockConfigPDA,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, updateDelayIxErr)

			result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			require.NotNil(t, result)

			var configAccount timelock.Config
			getConfigAccountErr := utils.GetAccountDataBorshInto(ctx, solanaGoClient, config.TimelockConfigPDA, config.DefaultCommitment, &configAccount)
			require.NoError(t, getConfigAccountErr, "failed to get account info")

			require.Equal(t, newMinDelay, configAccount.MinDelay, "MinDelay is not updated")
		})

		t.Run("setup: wsol transfer operation", func(t *testing.T) {
			requiredAmount := allowance.recipient

			fundPDAIx := system.NewTransferInstruction(allowance.timelockAuthority, admin.PublicKey(), config.TimelockSignerPDA).Build()

			createAdminATAIx, _, caErr := utils.CreateAssociatedTokenAccount(tokenProgram, wsol, admin.PublicKey(), admin.PublicKey())
			require.NoError(t, caErr)

			wrapSolIx := system.NewTransferInstruction(
				requiredAmount,
				admin.PublicKey(),
				adminATA,
			).Build()

			syncNativeIx, snErr := utils.SyncNative(
				tokenProgram,
				adminATA, // token account
			)
			require.NoError(t, snErr)

			// approve can't be deligated to timelock authority(CPI Guard)
			approveIx, aiErr := utils.TokenApproveChecked(
				requiredAmount*2, // double the requiredAmount for op2 + op3(op2 will be executed only)
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

		t.Run("success: schedule and execute operations", func(t *testing.T) {
			salt1, err := mcmsUtils.SimpleSalt()
			require.NoError(t, err)
			op1 := TimelockOperation{
				Predecessor: config.TimelockEmptyOpID,
				Salt:        salt1,
				Delay:       2,
			}
			cIx, _, ciErr := utils.CreateAssociatedTokenAccount(
				tokenProgram,
				wsol,
				recipient.PublicKey(),
				config.TimelockSignerPDA,
			)
			require.NoError(t, ciErr)
			op1.AddInstruction(cIx, []solana.PublicKey{solana.TokenProgramID, solana.SPLAssociatedTokenAccountProgramID})

			salt2, err := mcmsUtils.SimpleSalt()
			require.NoError(t, err)
			op2 := TimelockOperation{
				Predecessor: op1.OperationID(),
				Salt:        salt2,
				Delay:       2,
			}

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
			op2.AddInstruction(tIx, []solana.PublicKey{tokenProgram})

			salt3, err := mcmsUtils.SimpleSalt()
			require.NoError(t, err)
			op3 := TimelockOperation{
				Predecessor: op1.OperationID(),
				Salt:        salt3,
				Delay:       60,
			}

			anotherTransferIx, atErr := utils.TokenTransferChecked(
				allowance.recipient,
				wsolDecimal,
				tokenProgram,
				adminATA,
				wsol,
				recipientATA,
				config.TimelockSignerPDA,
				nil,
			)
			require.NoError(t, atErr)
			op3.AddInstruction(anotherTransferIx, []solana.PublicKey{tokenProgram})

			t.Run("schedule", func(t *testing.T) {
				proposer := roleMap[timelock.Proposer_Role].RandomPick()
				proposerAccessController := roleMap[timelock.Proposer_Role].AccessController.PublicKey()

				executor := roleMap[timelock.Executor_Role].RandomPick()
				executorAccessController := roleMap[timelock.Executor_Role].AccessController.PublicKey()

				t.Run("success: schedule all operations", func(t *testing.T) {
					for _, op := range []TimelockOperation{op1, op2, op3} {
						id := op.OperationID()
						operationPDA1 := op.OperationPDA()

						initOpIx, iErr := timelock.NewInitializeOperationInstruction(
							op.OperationID(),
							op.Predecessor,
							op.Salt,
							op.IxsCountU32(),
							config.TimelockConfigPDA,
							op.OperationPDA(),
							proposer.PublicKey(),
							proposer.PublicKey(),
							solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, iErr)

						utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initOpIx}, proposer, config.DefaultCommitment)

						for _, ixData := range op.ToInstructionData() {
							appendIxsIx, aErr := timelock.NewAppendInstructionsInstruction(
								op.OperationID(),
								[]timelock.InstructionData{ixData},
								op.OperationPDA(),
								proposer.PublicKey(),
								solana.SystemProgramID, // for reallocation
							).ValidateAndBuild()
							require.NoError(t, aErr)
							utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{appendIxsIx}, proposer, config.DefaultCommitment)
						}

						t.Run("clear & reappend op instructions", func(t *testing.T) {
							// clear instructions so that we can reinitialize the operation
							clearIx, ciErr := timelock.NewClearOperationInstruction(op.OperationID(), op.OperationPDA(), proposer.PublicKey()).ValidateAndBuild()
							require.NoError(t, ciErr)

							// send clear and check if it's closed
							utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{clearIx}, proposer, config.DefaultCommitment)
							utils.AssertClosedAccount(ctx, t, solanaGoClient, op.OperationPDA(), config.DefaultCommitment)

							// reinitialize operation
							utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initOpIx}, proposer, config.DefaultCommitment)

							// reappend instructions
							for _, ixData := range op.ToInstructionData() {
								appendIxsIx, aErr := timelock.NewAppendInstructionsInstruction(
									op.OperationID(),
									[]timelock.InstructionData{ixData},
									op.OperationPDA(),
									proposer.PublicKey(),
									solana.SystemProgramID, // for reallocation
								).ValidateAndBuild()
								require.NoError(t, aErr)
								utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{appendIxsIx}, proposer, config.DefaultCommitment)
							}
						})

						finIxIx, fErr := timelock.NewFinalizeOperationInstruction(
							op.OperationID(),
							op.OperationPDA(),
							proposer.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, fErr)

						utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{finIxIx}, proposer, config.DefaultCommitment)

						ix, ixVErr := timelock.NewScheduleBatchInstruction(
							id,
							op.Delay,
							config.TimelockConfigPDA,
							operationPDA1,
							proposerAccessController,
							proposer.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, ixVErr)

						result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment)
						require.NotNil(t, result)

						var opAccount timelock.Operation
						err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, operationPDA1, config.DefaultCommitment, &opAccount)
						if err != nil {
							require.NoError(t, err, "failed to get account info")
						}

						require.Equal(t,
							result.BlockTime.Time().Add(time.Duration(op.Delay)*time.Second).Unix(),
							int64(opAccount.Timestamp),
							"Scheduled Times don't match",
						)

						require.Equal(t,
							id,
							opAccount.Id,
							"Ids don't match",
						)
					}
				})

				t.Run("wait for operation 1 to be ready", func(t *testing.T) {
					// Wait for operations to be ready
					err := WaitForOperationToBeReady(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment)
					require.NoError(t, err)
				})

				t.Run("fail: OperationAlreadyExists", func(t *testing.T) {
					ix := timelock.NewScheduleBatchInstruction(
						op1.OperationID(),
						op1.Delay,
						config.TimelockConfigPDA,
						op1.OperationPDA(),
						proposerAccessController,
						proposer.PublicKey(),
					).Build()

					utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment, []string{"Error Code: " + timelock.OperationAlreadyScheduled_TimelockError.String()})
				})

				t.Run("wait for operation 2 to be ready", func(t *testing.T) {
					// Wait for operations to be ready
					err := WaitForOperationToBeReady(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment)
					require.NoError(t, err)
				})

				t.Run("fail: should provide the right dependency pda", func(t *testing.T) {
					ix := timelock.NewExecuteBatchInstruction(
						op2.OperationID(),
						config.TimelockConfigPDA,
						config.TimelockSignerPDA,
						op2.OperationPDA(),
						op2.OperationPDA(), // wrong dependency
						executorAccessController,
						executor.PublicKey(),
					)
					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

					vIx, err := ix.ValidateAndBuild()
					require.NoError(t, err)

					utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment, []string{"Error Code: " + timelock.InvalidInput_TimelockError.String()})
				})

				t.Run("fail: not able to execute op2 before dependency(op1) execution", func(t *testing.T) {
					ix := timelock.NewExecuteBatchInstruction(
						op2.OperationID(),
						config.TimelockConfigPDA,
						config.TimelockSignerPDA,
						op2.OperationPDA(),
						op1.OperationPDA(), // not executed yet
						executorAccessController,
						executor.PublicKey(),
					)

					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

					vIx, err := ix.ValidateAndBuild()
					require.NoError(t, err)

					utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment, []string{"Error Code: " + timelock.MissingDependency_TimelockError.String()})
				})

				t.Run("success: op1 executed", func(t *testing.T) {
					ix := timelock.NewExecuteBatchInstruction(
						op1.OperationID(),
						config.TimelockConfigPDA,
						config.TimelockSignerPDA,
						op1.OperationPDA(),
						config.TimelockEmptyOpID,
						executorAccessController,
						executor.PublicKey(),
					)

					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op1.RemainingAccounts()...)

					vIx, err := ix.ValidateAndBuild()
					require.NoError(t, err)

					result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment)
					require.NotNil(t, result)

					var opAccount timelock.Operation
					err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment, &opAccount)
					if err != nil {
						require.NoError(t, err, "failed to get account info")
					}

					require.Equal(t,
						config.TimelockOpDoneTimestamp,
						opAccount.Timestamp,
						"Executed operation's time should be 1(DONE_TIMESTAMP)",
					)
				})

				t.Run("success: op2 executed", func(t *testing.T) {
					ix := timelock.NewExecuteBatchInstruction(
						op2.OperationID(),
						config.TimelockConfigPDA,
						config.TimelockSignerPDA,
						op2.OperationPDA(),
						op1.OperationPDA(),
						executorAccessController,
						executor.PublicKey(),
					)

					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

					vIx, err := ix.ValidateAndBuild()
					require.NoError(t, err)

					result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment)
					require.NotNil(t, result)

					var opAccount timelock.Operation
					err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment, &opAccount)
					if err != nil {
						require.NoError(t, err, "failed to get account info")
					}

					require.Equal(t,
						config.TimelockOpDoneTimestamp,
						opAccount.Timestamp,
						"Executed operation's time should be 1(DONE_TIMESTAMP)",
					)

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

				t.Run("failure on execution try: op3 is not ready", func(t *testing.T) {
					ix := timelock.NewExecuteBatchInstruction(
						op3.OperationID(),
						config.TimelockConfigPDA,
						config.TimelockSignerPDA,
						op3.OperationPDA(),
						config.TimelockEmptyOpID,
						executorAccessController,
						executor.PublicKey(),
					)
					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op3.RemainingAccounts()...)

					vIx, vIxErr := ix.ValidateAndBuild()
					require.NoError(t, vIxErr)

					utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment, []string{"Error Code: " + timelock.OperationNotReady_TimelockError.String()})
				})
			})
		})
	})
}
