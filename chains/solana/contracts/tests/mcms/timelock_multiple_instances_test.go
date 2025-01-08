package contracts

import (
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
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/mcms"
	timelockutil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/timelock"
)

func TestTimelockMultipleInstances(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	timelock.SetProgramID(config.TimelockProgram)
	access_controller.SetProgramID(config.AccessControllerProgram)

	deployer, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, deployer)

	type TimelockInstance struct {
		ID        [32]byte
		Admin     solana.PrivateKey
		ConfigPDA solana.PublicKey
		SignerPDA solana.PublicKey
		Roles     []timelockutil.RoleAccounts
		RoleMap   timelockutil.RoleMap
	}

	admin1, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	admin2, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	testTimelockID1, err := mcms.PadString32("test_timelock_instance_1")
	require.NoError(t, err)

	testTimelockID2, err := mcms.PadString32("test_timelock_instance_2")
	require.NoError(t, err)

	roles1, roleMap1 := timelockutil.TestRoleAccounts(1)
	roles2, roleMap2 := timelockutil.TestRoleAccounts(1)

	timelockInstances := []TimelockInstance{
		{
			ID:        testTimelockID1,
			Admin:     admin1,
			ConfigPDA: timelockutil.GetConfigPDA(testTimelockID1),
			SignerPDA: timelockutil.GetSignerPDA(testTimelockID1),
			Roles:     roles1,
			RoleMap:   roleMap1,
		},
		{
			ID:        testTimelockID2,
			Admin:     admin2,
			ConfigPDA: timelockutil.GetConfigPDA(testTimelockID2),
			SignerPDA: timelockutil.GetSignerPDA(testTimelockID2),
			Roles:     roles2,
			RoleMap:   roleMap2,
		},
	}

	t.Run("setup:funding", func(t *testing.T) {
		all := []solana.PrivateKey{}
		all = append(all, deployer)
		for _, instance := range timelockInstances {
			all = append(all, instance.Admin)
			for _, role := range instance.Roles {
				all = append(all, role.Accounts...)
			}
		}
		testutils.FundAccounts(ctx, all, solanaGoClient, t)
	})

	t.Run("setup:init access controllers", func(t *testing.T) {
		for _, instance := range timelockInstances {
			for _, data := range instance.RoleMap {
				initAccIxs, ierr := timelockutil.GetInitAccessControllersIxs(ctx, data.AccessController.PublicKey(), instance.Admin, solanaGoClient)
				require.NoError(t, ierr)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, initAccIxs, instance.Admin, config.DefaultCommitment, common.AddSigners(data.AccessController))

				var ac access_controller.AccessController
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, data.AccessController.PublicKey(), config.DefaultCommitment, &ac)
				if err != nil {
					require.NoError(t, err, "failed to get account info")
				}
			}
		}
	})

	t.Run("setup:init timelock instances", func(t *testing.T) {
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

		for _, instance := range timelockInstances {
			initTimelockIx, ierr := timelock.NewInitializeInstruction(
				instance.ID,
				config.MinDelay,
				instance.ConfigPDA,
				deployer.PublicKey(), // deployer is the owner of the timelock program initially
				solana.SystemProgramID,
				config.TimelockProgram,
				programData.Address,
				config.AccessControllerProgram,
				instance.RoleMap[timelock.Proposer_Role].AccessController.PublicKey(),
				instance.RoleMap[timelock.Executor_Role].AccessController.PublicKey(),
				instance.RoleMap[timelock.Canceller_Role].AccessController.PublicKey(),
				instance.RoleMap[timelock.Bypasser_Role].AccessController.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, ierr)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initTimelockIx}, deployer, config.DefaultCommitment)

			var configAccount timelock.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, instance.ConfigPDA, config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}

			require.Equal(t, deployer.PublicKey(), configAccount.Owner, "The initial owner doesn't match")
			require.Equal(t, instance.ID, configAccount.TimelockId, "TimelockID doesn't match")
			require.Equal(t, config.MinDelay, configAccount.MinDelay, "MinDelay doesn't match")
			require.Equal(t, instance.RoleMap[timelock.Proposer_Role].AccessController.PublicKey(), configAccount.ProposerRoleAccessController, "ProposerRoleAccessController doesn't match")
			require.Equal(t, instance.RoleMap[timelock.Executor_Role].AccessController.PublicKey(), configAccount.ExecutorRoleAccessController, "ExecutorRoleAccessController doesn't match")
			require.Equal(t, instance.RoleMap[timelock.Canceller_Role].AccessController.PublicKey(), configAccount.CancellerRoleAccessController, "CancellerRoleAccessController doesn't match")
			require.Equal(t, instance.RoleMap[timelock.Bypasser_Role].AccessController.PublicKey(), configAccount.BypasserRoleAccessController, "BypasserRoleAccessController doesn't match")
		}
	})

	t.Run("transfer config ownership(admin role) to each admin", func(t *testing.T) {
		for _, instance := range timelockInstances {
			t.Run("transfer ownership", func(t *testing.T) {
				ix, ierr := timelock.NewTransferOwnershipInstruction(
					instance.ID,
					instance.Admin.PublicKey(),
					instance.ConfigPDA,
					deployer.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, ierr)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, deployer, config.DefaultCommitment)
				require.NotNil(t, result)
			})

			t.Run("instance.Admin becomes owner", func(t *testing.T) {
				ix, ierr := timelock.NewAcceptOwnershipInstruction(
					instance.ID,
					instance.ConfigPDA,
					instance.Admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, ierr)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, instance.Admin, config.DefaultCommitment)
				require.NotNil(t, result)

				// Validate proposed set to 0-address after accepting ownership
				var configAccount timelock.Config
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockutil.GetConfigPDA(instance.ID), config.DefaultCommitment, &configAccount)
				if err != nil {
					require.NoError(t, err, "failed to get account info")
				}
				require.Equal(t, instance.Admin.PublicKey(), configAccount.Owner)
				require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
			})
		}
	})

	t.Run("setup:register access list & verify", func(t *testing.T) {
		for _, instance := range timelockInstances {
			for role, data := range instance.RoleMap {
				addresses := []solana.PublicKey{}
				for _, account := range data.Accounts {
					addresses = append(addresses, account.PublicKey())
				}
				batchAddAccessIxs, baerr := timelockutil.GetBatchAddAccessIxs(ctx, instance.ID, data.AccessController.PublicKey(), role, addresses, instance.Admin, config.BatchAddAccessChunkSize, solanaGoClient)
				require.NoError(t, baerr)

				for _, ix := range batchAddAccessIxs {
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, instance.Admin, config.DefaultCommitment)
				}

				for _, account := range data.Accounts {
					found, ferr := accesscontroller.HasAccess(ctx, solanaGoClient, data.AccessController.PublicKey(), account.PublicKey(), config.DefaultCommitment)
					require.NoError(t, ferr)
					require.True(t, found, "Account %s not found in %s AccessList", account.PublicKey(), data.Role)
				}
			}
		}
	})

	t.Run("instance isolation tests", func(t *testing.T) {
		t.Run("config isolation", func(t *testing.T) {
			t.Run("cannot access other instance config using own roles", func(t *testing.T) {
				// try to use instance1's proposer to schedule on instance2
				proposer := timelockInstances[0].RoleMap[timelock.Proposer_Role].RandomPick()
				salt, err := mcms.SimpleSalt()
				require.NoError(t, err)

				// create test operation for instance2 but using instance1's proposer
				op := timelockutil.Operation{
					TimelockID:  timelockInstances[1].ID,
					Predecessor: config.TimelockEmptyOpID,
					Salt:        salt,
					Delay:       uint64(1),
				}

				// preload operation instructions
				preloadIxs, prierr := timelockutil.GetPreloadOperationIxs(
					timelockInstances[1].ID,
					op,
					proposer.PublicKey(),
				)
				require.NoError(t, prierr)

				for _, ix := range preloadIxs {
					testutils.SendAndConfirm(
						ctx,
						t,
						solanaGoClient,
						[]solana.Instruction{ix},
						proposer,
						config.DefaultCommitment,
					)
				}

				// Try to schedule operation on instance2 using instance1's proposer
				ix, ierr := timelock.NewScheduleBatchInstruction(
					timelockInstances[1].ID, // instance2's ID
					op.OperationID(),
					op.Delay,
					op.OperationPDA(),
					timelockInstances[1].ConfigPDA, // instance2's config
					timelockInstances[0].RoleMap[timelock.Proposer_Role].AccessController.PublicKey(), // instance1's proposer AC
					proposer.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, ierr)

				result := testutils.SendAndFailWith(
					ctx,
					t,
					solanaGoClient,
					[]solana.Instruction{ix},
					proposer, // signing with instance1's proposer
					config.DefaultCommitment,
					[]string{"Error Code: " + timelock.InvalidAccessController_TimelockError.String()},
				)
				require.NotNil(t, result)
			})

			t.Run("cannot update delay using other instance's admin", func(t *testing.T) {
				// try to update instance2's delay using instance1's admin
				newMinDelay := uint64(14000)

				ix, ierr := timelock.NewUpdateDelayInstruction(
					timelockInstances[1].ID, // instance2's ID
					newMinDelay,
					timelockInstances[1].ConfigPDA,
					timelockInstances[0].Admin.PublicKey(), // instance1's admin trying to update
				).ValidateAndBuild()
				require.NoError(t, ierr)

				result := testutils.SendAndFailWith(
					ctx,
					t,
					solanaGoClient,
					[]solana.Instruction{ix},
					timelockInstances[0].Admin, // signing with instance1's admin
					config.DefaultCommitment,
					[]string{"Error Code: " + timelockutil.UnauthorizedError.String()},
				)
				require.NotNil(t, result)
			})

			t.Run("cannot modify access controllers of another instance", func(t *testing.T) {
				// try to use instance1's admin to modify instance2's access controller
				randomKey, err := solana.NewRandomPrivateKey()
				require.NoError(t, err)

				// try to add access to instance2's proposer role using instance1's admin
				ix, ierr := access_controller.NewAddAccessInstruction(
					timelockInstances[1].RoleMap[timelock.Proposer_Role].AccessController.PublicKey(), // instance2's proposer AC
					timelockInstances[0].Admin.PublicKey(),                                            // instance1's admin trying to modify
					randomKey.PublicKey(),                                                             // random key to add
				).ValidateAndBuild()
				require.NoError(t, ierr)

				result := testutils.SendAndFailWith(
					ctx,
					t,
					solanaGoClient,
					[]solana.Instruction{ix},
					timelockInstances[0].Admin, // signing with instance1's admin
					config.DefaultCommitment,
					[]string{"Error Code: ConstraintHasOne"},
				)
				require.NotNil(t, result)
			})

			t.Run("cannot transfer ownership of another instance", func(t *testing.T) {
				// try to transfer instance2's ownership using instance1's admin
				ix, ierr := timelock.NewTransferOwnershipInstruction(
					timelockInstances[1].ID,                // instance2's ID
					timelockInstances[0].Admin.PublicKey(), // trying to transfer to instance1's admin
					timelockInstances[1].ConfigPDA,
					timelockInstances[0].Admin.PublicKey(), // instance1's admin trying to transfer
				).ValidateAndBuild()
				require.NoError(t, ierr)

				result := testutils.SendAndFailWith(
					ctx,
					t,
					solanaGoClient,
					[]solana.Instruction{ix},
					timelockInstances[0].Admin, // signing with instance1's admin
					config.DefaultCommitment,
					[]string{"Error Code: " + timelockutil.UnauthorizedError.String()},
				)
				require.NotNil(t, result)
			})
		})

		t.Run("operation isolation", func(t *testing.T) {
			t.Run("cannot execute operation scheduled on another instance", func(t *testing.T) {
				// Schedule operation on instance1
				proposer1 := timelockInstances[0].RoleMap[timelock.Proposer_Role].RandomPick()
				executor2 := timelockInstances[1].RoleMap[timelock.Executor_Role].RandomPick()

				salt, err := mcms.SimpleSalt()
				require.NoError(t, err)

				// Create and schedule operation on instance1
				op := timelockutil.Operation{
					TimelockID:  timelockInstances[0].ID,
					Predecessor: config.TimelockEmptyOpID,
					Salt:        salt,
					Delay:       uint64(1),
				}

				// Add simple transfer instruction
				ix := system.NewTransferInstruction(
					1*solana.LAMPORTS_PER_SOL,
					proposer1.PublicKey(),
					timelockInstances[0].SignerPDA,
				).Build()
				op.AddInstruction(ix, []solana.PublicKey{})

				// Schedule on instance1
				preloadIxs, prierr := timelockutil.GetPreloadOperationIxs(
					timelockInstances[0].ID,
					op,
					proposer1.PublicKey(),
				)
				require.NoError(t, prierr)
				for _, ix := range preloadIxs {
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer1, config.DefaultCommitment)
				}

				scheduleIx, ierr := timelock.NewScheduleBatchInstruction(
					timelockInstances[0].ID,
					op.OperationID(),
					op.Delay,
					op.OperationPDA(),
					timelockInstances[0].ConfigPDA,
					timelockInstances[0].RoleMap[timelock.Proposer_Role].AccessController.PublicKey(),
					proposer1.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, ierr)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{scheduleIx}, proposer1, config.DefaultCommitment)

				// Wait for operation to be ready
				err = timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, op.OperationPDA(), config.DefaultCommitment)
				require.NoError(t, err)

				// try to execute instance1's operation using instance2's executor
				executeIx := timelock.NewExecuteBatchInstruction(
					timelockInstances[1].ID, // Use instance2's ID
					op.OperationID(),
					op.OperationPDA(),              // instance1's operation PDA
					config.TimelockEmptyOpID,       // empty predecessor
					timelockInstances[1].ConfigPDA, // instance2's config
					timelockInstances[1].SignerPDA, // instance2's signer
					timelockInstances[1].RoleMap[timelock.Executor_Role].AccessController.PublicKey(),
					executor2.PublicKey(),
				)
				executeIx.AccountMetaSlice = append(executeIx.AccountMetaSlice, op.RemainingAccounts()...)

				vIx, err := executeIx.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(
					ctx,
					t,
					solanaGoClient,
					[]solana.Instruction{vIx},
					executor2,
					config.DefaultCommitment,
					[]string{"Error Code: ConstraintSeeds."},
				)
			})
			t.Run("cannot cancel operation from another instance", func(t *testing.T) {
				// Schedule operation on instance1
				proposer1 := timelockInstances[0].RoleMap[timelock.Proposer_Role].RandomPick()
				salt, err := mcms.SimpleSalt()
				require.NoError(t, err)

				// Create test operation for instance1
				op := timelockutil.Operation{
					TimelockID:  timelockInstances[0].ID,
					Predecessor: config.TimelockEmptyOpID,
					Salt:        salt,
					Delay:       uint64(1),
				}

				// Add simple transfer instruction
				ix := system.NewTransferInstruction(
					1*solana.LAMPORTS_PER_SOL,
					proposer1.PublicKey(),
					timelockInstances[0].SignerPDA,
				).Build()
				op.AddInstruction(ix, []solana.PublicKey{})

				// Schedule on instance1
				preloadIxs, prierr := timelockutil.GetPreloadOperationIxs(
					timelockInstances[0].ID,
					op,
					proposer1.PublicKey(),
				)
				require.NoError(t, prierr)
				for _, ix := range preloadIxs {
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, proposer1, config.DefaultCommitment)
				}

				scheduleIx, ierr := timelock.NewScheduleBatchInstruction(
					timelockInstances[0].ID,
					op.OperationID(),
					op.Delay,
					op.OperationPDA(),
					timelockInstances[0].ConfigPDA,
					timelockInstances[0].RoleMap[timelock.Proposer_Role].AccessController.PublicKey(),
					proposer1.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, ierr)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{scheduleIx}, proposer1, config.DefaultCommitment)

				// Try to cancel instance1's operation using instance2's canceller
				canceller2 := timelockInstances[1].RoleMap[timelock.Canceller_Role].RandomPick()

				cancelIx, cerr := timelock.NewCancelInstruction(
					timelockInstances[1].ID, // Using instance2's ID
					op.OperationID(),
					op.OperationPDA(),
					timelockInstances[1].ConfigPDA,
					timelockInstances[1].RoleMap[timelock.Canceller_Role].AccessController.PublicKey(),
					canceller2.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, cerr)

				testutils.SendAndFailWith(
					ctx,
					t,
					solanaGoClient,
					[]solana.Instruction{cancelIx},
					canceller2,
					config.DefaultCommitment,
					[]string{"Error Code: ConstraintSeeds"},
				)
			})
		})

		t.Run("role isolation", func(t *testing.T) {
			t.Run("proposer cannot schedule on other instance", func(t *testing.T) {})
			t.Run("executor cannot execute on other instance", func(t *testing.T) {})
			t.Run("canceller cannot cancel on other instance", func(t *testing.T) {})
			t.Run("bypasser cannot bypass on other instance", func(t *testing.T) {})
		})
	})
}
