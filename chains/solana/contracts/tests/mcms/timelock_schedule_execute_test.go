package contracts

import (
	"strconv"
	"sync"
	"testing"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/external_program_cpi_stub"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/accesscontroller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	timelockutil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestTimelockScheduleAndExecute(t *testing.T) {
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
		t.Run("setup: update delay for testing", func(t *testing.T) {
			newMinDelay := uint64(1)

			ix, updateDelayIxErr := timelock.NewUpdateDelayInstruction(
				config.TestTimelockID,
				newMinDelay,
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, updateDelayIxErr)

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			require.NotNil(t, result)

			var configAccount timelock.Config
			getConfigAccountErr := common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment, &configAccount)
			require.NoError(t, getConfigAccountErr, "failed to get account info")

			require.Equal(t, newMinDelay, configAccount.MinDelay, "MinDelay is not updated")
		})

		t.Run("setup: wsol transfer operation", func(t *testing.T) {
			requiredAmount := allowance.recipient

			fundPDAIx := system.NewTransferInstruction(allowance.timelockAuthority, admin.PublicKey(), timelockutil.GetSignerPDA(config.TestTimelockID)).Build()

			createAdminATAIx, _, caErr := tokens.CreateAssociatedTokenAccount(tokenProgram, wsol, admin.PublicKey(), admin.PublicKey())
			require.NoError(t, caErr)

			wrapSolIx := system.NewTransferInstruction(
				requiredAmount,
				admin.PublicKey(),
				adminATA,
			).Build()

			syncNativeIx, snErr := tokens.SyncNative(
				tokenProgram,
				adminATA, // token account
			)
			require.NoError(t, snErr)

			// approve can't be deligated to timelock authority(CPI Guard)
			approveIx, aiErr := tokens.TokenApproveChecked(
				requiredAmount*2, // double the requiredAmount for op2 + op3(op2 will be executed only)
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

		t.Run("success: schedule and execute operations", func(t *testing.T) {
			salt1, err := timelockutil.SimpleSalt()
			require.NoError(t, err)
			op1 := timelockutil.Operation{
				TimelockID:  config.TestTimelockID,
				Predecessor: config.TimelockEmptyOpID,
				Salt:        salt1,
				Delay:       2,
			}
			cIx, _, ciErr := tokens.CreateAssociatedTokenAccount(
				tokenProgram,
				wsol,
				recipient.PublicKey(),
				timelockutil.GetSignerPDA(config.TestTimelockID),
			)
			require.NoError(t, ciErr)
			op1.AddInstruction(cIx, []solana.PublicKey{solana.TokenProgramID, solana.SPLAssociatedTokenAccountProgramID})

			salt2, err := timelockutil.SimpleSalt()
			require.NoError(t, err)
			op2 := timelockutil.Operation{
				TimelockID:  config.TestTimelockID,
				Predecessor: op1.OperationID(),
				Salt:        salt2,
				Delay:       2,
			}

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
			op2.AddInstruction(tIx, []solana.PublicKey{tokenProgram})

			salt3, err := timelockutil.SimpleSalt()
			require.NoError(t, err)
			op3 := timelockutil.Operation{
				TimelockID:  config.TestTimelockID,
				Predecessor: op1.OperationID(),
				Salt:        salt3,
				Delay:       300, // enough delay to assert OperationNotReady error
			}

			anotherTransferIx, atErr := tokens.TokenTransferChecked(
				allowance.recipient,
				wsolDecimal,
				tokenProgram,
				adminATA,
				wsol,
				recipientATA,
				timelockutil.GetSignerPDA(config.TestTimelockID),
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
					for _, op := range []timelockutil.Operation{op1, op2, op3} {
						invalidIxs, ierr := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, op, proposer.PublicKey(), proposerAccessController)
						require.NoError(t, ierr)
						for _, ix := range invalidIxs {
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment)
						}

						// clear instructions so that we can reinitialize the operation
						clearIx, ciErr := timelock.NewClearOperationInstruction(
							config.TestTimelockID,
							op.OperationID(),
							op.OperationPDA(),
							timelockutil.GetConfigPDA(config.TestTimelockID),
							proposerAccessController,
							proposer.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, ciErr)

						// send clear and check if it's closed
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{clearIx}, proposer, config.DefaultCommitment)
						testutils.AssertClosedAccount(ctx, t, solanaGoClient, op.OperationPDA(), config.DefaultCommitment)

						// re-preload instructions
						ixs, err := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, op, proposer.PublicKey(), proposerAccessController)
						require.NoError(t, err)
						for _, ix := range ixs {
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment)
						}

						ix, ixVErr := timelock.NewScheduleBatchInstruction(
							config.TestTimelockID,
							op.OperationID(),
							op.Delay,
							op.OperationPDA(),
							timelockutil.GetConfigPDA(config.TestTimelockID),
							proposerAccessController,
							proposer.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, ixVErr)

						result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment)
						require.NotNil(t, result)

						var opAccount timelock.Operation
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, op.OperationPDA(), config.DefaultCommitment, &opAccount)
						require.NoError(t, err, "failed to get account info")

						require.Equal(t,
							result.BlockTime.Time().Add(time.Duration(op.Delay)*time.Second).Unix(),
							int64(opAccount.Timestamp),
							"Scheduled Times don't match",
						)

						require.Equal(t,
							op.OperationID(),
							opAccount.Id,
							"Ids don't match",
						)
					}
				})

				t.Run("wait for operation 1 to be ready", func(t *testing.T) {
					// Wait for operations to be ready
					err := timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment)
					require.NoError(t, err)
				})

				t.Run("fail: OperationAlreadyExists", func(t *testing.T) {
					ix := timelock.NewScheduleBatchInstruction(
						config.TestTimelockID,
						op1.OperationID(),
						op1.Delay,
						op1.OperationPDA(),
						timelockutil.GetConfigPDA(config.TestTimelockID),
						proposerAccessController,
						proposer.PublicKey(),
					).Build()

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment, []string{"Error Code: " + timelock.OperationAlreadyScheduled_TimelockError.String()})
				})

				t.Run("wait for operation 2 to be ready", func(t *testing.T) {
					// Wait for operations to be ready
					err := timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment)
					require.NoError(t, err)
				})

				t.Run("fail: should provide the right dependency pda", func(t *testing.T) {
					ix := timelock.NewExecuteBatchInstruction(
						config.TestTimelockID,
						op2.OperationID(),
						op2.OperationPDA(),
						op2.OperationPDA(), // wrong dependency
						timelockutil.GetConfigPDA(config.TestTimelockID),
						timelockutil.GetSignerPDA(config.TestTimelockID),
						executorAccessController,
						executor.PublicKey(),
					)
					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

					vIx, err := ix.ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment, []string{"Error Code: " + timelock.InvalidInput_TimelockError.String()})
				})

				t.Run("fail: not able to execute op2 before dependency(op1) execution", func(t *testing.T) {
					ix := timelock.NewExecuteBatchInstruction(
						config.TestTimelockID,
						op2.OperationID(),
						op2.OperationPDA(),
						op1.OperationPDA(), // not executed yet
						timelockutil.GetConfigPDA(config.TestTimelockID),
						timelockutil.GetSignerPDA(config.TestTimelockID),
						executorAccessController,
						executor.PublicKey(),
					)

					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

					vIx, err := ix.ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment, []string{"Error Code: " + timelock.MissingDependency_TimelockError.String()})
				})

				t.Run("success: op1 executed", func(t *testing.T) {
					ix := timelock.NewExecuteBatchInstruction(
						config.TestTimelockID,
						op1.OperationID(),
						op1.OperationPDA(),
						config.TimelockEmptyOpID,
						timelockutil.GetConfigPDA(config.TestTimelockID),
						timelockutil.GetSignerPDA(config.TestTimelockID),
						executorAccessController,
						executor.PublicKey(),
					)

					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op1.RemainingAccounts()...)

					vIx, err := ix.ValidateAndBuild()
					require.NoError(t, err)

					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment)
					require.NotNil(t, tx)

					parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
						[]common.EventMapping{
							common.EventMappingFor[timelockutil.CallExecuted]("CallExecuted"),
						},
					)

					for i, ixx := range op1.ToInstructionData() {
						event := parsedLogs[0].EventData[i].Data.(*timelockutil.CallExecuted)
						require.Equal(t, op1.OperationID(), event.ID)
						require.Equal(t, uint64(i), event.Index)
						require.Equal(t, ixx.ProgramId, event.Target)
						require.Equal(t, ixx.Data, common.NormalizeData(event.Data))
					}

					var opAccount timelock.Operation
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment, &opAccount)
					if err != nil {
						require.NoError(t, err, "failed to get account info")
					}

					require.Equal(t,
						timelock.Done_OperationState,
						opAccount.State,
						"Executed operation should be marked as done",
					)
				})

				t.Run("success: op2 executed", func(t *testing.T) {
					ix := timelock.NewExecuteBatchInstruction(
						config.TestTimelockID,
						op2.OperationID(),
						op2.OperationPDA(),
						op1.OperationPDA(),
						timelockutil.GetConfigPDA(config.TestTimelockID),
						timelockutil.GetSignerPDA(config.TestTimelockID),
						executorAccessController,
						executor.PublicKey(),
					)

					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

					vIx, err := ix.ValidateAndBuild()
					require.NoError(t, err)

					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment)
					require.NotNil(t, tx)

					parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
						[]common.EventMapping{
							common.EventMappingFor[timelockutil.CallExecuted]("CallExecuted"),
						},
					)

					for i, ixx := range op2.ToInstructionData() {
						event := parsedLogs[0].EventData[i].Data.(*timelockutil.CallExecuted)
						require.Equal(t, op2.OperationID(), event.ID)
						require.Equal(t, uint64(i), event.Index)
						require.Equal(t, ixx.ProgramId, event.Target)
						require.Equal(t, ixx.Data, common.NormalizeData(event.Data))
					}

					var opAccount timelock.Operation
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment, &opAccount)
					if err != nil {
						require.NoError(t, err, "failed to get account info")
					}

					require.Equal(t,
						timelock.Done_OperationState,
						opAccount.State,
						"Executed operation should be marked as done",
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
						config.TestTimelockID,
						op3.OperationID(),
						op3.OperationPDA(),
						config.TimelockEmptyOpID,
						timelockutil.GetConfigPDA(config.TestTimelockID),
						timelockutil.GetSignerPDA(config.TestTimelockID),
						executorAccessController,
						executor.PublicKey(),
					)
					ix.AccountMetaSlice = append(ix.AccountMetaSlice, op3.RemainingAccounts()...)

					vIx, vIxErr := ix.ValidateAndBuild()
					require.NoError(t, vIxErr)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{vIx}, executor, config.DefaultCommitment, []string{"Error Code: " + timelock.OperationNotReady_TimelockError.String()})
				})
			})
		})
	})

	t.Run("function blockers", func(t *testing.T) {
		t.Parallel()
		proposer := roleMap[timelock.Proposer_Role].RandomPick()
		proposerAccessController := roleMap[timelock.Proposer_Role].AccessController.PublicKey()

		salt, err := timelockutil.SimpleSalt()
		require.NoError(t, err)

		op := timelockutil.Operation{
			TimelockID:  config.TestTimelockID,
			Predecessor: config.TimelockEmptyOpID,
			Salt:        salt,
			Delay:       1,
		}

		ix, err := external_program_cpi_stub.NewInitializeInstruction(
			config.StubAccountPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		op.AddInstruction(ix, []solana.PublicKey{solana.TokenProgramID, solana.SPLAssociatedTokenAccountProgramID})

		t.Run("blocks initialize function", func(t *testing.T) {
			bIx, bIxErr := timelock.NewBlockFunctionSelectorInstruction(
				config.TestTimelockID,
				[8]uint8(external_program_cpi_stub.Instruction_Initialize.Bytes()),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, bIxErr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{bIx}, admin, config.DefaultCommitment)

			blockedSelectors, bserr := timelockutil.GetBlockedFunctionSelectors(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment)
			require.NoError(t, bserr)
			require.Contains(t, blockedSelectors, external_program_cpi_stub.Instruction_Initialize.Bytes())
		})

		t.Run("not able to block function that is already blocked", func(t *testing.T) {
			bbIx, bbIxErr := timelock.NewBlockFunctionSelectorInstruction(
				config.TestTimelockID,
				[8]uint8(external_program_cpi_stub.Instruction_Initialize.Bytes()),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, bbIxErr)
			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{bbIx}, admin, config.DefaultCommitment, []string{"Error Code: " + timelock.AlreadyBlocked_TimelockError.String()})
		})

		id := op.OperationID()
		operationPDA := op.OperationPDA()

		ixs, err := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, op, proposer.PublicKey(), proposerAccessController)
		require.NoError(t, err)
		for _, ix := range ixs {
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment)
		}

		scIx, scIxVErr := timelock.NewScheduleBatchInstruction(
			config.TestTimelockID,
			id,
			op.Delay,
			operationPDA,
			timelockutil.GetConfigPDA(config.TestTimelockID),
			proposerAccessController,
			proposer.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, scIxVErr)

		testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{scIx}, proposer, config.DefaultCommitment, []string{"Error Code: " + timelock.BlockedSelector_TimelockError.String()})

		t.Run("unblocks initialize function", func(t *testing.T) {
			bIx, bIxErr := timelock.NewUnblockFunctionSelectorInstruction(
				config.TestTimelockID,
				[8]uint8(external_program_cpi_stub.Instruction_Initialize.Bytes()),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, bIxErr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{bIx}, admin, config.DefaultCommitment)

			blockedSelectors, bserr := timelockutil.GetBlockedFunctionSelectors(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment)
			require.NoError(t, bserr)
			require.NotContains(t, blockedSelectors, external_program_cpi_stub.Instruction_Initialize.Bytes())
		})

		t.Run("not able to unblock function that is not blocked", func(t *testing.T) {
			bbIx, bbIxErr := timelock.NewUnblockFunctionSelectorInstruction(
				config.TestTimelockID,
				[8]uint8(external_program_cpi_stub.Instruction_Initialize.Bytes()),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, bbIxErr)
			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{bbIx}, admin, config.DefaultCommitment, []string{"Error Code: " + timelock.SelectorNotFound_TimelockError.String()})
		})

		t.Run("when unblocked, able to schedule operation", func(t *testing.T) {
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{scIx}, proposer, config.DefaultCommitment)
			require.NotNil(t, result)

			var opAccount timelock.Operation
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, operationPDA, config.DefaultCommitment, &opAccount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, op.OperationID(), opAccount.Id, "Ids don't match")
			require.Equal(t,
				result.BlockTime.Time().Add(time.Duration(op.Delay)*time.Second).Unix(),
				int64(opAccount.Timestamp),
				"Scheduled Times don't match",
			)
		})

		t.Run("can't register more than MAX_SELECTOR", func(t *testing.T) {
			// check if it's empty
			oldBlockedSelectors, gberr := timelockutil.GetBlockedFunctionSelectors(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment)
			require.NoError(t, gberr)
			require.Empty(t, oldBlockedSelectors)

			ixs := []solana.Instruction{}
			for i := 0; i < config.MaxFunctionSelectorLen; i++ {
				ix, nberr := timelock.NewBlockFunctionSelectorInstruction(
					config.TestTimelockID,
					[8]uint8{byte(i)},
					timelockutil.GetConfigPDA(config.TestTimelockID),
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, nberr)

				ixs = append(ixs, ix)
			}

			// max selectors at 128
			chunkSize := 18
			var wg sync.WaitGroup

			for i := 0; i < len(ixs); i += chunkSize {
				wg.Add(1)
				end := i + chunkSize
				if end > len(ixs) {
					end = len(ixs)
				}
				chunk := ixs[i:end]

				go func(chunk []solana.Instruction) {
					defer wg.Done()
					testutils.SendAndConfirm(ctx, t, solanaGoClient, chunk, admin, config.DefaultCommitment)
				}(chunk)
			}

			wg.Wait()

			// check if it's full
			blockedSelectors, bserr := timelockutil.GetBlockedFunctionSelectors(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment)
			require.NoError(t, bserr)
			require.Equal(t, config.MaxFunctionSelectorLen, len(blockedSelectors))

			// try one more
			ix, nberr := timelock.NewBlockFunctionSelectorInstruction(
				config.TestTimelockID,
				[8]uint8{255},
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, nberr)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, []string{"Error Code: " + timelock.MaxCapacityReached_TimelockError.String()})
		})
	})
}
