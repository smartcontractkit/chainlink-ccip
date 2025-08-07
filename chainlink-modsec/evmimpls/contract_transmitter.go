package evmimpls

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/evmimpls/gethwrappers"
	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/libmodsec/modsectypes"
)

type EVMContractTransmitter struct {
	signer                         *bind.TransactOpts
	client                         TransactorClient
	offRampProxyAddress            common.Address
	executeMethod                  abi.Method
	decodeAny2EVMMultiProofMessage abi.Arguments
}

// Transmit implements modsectypes.ContractTransmitter.
func (e *EVMContractTransmitter) Transmit(
	ctx context.Context,
	encodedMessage []byte,
	proofs [][]byte,
	_ []byte,
) error {
	if len(e.executeMethod.Inputs) != 2 {
		return fmt.Errorf("execute method must have exactly two inputs")
	}

	// first argument of the execute method is the Any2EVMMultiProofMessage
	any2EVMArguments := abi.Arguments{e.executeMethod.Inputs[0]}
	unpacked, err := any2EVMArguments.Unpack(encodedMessage)
	if err != nil {
		return fmt.Errorf("failed to unpack Any2EVMMultiProofMessage: %w", err)
	}
	message := *abi.ConvertType(unpacked[0], new(gethwrappers.CCIPMessageSentEmitterAny2EVMMultiProofMessage)).(*gethwrappers.CCIPMessageSentEmitterAny2EVMMultiProofMessage)

	callData, err := e.executeMethod.Inputs.Pack(message, proofs)
	if err != nil {
		return fmt.Errorf("failed to pack execute method inputs: %w", err)
	}

	callData = append(e.executeMethod.ID, callData...)

	nonce, err := e.client.PendingNonceAt(ctx, e.signer.From)
	if err != nil {
		return fmt.Errorf("failed to get pending nonce for from address %s: %w", e.signer.From, err)
	}

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}

	gas, err := e.client.EstimateGas(ctx, ethereum.CallMsg{
		From: e.signer.From,
		To:   &e.offRampProxyAddress,
		Data: callData,
	})
	if err != nil {
		return fmt.Errorf("failed to estimate gas: %w", err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &e.offRampProxyAddress,
		Value:    big.NewInt(0),
		Gas:      gas,
		GasPrice: gasPrice,
		Data:     callData,
	})

	signedTx, err := e.signer.Signer(e.signer.From, tx)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	err = e.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	return nil
}

var _ modsectypes.ContractTransmitter = (*EVMContractTransmitter)(nil)

type TransactorClient interface {
	ethereum.TransactionSender
	ethereum.PendingStateReader
	ethereum.GasPricer
	ethereum.GasEstimator
}

func NewEVMContractTransmitter(
	client TransactorClient,
	offRampProxyAddress common.Address,
	executeMethod abi.Method,
	transactOpts *bind.TransactOpts,
) *EVMContractTransmitter {
	return &EVMContractTransmitter{
		client:              client,
		offRampProxyAddress: offRampProxyAddress,
		executeMethod:       executeMethod,
		signer:              transactOpts,
	}
}
