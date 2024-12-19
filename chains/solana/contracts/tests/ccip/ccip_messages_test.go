package contracts

import (
	"encoding/hex"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
)

func TestHashEvmToSolanaMessage(t *testing.T) {
	t.Parallel()

	sender := make([]byte, 32)
	copy(sender, []byte{1, 2, 3})

	h, err := HashEvmToSolanaMessage(ccip_router.Any2SolanaRampMessage{
		Sender:   sender,
		Receiver: solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb"),
		Data:     []byte{4, 5, 6},
		Header: ccip_router.RampMessageHeader{
			MessageId:           [32]uint8{8, 5, 3},
			SourceChainSelector: 67,
			DestChainSelector:   78,
			SequenceNumber:      89,
			Nonce:               90,
		},
		ExtraArgs: ccip_router.SolanaExtraArgs{
			ComputeUnits: 1000,
			Accounts: []ccip_router.SolanaAccountMeta{
				{
					Pubkey:     solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb"),
					IsWritable: true,
				},
			},
		},
	}, config.OnRampAddress)
	require.NoError(t, err)
	require.Equal(t, "03da97f96c82237d8a8ab0f68d4f7ba02afe188b4a876f348278fbf2226312ed", hex.EncodeToString(h))
}
