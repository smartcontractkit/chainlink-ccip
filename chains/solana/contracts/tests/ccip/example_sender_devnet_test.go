package contracts

import (
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/example_ccip_sender"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
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
		require.NoError(t, common.GetAccountDataBorshInto(ctx, client, offrampPDAs.referenceAddresses, rpc.CommitmentConfirmed, &referenceAddresses))
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
		selector := devnetInfo.ChainSelectors.Sepolia

		ix, err := example_ccip_sender.NewInitChainConfigInstruction(
			selector,
			[]byte{1}, // recipient
			[]byte{2}, // extra args bytes
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
		selector := devnetInfo.ChainSelectors.Sepolia

		ix, err := example_ccip_sender.NewUpdateChainConfigInstruction(
			selector,
			[]byte{3}, // recipient
			[]byte{4}, // extra args bytes
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
}
