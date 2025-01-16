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

	t.Run("AnyToSVM", func(t *testing.T) {
		t.Parallel()
		h, err := HashAnyToSVMMessage(ccip_router.Any2SVMRampMessage{
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
			ExtraArgs: ccip_router.SVMExtraArgs{
				ComputeUnits:     1000,
				IsWritableBitmap: 1,
				Accounts: []solana.PublicKey{
					solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb"),
				},
			},
			TokenAmounts: []ccip_router.Any2SVMTokenTransfer{
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
		require.Equal(t, "cabe4b03c22b7e6c684748f4c7abe710f82180fede7ad6e48fd3640fb74f8cfd", hex.EncodeToString(h))
	})

	t.Run("SVMToAny", func(t *testing.T) {
		t.Parallel()

		h, err := HashSVMToAnyMessage(ccip_router.SVM2AnyRampMessage{
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
			TokenAmounts: []ccip_router.SVM2AnyTokenTransfer{
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
		require.Equal(t, "0b7c4cd9dba8b90737e734bced7bb9ea43d9e5bab22ceb2c962ab9b04e2798af", hex.EncodeToString(h))
	})
}
