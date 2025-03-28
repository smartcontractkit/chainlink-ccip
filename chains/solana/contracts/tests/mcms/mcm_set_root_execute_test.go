package contracts

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/external_program_cpi_stub"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/mcms"
)

type TestMcmOperation struct {
	Data              []byte
	ExpectedMethod    string
	ExpectedLogSubstr string
	RemainingAccounts []*solana.AccountMeta
	CheckExpectations func(instruction *common.AnchorInstruction) error
}

func TestMcmSetRootAndExecute(t *testing.T) {
	t.Parallel()
	mcm.SetProgramID(config.McmProgram)
	external_program_cpi_stub.SetProgramID(config.ExternalCpiStubProgram) // testing program

	ctx := tests.Context(t)

	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	user, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)

	t.Run("setup:funding", func(t *testing.T) {
		testutils.FundAccounts(ctx, []solana.PrivateKey{admin, user}, solanaGoClient, t)
	})

	t.Run("mcm:general test cases", func(t *testing.T) {
		// mcm name
		testMsigID := config.TestMsigID

		// test mcm pdas
		multisigConfigPDA := mcms.GetConfigPDA(testMsigID)
		rootMetadataPDA := mcms.GetRootMetadataPDA(testMsigID)
		expiringRootAndOpCountPDA := mcms.GetExpiringRootAndOpCountPDA(testMsigID)
		configSignersPDA := mcms.GetConfigSignersPDA(testMsigID)
		msigSignerPDA := mcms.GetSignerPDA(testMsigID)

		// fund the signer pda
		fundPDAIx := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), msigSignerPDA).Build()
		result := testutils.SendAndConfirm(ctx, t, solanaGoClient,
			[]solana.Instruction{fundPDAIx},
			admin, config.DefaultCommitment)
		require.NotNil(t, result)

		// helper to inject anchor discriminator into instruction data
		// NOTE: if ix is built with anchor-go we don't need it
		getAnchorInstructionData := func(method string, data []byte) []byte {
			discriminator := sha256.Sum256([]byte("global:" + method))
			return append(discriminator[:8], data...)
		}

		// NOTE: this list of operations is methods for testing program,
		// contracts/programs/external_program_cpi_stub
		// the other way to construct mcmTestOp is using IxToMcmTestOpNode
		var stupProgramTestMcmOps = []TestMcmOperation{
			{
				ExpectedMethod:    "Initialize",
				Data:              getAnchorInstructionData("initialize", nil),
				ExpectedLogSubstr: "Called `initialize`",
				RemainingAccounts: []*solana.AccountMeta{
					{
						PublicKey:  config.StubAccountPDA,
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PublicKey:  mcms.GetSignerPDA(config.TestMsigID),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PublicKey:  solana.SystemProgramID,
						IsSigner:   false,
						IsWritable: false,
					},
				},
			},
			{
				ExpectedMethod:    "Empty",
				Data:              getAnchorInstructionData("empty", nil),
				ExpectedLogSubstr: "Called `empty`",
			},
			{
				ExpectedMethod:    "U8InstructionData",
				Data:              getAnchorInstructionData("u8_instruction_data", []byte{123}),
				ExpectedLogSubstr: "Called `u8_instruction_data` Context { program_id: 2zZwzyptLqwFJFEFxjPvrdhiGpH9pJ3MfrrmZX6NTKxm, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data 123",
			},
			{
				ExpectedMethod:    "StructInstructionData",
				Data:              getAnchorInstructionData("struct_instruction_data", []byte{234}),
				ExpectedLogSubstr: "Called `struct_instruction_data` Context { program_id: 2zZwzyptLqwFJFEFxjPvrdhiGpH9pJ3MfrrmZX6NTKxm, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data Value { value: 234 }",
			},
			{
				ExpectedMethod: "AccountRead",
				Data:           getAnchorInstructionData("account_read", nil),
				RemainingAccounts: []*solana.AccountMeta{
					{
						PublicKey:  config.StubAccountPDA,
						IsSigner:   false,
						IsWritable: false,
					},
				},
				ExpectedLogSubstr: "Called `account_read`",
				CheckExpectations: func(instruction *common.AnchorInstruction) error {
					if !strings.Contains(instruction.Logs[0], "value: 1") {
						return fmt.Errorf("expected log to contain 'value: 1', got: %s", instruction.Logs[0])
					}
					return nil
				},
			},
			{
				ExpectedMethod: "AccountMut",
				Data:           getAnchorInstructionData("account_mut", nil),
				RemainingAccounts: []*solana.AccountMeta{
					{
						PublicKey:  config.StubAccountPDA,
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PublicKey:  mcms.GetSignerPDA(config.TestMsigID),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PublicKey:  solana.SystemProgramID,
						IsSigner:   false,
						IsWritable: false,
					},
				},
				ExpectedLogSubstr: "Called `account_mut`",
				CheckExpectations: func(instruction *common.AnchorInstruction) error {
					if !strings.Contains(instruction.Logs[0], "is_writable: true") {
						return fmt.Errorf("expected log to contain 'is_writable: true', got: %s", instruction.Logs[0])
					}
					return nil
				},
			},
		}

		t.Run("should not be able to init program from non-deployer", func(t *testing.T) {
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

			ix, initErr := mcm.NewInitializeInstruction(
				config.TestChainID,
				testMsigID,
				multisigConfigPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				config.McmProgram,
				programData.Address,
				rootMetadataPDA,
				expiringRootAndOpCountPDA,
			).ValidateAndBuild()
			require.NoError(t, initErr)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + mcms.UnauthorizedMcmError.String()})
			require.NotNil(t, result)
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

			ix, initErr := mcm.NewInitializeInstruction(
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
			require.NoError(t, initErr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

			// get config and validate
			var configAccount mcm.MultisigConfig
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, config.TestChainID, configAccount.ChainId)
			require.Equal(t, admin.PublicKey(), configAccount.Owner)
		})

		numSigners := 50
		signerPrivateKeys, err := eth.GenerateEthPrivateKeys(numSigners)

		t.Run("mcm:set_config:success", func(t *testing.T) {
			require.NoError(t, err)

			signerGroups := make([]byte, numSigners)
			for i := 0; i < len(signerGroups); i++ {
				signerGroups[i] = byte(i % 5)
			}

			// just use simple config for now
			groupQuorums := []uint8{1, 1, 1, 1, 1}
			groupParents := []uint8{0, 0, 0, 2, 0}

			mcmConfig, configErr := mcms.NewValidMcmConfig(
				testMsigID,
				signerPrivateKeys,
				signerGroups,
				groupQuorums,
				groupParents,
				config.ClearRoot,
			)
			require.NoError(t, configErr)

			signerAddresses := mcmConfig.SignerAddresses

			t.Run("mcm:preload signers", func(t *testing.T) {
				preloadIxs, pierr := mcms.GetPreloadSignersIxs(signerAddresses, testMsigID, multisigConfigPDA, configSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
				require.NoError(t, pierr)

				for _, ix := range preloadIxs {
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				var cfgSignersAccount mcm.ConfigSigners
				queryErr := common.GetAccountDataBorshInto(ctx, solanaGoClient, configSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
				require.NoError(t, queryErr, "failed to get account info")

				require.Equal(t, true, cfgSignersAccount.IsFinalized)

				// check if the addresses are registered correctly
				for i, signer := range cfgSignersAccount.SignerAddresses {
					require.Equal(t, signerAddresses[i], signer)
				}
			})

			// set config
			ix, configErr := mcm.NewSetConfigInstruction(
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

			require.NoError(t, configErr)

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			require.NotNil(t, result)

			// get config and validate
			var configAccount mcm.MultisigConfig
			configErr = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
			require.NoError(t, configErr, "failed to get account info")

			require.Equal(t, config.TestChainID, configAccount.ChainId)
			require.Equal(t, reflect.DeepEqual(configAccount.GroupParents, mcmConfig.GroupParents), true)
			require.Equal(t, reflect.DeepEqual(configAccount.GroupQuorums, mcmConfig.GroupQuorums), true)

			// check if the McmSigner struct is correct
			for i, signer := range configAccount.Signers {
				require.Equal(t, signer.EvmAddress, mcmConfig.SignerAddresses[i])
				require.Equal(t, signer.Index, uint8(i))
				require.Equal(t, signer.Group, mcmConfig.SignerGroups[i])
			}
		})

		t.Run("mcm set_root", func(t *testing.T) {
			var opNodes []mcms.McmOpNode

			for i, op := range stupProgramTestMcmOps {
				node := mcms.McmOpNode{
					Nonce:             uint64(i),
					Multisig:          multisigConfigPDA,
					To:                config.ExternalCpiStubProgram,
					Data:              op.Data,
					RemainingAccounts: op.RemainingAccounts,
				}
				opNodes = append(opNodes, node)
			}
			validUntil := uint32(0xffffffff)

			rootValidationData, rvErr := mcms.CreateMcmRootData(
				mcms.McmRootInput{
					Multisig:             multisigConfigPDA,
					Operations:           opNodes,
					PreOpCount:           0,
					PostOpCount:          uint64(len(opNodes)),
					ValidUntil:           validUntil,
					OverridePreviousRoot: false,
				},
			)
			require.NoError(t, rvErr)
			signaturesPDA := mcms.GetRootSignaturesPDA(testMsigID, rootValidationData.Root, validUntil, admin.PublicKey())

			t.Run("preload signatures", func(t *testing.T) {
				signers, getSignerErr := eth.GetEvmSigners(signerPrivateKeys)
				require.NoError(t, getSignerErr, "Failed to get signers")

				signatures, sigsErr := mcms.BulkSignOnMsgHash(signers, rootValidationData.EthMsgHash)
				require.NoError(t, sigsErr)

				//nolint:gosec
				parsedTotalSigs := uint8(len(signatures))
				initSigsIx, isErr := mcm.NewInitSignaturesInstruction(
					testMsigID,
					rootValidationData.Root,
					validUntil,
					parsedTotalSigs,
					signaturesPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, isErr)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initSigsIx}, admin, config.DefaultCommitment)

				// append signature hijacking scenario
				invalidAppendSigsIxs, iaserr := mcms.GetAppendSignaturesIxs(signatures, testMsigID, rootValidationData.Root, validUntil, user.PublicKey(), config.MaxAppendSignatureBatchSize)
				require.NoError(t, iaserr)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{invalidAppendSigsIxs[0]}, user, config.DefaultCommitment, []string{"AnchorError caused by account: signatures. Error Code: " + common.AccountNotInitialized_AnchorError.String()})

				// now try with valid authority
				appendSigsIxs, asErr := mcms.GetAppendSignaturesIxs(signatures, testMsigID, rootValidationData.Root, validUntil, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
				require.NoError(t, asErr)

				// partially register signatures
				for _, ix := range appendSigsIxs[:3] {
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				// clear uploaded signatures(this closes the account)
				clearIx, cErr := mcm.NewClearSignaturesInstruction(
					testMsigID,
					rootValidationData.Root,
					validUntil,
					signaturesPDA,
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, cErr)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{clearIx}, admin, config.DefaultCommitment)
				testutils.AssertClosedAccount(ctx, t, solanaGoClient, signaturesPDA, config.DefaultCommitment)

				// preload again
				preloadIxs, plerr := mcms.GetMcmPreloadSignaturesIxs(signatures, testMsigID, rootValidationData.Root, validUntil, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
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

			// legitimate root_signatures PDA but wrong authority
			invalidIx1, ivsrerr := mcm.NewSetRootInstruction(
				testMsigID,
				rootValidationData.Root,
				validUntil,
				rootValidationData.Metadata,
				rootValidationData.MetadataProof,
				signaturesPDA, // legitimate root_signatures PDA address, but with wrong authority -> seed constraint violation
				rootMetadataPDA,
				mcms.GetSeenSignedHashesPDA(testMsigID, rootValidationData.Root, validUntil),
				expiringRootAndOpCountPDA,
				multisigConfigPDA,
				user.PublicKey(), // invalid signer
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, ivsrerr)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{invalidIx1}, user, config.DefaultCommitment, []string{"Error Code: " + common.ConstraintSeeds_AnchorError.String()})

			// root_signatures PDA with matching authority, but not the same authority with preloading instructions
			// so in this case, the root_signatures PDA is not initialized
			invalidIx2, ivsrerr := mcm.NewSetRootInstruction(
				testMsigID,
				rootValidationData.Root,
				validUntil,
				rootValidationData.Metadata,
				rootValidationData.MetadataProof,
				mcms.GetRootSignaturesPDA(testMsigID, rootValidationData.Root, validUntil, user.PublicKey()),
				rootMetadataPDA,
				mcms.GetSeenSignedHashesPDA(testMsigID, rootValidationData.Root, validUntil),
				expiringRootAndOpCountPDA,
				multisigConfigPDA,
				user.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, ivsrerr)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{invalidIx2}, user, config.DefaultCommitment, []string{"AnchorError caused by account: root_signatures. Error Code: " + common.AccountNotInitialized_AnchorError.String()})

			newIx, setRootIxErr := mcm.NewSetRootInstruction(
				testMsigID,
				rootValidationData.Root,
				validUntil,
				rootValidationData.Metadata,
				rootValidationData.MetadataProof,
				signaturesPDA,
				rootMetadataPDA,
				mcms.GetSeenSignedHashesPDA(testMsigID, rootValidationData.Root, validUntil),
				expiringRootAndOpCountPDA,
				multisigConfigPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, setRootIxErr)

			cu := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{newIx}, admin, config.DefaultCommitment)
			tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{newIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(cu))
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
			require.Equal(t, multisigConfigPDA, event.MetadataMultisig)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, event.MetadataPreOpCount)
			require.Equal(t, rootValidationData.Metadata.PostOpCount, event.MetadataPostOpCount)
			require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, event.MetadataOverridePreviousRoot)

			var newRootAndOpCount mcm.ExpiringRootAndOpCount

			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, expiringRootAndOpCountPDA, config.DefaultCommitment, &newRootAndOpCount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, rootValidationData.Root, newRootAndOpCount.Root)
			require.Equal(t, validUntil, newRootAndOpCount.ValidUntil)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootAndOpCount.OpCount)

			// get config and validate
			var newRootMetadata mcm.RootMetadata
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootMetadataPDA, config.DefaultCommitment, &newRootMetadata)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, rootValidationData.Metadata.ChainId, newRootMetadata.ChainId)
			require.Equal(t, rootValidationData.Metadata.Multisig, newRootMetadata.Multisig)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootMetadata.PreOpCount)
			require.Equal(t, rootValidationData.Metadata.PostOpCount, newRootMetadata.PostOpCount)
			require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, newRootMetadata.OverridePreviousRoot)
		})

		var opNodes []mcms.McmOpNode

		t.Run("change config scenario, clear root and set a new root", func(t *testing.T) {
			// keep current op count
			var prevRootAndOpCount mcm.ExpiringRootAndOpCount
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, expiringRootAndOpCountPDA, config.DefaultCommitment, &prevRootAndOpCount)
			require.NoError(t, err, "failed to get account info")
			currentOpCount := prevRootAndOpCount.OpCount

			// new config with clear_root
			newMcmConfig, configErr := mcms.NewValidMcmConfig(
				testMsigID,
				config.SignerPrivateKeys,
				config.SignerGroups,
				config.GroupQuorums,
				config.GroupParents,
				true, // clear_root
			)
			require.NoError(t, configErr)

			signerAddresses := newMcmConfig.SignerAddresses

			t.Run("preload signers", func(t *testing.T) {
				preloadIxs, pierr := mcms.GetPreloadSignersIxs(signerAddresses, testMsigID, multisigConfigPDA, configSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
				require.NoError(t, pierr)

				for _, ix := range preloadIxs {
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				var cfgSignersAccount mcm.ConfigSigners
				queryErr := common.GetAccountDataBorshInto(ctx, solanaGoClient, configSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
				require.NoError(t, queryErr, "failed to get account info")
				require.Equal(t, true, cfgSignersAccount.IsFinalized)
				// check if the addresses are registered correctly
				for i, signer := range cfgSignersAccount.SignerAddresses {
					require.Equal(t, signerAddresses[i], signer)
				}
			})

			t.Run("set_config with clear_root", func(t *testing.T) {
				ix, configErr := mcm.NewSetConfigInstruction(
					newMcmConfig.MultisigID,
					newMcmConfig.SignerGroups,
					newMcmConfig.GroupQuorums,
					newMcmConfig.GroupParents,
					newMcmConfig.ClearRoot,
					multisigConfigPDA,
					configSignersPDA,
					rootMetadataPDA,
					expiringRootAndOpCountPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, configErr)

				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				require.NotNil(t, tx)

				parsedLogs := common.ParseLogMessages(tx.Meta.LogMessages,
					[]common.EventMapping{
						common.EventMappingFor[mcms.ConfigSet]("ConfigSet"),
					},
				)

				event := parsedLogs[0].EventData[0].Data.(*mcms.ConfigSet)
				require.Equal(t, newMcmConfig.GroupParents, event.GroupParents)
				require.Equal(t, newMcmConfig.GroupQuorums, event.GroupQuorums)
				require.Equal(t, true, event.IsRootCleared)
				for i, signer := range event.Signers {
					require.Equal(t, newMcmConfig.SignerAddresses[i], signer.EvmAddress)
					require.Equal(t, uint8(i), signer.Index)
					require.Equal(t, newMcmConfig.SignerGroups[i], signer.Group)
				}

				// get config and validate
				var configAccount mcm.MultisigConfig
				configErr = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
				require.NoError(t, configErr, "failed to get account info")

				require.Equal(t, config.TestChainID, configAccount.ChainId)
				require.Equal(t, reflect.DeepEqual(configAccount.GroupParents, newMcmConfig.GroupParents), true)
				require.Equal(t, reflect.DeepEqual(configAccount.GroupQuorums, newMcmConfig.GroupQuorums), true)

				// check if the McmSigner struct is correct
				for i, signer := range configAccount.Signers {
					require.Equal(t, signer.EvmAddress, newMcmConfig.SignerAddresses[i])
					require.Equal(t, signer.Index, uint8(i))
					require.Equal(t, signer.Group, newMcmConfig.SignerGroups[i])
				}

				// get root, metadata and validate
				var clearedRootAndOpCount mcm.ExpiringRootAndOpCount

				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, expiringRootAndOpCountPDA, config.DefaultCommitment, &clearedRootAndOpCount)
				require.NoError(t, err, "failed to get account info")

				require.Equal(t, config.McmEmptyRoot, clearedRootAndOpCount.Root)
				require.Equal(t, config.McmEmptyTimestamp, clearedRootAndOpCount.ValidUntil)
				require.Equal(t, currentOpCount, clearedRootAndOpCount.OpCount)

				var clearedRootMetadata mcm.RootMetadata
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootMetadataPDA, config.DefaultCommitment, &clearedRootMetadata)
				require.NoError(t, err, "failed to get account info")

				require.Equal(t, configAccount.ChainId, clearedRootMetadata.ChainId) // reset to config.chainid
				require.Equal(t, multisigConfigPDA, clearedRootMetadata.Multisig)    // should be the same multisig
				require.Equal(t, currentOpCount, clearedRootMetadata.PreOpCount)     // preserve op count
				require.Equal(t, currentOpCount, clearedRootMetadata.PostOpCount)    // preserve op count
				require.Equal(t, true, clearedRootMetadata.OverridePreviousRoot)     // should be true
			})

			t.Run("mcm set_root on new config", func(t *testing.T) {
				for i, op := range stupProgramTestMcmOps {
					node := mcms.McmOpNode{
						Nonce:             uint64(i),
						Multisig:          multisigConfigPDA,
						To:                config.ExternalCpiStubProgram,
						Data:              op.Data,
						RemainingAccounts: op.RemainingAccounts,
					}
					opNodes = append(opNodes, node)
				}
				// should be different timestamp from previous root(seen_signed_hash)
				validUntil := uint32(time.Now().Add(1 * time.Hour).Unix())

				rootValidationData, rvErr := mcms.CreateMcmRootData(
					mcms.McmRootInput{
						Multisig:             multisigConfigPDA,
						Operations:           opNodes,
						PreOpCount:           0,
						PostOpCount:          uint64(len(opNodes)),
						ValidUntil:           validUntil,
						OverridePreviousRoot: false,
					},
				)
				require.NoError(t, rvErr)
				signaturesPDA := mcms.GetRootSignaturesPDA(testMsigID, rootValidationData.Root, validUntil, admin.PublicKey())

				t.Run("preload signatures", func(t *testing.T) {
					signers, getSignerErr := eth.GetEvmSigners(config.SignerPrivateKeys)
					require.NoError(t, getSignerErr, "Failed to get signers")

					signatures, sigsErr := mcms.BulkSignOnMsgHash(signers, rootValidationData.EthMsgHash)
					require.NoError(t, sigsErr)

					preloadIxs, plerr := mcms.GetMcmPreloadSignaturesIxs(signatures, testMsigID, rootValidationData.Root, validUntil, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
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

				newIx, setRootIxErr := mcm.NewSetRootInstruction(
					testMsigID,
					rootValidationData.Root,
					validUntil,
					rootValidationData.Metadata,
					rootValidationData.MetadataProof,
					signaturesPDA,
					rootMetadataPDA,
					mcms.GetSeenSignedHashesPDA(testMsigID, rootValidationData.Root, validUntil),
					expiringRootAndOpCountPDA,
					multisigConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, setRootIxErr)

				cu := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{newIx}, admin, config.DefaultCommitment)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{newIx}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(cu))
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
				require.Equal(t, multisigConfigPDA, event.MetadataMultisig)
				require.Equal(t, rootValidationData.Metadata.PreOpCount, event.MetadataPreOpCount)
				require.Equal(t, rootValidationData.Metadata.PostOpCount, event.MetadataPostOpCount)
				require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, event.MetadataOverridePreviousRoot)

				var newRootAndOpCount mcm.ExpiringRootAndOpCount

				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, expiringRootAndOpCountPDA, config.DefaultCommitment, &newRootAndOpCount)
				require.NoError(t, err, "failed to get account info")

				require.Equal(t, rootValidationData.Root, newRootAndOpCount.Root)
				require.Equal(t, validUntil, newRootAndOpCount.ValidUntil)
				require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootAndOpCount.OpCount)

				// get config and validate
				var newRootMetadata mcm.RootMetadata
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootMetadataPDA, config.DefaultCommitment, &newRootMetadata)
				require.NoError(t, err, "failed to get account info")

				require.Equal(t, rootValidationData.Metadata.ChainId, newRootMetadata.ChainId)
				require.Equal(t, rootValidationData.Metadata.Multisig, newRootMetadata.Multisig)
				require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootMetadata.PreOpCount)
				require.Equal(t, rootValidationData.Metadata.PostOpCount, newRootMetadata.PostOpCount)
				require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, newRootMetadata.OverridePreviousRoot)
			})
		})

		t.Run("mcm execute happy path", func(t *testing.T) {
			for i, op := range opNodes {
				proofs, proofsErr := op.Proofs()
				require.NoError(t, proofsErr, "Failed to getting op proof")

				ix := mcm.NewExecuteInstruction(
					testMsigID,
					config.TestChainID,
					op.Nonce,
					op.Data,
					proofs,

					multisigConfigPDA,
					rootMetadataPDA,
					expiringRootAndOpCountPDA,
					config.ExternalCpiStubProgram,
					mcms.GetSignerPDA(testMsigID),
					admin.PublicKey(),
				)
				// append remaining accounts
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op.RemainingAccounts...)

				vIx, vIxErr := ix.ValidateAndBuild()
				require.NoError(t, vIxErr)

				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment)
				require.NotNil(t, tx.Meta)
				require.Nil(t, tx.Meta.Err, fmt.Sprintf("tx failed with: %+v", tx.Meta))
				parsedInstructions := common.ParseLogMessages(tx.Meta.LogMessages,
					[]common.EventMapping{
						common.EventMappingFor[mcms.OpExecuted]("OpExecuted"),
					},
				)

				require.Len(t, parsedInstructions, 1, "Expected 1 top-level instruction")

				topLevelInstruction := parsedInstructions[0]
				require.Equal(t, "Execute", topLevelInstruction.Name, "Top-level instruction should be Execute")
				require.Equal(t, config.McmProgram.String(), topLevelInstruction.ProgramID, "Top-level instruction should be executed by MCM program")

				require.Len(t, topLevelInstruction.InnerCalls, 1, "Expected 1 inner call")
				innerCall := topLevelInstruction.InnerCalls[0]

				require.Equal(t, stupProgramTestMcmOps[i].ExpectedMethod, innerCall.Name, "Inner call name should match the expected method")
				require.Equal(t, config.ExternalCpiStubProgram.String(), innerCall.ProgramID, "Inner call should be executed by external CPI stub program")

				require.NotEmpty(t, innerCall.Logs, "Inner call should have logs")
				require.Contains(t, innerCall.Logs[0], stupProgramTestMcmOps[i].ExpectedLogSubstr, "Inner call log should contain expected substring")

				if stupProgramTestMcmOps[i].CheckExpectations != nil {
					vIxErr = stupProgramTestMcmOps[i].CheckExpectations(innerCall)
					require.NoError(t, vIxErr, "Custom expectations check failed")
				}
			}

			var stubAccountValue external_program_cpi_stub.Value
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.StubAccountPDA, config.DefaultCommitment, &stubAccountValue)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, uint8(2), stubAccountValue.Value)
		})
	})

	t.Run("mcm set_root validations", func(t *testing.T) {
		type TestStage int

		const (
			InitSignatures TestStage = iota
			AppendSignatures
			FinalizeSignatures
			SetRoot
			Execute
		)

		type TxWithStage struct {
			Instructions []solana.Instruction
			Stage        TestStage
		}

		tests := []struct {
			name         string
			errorMsg     string
			failureStage TestStage
			modifyTxs    func(*[]TxWithStage)
			modifySigs   func(*[]mcm.Signature, *mcms.McmRootData)
		}{
			{
				name:         "should not be able to initialize signatures more than one time ",
				errorMsg:     "already in use Program 11111111111111111111111111111111 failed", // account creation failure from system program
				failureStage: InitSignatures,
				modifyTxs: func(txs *[]TxWithStage) {
					// find the index of the first initSignatures instruction
					var initSigIndex int
					for i, tx := range *txs {
						if tx.Stage == InitSignatures {
							initSigIndex = i
							break
						}
					}
					(*txs)[initSigIndex] = TxWithStage{
						Instructions: []solana.Instruction{(*txs)[initSigIndex].Instructions[0], (*txs)[initSigIndex].Instructions[0]},
						Stage:        InitSignatures,
					}
				},
			},
			{
				name:         "should not be able to finalize signatures before having all signatures",
				errorMsg:     "Error Code: " + mcm.SignatureCountMismatch_McmError.String(),
				failureStage: FinalizeSignatures,
				modifyTxs: func(txs *[]TxWithStage) {
					finalizeSigsIdx := -1
					appendSigsIdx := -1
					for i, tx := range *txs {
						if tx.Stage == FinalizeSignatures {
							finalizeSigsIdx = i
						}
						if tx.Stage == AppendSignatures {
							appendSigsIdx = i
						}
					}
					// move finalize before append in place
					if finalizeSigsIdx > appendSigsIdx {
						(*txs)[finalizeSigsIdx], (*txs)[appendSigsIdx] = (*txs)[appendSigsIdx], (*txs)[finalizeSigsIdx]
					}
				},
			},
			{
				name:         "attempt to append more signatures than defined in total",
				errorMsg:     "Error Code: " + mcm.TooManySignatures_McmError.String(),
				failureStage: AppendSignatures,
				modifyTxs: func(txs *[]TxWithStage) {
					appendSigsIdx := -1
					for i, tx := range *txs {
						if tx.Stage == AppendSignatures {
							appendSigsIdx = i
						}
					}
					(*txs)[appendSigsIdx] = TxWithStage{
						Instructions: []solana.Instruction{(*txs)[appendSigsIdx].Instructions[0], (*txs)[appendSigsIdx].Instructions[0]},
						Stage:        AppendSignatures,
					}
				},
			},
			{
				name:         "signatures not in ascending order",
				errorMsg:     "Error Code: " + mcm.SignersAddressesMustBeStrictlyIncreasing_McmError.String(),
				failureStage: SetRoot,
				modifySigs: func(sigs *[]mcm.Signature, _ *mcms.McmRootData) {
					slices.Reverse(*sigs)
				},
			},
			{
				name:         "should fail set_root when signatures don't meet group quorum",
				errorMsg:     "Error Code: " + mcm.InsufficientSigners_McmError.String(),
				failureStage: SetRoot,
				modifySigs: func(sigs *[]mcm.Signature, _ *mcms.McmRootData) {
					*sigs = (*sigs)[:1] // only keep first signature
				},
			},
			{
				name:         "when message hash is different from the one used to sign",
				errorMsg:     "Error Code: " + mcm.InvalidSigner_McmError.String(),
				failureStage: SetRoot,
				modifySigs: func(sigs *[]mcm.Signature, _ *mcms.McmRootData) {
					// same signers
					signers, signerErr := eth.GetEvmSigners(config.SignerPrivateKeys)
					require.NoError(t, signerErr)
					// but different signatures(wrong eth hash)
					// secp256k1_recover_from recovers a valid but different address --> invalidSigner
					signatures, sigErr := mcms.BulkSignOnMsgHash(signers, bytes.Repeat([]byte{1}, 32))
					require.NoError(t, sigErr)
					*sigs = signatures
				},
			},
			{
				name:         "invalid signature should fail ECDSA recovery",
				errorMsg:     "Error Code: " + mcm.FailedEcdsaRecover_McmError.String(),
				failureStage: SetRoot,
				modifySigs: func(sigs *[]mcm.Signature, _ *mcms.McmRootData) {
					invalidSig := (*sigs)[0]
					// corrupt V
					invalidSig.V = 26
					newSigs := make([]mcm.Signature, len(*sigs))
					for i := range newSigs {
						newSigs[i] = invalidSig
					}
					*sigs = newSigs
				},
			},
			{
				name:         "signatures from unauthorized signers should fail",
				errorMsg:     "Error Code: " + mcm.InvalidSigner_McmError.String(),
				failureStage: SetRoot,
				modifySigs: func(sigs *[]mcm.Signature, rootData *mcms.McmRootData) {
					wrongPrivateKeys, err := eth.GenerateEthPrivateKeys(len(*sigs))
					require.NoError(t, err)
					wrongSigners, err := eth.GetEvmSigners(wrongPrivateKeys)
					require.NoError(t, err)
					signatures, err := mcms.BulkSignOnMsgHash(wrongSigners, rootData.EthMsgHash)
					require.NoError(t, err)
					*sigs = signatures
				},
			},
		}

		for i, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// use different msig accounts per test
				testMsigID, err := mcms.PadString32(fmt.Sprintf("fail_sig_validation_test_%d", i))
				require.NoError(t, err)

				// test scoped mcm pdas
				multisigConfigPDA := mcms.GetConfigPDA(testMsigID)
				multisigSignerPDA := mcms.GetSignerPDA(testMsigID)
				rootMetadataPDA := mcms.GetRootMetadataPDA(testMsigID)
				expiringRootAndOpCountPDA := mcms.GetExpiringRootAndOpCountPDA(testMsigID)
				configSignersPDA := mcms.GetConfigSignersPDA(testMsigID)

				// fund the signer pda
				fundPDAIx, err := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), multisigSignerPDA).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient,
					[]solana.Instruction{fundPDAIx},
					admin, config.DefaultCommitment)
				require.NotNil(t, result)

				t.Run("setup:initialize mcm", func(t *testing.T) {
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

					ix, initErr := mcm.NewInitializeInstruction(
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
					require.NoError(t, initErr)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					// get config and validate
					var configAccount mcm.MultisigConfig
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
					require.NoError(t, err, "failed to get account info")

					require.Equal(t, config.TestChainID, configAccount.ChainId)
					require.Equal(t, admin.PublicKey(), configAccount.Owner)
				})

				mcmConfig, configErr := mcms.NewValidMcmConfig(
					testMsigID,
					config.SignerPrivateKeys,
					config.SignerGroups,
					config.GroupQuorums,
					config.GroupParents,
					config.ClearRoot,
				)
				require.NoError(t, configErr)

				t.Run("setup: load signers and set_config", func(t *testing.T) {
					preloadIxs, pierr := mcms.GetPreloadSignersIxs(mcmConfig.SignerAddresses, testMsigID, multisigConfigPDA, configSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
					require.NoError(t, pierr)

					for _, ix := range preloadIxs {
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
					}

					var cfgSignersAccount mcm.ConfigSigners
					queryErr := common.GetAccountDataBorshInto(ctx, solanaGoClient, configSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
					require.NoError(t, queryErr, "failed to get account info")

					require.Equal(t, true, cfgSignersAccount.IsFinalized)

					// check if the addresses are registered correctly
					for i, signer := range cfgSignersAccount.SignerAddresses {
						require.Equal(t, mcmConfig.SignerAddresses[i], signer)
					}

					// set config
					ix, configErr := mcm.NewSetConfigInstruction(
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

					require.NoError(t, configErr)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					// get config and validate
					var configAccount mcm.MultisigConfig
					configErr = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
					require.NoError(t, configErr, "failed to get account info")

					require.Equal(t, config.TestChainID, configAccount.ChainId)
					require.Equal(t, reflect.DeepEqual(configAccount.GroupParents, mcmConfig.GroupParents), true)
					require.Equal(t, reflect.DeepEqual(configAccount.GroupQuorums, mcmConfig.GroupQuorums), true)

					// check if the McmSigner struct is correct
					for i, signer := range configAccount.Signers {
						require.Equal(t, signer.EvmAddress, mcmConfig.SignerAddresses[i])
						require.Equal(t, signer.Index, uint8(i))
						require.Equal(t, signer.Group, mcmConfig.SignerGroups[i])
					}
				})

				var txs []TxWithStage

				// use simple program for testing
				stubProgramIx, err := external_program_cpi_stub.NewEmptyInstruction().ValidateAndBuild()
				node, err := mcms.IxToMcmTestOpNode(multisigConfigPDA, multisigSignerPDA, stubProgramIx, 0)
				require.NoError(t, err)

				validUntil := uint32(0xffffffff)

				// this will be used to generate proof on mcm::execute
				ops := []mcms.McmOpNode{
					node,
				}

				rootValidationData, rvErr := mcms.CreateMcmRootData(
					mcms.McmRootInput{
						Multisig:             multisigConfigPDA,
						Operations:           ops,
						PreOpCount:           0,
						PostOpCount:          1,
						ValidUntil:           validUntil,
						OverridePreviousRoot: false,
					},
				)
				require.NoError(t, rvErr)
				signaturesPDA := mcms.GetRootSignaturesPDA(testMsigID, rootValidationData.Root, validUntil, admin.PublicKey())

				signers, getSignerErr := eth.GetEvmSigners(config.SignerPrivateKeys)
				require.NoError(t, getSignerErr)
				signatures, err := mcms.BulkSignOnMsgHash(signers, rootValidationData.EthMsgHash)
				require.NoError(t, err)

				if tt.modifySigs != nil {
					tt.modifySigs(&signatures, &rootValidationData)
				}
				//nolint:gosec
				parsedTotalSigs := uint8(len(signatures))
				initSigsIx, err := mcm.NewInitSignaturesInstruction(
					testMsigID,
					rootValidationData.Root,
					validUntil,
					parsedTotalSigs,
					signaturesPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				txs = append(txs, TxWithStage{Instructions: []solana.Instruction{initSigsIx}, Stage: InitSignatures})

				appendSigsIxs, asErr := mcms.GetAppendSignaturesIxs(signatures, testMsigID, rootValidationData.Root, validUntil, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
				require.NoError(t, asErr)

				// one tx is enough since we only have 5 signers
				txs = append(txs, TxWithStage{Instructions: appendSigsIxs, Stage: AppendSignatures})

				finalizeSigsIx, err := mcm.NewFinalizeSignaturesInstruction(
					testMsigID,
					rootValidationData.Root,
					validUntil,
					signaturesPDA,
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				txs = append(txs, TxWithStage{Instructions: []solana.Instruction{finalizeSigsIx}, Stage: FinalizeSignatures})

				setRootIx, err := mcm.NewSetRootInstruction(
					testMsigID,
					rootValidationData.Root,
					validUntil,
					rootValidationData.Metadata,
					rootValidationData.MetadataProof,
					signaturesPDA,
					rootMetadataPDA,
					mcms.GetSeenSignedHashesPDA(testMsigID, rootValidationData.Root, validUntil),
					expiringRootAndOpCountPDA,
					multisigConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				// set compute budget for signature verification
				cuIx, err := computebudget.NewSetComputeUnitLimitInstruction(uint32(computebudget.MAX_COMPUTE_UNIT_LIMIT)).ValidateAndBuild()
				require.NoError(t, err)

				txs = append(txs, TxWithStage{Instructions: []solana.Instruction{cuIx, setRootIx}, Stage: SetRoot})

				// here only one op exists
				opNode := ops[0]
				proofs, proofsErr := opNode.Proofs()
				require.NoError(t, proofsErr, "Failed to getting op proof")

				executeIx := mcm.NewExecuteInstruction(
					testMsigID,
					config.TestChainID,
					opNode.Nonce,
					opNode.Data,
					proofs,

					multisigConfigPDA,
					rootMetadataPDA,
					expiringRootAndOpCountPDA,
					config.ExternalCpiStubProgram,
					multisigSignerPDA,
					admin.PublicKey(),
				)
				// append remaining accounts
				executeIx.AccountMetaSlice = append(executeIx.AccountMetaSlice, opNode.RemainingAccounts...)

				vIx, vIxErr := executeIx.ValidateAndBuild()
				require.NoError(t, vIxErr)

				txs = append(txs, TxWithStage{Instructions: []solana.Instruction{vIx}, Stage: Execute})

				if tt.modifyTxs != nil {
					tt.modifyTxs(&txs)
				}

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
