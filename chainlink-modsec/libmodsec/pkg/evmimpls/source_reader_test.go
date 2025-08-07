package evmimpls_test

import (
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/evmimpls"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/evmimpls/gethwrappers"
)

func Test_SourceReader(t *testing.T) {
	// 1. Setup simulated chain and deployer
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	require.NoError(t, err)

	genesisAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: new(big.Int).SetInt64(1000000000000000000)},
	}
	sim := simulated.NewBackend(genesisAlloc, simulated.WithBlockGasLimit(30e6))

	// 2. Deploy the contract
	addr, _, contract, err := gethwrappers.DeployCCIPMessageSentEmitter(auth, sim.Client())
	require.NoError(t, err)
	sim.Commit()

	// 3. Instantiate Source Reader
	abi, err := abi.JSON(strings.NewReader(evmimpls.InternalABI))
	require.NoError(t, err)

	codec := evmimpls.NewEVMEventCodec(abi)
	logger := log.Default()

	event, ok := abi.Events["CCIPMessageSent"]
	require.True(t, ok)

	sourceReader := evmimpls.NewEVMSourceReader(sim.Client(), addr, event.ID, logger, codec)

	// 4. Start the source reader
	err = sourceReader.Start(t.Context())
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, sourceReader.Close())
	})

	// 5. Emit an event
	testMsg := gethwrappers.CCIPMessageSentEmitterEVM2AnyVerifierMessage{
		Header: gethwrappers.CCIPMessageSentEmitterHeader{
			MessageId:           [32]byte{1},
			SourceChainSelector: 1,
			DestChainSelector:   2,
			SequenceNumber:      3,
		},
		Sender:         auth.From,
		Data:           []byte("test data"),
		Receiver:       []byte("test receiver"),
		FeeToken:       common.HexToAddress("0x123"),
		FeeTokenAmount: big.NewInt(100),
		FeeValueJuels:  big.NewInt(200),
		TokenTransfer: gethwrappers.CCIPMessageSentEmitterEVMTokenTransfer{
			SourceTokenAddress: common.HexToAddress("0x456"),
			SourcePoolAddress:  common.HexToAddress("0x789"),
			DestTokenAddress:   []byte("dest token address"),
			ExtraData:          []byte("extra data"),
			Amount:             big.NewInt(300),
			DestExecData:       []byte("dest exec data"),
			RequiredVerifierId: [32]byte{2},
		},
		Receipts: []gethwrappers.CCIPMessageSentEmitterReceipt{
			{
				ReceiptType:       0,
				Issuer:            common.HexToAddress("0xabc"),
				FeeTokenAmount:    big.NewInt(50),
				DestGasLimit:      100000,
				DestBytesOverhead: 200,
				ExtraArgs:         []byte("receipt extra args"),
			},
		},
	}

	_, err = contract.EmitCCIPMessageSent(auth, testMsg)
	require.NoError(t, err)
	sim.Commit()

	// 6. Verify the message is received
	select {
	case receivedMsg := <-sourceReader.Messages():
		require.Equal(t, testMsg.Header.MessageId, receivedMsg.Header.MessageID)
		require.Equal(t, testMsg.Header.SourceChainSelector, receivedMsg.Header.SourceChainSelector)
		require.Equal(t, testMsg.Header.DestChainSelector, receivedMsg.Header.DestChainSelector)
		require.Equal(t, testMsg.Header.SequenceNumber, receivedMsg.Header.SequenceNumber)
		require.Equal(t, testMsg.Sender.Bytes(), receivedMsg.Sender)
		require.Equal(t, testMsg.Data, receivedMsg.Data)
		require.Equal(t, testMsg.Receiver, receivedMsg.Receiver)
		require.Equal(t, testMsg.FeeToken.Bytes(), receivedMsg.FeeToken)
		require.Equal(t, testMsg.FeeTokenAmount, receivedMsg.FeeTokenAmount)
		require.Equal(t, testMsg.FeeValueJuels, receivedMsg.FeeValueJuels)
		require.Equal(t, testMsg.TokenTransfer.SourceTokenAddress.Bytes(), receivedMsg.TokenTransfer.SourceTokenAddress)
		require.Equal(t, testMsg.TokenTransfer.SourcePoolAddress.Bytes(), receivedMsg.TokenTransfer.SourcePoolAddress)
		require.Equal(t, testMsg.TokenTransfer.DestTokenAddress, receivedMsg.TokenTransfer.DestTokenAddress)
		require.Equal(t, testMsg.TokenTransfer.ExtraData, receivedMsg.TokenTransfer.ExtraData)
		require.Equal(t, testMsg.TokenTransfer.Amount, receivedMsg.TokenTransfer.Amount)
		require.Equal(t, testMsg.TokenTransfer.DestExecData, receivedMsg.TokenTransfer.DestExecData)
		require.Equal(t, testMsg.TokenTransfer.RequiredVerifierId, receivedMsg.TokenTransfer.RequiredVerifierID)
		require.Len(t, receivedMsg.Receipts, 1)
		require.Equal(t, testMsg.Receipts[0].ReceiptType, receivedMsg.Receipts[0].ReceiptType)
		require.Equal(t, testMsg.Receipts[0].Issuer.Bytes(), receivedMsg.Receipts[0].Issuer)
		require.Equal(t, testMsg.Receipts[0].FeeTokenAmount, receivedMsg.Receipts[0].FeeTokenAmount)
		require.Equal(t, testMsg.Receipts[0].DestGasLimit, receivedMsg.Receipts[0].DestGasLimit)
		require.Equal(t, testMsg.Receipts[0].DestBytesOverhead, receivedMsg.Receipts[0].DestBytesOverhead)
		require.Equal(t, testMsg.Receipts[0].ExtraArgs, receivedMsg.Receipts[0].ExtraArgs)
	case <-t.Context().Done():
		t.Fatal("timed out waiting for message")
	}
}
