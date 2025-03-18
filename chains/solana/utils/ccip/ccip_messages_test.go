package ccip

import (
	"encoding/hex"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
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
		h, err := HashAnyToSVMMessage(ccip_offramp.Any2SVMRampMessage{
			Sender:        sender,
			TokenReceiver: solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb"),
			Data:          []byte{4, 5, 6},
			Header: ccip_offramp.RampMessageHeader{
				MessageId:           [32]uint8{8, 5, 3},
				SourceChainSelector: 67,
				DestChainSelector:   78,
				SequenceNumber:      89,
				Nonce:               90,
			},
			ExtraArgs: ccip_offramp.Any2SVMRampExtraArgs{
				ComputeUnits:     1000,
				IsWritableBitmap: GenerateBitMapForIndexes([]int{0}),
			},
			TokenAmounts: []ccip_offramp.Any2SVMTokenTransfer{
				{
					SourcePoolAddress: []byte{0, 1, 2, 3},
					DestTokenAddress:  solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTc"),
					DestGasAmount:     100,
					ExtraData:         []byte{4, 5, 6},
					Amount:            ccip_offramp.CrossChainAmount{LeBytes: tokenAmount},
				},
			},
		},
			config.OnRampAddress,
			[]solana.PublicKey{
				solana.MustPublicKeyFromBase58("C8WSPj3yyus1YN3yNB6YA5zStYtbjQWtpmKadmvyUXq8"),
				solana.MustPublicKeyFromBase58("CtEVnHsQzhTNWav8skikiV2oF6Xx7r7uGGa8eCDQtTjH"),
			})

		require.NoError(t, err)
		require.Equal(t, "bd8025f7b32386d93be284b6b4eb6f36c7b46ea157c0228f00ccba38fe7a448e", hex.EncodeToString(h))
	})

	t.Run("SVMToAny", func(t *testing.T) {
		t.Parallel()

		extraArgs, err := SerializeExtraArgs(fee_quoter.GenericExtraArgsV2{
			GasLimit:                 bin.Uint128{Lo: 1},
			AllowOutOfOrderExecution: true,
		}, GenericExtraArgsV2Tag)
		require.NoError(t, err)

		h, err := HashSVMToAnyMessage(ccip_router.SVM2AnyRampMessage{
			Header: ccip_router.RampMessageHeader{
				MessageId:           [32]uint8{},
				SourceChainSelector: 10,
				DestChainSelector:   20,
				SequenceNumber:      30,
				Nonce:               40,
			},
			Sender:         solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTa"),
			Data:           []byte{4, 5, 6},
			Receiver:       sender,
			ExtraArgs:      extraArgs,
			FeeToken:       solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb"),
			FeeTokenAmount: ccip_router.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(50)},
			FeeValueJuels:  ccip_router.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(500)},
			TokenAmounts: []ccip_router.SVM2AnyTokenTransfer{
				{
					SourcePoolAddress: solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTc"),
					DestTokenAddress:  []byte{0, 1, 2, 3},
					ExtraData:         []byte{4, 5, 6},
					Amount:            ccip_router.CrossChainAmount{LeBytes: tokenAmount},
					DestExecData:      []byte{4, 5, 6},
				},
			},
		})
		require.NoError(t, err)
		require.Equal(t, "ab7f57fbf9979573a9fa1bcf1ad816449223c343ea8b938db9e917f2ca138a84", hex.EncodeToString(h))
	})
}
