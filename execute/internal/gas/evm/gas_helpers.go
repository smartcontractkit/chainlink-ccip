// Package evm provides an EVM implementation to the gas.EstimateProvider interface.
// TODO: Move this package into the EVM repo, chainlink-ccip should be chain agnostic.
package evm

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/internal/gas"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	EvmAddressLengthBytes           = 20
	EvmWordBytes                    = 32
	CalldataGasPerByte              = 16
	TokenAdminRegistryWarmupCost    = 2_500
	TokenAdminRegistryPoolLookupGas = 100 + // WARM_ACCESS_COST TokenAdminRegistry
		700 + // CALL cost for TokenAdminRegistry
		2_100 // COLD_SLOAD_COST loading the pool address
	SupportsInterfaceCheck = 2600 + // because the receiver will be untouched initially
		30_000*3 // supportsInterface of ERC165Checker library performs 3 static-calls of 30k gas each
	PerTokenOverheadGas = TokenAdminRegistryPoolLookupGas +
		SupportsInterfaceCheck +
		200_000 + // releaseOrMint using callWithExactGas
		50_000 // transfer using callWithExactGas
	RateLimiterOverheadGas = 2_100 + // COLD_SLOAD_COST for accessing token bucket
		5_000 // SSTORE_RESET_GAS for updating & decreasing token bucket
	ConstantMessagePartBytes            = 10 * 32 // A message consists of 10 abi encoded fields 32B each (after encoding)
	ExecutionStateProcessingOverheadGas = 2_100 + // COLD_SLOAD_COST for first reading the state
		20_000 + // SSTORE_SET_GAS for writing from 0 (untouched) to non-zero (in-progress)
		100 //# SLOAD_GAS = WARM_STORAGE_READ_COST for rewriting from non-zero (in-progress) to non-zero (success/failure)
)

var (
	abiUint32  = ABITypeOrPanic("uint32")
	abiUint256 = ABITypeOrPanic("uint256")
	abiBool    = ABITypeOrPanic("bool")

	// ExtraArgsV2ABI ABI types for EVMExtraArgsV2 (see Client.sol)
	ExtraArgsV2ABI = abi.Arguments{
		{
			Name: "gasLimit",
			Type: abiUint256,
		},
		{
			Name: "allowOutOfOrderExecution",
			Type: abiBool,
		},
	}

	TokenDestGasOverheadABI = abi.Arguments{
		{
			Type: abiUint32,
		},
	}
)

type EstimateProvider struct {
	lggr logger.Logger
}

var _ gas.EstimateProvider = &EstimateProvider{}

func NewEstimateProvider(lggr logger.Logger) EstimateProvider {
	return EstimateProvider{lggr}
}

// CalculateMerkleTreeGas estimates the merkle tree gas based on number of requests
func (gp EstimateProvider) CalculateMerkleTreeGas(numRequests int) uint64 {
	if numRequests == 0 {
		return 0
	}
	merkleProofBytes := (math.Ceil(math.Log2(float64(numRequests))))*32 + (1+2)*32 // only ever one outer root hash
	return uint64(merkleProofBytes * CalldataGasPerByte)
}

// return the size of bytes for msg tokens
func bytesForMsgTokens(numTokens int) int {
	// token address (address) + token amount (uint256)
	return (EvmAddressLengthBytes + EvmWordBytes) * numTokens
}

// CalculateMessageMaxGas computes the maximum gas overhead for a message.
func (gp EstimateProvider) CalculateMessageMaxGas(msg ccipocr3.Message) uint64 {
	numTokens := len(msg.TokenAmounts)
	var data []byte = msg.Data
	dataLength := len(data)

	// If msg is an EVM message, its ExtraArgs field is an encoded EVMExtraArgsV2 struct, defined in Client.sol here:
	//nolint:lll
	// https://github.com/smartcontractkit/ccip/blob/dee09c782c17de1e37e3a1cf625d430330532c6d/contracts/src/v0.8/ccip/libraries/Client.sol#L47
	gasLimit, err := decodeEVMExtraArgsV2GasLimit(msg.ExtraArgs)
	if err != nil {
		// TODO: metrics
		gp.lggr.Error("decoding EVMExtraArgsV2 gas limit", err)
	}

	var totalTokenDestGasOverhead uint64
	for _, rampTokenAmount := range msg.TokenAmounts {
		// If msg is an EVM message, its DestExecData field is an ABI encoded uint32, defined in FeeQuoter.sol here:
		//nolint:lll
		// https://github.com/smartcontractkit/ccip/blob/dee09c782c17de1e37e3a1cf625d430330532c6d/contracts/src/v0.8/ccip/FeeQuoter.sol#L958
		// This is the gas overhead for the token transfer.
		tokenDestGasOverhead, err := decodeTokenDestGasOverhead(rampTokenAmount.DestExecData)
		if err != nil {
			// TODO: metrics
			gp.lggr.Error("decoding token dest gas overhead", err)
		}
		totalTokenDestGasOverhead += uint64(tokenDestGasOverhead)
	}

	messageBytes := ConstantMessagePartBytes +
		bytesForMsgTokens(numTokens) +
		dataLength

	messageCallDataGas := uint64(messageBytes * CalldataGasPerByte)

	// Rate limiter only limits value in tokens. It's not called if there are no
	// tokens in the message. The same goes for the admin registry, it's only loaded
	// if there are tokens, and it's only loaded once.
	rateLimiterOverhead := uint64(0)
	adminRegistryOverhead := uint64(0)
	if numTokens >= 1 {
		rateLimiterOverhead = RateLimiterOverheadGas
		adminRegistryOverhead = TokenAdminRegistryWarmupCost
	}

	return messageCallDataGas +
		ExecutionStateProcessingOverheadGas +
		SupportsInterfaceCheck +
		adminRegistryOverhead +
		rateLimiterOverhead +
		PerTokenOverheadGas*uint64(numTokens) +
		gasLimit.Uint64() +
		totalTokenDestGasOverhead
}

func ABITypeOrPanic(t string) abi.Type {
	abiType, err := abi.NewType(t, "", nil)
	if err != nil {
		panic(err)
	}
	return abiType
}

// Decodes EVMExtraArgsV2, see Client.sol to learn how it's encoded:
// https://github.com/smartcontractkit/ccip/blob/dee09c782c17de1e37e3a1cf625d430330532c6d/contracts/src/v0.8/ccip/libraries/Client.sol#L52
//
//nolint:lll
func decodeEVMExtraArgsV2GasLimit(extraArgs []byte) (gasLimit *big.Int, err error) {
	// The first 4 bytes of the extra args are expected to be EVM_EXTRA_ARGS_V2_TAG
	if len(extraArgs) < 4 {
		return big.NewInt(0), fmt.Errorf("expected at least 4 bytes, got %d", len(extraArgs))
	}
	ifaces, err := ExtraArgsV2ABI.UnpackValues(extraArgs[4:])
	if err != nil {
		return big.NewInt(0), fmt.Errorf("abi decode EVMExtraArgsV2: %w", err)
	}
	// gas limit is always the first argument, and allow OOO isn't set explicitly
	// on the message.
	_, ok := ifaces[0].(*big.Int)
	if !ok {
		return big.NewInt(0), fmt.Errorf("expected *big.Int, got %T", ifaces[0])
	}
	return ifaces[0].(*big.Int), nil
}

// Decodes the given bytes into a uint32, based on the encoding of destGasAmount in FeeQuoter.sol here:
// https://github.com/smartcontractkit/ccip/blob/dee09c782c17de1e37e3a1cf625d430330532c6d/contracts/src/v0.8/ccip/FeeQuoter.sol#L958
//
//nolint:lll
func decodeTokenDestGasOverhead(destExecData []byte) (uint32, error) {
	ifaces, err := TokenDestGasOverheadABI.UnpackValues(destExecData)
	if err != nil {
		return 0, fmt.Errorf("abi decode TokenDestGasOverheadABI: %w", err)
	}
	_, ok := ifaces[0].(uint32)
	if !ok {
		return 0, fmt.Errorf("expected uint32, got %T", ifaces[0])
	}
	return ifaces[0].(uint32), nil
}
