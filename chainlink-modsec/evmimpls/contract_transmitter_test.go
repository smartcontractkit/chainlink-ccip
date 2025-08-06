package evmimpls_test

import (
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

	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/evmimpls"
	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/evmimpls/gethwrappers"
)

func Test_ContractTransmitter(t *testing.T) {
	// Setup
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	require.NoError(t, err)

	genesisAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: new(big.Int).SetInt64(1000000000000000000)},
	}
	sim := simulated.NewBackend(genesisAlloc, simulated.WithBlockGasLimit(30e6))

	addr, _, contract, err := gethwrappers.DeployCCIPMessageSentEmitter(auth, sim.Client())
	require.NoError(t, err)
	sim.Commit()

	emitterABI, err := gethwrappers.CCIPMessageSentEmitterMetaData.GetAbi()
	require.NoError(t, err)

	executeMethod, ok := emitterABI.Methods["execute"]
	require.True(t, ok)

	transmitter := evmimpls.NewEVMContractTransmitter(sim.Client(), addr, executeMethod, auth)

	messageCodec, err := evmimpls.NewEVMMessageCodec()
	require.NoError(t, err)

	internalABI, err := abi.JSON(strings.NewReader(evmimpls.InternalABI))
	require.NoError(t, err)
	eventCodec := evmimpls.NewEVMEventCodec(internalABI)

	// Call transmit
	tx, err := contract.EmitCCIPMessageSent(auth, gethwrappers.CCIPMessageSentEmitterEVM2AnyVerifierMessage{
		Header: gethwrappers.CCIPMessageSentEmitterHeader{
			MessageId:           [32]byte{1},
			SourceChainSelector: 1,
			DestChainSelector:   2,
			SequenceNumber:      3,
		},
		Sender:         auth.From,
		Data:           []byte("test data"),
		Receiver:       common.HexToAddress("0x456").Bytes(),
		FeeToken:       common.HexToAddress("0x123"),
		FeeTokenAmount: big.NewInt(100),
		FeeValueJuels:  big.NewInt(200),
		TokenTransfer: gethwrappers.CCIPMessageSentEmitterEVMTokenTransfer{
			SourceTokenAddress: common.HexToAddress("0xabc"),
			SourcePoolAddress:  common.HexToAddress("0xabc"),
			DestTokenAddress:   []byte("dest token address"),
			ExtraData:          []byte("extra data"),
			Amount:             big.NewInt(300),
			DestExecData:       []byte("dest exec data"),
			RequiredVerifierId: [32]byte{2},
		},
		Receipts: []gethwrappers.CCIPMessageSentEmitterReceipt{
			{
				ReceiptType:       0,
				Issuer:            common.HexToAddress("0xdef"),
				FeeTokenAmount:    big.NewInt(50),
				DestGasLimit:      100000,
				DestBytesOverhead: 200,
				ExtraArgs:         []byte("receipt extra args"),
			},
		},
	})
	require.NoError(t, err)
	sim.Commit()

	receipt, err := sim.Client().TransactionReceipt(t.Context(), tx.Hash())
	require.NoError(t, err)

	message, err := eventCodec.Decode(t.Context(), receipt.Logs[0].Data)
	require.NoError(t, err)

	encodedMessage, err := messageCodec.Encode(t.Context(), message)
	require.NoError(t, err)

	proofs := [][]byte{[]byte("proof1"), []byte("proof2")}
	err = transmitter.Transmit(t.Context(), encodedMessage, proofs, nil)
	require.NoError(t, err)
	sim.Commit()

	// Check for the executed event
	it, err := contract.FilterExecuted(&bind.FilterOpts{Context: t.Context(), Start: 0})
	require.NoError(t, err)
	require.True(t, it.Next())

	require.Equal(t, message.Header.MessageID, it.Event.Message.Header.MessageId)
	require.Equal(t, proofs, it.Event.Proofs)
}
