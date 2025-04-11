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
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_receiver"
)

func TestDevnet(t *testing.T) {
	devnetInfo, err := getDevnetInfo()
	require.NoError(t, err)

	ctx := tests.Context(t)
	client := rpc.New(devnetInfo.RPC)

	offrampAddress, err := solana.PublicKeyFromBase58(devnetInfo.Offramp)
	require.NoError(t, err)

	// this makes it so that instructions for the router target the right program id in devnet
	ccip_offramp.SetProgramID(offrampAddress)

	// offrampPDAs, err := getOfframpPDAs(offrampAddress)
	// require.NoError(t, err)

	// var referenceAddresses ccip_offramp.ReferenceAddresses

	admin := solana.PrivateKey(devnetInfo.PrivateKeys.Admin)
	require.True(t, admin.IsValid())

	testReceiverAddress := solana.MustPublicKeyFromBase58("D66KPujkoA2B81z2vUnmAa9PdzMbdXrVdiZ9otxUPevd")

	t.Run("Toggle Receiver", func(t *testing.T) {
		test_ccip_receiver.SetProgramID(testReceiverAddress)

		counterPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("counter")}, testReceiverAddress)
		require.NoError(t, err)

		ix, err := test_ccip_receiver.NewSetRejectAllInstruction(true, counterPDA, admin.PublicKey()).ValidateAndBuild()
		require.NoError(t, err)

		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)

		require.NotNil(t, result)

		fmt.Printf("Result: %s\n", result.Meta.LogMessages)
	})

	// t.Run("Check IDs", func(t *testing.T) {
	// 	ccip_router.SetProgramID(solana.MustPublicKeyFromBase58("HH36NJUF3fg9redq9LRQWJY2aAdceQg8z7YD7ZmDVbgK"))
	// 	ix, err := ccip_router.NewTestIdInstruction().ValidateAndBuild()
	// 	require.NoError(t, err)
	// 	result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, solana.PrivateKey(devnetInfo.PrivateKeys.User), rpc.CommitmentConfirmed)
	// 	require.NotNil(t, result)

	// 	fmt.Printf("Result: %s\n", result.Meta.LogMessages)

	// 	t.FailNow()
	// })

	// t.Run("Read event", func(t *testing.T) {
	// 	txsig, err := solana.SignatureFromBase58("dTWzQJ2FELaREGAQ4dMeuDTfXYwY42LSczTfUSRyqTqi3WvG4TBuxLCdZtcZnYtqWKheP5evXdQFrJgYVSG8SQi")
	// 	require.NoError(t, err)

	// 	v := uint64(0)
	// 	tx, err := client.GetTransaction(ctx, txsig, &rpc.GetTransactionOpts{
	// 		Commitment:                     rpc.CommitmentConfirmed,
	// 		MaxSupportedTransactionVersion: &v,
	// 	})
	// 	require.NoError(t, err)
	// 	require.NotNil(t, tx)

	// 	var event ccip.UsdPerUnitGasUpdated
	// 	require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerUnitGasUpdated", &event, true))
	// })

	// t.Run("Read Reference Addresses", func(t *testing.T) {
	// 	require.NoError(t, common.GetAccountDataBorshInto(ctx, client, offrampPDAs.referenceAddresses, rpc.CommitmentConfirmed, &referenceAddresses))
	// 	fmt.Printf("Reference Addresses: %+v\n", referenceAddresses)
	// })

	// ccip_router.SetProgramID(referenceAddresses.Router)
	// fee_quoter.SetProgramID(referenceAddresses.FeeQuoter)

	// user := solana.PrivateKey(devnetInfo.PrivateKeys.User)
	// require.True(t, user.IsValid())

	// configPDA, _, err := state.FindConfigPDA(referenceAddresses.Router)
	// require.NoError(t, err)

	// var routerConfig ccip_router.Config

	// t.Run("Read Router Config PDA", func(t *testing.T) {
	// 	require.NoError(t, common.GetAccountDataBorshInto(ctx, client, configPDA, rpc.CommitmentConfirmed, &routerConfig))
	// 	fmt.Printf("Router Config: %+v\n", routerConfig)
	// })

	// chainSelector := uint64(devnetInfo.ChainSelectors.Sepolia)

	// destinationChainStatePDA, err := state.FindDestChainStatePDA(chainSelector, referenceAddresses.Router)
	// require.NoError(t, err)

	// nonceEvmPDA, err := state.FindNoncePDA(chainSelector, user.PublicKey(), referenceAddresses.Router)
	// require.NoError(t, err)

	// billingSignerPDA, _, err := state.FindFeeBillingSignerPDA(referenceAddresses.Router)
	// require.NoError(t, err)

	// wsolReceiver, _, err := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, solana.SolMint, billingSignerPDA)
	// require.NoError(t, err)

	// fqConfigPDA, _, err := state.FindFqConfigPDA(routerConfig.FeeQuoter)
	// require.NoError(t, err)

	// fqDestChainPDA, _, err := state.FindFqDestChainPDA(chainSelector, routerConfig.FeeQuoter)
	// require.NoError(t, err)

	// fqWsolBillingTokenConfigPDA, _, err := state.FindFqBillingTokenConfigPDA(solana.SolMint, routerConfig.FeeQuoter)
	// require.NoError(t, err)

	// fqLinkBillingConfigPDA, _, err := state.FindFqBillingTokenConfigPDA(routerConfig.LinkTokenMint, routerConfig.FeeQuoter)
	// require.NoError(t, err)

	// externalTpSignerPDA, _, err := state.FindExternalTokenPoolsSignerPDA(referenceAddresses.Router)
	// require.NoError(t, err)

	// t.Run("Read other PDAs", func(t *testing.T) {
	// 	var calcDestChainState ccip_router.DestChain
	// 	require.NoError(t, common.GetAccountDataBorshInto(ctx, client, destinationChainStatePDA, rpc.CommitmentConfirmed, &calcDestChainState))
	// 	fmt.Printf("(Calculated) DestChainState: %+v\n", calcDestChainState)

	// 	var loggedDestChainState ccip_router.DestChain
	// 	require.NoError(t, common.GetAccountDataBorshInto(ctx, client, solana.MustPublicKeyFromBase58("6WwuJ4Z2RCkzZs2XY5TRhhGVGeBwSzC74sSBzoddfBPd"), rpc.CommitmentConfirmed, &loggedDestChainState))
	// 	fmt.Printf("(Logged) DestChainState: %+v\n", loggedDestChainState)
	// })
}
