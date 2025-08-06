package evmimpls_test

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/evmimpls"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/evmimpls/gethwrappers"
)

func Test_Decode(t *testing.T) {
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
	_, _, contract, err := gethwrappers.DeployCCIPMessageSentEmitter(auth, sim.Client())
	require.NoError(t, err)
	sim.Commit()

	// 3. Create a test message and emit the event
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
				DestGasLimit:      100_000,
				DestBytesOverhead: 200,
				ExtraArgs:         []byte("receipt extra args"),
			},
		},
	}

	tx, err := contract.EmitCCIPMessageSent(auth, testMsg)
	require.NoError(t, err)
	sim.Commit()

	receipt, err := sim.Client().TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Len(t, receipt.Logs, 1)

	// 4. Decode the log and assert
	abi, err := abi.JSON(strings.NewReader(evmimpls.InternalABI))
	require.NoError(t, err)

	codec := evmimpls.NewEVMEventCodec(abi)

	decodedMsg, err := codec.Decode(context.Background(), receipt.Logs[0].Data)
	require.NoError(t, err)

	require.Equal(t, testMsg.Header.MessageId, decodedMsg.Header.MessageID)
	require.Equal(t, testMsg.Header.SourceChainSelector, decodedMsg.Header.SourceChainSelector)
	require.Equal(t, testMsg.Header.DestChainSelector, decodedMsg.Header.DestChainSelector)
	require.Equal(t, testMsg.Header.SequenceNumber, decodedMsg.Header.SequenceNumber)
	require.Equal(t, testMsg.Sender.Bytes(), decodedMsg.Sender)
	require.Equal(t, testMsg.Data, decodedMsg.Data)
	require.Equal(t, testMsg.Receiver, decodedMsg.Receiver)
	require.Equal(t, testMsg.FeeToken.Bytes(), decodedMsg.FeeToken)
	require.Equal(t, testMsg.FeeTokenAmount, decodedMsg.FeeTokenAmount)
	require.Equal(t, testMsg.FeeValueJuels, decodedMsg.FeeValueJuels)
	require.Equal(t, testMsg.TokenTransfer.SourceTokenAddress.Bytes(), decodedMsg.TokenTransfer.SourceTokenAddress)
	require.Equal(t, testMsg.TokenTransfer.SourcePoolAddress.Bytes(), decodedMsg.TokenTransfer.SourcePoolAddress)
	require.Equal(t, testMsg.TokenTransfer.DestTokenAddress, decodedMsg.TokenTransfer.DestTokenAddress)
	require.Equal(t, testMsg.TokenTransfer.ExtraData, decodedMsg.TokenTransfer.ExtraData)
	require.Equal(t, testMsg.TokenTransfer.Amount, decodedMsg.TokenTransfer.Amount)
	require.Equal(t, testMsg.TokenTransfer.DestExecData, decodedMsg.TokenTransfer.DestExecData)
	require.Equal(t, testMsg.TokenTransfer.RequiredVerifierId, decodedMsg.TokenTransfer.RequiredVerifierID)
	require.Len(t, decodedMsg.Receipts, 1)
	require.Equal(t, testMsg.Receipts[0].ReceiptType, decodedMsg.Receipts[0].ReceiptType)
	require.Equal(t, testMsg.Receipts[0].Issuer.Bytes(), decodedMsg.Receipts[0].Issuer)
	require.Equal(t, testMsg.Receipts[0].FeeTokenAmount, decodedMsg.Receipts[0].FeeTokenAmount)
	require.Equal(t, testMsg.Receipts[0].DestGasLimit, decodedMsg.Receipts[0].DestGasLimit)
	require.Equal(t, testMsg.Receipts[0].DestBytesOverhead, decodedMsg.Receipts[0].DestBytesOverhead)
	require.Equal(t, testMsg.Receipts[0].ExtraArgs, decodedMsg.Receipts[0].ExtraArgs)
}
