package contracts

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/external_program_cpi_stub"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/accesscontroller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/mcms"
	timelockutil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestMcmsCapacity(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	mcm.SetProgramID(config.McmProgram)
	timelock.SetProgramID(config.TimelockProgram)
	access_controller.SetProgramID(config.AccessControllerProgram)
	external_program_cpi_stub.SetProgramID(config.ExternalCpiStubProgram) // test target program

	// initial admin
	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	anyone, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)

	msigs := map[timelock.Role]timelockutil.RoleMultisigs{
		timelock.Proposer_Role:  timelockutil.CreateRoleMultisigs(timelock.Proposer_Role, 1),
		timelock.Canceller_Role: timelockutil.CreateRoleMultisigs(timelock.Canceller_Role, 1),
		timelock.Executor_Role:  timelockutil.CreateRoleMultisigs(timelock.Executor_Role, 1),
		timelock.Bypasser_Role:  timelockutil.CreateRoleMultisigs(timelock.Bypasser_Role, 1),
	}

	require.NoError(t, err)

	t.Run("setup:funding", func(t *testing.T) {
		// fund EOAs
		testutils.FundAccounts(ctx, []solana.PrivateKey{admin, anyone}, solanaGoClient, t)
		pdaSignerAllowance := 1 * solana.LAMPORTS_PER_SOL

		// fund msig signers
		for _, roleMsigs := range msigs {
			ixs := make([]solana.Instruction, 0)
			for _, msig := range roleMsigs.Multisigs {
				fundPDAIx := system.NewTransferInstruction(pdaSignerAllowance, admin.PublicKey(), msig.SignerPDA).Build()
				ixs = append(ixs, fundPDAIx)
			}
			testutils.SendAndConfirm(ctx, t, solanaGoClient,
				ixs,
				admin, config.DefaultCommitment)
		}

		// fund timelock signer
		timelockSigner := timelockutil.GetSignerPDA(config.TestTimelockID)
		fundPDAIx := system.NewTransferInstruction(pdaSignerAllowance, admin.PublicKey(), timelockSigner).Build()
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{fundPDAIx}, admin, config.DefaultCommitment)
	})

	t.Run("setup: initialize mcm multisigs for each timelock role", func(t *testing.T) {
		for role, roleMsigs := range msigs {
			for _, msig := range roleMsigs.Multisigs {
				t.Run(fmt.Sprintf("init mcm for role %s with multisig %s", role.String(), mcms.UnpadString32(msig.PaddedID)), func(t *testing.T) {
					t.Parallel()
					// get program data account
					data, accErr := solanaGoClient.GetAccountInfoWithOpts(ctx, config.McmProgram, &rpc.GetAccountInfoOpts{
						Commitment: config.DefaultCommitment,
					})
					require.NoError(t, accErr)

					// decode program data
					var programData struct {
						DataType uint32
						Address  solana.PublicKey
					}
					require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

					ix, initIxErr := mcm.NewInitializeInstruction(
						config.TestChainID,
						msig.PaddedID,
						msig.ConfigPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
						config.McmProgram,
						programData.Address,
						msig.RootMetadataPDA,
						msig.ExpiringRootAndOpCountPDA,
					).ValidateAndBuild()
					require.NoError(t, initIxErr)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					// get config and validate
					var configAccount mcm.MultisigConfig
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigPDA, config.DefaultCommitment, &configAccount)
					require.NoError(t, err, "failed to get account info")

					require.Equal(t, config.TestChainID, configAccount.ChainId)
					require.Equal(t, admin.PublicKey(), configAccount.Owner)
					require.Equal(t, msig.PaddedID, configAccount.MultisigId)
				})
			}
		}
	})

	t.Run("setup: set_config for each mcm multisigs", func(t *testing.T) {
		for role, roleMsigs := range msigs {
			for _, msig := range roleMsigs.Multisigs {
				t.Run(fmt.Sprintf("set_config of role %s with multisig %s", role.String(), mcms.UnpadString32(msig.PaddedID)), func(t *testing.T) {
					t.Parallel()
					signerAddresses := msig.RawConfig.SignerAddresses

					t.Run("preload signers on PDA", func(t *testing.T) {
						ixs := make([]solana.Instruction, 0)
						//nolint:gosec
						parsedTotalSigners := uint8(len(signerAddresses))
						initSignersIx, initSignersIxErr := mcm.NewInitSignersInstruction(
							msig.PaddedID,
							parsedTotalSigners,
							msig.ConfigPDA,
							msig.ConfigSignersPDA,
							admin.PublicKey(),
							solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, initSignersIxErr)
						ixs = append(ixs, initSignersIx)

						appendSignersIxs, appendSignersIxsErr := mcms.GetAppendSignersIxs(signerAddresses, msig.PaddedID, msig.ConfigPDA, msig.ConfigSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
						require.NoError(t, appendSignersIxsErr)
						ixs = append(ixs, appendSignersIxs...)

						finalizeSignersIx, finSignersIxErr := mcm.NewFinalizeSignersInstruction(
							msig.PaddedID,
							msig.ConfigPDA,
							msig.ConfigSignersPDA,
							admin.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, finSignersIxErr)
						ixs = append(ixs, finalizeSignersIx)

						for _, ix := range ixs {
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
						}

						var cfgSignersAccount mcm.ConfigSigners
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
						require.NoError(t, err, "failed to get account info")

						require.Equal(t, true, cfgSignersAccount.IsFinalized)

						// check if the addresses are registered correctly
						for i, signer := range cfgSignersAccount.SignerAddresses {
							require.Equal(t, signerAddresses[i], signer)
						}
					})

					t.Run("success:set_config", func(t *testing.T) {
						// set config
						ix, setConfigErr := mcm.NewSetConfigInstruction(
							msig.PaddedID,
							msig.RawConfig.SignerGroups,
							msig.RawConfig.GroupQuorums,
							msig.RawConfig.GroupParents,
							msig.RawConfig.ClearRoot,
							msig.ConfigPDA,
							msig.ConfigSignersPDA,
							msig.RootMetadataPDA,
							msig.ExpiringRootAndOpCountPDA,
							admin.PublicKey(),
							solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, setConfigErr)

						tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
						require.NotNil(t, tx)

						parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
							[]common.EventMapping{
								common.EventMappingFor[mcms.ConfigSet]("ConfigSet"),
							},
						)

						event := parsedLogs[0].EventData[0].Data.(*mcms.ConfigSet)
						require.Equal(t, msig.RawConfig.GroupParents, event.GroupParents)
						require.Equal(t, msig.RawConfig.GroupQuorums, event.GroupQuorums)
						require.Equal(t, msig.RawConfig.ClearRoot, event.IsRootCleared)
						for i, signer := range event.Signers {
							require.Equal(t, msig.RawConfig.SignerAddresses[i], signer.EvmAddress)
							require.Equal(t, uint8(i), signer.Index)
							require.Equal(t, msig.RawConfig.SignerGroups[i], signer.Group)
						}

						// get config and validate
						var configAccount mcm.MultisigConfig
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigPDA, config.DefaultCommitment, &configAccount)
						require.NoError(t, err, "failed to get account info")

						require.Equal(t, config.TestChainID, configAccount.ChainId)
						require.Equal(t, reflect.DeepEqual(configAccount.GroupParents, msig.RawConfig.GroupParents), true)
						require.Equal(t, reflect.DeepEqual(configAccount.GroupQuorums, msig.RawConfig.GroupQuorums), true)

						// check if the McmSigner struct is correct
						for i, signer := range configAccount.Signers {
							require.Equal(t, signer.EvmAddress, msig.RawConfig.SignerAddresses[i])
							require.Equal(t, signer.Index, uint8(i))
							require.Equal(t, signer.Group, msig.RawConfig.SignerGroups[i])
						}

						// signers pda closed after set_config
						testutils.AssertClosedAccount(ctx, t, solanaGoClient, msig.ConfigSignersPDA, config.DefaultCommitment)
					})
				})
			}
		}
	})

	t.Run("setup: timelock", func(t *testing.T) {
		for role, roleMsigs := range msigs {
			t.Run(fmt.Sprintf("init access controller for role %s", role.String()), func(t *testing.T) {
				initAccIxs, initAccErr := timelockutil.GetInitAccessControllersIxs(ctx, roleMsigs.AccessController.PublicKey(), admin, solanaGoClient)
				require.NoError(t, initAccErr)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, initAccIxs, admin, config.DefaultCommitment, common.AddSigners(roleMsigs.AccessController))

				var ac access_controller.AccessController
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, roleMsigs.AccessController.PublicKey(), config.DefaultCommitment, &ac)
				if err != nil {
					require.NoError(t, err, "failed to get account info")
				}
			})
		}

		t.Run("initialize timelock with access controllers", func(t *testing.T) {
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

			initTimelockIx, initTimelockErr := timelock.NewInitializeInstruction(
				config.TestTimelockID,
				config.MinDelay,
				timelockutil.GetConfigPDA(config.TestTimelockID),
				admin.PublicKey(),
				solana.SystemProgramID,
				config.TimelockProgram,
				programData.Address,
				config.AccessControllerProgram,
				msigs[timelock.Proposer_Role].AccessController.PublicKey(),
				msigs[timelock.Executor_Role].AccessController.PublicKey(),
				msigs[timelock.Canceller_Role].AccessController.PublicKey(),
				msigs[timelock.Bypasser_Role].AccessController.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, initTimelockErr)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initTimelockIx}, admin, config.DefaultCommitment)

			var configAccount timelock.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockutil.GetConfigPDA(config.TestTimelockID), config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}

			require.Equal(t, admin.PublicKey(), configAccount.Owner, "Owner doesn't match")
			require.Equal(t, config.MinDelay, configAccount.MinDelay, "MinDelay doesn't match")
			require.Equal(t, msigs[timelock.Proposer_Role].AccessController.PublicKey(), configAccount.ProposerRoleAccessController, "ProposerRoleAccessController doesn't match")
			require.Equal(t, msigs[timelock.Executor_Role].AccessController.PublicKey(), configAccount.ExecutorRoleAccessController, "ExecutorRoleAccessController doesn't match")
			require.Equal(t, msigs[timelock.Canceller_Role].AccessController.PublicKey(), configAccount.CancellerRoleAccessController, "CancellerRoleAccessController doesn't match")
			require.Equal(t, msigs[timelock.Bypasser_Role].AccessController.PublicKey(), configAccount.BypasserRoleAccessController, "BypasserRoleAccessController doesn't match")
		})

		t.Run("register msig signers to each role", func(t *testing.T) {
			for role, roleMsigs := range msigs {
				t.Run(fmt.Sprintf("registering role %s", role.String()), func(t *testing.T) {
					t.Parallel()
					addresses := []solana.PublicKey{}
					for _, msig := range roleMsigs.Multisigs {
						addresses = append(addresses, msig.SignerPDA)
					}
					batchAddAccessIxs, batchAddAccessErr := timelockutil.GetBatchAddAccessIxs(ctx, config.TestTimelockID, roleMsigs.AccessController.PublicKey(), role, addresses, admin, config.BatchAddAccessChunkSize, solanaGoClient)
					require.NoError(t, batchAddAccessErr)

					for _, ix := range batchAddAccessIxs {
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
					}

					for _, msig := range roleMsigs.Multisigs {
						found, ferr := accesscontroller.HasAccess(ctx, solanaGoClient, roleMsigs.AccessController.PublicKey(), msig.SignerPDA, config.DefaultCommitment)
						require.NoError(t, ferr)
						require.True(t, found, "Account %s not found in %s AccessList", msig.SignerPDA, role)
					}
				})
			}
		})
	})

	t.Run("setup: transfer ownership multisigs to timelock signer", func(t *testing.T) {
		for role, roleMsigs := range msigs {
			t.Run(fmt.Sprintf("transfer ownership of role %s", role.String()), func(t *testing.T) {
				t.Parallel()
				for _, msig := range roleMsigs.Multisigs {
					t.Run(fmt.Sprintf(role.String(), mcms.UnpadString32(msig.PaddedID)), func(t *testing.T) {
						ix, transferOwnershipErr := mcm.NewTransferOwnershipInstruction(
							msig.PaddedID,
							timelockutil.GetSignerPDA(config.TestTimelockID), // new proposed owner
							msig.ConfigPDA,
							admin.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, transferOwnershipErr)

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

						var configAccount mcm.MultisigConfig
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigPDA, config.DefaultCommitment, &configAccount)
						if err != nil {
							require.NoError(t, err, "failed to get account info")
						}
						require.Equal(t, admin.PublicKey(), configAccount.Owner)
						require.Equal(t, timelockutil.GetSignerPDA(config.TestTimelockID), configAccount.ProposedOwner)

						acceptOwnershipIx, acceptOwnershipixErr := mcm.NewAcceptOwnershipInstruction(
							msig.PaddedID,
							msig.ConfigPDA,
							timelockutil.GetSignerPDA(config.TestTimelockID),
						).ValidateAndBuild()
						require.NoError(t, acceptOwnershipixErr)

						salt, sErr := timelockutil.SimpleSalt()
						require.NoError(t, sErr)
						acceptOwnershipOp := timelockutil.Operation{
							TimelockID:   config.TestTimelockID,
							Predecessor:  config.TimelockEmptyOpID,
							Salt:         salt,
							Delay:        uint64(1),
							IsBypasserOp: true,
						}

						acceptOwnershipOp.AddInstruction(acceptOwnershipIx, []solana.PublicKey{config.McmProgram})

						id := acceptOwnershipOp.OperationID()
						operationPDA := acceptOwnershipOp.OperationPDA()

						ixs, ierr := timelockutil.GetPreloadBypasserOperationIxs(config.TestTimelockID, acceptOwnershipOp, admin.PublicKey(), roleMsigs.AccessController.PublicKey())
						require.NoError(t, ierr)
						for _, ix := range ixs {
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
						}

						var opAccount timelock.Operation
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, operationPDA, config.DefaultCommitment, &opAccount)
						if err != nil {
							require.NoError(t, err, "failed to get account info")
						}

						require.Equal(t,
							id,
							opAccount.Id,
							"Ids don't match",
						)

						bypassExeIx := timelock.NewBypasserExecuteBatchInstruction(
							config.TestTimelockID,
							acceptOwnershipOp.OperationID(),
							acceptOwnershipOp.OperationPDA(),
							timelockutil.GetConfigPDA(config.TestTimelockID),
							timelockutil.GetSignerPDA(config.TestTimelockID),
							roleMsigs.AccessController.PublicKey(),
							admin.PublicKey(), // bypass execute with admin previledges
						)
						bypassExeIx.AccountMetaSlice = append(bypassExeIx.AccountMetaSlice, acceptOwnershipOp.RemainingAccounts()...)

						vIx, vIxErr := bypassExeIx.ValidateAndBuild()
						require.NoError(t, vIxErr)

						acceptTx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment)

						parsedLogs := common.ParseLogMessages(acceptTx.Meta.LogMessages,
							[]common.EventMapping{
								common.EventMappingFor[timelockutil.BypasserCallExecuted]("BypasserCallExecuted"),
							},
						)

						for i, ixx := range acceptOwnershipOp.ToInstructionData() {
							event := parsedLogs[0].EventData[i].Data.(*timelockutil.BypasserCallExecuted)
							require.Equal(t, uint64(i), event.Index)
							require.Equal(t, ixx.ProgramId, event.Target)
							require.Equal(t, ixx.Data, common.NormalizeData(event.Data))
						}

						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigPDA, config.DefaultCommitment, &configAccount)
						if err != nil {
							require.NoError(t, err, "failed to get account info")
						}
						require.Equal(t, timelockutil.GetSignerPDA(config.TestTimelockID), configAccount.Owner)
						require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
					})
				}
			})
		}
	})

	t.Run("overhead analysis", func(t *testing.T) {
		// -------------------------
		// Direct No-Op Overhead
		// -------------------------
		directNoOpOverhead := func(t *testing.T) uint32 {
			// Create the no-op instruction.
			noOpDirectIx, nierr := external_program_cpi_stub.NewNoOpInstruction().ValidateAndBuild()
			require.NoError(t, nierr)

			// Measure its direct execution cost.
			directCU := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{noOpDirectIx}, admin, config.DefaultCommitment)
			t.Logf("Direct no-op CU: %d", directCU)

			// Execute to confirm it works.
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{noOpDirectIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(directCU))
			return uint32(directCU)
		}

		// -------------------------
		// MCM Execute Overhead
		// -------------------------
		mcmExecuteOverhead := func(t *testing.T, directCU uint32) uint32 {
			// Retrieve the executor multisig info.
			executorMsig := msigs[timelock.Executor_Role].GetAnyMultisig()

			// Get the current operation count.
			var prevRootAndOpCount mcm.ExpiringRootAndOpCount
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, executorMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &prevRootAndOpCount)
			require.NoError(t, err, "failed to get account info")
			currOpCount := prevRootAndOpCount.OpCount

			// Create a no-op instruction.
			noOpIx, nierr := external_program_cpi_stub.NewNoOpInstruction().ValidateAndBuild()
			require.NoError(t, nierr)

			// Wrap the no-op in an MCM operation node.
			node, nerr := mcms.IxToMcmTestOpNode(
				executorMsig.ConfigPDA,
				executorMsig.SignerPDA,
				noOpIx,
				currOpCount,
			)
			require.NoError(t, nerr)
			ops := []mcms.McmOpNode{node}

			// Create the MCM root data.
			validUntil := uint32(0xffffffff)
			rootValidationData, rverr := mcms.CreateMcmRootData(mcms.McmRootInput{
				Multisig:             executorMsig.ConfigPDA,
				Operations:           ops,
				PreOpCount:           currOpCount,
				PostOpCount:          currOpCount + 1,
				ValidUntil:           validUntil,
				OverridePreviousRoot: false,
			})
			require.NoError(t, rverr)

			// Sign and set the root.
			signatures, serr := mcms.BulkSignOnMsgHash(executorMsig.Signers, rootValidationData.EthMsgHash)
			require.NoError(t, serr)
			signaturesPDA := executorMsig.RootSignaturesPDA(rootValidationData.Root, validUntil, admin.PublicKey())
			preloadIxs, perr := mcms.GetMcmPreloadSignaturesIxs(
				signatures,
				executorMsig.PaddedID,
				rootValidationData.Root,
				validUntil,
				admin.PublicKey(),
				config.MaxAppendSignatureBatchSize,
			)
			require.NoError(t, perr)
			for _, ix := range preloadIxs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			setRootIx, srerr := mcm.NewSetRootInstruction(
				executorMsig.PaddedID,
				rootValidationData.Root,
				validUntil,
				rootValidationData.Metadata,
				rootValidationData.MetadataProof,
				signaturesPDA,
				executorMsig.RootMetadataPDA,
				mcms.GetSeenSignedHashesPDA(executorMsig.PaddedID, rootValidationData.Root, validUntil),
				executorMsig.ExpiringRootAndOpCountPDA,
				executorMsig.ConfigPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, srerr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{setRootIx}, admin, config.DefaultCommitment)

			// Build the MCM execute instruction.
			op := ops[0]
			proofs, perr := op.Proofs()
			require.NoError(t, perr)
			executeIx := mcm.NewExecuteInstruction(
				executorMsig.PaddedID,
				config.TestChainID,
				op.Nonce,
				op.Data,
				proofs,
				executorMsig.ConfigPDA,
				executorMsig.RootMetadataPDA,
				executorMsig.ExpiringRootAndOpCountPDA,
				op.To,
				executorMsig.SignerPDA,
				admin.PublicKey(),
			)
			executeIx.AccountMetaSlice = append(executeIx.AccountMetaSlice, op.RemainingAccounts...)
			vIx, vierr := executeIx.ValidateAndBuild()
			require.NoError(t, vierr)

			// Measure and log the compute units (CU) for MCM execution.
			mcmCU := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment)
			t.Logf("MCM execute CU: %d", mcmCU)

			// Execute the MCM operation.
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(mcmCU))
			overhead := uint32(mcmCU) - directCU
			t.Logf("MCM overhead: %d", overhead)
			return overhead
		}

		// -------------------------
		// Timelock Execute Overhead
		// -------------------------
		timelockExecuteOverhead := func(t *testing.T, directCU uint32) uint32 {
			// Build a timelock operation that wraps a no-op.
			salt, serr := timelockutil.SimpleSalt()
			require.NoError(t, serr)
			op := timelockutil.Operation{
				TimelockID:  config.TestTimelockID,
				Predecessor: config.TimelockEmptyOpID,
				Salt:        salt,
				Delay:       uint64(1),
			}
			noOpIx, nierr := external_program_cpi_stub.NewNoOpInstruction().ValidateAndBuild()
			require.NoError(t, nierr)
			op.AddInstruction(noOpIx, []solana.PublicKey{config.ExternalCpiStubProgram})

			// Preload and schedule the timelock operation.
			opIxs, oierr := timelockutil.GetPreloadOperationIxs(
				config.TestTimelockID,
				op,
				admin.PublicKey(),
				msigs[timelock.Proposer_Role].AccessController.PublicKey(),
			)
			require.NoError(t, oierr)
			for _, ix := range opIxs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}
			scheduleIx, scerr := timelock.NewScheduleBatchInstruction(
				config.TestTimelockID,
				op.OperationID(),
				op.Delay,
				op.OperationPDA(),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				msigs[timelock.Proposer_Role].AccessController.PublicKey(),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, scerr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{scheduleIx}, admin, config.DefaultCommitment)

			// Wait until the timelock operation is ready.
			err = timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, op.OperationPDA(), config.DefaultCommitment)
			require.NoError(t, err)

			// Build and measure the timelock execute instruction.
			timelockExeIx := timelock.NewExecuteBatchInstruction(
				config.TestTimelockID,
				op.OperationID(),
				op.OperationPDA(),
				config.TimelockEmptyOpID,
				timelockutil.GetConfigPDA(config.TestTimelockID),
				timelockutil.GetSignerPDA(config.TestTimelockID),
				msigs[timelock.Executor_Role].AccessController.PublicKey(),
				admin.PublicKey(),
			)
			timelockExeIx.AccountMetaSlice = append(timelockExeIx.AccountMetaSlice, op.RemainingAccounts()...)
			timelockExeVIx, tverr := timelockExeIx.ValidateAndBuild()
			require.NoError(t, tverr)

			timelockCU := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{timelockExeVIx}, admin, config.DefaultCommitment)
			t.Logf("Timelock execute CU: %d", timelockCU)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{timelockExeVIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(timelockCU))
			overhead := uint32(timelockCU) - directCU
			t.Logf("Timelock overhead: %d", overhead)
			return overhead
		}

		// -------------------------
		// Helper function for the scaling test using compute-heavy instructions.
		// Instead of adding many no-op instructions (which causes OOM),
		// we create each instruction as a compute-heavy instruction.
		// 'iterationsPerInstr' can be tuned to achieve desired CU usage.
		// -------------------------
		createTimelockOpWithHeavyOps := func(t *testing.T, count int, iterationsPerInstr uint32) timelockutil.Operation {
			salt, serr := timelockutil.SimpleSalt()
			require.NoError(t, serr)
			op := timelockutil.Operation{
				TimelockID:  config.TestTimelockID,
				Predecessor: config.TimelockEmptyOpID,
				Salt:        salt,
				Delay:       uint64(1),
			}
			for i := 0; i < count; i++ {
				heavyIx, hverr := external_program_cpi_stub.NewComputeHeavyInstruction(iterationsPerInstr).ValidateAndBuild()
				require.NoError(t, hverr)
				op.AddInstruction(heavyIx, []solana.PublicKey{config.ExternalCpiStubProgram})
			}
			return op
		}

		executeTimelockOpAndGetCU := func(t *testing.T, op timelockutil.Operation) uint32 {
			opIxs, oierr := timelockutil.GetPreloadOperationIxs(
				config.TestTimelockID,
				op,
				admin.PublicKey(),
				msigs[timelock.Proposer_Role].AccessController.PublicKey(),
			)
			require.NoError(t, oierr)
			for _, ix := range opIxs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}
			scheduleIx, scerr := timelock.NewScheduleBatchInstruction(
				config.TestTimelockID,
				op.OperationID(),
				op.Delay,
				op.OperationPDA(),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				msigs[timelock.Proposer_Role].AccessController.PublicKey(),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, scerr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{scheduleIx}, admin, config.DefaultCommitment)
			err = timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, op.OperationPDA(), config.DefaultCommitment)
			require.NoError(t, err)
			timelockExeIx := timelock.NewExecuteBatchInstruction(
				config.TestTimelockID,
				op.OperationID(),
				op.OperationPDA(),
				config.TimelockEmptyOpID,
				timelockutil.GetConfigPDA(config.TestTimelockID),
				timelockutil.GetSignerPDA(config.TestTimelockID),
				msigs[timelock.Executor_Role].AccessController.PublicKey(),
				admin.PublicKey(),
			)
			timelockExeIx.AccountMetaSlice = append(timelockExeIx.AccountMetaSlice, op.RemainingAccounts()...)
			timelockExeVIx, tverr := timelockExeIx.ValidateAndBuild()
			require.NoError(t, tverr)
			cu := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{timelockExeVIx}, admin, config.DefaultCommitment)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{timelockExeVIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(cu))
			return uint32(cu)
		}

		// -------------------------
		// Run the tests.
		// -------------------------
		t.Run("Direct No-Op", func(t *testing.T) {
			directCU := directNoOpOverhead(t)
			t.Logf("Direct no-op CU: %d", directCU)
		})

		t.Run("MCM Execute Overhead", func(t *testing.T) {
			directCU := directNoOpOverhead(t)
			overhead := mcmExecuteOverhead(t, directCU)
			t.Logf("MCM Execute Overhead: %d", overhead)
		})

		t.Run("Timelock Execute Overhead", func(t *testing.T) {
			directCU := directNoOpOverhead(t)
			overhead := timelockExecuteOverhead(t, directCU)
			t.Logf("Timelock Execute Overhead: %d", overhead)
		})

		t.Run("Timelock Scaling Test (Compute-Heavy Ops)", func(t *testing.T) {
			// For scaling tests, we now use compute-heavy instructions instead of many no-ops.
			// Each instruction in the timelock op will be a compute-heavy instruction with a fixed iteration count.
			const iterationsPerInstr uint32 = 1000
			directCU := directNoOpOverhead(t)

			// Measure a timelock op with 1 compute-heavy instruction to establish baseline.
			totalCU1 := executeTimelockOpAndGetCU(t, createTimelockOpWithHeavyOps(t, 1, iterationsPerInstr))
			baselineOverhead := totalCU1 - directCU
			t.Logf("Baseline overhead (1 heavy op): %d", baselineOverhead)

			// Test with multiple instructions.
			counts := []int{7, 8} // 9 sometimes fails due to exceeding the compute unit limit

			for _, count := range counts {
				totalCU := executeTimelockOpAndGetCU(t, createTimelockOpWithHeavyOps(t, count, iterationsPerInstr))
				perInstrOverhead := (totalCU - directCU - baselineOverhead) / uint32(count-1)
				t.Logf("For %d heavy ops: totalCU=%d, baseline=%d, per-instruction overhead=%d",
					count, totalCU, baselineOverhead, perInstrOverhead)
			}
		})
	})

	/*
		NOTE: Testing Transaction Size Limitations
		- Analyzing maximum transaction size constraints through two approaches:
			1. Single instruction capacity (mcm::execute with memo program)
			2. Multiple instruction batching with multiple mints (mcm::execute with timelock::schedule_batch / timelock::execute_batch)
			3. Multiple instruction batching with memo program - for operation payload test (mcm::execute with timelock::schedule_batch / timelock::execute_batch)

		Resource Constraints
		- Compute Units (CU): Maximum 1.4M CU per transaction(hard cap)
		- Memory: Program specific heap allocation limits
		- Transaction Size: 1232 bytes maximum on-chain packet size

		Measurement Methodology
		- Account meta size: 32 bytes (pubkey) + 1 byte (signer) + 1 byte (writable)
		- Instruction data measured separately from account metadata
		- Note: Measurement excludes transaction overhead (~128 bytes for signatures, message header)
		- Measured size = account metas + instruction data

		Testing Approach
		1. MCM Test with Memo Program:
			- Tests raw instruction data capacity
			- Single account structure (minimal metadata overhead)
			- Finds maximum data payload size for mcm::execute

		2. MCM + Timelock Test:
			- Tests combined instruction data and account metadata limits
			- Analysis of multi-instruction batching capacity
			- Memory constraints from operation size and account lookups
			- Real-world throughput limitations for token transfers

		3. MCM + Timelock Test with Memo Program:
			- Tests raw instruction data capacity
			- Finds maximum data payload size for timelock preloading sequences with mcm::execute { timelock::schedule_batch }
	*/
	t.Run("mcms tx size analysis", func(t *testing.T) {
		// helper to measure instruction size
		measureInstructionSize := func(ix solana.Instruction) int {
			metaSize := 0
			for range ix.Accounts() {
				metaSize += 32 + 1 + 1 // account keys, is_signer, is_writable
			}
			data, derr := ix.Data()
			require.NoError(t, derr)
			return metaSize + len(data)
		}

		t.Run("mcm::execute max tx size analysis", func(t *testing.T) {
			t.Parallel()
			// use executor multisig for testing
			executorMsig := msigs[timelock.Executor_Role].GetAnyMultisig()

			type memoTestSummary struct {
				name        string
				memoSize    int
				mcmOpSize   int
				finalTxSize int
				expectError bool
			}

			testCases := []struct {
				name        string
				memoSize    int
				expectError bool
			}{
				{"max_memo", 759, false},
				{"too_big_memo", 760, true},
			}

			testSummaries := make([]memoTestSummary, 0, len(testCases))

			// get current root and op count for nonce management
			var prevRootAndOpCount mcm.ExpiringRootAndOpCount
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, executorMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &prevRootAndOpCount)
			require.NoError(t, err, "failed to get account info")

			for opIdx, tc := range testCases {
				currOpCount := prevRootAndOpCount.OpCount + uint64(opIdx)
				tcResult := memoTestSummary{
					name:        tc.name,
					memoSize:    tc.memoSize,
					expectError: tc.expectError,
				}

				t.Run(tc.name, func(t *testing.T) {
					// create memo data of specified size
					memoText := strings.Repeat("A", tc.memoSize)

					memoIx := solana.NewInstruction(
						solana.MemoProgramID,
						[]*solana.AccountMeta{
							solana.Meta(executorMsig.SignerPDA).SIGNER(),
						},
						[]byte(memoText),
					)

					// convert to MCM operation
					node, nerr := mcms.IxToMcmTestOpNode(
						executorMsig.ConfigPDA,
						executorMsig.SignerPDA,
						memoIx,
						currOpCount, // nonce
					)
					require.NoError(t, nerr)

					// measure MCM operation size
					tcResult.mcmOpSize = len(node.Data)

					// build verifiable ops
					ops := []mcms.McmOpNode{node}

					// create and validate root
					validUntil := uint32(0xffffffff)
					rootValidationData, rverr := mcms.CreateMcmRootData(mcms.McmRootInput{
						Multisig:             executorMsig.ConfigPDA,
						Operations:           ops,
						PreOpCount:           currOpCount,
						PostOpCount:          currOpCount + 1,
						ValidUntil:           validUntil,
						OverridePreviousRoot: false,
					})
					require.NoError(t, rverr)

					// sign and set root
					signatures, sigerr := mcms.BulkSignOnMsgHash(executorMsig.Signers, rootValidationData.EthMsgHash)
					require.NoError(t, sigerr)

					signaturesPDA := executorMsig.RootSignaturesPDA(rootValidationData.Root, validUntil, admin.PublicKey())

					// preload signatures
					preloadIxs, prerr := mcms.GetMcmPreloadSignaturesIxs(
						signatures,
						executorMsig.PaddedID,
						rootValidationData.Root,
						validUntil,
						admin.PublicKey(),
						config.MaxAppendSignatureBatchSize,
					)
					require.NoError(t, prerr)

					for _, ix := range preloadIxs {
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
					}

					// Set root
					setRootIx, srerr := mcm.NewSetRootInstruction(
						executorMsig.PaddedID,
						rootValidationData.Root,
						validUntil,
						rootValidationData.Metadata,
						rootValidationData.MetadataProof,
						signaturesPDA,
						executorMsig.RootMetadataPDA,
						mcms.GetSeenSignedHashesPDA(executorMsig.PaddedID, rootValidationData.Root, validUntil),
						executorMsig.ExpiringRootAndOpCountPDA,
						executorMsig.ConfigPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, srerr)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{setRootIx}, admin, config.DefaultCommitment)

					// Execute MCM operation
					proofs, prferr := ops[0].Proofs()
					require.NoError(t, prferr)

					executeIx := mcm.NewExecuteInstruction(
						executorMsig.PaddedID,
						config.TestChainID,
						node.Nonce,
						node.Data,
						proofs,
						executorMsig.ConfigPDA,
						executorMsig.RootMetadataPDA,
						executorMsig.ExpiringRootAndOpCountPDA,
						node.To,
						executorMsig.SignerPDA,
						admin.PublicKey(),
					)
					executeIx.AccountMetaSlice = append(executeIx.AccountMetaSlice, node.RemainingAccounts...)

					vIx, verr := executeIx.ValidateAndBuild()
					require.NoError(t, verr)

					// measure final transaction size
					tcResult.finalTxSize = measureInstructionSize(vIx)

					if tc.expectError {
						testutils.SendAndFailWithRPCError(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment, []string{"solana_sdk::transaction::versioned::VersionedTransaction too large"}, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
					} else {
						tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))

						parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
							[]common.EventMapping{
								common.EventMappingFor[mcms.OpExecuted]("OpExecuted"),
							},
						)
						event := parsedLogs[0].EventData[0].Data.(*mcms.OpExecuted)
						require.Equal(t, node.Nonce, event.Nonce)
						require.Equal(t, node.To, event.To)
						require.Equal(t, node.Data, common.NormalizeData(event.Data))
					}

					testSummaries = append(testSummaries, tcResult)
				})
			}

			t.Run("result_summary", func(t *testing.T) {
				t.Logf("\n================= MCM Execute Size Analysis =================")
				t.Logf("%-20s %-12s %-16s %-14s",
					"Test Name",
					"Message Size",
					"McmOpSize",
					"Final TxSize",
				)

				for _, result := range testSummaries {
					t.Logf("%-20s %-12d %-16d %-14d",
						result.name,
						result.memoSize,
						result.mcmOpSize,
						result.finalTxSize,
					)
				}
			})
		})

		t.Run("mcm::execute{timelock::schedule_batch}/timelock::execute_batch max txs analysis", func(t *testing.T) {
			t.Parallel()
			mintKeypair, mkerr := solana.NewRandomPrivateKey()
			require.NoError(t, mkerr)
			mint := mintKeypair.PublicKey()

			tokenProgram := config.Token2022Program

			createTokenIxs, cterr := tokens.CreateToken(
				ctx,
				tokenProgram,
				mint,
				admin.PublicKey(),
				9, // decimals
				solanaGoClient,
				config.DefaultCommitment,
			)
			require.NoError(t, cterr)

			treasury, pkerr := solana.NewRandomPrivateKey()
			require.NoError(t, pkerr)

			fundIx := system.NewTransferInstruction(
				1*solana.LAMPORTS_PER_SOL,
				admin.PublicKey(),
				treasury.PublicKey(),
			).Build()

			treasuryIx, treasuryATA, taerr := tokens.CreateAssociatedTokenAccount(
				tokenProgram,
				mint,
				treasury.PublicKey(),
				admin.PublicKey(),
			)
			require.NoError(t, taerr)

			// mint initial tokens to treasury
			mintIx, merr := tokens.MintTo(
				1000*solana.LAMPORTS_PER_SOL,
				tokenProgram,
				mint,
				treasuryATA,
				admin.PublicKey(),
			)
			require.NoError(t, merr)

			setupIxs := append(createTokenIxs, fundIx, treasuryIx, mintIx)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, setupIxs, admin, config.DefaultCommitment, common.AddSigners(mintKeypair))

			// setup mint authority to Timelock signer
			authIx, auerr := tokens.SetTokenMintAuthority(tokenProgram, timelockutil.GetSignerPDA(config.TestTimelockID), mint, admin.PublicKey())
			require.NoError(t, auerr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{authIx}, admin, config.DefaultCommitment)

			// use proposer multisig for testing
			proposerMsig := msigs[timelock.Proposer_Role].GetAnyMultisig()

			testCases := []struct {
				name         string
				numTransfers int
				expectError  bool
			}{
				{"max_batch_16", 16, false},
				{"oom_batch_17", 17, true}, // OOM on timelock::execute_batch
			}

			type timelockTestSummary struct {
				name                         string
				numTransfers                 int
				totalTransferInstructionSize int
				mcmOpSize                    int
				finalMcmTxSize               int
				timelockOpPDASize            int
				timelockTxSize               int
				expectError                  bool
			}

			allowance := uint64(0)
			for _, tc := range testCases {
				allowance += uint64(tc.numTransfers) * solana.LAMPORTS_PER_SOL
			}

			approveIx, aperr := tokens.TokenApproveChecked(
				allowance,
				9,
				tokenProgram,
				treasuryATA,
				mint,
				timelockutil.GetSignerPDA(config.TestTimelockID),
				treasury.PublicKey(),
				nil,
			)
			require.NoError(t, aperr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{approveIx}, treasury, config.DefaultCommitment)

			testSummaries := make([]timelockTestSummary, 0, len(testCases))

			// get current root and op count
			var prevRootAndOpCount mcm.ExpiringRootAndOpCount
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &prevRootAndOpCount)
			require.NoError(t, err, "failed to get account info")

			for opIdx, tc := range testCases {
				currOpCount := prevRootAndOpCount.OpCount + uint64(opIdx)
				tcResult := timelockTestSummary{
					name:         tc.name,
					numTransfers: tc.numTransfers,
					expectError:  tc.expectError,
				}

				t.Run(tc.name, func(t *testing.T) {
					recipients := make([]struct {
						key solana.PrivateKey
						ata solana.PublicKey
					}, tc.numTransfers)

					createAtaIxs := make([]solana.Instruction, 0, tc.numTransfers)
					transferIxs := make([]solana.Instruction, 0, tc.numTransfers)
					totalSize := 0

					// create recipient accounts and transfer instructions
					for i := 0; i < tc.numTransfers; i++ {
						recipientKey, rkerr := solana.NewRandomPrivateKey()
						require.NoError(t, rkerr)
						recipients[i].key = recipientKey

						// create ATA for recipient
						ataIx, ata, aterr := tokens.CreateAssociatedTokenAccount(
							tokenProgram,
							mint,
							recipientKey.PublicKey(),
							admin.PublicKey(),
						)
						require.NoError(t, aterr)

						createAtaIxs = append(createAtaIxs, ataIx)
						recipients[i].ata = ata
						// Create transfer instruction
						transferIx, tierr := tokens.TokenTransferChecked(
							1*solana.LAMPORTS_PER_SOL, // Amount
							9,                         // Decimals
							tokenProgram,
							treasuryATA,
							mint,
							ata,
							timelockutil.GetSignerPDA(config.TestTimelockID),
							[]solana.PublicKey{},
						)
						require.NoError(t, tierr)

						transferIxs = append(transferIxs, transferIx)

						ixSize := measureInstructionSize(transferIx)
						totalSize += ixSize
					}

					tcResult.totalTransferInstructionSize = totalSize

					batchSize := 10
					for i := 0; i < len(createAtaIxs); i += batchSize {
						end := i + batchSize
						if end > len(createAtaIxs) {
							end = len(createAtaIxs)
						}

						batch := createAtaIxs[i:end]
						testutils.SendAndConfirm(ctx, t, solanaGoClient, batch, admin, config.DefaultCommitment)
					}

					// create timelock operation for the batch
					salt, serr := timelockutil.SimpleSalt()
					require.NoError(t, serr)

					op := timelockutil.Operation{
						TimelockID:  config.TestTimelockID,
						Predecessor: config.TimelockEmptyOpID,
						Salt:        salt,
						Delay:       uint64(1),
					}

					// add all transfer instructions to the operation
					for _, transferIx := range transferIxs {
						op.AddInstruction(transferIx, []solana.PublicKey{tokenProgram})
					}

					// create and initialize operation accounts
					ixs, prerr := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, op, admin.PublicKey(), msigs[timelock.Proposer_Role].AccessController.PublicKey())
					require.NoError(t, prerr)
					for _, ix := range ixs {
						cu := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(cu))
					}

					info, ierr := solanaGoClient.GetAccountInfoWithOpts(ctx, op.OperationPDA(), &rpc.GetAccountInfoOpts{
						Commitment: config.DefaultCommitment,
					})
					require.NoError(t, ierr)
					tcResult.timelockOpPDASize = len(info.Value.Data.GetBinary())

					// schedule the operation
					scheduleIx, scerr := timelock.NewScheduleBatchInstruction(
						config.TestTimelockID,
						op.OperationID(),
						op.Delay,
						op.OperationPDA(),
						timelockutil.GetConfigPDA(config.TestTimelockID),
						msigs[timelock.Proposer_Role].AccessController.PublicKey(),
						proposerMsig.SignerPDA,
					).ValidateAndBuild()
					require.NoError(t, scerr)

					// convert to MCM operation
					node, nerr := mcms.IxToMcmTestOpNode(
						proposerMsig.ConfigPDA,
						proposerMsig.SignerPDA,
						scheduleIx,
						currOpCount, // nonce
					)
					require.NoError(t, nerr)

					// measure MCM operation size
					tcResult.mcmOpSize = len(node.Data)

					// build verifiable ops
					ops := []mcms.McmOpNode{node}

					// create and validate root
					validUntil := uint32(0xffffffff)
					rootValidationData, rverr := mcms.CreateMcmRootData(mcms.McmRootInput{
						Multisig:             proposerMsig.ConfigPDA,
						Operations:           ops,
						PreOpCount:           currOpCount,
						PostOpCount:          currOpCount + 1,
						ValidUntil:           validUntil,
						OverridePreviousRoot: false,
					})
					require.NoError(t, rverr)

					// sign and set root
					signatures, serr := mcms.BulkSignOnMsgHash(proposerMsig.Signers, rootValidationData.EthMsgHash)
					require.NoError(t, serr)

					signaturesPDA := proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil, admin.PublicKey())

					// preload signatures
					preloadIxs, prerr := mcms.GetMcmPreloadSignaturesIxs(
						signatures,
						proposerMsig.PaddedID,
						rootValidationData.Root,
						validUntil,
						admin.PublicKey(),
						config.MaxAppendSignatureBatchSize,
					)
					require.NoError(t, prerr)

					for _, ix := range preloadIxs {
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
					}

					// set root
					setRootIx, srerr := mcm.NewSetRootInstruction(
						proposerMsig.PaddedID,
						rootValidationData.Root,
						validUntil,
						rootValidationData.Metadata,
						rootValidationData.MetadataProof,
						signaturesPDA,
						proposerMsig.RootMetadataPDA,
						mcms.GetSeenSignedHashesPDA(proposerMsig.PaddedID, rootValidationData.Root, validUntil),
						proposerMsig.ExpiringRootAndOpCountPDA,
						proposerMsig.ConfigPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, srerr)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{setRootIx}, admin, config.DefaultCommitment)

					// execute MCM operation
					proofs, prerr := ops[0].Proofs()
					require.NoError(t, prerr)

					executeIx := mcm.NewExecuteInstruction(
						proposerMsig.PaddedID,
						config.TestChainID,
						node.Nonce,
						node.Data,
						proofs,
						proposerMsig.ConfigPDA,
						proposerMsig.RootMetadataPDA,
						proposerMsig.ExpiringRootAndOpCountPDA,
						node.To,
						proposerMsig.SignerPDA,
						admin.PublicKey(),
					)
					executeIx.AccountMetaSlice = append(executeIx.AccountMetaSlice, node.RemainingAccounts...)

					vIx, verr := executeIx.ValidateAndBuild()
					require.NoError(t, verr)

					// measure final transaction size
					tcResult.finalMcmTxSize = measureInstructionSize(vIx)

					// mcm::execute { timelock::schedule_batch }
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment)

					err = timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, op.OperationPDA(), config.DefaultCommitment)
					require.NoError(t, err)

					tlExeIx := timelock.NewExecuteBatchInstruction(
						config.TestTimelockID,
						op.OperationID(),
						op.OperationPDA(),
						config.TimelockEmptyOpID,
						timelockutil.GetConfigPDA(config.TestTimelockID),
						timelockutil.GetSignerPDA(config.TestTimelockID),
						msigs[timelock.Executor_Role].AccessController.PublicKey(),
						admin.PublicKey(),
					)
					tlExeIx.AccountMetaSlice = append(tlExeIx.AccountMetaSlice, op.RemainingAccounts()...)

					tivIx, txerr := tlExeIx.ValidateAndBuild()
					require.NoError(t, txerr)

					// measure execution size
					tcResult.timelockTxSize = measureInstructionSize(tivIx)

					if tc.expectError {
						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{tivIx}, admin, config.DefaultCommitment, []string{"Program log: Error: memory allocation failed, out of memory"}, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
					} else {
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{tivIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))

						// verify recipients final balances,
						// note: we can't verify the exact amount of treasury balance(tests in parallel execution)
						for _, recipient := range recipients {
							_, balance, berr := tokens.TokenBalance(ctx, solanaGoClient, recipient.ata, config.DefaultCommitment)
							require.NoError(t, berr)
							require.Equal(t, 1*int(solana.LAMPORTS_PER_SOL), balance, "Recipient should have received 1 token")
						}
					}
					testSummaries = append(testSummaries, tcResult)
				})
			}

			t.Run("result_summary", func(t *testing.T) {
				t.Logf("\n================= Timelock Test Summary =================")
				t.Logf("%-20s %-12s %-24s %-16s %-14s %-16s %-16s",
					"Test Name",
					"#Transfers",
					"TransferInstrSize",
					"McmOpSize",
					"McmTxSize",
					"timelockOpPDA",
					"TimelockTxSize",
				)

				for _, result := range testSummaries {
					t.Logf("%-20s %-12d %-24d %-16d %-14d %-16d %-16d",
						result.name,
						result.numTransfers,
						result.totalTransferInstructionSize,
						result.mcmOpSize,
						result.finalMcmTxSize,
						result.timelockOpPDASize,
						result.timelockTxSize,
					)
				}
			})
		})

		t.Run("mcm::execute{timelock::schedule_batch}/timelock::execute_batch for operation capacity analysis", func(t *testing.T) {
			proposerMsig := msigs[timelock.Proposer_Role].GetAnyMultisig()

			testCases := []struct {
				name        string
				memoSize    int
				expectError bool
			}{
				{"oom_cpi_instruction", 5096, false},
				{"oom_cpi_instruction", 5097, true}, // OOM on timelock::execute_batch
			}

			for _, tc := range testCases {
				// get current root and op count for nonce management
				var proposerPrevRootAndOpCount mcm.ExpiringRootAndOpCount
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &proposerPrevRootAndOpCount)
				require.NoError(t, err, "failed to get account info")

				currOpCount := proposerPrevRootAndOpCount.OpCount

				t.Run(tc.name, func(t *testing.T) {
					// pure test function that only log the length of the instruction data
					bix, berr := external_program_cpi_stub.NewBigInstructionDataInstruction(
						[]byte(strings.Repeat("A", tc.memoSize)),
					).ValidateAndBuild()
					require.NoError(t, berr)

					salt, serr := timelockutil.SimpleSalt()
					require.NoError(t, serr)

					timelockOp := timelockutil.Operation{
						TimelockID:  config.TestTimelockID,
						Predecessor: config.TimelockEmptyOpID,
						Salt:        salt,
						Delay:       uint64(1),
					}
					timelockOp.AddInstruction(bix, []solana.PublicKey{config.ExternalCpiStubProgram})

					ixs, perr := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, timelockOp, proposerMsig.SignerPDA, msigs[timelock.Proposer_Role].AccessController.PublicKey())
					require.NoError(t, perr)

					mcmOpNodes := []mcms.McmOpNode{}
					for i, ix := range ixs {
						opNode, oerr := mcms.IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, ix, currOpCount+uint64(i))
						require.NoError(t, oerr)
						// append preloading instructions
						mcmOpNodes = append(mcmOpNodes, opNode)
					}

					scheduleIx, scerr := timelock.NewScheduleBatchInstruction(
						config.TestTimelockID,
						timelockOp.OperationID(),
						timelockOp.Delay,
						timelockOp.OperationPDA(),
						timelockutil.GetConfigPDA(config.TestTimelockID),
						msigs[timelock.Proposer_Role].AccessController.PublicKey(),
						proposerMsig.SignerPDA,
					).ValidateAndBuild()
					require.NoError(t, scerr)

					opNode, onerr := mcms.IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, scheduleIx, currOpCount+uint64(len(mcmOpNodes)))
					require.NoError(t, onerr)
					// append schedule instruction
					mcmOpNodes = append(mcmOpNodes, opNode)

					// Create and validate root data
					validUntil := uint32(0xffffffff)

					rootValidationData, rverr := mcms.CreateMcmRootData(mcms.McmRootInput{
						Multisig:             proposerMsig.ConfigPDA,
						Operations:           mcmOpNodes,
						PreOpCount:           currOpCount,
						PostOpCount:          currOpCount + uint64(len(mcmOpNodes)),
						ValidUntil:           validUntil,
						OverridePreviousRoot: false,
					})
					require.NoError(t, rverr)

					signatures, serr := mcms.BulkSignOnMsgHash(proposerMsig.Signers, rootValidationData.EthMsgHash)
					require.NoError(t, serr)

					signaturesPDA := proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil, admin.PublicKey())

					// preload signatures
					preloadIxs, plerr := mcms.GetMcmPreloadSignaturesIxs(signatures, proposerMsig.PaddedID, rootValidationData.Root, validUntil, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
					require.NoError(t, plerr)
					for _, ix := range preloadIxs {
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
					}

					// set root
					setRootIx, srerr := mcm.NewSetRootInstruction(
						proposerMsig.PaddedID,
						rootValidationData.Root,
						validUntil,
						rootValidationData.Metadata,
						rootValidationData.MetadataProof,
						signaturesPDA,
						proposerMsig.RootMetadataPDA,
						mcms.GetSeenSignedHashesPDA(proposerMsig.PaddedID, rootValidationData.Root, validUntil),
						proposerMsig.ExpiringRootAndOpCountPDA,
						proposerMsig.ConfigPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, srerr)

					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{setRootIx}, admin, config.DefaultCommitment)
					require.NotNil(t, tx)

					parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
						[]common.EventMapping{
							common.EventMappingFor[mcms.NewRoot]("NewRoot"),
						},
					)
					event := parsedLogs[0].EventData[0].Data.(*mcms.NewRoot)
					require.Equal(t, rootValidationData.Root, event.Root)
					require.Equal(t, validUntil, event.ValidUntil)
					require.Equal(t, rootValidationData.Metadata.ChainId, event.MetadataChainID)
					require.Equal(t, proposerMsig.ConfigPDA, event.MetadataMultisig)
					require.Equal(t, rootValidationData.Metadata.PreOpCount, event.MetadataPreOpCount)
					require.Equal(t, rootValidationData.Metadata.PostOpCount, event.MetadataPostOpCount)
					require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, event.MetadataOverridePreviousRoot)

					for _, op := range mcmOpNodes {
						// execute mcm operation to preload and schedule the timelock operation
						proofs, perr := op.Proofs()
						require.NoError(t, perr)

						executeIx := mcm.NewExecuteInstruction(
							proposerMsig.PaddedID,
							config.TestChainID,
							op.Nonce,
							op.Data,
							proofs,
							proposerMsig.ConfigPDA,
							proposerMsig.RootMetadataPDA,
							proposerMsig.ExpiringRootAndOpCountPDA,
							op.To,
							proposerMsig.SignerPDA,
							anyone.PublicKey(),
						)
						executeIx.AccountMetaSlice = append(executeIx.AccountMetaSlice, op.RemainingAccounts...)

						vIx, verr := executeIx.ValidateAndBuild()
						require.NoError(t, verr)

						exeTx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, anyone, config.DefaultCommitment)
						require.NotNil(t, exeTx)

						if bytes.Equal(op.Data[:8], timelock.Instruction_ScheduleBatch[:]) {
							parsedExeLogs := common.ParseLogMessages(exeTx.Meta.LogMessages,
								[]common.EventMapping{
									common.EventMappingFor[mcms.OpExecuted]("OpExecuted"),
									common.EventMappingFor[timelockutil.CallScheduled]("CallScheduled"),
								},
							)
							exeEvent := parsedExeLogs[0].EventData[0].Data.(*mcms.OpExecuted)
							require.Equal(t, op.Nonce, exeEvent.Nonce)
							require.Equal(t, op.To, exeEvent.To)
							require.Equal(t, op.Data, common.NormalizeData(exeEvent.Data))

							// check inner CallScheduled events
							opIxData := timelockOp.ToInstructionData()

							require.Equal(t, len(opIxData), len(parsedExeLogs[0].InnerCalls[0].EventData), "Number of actual CallScheduled events does not match expected for operation")

							for j, ix := range opIxData {
								timelockEvent := parsedExeLogs[0].InnerCalls[0].EventData[j].Data.(*timelockutil.CallScheduled)
								require.Equal(t, timelockOp.OperationID(), timelockEvent.ID, "ID does not match")
								require.Equal(t, uint64(j), timelockEvent.Index, "Index does not match")
								require.Equal(t, ix.ProgramId, timelockEvent.Target, "Target does not match")
								require.Equal(t, timelockOp.Predecessor, timelockEvent.Predecessor, "Predecessor does not match")
								require.Equal(t, timelockOp.Salt, timelockEvent.Salt, "Salt does not match")
								require.Equal(t, timelockOp.Delay, timelockEvent.Delay, "Delay does not match")
								require.Equal(t, ix.Data, common.NormalizeData(timelockEvent.Data), "Data does not match")
							}
						}
					}

					// read timelock operation
					var opAccount timelock.Operation
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockOp.OperationPDA(), config.DefaultCommitment, &opAccount)
					if err != nil {
						require.NoError(t, err, "failed to get account info")
					}
					require.Equal(t, timelockOp.OperationID(), opAccount.Id, "Operation ID does not match")
					require.Equal(t, timelock.Scheduled_OperationState, opAccount.State, "Operation state is invalid")

					werr := timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, timelockOp.OperationPDA(), config.DefaultCommitment)
					require.NoError(t, werr)

					ix := timelock.NewExecuteBatchInstruction(
						config.TestTimelockID,
						timelockOp.OperationID(),
						timelockOp.OperationPDA(),
						config.TimelockEmptyOpID,
						timelockutil.GetConfigPDA(config.TestTimelockID),
						timelockutil.GetSignerPDA(config.TestTimelockID),
						msigs[timelock.Executor_Role].AccessController.PublicKey(),
						admin.PublicKey(), // any authority, timelock worker
					)

					ix.AccountMetaSlice = append(ix.AccountMetaSlice, timelockOp.RemainingAccounts()...)

					vIx, vErr := ix.ValidateAndBuild()
					require.NoError(t, vErr)

					if tc.expectError {
						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment, []string{"Program log: Error: memory allocation failed, out of memory"}, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
					} else {
						tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))

						parsedLogs = common.ParseLogMessages(tx.Meta.LogMessages,
							[]common.EventMapping{
								common.EventMappingFor[timelockutil.CallExecuted]("CallExecuted"),
							},
						)

						for i, ixx := range timelockOp.ToInstructionData() {
							event := parsedLogs[0].EventData[i].Data.(*timelockutil.CallExecuted)
							require.Equal(t, timelockOp.OperationID(), event.ID)
							require.Equal(t, uint64(i), event.Index)
							require.Equal(t, ixx.ProgramId, event.Target)
							require.Equal(t, ixx.Data, common.NormalizeData(event.Data))
						}

						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, timelockOp.OperationPDA(), config.DefaultCommitment, &opAccount)
						if err != nil {
							require.NoError(t, err, "failed to get account info")
						}

						require.Equal(t,
							timelock.Done_OperationState,
							opAccount.State,
							"Executed operation should be marked as done",
						)
					}
				})
			}
		})
	})
}
