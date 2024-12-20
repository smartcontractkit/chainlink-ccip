package contracts

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"reflect"
	"slices"
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

	t.Run("setup:funding", func(t *testing.T) {
		utils.FundAccounts(ctx, []solana.PrivateKey{admin, user}, solanaGoClient, t)
	})

	t.Run("mcm:general test cases", func(t *testing.T) {
		// mcm name
		testMsigName := config.TestMsigNamePaddedBuffer

		// test mcm pdas
		multisigConfigPDA := McmConfigAddress(testMsigName)
		rootMetadataPDA := RootMetadataAddress(testMsigName)
		expiringRootAndOpCountPDA := ExpiringRootAndOpCountAddress(testMsigName)
		configSignersPDA := McmConfigSignersAddress(testMsigName)
		msigSignerPDA := McmSignerAddress(testMsigName)

		// fund the signer pda
		fundPDAIx := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), msigSignerPDA).Build()
		result := utils.SendAndConfirm(ctx, t, solanaGoClient,
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
				testMsigName,
				multisigConfigPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				config.McmProgram,
				programData.Address,
				rootMetadataPDA,
				expiringRootAndOpCountPDA,
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
				testMsigName,
				multisigConfigPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
				config.McmProgram,
				programData.Address,
				rootMetadataPDA,
				expiringRootAndOpCountPDA,
			).ValidateAndBuild()
			require.NoError(t, initErr)
			utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

			// get config and validate
			var configAccount mcm.MultisigConfig
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
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
				testMsigName,
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
					testMsigName,
					parsedTotalSigners,
					multisigConfigPDA,
					configSignersPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, isErr)
				ixs = append(ixs, initSignersIx)

				appendSignersIxs, asErr := AppendSignersIxs(signerAddresses, testMsigName, multisigConfigPDA, configSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
				require.NoError(t, asErr)
				ixs = append(ixs, appendSignersIxs...)

				finalizeSignersIx, fsErr := mcm.NewFinalizeSignersInstruction(
					testMsigName,
					multisigConfigPDA,
					configSignersPDA,
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, fsErr)
				ixs = append(ixs, finalizeSignersIx)

				for _, ix := range ixs {
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				var cfgSignersAccount mcm.ConfigSigners
				queryErr := utils.GetAccountDataBorshInto(ctx, solanaGoClient, configSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
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
				multisigConfigPDA,
				configSignersPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()

			require.NoError(t, configErr)

			result := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			require.NotNil(t, result)

			// get config and validate
			var configAccount mcm.MultisigConfig
			configErr = utils.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
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

		t.Run("mcm set_root happy path", func(t *testing.T) {
			for i, op := range stupProgramTestMcmOps {
				node := mcmsUtils.McmOpNode{
					Nonce:             uint64(i),
					Multisig:          multisigConfigPDA,
					To:                config.ExternalCpiStubProgram,
					Data:              op.Data,
					RemainingAccounts: op.RemainingAccounts,
				}
				opNodes = append(opNodes, node)
			}
			validUntil := uint32(0xffffffff)

			rootValidationData, rvErr := CreateMcmRootData(
				McmRootInput{
					Multisig:             multisigConfigPDA,
					Operations:           opNodes,
					PreOpCount:           0,
					PostOpCount:          uint64(len(opNodes)),
					ValidUntil:           validUntil,
					OverridePreviousRoot: false,
				},
			)
			require.NoError(t, rvErr)
			signaturesPDA := RootSignaturesAddress(testMsigName, rootValidationData.Root, validUntil)

			t.Run("preload signatures", func(t *testing.T) {
				signers, getSignerErr := eth.GetEvmSigners(signerPrivateKeys)
				require.NoError(t, getSignerErr, "Failed to get signers")

				signatures, sigsErr := BulkSignOnMsgHash(signers, rootValidationData.EthMsgHash)
				require.NoError(t, sigsErr)

				parsedTotalSigs, pErr := mcmsUtils.SafeToUint8(len(signatures))
				require.NoError(t, pErr)

				initSigsIx, isErr := mcm.NewInitSignaturesInstruction(
					testMsigName,
					rootValidationData.Root,
					validUntil,
					parsedTotalSigs,
					signaturesPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, isErr)
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initSigsIx}, admin, config.DefaultCommitment)

				appendSigsIxs, asErr := AppendSignaturesIxs(signatures, testMsigName, rootValidationData.Root, validUntil, signaturesPDA, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
				require.NoError(t, asErr)

				// partially register signatures
				for _, ix := range appendSigsIxs[:3] {
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				// clear uploaded signatures(this closes the account)
				clearIx, cErr := mcm.NewClearSignaturesInstruction(
					testMsigName,
					rootValidationData.Root,
					validUntil,
					signaturesPDA,
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, cErr)

				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{clearIx}, admin, config.DefaultCommitment)
				utils.AssertClosedAccount(ctx, t, solanaGoClient, signaturesPDA, config.DefaultCommitment)

				reInitSigsIx, rIsErr := mcm.NewInitSignaturesInstruction(
					testMsigName,
					rootValidationData.Root,
					validUntil,
					parsedTotalSigs,
					signaturesPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()

				require.NoError(t, rIsErr)
				utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{reInitSigsIx}, admin, config.DefaultCommitment)

				// register all signatures again
				for _, ix := range appendSigsIxs {
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				}

				finalizeSigsIx, fsErr := mcm.NewFinalizeSignaturesInstruction(
					testMsigName,
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
				testMsigName,
				rootValidationData.Root,
				validUntil,
				rootValidationData.Metadata,
				rootValidationData.MetadataProof,
				signaturesPDA,
				rootMetadataPDA,
				SeenSignedHashesAddress(testMsigName, rootValidationData.Root, validUntil),
				expiringRootAndOpCountPDA,
				multisigConfigPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, setRootIxErr)

			tx := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{newIx}, admin, config.DefaultCommitment, utils.AddComputeUnitLimit(1_400_000))
			require.NotNil(t, tx)

			parsedLogs := utils.ParseLogMessages(tx.Meta.LogMessages,
				[]utils.EventMapping{
					utils.EventMappingFor[NewRoot]("NewRoot"),
				},
			)
			event := parsedLogs[0].EventData[0].Data.(*NewRoot)
			require.Equal(t, rootValidationData.Root, event.Root)
			require.Equal(t, validUntil, event.ValidUntil)
			require.Equal(t, rootValidationData.Metadata.ChainId, event.MetadataChainID)
			require.Equal(t, multisigConfigPDA, event.MetadataMultisig)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, event.MetadataPreOpCount)
			require.Equal(t, rootValidationData.Metadata.PostOpCount, event.MetadataPostOpCount)
			require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, event.MetadataOverridePreviousRoot)

			var newRootAndOpCount mcm.ExpiringRootAndOpCount

			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, expiringRootAndOpCountPDA, config.DefaultCommitment, &newRootAndOpCount)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, rootValidationData.Root, newRootAndOpCount.Root)
			require.Equal(t, validUntil, newRootAndOpCount.ValidUntil)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootAndOpCount.OpCount)

			// get config and validate
			var newRootMetadata mcm.RootMetadata
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, rootMetadataPDA, config.DefaultCommitment, &newRootMetadata)
			require.NoError(t, err, "failed to get account info")

			require.Equal(t, rootValidationData.Metadata.ChainId, newRootMetadata.ChainId)
			require.Equal(t, rootValidationData.Metadata.Multisig, newRootMetadata.Multisig)
			require.Equal(t, rootValidationData.Metadata.PreOpCount, newRootMetadata.PreOpCount)
			require.Equal(t, rootValidationData.Metadata.PostOpCount, newRootMetadata.PostOpCount)
			require.Equal(t, rootValidationData.Metadata.OverridePreviousRoot, newRootMetadata.OverridePreviousRoot)
		})

		t.Run("mcm execute happy path", func(t *testing.T) {
			for i, op := range opNodes {
				proofs, proofsErr := op.Proofs()
				require.NoError(t, proofsErr, "Failed to getting op proof")

				ix := mcm.NewExecuteInstruction(
					testMsigName,
					config.TestChainID,
					op.Nonce,
					op.Data,
					proofs,

					multisigConfigPDA,
					rootMetadataPDA,
					expiringRootAndOpCountPDA,
					config.ExternalCpiStubProgram,
					McmSignerAddress(testMsigName),
					admin.PublicKey(),
				)
				// append remaining accounts
				ix.AccountMetaSlice = append(ix.AccountMetaSlice, op.RemainingAccounts...)

				vIx, vIxErr := ix.ValidateAndBuild()
				require.NoError(t, vIxErr)

				tx := utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{vIx}, admin, config.DefaultCommitment)
				require.NotNil(t, tx.Meta)
				require.Nil(t, tx.Meta.Err, fmt.Sprintf("tx failed with: %+v", tx.Meta))
				parsedInstructions := utils.ParseLogMessages(tx.Meta.LogMessages,
					[]utils.EventMapping{
						utils.EventMappingFor[OpExecuted]("OpExecuted"),
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
			err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, config.StubAccountPDA, config.DefaultCommitment, &stubAccountValue)
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
			modifySigs   func(*[]mcm.Signature, *McmRootData)
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
				modifySigs: func(sigs *[]mcm.Signature, _ *McmRootData) {
					slices.Reverse(*sigs)
				},
			},
			{
				name:         "should fail set_root when signatures don't meet group quorum",
				errorMsg:     "Error Code: " + mcm.InsufficientSigners_McmError.String(),
				failureStage: SetRoot,
				modifySigs: func(sigs *[]mcm.Signature, _ *McmRootData) {
					*sigs = (*sigs)[:1] // only keep first signature
				},
			},
			{
				name:         "when message hash is different from the one used to sign",
				errorMsg:     "Error Code: " + mcm.InvalidSigner_McmError.String(),
				failureStage: SetRoot,
				modifySigs: func(sigs *[]mcm.Signature, _ *McmRootData) {
					// same signers
					signers, signerErr := eth.GetEvmSigners(config.SignerPrivateKeys)
					require.NoError(t, signerErr)
					// but different signatures(wrong eth hash)
					// secp256k1_recover_from recovers a valid but different address --> invalidSigner
					signatures, sigErr := BulkSignOnMsgHash(signers, bytes.Repeat([]byte{1}, 32))
					require.NoError(t, sigErr)
					*sigs = signatures
				},
			},
			{
				name:         "invalid signature should fail ECDSA recovery",
				errorMsg:     "Error Code: " + mcm.FailedEcdsaRecover_McmError.String(),
				failureStage: SetRoot,
				modifySigs: func(sigs *[]mcm.Signature, _ *McmRootData) {
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
				modifySigs: func(sigs *[]mcm.Signature, rootData *McmRootData) {
					wrongPrivateKeys, err := eth.GenerateEthPrivateKeys(len(*sigs))
					require.NoError(t, err)
					wrongSigners, err := eth.GetEvmSigners(wrongPrivateKeys)
					require.NoError(t, err)
					signatures, err := BulkSignOnMsgHash(wrongSigners, rootData.EthMsgHash)
					require.NoError(t, err)
					*sigs = signatures
				},
			},
		}

		for i, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// use different msig accounts per test
				testMsigName, err := mcmsUtils.PadString32(fmt.Sprintf("fail_sig_validation_test_%d", i))
				require.NoError(t, err)

				// test scoped mcm pdas
				multisigConfigPDA := McmConfigAddress(testMsigName)
				multisigSignerPDA := McmSignerAddress(testMsigName)
				rootMetadataPDA := RootMetadataAddress(testMsigName)
				expiringRootAndOpCountPDA := ExpiringRootAndOpCountAddress(testMsigName)
				configSignersPDA := McmConfigSignersAddress(testMsigName)

				// fund the signer pda
				fundPDAIx, err := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), multisigSignerPDA).ValidateAndBuild()
				require.NoError(t, err)
				result := utils.SendAndConfirm(ctx, t, solanaGoClient,
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
						testMsigName,
						multisigConfigPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
						config.McmProgram,
						programData.Address,
						rootMetadataPDA,
						expiringRootAndOpCountPDA,
					).ValidateAndBuild()
					require.NoError(t, initErr)
					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					// get config and validate
					var configAccount mcm.MultisigConfig
					err = utils.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
					require.NoError(t, err, "failed to get account info")

					require.Equal(t, config.TestChainID, configAccount.ChainId)
					require.Equal(t, admin.PublicKey(), configAccount.Owner)
				})

				mcmConfig, configErr := mcmsUtils.NewValidMcmConfig(
					testMsigName,
					config.SignerPrivateKeys,
					config.SignerGroups,
					config.GroupQuorums,
					config.GroupParents,
					config.ClearRoot,
				)
				require.NoError(t, configErr)

				t.Run("setup: load signers and set_config", func(t *testing.T) {
					ixs := make([]solana.Instruction, 0)

					parsedTotalSigners, pErr := mcmsUtils.SafeToUint8(len(mcmConfig.SignerAddresses))
					require.NoError(t, pErr)
					initSignersIx, isErr := mcm.NewInitSignersInstruction(
						testMsigName,
						parsedTotalSigners,
						multisigConfigPDA,
						configSignersPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()

					require.NoError(t, isErr)
					ixs = append(ixs, initSignersIx)

					appendSignersIxs, asErr := AppendSignersIxs(mcmConfig.SignerAddresses, testMsigName, multisigConfigPDA, configSignersPDA, admin.PublicKey(), config.MaxAppendSignerBatchSize)
					require.NoError(t, asErr)
					ixs = append(ixs, appendSignersIxs...)

					finalizeSignersIx, fsErr := mcm.NewFinalizeSignersInstruction(
						testMsigName,
						multisigConfigPDA,
						configSignersPDA,
						admin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, fsErr)
					ixs = append(ixs, finalizeSignersIx)

					for _, ix := range ixs {
						utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
					}

					var cfgSignersAccount mcm.ConfigSigners
					queryErr := utils.GetAccountDataBorshInto(ctx, solanaGoClient, configSignersPDA, config.DefaultCommitment, &cfgSignersAccount)
					require.NoError(t, queryErr, "failed to get account info")

					require.Equal(t, true, cfgSignersAccount.IsFinalized)

					// check if the addresses are registered correctly
					for i, signer := range cfgSignersAccount.SignerAddresses {
						require.Equal(t, mcmConfig.SignerAddresses[i], signer)
					}

					// set config
					ix, configErr := mcm.NewSetConfigInstruction(
						mcmConfig.MultisigName,
						mcmConfig.SignerGroups,
						mcmConfig.GroupQuorums,
						mcmConfig.GroupParents,
						mcmConfig.ClearRoot,
						multisigConfigPDA,
						configSignersPDA,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()

					require.NoError(t, configErr)

					utils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					// get config and validate
					var configAccount mcm.MultisigConfig
					configErr = utils.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
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
				node, err := IxToMcmTestOpNode(multisigConfigPDA, multisigSignerPDA, stubProgramIx, 0)
				require.NoError(t, err)

				validUntil := uint32(0xffffffff)

				// this will be used to generate proof on mcm::execute
				ops := []mcmsUtils.McmOpNode{
					node,
				}

				rootValidationData, rvErr := CreateMcmRootData(
					McmRootInput{
						Multisig:             multisigConfigPDA,
						Operations:           ops,
						PreOpCount:           0,
						PostOpCount:          1,
						ValidUntil:           validUntil,
						OverridePreviousRoot: false,
					},
				)
				require.NoError(t, rvErr)
				signaturesPDA := RootSignaturesAddress(testMsigName, rootValidationData.Root, validUntil)

				signers, getSignerErr := eth.GetEvmSigners(config.SignerPrivateKeys)
				require.NoError(t, getSignerErr)
				signatures, err := BulkSignOnMsgHash(signers, rootValidationData.EthMsgHash)
				require.NoError(t, err)

				if tt.modifySigs != nil {
					tt.modifySigs(&signatures, &rootValidationData)
				}

				parsedTotalSigs, err := mcmsUtils.SafeToUint8(len(signatures))
				require.NoError(t, err)

				initSigsIx, err := mcm.NewInitSignaturesInstruction(
					testMsigName,
					rootValidationData.Root,
					validUntil,
					parsedTotalSigs,
					signaturesPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				txs = append(txs, TxWithStage{Instructions: []solana.Instruction{initSigsIx}, Stage: InitSignatures})

				appendSigsIxs, asErr := AppendSignaturesIxs(signatures, testMsigName, rootValidationData.Root, validUntil, signaturesPDA, admin.PublicKey(), config.MaxAppendSignatureBatchSize)
				require.NoError(t, asErr)

				// one tx is enough since we only have 5 signers
				txs = append(txs, TxWithStage{Instructions: appendSigsIxs, Stage: AppendSignatures})

				finalizeSigsIx, err := mcm.NewFinalizeSignaturesInstruction(
					testMsigName,
					rootValidationData.Root,
					validUntil,
					signaturesPDA,
					admin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				txs = append(txs, TxWithStage{Instructions: []solana.Instruction{finalizeSigsIx}, Stage: FinalizeSignatures})

				setRootIx, err := mcm.NewSetRootInstruction(
					testMsigName,
					rootValidationData.Root,
					validUntil,
					rootValidationData.Metadata,
					rootValidationData.MetadataProof,
					signaturesPDA,
					rootMetadataPDA,
					SeenSignedHashesAddress(testMsigName, rootValidationData.Root, validUntil),
					expiringRootAndOpCountPDA,
					multisigConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				// set compute budget for signature verification
				cuIx, err := computebudget.NewSetComputeUnitLimitInstruction(uint32(1_400_000)).ValidateAndBuild()
				require.NoError(t, err)

				txs = append(txs, TxWithStage{Instructions: []solana.Instruction{cuIx, setRootIx}, Stage: SetRoot})

				// here only one op exists
				opNode := ops[0]
				proofs, proofsErr := opNode.Proofs()
				require.NoError(t, proofsErr, "Failed to getting op proof")

				executeIx := mcm.NewExecuteInstruction(
					testMsigName,
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
						result := utils.SendAndFailWith(ctx, t, solanaGoClient,
							tx.Instructions,
							admin,
							rpc.CommitmentConfirmed,
							[]string{tt.errorMsg},
						)
						require.NotNil(t, result)
						break
					}

					// all other instructions should succeed
					utils.SendAndConfirm(ctx, t, solanaGoClient,
						tx.Instructions,
						admin,
						config.DefaultCommitment,
					)
				}
			})
		}
	})
}