package contracts

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"

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

func TestMcmWithTimelock(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	mcm.SetProgramID(config.McmProgram)
	timelock.SetProgramID(config.TimelockProgram)
	access_controller.SetProgramID(config.AccessControllerProgram)

	// test target program
	external_program_cpi_stub.SetProgramID(config.ExternalCpiStubProgram)

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
		testutils.FundAccounts(ctx, []solana.PrivateKey{admin, anyone}, solanaGoClient, t)

		// fund msig PDA signers
		for _, roleMsigs := range msigs {
			ixs := make([]solana.Instruction, 0)
			for _, msig := range roleMsigs.Multisigs {
				fundPDAIx := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), msig.SignerPDA).Build()
				ixs = append(ixs, fundPDAIx)
			}
			testutils.SendAndConfirm(ctx, t, solanaGoClient,
				ixs,
				admin, config.DefaultCommitment)
		}
		// fund timelock signer
		fundPDAIx := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), timelockutil.GetSignerPDA(config.TestTimelockID)).Build()
		testutils.SendAndConfirm(ctx, t, solanaGoClient,
			[]solana.Instruction{fundPDAIx},
			admin, config.DefaultCommitment)
	})

	t.Run("setup: initialize mcm multisigs", func(t *testing.T) {
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

						// pda closed after set_config
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

	// shared proposer msig for testing OpCount in metadata
	proposerMsig := msigs[timelock.Proposer_Role].GetAnyMultisig()
	// keep track of operation count, this will be updated after each operation
	currentOpCount := 0

	// NOTE: These tests are not designed to run in parallel, we're testing opCount also(using currentOpCount).
	t.Run("mcm + timelock mint operation", func(t *testing.T) {
		for _, v := range []struct {
			tokenName    string
			tokenProgram solana.PublicKey
		}{
			{tokenName: "spl-token", tokenProgram: solana.TokenProgramID},
			{tokenName: "spl-token-2022", tokenProgram: config.Token2022Program},
		} {
			t.Run(v.tokenName, func(t *testing.T) {
				mintKeypair, mintKeypairErr := solana.NewRandomPrivateKey()
				require.NoError(t, mintKeypairErr)
				mint := mintKeypair.PublicKey()

				// Use CreateToken utility to get initialization instructions
				// NOTE: can't create token with cpi(mint signature required)
				createTokenIxs, createTokenErr := tokens.CreateToken(
					ctx,
					v.tokenProgram,    // token program
					mint,              // mint account
					admin.PublicKey(), // initial mint owner(admin)
					9,                 // decimals
					solanaGoClient,
					config.DefaultCommitment,
				)
				require.NoError(t, createTokenErr)

				for _, ix := range createTokenIxs {
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, common.AddSigners(mintKeypair))
				}

				t.Run("mint ixs", func(t *testing.T) {
					// 1. mcm set_root and execute through multisigs
					// 1-1. mcm:: set_root { root of { timelock::schedule_batch { spl::mint }}}
					// 1-2. pre-create operation PDA & upload instructions with timelock::initialize_operation, append_instructions, finalize_operation
					// 1-3. mcm:: execute { timelock::schedule_batch { spl::mint }}
					// 2. execute scheduled transaction
					// 2-1. timelock worker -> timelock::execute_batch { spl::mint } by op id

					recipient, kerr := solana.NewRandomPrivateKey()
					require.NoError(t, kerr)

					rIxATA, rAta, rAtaIxErr := tokens.CreateAssociatedTokenAccount(v.tokenProgram, mint, recipient.PublicKey(), admin.PublicKey())
					require.NoError(t, rAtaIxErr)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rIxATA}, admin, config.DefaultCommitment)

					_, rInitBal, bErr := tokens.TokenBalance(ctx, solanaGoClient, rAta, config.DefaultCommitment)
					require.NoError(t, bErr)
					require.Equal(t, 0, rInitBal)

					// mint authority to timelock
					authIx, aErr := tokens.SetTokenMintAuthority(v.tokenProgram, timelockutil.GetSignerPDA(config.TestTimelockID), mint, admin.PublicKey())
					require.NoError(t, aErr)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{authIx}, admin, config.DefaultCommitment)

					numMintIxs := 18

					salt, sErr := timelockutil.SimpleSalt()
					require.NoError(t, sErr)
					opToSchedule := timelockutil.Operation{
						TimelockID:  config.TestTimelockID,
						Predecessor: config.TimelockEmptyOpID, // no predecessor
						Salt:        salt,
						Delay:       uint64(1),
					}

					mcmOpNodes := []mcms.McmOpNode{}

					for i := 0; i < numMintIxs; i++ {
						// timelock signer can mint token (transferred authority)
						ix, mIxErr := tokens.MintTo(1*solana.LAMPORTS_PER_SOL, v.tokenProgram, mint, rAta, timelockutil.GetSignerPDA(config.TestTimelockID))
						require.NoError(t, mIxErr)

						// add instruction to timelock operation
						opToSchedule.AddInstruction(ix, []solana.PublicKey{v.tokenProgram})
					}

					ixs, ierr := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, opToSchedule, proposerMsig.SignerPDA, msigs[timelock.Proposer_Role].AccessController.PublicKey())
					require.NoError(t, ierr)
					for i, ix := range ixs {
						opNode, oerr := mcms.IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, ix, uint64(currentOpCount+i))
						require.NoError(t, oerr)
						mcmOpNodes = append(mcmOpNodes, opNode)
					}

					// Schedule the operation
					scheduleIx, scErr := timelock.NewScheduleBatchInstruction(
						config.TestTimelockID,
						opToSchedule.OperationID(),
						opToSchedule.Delay,
						opToSchedule.OperationPDA(),
						timelockutil.GetConfigPDA(config.TestTimelockID),
						msigs[timelock.Proposer_Role].AccessController.PublicKey(),
						proposerMsig.SignerPDA, // msig signer since we're going to run this ix with mcm::execute
					).ValidateAndBuild()
					require.NoError(t, scErr)

					node, cErr := mcms.IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, scheduleIx, uint64(currentOpCount+len(mcmOpNodes))) // operation nonce
					require.NoError(t, cErr)

					mcmOpNodes = append(mcmOpNodes, node)

					validUntil := uint32(0xffffffff)

					rootValidationData, rErr := mcms.CreateMcmRootData(mcms.McmRootInput{
						Multisig:             proposerMsig.ConfigPDA,
						Operations:           mcmOpNodes,
						PreOpCount:           uint64(currentOpCount),
						PostOpCount:          uint64(currentOpCount + len(mcmOpNodes)),
						ValidUntil:           validUntil,
						OverridePreviousRoot: false,
					})
					require.NoError(t, rErr)

					currentOpCount += len(mcmOpNodes)

					signatures, bulkSignErr := mcms.BulkSignOnMsgHash(proposerMsig.Signers, rootValidationData.EthMsgHash)
					require.NoError(t, bulkSignErr)
					signaturesPDA := proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil, admin.PublicKey())

					t.Run("mcm:preload signatures", func(t *testing.T) {
						preloadIxs, plerr := mcms.GetMcmPreloadSignaturesIxs(signatures, proposerMsig.PaddedID, rootValidationData.Root, validUntil, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
						require.NoError(t, plerr)

						for _, ix := range preloadIxs {
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
						}

						var sigAccount mcm.RootSignatures
						queryErr := common.GetAccountDataBorshInto(ctx, solanaGoClient, signaturesPDA, config.DefaultCommitment, &sigAccount)
						require.NoError(t, queryErr, "failed to get account info")

						require.Equal(t, true, sigAccount.IsFinalized)
						require.Equal(t, true, sigAccount.TotalSignatures == uint8(len(signatures)))

						// check if the sigs are registered correctly
						for i, sig := range sigAccount.Signatures {
							require.Equal(t, signatures[i], sig)
						}
					})

					t.Run("mcm:set_root", func(t *testing.T) {
						newIx, setRootIxErr := mcm.NewSetRootInstruction(
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
						require.NoError(t, setRootIxErr)

						tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{newIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
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

						var newRootAndOpCount mcm.ExpiringRootAndOpCount

						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &newRootAndOpCount)
						require.NoError(t, err, "failed to get account info")

						require.Equal(t, rootValidationData.Root, newRootAndOpCount.Root)
						require.Equal(t, validUntil, newRootAndOpCount.ValidUntil)
						require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootAndOpCount.OpCount)

						var newRootMetadata mcm.RootMetadata
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.RootMetadataPDA, config.DefaultCommitment, &newRootMetadata)
						require.NoError(t, err, "failed to get account info")

						require.Equal(t, rootValidationData.Metadata.ChainId, newRootMetadata.ChainId)
						require.Equal(t, rootValidationData.Metadata.Multisig, newRootMetadata.Multisig)
						require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootMetadata.PreOpCount)
						require.Equal(t, rootValidationData.Metadata.PostOpCount, newRootMetadata.PostOpCount)
						require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, newRootMetadata.OverridePreviousRoot)
					})

					t.Run("mcm:execute -> timelock::schedule_batch", func(t *testing.T) {
						// this includes timelock::initialize, append, finalize operation, and timelock::schedule_batch ixs
						t.Run("check if timelock config is correct", func(t *testing.T) {
							info, infoErr := solanaGoClient.GetAccountInfoWithOpts(ctx, timelockutil.GetConfigPDA(config.TestTimelockID), &rpc.GetAccountInfoOpts{
								Commitment: config.DefaultCommitment,
							})
							require.NoError(t, infoErr)
							require.Equal(t, info.Value.Owner, config.TimelockProgram, "Timelock config owner doesn't match")
						})

						for _, op := range mcmOpNodes {
							proofs, proofsErr := op.Proofs()
							require.NoError(t, proofsErr, "Failed to getting op proof")

							ix := mcm.NewExecuteInstruction(
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

							ix.AccountMetaSlice = append(ix.AccountMetaSlice, op.RemainingAccounts...)

							vIx, vIxErr := ix.ValidateAndBuild()
							require.NoError(t, vIxErr)

							tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, anyone, config.DefaultCommitment)
							require.NotNil(t, tx.Meta)
							require.Nil(t, tx.Meta.Err, fmt.Sprintf("tx failed with: %+v", tx.Meta))

							if bytes.Equal(op.Data[:8], timelock.Instruction_ScheduleBatch[:]) {
								parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
									[]common.EventMapping{
										common.EventMappingFor[mcms.OpExecuted]("OpExecuted"),
										common.EventMappingFor[timelockutil.CallScheduled]("CallScheduled"),
									},
								)

								// check opExecuted event
								event := parsedLogs[0].EventData[0].Data.(*mcms.OpExecuted)
								require.Equal(t, op.Nonce, event.Nonce)
								require.Equal(t, op.To, event.To)
								require.Equal(t, op.Data, common.NormalizeData(event.Data))

								// check inner CallScheduled events
								opIxData := opToSchedule.ToInstructionData()
								require.Equal(t, len(opIxData), len(parsedLogs[0].InnerCalls[0].EventData), "Number of actual CallScheduled events does not match expected for operation")

								for j, ix := range opIxData {
									timelockEvent := parsedLogs[0].InnerCalls[0].EventData[j].Data.(*timelockutil.CallScheduled)
									require.Equal(t, opToSchedule.OperationID(), timelockEvent.ID, "ID does not match")
									require.Equal(t, uint64(j), timelockEvent.Index, "Index does not match")
									require.Equal(t, ix.ProgramId, timelockEvent.Target, "Target does not match")
									require.Equal(t, opToSchedule.Predecessor, timelockEvent.Predecessor, "Predecessor does not match")
									require.Equal(t, opToSchedule.Salt, timelockEvent.Salt, "Salt does not match")
									require.Equal(t, opToSchedule.Delay, timelockEvent.Delay, "Delay does not match")
									require.Equal(t, ix.Data, common.NormalizeData(timelockEvent.Data), "Data does not match")
								}

								var opAccount timelock.Operation
								err = common.GetAccountDataBorshInto(ctx, solanaGoClient, opToSchedule.OperationPDA(), config.DefaultCommitment, &opAccount)
								if err != nil {
									require.NoError(t, err, "failed to get account info")
								}

								require.Equal(t,
									tx.BlockTime.Time().Add(time.Duration(opToSchedule.Delay)*time.Second).Unix(),
									int64(opAccount.Timestamp),
									"Scheduled Times don't match",
								)

								require.Equal(t,
									opToSchedule.OperationID(),
									opAccount.Id,
									"Ids don't match",
								)
							}
						}
					})

					t.Run("success: wait for operations to be ready", func(t *testing.T) {
						wErr := timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, opToSchedule.OperationPDA(), config.DefaultCommitment)
						require.NoError(t, wErr)
					})

					t.Run("timelock worker -> timelock::execute_batch", func(t *testing.T) {
						ix := timelock.NewExecuteBatchInstruction(
							config.TestTimelockID,
							opToSchedule.OperationID(),
							opToSchedule.OperationPDA(),
							config.TimelockEmptyOpID,
							timelockutil.GetConfigPDA(config.TestTimelockID),
							timelockutil.GetSignerPDA(config.TestTimelockID),
							msigs[timelock.Executor_Role].AccessController.PublicKey(),
							admin.PublicKey(), // timelock worker authority
						)

						ix.AccountMetaSlice = append(ix.AccountMetaSlice, opToSchedule.RemainingAccounts()...)

						vIx, vErr := ix.ValidateAndBuild()
						require.NoError(t, vErr)

						tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
						require.NotNil(t, tx)

						parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
							[]common.EventMapping{
								common.EventMappingFor[timelockutil.CallExecuted]("CallExecuted"),
							},
						)

						for i, ixx := range opToSchedule.ToInstructionData() {
							event := parsedLogs[0].EventData[i].Data.(*timelockutil.CallExecuted)
							require.Equal(t, opToSchedule.OperationID(), event.ID)
							require.Equal(t, uint64(i), event.Index)
							require.Equal(t, ixx.ProgramId, event.Target)
							require.Equal(t, ixx.Data, common.NormalizeData(event.Data))
						}

						var opAccount timelock.Operation
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, opToSchedule.OperationPDA(), config.DefaultCommitment, &opAccount)
						if err != nil {
							require.NoError(t, err, "failed to get account info")
						}

						require.Equal(t,
							timelock.Done_OperationState,
							opAccount.State,
							"Executed operation's time should be 1(DONE_TIMESTAMP)",
						)

						_, rInitBal, bErr := tokens.TokenBalance(ctx, solanaGoClient, rAta, config.DefaultCommitment)
						require.NoError(t, bErr)
						require.Equal(t, numMintIxs*int(solana.LAMPORTS_PER_SOL), rInitBal)
					})
				})
			})
		}

		var rootAccount mcm.ExpiringRootAndOpCount
		err = common.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &rootAccount)
		require.NoError(t, err, "failed to get account info")

		require.Equal(t, uint64(currentOpCount), rootAccount.OpCount)
	})

	t.Run("scheduled token vesting scenario", func(t *testing.T) {
		///////////////////////////////////////////
		// Setup - Create Token & Pass Authority //
		///////////////////////////////////////////
		mintKeypair, mkerr := solana.NewRandomPrivateKey()
		require.NoError(t, mkerr)
		mint := mintKeypair.PublicKey()

		tokenProgram := config.Token2022Program

		// Use CreateToken utility to get initialization instructions
		// NOTE: can't create token with cpi(mint signature required)
		createTokenIxs, cterr := tokens.CreateToken(
			ctx,
			tokenProgram,      // token program
			mint,              // mint account
			admin.PublicKey(), // initial mint owner(admin)
			9,                 // decimals
			solanaGoClient,
			config.DefaultCommitment,
		)
		require.NoError(t, cterr)

		authIx, aterr := tokens.SetTokenMintAuthority(tokenProgram, timelockutil.GetSignerPDA(config.TestTimelockID), mint, admin.PublicKey())
		require.NoError(t, aterr)

		setupIxs := append(createTokenIxs, authIx)

		testutils.SendAndConfirm(ctx, t, solanaGoClient, setupIxs, admin, config.DefaultCommitment, common.AddSigners(mintKeypair))

		/////////////////////////////////////////
		// Timelock Operation 1 - Initial Mint //
		/////////////////////////////////////////
		treasury, kerr := solana.NewRandomPrivateKey()
		require.NoError(t, kerr)

		ix1, treasuryATA, taerr := tokens.CreateAssociatedTokenAccount(tokenProgram, mint, treasury.PublicKey(), timelockutil.GetSignerPDA(config.TestTimelockID))
		require.NoError(t, taerr)

		ix2, tmerr := tokens.MintTo(1000*solana.LAMPORTS_PER_SOL, tokenProgram, mint, treasuryATA, timelockutil.GetSignerPDA(config.TestTimelockID))
		require.NoError(t, tmerr)

		salt1, serr := timelockutil.SimpleSalt()
		require.NoError(t, serr)
		op1 := timelockutil.Operation{
			TimelockID:  config.TestTimelockID,
			Predecessor: config.TimelockEmptyOpID, // no predecessor
			Salt:        salt1,
			Delay:       uint64(1),
		}
		op1.AddInstruction(ix1, []solana.PublicKey{tokenProgram, solana.SPLAssociatedTokenAccountProgramID})
		op1.AddInstruction(ix2, []solana.PublicKey{tokenProgram})

		////////////////////////////////////////////////////////////////////////////
		// Timelock Operation 2 - Schedule team associated token account creation //
		////////////////////////////////////////////////////////////////////////////
		team1, t1err := solana.NewRandomPrivateKey()
		require.NoError(t, t1err)
		team2, t2err := solana.NewRandomPrivateKey()
		require.NoError(t, t2err)
		team3, t3err := solana.NewRandomPrivateKey()
		require.NoError(t, t3err)

		ix3, team1ATA, t1cerr := tokens.CreateAssociatedTokenAccount(
			tokenProgram, mint, team1.PublicKey(),
			timelockutil.GetSignerPDA(config.TestTimelockID),
		)
		require.NoError(t, t1cerr)

		ix4, team2ATA, t2cerr := tokens.CreateAssociatedTokenAccount(
			tokenProgram, mint, team2.PublicKey(),
			timelockutil.GetSignerPDA(config.TestTimelockID),
		)
		require.NoError(t, t2cerr)

		ix5, team3ATA, t3cerr := tokens.CreateAssociatedTokenAccount(
			tokenProgram, mint, team3.PublicKey(),
			timelockutil.GetSignerPDA(config.TestTimelockID),
		)
		require.NoError(t, t3cerr)

		salt2, s2err := timelockutil.SimpleSalt()
		require.NoError(t, s2err)
		op2 := timelockutil.Operation{
			TimelockID:  config.TestTimelockID,
			Predecessor: op1.OperationID(), // must happen after initial mint
			Salt:        salt2,
			Delay:       uint64(1),
		}
		op2.AddInstruction(ix3, []solana.PublicKey{tokenProgram, solana.SPLAssociatedTokenAccountProgramID})
		op2.AddInstruction(ix4, []solana.PublicKey{tokenProgram, solana.SPLAssociatedTokenAccountProgramID})
		op2.AddInstruction(ix5, []solana.PublicKey{tokenProgram, solana.SPLAssociatedTokenAccountProgramID})

		//////////////////////////////////////////////////////////////
		// Timelock Operation 3 - Schedule team token distribution //
		//////////////////////////////////////////////////////////////
		ix6, i6err := tokens.TokenTransferChecked(100*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team1ATA, timelockutil.GetSignerPDA(config.TestTimelockID), []solana.PublicKey{})
		require.NoError(t, i6err)
		ix7, i7err := tokens.TokenTransferChecked(200*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team2ATA, timelockutil.GetSignerPDA(config.TestTimelockID), []solana.PublicKey{})
		require.NoError(t, i7err)
		ix8, i8err := tokens.TokenTransferChecked(300*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team3ATA, timelockutil.GetSignerPDA(config.TestTimelockID), []solana.PublicKey{})
		require.NoError(t, i8err)

		// add all team distribution instructions
		salt3, s3err := timelockutil.SimpleSalt()
		require.NoError(t, s3err)
		op3 := timelockutil.Operation{
			TimelockID:  config.TestTimelockID,
			Predecessor: op2.OperationID(), // must happen after ata creation
			Salt:        salt3,
			Delay:       uint64(1),
		}
		op3.AddInstruction(ix6, []solana.PublicKey{tokenProgram})
		op3.AddInstruction(ix7, []solana.PublicKey{tokenProgram})
		op3.AddInstruction(ix8, []solana.PublicKey{tokenProgram})

		require.NotEqual(t, op1.OperationID(), op2.OperationID(), "Operation IDs should be different")
		require.NotEqual(t, op1.OperationID(), op3.OperationID(), "Operation IDs should be different")
		require.NotEqual(t, op2.OperationID(), op3.OperationID(), "Operation IDs should be different")

		////////////////////////////////////////
		// Pre-create Timelock Operation PDAs //
		////////////////////////////////////////
		opNodes := []mcms.McmOpNode{}
		timelockOps := []timelockutil.Operation{op1, op2, op3}

		nextNonce := uint64(currentOpCount)

		for i, op := range timelockOps {
			t.Run(fmt.Sprintf("prepare mcm op node %d with timelock::schedule_batch ix", i), func(t *testing.T) {
				// base nonce for this operation group
				baseNonce := nextNonce + uint64(i)

				preloadIxs, ierr := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, op, proposerMsig.SignerPDA, msigs[timelock.Proposer_Role].AccessController.PublicKey())
				require.NoError(t, ierr)

				for j, ix := range preloadIxs {
					opNode, oerr := mcms.IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, ix, baseNonce+uint64(j))
					require.NoError(t, oerr)
					opNodes = append(opNodes, opNode)
				}

				scheduleOpIx, scErr := timelock.NewScheduleBatchInstruction(
					config.TestTimelockID,
					op.OperationID(),
					op.Delay,
					op.OperationPDA(),
					timelockutil.GetConfigPDA(config.TestTimelockID),
					msigs[timelock.Proposer_Role].AccessController.PublicKey(),
					proposerMsig.SignerPDA,
				).ValidateAndBuild()
				require.NoError(t, scErr)

				opNode, cErr := mcms.IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, scheduleOpIx, baseNonce+uint64(len(preloadIxs)))
				require.NoError(t, cErr)

				opNodes = append(opNodes, opNode)

				// update nextNonce to start after this operation group
				nextNonce += uint64(len(preloadIxs))
			})
		}

		//////////////////////////////////
		// mcm - Prepare root & root metadata //
		//////////////////////////////////
		validUntil := uint32(0xffffffff)

		rootValidationData, rverr := mcms.CreateMcmRootData(mcms.McmRootInput{
			Multisig:             proposerMsig.ConfigPDA,
			Operations:           opNodes,
			PreOpCount:           uint64(currentOpCount),
			PostOpCount:          uint64(currentOpCount + len(opNodes)),
			ValidUntil:           validUntil,
			OverridePreviousRoot: false,
		})
		require.NoError(t, rverr)

		// update currentOpCount
		currentOpCount += len(opNodes)

		t.Run("offchain: bulk sign on root and upload signatures", func(t *testing.T) {
			// sign the root
			signatures, signErr := mcms.BulkSignOnMsgHash(proposerMsig.Signers, rootValidationData.EthMsgHash)
			require.NoError(t, signErr)

			////////////////////////////////////////////////
			// mcm::set_root - with preloading signatures //
			////////////////////////////////////////////////
			preloadIxs, plerr := mcms.GetMcmPreloadSignaturesIxs(signatures, proposerMsig.PaddedID, rootValidationData.Root, validUntil, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
			require.NoError(t, plerr)

			for _, ix := range preloadIxs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			var sigAccount mcm.RootSignatures
			queryErr := common.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil, admin.PublicKey()), config.DefaultCommitment, &sigAccount)
			require.NoError(t, queryErr, "failed to get account info")

			require.Equal(t, true, sigAccount.IsFinalized)
			require.Equal(t, true, sigAccount.TotalSignatures == uint8(len(signatures)))

			// check if the sigs are registered correctly
			for i, sig := range sigAccount.Signatures {
				require.Equal(t, signatures[i], sig)
			}
		})

		t.Run("mcm::set_root", func(t *testing.T) {
			newIx, setRootIxErr := mcm.NewSetRootInstruction(
				proposerMsig.PaddedID,
				rootValidationData.Root,
				validUntil,
				rootValidationData.Metadata,
				rootValidationData.MetadataProof,
				proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil, admin.PublicKey()),
				proposerMsig.RootMetadataPDA,
				mcms.GetSeenSignedHashesPDA(proposerMsig.PaddedID, rootValidationData.Root, validUntil),
				proposerMsig.ExpiringRootAndOpCountPDA,
				proposerMsig.ConfigPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, setRootIxErr)

			tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{newIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
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

			var newRootAndOpCount mcm.ExpiringRootAndOpCount

			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &newRootAndOpCount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, rootValidationData.Root, newRootAndOpCount.Root)
			require.Equal(t, validUntil, newRootAndOpCount.ValidUntil)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootAndOpCount.OpCount)

			// get config and validate
			var newRootMetadata mcm.RootMetadata
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.RootMetadataPDA, config.DefaultCommitment, &newRootMetadata)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, rootValidationData.Metadata.ChainId, newRootMetadata.ChainId)
			require.Equal(t, rootValidationData.Metadata.Multisig, newRootMetadata.Multisig)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootMetadata.PreOpCount)
			require.Equal(t, rootValidationData.Metadata.PostOpCount, newRootMetadata.PostOpCount)
			require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, newRootMetadata.OverridePreviousRoot)
		})

		t.Run("mcm::execute to preload operation and schedule timelock operations", func(t *testing.T) {
			// this includes timelock::initialize, append, finalize operation, and timelock::schedule_batch ixs
			foundCallScheduledIdx := 0
			for i, op := range opNodes {
				proofs, proofsErr := op.Proofs()
				require.NoError(t, proofsErr)

				ix := mcm.NewExecuteInstruction(
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
				// append remaining accounts
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op.RemainingAccounts...)

				vIx, vIxErr := ix.ValidateAndBuild()
				require.NoError(t, vIxErr)

				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, anyone, config.DefaultCommitment)

				if bytes.Equal(op.Data[:8], timelock.Instruction_ScheduleBatch[:]) {
					parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
						[]common.EventMapping{
							common.EventMappingFor[mcms.OpExecuted]("OpExecuted"),
							common.EventMappingFor[timelockutil.CallScheduled]("CallScheduled"),
						},
					)

					// check opExecuted event
					event := parsedLogs[0].EventData[0].Data.(*mcms.OpExecuted)
					require.Equal(t, op.Nonce, event.Nonce)
					require.Equal(t, op.To, event.To)
					require.Equal(t, op.Data, common.NormalizeData(event.Data))

					// check inner CallScheduled events
					currentOp := timelockOps[foundCallScheduledIdx] // match the Operation with the current opNode

					opIxData := currentOp.ToInstructionData()

					require.Equal(t, len(opIxData), len(parsedLogs[0].InnerCalls[0].EventData), "Number of actual CallScheduled events does not match expected for operation %d", i)

					for j, ix := range opIxData {
						timelockEvent := parsedLogs[0].InnerCalls[0].EventData[j].Data.(*timelockutil.CallScheduled)
						require.Equal(t, currentOp.OperationID(), timelockEvent.ID, "ID does not match")
						require.Equal(t, uint64(j), timelockEvent.Index, "Index does not match")
						require.Equal(t, ix.ProgramId, timelockEvent.Target, "Target does not match")
						require.Equal(t, currentOp.Predecessor, timelockEvent.Predecessor, "Predecessor does not match")
						require.Equal(t, currentOp.Salt, timelockEvent.Salt, "Salt does not match")
						require.Equal(t, currentOp.Delay, timelockEvent.Delay, "Delay does not match")
						require.Equal(t, ix.Data, common.NormalizeData(timelockEvent.Data), "Data does not match")
					}
					foundCallScheduledIdx++
				}
			}
		})

		var newOp3 timelockutil.Operation

		t.Run("cancel and reschedule token distribution with corrected amounts", func(t *testing.T) {
			t.Run("cancel existing distribution operation through multisig", func(t *testing.T) {
				canceller := msigs[timelock.Canceller_Role].GetAnyMultisig()

				cancelIx, cerr := timelock.NewCancelInstruction(
					config.TestTimelockID,
					op3.OperationID(),
					op3.OperationPDA(),
					timelockutil.GetConfigPDA(config.TestTimelockID),
					msigs[timelock.Canceller_Role].AccessController.PublicKey(),
					canceller.SignerPDA,
				).ValidateAndBuild()
				require.NoError(t, cerr)

				// create MCM operation node for the cancel instruction
				// NOTE: nonce is 0 since it's the first operation
				node, nerr := mcms.IxToMcmTestOpNode(canceller.ConfigPDA, canceller.SignerPDA, cancelIx, uint64(0))
				require.NoError(t, nerr)

				cancleOpNodes := []mcms.McmOpNode{node}

				// create and validate root data for the cancel operation
				validUntil := uint32(0xffffffff)
				rootValidationData, rverr := mcms.CreateMcmRootData(mcms.McmRootInput{
					Multisig:             canceller.ConfigPDA,
					Operations:           cancleOpNodes,
					PreOpCount:           uint64(0),
					PostOpCount:          uint64(1),
					ValidUntil:           validUntil,
					OverridePreviousRoot: false,
				})
				require.NoError(t, rverr)

				signatures, serr := mcms.BulkSignOnMsgHash(canceller.Signers, rootValidationData.EthMsgHash)
				require.NoError(t, serr)

				signaturesPDA := canceller.RootSignaturesPDA(rootValidationData.Root, validUntil, admin.PublicKey())

				preloadIxs, plerr := mcms.GetMcmPreloadSignaturesIxs(signatures, canceller.PaddedID, rootValidationData.Root, validUntil, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
				require.NoError(t, plerr)

				for _, ix := range preloadIxs {
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				setRootIx, srerr := mcm.NewSetRootInstruction(
					canceller.PaddedID,
					rootValidationData.Root,
					validUntil,
					rootValidationData.Metadata,
					rootValidationData.MetadataProof,
					signaturesPDA,
					canceller.RootMetadataPDA,
					mcms.GetSeenSignedHashesPDA(canceller.PaddedID, rootValidationData.Root, validUntil),
					canceller.ExpiringRootAndOpCountPDA,
					canceller.ConfigPDA,
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
				require.Equal(t, canceller.ConfigPDA, event.MetadataMultisig)
				require.Equal(t, rootValidationData.Metadata.PreOpCount, event.MetadataPreOpCount)
				require.Equal(t, rootValidationData.Metadata.PostOpCount, event.MetadataPostOpCount)
				require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, event.MetadataOverridePreviousRoot)

				// execute mcm operation to cancel the timelock operation
				proofs, perr := cancleOpNodes[0].Proofs()
				require.NoError(t, perr)

				executeIx := mcm.NewExecuteInstruction(
					canceller.PaddedID,
					config.TestChainID,
					node.Nonce,
					node.Data,
					proofs,
					canceller.ConfigPDA,
					canceller.RootMetadataPDA,
					canceller.ExpiringRootAndOpCountPDA,
					node.To,
					canceller.SignerPDA,
					anyone.PublicKey(),
				)
				executeIx.AccountMetaSlice = append(executeIx.AccountMetaSlice, node.RemainingAccounts...)

				vIx, verr := executeIx.ValidateAndBuild()
				require.NoError(t, verr)

				exeTx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, anyone, config.DefaultCommitment)
				require.NotNil(t, exeTx)

				parsedExeLogs := common.ParseLogMessages(exeTx.Meta.LogMessages,
					[]common.EventMapping{
						common.EventMappingFor[mcms.OpExecuted]("OpExecuted"),
						common.EventMappingFor[timelockutil.Cancelled]("Cancelled"),
					},
				)

				// check opExecuted event
				exeEvent := parsedExeLogs[0].EventData[0].Data.(*mcms.OpExecuted)
				require.Equal(t, node.Nonce, exeEvent.Nonce)
				require.Equal(t, node.To, exeEvent.To)
				require.Equal(t, node.Data, common.NormalizeData(exeEvent.Data))

				// check inner Cancelled event
				timelockEvent := parsedExeLogs[0].InnerCalls[0].EventData[0].Data.(*timelockutil.Cancelled)
				require.Equal(t, op3.OperationID(), timelockEvent.ID, "ID does not match")

				// check if operation pda is closed
				testutils.AssertClosedAccount(ctx, t, solanaGoClient, op3.OperationPDA(), config.DefaultCommitment)
			})

			t.Run("create new operation with corrected amounts", func(t *testing.T) {
				// Create corrected transfer instructions with new amounts
				ix1, i1err := tokens.TokenTransferChecked(150*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team1ATA, timelockutil.GetSignerPDA(config.TestTimelockID), []solana.PublicKey{})
				require.NoError(t, i1err)
				ix2, i2err := tokens.TokenTransferChecked(150*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team2ATA, timelockutil.GetSignerPDA(config.TestTimelockID), []solana.PublicKey{})
				require.NoError(t, i2err)
				ix3, i3err := tokens.TokenTransferChecked(100*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team3ATA, timelockutil.GetSignerPDA(config.TestTimelockID), []solana.PublicKey{})
				require.NoError(t, i3err)

				// Create new operation
				salt, serr := timelockutil.SimpleSalt()
				require.NoError(t, serr)
				newOp3 = timelockutil.Operation{
					TimelockID:  config.TestTimelockID,
					Predecessor: op2.OperationID(),
					Salt:        salt,
					Delay:       uint64(1),
				}

				newOp3.AddInstruction(ix1, []solana.PublicKey{tokenProgram})
				newOp3.AddInstruction(ix2, []solana.PublicKey{tokenProgram})
				newOp3.AddInstruction(ix3, []solana.PublicKey{tokenProgram})

				newOpNodes := []mcms.McmOpNode{}

				ixs, perr := timelockutil.GetPreloadOperationIxs(config.TestTimelockID, newOp3, proposerMsig.SignerPDA, msigs[timelock.Proposer_Role].AccessController.PublicKey())
				require.NoError(t, perr)
				for i, ix := range ixs {
					opNode, oerr := mcms.IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, ix, uint64(currentOpCount+i))
					require.NoError(t, oerr)
					// append preloading instructions
					newOpNodes = append(newOpNodes, opNode)
				}

				// Create mcm operation node for scheduling
				scheduleIx, scerr := timelock.NewScheduleBatchInstruction(
					config.TestTimelockID,
					newOp3.OperationID(),
					newOp3.Delay,
					newOp3.OperationPDA(),
					timelockutil.GetConfigPDA(config.TestTimelockID),
					msigs[timelock.Proposer_Role].AccessController.PublicKey(),
					proposerMsig.SignerPDA,
				).ValidateAndBuild()
				require.NoError(t, scerr)

				opNode, onerr := mcms.IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, scheduleIx, uint64(currentOpCount+len(newOpNodes)))
				require.NoError(t, onerr)
				// append schedule instruction
				newOpNodes = append(newOpNodes, opNode)

				// Create and validate root data
				validUntil := uint32(0xffffffff)
				rootValidationData, rverr := mcms.CreateMcmRootData(mcms.McmRootInput{
					Multisig:             proposerMsig.ConfigPDA,
					Operations:           newOpNodes,
					PreOpCount:           uint64(currentOpCount),
					PostOpCount:          uint64(currentOpCount + len(newOpNodes)),
					ValidUntil:           validUntil,
					OverridePreviousRoot: false,
				})
				require.NoError(t, rverr)

				currentOpCount += len(newOpNodes)

				// Sign and set root
				signatures, serr := mcms.BulkSignOnMsgHash(proposerMsig.Signers, rootValidationData.EthMsgHash)
				require.NoError(t, serr)

				signaturesPDA := proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil, admin.PublicKey())

				// preload signatures
				preloadIxs, plerr := mcms.GetMcmPreloadSignaturesIxs(signatures, proposerMsig.PaddedID, rootValidationData.Root, validUntil, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
				require.NoError(t, plerr)
				for _, ix := range preloadIxs {
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				// Set root
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

				for _, op := range newOpNodes {
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
						opIxData := newOp3.ToInstructionData()

						require.Equal(t, len(opIxData), len(parsedExeLogs[0].InnerCalls[0].EventData), "Number of actual CallScheduled events does not match expected for operation")

						for j, ix := range opIxData {
							timelockEvent := parsedExeLogs[0].InnerCalls[0].EventData[j].Data.(*timelockutil.CallScheduled)
							require.Equal(t, newOp3.OperationID(), timelockEvent.ID, "ID does not match")
							require.Equal(t, uint64(j), timelockEvent.Index, "Index does not match")
							require.Equal(t, ix.ProgramId, timelockEvent.Target, "Target does not match")
							require.Equal(t, newOp3.Predecessor, timelockEvent.Predecessor, "Predecessor does not match")
							require.Equal(t, newOp3.Salt, timelockEvent.Salt, "Salt does not match")
							require.Equal(t, newOp3.Delay, timelockEvent.Delay, "Delay does not match")
							require.Equal(t, ix.Data, common.NormalizeData(timelockEvent.Data), "Data does not match")
						}
					}
				}
			})
		})

		t.Run("execute timelock operations", func(t *testing.T) {
			// Wait for operations to be ready
			werr := timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment)
			require.NoError(t, werr)

			rErr := timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, op2.OperationPDA(), config.DefaultCommitment)
			require.NoError(t, rErr)

			t.Run("op2: cannot be executed before op1", func(t *testing.T) {
				ix := timelock.NewExecuteBatchInstruction(
					config.TestTimelockID,
					op2.OperationID(),
					op2.OperationPDA(),
					op1.OperationPDA(), // provide op1 PDA as predecessor
					timelockutil.GetConfigPDA(config.TestTimelockID),
					timelockutil.GetSignerPDA(config.TestTimelockID),
					msigs[timelock.Executor_Role].AccessController.PublicKey(),
					admin.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

				vIx, verr := ix.ValidateAndBuild()
				require.NoError(t, verr)

				testutils.SendAndFailWith(ctx, t, solanaGoClient,
					[]solana.Instruction{vIx},
					admin,
					config.DefaultCommitment,
					[]string{"Error Code: " + timelock.MissingDependency_TimelockError.String()},
				)
			})

			t.Run("op1: initial mint to treasury", func(t *testing.T) {
				ix := timelock.NewExecuteBatchInstruction(
					config.TestTimelockID,
					op1.OperationID(),
					op1.OperationPDA(),
					config.TimelockEmptyOpID,
					timelockutil.GetConfigPDA(config.TestTimelockID),
					timelockutil.GetSignerPDA(config.TestTimelockID),
					msigs[timelock.Executor_Role].AccessController.PublicKey(),
					admin.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op1.RemainingAccounts()...)

				vIx, verr := ix.ValidateAndBuild()
				require.NoError(t, verr)

				cu := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient,
					[]solana.Instruction{vIx},
					admin,
					config.DefaultCommitment,
					common.AddComputeUnitLimit(cu),
				)
				require.NotNil(t, tx)
				parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
					[]common.EventMapping{
						common.EventMappingFor[timelockutil.CallExecuted]("CallExecuted"),
					},
				)
				for i, ix := range op1.ToInstructionData() {
					event := parsedLogs[0].EventData[i].Data.(*timelockutil.CallExecuted)
					require.Equal(t, op1.OperationID(), event.ID)
					require.Equal(t, uint64(i), event.Index)
					require.Equal(t, ix.ProgramId, event.Target)
					require.Equal(t, ix.Data, common.NormalizeData(event.Data))
				}

				// Verify operation status
				var opAccount timelock.Operation
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment, &opAccount)
				require.NoError(t, err)
				require.Equal(t, timelock.Done_OperationState, opAccount.State, "Op1 should be marked as done")

				// Verify treasury balance
				_, treasuryBalance, tberr := tokens.TokenBalance(ctx, solanaGoClient, treasuryATA, config.DefaultCommitment)
				require.NoError(t, tberr)
				require.Equal(t, 1000*int(solana.LAMPORTS_PER_SOL), treasuryBalance,
					"Treasury should have received 1000 tokens")
			})

			t.Run("token approval to timelock signer", func(t *testing.T) {
				// fund treasury account first
				fundIx, ferr := system.NewTransferInstruction(
					1*solana.LAMPORTS_PER_SOL, // 1 SOL should be enough
					admin.PublicKey(),
					treasury.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, ferr)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{fundIx}, admin, config.DefaultCommitment)

				// approve can't be deligated to timelock authority(security - CPI Guard)
				approveIx, aerr := tokens.TokenApproveChecked(
					600*solana.LAMPORTS_PER_SOL,
					9,
					tokenProgram,
					treasuryATA,
					mint,
					timelockutil.GetSignerPDA(config.TestTimelockID),
					treasury.PublicKey(),
					nil,
				)
				require.NoError(t, aerr)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{approveIx}, treasury, config.DefaultCommitment)
			})

			t.Run("op2: should provide the correct predecessor pda address", func(t *testing.T) {
				ix := timelock.NewExecuteBatchInstruction(
					config.TestTimelockID,
					op2.OperationID(),
					op2.OperationPDA(),
					op1.OperationID(), // provide op1 ID as predecessor
					timelockutil.GetConfigPDA(config.TestTimelockID),
					timelockutil.GetSignerPDA(config.TestTimelockID),
					msigs[timelock.Executor_Role].AccessController.PublicKey(),
					admin.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

				vIx, verr := ix.ValidateAndBuild()
				require.NoError(t, verr)

				testutils.SendAndFailWith(ctx, t, solanaGoClient,
					[]solana.Instruction{vIx},
					admin,
					config.DefaultCommitment,
					[]string{"Error Code: " + timelock.InvalidInput_TimelockError.String()},
				)
			})

			t.Run("op2: team ata creation", func(t *testing.T) {
				ix := timelock.NewExecuteBatchInstruction(
					config.TestTimelockID,
					op2.OperationID(),
					op2.OperationPDA(),
					op1.OperationPDA(), // provide op1 PDA as predecessor
					timelockutil.GetConfigPDA(config.TestTimelockID),
					timelockutil.GetSignerPDA(config.TestTimelockID),
					msigs[timelock.Executor_Role].AccessController.PublicKey(),
					admin.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

				vIx, verr := ix.ValidateAndBuild()
				require.NoError(t, verr)

				cu := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient,
					[]solana.Instruction{vIx},
					admin,
					config.DefaultCommitment,
					common.AddComputeUnitLimit(cu),
				)
				require.NotNil(t, tx)
				parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
					[]common.EventMapping{
						common.EventMappingFor[timelockutil.CallExecuted]("CallExecuted"),
					},
				)
				for i, ix := range op2.ToInstructionData() {
					event := parsedLogs[0].EventData[i].Data.(*timelockutil.CallExecuted)
					require.Equal(t, op2.OperationID(), event.ID)
					require.Equal(t, uint64(i), event.Index)
					require.Equal(t, ix.ProgramId, event.Target)
					require.Equal(t, ix.Data, common.NormalizeData(event.Data))
				}

				// verify operation status
				var opAccount timelock.Operation
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, op2.OperationPDA(), config.DefaultCommitment, &opAccount)
				require.NoError(t, err)
				require.Equal(t, timelock.Done_OperationState, opAccount.State, "Op2 should be marked as done")
			})
		})

		t.Run("op3: team token distribution", func(t *testing.T) {
			// Wait for delay and execute the timelock operation
			werr := timelockutil.WaitForOperationToBeReady(ctx, solanaGoClient, newOp3.OperationPDA(), config.DefaultCommitment)
			require.NoError(t, werr)

			executeTimelockIx := timelock.NewExecuteBatchInstruction(
				config.TestTimelockID,
				newOp3.OperationID(),
				newOp3.OperationPDA(),
				op2.OperationPDA(),
				timelockutil.GetConfigPDA(config.TestTimelockID),
				timelockutil.GetSignerPDA(config.TestTimelockID),
				msigs[timelock.Executor_Role].AccessController.PublicKey(),
				admin.PublicKey(),
			)
			executeTimelockIx.AccountMetaSlice = append(executeTimelockIx.AccountMetaSlice, newOp3.RemainingAccounts()...)

			vTimelockIx, verr := executeTimelockIx.ValidateAndBuild()
			require.NoError(t, verr)

			tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vTimelockIx}, admin, config.DefaultCommitment)
			require.NotNil(t, tx)
			parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
				[]common.EventMapping{
					common.EventMappingFor[timelockutil.CallExecuted]("CallExecuted"),
				},
			)
			for i, ix := range newOp3.ToInstructionData() {
				event := parsedLogs[0].EventData[i].Data.(*timelockutil.CallExecuted)
				require.NotEqual(t, op3.OperationID(), event.ID)
				require.Equal(t, newOp3.OperationID(), event.ID)
				require.Equal(t, uint64(i), event.Index)
				require.Equal(t, ix.ProgramId, event.Target)
				require.Equal(t, ix.Data, common.NormalizeData(event.Data))
			}
			// Verify final balances
			_, treasuryBalance, tberr := tokens.TokenBalance(ctx, solanaGoClient, treasuryATA, config.DefaultCommitment)
			require.NoError(t, tberr)
			require.Equal(t, 600*int(solana.LAMPORTS_PER_SOL), treasuryBalance,
				"Treasury should have 600 tokens remaining after distributions")

			_, team1Balance, t1berr := tokens.TokenBalance(ctx, solanaGoClient, team1ATA, config.DefaultCommitment)
			require.NoError(t, t1berr)
			require.Equal(t, 150*int(solana.LAMPORTS_PER_SOL), team1Balance,
				"Team1 should have received 150 tokens")

			_, team2Balance, t2berr := tokens.TokenBalance(ctx, solanaGoClient, team2ATA, config.DefaultCommitment)
			require.NoError(t, t2berr)
			require.Equal(t, 150*int(solana.LAMPORTS_PER_SOL), team2Balance,
				"Team2 should have received 150 tokens")

			_, team3Balance, t3berr := tokens.TokenBalance(ctx, solanaGoClient, team3ATA, config.DefaultCommitment)
			require.NoError(t, t3berr)
			require.Equal(t, 100*int(solana.LAMPORTS_PER_SOL), team3Balance,
				"Team3 should have received 100 tokens")
		})
	})
}
