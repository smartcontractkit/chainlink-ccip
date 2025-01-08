package contracts

import (
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
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
}
