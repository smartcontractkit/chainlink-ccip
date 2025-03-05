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
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
)

func TestConfigureDevnet(t *testing.T) {
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

	admin := solana.PrivateKey(devnetInfo.PrivateKeys.Admin)
	require.True(t, admin.IsValid())

	t.Run("Override commit OCR config", func(t *testing.T) {
		var initialConfig ccip_offramp.Config
		require.NoError(t, common.GetAccountDataBorshInto(ctx, client, offrampPDAs.config, rpc.CommitmentConfirmed, &initialConfig))

		index := uint8(testutils.OcrCommitPlugin)

		fmt.Printf("initial OCR config: %+v\n", initialConfig.Ocr3[index])

		commitConfig := initialConfig.Ocr3[index]

		transmitters := []solana.PublicKey{}
		for _, transmitter := range commitConfig.Transmitters {
			if (solana.PublicKey{}.Equals(transmitter)) {
				continue
			}
			transmitters = append(transmitters, transmitter)
		}
		foundTransmitter := false
		newTransmitter := solana.PrivateKey(devnetInfo.PrivateKeys.Transmitter).PublicKey()
		for _, transmitter := range transmitters {
			if transmitter.Equals(newTransmitter) {
				foundTransmitter = true
				break
			}
		}
		if !foundTransmitter {
			transmitters = append(transmitters, newTransmitter)
		}

		ocrConfigInfo := commitConfig.ConfigInfo

		signers := [][20]uint8{}
		for _, signer := range commitConfig.Signers {
			if [20]byte{} == signer {
				continue
			}
			signers = append(signers, signer)
		}
		for _, signerPrivK := range devnetInfo.PrivateKeys.Signers {
			signer, err := eth.GetSignerFromPk(signerPrivK)
			require.NoError(t, err)
			// check signer address is not in signers already
			found := false
			for _, s := range signers {
				if s == signer.Address {
					found = true
					break
				}
			}
			if !found {
				signers = append(signers, signer.Address)
			}
		}

		ix, err := ccip_offramp.NewSetOcrConfigInstruction(
			index,
			ocrConfigInfo,
			signers,      // signers
			transmitters, // transmitters
			offrampPDAs.config,
			offrampPDAs.state,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)
	})
}
