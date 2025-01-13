package ccip

import (
	"encoding/hex"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
)

func TestMessageHashing(t *testing.T) {
	t.Parallel()

	sender := make([]byte, 32)
	copy(sender, []byte{1, 2, 3})
	tokenAmount := [32]uint8{}
	for i := range tokenAmount {
		tokenAmount[i] = 1
	}

	t.Run("EvmToSolana", func(t *testing.T) {
		t.Parallel()
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
			TokenAmounts: []ccip_router.Any2SolanaTokenTransfer{
				{
					SourcePoolAddress: []byte{0, 1, 2, 3},
					DestTokenAddress:  solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTc"),
					DestGasAmount:     100,
					ExtraData:         []byte{4, 5, 6},
					Amount:            tokenAmount,
				},
			},
		}, config.OnRampAddress)

		require.NoError(t, err)
		require.Equal(t, "fb47aed864f6e050f05ad851fdc0015c2e946f05e25093d150884cfb995834d0", hex.EncodeToString(h))
	})

	t.Run("SolanaToEvm", func(t *testing.T) {
		t.Parallel()

		h, err := HashSolanaToEvmMessage(ccip_router.Solana2AnyRampMessage{
			Header: ccip_router.RampMessageHeader{
				MessageId:           [32]uint8{},
				SourceChainSelector: 10,
				DestChainSelector:   20,
				SequenceNumber:      30,
				Nonce:               40,
			},
			Sender:   solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTa"),
			Data:     []byte{4, 5, 6},
			Receiver: sender,
			ExtraArgs: ccip_router.AnyExtraArgs{
				GasLimit:                 bin.Uint128{Lo: 1},
				AllowOutOfOrderExecution: true,
			},
			FeeToken:       solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb"),
			FeeTokenAmount: 50,
			TokenAmounts: []ccip_router.Solana2AnyTokenTransfer{
				{
					SourcePoolAddress: solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTc"),
					DestTokenAddress:  []byte{0, 1, 2, 3},
					ExtraData:         []byte{4, 5, 6},
					Amount:            tokenAmount,
					DestExecData:      []byte{4, 5, 6},
				},
			},
		})
		require.NoError(t, err)
		require.Equal(t, "9296d0ab425d1715b7709d1350e9486edc4ea235c47eed096b54bff20d07c692", hex.EncodeToString(h))
	})
}
