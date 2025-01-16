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
			Sender:        sender,
			TokenReceiver: solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb"),
			LogicReceiver: solana.MustPublicKeyFromBase58("C8WSPj3yyus1YN3yNB6YA5zStYtbjQWtpmKadmvyUXq8"),
			Data:          []byte{4, 5, 6},
			Header: ccip_router.RampMessageHeader{
				MessageId:           [32]uint8{8, 5, 3},
				SourceChainSelector: 67,
				DestChainSelector:   78,
				SequenceNumber:      89,
				Nonce:               90,
			},
			ExtraArgs: ccip_router.SolanaExtraArgs{
				ComputeUnits:     1000,
				IsWritableBitmap: 1,
				Accounts: []solana.PublicKey{
					solana.MustPublicKeyFromBase58("CtEVnHsQzhTNWav8skikiV2oF6Xx7r7uGGa8eCDQtTjH"),
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
		require.Equal(t, "46931be172374199bbf69f7138e18360a744bc5cf1159ffeacf43aaa53d427db", hex.EncodeToString(h))
	})

	t.Run("SolanaToEvm", func(t *testing.T) {
		t.Parallel()

		h, err := HashSolanaToAnyMessage(ccip_router.Solana2AnyRampMessage{
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
		require.Equal(t, "557e0080a3616647be8f376859d4c991778a21859e266ab3c92edfa04655f5dc", hex.EncodeToString(h))
	})
}
