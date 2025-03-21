package contracts

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/mcms"
)

func TestMcmSetConfig(t *testing.T) {
	t.Parallel()
	mcm.SetProgramID(config.McmProgram)

	ctx := tests.Context(t)

	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	anotherAdmin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	user, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)

	// mcm name
	testMsigID := config.TestMsigID

	// test mcm pdas
	multisigConfigPDA := mcms.GetConfigPDA(testMsigID)
	rootMetadataPDA := mcms.GetRootMetadataPDA(testMsigID)
	expiringRootAndOpCountPDA := mcms.GetExpiringRootAndOpCountPDA(testMsigID)
	configSignersPDA := mcms.GetConfigSignersPDA(testMsigID)

	t.Run("setup:funding", func(t *testing.T) {
		testutils.FundAccounts(ctx, []solana.PrivateKey{admin, anotherAdmin, user}, solanaGoClient, t)
	})

	t.Run("setup:mcm", func(t *testing.T) {
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

		ix, err := mcm.NewInitializeInstruction(
			config.TestChainID,
			testMsigID,
			multisigConfigPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
			config.McmProgram,
			programData.Address,
			rootMetadataPDA,
			expiringRootAndOpCountPDA,
		).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

		// get config and validate
		var configAccount mcm.MultisigConfig
		err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
		require.NoError(t, err, "failed to get account info")

		require.Equal(t, config.TestChainID, configAccount.ChainId)
		require.Equal(t, admin.PublicKey(), configAccount.Owner)
	})

	t.Run("mcm:ownership", func(t *testing.T) {
		// Fail to transfer ownership when not owner
		instruction, err := mcm.NewTransferOwnershipInstruction(
			testMsigID,
			anotherAdmin.PublicKey(),
			multisigConfigPDA,
			user.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + mcms.UnauthorizedMcmError.String()})
		require.NotNil(t, result)

		// successfully transfer ownership
		instruction, err = mcm.NewTransferOwnershipInstruction(
			testMsigID,
			anotherAdmin.PublicKey(),
			multisigConfigPDA,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment)
		require.NotNil(t, result)

		// Fail to accept ownership when not proposed_owner
		instruction, err = mcm.NewAcceptOwnershipInstruction(
			testMsigID,
			multisigConfigPDA,
			user.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + mcms.UnauthorizedMcmError.String()})
		require.NotNil(t, result)

		// Successfully accept ownership
		// anotherAdmin becomes owner for remaining tests
		instruction, err = mcm.NewAcceptOwnershipInstruction(
			testMsigID,
			multisigConfigPDA,
			anotherAdmin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment)
		require.NotNil(t, result)

		// Current owner cannot propose self
		instruction, err = mcm.NewTransferOwnershipInstruction(
			testMsigID,
			anotherAdmin.PublicKey(),
			multisigConfigPDA,
			anotherAdmin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment, []string{"Error Code: " + mcm.InvalidInputs_McmError.String()})
		require.NotNil(t, result)

		// Validate proposed set to 0-address after accepting ownership
		var configAccount mcm.MultisigConfig
		err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
		if err != nil {
			require.NoError(t, err, "failed to get account info")
		}
		require.Equal(t, anotherAdmin.PublicKey(), configAccount.Owner)
		require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)

		// get it back
		instruction, err = mcm.NewTransferOwnershipInstruction(
			testMsigID,
			admin.PublicKey(),
			multisigConfigPDA,
			anotherAdmin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment)
		require.NotNil(t, result)

		instruction, err = mcm.NewAcceptOwnershipInstruction(
			testMsigID,
			multisigConfigPDA,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment)
		require.NotNil(t, result)

		err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
		if err != nil {
			require.NoError(t, err, "failed to get account info")
		}

		require.Equal(t, admin.PublicKey(), configAccount.Owner)
		require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
	})

	t.Run("mcm:set_config", func(t *testing.T) {
		numSigners := config.MaxNumSigners
		signerPrivateKeys, err := eth.GenerateEthPrivateKeys(numSigners)
		require.NoError(t, err)

		signerGroups := make([]byte, numSigners)
		for i := 0; i < len(signerGroups); i++ {
			signerGroups[i] = byte(i % 10)
		}

		// just use simple config for now
		groupQuorums := []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		groupParents := []uint8{0, 0, 0, 2, 0, 0, 0, 0, 0, 0}

		mcmConfig, err := mcms.NewValidMcmConfig(
			testMsigID,
			signerPrivateKeys,
			signerGroups,
			groupQuorums,
			groupParents,
			config.ClearRoot,
		)
		require.NoError(t, err)

		signerAddresses := mcmConfig.SignerAddresses

		t.Run("mcm:set_config: preload signers on PDA", func(t *testing.T) {
			preloadIxs, pierr := mcms.GetPreloadSignersIxs(signerAddresses, testMsigID, multisigConfigPDA, configSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
			require.NoError(t, pierr)

			for _, ix := range preloadIxs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			var cfgSignersAccount mcm.ConfigSigners
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, configSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, true, cfgSignersAccount.IsFinalized)

			// check if the addresses are registered correctly
			for i, signer := range cfgSignersAccount.SignerAddresses {
				require.Equal(t, signerAddresses[i], signer)
			}
		})

		t.Run("mcm:set_config:admin authorization", func(t *testing.T) {
			t.Run("fail:set_config from unauthorized user", func(t *testing.T) {
				ix, err := mcm.NewSetConfigInstruction(
					mcmConfig.MultisigID,
					mcmConfig.SignerGroups,
					mcmConfig.GroupQuorums,
					mcmConfig.GroupParents,
					mcmConfig.ClearRoot,
					multisigConfigPDA,
					configSignersPDA,
					rootMetadataPDA,
					expiringRootAndOpCountPDA,
					user.PublicKey(), // unauthorized user
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + mcms.UnauthorizedMcmError.String()})
				require.NotNil(t, result)
			})

			t.Run("success:set_config from admin", func(t *testing.T) {
				// set config
				ix, err := mcm.NewSetConfigInstruction(
					mcmConfig.MultisigID,
					mcmConfig.SignerGroups,
					mcmConfig.GroupQuorums,
					mcmConfig.GroupParents,
					mcmConfig.ClearRoot,
					multisigConfigPDA,
					configSignersPDA,
					rootMetadataPDA,
					expiringRootAndOpCountPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, err)

				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				require.NotNil(t, tx)

				parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
					[]common.EventMapping{
						common.EventMappingFor[mcms.ConfigSet]("ConfigSet"),
					},
				)

				event := parsedLogs[0].EventData[0].Data.(*mcms.ConfigSet)
				require.Equal(t, mcmConfig.GroupParents, event.GroupParents)
				require.Equal(t, mcmConfig.GroupQuorums, event.GroupQuorums)
				require.Equal(t, mcmConfig.ClearRoot, event.IsRootCleared)

				for i, signer := range event.Signers {
					require.Equal(t, mcmConfig.SignerAddresses[i], signer.EvmAddress)
					require.Equal(t, uint8(i), signer.Index)
					require.Equal(t, mcmConfig.SignerGroups[i], signer.Group)
				}

				// get config and validate
				var configAccount mcm.MultisigConfig
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
				require.NoError(t, err, "failed to get account info")

				require.Equal(t, config.TestChainID, configAccount.ChainId)
				require.Equal(t, reflect.DeepEqual(configAccount.GroupParents, mcmConfig.GroupParents), true)
				require.Equal(t, reflect.DeepEqual(configAccount.GroupQuorums, mcmConfig.GroupQuorums), true)

				// check if the McmSigner struct is correct
				for i, signer := range configAccount.Signers {
					require.Equal(t, signer.EvmAddress, mcmConfig.SignerAddresses[i])
					require.Equal(t, signer.Index, uint8(i))
					require.Equal(t, signer.Group, (mcmConfig.SignerGroups)[i])
				}

				// pda closed after set_config
				testutils.AssertClosedAccount(ctx, t, solanaGoClient, configSignersPDA, config.DefaultCommitment)
			})
		})
	})

	t.Run("mcm:set_config with reinitializing signers pda(closed)", func(t *testing.T) {
		numSigners := config.MaxNumSigners
		signerPrivateKeys, err := eth.GenerateEthPrivateKeys(numSigners)
		require.NoError(t, err)

		signerGroups := make([]byte, numSigners)
		for i := 0; i < len(signerGroups); i++ {
			signerGroups[i] = byte(i % 10)
		}

		// just use simple config for now
		groupQuorums := []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		groupParents := []uint8{0, 0, 0, 2, 0, 0, 0, 0, 0, 0}

		mcmConfig, err := mcms.NewValidMcmConfig(
			testMsigID,
			signerPrivateKeys,
			signerGroups,
			groupQuorums,
			groupParents,
			config.ClearRoot,
		)
		require.NoError(t, err)

		signerAddresses := mcmConfig.SignerAddresses

		t.Run("mcm:set_config: preload signers on PDA", func(t *testing.T) {
			// ConfigSignersPDA should be closed before reinitializing
			testutils.AssertClosedAccount(ctx, t, solanaGoClient, configSignersPDA, config.DefaultCommitment)
			//nolint:gosec
			parsedTotalSigners := uint8(len(signerAddresses))
			initSignersIx, err := mcm.NewInitSignersInstruction(
				testMsigID,
				parsedTotalSigners,
				multisigConfigPDA,
				configSignersPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()

			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initSignersIx}, admin, config.DefaultCommitment)

			appendSignersIxs, err := mcms.GetAppendSignersIxs(signerAddresses, testMsigID, multisigConfigPDA, configSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
			require.NoError(t, err)

			// partially register signers
			for _, ix := range appendSignersIxs[:3] {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			// clear signers(this closes the account)
			clearIx, err := mcm.NewClearSignersInstruction(
				testMsigID,
				multisigConfigPDA,
				configSignersPDA,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{clearIx}, admin, config.DefaultCommitment)
			testutils.AssertClosedAccount(ctx, t, solanaGoClient, configSignersPDA, config.DefaultCommitment)

			// preload signers again
			preloadIxs, pierr := mcms.GetPreloadSignersIxs(signerAddresses, testMsigID, multisigConfigPDA, configSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
			require.NoError(t, pierr)

			for _, ix := range preloadIxs {
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			var cfgSignersAccount mcm.ConfigSigners
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, configSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, true, cfgSignersAccount.IsFinalized)

			// check if the addresses are registered correctly
			for i, signer := range cfgSignersAccount.SignerAddresses {
				require.Equal(t, signerAddresses[i], signer)
			}
		})

		t.Run("success:set_config", func(t *testing.T) {
			ix, err := mcm.NewSetConfigInstruction(
				testMsigID,
				mcmConfig.SignerGroups,
				mcmConfig.GroupQuorums,
				mcmConfig.GroupParents,
				mcmConfig.ClearRoot,
				multisigConfigPDA,
				configSignersPDA,
				rootMetadataPDA,
				expiringRootAndOpCountPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()

			require.NoError(t, err)

			tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			require.NotNil(t, tx)

			parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
				[]common.EventMapping{
					common.EventMappingFor[mcms.ConfigSet]("ConfigSet"),
				},
			)

			event := parsedLogs[0].EventData[0].Data.(*mcms.ConfigSet)
			require.Equal(t, mcmConfig.GroupParents, event.GroupParents)
			require.Equal(t, mcmConfig.GroupQuorums, event.GroupQuorums)
			require.Equal(t, mcmConfig.ClearRoot, event.IsRootCleared)
			for i, signer := range event.Signers {
				require.Equal(t, mcmConfig.SignerAddresses[i], signer.EvmAddress)
				require.Equal(t, uint8(i), signer.Index)
				require.Equal(t, mcmConfig.SignerGroups[i], signer.Group)
			}

			// get config and validate
			var configAccount mcm.MultisigConfig
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, config.TestChainID, configAccount.ChainId)
			require.Equal(t, reflect.DeepEqual(configAccount.GroupParents, mcmConfig.GroupParents), true)
			require.Equal(t, reflect.DeepEqual(configAccount.GroupQuorums, mcmConfig.GroupQuorums), true)

			// check if the McmSigner struct is correct
			for i, signer := range configAccount.Signers {
				require.Equal(t, signer.EvmAddress, mcmConfig.SignerAddresses[i])
				require.Equal(t, signer.Index, uint8(i))
				require.Equal(t, signer.Group, (mcmConfig.SignerGroups)[i])
			}

			// pda closed after set_config
			testutils.AssertClosedAccount(ctx, t, solanaGoClient, configSignersPDA, config.DefaultCommitment)
		})
	})

	t.Run("set_config validation", func(t *testing.T) {
		tests := []struct {
			name                string
			errorMsg            string
			modifyConfig        func(*mcms.McmConfigArgs)
			skipPreloadSigners  bool
			skipFinalizeSigners bool
		}{
			{
				name:               "should not be able to call set_config without preloading config_signers",
				errorMsg:           "Error Code: " + common.AccountNotInitialized_AnchorError.String(),
				modifyConfig:       func(c *mcms.McmConfigArgs) {},
				skipPreloadSigners: true,
			},
			{
				name:                "should not be able to call set_config without finalized config_signers",
				errorMsg:            "Error Code: " + mcm.SignersNotFinalized_McmError.String(),
				modifyConfig:        func(c *mcms.McmConfigArgs) {},
				skipFinalizeSigners: true,
			},
			{
				name:     "length of signer addresses and signer groups length should be equal",
				errorMsg: "Error Code: " + mcm.MismatchedInputSignerVectorsLength_McmError.String(),
				modifyConfig: func(c *mcms.McmConfigArgs) {
					c.SignerGroups = append(c.SignerGroups, 1)
				},
			},
			{
				name:     "every group index in signer group should be less than NUM_GROUPS",
				errorMsg: "Error Code: " + mcm.MismatchedInputGroupArraysLength_McmError.String(),
				modifyConfig: func(c *mcms.McmConfigArgs) {
					(c.SignerGroups)[0] = 32
				},
			},
			{
				name:     "the parent of root has to be 0",
				errorMsg: "Error Code: " + mcm.GroupTreeNotWellFormed_McmError.String(),
				modifyConfig: func(c *mcms.McmConfigArgs) {
					(c.GroupParents)[0] = 1
				},
			},
			{
				name:     "the parent group should be at a higher index than the child group",
				errorMsg: "Error Code: " + mcm.GroupTreeNotWellFormed_McmError.String(),
				modifyConfig: func(c *mcms.McmConfigArgs) {
					(c.GroupParents)[1] = 2
				},
			},
			{
				name:     "disabled group(with 0 quorum) should not have a signer",
				errorMsg: "Error Code: " + mcm.SignerInDisabledGroup_McmError.String(),
				modifyConfig: func(c *mcms.McmConfigArgs) {
					(c.GroupQuorums)[3] = 0 // set quorum of group 3 to 0, but we still have signers in group 3
				},
			},
			{
				name:     "the group quorum should be able to be met(i.e. have more signers than the quorum)",
				errorMsg: "Error Code: " + mcm.OutOfBoundsGroupQuorum_McmError.String(),
				modifyConfig: func(c *mcms.McmConfigArgs) {
					(c.GroupQuorums)[3] = 3 // set quorum of group 3 to 3, but we have two signers in group 3
				},
			},
		}

		for i, tt := range tests {
			t.Run(fmt.Sprintf("set_config validation:%s", tt.name), func(t *testing.T) {
				t.Parallel()

				// use different msig accounts per test
				failTestMsigID, err := mcms.PadString32(fmt.Sprintf("fail_test_%d", i))
				require.NoError(t, err)

				// test scope mcm pdas
				failMultisigConfigPDA := mcms.GetConfigPDA(failTestMsigID)
				failRootMetadataPDA := mcms.GetRootMetadataPDA(failTestMsigID)
				failExpiringRootAndOpCountPDA := mcms.GetExpiringRootAndOpCountPDA(failTestMsigID)
				failConfigSignersPDA := mcms.GetConfigSignersPDA(failTestMsigID)

				t.Run(fmt.Sprintf("msig initialization:%s", tt.name), func(t *testing.T) {
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

					// initialize msig
					ix, initIxErr := mcm.NewInitializeInstruction(
						config.TestChainID,
						failTestMsigID,
						failMultisigConfigPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
						config.McmProgram,
						programData.Address,
						failRootMetadataPDA,
						failExpiringRootAndOpCountPDA,
					).ValidateAndBuild()
					require.NoError(t, initIxErr)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				})

				cfg, err := mcms.NewValidMcmConfig(
					failTestMsigID,
					config.SignerPrivateKeys,
					config.SignerGroups,
					config.GroupQuorums,
					config.GroupParents,
					config.ClearRoot,
				)
				require.NoError(t, err)

				t.Run("preload signers for validation tests", func(t *testing.T) {
					if tt.skipPreloadSigners {
						return
					}
					//nolint:gosec
					parsedTotalSigners := uint8(len(cfg.SignerAddresses))
					initSignersIx, initSignersErr := mcm.NewInitSignersInstruction(
						failTestMsigID,
						parsedTotalSigners,
						failMultisigConfigPDA,
						failConfigSignersPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()

					require.NoError(t, initSignersErr)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initSignersIx}, admin, config.DefaultCommitment)

					appendSignersIxs, appendSignersIxsErr := mcms.GetAppendSignersIxs(cfg.SignerAddresses, failTestMsigID, failMultisigConfigPDA, failConfigSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
					require.NoError(t, appendSignersIxsErr)
					for _, ix := range appendSignersIxs {
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
					}

					if !tt.skipFinalizeSigners {
						finalizeSignersIx, finSignersIxErr := mcm.NewFinalizeSignersInstruction(
							failTestMsigID,
							failMultisigConfigPDA,
							failConfigSignersPDA,
							admin.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, finSignersIxErr)

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{finalizeSignersIx}, admin, config.DefaultCommitment)

						var cfgSignersAccount mcm.ConfigSigners
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, failConfigSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
						require.NoError(t, err, "failed to get account info")

						require.Equal(t, true, cfgSignersAccount.IsFinalized)

						// check if the addresses are registered correctly
						for i, signer := range cfgSignersAccount.SignerAddresses {
							require.Equal(t, cfg.SignerAddresses[i], signer)
						}
					}
				})

				// corrupt the config
				tt.modifyConfig(cfg)

				ix, err := mcm.NewSetConfigInstruction(
					cfg.MultisigID,
					cfg.SignerGroups,
					cfg.GroupQuorums,
					cfg.GroupParents,
					cfg.ClearRoot,
					failMultisigConfigPDA,
					failConfigSignersPDA,
					failRootMetadataPDA,
					failExpiringRootAndOpCountPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed, []string{tt.errorMsg})
				require.NotNil(t, result)
			})
		}
	})

	t.Run("pre-uploading config_signers validations", func(t *testing.T) {
		type TestStage int

		const (
			InitStage TestStage = iota
			AppendStage
			FinalizeStage
		)

		type TxWithStage struct {
			Instructions []solana.Instruction
			Stage        TestStage
		}

		tests := []struct {
			name               string
			errorMsg           string
			modifyConfig       func(*mcms.McmConfigArgs)
			failureStage       TestStage
			skipInitSigners    bool
			totalSignersOffset int
		}{
			{
				name:     "should not be able to initialize config_signers with empty",
				errorMsg: "Error Code: " + mcm.OutOfBoundsNumOfSigners_McmError.String(),
				modifyConfig: func(c *mcms.McmConfigArgs) {
					// empty cfg.SignerAddresses
					c.SignerAddresses = make([][20]byte, 0)
				},
				failureStage: InitStage,
			},
			{
				name:     "should not be able to initialize config_signers with more than MAX_NUM_SIGNERS",
				errorMsg: "Error Code: " + mcm.OutOfBoundsNumOfSigners_McmError.String(),
				modifyConfig: func(c *mcms.McmConfigArgs) {
					// replace cfg.SignerAddresses with more than MAX_NUM_SIGNERS(200)
					privateKeys, err := eth.GenerateEthPrivateKeys(201)
					require.NoError(t, err)
					signers, err := eth.GetEvmSigners(privateKeys)
					require.NoError(t, err)
					signerAddresses := make([][20]byte, 0)
					for _, signer := range signers {
						signerAddresses = append(signerAddresses, signer.Address)
					}
					c.SignerAddresses = signerAddresses
				},
				failureStage: InitStage,
			},
			{
				name:            "should not be able to append signers without initializing",
				errorMsg:        "Error Code: " + common.AccountNotInitialized_AnchorError.String(),
				modifyConfig:    func(c *mcms.McmConfigArgs) {},
				failureStage:    AppendStage,
				skipInitSigners: true,
			},
			{
				name:     "should not be able to append unsorted signer",
				errorMsg: "Error Code: " + mcm.SignersAddressesMustBeStrictlyIncreasing_McmError.String(),
				modifyConfig: func(c *mcms.McmConfigArgs) {
					slices.Reverse(c.SignerAddresses)
				},
				failureStage: AppendStage,
			},
			{
				name:               "should not be able to append more signers than specified in total_signers",
				errorMsg:           "Error Code: " + mcm.OutOfBoundsNumOfSigners_McmError.String(),
				modifyConfig:       func(c *mcms.McmConfigArgs) {},
				failureStage:       AppendStage,
				totalSignersOffset: -2,
			},
			{
				name:               "should not be able to finalize unmatched total signers",
				errorMsg:           "Error Code: " + mcm.OutOfBoundsNumOfSigners_McmError.String(),
				modifyConfig:       func(c *mcms.McmConfigArgs) {},
				failureStage:       FinalizeStage,
				totalSignersOffset: 2,
			},
		}

		for i, tt := range tests {
			t.Run(fmt.Sprintf("set_config validation:%s", tt.name), func(t *testing.T) {
				t.Parallel()

				// use different msig accounts per test
				failTestMsigID, err := mcms.PadString32(fmt.Sprintf("fail_preupload_signer_test_%d", i))
				require.NoError(t, err)

				// test scope mcm pdas
				failMultisigConfigPDA := mcms.GetConfigPDA(failTestMsigID)
				failRootMetadataPDA := mcms.GetRootMetadataPDA(failTestMsigID)
				failExpiringRootAndOpCountPDA := mcms.GetExpiringRootAndOpCountPDA(failTestMsigID)
				failConfigSignersPDA := mcms.GetConfigSignersPDA(failTestMsigID)

				t.Run(fmt.Sprintf("msig initialization:%s", tt.name), func(t *testing.T) {
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

					// initialize msig
					ix, initIxErr := mcm.NewInitializeInstruction(
						config.TestChainID,
						failTestMsigID,
						failMultisigConfigPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
						config.McmProgram,
						programData.Address,
						failRootMetadataPDA,
						failExpiringRootAndOpCountPDA,
					).ValidateAndBuild()
					require.NoError(t, initIxErr)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				})

				cfg, err := mcms.NewValidMcmConfig(
					failTestMsigID,
					config.SignerPrivateKeys,
					config.SignerGroups,
					config.GroupQuorums,
					config.GroupParents,
					config.ClearRoot,
				)
				require.NoError(t, err)

				// corrupt the config if needed
				tt.modifyConfig(cfg)

				var txs []TxWithStage

				if !tt.skipInitSigners {
					actualLength := len(cfg.SignerAddresses)
					//nolint:gosec
					totalSigners := uint8(actualLength + tt.totalSignersOffset) // offset for the test

					initSignersIx, _ := mcm.NewInitSignersInstruction(
						failTestMsigID,
						totalSigners,
						failMultisigConfigPDA,
						failConfigSignersPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					txs = append(txs, TxWithStage{
						Instructions: []solana.Instruction{initSignersIx},
						Stage:        InitStage,
					})
				}

				appendIxs, _ := mcms.GetAppendSignersIxs(
					cfg.SignerAddresses,
					failTestMsigID,
					failMultisigConfigPDA,
					failConfigSignersPDA,
					admin.PublicKey(),
					config.MaxAppendSignerBatchSize,
				)
				for _, ix := range appendIxs {
					txs = append(txs, TxWithStage{
						Instructions: []solana.Instruction{ix},
						Stage:        AppendStage,
					})
				}

				finalizeIx, _ := mcm.NewFinalizeSignersInstruction(
					failTestMsigID,
					failMultisigConfigPDA,
					failConfigSignersPDA,
					admin.PublicKey(),
				).ValidateAndBuild()
				txs = append(txs, TxWithStage{
					Instructions: []solana.Instruction{finalizeIx},
					Stage:        FinalizeStage,
				})

				for _, tx := range txs {
					if tx.Stage == tt.failureStage {
						// this stage should fail
						result := testutils.SendAndFailWith(ctx, t, solanaGoClient,
							tx.Instructions,
							admin,
							rpc.CommitmentConfirmed,
							[]string{tt.errorMsg},
						)
						require.NotNil(t, result)
						break
					}

					// all other instructions should succeed
					testutils.SendAndConfirm(ctx, t, solanaGoClient,
						tx.Instructions,
						admin,
						config.DefaultCommitment,
					)
				}
			})
		}
	})
}
