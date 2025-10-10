package contracts

import (
	"testing"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/mcms"
)

func TestMcmSelfOwnership(t *testing.T) {
	t.Parallel()
	mcm.SetProgramID(config.McmProgram)

	ctx := tests.Context(t)

	// create initial admin
	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)

	// mcm instance setup
	testMsigID := config.TestMsigID
	multisigConfigPDA := mcms.GetConfigPDA(testMsigID)
	rootMetadataPDA := mcms.GetRootMetadataPDA(testMsigID)
	expiringRootAndOpCountPDA := mcms.GetExpiringRootAndOpCountPDA(testMsigID)
	configSignersPDA := mcms.GetConfigSignersPDA(testMsigID)
	msigSignerPDA := mcms.GetSignerPDA(testMsigID)

	t.Run("setup:funding", func(t *testing.T) {
		testutils.FundAccounts(ctx, []solana.PrivateKey{admin}, solanaGoClient, t)

		// fund the signer PDA
		fundPDAIx := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), msigSignerPDA).Build()
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{fundPDAIx}, admin, config.DefaultCommitment)
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

		// initialize MCM
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

		// verify initial owner
		var configAccount mcm.MultisigConfig
		err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
		require.NoError(t, err)
		require.Equal(t, admin.PublicKey(), configAccount.Owner)
	})

	t.Run("setup:signers_and_config", func(t *testing.T) {
		// get predefined test signers from config
		evmSigners, err := eth.GetEvmSigners(config.SignerPrivateKeys)
		require.NoError(t, err)

		// convert to address format for MCM
		signers := make([][20]byte, len(evmSigners))
		for i, signer := range evmSigners {
			copy(signers[i][:], signer.Address[:])
		}

		// initialize signers
		initSignersIx, err := mcm.NewInitSignersInstruction(
			testMsigID,
			uint8(len(signers)),
			multisigConfigPDA,
			configSignersPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initSignersIx}, admin, config.DefaultCommitment)

		// append signers
		appendSignersIx, err := mcm.NewAppendSignersInstruction(
			testMsigID,
			signers,
			multisigConfigPDA,
			configSignersPDA,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{appendSignersIx}, admin, config.DefaultCommitment)

		// finalize signers
		finalizeSignersIx, err := mcm.NewFinalizeSignersInstruction(
			testMsigID,
			multisigConfigPDA,
			configSignersPDA,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{finalizeSignersIx}, admin, config.DefaultCommitment)

		// set config using predefined test configuration
		var groupQuorums [32]uint8
		var groupParents [32]uint8
		copy(groupQuorums[:], config.GroupQuorums)
		copy(groupParents[:], config.GroupParents)

		setConfigIx, err := mcm.NewSetConfigInstruction(
			testMsigID,
			config.SignerGroups,
			groupQuorums,
			groupParents,
			false, // don't clear root
			multisigConfigPDA,
			configSignersPDA,
			rootMetadataPDA,
			expiringRootAndOpCountPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{setConfigIx}, admin, config.DefaultCommitment)
	})

	t.Run("test:self_ownership_via_execute", func(t *testing.T) {
		// step 1: transfer ownership to the multisig signer PDA
		transferOwnershipIx, err := mcm.NewTransferOwnershipInstruction(
			testMsigID,
			msigSignerPDA, // propose the multisig signer PDA as new owner
			multisigConfigPDA,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{transferOwnershipIx}, admin, config.DefaultCommitment)

		// verify proposed owner is set
		var configAccount mcm.MultisigConfig
		err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
		require.NoError(t, err)
		require.Equal(t, msigSignerPDA, configAccount.ProposedOwner)
		require.Equal(t, admin.PublicKey(), configAccount.Owner) // still the old owner

		// step 2: create accept_ownership operation to be executed via MCM
		acceptOwnershipIx, err := mcm.NewAcceptOwnershipInstruction(
			testMsigID,
			multisigConfigPDA,
			msigSignerPDA, // the multisig signer PDA accepts ownership
		).ValidateAndBuild()
		require.NoError(t, err)

		// convert instruction to MCM operation
		acceptOwnershipNode, err := mcms.IxToMcmTestOpNode(
			multisigConfigPDA,
			msigSignerPDA,
			acceptOwnershipIx,
			0, // nonce
		)
		require.NoError(t, err)

		// create operations list
		ops := []mcms.McmOpNode{acceptOwnershipNode}

		// create root validation data
		validUntil := uint32(time.Now().Unix() + 3600) // 1 hour from now
		rootValidationData, err := mcms.CreateMcmRootData(
			mcms.McmRootInput{
				Multisig:             multisigConfigPDA,
				Operations:           ops,
				PreOpCount:           0,
				PostOpCount:          1,
				ValidUntil:           validUntil,
				OverridePreviousRoot: false,
			},
		)
		require.NoError(t, err)

		// update the operation node with the tree information
		acceptOwnershipNode = ops[0] // get the updated node with tree information

		// get signers and sign the root
		signers, err := eth.GetEvmSigners(config.SignerPrivateKeys)
		require.NoError(t, err)
		signatures, err := mcms.BulkSignOnMsgHash(signers, rootValidationData.EthMsgHash)
		require.NoError(t, err)

		// preload signatures
		rootSignaturesPDA := mcms.GetRootSignaturesPDA(testMsigID, rootValidationData.Root, validUntil, admin.PublicKey())
		preloadIxs, err := mcms.GetMcmPreloadSignaturesIxs(
			signatures,
			testMsigID,
			rootValidationData.Root,
			validUntil,
			admin.PublicKey(),
			config.MaxAppendSignatureBatchSize,
		)
		require.NoError(t, err)

		// execute preload instructions
		for _, ix := range preloadIxs {
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
		}

		// set root
		seenSignedHashesPDA := mcms.GetSeenSignedHashesPDA(testMsigID, rootValidationData.Root, validUntil)
		setRootIx, err := mcm.NewSetRootInstruction(
			testMsigID,
			rootValidationData.Root,
			validUntil,
			rootValidationData.Metadata,
			rootValidationData.MetadataProof,
			rootSignaturesPDA,
			rootMetadataPDA,
			seenSignedHashesPDA,
			expiringRootAndOpCountPDA,
			multisigConfigPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{setRootIx}, admin, config.DefaultCommitment)

		// step 3: execute the accept_ownership operation via MCM
		proofs, err := acceptOwnershipNode.Proofs()
		require.NoError(t, err)

		executeIx := mcm.NewExecuteInstruction(
			testMsigID,
			config.TestChainID,
			acceptOwnershipNode.Nonce,
			acceptOwnershipNode.Data,
			proofs,
			multisigConfigPDA,
			rootMetadataPDA,
			expiringRootAndOpCountPDA,
			config.McmProgram, // CPI back to MCM program
			msigSignerPDA,
			admin.PublicKey(),
		)

		// add remaining accounts for the CPI
		executeIx.AccountMetaSlice = append(executeIx.AccountMetaSlice, acceptOwnershipNode.RemainingAccounts...)

		// build and execute the instruction
		builtExecuteIx, err := executeIx.ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{builtExecuteIx}, admin, config.DefaultCommitment)

		// step 4: verify the ownership transfer
		err = common.GetAccountDataBorshInto(ctx, solanaGoClient, multisigConfigPDA, config.DefaultCommitment, &configAccount)
		require.NoError(t, err)

		// with reload() fix, owner should be msigSignerPDA
		t.Logf("Final owner: %s", configAccount.Owner.String())
		t.Logf("Expected owner (msigSignerPDA): %s", msigSignerPDA.String())
		t.Logf("Old owner (admin): %s", admin.PublicKey().String())

		// check if the ownership transfer worked
		if configAccount.Owner.Equals(msigSignerPDA) {
			t.Logf("SUCCESS: MCM is now owned by its own signer PDA")
			// proposed owner should be reset to default
			require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
		} else if configAccount.Owner.Equals(admin.PublicKey()) {
			t.Logf("FAILURE: MCM is still owned by admin - demonstrates account reloading bug")
			t.Logf("without reload(), CPI-modified account state gets overwritten by stale outer context")
			t.Fail() // mark test as failed but continue to show the issue
		} else {
			t.Fatalf("Unexpected owner: %s", configAccount.Owner.String())
		}
	})
}
