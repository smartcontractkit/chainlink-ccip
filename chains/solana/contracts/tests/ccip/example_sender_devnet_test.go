package contracts

import (
	"context"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	ag_binary "github.com/gagliardetto/binary"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/example_ccip_sender"
	solcommon "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

func TestSenderDevnet(t *testing.T) {
	devnetInfo, err := getDevnetInfo()
	require.NoError(t, err)

	ctx := tests.Context(t)
	client := rpc.New(devnetInfo.RPC)

	admin := solana.PrivateKey(devnetInfo.PrivateKeys.Admin)
	require.True(t, admin.IsValid())

	offrampAddress, err := solana.PublicKeyFromBase58(devnetInfo.Offramp)
	require.NoError(t, err)

	offrampPDAs, err := getOfframpPDAs(offrampAddress)
	require.NoError(t, err)

	var referenceAddresses ccip_offramp.ReferenceAddresses

	t.Run("Read Reference Addresses", func(t *testing.T) {
		require.NoError(t, solcommon.GetAccountDataBorshInto(ctx, client, offrampPDAs.referenceAddresses, rpc.CommitmentConfirmed, &referenceAddresses))
		fmt.Printf("Reference Addresses: %+v\n", referenceAddresses)
	})

	exampleSenderAddress := solana.MustPublicKeyFromBase58(devnetInfo.ExampleSender)
	example_ccip_sender.SetProgramID(exampleSenderAddress)

	senderPDAs, err := getSenderPDAs(exampleSenderAddress)
	require.NoError(t, err)

	t.Run("Initialize Sender", func(t *testing.T) {
		t.Skip()
		ix, err := example_ccip_sender.NewInitializeInstruction(
			referenceAddresses.Router,
			senderPDAs.state,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
		fmt.Printf("Result: %s\n", result.Meta.LogMessages)
	})

	t.Run("Initialize Sepolia Dest Chain", func(t *testing.T) {
		// t.Skip()

		selector := devnetInfo.ChainSelectors.Sepolia

		addy, err := createEVMSpecificRecipientBytes("0x59bb22C50900291f0428d712eE48529eb4B159cC")
		require.NoError(t, err)

		ix, err := example_ccip_sender.NewInitChainConfigInstruction(
			selector,
			addy,
			MakeSVM2EVMExtraArgsV2(0, true), // extra args bytes
			senderPDAs.state,
			senderPDAs.FindChainConfig(t, selector),
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
		fmt.Printf("Result: %s\n", result.Meta.LogMessages)
	})

	t.Run("Update Sepolia Dest Chain", func(t *testing.T) {
		t.Skip()

		selector := devnetInfo.ChainSelectors.Sepolia

		addy, err := createEVMSpecificRecipientBytes("0x59bb22C50900291f0428d712eE48529eb4B159cC")
		require.NoError(t, err)

		ix, err := example_ccip_sender.NewUpdateChainConfigInstruction(
			selector,
			addy,
			MakeSVM2EVMExtraArgsV2(0, true), // extra args bytes
			senderPDAs.state,
			senderPDAs.FindChainConfig(t, selector),
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
		fmt.Printf("Result: %s\n", result.Meta.LogMessages)
	})

	t.Run("Transfer ownership", func(t *testing.T) {
		t.Skip()

		ix, err := example_ccip_sender.NewTransferOwnershipInstruction(
			solana.MustPublicKeyFromBase58("<REPLACE_WITH_NEW_OWNER>"),
			senderPDAs.state,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
		fmt.Printf("Result: %s\n", result.Meta.LogMessages)
	})

	t.Run("Accept ownership", func(t *testing.T) {
		t.Skip()

		ix, err := example_ccip_sender.NewAcceptOwnershipInstruction(
			senderPDAs.state,
			admin.PublicKey(), // make sure the admin is the new owner
		).ValidateAndBuild()
		require.NoError(t, err)
		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
		fmt.Printf("Result: %s\n", result.Meta.LogMessages)
	})

	t.Run("Verify Sepolia Dest Chain Config", func(t *testing.T) {
		t.Skip()
		selector := devnetInfo.ChainSelectors.Sepolia
		expectedRecipient := common.HexToAddress("0x59bb22C50900291f0428d712eE48529eb4B159cC") // Use the correct EVM address
		expectedAllowOOO := true                                                               // Or false if that's expected

		verifyChainConfig(ctx, t, client, senderPDAs, selector, expectedRecipient, expectedAllowOOO)
	})
}

func MakeSVM2EVMExtraArgsV2(gasLimit uint64, allowOOO bool) []byte {
	var evmExtraArgsV2Tag = hexutil.MustDecode("0x181dcf10")

	// extra args is the tag followed by the gas limit and allowOOO borsh-encoded.
	var extraArgs []byte
	extraArgs = append(extraArgs, evmExtraArgsV2Tag...)

	gasLimitBytes := make([]byte, 8) // uint64 requires 8 bytes
	binary.LittleEndian.PutUint64(gasLimitBytes, gasLimit)
	gasLimitBytes = common.RightPadBytes(gasLimitBytes, 16) // little-endian borsh-encoding of u128

	// borsh-encode allowOOO
	var allowOOOBytes []byte
	if allowOOO {
		allowOOOBytes = append(allowOOOBytes, 1)
	} else {
		allowOOOBytes = append(allowOOOBytes, 0)
	}

	extraArgs = append(extraArgs, gasLimitBytes...)
	extraArgs = append(extraArgs, allowOOOBytes...)
	return extraArgs
}

// verifyChainConfig fetches and verifies the on-chain configuration for a specific destination chain.
// verifyChainConfig fetches and verifies the on-chain configuration for a specific destination chain.
func verifyChainConfig(
	ctx context.Context,
	t *testing.T,
	client *rpc.Client,
	senderPDAs senderPDAs, // Assuming senderPDAs type holds methods like FindChainConfig
	selector uint64,
	expectedRecipient common.Address, // Still expecting EVM for Sepolia recipient
	expectedAllowOOO bool,
) {
	t.Helper() // Mark this as a test helper function

	// Find the chain config PDA
	chainConfigPDA := senderPDAs.FindChainConfig(t, selector)
	fmt.Printf("\n--- Verifying ChainConfig ---\n")
	fmt.Printf("Selector: %d\n", selector)
	fmt.Printf("Expected Recipient: %s\n", expectedRecipient.Hex())
	fmt.Printf("Expected AllowOOO: %t\n", expectedAllowOOO)
	fmt.Printf("ChainConfig PDA: %s\n", chainConfigPDA.String())

	// Get account info
	accountInfo, err := client.GetAccountInfoWithOpts(ctx, chainConfigPDA, &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentConfirmed,
	})
	require.NoError(t, err, "Failed to get account info for ChainConfig PDA")
	require.NotNil(t, accountInfo, "AccountInfo is nil for ChainConfig PDA")
	require.NotNil(t, accountInfo.Value, "AccountInfo.Value is nil for ChainConfig PDA")

	accountData := accountInfo.Value.Data.GetBinary()
	fmt.Printf("Raw data length: %d bytes\n", len(accountData))
	require.GreaterOrEqual(t, len(accountData), 8, "Account data too short to contain discriminator")

	// Define a struct matching the expected on-chain layout (after discriminator)
	type chainConfigData struct {
		Recipient      []byte
		ExtraArgsBytes []byte
	}
	var chainConfig chainConfigData

	// Decode the account data (excluding the 8-byte discriminator)
	decoder := ag_binary.NewBorshDecoder(accountData[8:])
	err = decoder.Decode(&chainConfig)
	require.NoError(t, err, "Failed to decode ChainConfig data")

	// --- Verification ---
	fmt.Printf("Decoded Recipient length: %d\n", len(chainConfig.Recipient))
	require.Equal(t, expectedRecipient.Bytes(), chainConfig.Recipient, "Recipient mismatch")
	fmt.Printf("✅ Recipient verified: %s\n", common.BytesToAddress(chainConfig.Recipient).Hex())

	fmt.Printf("Decoded ExtraArgsBytes length: %d\n", len(chainConfig.ExtraArgsBytes))
	// Generate expected args using the SVM format function
	expectedExtraArgs := MakeSVM2EVMExtraArgsV2(0, expectedAllowOOO)
	require.Equal(t, expectedExtraArgs, chainConfig.ExtraArgsBytes, "ExtraArgsBytes mismatch (full comparison)")

	// --- Explicitly decode allowOOO based on SVM format ---
	expectedSvmExtraArgsLen := 4 + 16 + 1 // Tag + Gas (u128 borsh) + AllowOOO (bool borsh)
	require.Equal(t, expectedSvmExtraArgsLen, len(chainConfig.ExtraArgsBytes), "ExtraArgsBytes length mismatch for SVM format")
	// In SVM format, the bool is the last byte
	actualAllowOOO := chainConfig.ExtraArgsBytes[len(chainConfig.ExtraArgsBytes)-1] == 1
	require.Equal(t, expectedAllowOOO, actualAllowOOO, "Decoded allowOOO flag mismatch")

	fmt.Printf("✅ ExtraArgsBytes verified (allowOOO=%t): %x\n", actualAllowOOO, chainConfig.ExtraArgsBytes)
	fmt.Printf("--- Verification Complete ---\n\n")
}

func createEVMSpecificRecipientBytes(recipientStr string) ([]byte, error) {
	if recipientStr == "" {
		return nil, fmt.Errorf("recipient string cannot be empty")
	}

	// Validate and parse the EVM address string
	if !common.IsHexAddress(recipientStr) {
		return nil, fmt.Errorf("invalid EVM address format for recipient: %s", recipientStr)
	}
	evmAddress := common.HexToAddress(recipientStr) // This is [20]byte

	// Pad the 20 bytes to 32 bytes (left-padded with zeros), mimicking the snippet's logic.
	paddedBytes := leftPadBytes(evmAddress.Bytes(), 32) // Using the solcommon alias

	return paddedBytes, nil
}

func leftPadBytes(input []byte, length int) []byte {
	if len(input) >= length {
		return input
	}

	pad := make([]byte, length-len(input))
	return append(pad, input...)
}
