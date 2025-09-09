package contract_test

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var testContractType = cldf_deployment.ContractType("TestContract")

var OwnerAddress = common.HexToAddress("0x02")

// Implements rpc.DataError from go-ethereum/rpc
// Enables ABI decoding of revert reasons
type rpcError struct {
	Data interface{}
}

func (e *rpcError) Error() string {
	return ""
}

func (e *rpcError) ErrorData() interface{} {
	return e.Data
}

type testContract struct {
	address common.Address
	owner   common.Address
	value   int
}

func newTestContract(address common.Address, backend bind.ContractBackend) (*testContract, error) {
	return &testContract{
		address: address,
		value:   0,
	}, nil
}

func (t *testContract) Read(opts *bind.CallOpts, value int) (string, error) {
	if value%2 == 0 {
		return "even", nil
	}
	return "", fmt.Errorf("odd value: %d", value)
}

func (t *testContract) Write(opts *bind.TransactOpts, value int) (*types.Transaction, error) {
	// Not caught by operation validation, revert reason should be surfaced
	if value == 10 {
		return &types.Transaction{}, &rpcError{
			Data: common.Bytes2Hex(append(
				crypto.Keccak256([]byte("InvalidValue(uint256)"))[:4],
				common.LeftPadBytes([]byte{1}, 32)...,
			)),
		}
	}

	if value%2 == 0 {
		t.value = value
		return types.NewTx(&types.LegacyTx{
			To:   &t.address,
			Data: []byte{0xDE, 0xAD, 0xBE, 0xEF},
		}), nil
	}
	return &types.Transaction{}, fmt.Errorf("odd value: %d", value)
}

func (t *testContract) Owner(opts *bind.CallOpts) (common.Address, error) {
	return OwnerAddress, nil
}

func (t *testContract) Address() common.Address {
	return t.address
}
