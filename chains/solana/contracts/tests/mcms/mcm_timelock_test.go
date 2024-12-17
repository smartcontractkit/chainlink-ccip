package contracts

import (
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

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/access_controller"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils"
	mcmsUtils "github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/mcms"
)

func TestMcmWithTimelock(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	mcm.SetProgramID(config.McmProgram)
	timelock.SetProgramID(config.TimelockProgram)
	access_controller.SetProgramID(config.AccessControllerProgram)

	// initial admin
	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	anyone, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	solanaGoClient := utils.DeployAllPrograms(t, utils.PathToAnchorConfig, admin)

	msigs := map[timelock.Role]RoleMultisigs{
		timelock.Proposer_Role:  CreateRoleMultisigs(timelock.Proposer_Role, 1),
		timelock.Canceller_Role: CreateRoleMultisigs(timelock.Canceller_Role, 1),
		timelock.Executor_Role:  CreateRoleMultisigs(timelock.Executor_Role, 1),
		timelock.Bypasser_Role:  CreateRoleMultisigs(timelock.Bypasser_Role, 1),
	}

	require.NoError(t, err)

	t.Run("setup:funding", func(t *testing.T) {
		utils.FundAccounts(ctx, []solana.PrivateKey{admin, anyone}, solanaGoClient, t)

		// fund msig PDA signers
		for _, roleMsigs := range msigs {
			ixs := make([]solana.Instruction, 0)
			for _, msig := range roleMsigs.Multisigs {
				fundPDAIx := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), msig.SignerPDA).Build()
				ixs = append(ixs, fundPDAIx)
			}
			utils.SendAndConfirm(ctx, t, solanaGoClient,
				ixs,
				admin, config.DefaultCommitment)
		}
		// fund timelock signer
		fundPDAIx := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), config.TimelockSignerPDA).Build()
		utils.SendAndConfirm(ctx, t, solanaGoClient,
			[]solana.Instruction{fundPDAIx},
			admin, config.DefaultCommitment)
	})

	t.Run("setup: initialize mcm multisigs", func(t *testing.T) {
		for role, roleMsigs := range msigs {
			for _, msig := range roleMsigs.Multisigs {
				t.Run(fmt.Sprintf("init mcm for role %s with multisig %s", role.String(), mcmsUtils.UnpadString32(msig.PaddedName)), func(t *testing.T) {
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
						msig.PaddedName,
						msig.ConfigPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
						config.McmProgram,
						programData.Address,
						msig.RootMetadataPDA,
						msig.ExpiringRootAndOpCountPDA,
					).ValidateAndBuild()
					require.NoError(t, initIxErr)
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					// get config and validate
					var configAccount mcm.MultisigConfig
					err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigPDA, config.DefaultCommitment, &configAccount)
					require.NoError(t, err, "failed to get account info")

					require.Equal(t, config.TestChainID, configAccount.ChainId)
					require.Equal(t, admin.PublicKey(), configAccount.Owner)
					require.Equal(t, msig.PaddedName, configAccount.MultisigName)
				})
			}
		}
	})

	t.Run("setup: set_config for each mcm multisigs", func(t *testing.T) {
		for role, roleMsigs := range msigs {
			for _, msig := range roleMsigs.Multisigs {
				t.Run(fmt.Sprintf("set_config of role %s with multisig %s", role.String(), mcmsUtils.UnpadString32(msig.PaddedName)), func(t *testing.T) {
					signerAddresses := msig.RawConfig.SignerAddresses

					t.Run("preload signers on PDA", func(t *testing.T) {
						ixs := make([]solana.Instruction, 0)

						parsedTotalSigners, parseErr := mcmsUtils.SafeToUint8(len(signerAddresses))
						require.NoError(t, parseErr)

						initSignersIx, initSignersIxErr := mcm.NewInitSignersInstruction(
							msig.PaddedName,
							parsedTotalSigners,
							msig.ConfigPDA,
							msig.ConfigSignersPDA,
							admin.PublicKey(),
							solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, initSignersIxErr)
						ixs = append(ixs, initSignersIx)

						appendSignersIxs, appendSignersIxsErr := AppendSignersIxs(signerAddresses, msig.PaddedName, msig.ConfigPDA, msig.ConfigSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
						require.NoError(t, appendSignersIxsErr)
						ixs = append(ixs, appendSignersIxs...)

						finalizeSignersIx, finSignersIxErr := mcm.NewFinalizeSignersInstruction(
							msig.PaddedName,
							msig.ConfigPDA,
							msig.ConfigSignersPDA,
							admin.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, finSignersIxErr)
						ixs = append(ixs, finalizeSignersIx)

						for _, ix := range ixs {
							utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
						}

						var cfgSignersAccount mcm.ConfigSigners
						err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
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
							msig.PaddedName,
							msig.RawConfig.SignerGroups,
							msig.RawConfig.GroupQuorums,
							msig.RawConfig.GroupParents,
							msig.RawConfig.ClearRoot,
							msig.ConfigPDA,
							msig.ConfigSignersPDA,
							admin.PublicKey(),
							solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, setConfigErr)

						result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
						require.NotNil(t, result)

						// get config and validate
						var configAccount mcm.MultisigConfig
						err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigPDA, config.DefaultCommitment, &configAccount)
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
						utils.AssertClosedAccount(ctx, t, solanaGoClient, msig.ConfigSignersPDA, config.DefaultCommitment)
					})
				})
			}
		}
	})

	t.Run("setup: timelock", func(t *testing.T) {
		for role, roleMsigs := range msigs {
			t.Run(fmt.Sprintf("init access controller for role %s", role.String()), func(t *testing.T) {
				initAccIxs, initAccErr := InitAccessControllersIxs(ctx, roleMsigs.AccessController.PublicKey(), admin, solanaGoClient)
				require.NoError(t, initAccErr)

				utils.SendAndConfirm(ctx, t, solanaGoClient, initAccIxs, admin, config.DefaultCommitment, utils.AddSigners(roleMsigs.AccessController))

				var ac access_controller.AccessController
				err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, roleMsigs.AccessController.PublicKey(), config.DefaultCommitment, &ac)
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
				config.MinDelay,
				config.TimelockConfigPDA,
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

			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initTimelockIx}, admin, config.DefaultCommitment)

			var configAccount timelock.Config
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, config.TimelockConfigPDA, config.DefaultCommitment, &configAccount)
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
					batchAddAccessIxs, batchAddAccessErr := TimelockBatchAddAccessIxs(ctx, roleMsigs.AccessController.PublicKey(), role, addresses, admin, config.BatchAddAccessChunkSize, solanaGoClient)
					require.NoError(t, batchAddAccessErr)

					for _, ix := range batchAddAccessIxs {
						utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
					}

					var ac access_controller.AccessController
					err = utils.GetAccountDataBorshInto(
						ctx,
						solanaGoClient,
						roleMsigs.AccessController.PublicKey(),
						config.DefaultCommitment,
						&ac,
					)
					require.NoError(t, err)

					require.Equal(t, uint64(len(roleMsigs.Multisigs)), ac.AccessList.Len,
						"AccessList length mismatch for %s", role)

					for _, msig := range roleMsigs.Multisigs {
						targetPubKey := msig.SignerPDA
						_, found := mcmsUtils.FindInSortedList(ac.AccessList.Xs[:ac.AccessList.Len], targetPubKey)
						require.True(t, found, "Account %s not found in %s AccessList",
							targetPubKey, role)
					}
				})
			}
		})
	})

	t.Run("setup: transfer ownership multisigs to timelock signer", func(t *testing.T) {
		for role, roleMsigs := range msigs {
			for _, msig := range roleMsigs.Multisigs {
				t.Run(fmt.Sprintf("transfer ownership of role %s multisig %s to timelock signer", role.String(), mcmsUtils.UnpadString32(msig.PaddedName)), func(t *testing.T) {
					t.Parallel()
					ix, transferOwnershipErr := mcm.NewTransferOwnershipInstruction(
						msig.PaddedName,
						config.TimelockSignerPDA, // new proposed owner
						msig.ConfigPDA,
						admin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, transferOwnershipErr)

					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					var configAccount mcm.MultisigConfig
					err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigPDA, config.DefaultCommitment, &configAccount)
					if err != nil {
						require.NoError(t, err, "failed to get account info")
					}
					require.Equal(t, admin.PublicKey(), configAccount.Owner)
					require.Equal(t, config.TimelockSignerPDA, configAccount.ProposedOwner)

					acceptOwnershipIx, acceptOwnershipixErr := mcm.NewAcceptOwnershipInstruction(
						msig.PaddedName,
						msig.ConfigPDA,
						config.TimelockSignerPDA,
					).ValidateAndBuild()
					require.NoError(t, acceptOwnershipixErr)

					salt, sErr := mcmsUtils.SimpleSalt()
					require.NoError(t, sErr)
					acceptOwnershipOp := TimelockOperation{
						Predecessor: config.TimelockEmptyOpID,
						Salt:        salt,
						Delay:       uint64(1),
					}

					acceptOwnershipOp.AddInstruction(acceptOwnershipIx, []solana.PublicKey{config.McmProgram})

					id := acceptOwnershipOp.OperationID()
					operationPDA := acceptOwnershipOp.OperationPDA()

					ixs := make([]solana.Instruction, 0)
					initOpIx, initOpErr := timelock.NewInitializeOperationInstruction(
						acceptOwnershipOp.OperationID(),
						acceptOwnershipOp.Predecessor,
						acceptOwnershipOp.Salt,
						acceptOwnershipOp.IxsCountU32(),
						config.TimelockConfigPDA,
						operationPDA,
						admin.PublicKey(),
						admin.PublicKey(), // proposer - direct schedule batch here
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, initOpErr)
					ixs = append(ixs, initOpIx)

					appendIxIx, apErr := timelock.NewAppendInstructionsInstruction(
						acceptOwnershipOp.OperationID(),
						acceptOwnershipOp.ToInstructionData(),
						operationPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, apErr)

					ixs = append(ixs, appendIxIx)

					finIxIx, finErr := timelock.NewFinalizeOperationInstruction(
						acceptOwnershipOp.OperationID(),
						operationPDA,
						admin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, finErr)
					ixs = append(ixs, finIxIx)

					utils.SendAndConfirm(ctx, t, solanaGoClient, ixs, admin, config.DefaultCommitment)

					scheduleBatchIx, scErr := timelock.NewScheduleBatchInstruction(
						acceptOwnershipOp.OperationID(),
						acceptOwnershipOp.Delay,
						config.TimelockConfigPDA,
						operationPDA,
						roleMsigs.AccessController.PublicKey(),
						admin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, scErr)

					result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{scheduleBatchIx}, admin, config.DefaultCommitment)

					var opAccount timelock.Operation
					err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, operationPDA, config.DefaultCommitment, &opAccount)
					if err != nil {
						require.NoError(t, err, "failed to get account info")
					}

					require.Equal(t,
						result.BlockTime.Time().Add(time.Duration(acceptOwnershipOp.Delay)*time.Second).Unix(),
						int64(opAccount.Timestamp),
						"Scheduled Times don't match",
					)

					require.Equal(t,
						id,
						opAccount.Id,
						"Ids don't match",
					)

					bypassExeIx := timelock.NewBypasserExecuteBatchInstruction(
						acceptOwnershipOp.OperationID(),
						config.TimelockConfigPDA,
						config.TimelockSignerPDA,
						acceptOwnershipOp.OperationPDA(),
						roleMsigs.AccessController.PublicKey(),
						admin.PublicKey(), // bypass execute with admin previledges
					)
					bypassExeIx.AccountMetaSlice = append(bypassExeIx.AccountMetaSlice, acceptOwnershipOp.RemainingAccounts()...)

					vIx, vIxErr := bypassExeIx.ValidateAndBuild()
					require.NoError(t, vIxErr)

					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment)

					err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, msig.ConfigPDA, config.DefaultCommitment, &configAccount)
					if err != nil {
						require.NoError(t, err, "failed to get account info")
					}
					require.Equal(t, config.TimelockSignerPDA, configAccount.Owner)
					require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
				})
			}
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
				createTokenIxs, createTokenErr := utils.CreateToken(
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
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, utils.AddSigners(mintKeypair))
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

					rIxATA, rAta, rAtaIxErr := utils.CreateAssociatedTokenAccount(v.tokenProgram, mint, recipient.PublicKey(), admin.PublicKey())
					require.NoError(t, rAtaIxErr)

					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rIxATA}, admin, config.DefaultCommitment)

					_, rInitBal, bErr := utils.TokenBalance(ctx, solanaGoClient, rAta, config.DefaultCommitment)
					require.NoError(t, bErr)
					require.Equal(t, 0, rInitBal)

					// mint authority to timelock
					authIx, aErr := utils.SetTokenMintAuthority(v.tokenProgram, config.TimelockSignerPDA, mint, admin.PublicKey())
					require.NoError(t, aErr)

					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{authIx}, admin, config.DefaultCommitment)

					numMintIxs := 18

					salt, sErr := mcmsUtils.SimpleSalt()
					require.NoError(t, sErr)
					opToSchedule := TimelockOperation{
						Predecessor: config.TimelockEmptyOpID, // no predecessor
						Salt:        salt,
						Delay:       uint64(1),
					}

					for i := 0; i < numMintIxs; i++ {
						// timelock signer can mint token (transferred authority)
						ix, mIxErr := utils.MintTo(1*solana.LAMPORTS_PER_SOL, v.tokenProgram, mint, rAta, config.TimelockSignerPDA)
						require.NoError(t, mIxErr)

						// add instruction to timelock operation
						opToSchedule.AddInstruction(ix, []solana.PublicKey{v.tokenProgram})
					}

					initOpIx, ioErr := timelock.NewInitializeOperationInstruction(
						opToSchedule.OperationID(),
						opToSchedule.Predecessor,
						opToSchedule.Salt,
						uint32(len(opToSchedule.instructions)),
						config.TimelockConfigPDA,
						opToSchedule.OperationPDA(),
						admin.PublicKey(),
						proposerMsig.SignerPDA,
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, ioErr)

					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initOpIx}, admin, config.DefaultCommitment)

					for _, instruction := range opToSchedule.ToInstructionData() {
						appendIxsIx, apErr := timelock.NewAppendInstructionsInstruction(
							opToSchedule.OperationID(),
							[]timelock.InstructionData{instruction}, // this should be a slice of instruction within 1232 bytes
							opToSchedule.OperationPDA(),
							admin.PublicKey(),
							solana.SystemProgramID, // for reallocation
						).ValidateAndBuild()
						require.NoError(t, apErr)
						utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{appendIxsIx}, admin, config.DefaultCommitment)
					}

					finOpIx, foErr := timelock.NewFinalizeOperationInstruction(
						opToSchedule.OperationID(),
						opToSchedule.OperationPDA(),
						admin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, foErr)
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{finOpIx}, admin, config.DefaultCommitment)

					// Schedule the operation
					scheduleIx, scErr := timelock.NewScheduleBatchInstruction(
						opToSchedule.OperationID(),
						opToSchedule.Delay,
						config.TimelockConfigPDA,
						opToSchedule.OperationPDA(),
						msigs[timelock.Proposer_Role].AccessController.PublicKey(),
						proposerMsig.SignerPDA, // msig signer since we're going to run this ix with mcm::execute
					).ValidateAndBuild()
					require.NoError(t, scErr)

					node, cErr := IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, scheduleIx, uint64(currentOpCount)) // operation nonce
					require.NoError(t, cErr)
					mcmOpNodes := []mcmsUtils.McmOpNode{node} // only one mcm op node
					validUntil := uint32(0xffffffff)

					rootValidationData, rErr := CreateMcmRootData(McmRootInput{
						Multisig:             proposerMsig.ConfigPDA,
						Operations:           mcmOpNodes,
						PreOpCount:           uint64(currentOpCount),
						PostOpCount:          uint64(currentOpCount + len(mcmOpNodes)),
						ValidUntil:           validUntil,
						OverridePreviousRoot: false,
					})
					require.NoError(t, rErr)

					currentOpCount += len(mcmOpNodes)

					signatures, bulkSignErr := BulkSignOnMsgHash(proposerMsig.Signers, rootValidationData.EthMsgHash)
					require.NoError(t, bulkSignErr)
					signaturesPDA := proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil)

					t.Run("mcm:preload signatures", func(t *testing.T) {
						parsedTotalSigs, pErr := mcmsUtils.SafeToUint8(len(signatures))
						require.NoError(t, pErr)

						ixs := make([]solana.Instruction, 0)

						initSigsIx, isErr := mcm.NewInitSignaturesInstruction(
							proposerMsig.PaddedName,
							rootValidationData.Root,
							validUntil,
							parsedTotalSigs,
							signaturesPDA,
							admin.PublicKey(), // auth from someone who call set_root
							solana.SystemProgramID,
						).ValidateAndBuild()

						require.NoError(t, isErr)
						ixs = append(ixs, initSigsIx)

						appendSigsIxs, asErr := AppendSignaturesIxs(signatures, proposerMsig.PaddedName, rootValidationData.Root, validUntil, signaturesPDA, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
						require.NoError(t, asErr)
						ixs = append(ixs, appendSigsIxs...)

						finalizeSigsIx, fsErr := mcm.NewFinalizeSignaturesInstruction(
							proposerMsig.PaddedName,
							rootValidationData.Root,
							validUntil,
							signaturesPDA,
							admin.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, fsErr)
						ixs = append(ixs, finalizeSigsIx)

						for _, ix := range ixs {
							utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
						}

						var sigAccount mcm.RootSignatures
						queryErr := utils.GetAccountDataBorshInto(ctx, solanaGoClient, signaturesPDA, config.DefaultCommitment, &sigAccount)
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
							proposerMsig.PaddedName,
							rootValidationData.Root,
							validUntil,
							rootValidationData.Metadata,
							rootValidationData.MetadataProof,
							signaturesPDA,
							proposerMsig.RootMetadataPDA,
							SeenSignedHashesAddress(proposerMsig.PaddedName, rootValidationData.Root, validUntil),
							proposerMsig.ExpiringRootAndOpCountPDA,
							proposerMsig.ConfigPDA,
							admin.PublicKey(),
							solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, setRootIxErr)

						cuIx, cuErr := computebudget.NewSetComputeUnitLimitInstruction(1_400_000).ValidateAndBuild()
						require.NoError(t, cuErr)
						result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{cuIx, newIx}, admin, config.DefaultCommitment)
						require.NotNil(t, result)

						var newRootAndOpCount mcm.ExpiringRootAndOpCount

						err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &newRootAndOpCount)
						require.NoError(t, err, "failed to get account info")

						require.Equal(t, rootValidationData.Root, newRootAndOpCount.Root)
						require.Equal(t, validUntil, newRootAndOpCount.ValidUntil)
						require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootAndOpCount.OpCount)

						// get config and validate
						var newRootMetadata mcm.RootMetadata
						err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.RootMetadataPDA, config.DefaultCommitment, &newRootMetadata)
						require.NoError(t, err, "failed to get account info")

						require.Equal(t, rootValidationData.Metadata.ChainId, newRootMetadata.ChainId)
						require.Equal(t, rootValidationData.Metadata.Multisig, newRootMetadata.Multisig)
						require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootMetadata.PreOpCount)
						require.Equal(t, rootValidationData.Metadata.PostOpCount, newRootMetadata.PostOpCount)
						require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, newRootMetadata.OverridePreviousRoot)
					})

					t.Run("mcm:execute -> timelock::schedule_batch", func(t *testing.T) {
						t.Run("check if timelock config is correct", func(t *testing.T) {
							info, infoErr := solanaGoClient.GetAccountInfoWithOpts(ctx, config.TimelockConfigPDA, &rpc.GetAccountInfoOpts{
								Commitment: config.DefaultCommitment,
							})
							require.NoError(t, infoErr)
							require.Equal(t, info.Value.Owner, config.TimelockProgram, "Timelock config owner doesn't match")
						})

						for _, op := range mcmOpNodes {
							proofs, proofsErr := op.Proofs()
							require.NoError(t, proofsErr, "Failed to getting op proof")

							ix := mcm.NewExecuteInstruction(
								proposerMsig.PaddedName,
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

							tx := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, anyone, config.DefaultCommitment)
							require.NotNil(t, tx.Meta)
							require.Nil(t, tx.Meta.Err, fmt.Sprintf("tx failed with: %+v", tx.Meta))

							var opAccount timelock.Operation
							err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, opToSchedule.OperationPDA(), config.DefaultCommitment, &opAccount)
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
					})

					t.Run("success: wait for operations to be ready", func(t *testing.T) {
						wErr := WaitForOperationToBeReady(ctx, solanaGoClient, opToSchedule.OperationPDA(), config.DefaultCommitment)
						require.NoError(t, wErr)
					})

					t.Run("timelock worker -> timelock::execute_batch", func(t *testing.T) {
						ix := timelock.NewExecuteBatchInstruction(
							opToSchedule.OperationID(),
							config.TimelockConfigPDA,
							config.TimelockSignerPDA,
							opToSchedule.OperationPDA(),
							config.TimelockEmptyOpID,
							msigs[timelock.Executor_Role].AccessController.PublicKey(),
							admin.PublicKey(), // timelock worker authority
						)

						ix.AccountMetaSlice = append(ix.AccountMetaSlice, opToSchedule.RemainingAccounts()...)

						vIx, vErr := ix.ValidateAndBuild()
						require.NoError(t, vErr)

						result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment, utils.AddComputeUnitLimit(1_400_000))
						require.NotNil(t, result)

						var opAccount timelock.Operation
						err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, opToSchedule.OperationPDA(), config.DefaultCommitment, &opAccount)
						if err != nil {
							require.NoError(t, err, "failed to get account info")
						}

						require.Equal(t,
							config.TimelockOpDoneTimestamp,
							opAccount.Timestamp,
							"Executed operation's time should be 1(DONE_TIMESTAMP)",
						)

						_, rInitBal, bErr := utils.TokenBalance(ctx, solanaGoClient, rAta, config.DefaultCommitment)
						require.NoError(t, bErr)
						require.Equal(t, numMintIxs*int(solana.LAMPORTS_PER_SOL), rInitBal)
					})
				})
			})
		}

		var rootAccount mcm.ExpiringRootAndOpCount
		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &rootAccount)
		require.NoError(t, err, "failed to get account info")

		require.Equal(t, uint64(currentOpCount), rootAccount.OpCount)
	})

	t.Run("scheduled token vesting scenario", func(t *testing.T) {
		///////////////////////////////////////////
		// Setup - Create Token & Pass Authority //
		///////////////////////////////////////////
		mintKeypair, err := solana.NewRandomPrivateKey()
		require.NoError(t, err)
		mint := mintKeypair.PublicKey()

		tokenProgram := config.Token2022Program

		// Use CreateToken utility to get initialization instructions
		// NOTE: can't create token with cpi(mint signature required)
		createTokenIxs, err := utils.CreateToken(
			ctx,
			tokenProgram,      // token program
			mint,              // mint account
			admin.PublicKey(), // initial mint owner(admin)
			9,                 // decimals
			solanaGoClient,
			config.DefaultCommitment,
		)
		require.NoError(t, err)

		authIx, err := utils.SetTokenMintAuthority(tokenProgram, config.TimelockSignerPDA, mint, admin.PublicKey())
		require.NoError(t, err)

		setupIxs := append(createTokenIxs, authIx)

		utils.SendAndConfirm(ctx, t, solanaGoClient, setupIxs, admin, config.DefaultCommitment, utils.AddSigners(mintKeypair))

		/////////////////////////////////////////
		// Timelock Operation 1 - Initial Mint //
		/////////////////////////////////////////
		treasury, kerr := solana.NewRandomPrivateKey()
		require.NoError(t, kerr)

		ix1, treasuryATA, err := utils.CreateAssociatedTokenAccount(tokenProgram, mint, treasury.PublicKey(), config.TimelockSignerPDA)
		require.NoError(t, err)

		ix2, err := utils.MintTo(1000*solana.LAMPORTS_PER_SOL, tokenProgram, mint, treasuryATA, config.TimelockSignerPDA)
		require.NoError(t, err)

		salt1, err := mcmsUtils.SimpleSalt()
		require.NoError(t, err)
		op1 := TimelockOperation{
			Predecessor: config.TimelockEmptyOpID, // no predecessor
			Salt:        salt1,
			Delay:       uint64(1),
		}
		op1.AddInstruction(ix1, []solana.PublicKey{tokenProgram, solana.SPLAssociatedTokenAccountProgramID})
		op1.AddInstruction(ix2, []solana.PublicKey{tokenProgram})

		////////////////////////////////////////////////////////////////////////////
		// Timelock Operation 2 - Schedule team associated token account creation //
		////////////////////////////////////////////////////////////////////////////
		team1, err := solana.NewRandomPrivateKey()
		require.NoError(t, err)
		team2, err := solana.NewRandomPrivateKey()
		require.NoError(t, err)
		team3, err := solana.NewRandomPrivateKey()
		require.NoError(t, err)

		ix3, team1ATA, err := utils.CreateAssociatedTokenAccount(
			tokenProgram, mint, team1.PublicKey(),
			config.TimelockSignerPDA,
		)
		require.NoError(t, err)

		ix4, team2ATA, err := utils.CreateAssociatedTokenAccount(
			tokenProgram, mint, team2.PublicKey(),
			config.TimelockSignerPDA,
		)
		require.NoError(t, err)

		ix5, team3ATA, err := utils.CreateAssociatedTokenAccount(
			tokenProgram, mint, team3.PublicKey(),
			config.TimelockSignerPDA,
		)
		require.NoError(t, err)

		salt2, err := mcmsUtils.SimpleSalt()
		require.NoError(t, err)
		op2 := TimelockOperation{
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
		ix6, err := utils.TokenTransferChecked(100*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team1ATA, config.TimelockSignerPDA, []solana.PublicKey{})
		require.NoError(t, err)
		ix7, err := utils.TokenTransferChecked(200*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team2ATA, config.TimelockSignerPDA, []solana.PublicKey{})
		require.NoError(t, err)
		ix8, err := utils.TokenTransferChecked(300*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team3ATA, config.TimelockSignerPDA, []solana.PublicKey{})
		require.NoError(t, err)

		// add all team distribution instructions
		salt3, err := mcmsUtils.SimpleSalt()
		require.NoError(t, err)
		op3 := TimelockOperation{
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
		opNodes := []mcmsUtils.McmOpNode{}

		for i, op := range []TimelockOperation{op1, op2, op3} {
			t.Run(fmt.Sprintf("prepare mcm op node %d with timelock::schedule_batch ix", i), func(t *testing.T) {
				prepareOpIxs := make([]solana.Instruction, 0)

				initOpIx, initOpIxErr := timelock.NewInitializeOperationInstruction(
					op.OperationID(),
					op.Predecessor,
					op.Salt,
					op.IxsCountU32(),
					config.TimelockConfigPDA,
					op.OperationPDA(),
					admin.PublicKey(),
					proposerMsig.SignerPDA,
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, initOpIxErr)
				prepareOpIxs = append(prepareOpIxs, initOpIx)

				for _, ixData := range op.ToInstructionData() {
					appendIxsIx, apErr := timelock.NewAppendInstructionsInstruction(
						op.OperationID(),
						[]timelock.InstructionData{ixData}, // this should be a slice of instruction within 1232 bytes
						op.OperationPDA(),
						admin.PublicKey(),
						solana.SystemProgramID, // for reallocation
					).ValidateAndBuild()
					require.NoError(t, apErr)
					prepareOpIxs = append(prepareOpIxs, appendIxsIx)
				}

				finOpIx, finOpErr := timelock.NewFinalizeOperationInstruction(
					op.OperationID(),
					op.OperationPDA(),
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, finOpErr)
				prepareOpIxs = append(prepareOpIxs, finOpIx)

				for _, ix := range prepareOpIxs {
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				///////////////////////////////////////////
				// Construct schedule_batch instruction //
				///////////////////////////////////////////
				scheduleOpIx, scErr := timelock.NewScheduleBatchInstruction(
					op.OperationID(),
					op.Delay,
					config.TimelockConfigPDA,
					op.OperationPDA(),
					msigs[timelock.Proposer_Role].AccessController.PublicKey(),
					proposerMsig.SignerPDA,
				).ValidateAndBuild()
				require.NoError(t, scErr)

				opNode, cErr := IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, scheduleOpIx, uint64(currentOpCount+i))
				require.NoError(t, cErr)
				// fmt.Println("opNode", opNode)
				opNodes = append(opNodes, opNode)
			})
		}

		//////////////////////////////////
		// mcm - Prepare root & root metadata //
		//////////////////////////////////
		validUntil := uint32(0xffffffff)

		rootValidationData, err := CreateMcmRootData(McmRootInput{
			Multisig:             proposerMsig.ConfigPDA,
			Operations:           opNodes,
			PreOpCount:           uint64(currentOpCount),
			PostOpCount:          uint64(currentOpCount + len(opNodes)),
			ValidUntil:           validUntil,
			OverridePreviousRoot: false,
		})
		require.NoError(t, err)

		// update currentOpCount
		currentOpCount += len(opNodes)

		t.Run("offchain: bulk sign on root and upload signatures", func(t *testing.T) {
			// sign the root
			signatures, signErr := BulkSignOnMsgHash(proposerMsig.Signers, rootValidationData.EthMsgHash)
			require.NoError(t, signErr)

			////////////////////////////////////////////////
			// mcm::set_root - with preloading signatures //
			////////////////////////////////////////////////
			parsedTotalSigs, pErr := mcmsUtils.SafeToUint8(len(signatures))
			require.NoError(t, pErr)

			ixs := make([]solana.Instruction, 0)

			initSigsIx, isErr := mcm.NewInitSignaturesInstruction(
				proposerMsig.PaddedName,
				rootValidationData.Root,
				validUntil,
				parsedTotalSigs,
				proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil),
				admin.PublicKey(), // auth from someone who call set_root
				solana.SystemProgramID,
			).ValidateAndBuild()

			require.NoError(t, isErr)
			ixs = append(ixs, initSigsIx)

			appendSigsIxs, asErr := AppendSignaturesIxs(signatures, proposerMsig.PaddedName, rootValidationData.Root, validUntil, proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil), admin.PublicKey(), config.MaxAppendSignatureBatchSize)
			require.NoError(t, asErr)
			ixs = append(ixs, appendSigsIxs...)

			finalizeSigsIx, fsErr := mcm.NewFinalizeSignaturesInstruction(
				proposerMsig.PaddedName,
				rootValidationData.Root,
				validUntil,
				proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil),
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, fsErr)
			ixs = append(ixs, finalizeSigsIx)

			for _, ix := range ixs {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			var sigAccount mcm.RootSignatures
			queryErr := utils.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil), config.DefaultCommitment, &sigAccount)
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
				proposerMsig.PaddedName,
				rootValidationData.Root,
				validUntil,
				rootValidationData.Metadata,
				rootValidationData.MetadataProof,
				proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil),
				proposerMsig.RootMetadataPDA,
				SeenSignedHashesAddress(proposerMsig.PaddedName, rootValidationData.Root, validUntil),
				proposerMsig.ExpiringRootAndOpCountPDA,
				proposerMsig.ConfigPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, setRootIxErr)

			cuIx, cuErr := computebudget.NewSetComputeUnitLimitInstruction(1_400_000).ValidateAndBuild()
			require.NoError(t, cuErr)
			result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{cuIx, newIx}, admin, config.DefaultCommitment)
			require.NotNil(t, result)

			var newRootAndOpCount mcm.ExpiringRootAndOpCount

			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.ExpiringRootAndOpCountPDA, config.DefaultCommitment, &newRootAndOpCount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, rootValidationData.Root, newRootAndOpCount.Root)
			require.Equal(t, validUntil, newRootAndOpCount.ValidUntil)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootAndOpCount.OpCount)

			// get config and validate
			var newRootMetadata mcm.RootMetadata
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, proposerMsig.RootMetadataPDA, config.DefaultCommitment, &newRootMetadata)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, rootValidationData.Metadata.ChainId, newRootMetadata.ChainId)
			require.Equal(t, rootValidationData.Metadata.Multisig, newRootMetadata.Multisig)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootMetadata.PreOpCount)
			require.Equal(t, rootValidationData.Metadata.PostOpCount, newRootMetadata.PostOpCount)
			require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, newRootMetadata.OverridePreviousRoot)
		})

		t.Run("mcm::execute to schedule timelock operations", func(t *testing.T) {
			for _, op := range opNodes {
				proofs, proofsErr := op.Proofs()
				require.NoError(t, proofsErr)

				ix := mcm.NewExecuteInstruction(
					proposerMsig.PaddedName,
					config.TestChainID,
					op.Nonce,
					op.Data, // this is timelock::schedule_batch ix
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

				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, anyone, config.DefaultCommitment)
			}
		})

		var newOp3 TimelockOperation

		t.Run("cancel and reschedule token distribution with corrected amounts", func(t *testing.T) {
			t.Run("cancel existing distribution operation through multisig", func(t *testing.T) {
				canceller := msigs[timelock.Canceller_Role].GetAnyMultisig()

				cancelIx, err := timelock.NewCancelInstruction(
					op3.OperationID(),
					config.TimelockConfigPDA,
					op3.OperationPDA(),
					msigs[timelock.Canceller_Role].AccessController.PublicKey(),
					canceller.SignerPDA,
				).ValidateAndBuild()
				require.NoError(t, err)

				// create MCM operation node for the cancel instruction
				// NOTE: nonce is 0 since it's the first operation
				node, err := IxToMcmTestOpNode(canceller.ConfigPDA, canceller.SignerPDA, cancelIx, uint64(0))
				require.NoError(t, err)

				cancleOpNodes := []mcmsUtils.McmOpNode{node}

				// create and validate root data for the cancel operation
				validUntil := uint32(0xffffffff)
				rootValidationData, err := CreateMcmRootData(McmRootInput{
					Multisig:             canceller.ConfigPDA,
					Operations:           cancleOpNodes,
					PreOpCount:           uint64(0),
					PostOpCount:          uint64(1),
					ValidUntil:           validUntil,
					OverridePreviousRoot: false,
				})
				require.NoError(t, err)

				signatures, err := BulkSignOnMsgHash(canceller.Signers, rootValidationData.EthMsgHash)
				require.NoError(t, err)

				signaturesPDA := canceller.RootSignaturesPDA(rootValidationData.Root, validUntil)

				parsedTotalSigs, err := mcmsUtils.SafeToUint8(len(signatures))
				require.NoError(t, err)

				initSigsIx, err := mcm.NewInitSignaturesInstruction(
					canceller.PaddedName,
					rootValidationData.Root,
					validUntil,
					parsedTotalSigs,
					signaturesPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initSigsIx}, admin, config.DefaultCommitment)

				appendSigsIxs, err := AppendSignaturesIxs(
					signatures,
					canceller.PaddedName,
					rootValidationData.Root,
					validUntil,
					signaturesPDA,
					admin.PublicKey(),
					config.MaxAppendSignatureBatchSize,
				)
				require.NoError(t, err)
				for _, ix := range appendSigsIxs {
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				finalizeSigsIx, err := mcm.NewFinalizeSignaturesInstruction(
					canceller.PaddedName,
					rootValidationData.Root,
					validUntil,
					signaturesPDA,
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{finalizeSigsIx}, admin, config.DefaultCommitment)

				setRootIx, err := mcm.NewSetRootInstruction(
					canceller.PaddedName,
					rootValidationData.Root,
					validUntil,
					rootValidationData.Metadata,
					rootValidationData.MetadataProof,
					signaturesPDA,
					canceller.RootMetadataPDA,
					SeenSignedHashesAddress(canceller.PaddedName, rootValidationData.Root, validUntil),
					canceller.ExpiringRootAndOpCountPDA,
					canceller.ConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{setRootIx}, admin, config.DefaultCommitment)

				// execute mcm operation to cancel the timelock operation
				proofs, err := cancleOpNodes[0].Proofs()
				require.NoError(t, err)

				executeIx := mcm.NewExecuteInstruction(
					canceller.PaddedName,
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

				vIx, err := executeIx.ValidateAndBuild()
				require.NoError(t, err)

				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, anyone, config.DefaultCommitment)

				utils.AssertClosedAccount(ctx, t, solanaGoClient, op3.OperationPDA(), config.DefaultCommitment)
			})

			t.Run("create new operation with corrected amounts", func(t *testing.T) {
				// Create corrected transfer instructions with new amounts
				ix1, err := utils.TokenTransferChecked(150*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team1ATA, config.TimelockSignerPDA, []solana.PublicKey{})
				require.NoError(t, err)
				ix2, err := utils.TokenTransferChecked(150*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team2ATA, config.TimelockSignerPDA, []solana.PublicKey{})
				require.NoError(t, err)
				ix3, err := utils.TokenTransferChecked(100*solana.LAMPORTS_PER_SOL, 9, tokenProgram, treasuryATA, mint, team3ATA, config.TimelockSignerPDA, []solana.PublicKey{})
				require.NoError(t, err)

				// Create new operation
				salt, err := mcmsUtils.SimpleSalt()
				require.NoError(t, err)
				newOp3 = TimelockOperation{
					Predecessor: op2.OperationID(),
					Salt:        salt,
					Delay:       uint64(1),
				}

				newOp3.AddInstruction(ix1, []solana.PublicKey{tokenProgram})
				newOp3.AddInstruction(ix2, []solana.PublicKey{tokenProgram})
				newOp3.AddInstruction(ix3, []solana.PublicKey{tokenProgram})

				// Initialize operation account
				initOpIx, err := timelock.NewInitializeOperationInstruction(
					newOp3.OperationID(),
					newOp3.Predecessor,
					newOp3.Salt,
					newOp3.IxsCountU32(),
					config.TimelockConfigPDA,
					newOp3.OperationPDA(),
					admin.PublicKey(),
					proposerMsig.SignerPDA,
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initOpIx}, admin, config.DefaultCommitment)

				// Append instructions
				for _, ixData := range newOp3.ToInstructionData() {
					appendIx, appendIxErr := timelock.NewAppendInstructionsInstruction(
						newOp3.OperationID(),
						[]timelock.InstructionData{ixData},
						newOp3.OperationPDA(),
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, appendIxErr)
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{appendIx}, admin, config.DefaultCommitment)
				}

				// Finalize operation
				finalizeIx, err := timelock.NewFinalizeOperationInstruction(
					newOp3.OperationID(),
					newOp3.OperationPDA(),
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{finalizeIx}, admin, config.DefaultCommitment)

				// Create mcm operation node for scheduling
				scheduleIx, err := timelock.NewScheduleBatchInstruction(
					newOp3.OperationID(),
					newOp3.Delay,
					config.TimelockConfigPDA,
					newOp3.OperationPDA(),
					msigs[timelock.Proposer_Role].AccessController.PublicKey(),
					proposerMsig.SignerPDA,
				).ValidateAndBuild()
				require.NoError(t, err)

				node, err := IxToMcmTestOpNode(proposerMsig.ConfigPDA, proposerMsig.SignerPDA, scheduleIx, uint64(currentOpCount))
				require.NoError(t, err)

				newOpNodes := []mcmsUtils.McmOpNode{node}

				// Create and validate root data
				validUntil := uint32(0xffffffff)
				rootValidationData, err := CreateMcmRootData(McmRootInput{
					Multisig:             proposerMsig.ConfigPDA,
					Operations:           newOpNodes,
					PreOpCount:           uint64(currentOpCount),
					PostOpCount:          uint64(currentOpCount + 1),
					ValidUntil:           validUntil,
					OverridePreviousRoot: false,
				})
				require.NoError(t, err)

				currentOpCount++

				// Sign and set root
				signatures, err := BulkSignOnMsgHash(proposerMsig.Signers, rootValidationData.EthMsgHash)
				require.NoError(t, err)

				signaturesPDA := proposerMsig.RootSignaturesPDA(rootValidationData.Root, validUntil)

				// Initialize signatures
				parsedTotalSigs, err := mcmsUtils.SafeToUint8(len(signatures))
				require.NoError(t, err)

				initSigsIx, err := mcm.NewInitSignaturesInstruction(
					proposerMsig.PaddedName,
					rootValidationData.Root,
					validUntil,
					parsedTotalSigs,
					signaturesPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initSigsIx}, admin, config.DefaultCommitment)

				// Append signatures
				appendSigsIxs, err := AppendSignaturesIxs(
					signatures,
					proposerMsig.PaddedName,
					rootValidationData.Root,
					validUntil,
					signaturesPDA,
					admin.PublicKey(),
					config.MaxAppendSignatureBatchSize,
				)
				require.NoError(t, err)
				for _, ix := range appendSigsIxs {
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				// Finalize signatures
				finalizeSigsIx, err := mcm.NewFinalizeSignaturesInstruction(
					proposerMsig.PaddedName,
					rootValidationData.Root,
					validUntil,
					signaturesPDA,
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{finalizeSigsIx}, admin, config.DefaultCommitment)

				// Set root
				setRootIx, err := mcm.NewSetRootInstruction(
					proposerMsig.PaddedName,
					rootValidationData.Root,
					validUntil,
					rootValidationData.Metadata,
					rootValidationData.MetadataProof,
					signaturesPDA,
					proposerMsig.RootMetadataPDA,
					SeenSignedHashesAddress(proposerMsig.PaddedName, rootValidationData.Root, validUntil),
					proposerMsig.ExpiringRootAndOpCountPDA,
					proposerMsig.ConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{setRootIx}, admin, config.DefaultCommitment)

				// Execute mcm operation to schedule the timelock operation
				proofs, err := newOpNodes[0].Proofs()
				require.NoError(t, err)

				executeIx := mcm.NewExecuteInstruction(
					proposerMsig.PaddedName,
					config.TestChainID,
					node.Nonce,
					node.Data,
					proofs,
					proposerMsig.ConfigPDA,
					proposerMsig.RootMetadataPDA,
					proposerMsig.ExpiringRootAndOpCountPDA,
					node.To,
					proposerMsig.SignerPDA,
					anyone.PublicKey(),
				)
				executeIx.AccountMetaSlice = append(executeIx.AccountMetaSlice, node.RemainingAccounts...)

				vIx, err := executeIx.ValidateAndBuild()
				require.NoError(t, err)

				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, anyone, config.DefaultCommitment)
			})
		})

		t.Run("execute timelock operations", func(t *testing.T) {
			// Wait for operations to be ready
			err := WaitForOperationToBeReady(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment)
			require.NoError(t, err)

			rErr := WaitForOperationToBeReady(ctx, solanaGoClient, op2.OperationPDA(), config.DefaultCommitment)
			require.NoError(t, rErr)

			t.Run("op2: cannot be executed before op1", func(t *testing.T) {
				ix := timelock.NewExecuteBatchInstruction(
					op2.OperationID(),
					config.TimelockConfigPDA,
					config.TimelockSignerPDA,
					op2.OperationPDA(),
					op1.OperationPDA(), // provide op1 PDA as predecessor
					msigs[timelock.Executor_Role].AccessController.PublicKey(),
					admin.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

				vIx, err := ix.ValidateAndBuild()
				require.NoError(t, err)

				utils.SendAndFailWith(ctx, t, solanaGoClient,
					[]solana.Instruction{vIx},
					admin,
					config.DefaultCommitment,
					[]string{"Error Code: " + timelock.MissingDependency_TimelockError.String()},
				)
			})

			t.Run("op1: initial mint to treasury", func(t *testing.T) {
				ix := timelock.NewExecuteBatchInstruction(
					op1.OperationID(),
					config.TimelockConfigPDA,
					config.TimelockSignerPDA,
					op1.OperationPDA(),
					config.TimelockEmptyOpID,
					msigs[timelock.Executor_Role].AccessController.PublicKey(),
					admin.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op1.RemainingAccounts()...)

				vIx, err := ix.ValidateAndBuild()
				require.NoError(t, err)

				result := utils.SendAndConfirm(ctx, t, solanaGoClient,
					[]solana.Instruction{vIx},
					admin,
					config.DefaultCommitment,
					utils.AddComputeUnitLimit(1_400_000),
				)
				require.NotNil(t, result)

				// Verify operation status
				var opAccount timelock.Operation
				err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, op1.OperationPDA(), config.DefaultCommitment, &opAccount)
				require.NoError(t, err)
				require.Equal(t, config.TimelockOpDoneTimestamp, opAccount.Timestamp, "Op1 should be marked as executed")

				// Verify treasury balance
				_, treasuryBalance, err := utils.TokenBalance(ctx, solanaGoClient, treasuryATA, config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 1000*int(solana.LAMPORTS_PER_SOL), treasuryBalance,
					"Treasury should have received 1000 tokens")
			})

			t.Run("token approval to timelock signer", func(t *testing.T) {
				// fund treasury account first
				fundIx, err := system.NewTransferInstruction(
					1*solana.LAMPORTS_PER_SOL, // 1 SOL should be enough
					admin.PublicKey(),
					treasury.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{fundIx}, admin, config.DefaultCommitment)

				// approve can't be deligated to timelock authority(security - CPI Guard)
				approveIx, err := utils.TokenApproveChecked(
					600*solana.LAMPORTS_PER_SOL,
					9,
					tokenProgram,
					treasuryATA,
					mint,
					config.TimelockSignerPDA,
					treasury.PublicKey(),
					nil,
				)
				require.NoError(t, err)
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{approveIx}, treasury, config.DefaultCommitment)
			})

			t.Run("op2: should provide the correct predecessor pda address", func(t *testing.T) {
				ix := timelock.NewExecuteBatchInstruction(
					op2.OperationID(),
					config.TimelockConfigPDA,
					config.TimelockSignerPDA,
					op2.OperationPDA(),
					op1.OperationID(), // provide op1 ID as predecessor
					msigs[timelock.Executor_Role].AccessController.PublicKey(),
					admin.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

				vIx, err := ix.ValidateAndBuild()
				require.NoError(t, err)

				utils.SendAndFailWith(ctx, t, solanaGoClient,
					[]solana.Instruction{vIx},
					admin,
					config.DefaultCommitment,
					[]string{"Error Code: " + timelock.InvalidInput_TimelockError.String()},
				)
			})

			t.Run("op2: team ata creation", func(t *testing.T) {
				ix := timelock.NewExecuteBatchInstruction(
					op2.OperationID(),
					config.TimelockConfigPDA,
					config.TimelockSignerPDA,
					op2.OperationPDA(),
					op1.OperationPDA(), // provide op1 PDA as predecessor
					msigs[timelock.Executor_Role].AccessController.PublicKey(),
					admin.PublicKey(),
				)
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op2.RemainingAccounts()...)

				vIx, err := ix.ValidateAndBuild()
				require.NoError(t, err)

				result := utils.SendAndConfirm(ctx, t, solanaGoClient,
					[]solana.Instruction{vIx},
					admin,
					config.DefaultCommitment,
					utils.AddComputeUnitLimit(1_400_000),
				)
				require.NotNil(t, result)

				// verify operation status
				var opAccount timelock.Operation
				err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, op2.OperationPDA(), config.DefaultCommitment, &opAccount)
				require.NoError(t, err)
				require.Equal(t, config.TimelockOpDoneTimestamp, opAccount.Timestamp, "Op2 should be marked as executed")
			})
		})

		t.Run("op3: team token distribution", func(t *testing.T) {
			// Wait for delay and execute the timelock operation
			time.Sleep(time.Duration(newOp3.Delay) * time.Second)

			executeTimelockIx := timelock.NewExecuteBatchInstruction(
				newOp3.OperationID(),
				config.TimelockConfigPDA,
				config.TimelockSignerPDA,
				newOp3.OperationPDA(),
				op2.OperationPDA(),
				msigs[timelock.Executor_Role].AccessController.PublicKey(),
				admin.PublicKey(),
			)
			executeTimelockIx.AccountMetaSlice = append(executeTimelockIx.AccountMetaSlice, newOp3.RemainingAccounts()...)

			vTimelockIx, err := executeTimelockIx.ValidateAndBuild()
			require.NoError(t, err)

			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vTimelockIx}, admin, config.DefaultCommitment)

			// Verify final balances
			_, treasuryBalance, err := utils.TokenBalance(ctx, solanaGoClient, treasuryATA, config.DefaultCommitment)
			require.NoError(t, err)
			require.Equal(t, 600*int(solana.LAMPORTS_PER_SOL), treasuryBalance,
				"Treasury should have 600 tokens remaining after distributions")

			_, team1Balance, err := utils.TokenBalance(ctx, solanaGoClient, team1ATA, config.DefaultCommitment)
			require.NoError(t, err)
			require.Equal(t, 150*int(solana.LAMPORTS_PER_SOL), team1Balance,
				"Team1 should have received 150 tokens")

			_, team2Balance, err := utils.TokenBalance(ctx, solanaGoClient, team2ATA, config.DefaultCommitment)
			require.NoError(t, err)
			require.Equal(t, 150*int(solana.LAMPORTS_PER_SOL), team2Balance,
				"Team2 should have received 150 tokens")

			_, team3Balance, err := utils.TokenBalance(ctx, solanaGoClient, team3ATA, config.DefaultCommitment)
			require.NoError(t, err)
			require.Equal(t, 100*int(solana.LAMPORTS_PER_SOL), team3Balance,
				"Team3 should have received 100 tokens")
		})
	})
}
