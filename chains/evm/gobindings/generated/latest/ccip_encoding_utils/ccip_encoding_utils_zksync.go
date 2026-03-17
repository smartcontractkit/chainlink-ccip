// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package ccip_encoding_utils

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/types"
)

func DeployEncodingUtilsZk(deployOpts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ...interface{}) (common.Address, *types.Receipt, *EncodingUtils, error) {
	var calldata []byte
	if len(args) > 0 {
		abi, err := EncodingUtilsMetaData.GetAbi()
		if err != nil {
			return common.Address{}, nil, nil, err
		}
		calldata, err = abi.Pack("", args...)
		if err != nil {
			return common.Address{}, nil, nil, err
		}
	}

	salt := make([]byte, 32)
	n, err := rand.Read(salt)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if n != len(salt) {
		return common.Address{}, nil, nil, fmt.Errorf("failed to read random bytes: expected %d, got %d", len(salt), n)
	}

	txHash, err := wallet.Deploy(deployOpts, accounts.Create2Transaction{
		Bytecode: ZkBytecode,
		Calldata: calldata,
		Salt:     salt,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	receipt, err := client.WaitMined(context.Background(), txHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address := receipt.ContractAddress
	contract, err := NewEncodingUtils(address, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	return address, receipt, contract, nil
}

var ZkBytecode = common.Hex2Bytes("")
