package contracts

import (
	"reflect"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/generated/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/eth"
	mcmsUtils "github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/mcms"
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

	solanaGoClient := utils.DeployAllPrograms(t, utils.PathToAnchorConfig, admin)

	// mcm name
	TestMsigName := config.TestMsigNamePaddedBuffer

	// test mcm pdas
	MultisigConfigPDA := McmConfigAddress(TestMsigName)
	RootMetadataPDA := RootMetadataAddress(TestMsigName)
	ExpiringRootAndOpCountPDA := ExpiringRootAndOpCountAddress(TestMsigName)
	ConfigSignersPDA := McmConfigSignersAddress(TestMsigName)

	t.Run("setup:funding", func(t *testing.T) {
		utils.FundAccounts(ctx, []solana.PrivateKey{admin, anotherAdmin, user}, solanaGoClient, t)
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
			TestMsigName,
			MultisigConfigPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
			config.McmProgram,
			programData.Address,
			RootMetadataPDA,
			ExpiringRootAndOpCountPDA,
		).ValidateAndBuild()
		require.NoError(t, err)
		utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

		// get config and validate
		var configAccount mcm.MultisigConfig
		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA, config.DefaultCommitment, &configAccount)
		require.NoError(t, err, "failed to get account info")

		require.Equal(t, config.TestChainID, configAccount.ChainId)
		require.Equal(t, admin.PublicKey(), configAccount.Owner)
	})

	t.Run("mcm:ownership", func(t *testing.T) {
		// Fail to transfer ownership when not owner
		instruction, err := mcm.NewTransferOwnershipInstruction(
			TestMsigName,
			anotherAdmin.PublicKey(),
			MultisigConfigPDA,
			user.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result := utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + UnauthorizedMcmError.String()})
		require.NotNil(t, result)

		// successfully transfer ownership
		instruction, err = mcm.NewTransferOwnershipInstruction(
			TestMsigName,
			anotherAdmin.PublicKey(),
			MultisigConfigPDA,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment)
		require.NotNil(t, result)

		// Fail to accept ownership when not proposed_owner
		instruction, err = mcm.NewAcceptOwnershipInstruction(
			TestMsigName,
			MultisigConfigPDA,
			user.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + UnauthorizedMcmError.String()})
		require.NotNil(t, result)

		// Successfully accept ownership
		// anotherAdmin becomes owner for remaining tests
		instruction, err = mcm.NewAcceptOwnershipInstruction(
			TestMsigName,
			MultisigConfigPDA,
			anotherAdmin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment)
		require.NotNil(t, result)

		// Current owner cannot propose self
		instruction, err = mcm.NewTransferOwnershipInstruction(
			TestMsigName,
			anotherAdmin.PublicKey(),
			MultisigConfigPDA,
			anotherAdmin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment, []string{"Error Code: " + InvalidInputsMcmError.String()})
		require.NotNil(t, result)

		// Validate proposed set to 0-address after accepting ownership
		var configAccount mcm.MultisigConfig
		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA, config.DefaultCommitment, &configAccount)
		if err != nil {
			require.NoError(t, err, "failed to get account info")
		}
		require.Equal(t, anotherAdmin.PublicKey(), configAccount.Owner)
		require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)

		// get it back
		instruction, err = mcm.NewTransferOwnershipInstruction(
			TestMsigName,
			admin.PublicKey(),
			MultisigConfigPDA,
			anotherAdmin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment)
		require.NotNil(t, result)

		instruction, err = mcm.NewAcceptOwnershipInstruction(
			TestMsigName,
			MultisigConfigPDA,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result = utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment)
		require.NotNil(t, result)

		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA, config.DefaultCommitment, &configAccount)
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

		mcmConfig, err := mcmsUtils.NewValidMcmConfig(
			TestMsigName,
			signerPrivateKeys,
			signerGroups,
			groupQuorums,
			groupParents,
			config.ClearRoot,
		)
		require.NoError(t, err)

		signerAddresses := mcmConfig.SignerAddresses

		t.Run("mcm:set_config: preload signers on PDA", func(t *testing.T) {
			ixs := make([]solana.Instruction, 0)
			parsedTotalSigners, err := mcmsUtils.SafeToUint8(len(signerAddresses))
			require.NoError(t, err)

			initSignersIx, err := mcm.NewInitSignersInstruction(
				TestMsigName,
				parsedTotalSigners,
				MultisigConfigPDA,
				ConfigSignersPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()

			require.NoError(t, err)
			ixs = append(ixs, initSignersIx)

			appendSignersIxs, err := AppendSignersIxs(signerAddresses, TestMsigName, MultisigConfigPDA, ConfigSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
			require.NoError(t, err)
			ixs = append(ixs, appendSignersIxs...)

			finalizeSignersIx, err := mcm.NewFinalizeSignersInstruction(
				TestMsigName,
				MultisigConfigPDA,
				ConfigSignersPDA,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			ixs = append(ixs, finalizeSignersIx)

			for _, ix := range ixs {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			var cfgSignersAccount mcm.ConfigSigners
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, ConfigSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
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
					mcmConfig.MultisigName,
					mcmConfig.SignerGroups,
					mcmConfig.GroupQuorums,
					mcmConfig.GroupParents,
					mcmConfig.ClearRoot,
					MultisigConfigPDA,
					ConfigSignersPDA,
					user.PublicKey(), // unauthorized user
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				result := utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + UnauthorizedMcmError.String()})
				require.NotNil(t, result)
			})

			t.Run("success:set_config from admin", func(t *testing.T) {
				// set config
				ix, err := mcm.NewSetConfigInstruction(
					mcmConfig.MultisigName,
					mcmConfig.SignerGroups,
					mcmConfig.GroupQuorums,
					mcmConfig.GroupParents,
					mcmConfig.ClearRoot,
					MultisigConfigPDA,
					ConfigSignersPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, err)

				result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				require.NotNil(t, result)

				// get config and validate
				var configAccount mcm.MultisigConfig
				err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA, config.DefaultCommitment, &configAccount)
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
				utils.AssertClosedAccount(ctx, t, solanaGoClient, ConfigSignersPDA, config.DefaultCommitment)
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

		mcmConfig, err := mcmsUtils.NewValidMcmConfig(
			TestMsigName,
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
			utils.AssertClosedAccount(ctx, t, solanaGoClient, ConfigSignersPDA, config.DefaultCommitment)

			parsedTotalSigners, err := mcmsUtils.SafeToUint8(len(signerAddresses))
			require.NoError(t, err)

			initSignersIx, err := mcm.NewInitSignersInstruction(
				TestMsigName,
				parsedTotalSigners,
				MultisigConfigPDA,
				ConfigSignersPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()

			require.NoError(t, err)
			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initSignersIx}, admin, config.DefaultCommitment)

			appendSignersIxs, err := AppendSignersIxs(signerAddresses, TestMsigName, MultisigConfigPDA, ConfigSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
			require.NoError(t, err)

			// partially register signers
			for _, ix := range appendSignersIxs[:3] {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			// clear signers
			clearIx, err := mcm.NewClearSignersInstruction(
				TestMsigName,
				MultisigConfigPDA,
				ConfigSignersPDA,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{clearIx}, admin, config.DefaultCommitment)

			// register all signers
			for _, ix := range appendSignersIxs {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			// finalize registration
			finalizeSignersIx, err := mcm.NewFinalizeSignersInstruction(
				TestMsigName,
				MultisigConfigPDA,
				ConfigSignersPDA,
				admin.PublicKey(),
			).ValidateAndBuild()

			require.NoError(t, err)
			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{finalizeSignersIx}, admin, config.DefaultCommitment)

			var cfgSignersAccount mcm.ConfigSigners
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, ConfigSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, true, cfgSignersAccount.IsFinalized)

			// check if the addresses are registered correctly
			for i, signer := range cfgSignersAccount.SignerAddresses {
				require.Equal(t, signerAddresses[i], signer)
			}
		})

		t.Run("success:set_config", func(t *testing.T) {
			ix, err := mcm.NewSetConfigInstruction(
				TestMsigName,
				mcmConfig.SignerGroups,
				mcmConfig.GroupQuorums,
				mcmConfig.GroupParents,
				mcmConfig.ClearRoot,
				MultisigConfigPDA,
				ConfigSignersPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()

			require.NoError(t, err)

			result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			require.NotNil(t, result)

			// get config and validate
			var configAccount mcm.MultisigConfig
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA, config.DefaultCommitment, &configAccount)
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
			utils.AssertClosedAccount(ctx, t, solanaGoClient, ConfigSignersPDA, config.DefaultCommitment)
		})
	})
	// todo: negative tests
}
