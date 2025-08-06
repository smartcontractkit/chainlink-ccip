package evmimpls

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/modsectypes"
)

type EVMDestReader struct {
	caller                ethereum.ContractCaller
	offRampProxyAddress   common.Address
	commitVerifierAddress common.Address
	getNonceMethod        abi.Method
	isExecutedMethod      abi.Method
}

// Close implements modsectypes.DestReader.
func (e *EVMDestReader) Close() error {
	panic("unimplemented")
}

// GetNonce implements modsectypes.DestReader.
func (e *EVMDestReader) GetNonce(ctx context.Context, sourceChainSelector uint64, account []byte) (uint64, error) {
	callData, err := e.getNonceMethod.Inputs.Pack(sourceChainSelector, common.BytesToAddress(account))
	if err != nil {
		return 0, fmt.Errorf("failed to pack getNonce method inputs: %w", err)
	}

	callData = append(e.getNonceMethod.ID, callData...)

	callMsg := ethereum.CallMsg{
		To:   &e.offRampProxyAddress,
		Data: callData,
	}

	encodedNonce, err := e.caller.CallContract(ctx, callMsg, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to call getNonce method: %w", err)
	}

	output, err := e.getNonceMethod.Outputs.Unpack(encodedNonce)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack getNonce method outputs: %w", err)
	}

	nonce := *abi.ConvertType(output[0], new(uint64)).(*uint64)

	return nonce, nil
}

// IsExecuted implements modsectypes.DestReader.
func (e *EVMDestReader) IsExecuted(ctx context.Context, message modsectypes.Message) (bool, error) {
	callData, err := e.isExecutedMethod.Inputs.Pack(message.Header.SourceChainSelector, message.Header.SequenceNumber)
	if err != nil {
		return false, fmt.Errorf("failed to pack isExecuted method inputs: %w", err)
	}

	callData = append(e.isExecutedMethod.ID, callData...)

	callMsg := ethereum.CallMsg{
		To:   &e.offRampProxyAddress,
		Data: callData,
	}

	encodedExecuted, err := e.caller.CallContract(ctx, callMsg, nil)
	if err != nil {
		return false, fmt.Errorf("failed to call isExecuted method: %w", err)
	}

	output, err := e.isExecutedMethod.Outputs.Unpack(encodedExecuted)
	if err != nil {
		return false, fmt.Errorf("failed to unpack isExecuted method outputs: %w", err)
	}

	executed := *abi.ConvertType(output[0], new(bool)).(*bool)

	return executed, nil
}

// Start implements modsectypes.DestReader.
func (e *EVMDestReader) Start(ctx context.Context) error {
	return nil
}

var _ modsectypes.DestReader = (*EVMDestReader)(nil)

func NewEVMDestReader(
	caller ethereum.ContractCaller,
	offRampProxyAddress common.Address,
	commitVerifierAddress common.Address,
	getNonceMethod, isExecutedMethod abi.Method,
) *EVMDestReader {
	return &EVMDestReader{
		caller:                caller,
		offRampProxyAddress:   offRampProxyAddress,
		commitVerifierAddress: commitVerifierAddress,
		getNonceMethod:        getNonceMethod,
		isExecutedMethod:      isExecutedMethod,
	}
}
