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

func TestMcmMultipleMultisigs(t *testing.T) {
	t.Parallel()
	mcm.SetProgramID(config.McmProgram)

	ctx := tests.Context(t)

	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	solanaGoClient := utils.DeployAllPrograms(t, utils.PathToAnchorConfig, admin)

	// mcm multisig 1
	TestMsigName1, err := mcmsUtils.PadString32("test_mcm_instance_1")
	require.NoError(t, err)
	MultisigConfigPDA1 := McmConfigAddress(TestMsigName1)
	RootMetadataPDA1 := RootMetadataAddress(TestMsigName1)
	ExpiringRootAndOpCountPDA1 := ExpiringRootAndOpCountAddress(TestMsigName1)
	ConfigSignersPDA1 := McmConfigSignersAddress(TestMsigName1)

	// mcm multisig 2
	TestMsigName2, err := mcmsUtils.PadString32("test_mcm_instance_2")
	require.NoError(t, err)
	MultisigConfigPDA2 := McmConfigAddress(TestMsigName2)
	RootMetadataPDA2 := RootMetadataAddress(TestMsigName2)
	ExpiringRootAndOpCountPDA2 := ExpiringRootAndOpCountAddress(TestMsigName2)
	ConfigSignersPDA2 := McmConfigSignersAddress(TestMsigName2)

	t.Run("setup:funding", func(t *testing.T) {
		utils.FundAccounts(ctx, []solana.PrivateKey{admin}, solanaGoClient, t)
	})

	t.Run("setup:test_mcm_instance_1", func(t *testing.T) {
		t.Run("setup:test_mcm_instance_1 init", func(t *testing.T) {
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
				TestMsigName1,
				MultisigConfigPDA1,
				admin.PublicKey(),
				solana.SystemProgramID,
				config.McmProgram,
				programData.Address,
				RootMetadataPDA1,
				ExpiringRootAndOpCountPDA1,
			).ValidateAndBuild()
			require.NoError(t, err)
			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

			// get config and validate
			var configAccount mcm.MultisigConfig
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA1, config.DefaultCommitment, &configAccount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, config.TestChainID, configAccount.ChainId)
			require.Equal(t, admin.PublicKey(), configAccount.Owner)
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
				TestMsigName1,
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
					TestMsigName1,
					parsedTotalSigners,
					MultisigConfigPDA1,
					ConfigSignersPDA1,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, err)
				ixs = append(ixs, initSignersIx)

				appendSignersIxs, err := AppendSignersIxs(signerAddresses, TestMsigName1, MultisigConfigPDA1, ConfigSignersPDA1, admin.PublicKey(), config.MaxAppendSignerBatchSize)
				require.NoError(t, err)
				ixs = append(ixs, appendSignersIxs...)

				finalizeSignersIx, err := mcm.NewFinalizeSignersInstruction(
					TestMsigName1,
					MultisigConfigPDA1,
					ConfigSignersPDA1,
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				ixs = append(ixs, finalizeSignersIx)

				for _, ix := range ixs {
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				var cfgSignersAccount mcm.ConfigSigners
				err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, ConfigSignersPDA1, config.DefaultCommitment, &cfgSignersAccount)
				require.NoError(t, err, "failed to get account info")

				require.Equal(t, true, cfgSignersAccount.IsFinalized)

				// check if the addresses are registered correctly
				for i, signer := range cfgSignersAccount.SignerAddresses {
					require.Equal(t, signerAddresses[i], signer)
				}
			})

			t.Run("success:set_config", func(t *testing.T) {
				ix, err := mcm.NewSetConfigInstruction(
					TestMsigName1,
					mcmConfig.SignerGroups,
					mcmConfig.GroupQuorums,
					mcmConfig.GroupParents,
					mcmConfig.ClearRoot,
					MultisigConfigPDA1,
					ConfigSignersPDA1,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, err)

				result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				require.NotNil(t, result)

				// get config and validate
				var configAccount mcm.MultisigConfig
				err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA1, config.DefaultCommitment, &configAccount)
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
				utils.AssertClosedAccount(ctx, t, solanaGoClient, ConfigSignersPDA1, config.DefaultCommitment)
			})
		})
	})
	t.Run("setup:test_mcm_instance_2", func(t *testing.T) {
		t.Run("setup:test_mcm_instance_2 init", func(t *testing.T) {
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
				TestMsigName2,
				MultisigConfigPDA2,
				admin.PublicKey(),
				solana.SystemProgramID,
				config.McmProgram,
				programData.Address,
				RootMetadataPDA2,
				ExpiringRootAndOpCountPDA2,
			).ValidateAndBuild()
			require.NoError(t, err)
			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

			// get config and validate
			var configAccount mcm.MultisigConfig
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA2, config.DefaultCommitment, &configAccount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, config.TestChainID, configAccount.ChainId)
			require.Equal(t, admin.PublicKey(), configAccount.Owner)
		})

		t.Run("mcm:set_config", func(t *testing.T) {
			numSigners := config.MaxNumSigners
			signerPrivateKeys, err := eth.GenerateEthPrivateKeys(numSigners)
			require.NoError(t, err)

			signerGroups := make([]byte, numSigners)
			for i := 0; i < len(signerGroups); i++ {
				signerGroups[i] = byte(i % 10)
			}

			groupQuorums := []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
			groupParents := []uint8{0, 0, 0, 2, 0, 0, 0, 0, 0, 0}

			mcmConfig, err := mcmsUtils.NewValidMcmConfig(
				TestMsigName2,
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
					TestMsigName2,
					parsedTotalSigners,
					MultisigConfigPDA2,
					ConfigSignersPDA2,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, err)
				ixs = append(ixs, initSignersIx)

				appendSignersIxs, err := AppendSignersIxs(signerAddresses, TestMsigName2, MultisigConfigPDA2, ConfigSignersPDA2, admin.PublicKey(), config.MaxAppendSignerBatchSize)
				require.NoError(t, err)
				ixs = append(ixs, appendSignersIxs...)

				finalizeSignersIx, err := mcm.NewFinalizeSignersInstruction(
					TestMsigName2,
					MultisigConfigPDA2,
					ConfigSignersPDA2,
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				ixs = append(ixs, finalizeSignersIx)

				for _, ix := range ixs {
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				var cfgSignersAccount mcm.ConfigSigners
				err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, ConfigSignersPDA2, config.DefaultCommitment, &cfgSignersAccount)
				require.NoError(t, err, "failed to get account info")

				require.Equal(t, true, cfgSignersAccount.IsFinalized)

				// check if the addresses are registered correctly
				for i, signer := range cfgSignersAccount.SignerAddresses {
					require.Equal(t, signerAddresses[i], signer)
				}
			})

			t.Run("fail:set_config with invalid seeds", func(t *testing.T) {
				ix, err := mcm.NewSetConfigInstruction(
					TestMsigName1,
					mcmConfig.SignerGroups,
					mcmConfig.GroupQuorums,
					mcmConfig.GroupParents,
					mcmConfig.ClearRoot,
					MultisigConfigPDA2,
					ConfigSignersPDA2,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, err)

				result := utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, []string{"Error Code: " + "ConstraintSeeds"})
				require.NotNil(t, result)
			})

			t.Run("success:set_config", func(t *testing.T) {
				ix, err := mcm.NewSetConfigInstruction(
					TestMsigName2,
					mcmConfig.SignerGroups,
					mcmConfig.GroupQuorums,
					mcmConfig.GroupParents,
					mcmConfig.ClearRoot,
					MultisigConfigPDA2,
					ConfigSignersPDA2,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, err)

				result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				require.NotNil(t, result)

				// get config and validate
				var configAccount mcm.MultisigConfig
				err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA2, config.DefaultCommitment, &configAccount)
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
				utils.AssertClosedAccount(ctx, t, solanaGoClient, ConfigSignersPDA2, config.DefaultCommitment)
			})
		})
	})
}
