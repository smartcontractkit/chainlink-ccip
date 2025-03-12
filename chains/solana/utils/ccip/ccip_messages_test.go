package ccip

import (
	"encoding/hex"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

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
			[]solana.PublicKey{
				solana.MustPublicKeyFromBase58("7gtMT88cNmZVmVUzkcC6MKZ4NY27ynwodKBVJFQdq8R2"),
				solana.MustPublicKeyFromBase58("EvhgrPhTDt4LcSPS2kfJgH6T6XWZ6wT3X9ncDGLT1vui"),
			})

		require.NoError(t, err)
		require.Equal(t, "088a485369cc2f2222e9fd1c997d317d6de29ff9d1d4181fbb4a45018277a3da", hex.EncodeToString(h))
	})

	t.Run("SVMToAny", func(t *testing.T) {
		t.Parallel()

		extraArgs, err := SerializeExtraArgs(fee_quoter.EVMExtraArgsV2{
			GasLimit:                 bin.Uint128{Lo: 1},
			AllowOutOfOrderExecution: true,
		}, EVMExtraArgsV2Tag)
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
		require.Equal(t, "a43d7d70b20ea71a61a1ac3eb1de4b04126678460487157b6a93d4cb10b73609", hex.EncodeToString(h))
	})
}
