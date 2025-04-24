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
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/redirecting_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

func TestSetupRedirectingReceiver(t *testing.T) {
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

	redirectingReceiverAddress := solana.MustPublicKeyFromBase58(devnetInfo.RedirectingReceiver)

	redirecting_ccip_receiver.SetProgramID(redirectingReceiverAddress)

	statePDA, _, err := solana.FindProgramAddress([][]byte{[]byte("state")}, redirectingReceiverAddress)
	require.NoError(t, err)
	tokenAdminPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("receiver_token_admin")}, redirectingReceiverAddress)
	require.NoError(t, err)

	t.Run("Initialize redirecting receiver", func(t *testing.T) {
		t.Skip()
		ix, err := redirecting_ccip_receiver.NewInitializeInstruction(
			referenceAddresses.Router,
			statePDA,
			tokenAdminPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
		fmt.Printf("Result: %s\n", result.Meta.LogMessages)
	})

	t.Log()
}
