package contracts

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
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

	offrampPDAs, err := getOfframpPDAs(offrampAddress)
	require.NoError(t, err)

	var referenceAddresses ccip_offramp.ReferenceAddresses
	t.Run("Read Reference Addresses", func(t *testing.T) {
		require.NoError(t, common.GetAccountDataBorshInto(ctx, client, offrampPDAs.referenceAddresses, rpc.CommitmentConfirmed, &referenceAddresses))
		fmt.Printf("Reference Addresses: %+v\n", referenceAddresses)
	})

	ccip_router.SetProgramID(referenceAddresses.Router)

	user := solana.PrivateKey(devnetInfo.PrivateKeys.User)
	require.True(t, user.IsValid())

	configPDA, _, err := state.FindConfigPDA(referenceAddresses.Router)
	require.NoError(t, err)

	var routerConfig ccip_router.Config

	t.Run("Read Router Config PDA", func(t *testing.T) {
		require.NoError(t, common.GetAccountDataBorshInto(ctx, client, configPDA, rpc.CommitmentConfirmed, &routerConfig))
		fmt.Printf("Router Config: %+v\n", routerConfig)
	})

	destinationChainSelector := uint64(devnetInfo.ChainSelectors.Sepolia)

	destinationChainStatePDA, err := state.FindDestChainStatePDA(destinationChainSelector, referenceAddresses.Router)
	require.NoError(t, err)

	nonceEvmPDA, err := state.FindNoncePDA(destinationChainSelector, user.PublicKey(), referenceAddresses.Router)
	require.NoError(t, err)

	billingSignerPDA, _, err := state.FindFeeBillingSignerPDA(referenceAddresses.Router)
	require.NoError(t, err)

	wsolReceiver, _, err := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, solana.SolMint, billingSignerPDA)
	require.NoError(t, err)

	fqConfigPDA, _, err := state.FindFqConfigPDA(routerConfig.FeeQuoter)
	require.NoError(t, err)

	fqDestChainPDA, _, err := state.FindFqDestChainPDA(destinationChainSelector, routerConfig.FeeQuoter)
	require.NoError(t, err)

	fqWsolBillingTokenConfigPDA, _, err := state.FindFqBillingTokenConfigPDA(solana.SolMint, routerConfig.FeeQuoter)
	require.NoError(t, err)

	fqLinkBillingConfigPDA, _, err := state.FindFqBillingTokenConfigPDA(routerConfig.LinkTokenMint, routerConfig.FeeQuoter)
	require.NoError(t, err)

	externalTpSignerPDA, _, err := state.FindExternalTokenPoolsSignerPDA(referenceAddresses.Router)
	require.NoError(t, err)

	// t.Run("Read other PDAs", func(t *testing.T) {
	// 	var calcDestChainState ccip_router.DestChain
	// 	require.NoError(t, common.GetAccountDataBorshInto(ctx, client, destinationChainStatePDA, rpc.CommitmentConfirmed, &calcDestChainState))
	// 	fmt.Printf("(Calculated) DestChainState: %+v\n", calcDestChainState)

	// 	var loggedDestChainState ccip_router.DestChain
	// 	require.NoError(t, common.GetAccountDataBorshInto(ctx, client, solana.MustPublicKeyFromBase58("6WwuJ4Z2RCkzZs2XY5TRhhGVGeBwSzC74sSBzoddfBPd"), rpc.CommitmentConfirmed, &loggedDestChainState))
	// 	fmt.Printf("(Logged) DestChainState: %+v\n", loggedDestChainState)
	// })

	t.Run("Read event", func(t *testing.T) {
		txsig, err := solana.SignatureFromBase58("61qFfCUu5wsoLZSt9B9MjAiokFFxSHywboZk1rFtE57ndF2GhYmQkJDqgJ4xC14iizsaX2prQXtfrmmN4pmMJwtx")
		require.NoError(t, err)

		v := uint64(0)
		tx, err := client.GetTransaction(ctx, txsig, &rpc.GetTransactionOpts{
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &v,
		})
		require.NoError(t, err)
		require.NotNil(t, tx)

		var event ccip.EventConfigSet
		require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "ConfigSet", &event, true))
	})

	t.Run("Offramp: Commit prices only", func(t *testing.T) {
		report := ccip_offramp.CommitInput{
			MerkleRoot: nil,
			PriceUpdates: ccip_offramp.PriceUpdates{
				TokenPriceUpdates: []ccip_offramp.TokenPriceUpdate{
					{SourceToken: solana.SolMint, UsdPerToken: common.To28BytesBE(3)},
					{SourceToken: routerConfig.LinkTokenMint, UsdPerToken: common.To28BytesBE(4)},
				},
				GasPriceUpdates: []ccip_offramp.GasPriceUpdate{
					{DestChainSelector: destinationChainSelector, UsdPerUnitGas: common.To28BytesBE(5)},
				},
			},
		}

		// signers, _, _ := testutils.GenerateSignersAndTransmitters(t, 16) // TODO

		var initialConfig ccip_offramp.Config
		require.NoError(t, common.GetAccountDataBorshInto(ctx, client, offrampPDAs.config, rpc.CommitmentConfirmed, &initialConfig))

		index := uint8(testutils.OcrCommitPlugin)
		commitConfig := initialConfig.Ocr3[index]

		var sourceChain ccip_offramp.SourceChain
		sequence := sourceChain.State.MinSeqNr

		empty24byte := [24]byte{}
		var reportContext [2][32]byte
		reportContext[0] = commitConfig.ConfigInfo.ConfigDigest                              // match the onchain config digest
		reportContext[1] = [32]byte(binary.BigEndian.AppendUint64(empty24byte[:], sequence)) // sequence number

		signers := []eth.Signer{}
		for _, signerPrivK := range devnetInfo.PrivateKeys.Signers {
			signer, err := eth.GetSignerFromPk(string(signerPrivK))
			require.NoError(t, err)
			signers = append(signers, signer)
		}
		sigs, err := ccip.SignCommitReport(reportContext, report, signers)
		require.NoError(t, err)

		transmitter := solana.PrivateKey(devnetInfo.PrivateKeys.Transmitter)

		fqAllowedPriceUpdater, _, err := state.FindFqAllowedPriceUpdaterPDA(offrampPDAs.billingSigner, referenceAddresses.FeeQuoter)
		require.NoError(t, err)

		fqConfigPDA, _, err := state.FindFqConfigPDA(referenceAddresses.FeeQuoter)
		require.NoError(t, err)

		raw := ccip_offramp.NewCommitPriceOnlyInstruction(
			reportContext,
			testutils.MustMarshalBorsh(t, report),
			sigs.Rs,
			sigs.Ss,
			sigs.RawVs,
			offrampPDAs.config,
			offrampPDAs.referenceAddresses,
			transmitter.PublicKey(),
			solana.SystemProgramID,
			solana.SysVarInstructionsPubkey,
			offrampPDAs.billingSigner,
			referenceAddresses.FeeQuoter,
			fqAllowedPriceUpdater,
			fqConfigPDA,
		)

		remainingAccounts := []solana.PublicKey{
			offrampPDAs.state,
			fqWsolBillingTokenConfigPDA,
			fqLinkBillingConfigPDA,
			fqDestChainPDA,
		}

		for _, pubkey := range remainingAccounts {
			raw.AccountMetaSlice.Append(solana.Meta(pubkey).WRITE())
		}

		instruction, err := raw.ValidateAndBuild()
		require.NoError(t, err)

		lookupTableEntries, err := common.GetAddressLookupTable(ctx, client, referenceAddresses.OfframpLookupTable)
		require.NoError(t, err)
		table := map[solana.PublicKey]solana.PublicKeySlice{
			referenceAddresses.OfframpLookupTable: lookupTableEntries,
		}

		tx := testutils.SendAndConfirmWithLookupTables(ctx, t, client, []solana.Instruction{instruction}, transmitter, rpc.CommitmentConfirmed, table, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
		commitEvent := ccip.EventCommitReportAccepted{}
		require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &commitEvent, true))
	})

	t.Run("OnRamp: CCIP Send", func(t *testing.T) {
		t.Skip("This test is not yet to be run")
		message := ccip_router.SVM2AnyMessage{
			FeeToken:  solana.PublicKey{},
			Receiver:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 33, 12, 248, 134, 206, 65, 149, 35, 22, 68, 26, 228, 202, 195, 95, 0, 240, 232, 130, 166}, // just an example EVM address
			Data:      []byte{4, 5, 6},
			ExtraArgs: []byte{},
		}

		raw := ccip_router.NewCcipSendInstruction(
			destinationChainSelector,
			message,
			[]byte{},
			configPDA,
			destinationChainStatePDA,
			nonceEvmPDA,
			user.PublicKey(),
			solana.SystemProgramID,
			solana.TokenProgramID,
			solana.SolMint,
			solana.PublicKey{},
			wsolReceiver,
			billingSignerPDA,
			routerConfig.FeeQuoter,
			fqConfigPDA,
			fqDestChainPDA,
			fqWsolBillingTokenConfigPDA,
			fqLinkBillingConfigPDA,
			externalTpSignerPDA,
		)
		// raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
		instruction, err := raw.ValidateAndBuild()
		require.NoError(t, err)
		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{instruction}, user, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
	})
}
