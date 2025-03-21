package contracts

import (
	"fmt"
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
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/accesscontroller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	timelockutil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/timelock"
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

	roles, roleMap := timelockutil.TestRoleAccounts(config.NumAccountsPerRole)
	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)

	t.Run("setup:funding", func(t *testing.T) {
		all := []solana.PrivateKey{}
		all = append(all, admin, anotherAdmin, user)
		// utils.FundAccounts(ctx, []solana.PrivateKey{admin}, client, t)
		for _, role := range roles {
			all = append(all, role.Accounts...)
		}
		testutils.FundAccounts(ctx, all, solanaGoClient, t)
	})

	t.Run("setup:init access controllers", func(t *testing.T) {
		for _, data := range roleMap {
			initAccIxs, ierr := timelockutil.GetInitAccessControllersIxs(ctx, data.AccessController.PublicKey(), admin, solanaGoClient)
			require.NoError(t, ierr)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, initAccIxs, admin, config.DefaultCommitment, common.AddSigners(data.AccessController))

			var ac access_controller.AccessController
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, data.AccessController.PublicKey(), config.DefaultCommitment, &ac)
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

		initTimelockIx, ierr := timelock.NewInitializeInstruction(
			config.TestTimelockID,
			config.MinDelay,
			timelockutil.GetConfigPDA(config.TestTimelockID),
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
		require.NoError(t, ierr)

		result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{initTimelockIx}, anotherAdmin, config.DefaultCommitment, []string{"Error Code: " + timelockutil.UnauthorizedError.String()})
		require.NotNil(t, result)
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

		initTimelockIx, ierr := timelock.NewInitializeInstruction(
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
		require.NoError(t, ierr)

		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initTimelockIx}, admin, config.DefaultCommitment)

		var configAccount timelock.Config
		err = common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment, &configAccount)
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
		t.Run("fail to transfer ownership when not owner", func(t *testing.T) {
			instruction, ierr := timelock.NewTransferOwnershipInstruction(
				config.TestTimelockID,
				anotherAdmin.PublicKey(),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, ierr)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + timelockutil.UnauthorizedError.String()})
			require.NotNil(t, result)
		})

		t.Run("Current owner cannot propose self", func(t *testing.T) {
			instruction, ierr := timelock.NewTransferOwnershipInstruction(
				config.TestTimelockID,
				admin.PublicKey(),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, ierr)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment, []string{"Error Code: " + timelock.InvalidInput_TimelockError.String()})
			require.NotNil(t, result)
		})

		t.Run("successfully transfer ownership", func(t *testing.T) {
			instruction, ierr := timelock.NewTransferOwnershipInstruction(
				config.TestTimelockID,
				anotherAdmin.PublicKey(),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, ierr)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment)
			require.NotNil(t, result)
		})

		t.Run("Fail to accept ownership when not proposed_owner", func(t *testing.T) {
			instruction, ierr := timelock.NewAcceptOwnershipInstruction(
				config.TestTimelockID,
				timelockutil.GetConfigPDA(config.TestTimelockID),
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, ierr)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + timelockutil.UnauthorizedError.String()})
			require.NotNil(t, result)
		})

		t.Run("anotherAdmin becomes owner", func(t *testing.T) {
			instruction, ierr := timelock.NewAcceptOwnershipInstruction(
				config.TestTimelockID,
				timelockutil.GetConfigPDA(config.TestTimelockID),
				anotherAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, ierr)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			// Validate proposed set to 0-address after accepting ownership
			var configAccount timelock.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
			require.Equal(t, anotherAdmin.PublicKey(), configAccount.Owner)
			require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
		})

		// get it back
		t.Run("retrieve back ownership to admin", func(t *testing.T) {
			tix, ierr := timelock.NewTransferOwnershipInstruction(
				config.TestTimelockID,
				admin.PublicKey(),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				anotherAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, ierr)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{tix}, anotherAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			aix, aerr := timelock.NewAcceptOwnershipInstruction(
				config.TestTimelockID,
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, aerr)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{aix}, admin, config.DefaultCommitment)
			require.NotNil(t, result)

			var configAccount timelock.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}

			require.Equal(t, admin.PublicKey(), configAccount.Owner)
			require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
		})
	})

	t.Run("setup:register access list & verify", func(t *testing.T) {
		for role, data := range roleMap {
			addresses := []solana.PublicKey{}
			for _, account := range data.Accounts {
				addresses = append(addresses, account.PublicKey())
			}
			batchAddAccessIxs, baerr := timelockutil.GetBatchAddAccessIxs(ctx, config.TestTimelockID, data.AccessController.PublicKey(), role, addresses, admin, config.BatchAddAccessChunkSize, solanaGoClient)
			require.NoError(t, baerr)

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

	t.Run("rbac: schedule and cancel a timelock operation", func(t *testing.T) {
		salt, serr := timelockutil.SimpleSalt()
		require.NoError(t, serr)
		nonExecutableOp := timelockutil.Operation{
			TimelockID:  config.TestTimelockID,
			Predecessor: config.TimelockEmptyOpID,
			Salt:        salt,
			Delay:       uint64(1),
		}

		ix := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), timelockutil.GetSignerPDA(config.TestTimelockID)).Build()
		nonExecutableOp.AddInstruction(ix, []solana.PublicKey{})

		t.Run("rbac: when try to schedule from non proposer role, it fails", func(t *testing.T) {
			proposer := roleMap[timelock.Proposer_Role].RandomPick()
			nonProposer := roleMap[timelock.Executor_Role].RandomPick()
			ac := roleMap[timelock.Proposer_Role].AccessController

			// preload the operation with proposer for access testing on schedule_batch
			prixs, prierr := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, nonExecutableOp, proposer.PublicKey(), ac.PublicKey())
			require.NoError(t, prierr)
			for _, ix := range prixs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment)
			}

			ix, scerr := timelock.NewScheduleBatchInstruction(
				config.TestTimelockID,
				nonExecutableOp.OperationID(),
				nonExecutableOp.Delay,
				nonExecutableOp.OperationPDA(),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				ac.PublicKey(),
				nonProposer.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, scerr)
			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, nonProposer, config.DefaultCommitment, []string{"Error Code: " + timelockutil.UnauthorizedError.String()})
		})

		t.Run("rbac: Should able to schedule tx with proposer role", func(t *testing.T) {
			proposer := roleMap[timelock.Proposer_Role].RandomPick()
			ac := roleMap[timelock.Proposer_Role].AccessController

			t.Run("rbac: when proposer's access is removed, it should not be able to schedule", func(t *testing.T) {
				raIx, raerr := access_controller.NewRemoveAccessInstruction(
					ac.PublicKey(),
					admin.PublicKey(),
					proposer.PublicKey(), // remove access of proposer
				).ValidateAndBuild()
				require.NoError(t, raerr)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{raIx}, admin, config.DefaultCommitment)

				found, ferr := accesscontroller.HasAccess(ctx, solanaGoClient, ac.PublicKey(), proposer.PublicKey(), config.DefaultCommitment)
				require.NoError(t, ferr)
				require.False(t, found, "Account %s should not be in the AccessList", proposer.PublicKey())

				ix, scerr := timelock.NewScheduleBatchInstruction(
					config.TestTimelockID,
					nonExecutableOp.OperationID(),
					nonExecutableOp.Delay,
					nonExecutableOp.OperationPDA(),
					timelockutil.GetConfigPDA(config.TestTimelockID),
					ac.PublicKey(),
					proposer.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, scerr)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment, []string{"Error Code: " + timelockutil.UnauthorizedError.String()})
			})

			raIx, raerr := access_controller.NewAddAccessInstruction(
				ac.PublicKey(),
				admin.PublicKey(),
				proposer.PublicKey(), // add access of proposer again
			).ValidateAndBuild()
			require.NoError(t, raerr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{raIx}, admin, config.DefaultCommitment)

			found, ferr := accesscontroller.HasAccess(ctx, solanaGoClient, ac.PublicKey(), proposer.PublicKey(), config.DefaultCommitment)
			require.NoError(t, ferr)
			require.True(t, found, "Account %s should be in the AccessList", proposer.PublicKey())

			salt, serr := timelockutil.SimpleSalt()
			require.NoError(t, serr)
			nonExecutableOp2 := timelockutil.Operation{
				TimelockID:  config.TestTimelockID,
				Predecessor: config.TimelockEmptyOpID,
				Salt:        salt,
				Delay:       uint64(1),
			}
			ix := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), timelockutil.GetSignerPDA(config.TestTimelockID)).Build()
			nonExecutableOp2.AddInstruction(ix, []solana.PublicKey{})

			ixs, prerr := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, nonExecutableOp2, proposer.PublicKey(), ac.PublicKey())
			require.NoError(t, prerr)
			for _, ix := range ixs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer, config.DefaultCommitment)
			}

			sbix, sberr := timelock.NewScheduleBatchInstruction(
				config.TestTimelockID,
				nonExecutableOp2.OperationID(),
				nonExecutableOp2.Delay,
				nonExecutableOp2.OperationPDA(), // formerly uploaded
				timelockutil.GetConfigPDA(config.TestTimelockID),
				ac.PublicKey(),
				proposer.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, sberr)

			tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{sbix}, proposer, config.DefaultCommitment)
			require.NotNil(t, tx)

			parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
				[]common.EventMapping{
					common.EventMappingFor[timelockutil.CallScheduled]("CallScheduled"),
				},
			)

			for i, ixx := range nonExecutableOp2.ToInstructionData() {
				event := parsedLogs[0].EventData[i].Data.(*timelockutil.CallScheduled)
				require.Equal(t, nonExecutableOp2.OperationID(), event.ID)
				require.Equal(t, uint64(i), event.Index)
				require.Equal(t, ixx.ProgramId, event.Target)
				require.Equal(t, nonExecutableOp2.Predecessor, event.Predecessor)
				require.Equal(t, nonExecutableOp2.Salt, event.Salt)
				require.Equal(t, nonExecutableOp2.Delay, event.Delay)
				require.Equal(t, ixx.Data, common.NormalizeData(event.Data))
			}

			var opAccount timelock.Operation
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, nonExecutableOp2.OperationPDA(), config.DefaultCommitment, &opAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}

			require.Equal(t,
				tx.BlockTime.Time().Add(time.Duration(nonExecutableOp2.Delay)*time.Second).Unix(),
				int64(opAccount.Timestamp),
				"Scheduled Times don't match",
			)
			require.Equal(t,
				nonExecutableOp2.OperationID(),
				opAccount.Id,
				"Ids don't match",
			)

			t.Run("rbac: cancel scheduled tx", func(t *testing.T) {
				t.Run("fail: should feed the right role access controller", func(t *testing.T) {
					signer := roleMap[timelock.Canceller_Role].RandomPick()
					ac := roleMap[timelock.Proposer_Role].AccessController

					ix, cerr := timelock.NewCancelInstruction(
						config.TestTimelockID,
						nonExecutableOp2.OperationID(),
						nonExecutableOp2.OperationPDA(),
						timelockutil.GetConfigPDA(config.TestTimelockID),
						ac.PublicKey(),
						signer.PublicKey(),
					).ValidateAndBuild()

					require.NoError(t, cerr)

					result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment, []string{"Error Code: " + "InvalidAccessController."})
					require.NotNil(t, result)
				})

				t.Run("fail: unauthorized on cancel attempt from non-canceller(proposer)", func(t *testing.T) {
					signer := roleMap[timelock.Proposer_Role].RandomPick()
					ac := roleMap[timelock.Canceller_Role].AccessController

					ix, cerr := timelock.NewCancelInstruction(
						config.TestTimelockID,
						nonExecutableOp2.OperationID(),
						nonExecutableOp2.OperationPDA(),
						timelockutil.GetConfigPDA(config.TestTimelockID),
						ac.PublicKey(),
						signer.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, cerr)

					result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment, []string{"Error Code: " + timelockutil.UnauthorizedError.String()})
					require.NotNil(t, result)
				})

				t.Run("success: Should able to cancel scheduled tx: PDA closed", func(t *testing.T) {
					signer := roleMap[timelock.Canceller_Role].RandomPick()
					ac := roleMap[timelock.Canceller_Role].AccessController

					ix, cerr := timelock.NewCancelInstruction(
						config.TestTimelockID,
						nonExecutableOp2.OperationID(),
						nonExecutableOp2.OperationPDA(),
						timelockutil.GetConfigPDA(config.TestTimelockID),
						ac.PublicKey(),
						signer.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, cerr)

					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment)
					require.NotNil(t, tx)

					parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
						[]common.EventMapping{
							common.EventMappingFor[timelockutil.Cancelled]("Cancelled"),
						},
					)

					for i := range nonExecutableOp2.ToInstructionData() {
						event := parsedLogs[0].EventData[i].Data.(*timelockutil.Cancelled)
						require.Equal(t, nonExecutableOp2.OperationID(), event.ID)
					}

					testutils.AssertClosedAccount(ctx, t, solanaGoClient, nonExecutableOp2.OperationPDA(), config.DefaultCommitment)
				})
			})
		})
	})

	t.Run("rbac: only_admin authentication ", func(t *testing.T) {
		newMinDelay := uint64(14000)

		t.Run("fail: only admin can call functions with only_admin macro", func(t *testing.T) {
			signer := roleMap[timelock.Proposer_Role].RandomPick()

			ix, ierr := timelock.NewUpdateDelayInstruction(
				config.TestTimelockID,
				newMinDelay,
				timelockutil.GetConfigPDA(config.TestTimelockID),
				signer.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, ierr)

			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment, []string{"Error Code: " + timelockutil.UnauthorizedError.String()})
			require.NotNil(t, result)
		})

		t.Run("success: only admin can call functions with only_admin macro", func(t *testing.T) {
			signer := admin

			var oldConfigAccount timelock.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment, &oldConfigAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}

			ix, err := timelock.NewUpdateDelayInstruction(
				config.TestTimelockID,
				newMinDelay,
				timelockutil.GetConfigPDA(config.TestTimelockID),
				signer.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment)
			require.NotNil(t, tx)

			parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
				[]common.EventMapping{
					common.EventMappingFor[timelockutil.MinDelayChange]("MinDelayChange"),
				},
			)

			event := parsedLogs[0].EventData[0].Data.(*timelockutil.MinDelayChange)
			require.Equal(t, oldConfigAccount.MinDelay, event.OldDuration)
			require.Equal(t, newMinDelay, event.NewDuration)

			var newConfigAccount timelock.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment, &newConfigAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
			require.Equal(t, newMinDelay, newConfigAccount.MinDelay, "MinDelay is not updated")
		})

		t.Run("rbac: preloading normal operation is only accessible from proposer and admin", func(t *testing.T) {
			cfgs := []struct {
				Name          string
				Ac            solana.PrivateKey
				Signer        solana.PrivateKey
				ShouldFail    bool
				ExpectedError string
			}{
				{
					Name:   "authorized proposer",
					Ac:     roleMap[timelock.Proposer_Role].AccessController,
					Signer: roleMap[timelock.Proposer_Role].RandomPick(),
				},
				{
					Name:          "proper proposer ac, unauthorized signer",
					Ac:            roleMap[timelock.Proposer_Role].AccessController, // valid
					Signer:        roleMap[timelock.Bypasser_Role].RandomPick(),     // unauthorized signer
					ShouldFail:    true,
					ExpectedError: "Error Code: " + timelockutil.UnauthorizedError.String(),
				},
				{
					Name:          "invalid access controller (bypasser)",
					Ac:            roleMap[timelock.Bypasser_Role].AccessController, // invalid access controller
					Signer:        roleMap[timelock.Bypasser_Role].RandomPick(),
					ShouldFail:    true,
					ExpectedError: "Error Code: " + timelock.InvalidAccessController_TimelockError.String(),
				},
				{
					Name:          "invalid access controller (canceller)",
					Ac:            roleMap[timelock.Canceller_Role].AccessController, // invalid access controller
					Signer:        roleMap[timelock.Canceller_Role].RandomPick(),
					ShouldFail:    true,
					ExpectedError: "Error Code: " + timelock.InvalidAccessController_TimelockError.String(),
				},
				{
					Name:          "invalid access controller (executor)",
					Ac:            roleMap[timelock.Executor_Role].AccessController, // invalid access controller
					Signer:        roleMap[timelock.Executor_Role].RandomPick(),
					ShouldFail:    true,
					ExpectedError: "Error Code: " + timelock.InvalidAccessController_TimelockError.String(),
				},
			}

			for _, cfg := range cfgs {
				t.Run(fmt.Sprintf("rbac: preloading operation test case - %s", cfg.Name), func(t *testing.T) {
					t.Parallel()
					signer := cfg.Signer
					ac := cfg.Ac

					salt, serr := timelockutil.SimpleSalt()
					require.NoError(t, serr)
					op := timelockutil.Operation{
						TimelockID:  config.TestTimelockID,
						Predecessor: config.TimelockEmptyOpID,
						Salt:        salt,
						Delay:       uint64(1),
					}
					ix := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), timelockutil.GetSignerPDA(config.TestTimelockID)).Build()
					op.AddInstruction(ix, []solana.PublicKey{})

					ixs, prerr := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, op, signer.PublicKey(), ac.PublicKey())
					require.NoError(t, prerr)
					if cfg.ShouldFail {
						testutils.SendAndFailWith(ctx, t, solanaGoClient, ixs, signer, config.DefaultCommitment, []string{cfg.ExpectedError})
					} else {
						for _, ix := range ixs {
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment)
						}
					}
				})
			}
		})
		t.Run("rbac: preloading bypasser operation is only accessible from bypasser and admin", func(t *testing.T) {
			cfgs := []struct {
				Name          string
				Ac            solana.PrivateKey
				Signer        solana.PrivateKey
				ShouldFail    bool
				ExpectedError string
			}{
				{
					Name:   "authorized bypasser",
					Ac:     roleMap[timelock.Bypasser_Role].AccessController,
					Signer: roleMap[timelock.Bypasser_Role].RandomPick(),
				},
				{
					Name:          "proper bypasser ac, unauthorized signer",
					Ac:            roleMap[timelock.Bypasser_Role].AccessController, // valid
					Signer:        roleMap[timelock.Proposer_Role].RandomPick(),     // unauthorized signer
					ShouldFail:    true,
					ExpectedError: "Error Code: " + timelockutil.UnauthorizedError.String(),
				},
				{
					Name:          "invalid access controller (canceller)",
					Ac:            roleMap[timelock.Canceller_Role].AccessController, // invalid access controller
					Signer:        roleMap[timelock.Canceller_Role].RandomPick(),
					ShouldFail:    true,
					ExpectedError: "Error Code: " + timelock.InvalidAccessController_TimelockError.String(),
				},
				{
					Name:          "invalid access controller (executor)",
					Ac:            roleMap[timelock.Executor_Role].AccessController, // invalid access controller
					Signer:        roleMap[timelock.Executor_Role].RandomPick(),
					ShouldFail:    true,
					ExpectedError: "Error Code: " + timelock.InvalidAccessController_TimelockError.String(),
				},
				{
					Name:          "invalid access controller (proposer)",
					Ac:            roleMap[timelock.Proposer_Role].AccessController, // invalid access controller
					Signer:        roleMap[timelock.Proposer_Role].RandomPick(),
					ShouldFail:    true,
					ExpectedError: "Error Code: " + timelock.InvalidAccessController_TimelockError.String(),
				},
			}

			for _, cfg := range cfgs {
				t.Run(fmt.Sprintf("rbac: preloading bypasser operation test case - %s", cfg.Name), func(t *testing.T) {
					t.Parallel()
					signer := cfg.Signer
					ac := cfg.Ac

					salt, serr := timelockutil.SimpleSalt()
					require.NoError(t, serr)
					op := timelockutil.Operation{
						TimelockID:   config.TestTimelockID,
						Predecessor:  config.TimelockEmptyOpID,
						Salt:         salt,
						Delay:        uint64(1),
						IsBypasserOp: true,
					}
					ix := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), timelockutil.GetSignerPDA(config.TestTimelockID)).Build()
					op.AddInstruction(ix, []solana.PublicKey{})

					ixs, prerr := timelockutil.GetPreloadBypasserOperationIxs(config.TestTimelockID, op, signer.PublicKey(), ac.PublicKey())
					require.NoError(t, prerr)
					if cfg.ShouldFail {
						testutils.SendAndFailWith(ctx, t, solanaGoClient, ixs, signer, config.DefaultCommitment, []string{cfg.ExpectedError})
					} else {
						for _, ix := range ixs {
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, signer, config.DefaultCommitment)
						}
					}
				})
			}
		})
	})
}
