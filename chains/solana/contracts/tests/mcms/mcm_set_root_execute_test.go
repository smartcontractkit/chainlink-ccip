package contracts

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strings"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/eth"
	mcmsUtils "github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/external_program_cpi_stub"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/mcm"
)

type TestMcmOperation struct {
	Data              []byte
	ExpectedMethod    string
	ExpectedLogSubstr string
	RemainingAccounts []*solana.AccountMeta
	CheckExpectations func(instruction *utils.AnchorInstruction) error
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

	solanaGoClient := utils.DeployAllPrograms(t, utils.PathToAnchorConfig, admin)

	// mcm name
	TestMsigName := config.TestMsigNamePaddedBuffer

	// test mcm pdas
	MultisigConfigPDA := McmConfigAddress(TestMsigName)
	RootMetadataPDA := RootMetadataAddress(TestMsigName)
	ExpiringRootAndOpCountPDA := ExpiringRootAndOpCountAddress(TestMsigName)
	ConfigSignersPDA := McmConfigSignersAddress(TestMsigName)
	MsigSignerPDA := McmSignerAddress(TestMsigName)

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
					PublicKey:  McmSignerAddress(config.TestMsigNamePaddedBuffer),
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
			ExpectedLogSubstr: "Called `u8_instruction_data` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data 123",
		},
		{
			ExpectedMethod:    "StructInstructionData",
			Data:              getAnchorInstructionData("struct_instruction_data", []byte{234}),
			ExpectedLogSubstr: "Called `struct_instruction_data` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data Value { value: 234 }",
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
			CheckExpectations: func(instruction *utils.AnchorInstruction) error {
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
					PublicKey:  McmSignerAddress(config.TestMsigNamePaddedBuffer),
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
			CheckExpectations: func(instruction *utils.AnchorInstruction) error {
				if !strings.Contains(instruction.Logs[0], "is_writable: true") {
					return fmt.Errorf("expected log to contain 'is_writable: true', got: %s", instruction.Logs[0])
				}
				return nil
			},
		},
	}

	t.Run("setup:funding", func(t *testing.T) {
		utils.FundAccounts(ctx, []solana.PrivateKey{admin, user}, solanaGoClient, t)
		// fund the signer pda
		fundPDAIx := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), MsigSignerPDA).Build()
		result := utils.SendAndConfirm(ctx, t, solanaGoClient,
			[]solana.Instruction{fundPDAIx},
			admin, config.DefaultCommitment)
		require.NotNil(t, result)
	})

	t.Run("fail: NOT able to init program from non-deployer user", func(t *testing.T) {
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
			TestMsigName,
			MultisigConfigPDA,
			user.PublicKey(),
			solana.SystemProgramID,
			config.McmProgram,
			programData.Address,
			RootMetadataPDA,
			ExpiringRootAndOpCountPDA,
		).ValidateAndBuild()
		require.NoError(t, initErr)
		result := utils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + UnauthorizedMcmError.String()})
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
			TestMsigName,
			MultisigConfigPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
			config.McmProgram,
			programData.Address,
			RootMetadataPDA,
			ExpiringRootAndOpCountPDA,
		).ValidateAndBuild()
		require.NoError(t, initErr)
		utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

		// get config and validate
		var configAccount mcm.MultisigConfig
		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA, config.DefaultCommitment, &configAccount)
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

		mcmConfig, configErr := mcmsUtils.NewValidMcmConfig(
			TestMsigName,
			signerPrivateKeys,
			signerGroups,
			groupQuorums,
			groupParents,
			config.ClearRoot,
		)
		require.NoError(t, configErr)

		signerAddresses := mcmConfig.SignerAddresses

		t.Run("mcm:preload signers", func(t *testing.T) {
			ixs := make([]solana.Instruction, 0)

			parsedTotalSigners, pErr := mcmsUtils.SafeToUint8(len(signerAddresses))
			require.NoError(t, pErr)
			initSignersIx, isErr := mcm.NewInitSignersInstruction(
				TestMsigName,
				parsedTotalSigners,
				MultisigConfigPDA,
				ConfigSignersPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()

			require.NoError(t, isErr)
			ixs = append(ixs, initSignersIx)

			appendSignersIxs, asErr := AppendSignersIxs(signerAddresses, TestMsigName, MultisigConfigPDA, ConfigSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
			require.NoError(t, asErr)
			ixs = append(ixs, appendSignersIxs...)

			finalizeSignersIx, fsErr := mcm.NewFinalizeSignersInstruction(
				TestMsigName,
				MultisigConfigPDA,
				ConfigSignersPDA,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, fsErr)
			ixs = append(ixs, finalizeSignersIx)

			for _, ix := range ixs {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			var cfgSignersAccount mcm.ConfigSigners
			queryErr := utils.GetAccountDataBorshInto(ctx, solanaGoClient, ConfigSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
			require.NoError(t, queryErr, "failed to get account info")

			require.Equal(t, true, cfgSignersAccount.IsFinalized)

			// check if the addresses are registered correctly
			for i, signer := range cfgSignersAccount.SignerAddresses {
				require.Equal(t, signerAddresses[i], signer)
			}
		})

		// set config
		ix, configErr := mcm.NewSetConfigInstruction(
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

		require.NoError(t, configErr)

		result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
		require.NotNil(t, result)

		// get config and validate
		var configAccount mcm.MultisigConfig
		configErr = utils.GetAccountDataBorshInto(ctx, solanaGoClient, MultisigConfigPDA, config.DefaultCommitment, &configAccount)
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

	var opNodes []mcmsUtils.McmOpNode

	t.Run("mcm:set_root:success", func(t *testing.T) {
		for i, op := range stupProgramTestMcmOps {
			node := mcmsUtils.McmOpNode{
				Nonce:             uint64(i),
				Multisig:          MultisigConfigPDA,
				To:                config.ExternalCpiStubProgram,
				Data:              op.Data,
				RemainingAccounts: op.RemainingAccounts,
			}
			opNodes = append(opNodes, node)
		}
		validUntil := uint32(0xffffffff)

		rootValidationData, rvErr := CreateMcmRootData(
			McmRootInput{
				Multisig:             MultisigConfigPDA,
				Operations:           opNodes,
				PreOpCount:           0,
				PostOpCount:          uint64(len(opNodes)),
				ValidUntil:           validUntil,
				OverridePreviousRoot: false,
			},
		)
		require.NoError(t, rvErr)
		signaturesPDA := RootSignaturesAddress(TestMsigName, rootValidationData.Root, validUntil)

		t.Run("mcm:preload signatures", func(t *testing.T) {
			signers, getSignerErr := eth.GetEvmSigners(signerPrivateKeys)
			require.NoError(t, getSignerErr, "Failed to get signers")

			signatures, sigsErr := BulkSignOnMsgHash(signers, rootValidationData.EthMsgHash)
			require.NoError(t, sigsErr)

			parsedTotalSigs, pErr := mcmsUtils.SafeToUint8(len(signatures))
			require.NoError(t, pErr)

			initSigsIx, isErr := mcm.NewInitSignaturesInstruction(
				TestMsigName,
				rootValidationData.Root,
				validUntil,
				parsedTotalSigs,
				signaturesPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()

			require.NoError(t, isErr)
			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initSigsIx}, admin, config.DefaultCommitment)

			appendSigsIxs, asErr := AppendSignaturesIxs(signatures, TestMsigName, rootValidationData.Root, validUntil, signaturesPDA, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
			require.NoError(t, asErr)

			// partially register signatures
			for _, ix := range appendSigsIxs[:3] {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			// clear uploaded signatures
			clearIx, cErr := mcm.NewClearSignaturesInstruction(
				TestMsigName,
				rootValidationData.Root,
				validUntil,
				signaturesPDA,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, cErr)

			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{clearIx}, admin, config.DefaultCommitment)

			// register all signatures again
			for _, ix := range appendSigsIxs {
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			}

			finalizeSigsIx, fsErr := mcm.NewFinalizeSignaturesInstruction(
				TestMsigName,
				rootValidationData.Root,
				validUntil,
				signaturesPDA,
				admin.PublicKey(),
			).ValidateAndBuild()

			require.NoError(t, fsErr)
			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{finalizeSigsIx}, admin, config.DefaultCommitment)

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

		newIx, setRootIxErr := mcm.NewSetRootInstruction(
			TestMsigName,
			rootValidationData.Root,
			validUntil,
			rootValidationData.Metadata,
			rootValidationData.MetadataProof,
			signaturesPDA,
			RootMetadataPDA,
			SeenSignedHashesAddress(TestMsigName, rootValidationData.Root, validUntil),
			ExpiringRootAndOpCountPDA,
			MultisigConfigPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, setRootIxErr)
		cu := uint32(numSigners * 28_000) //estimated cu per signer
		require.True(t, cu <= 1_400_000, "maximum CU limit exceeded")
		cuIx, cuErr := computebudget.NewSetComputeUnitLimitInstruction(cu).ValidateAndBuild()
		require.NoError(t, cuErr)
		result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{cuIx, newIx}, admin, config.DefaultCommitment)
		require.NotNil(t, result)

		var newRootAndOpCount mcm.ExpiringRootAndOpCount

		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, ExpiringRootAndOpCountPDA, config.DefaultCommitment, &newRootAndOpCount)
		require.NoError(t, err, "failed to get account info")

		require.Equal(t, rootValidationData.Root, newRootAndOpCount.Root)
		require.Equal(t, validUntil, newRootAndOpCount.ValidUntil)
		require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootAndOpCount.OpCount)

		// get config and validate
		var newRootMetadata mcm.RootMetadata
		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, RootMetadataPDA, config.DefaultCommitment, &newRootMetadata)
		require.NoError(t, err, "failed to get account info")

		require.Equal(t, rootValidationData.Metadata.ChainId, newRootMetadata.ChainId)
		require.Equal(t, rootValidationData.Metadata.Multisig, newRootMetadata.Multisig)
		require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootMetadata.PreOpCount)
		require.Equal(t, rootValidationData.Metadata.PostOpCount, newRootMetadata.PostOpCount)
		require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, newRootMetadata.OverridePreviousRoot)
	})

	t.Run("mcm:execute:success", func(t *testing.T) {
		for i, op := range opNodes {
			proofs, proofsErr := op.Proofs()
			require.NoError(t, proofsErr, "Failed to getting op proof")

			ix := mcm.NewExecuteInstruction(
				TestMsigName,
				config.TestChainID,
				op.Nonce,
				op.Data,
				proofs,

				MultisigConfigPDA,
				RootMetadataPDA,
				ExpiringRootAndOpCountPDA,
				config.ExternalCpiStubProgram,
				McmSignerAddress(TestMsigName),
				admin.PublicKey(),
			)
			// append remaining accounts
			ix.AccountMetaSlice = append(ix.AccountMetaSlice, op.RemainingAccounts...)

			vIx, vIxErr := ix.ValidateAndBuild()
			require.NoError(t, vIxErr)

			tx := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment)
			require.NotNil(t, tx.Meta)
			require.Nil(t, tx.Meta.Err, fmt.Sprintf("tx failed with: %+v", tx.Meta))

			parsedInstructions := utils.ParseLogMessages(tx.Meta.LogMessages)

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
		err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, config.StubAccountPDA, config.DefaultCommitment, &stubAccountValue)
		require.NoError(t, err, "failed to get account info")

		require.Equal(t, uint8(2), stubAccountValue.Value)
	})
}
