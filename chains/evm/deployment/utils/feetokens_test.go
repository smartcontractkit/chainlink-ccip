package utils_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	evm_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
)

const testChainSelector = uint64(909606746561742123)

func TestEVMFeeTokenAddresses(t *testing.T) {
	linkAddr := "0x779877A7B0D9E8603169DdbD7836e478b4624789"
	wethAddr := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

	t.Run("converts multiple tokens preserving order", func(t *testing.T) {
		got, err := evm_utils.EVMFeeTokenAddresses([]fees.FeeTokenWithdrawal{
			{Token: linkAddr},
			{Token: wethAddr},
		}, testChainSelector)
		require.NoError(t, err)
		require.Equal(t, []common.Address{
			common.HexToAddress(linkAddr),
			common.HexToAddress(wethAddr),
		}, got)
	})

	t.Run("accepts non-checksummed addresses", func(t *testing.T) {
		got, err := evm_utils.EVMFeeTokenAddresses([]fees.FeeTokenWithdrawal{
			{Token: "0x779877a7b0d9e8603169ddbd7836e478b4624789"},
		}, testChainSelector)
		require.NoError(t, err)
		require.Equal(t, []common.Address{common.HexToAddress(linkAddr)}, got)
	})

	t.Run("rejects an empty token list", func(t *testing.T) {
		_, err := evm_utils.EVMFeeTokenAddresses(nil, testChainSelector)
		require.Error(t, err)
		require.Contains(t, err.Error(), "no fee tokens specified")
	})

	t.Run("rejects a malformed address", func(t *testing.T) {
		_, err := evm_utils.EVMFeeTokenAddresses([]fees.FeeTokenWithdrawal{
			{Token: linkAddr},
			{Token: "not-an-address"},
		}, testChainSelector)
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid fee token address")
		require.Contains(t, err.Error(), "feeTokens[1]")
	})

	// withdrawFeeTokens always sweeps the full balance, so honouring an amount is impossible.
	// Silently withdrawing everything would move more than the operator asked for.
	t.Run("rejects a per-token amount", func(t *testing.T) {
		_, err := evm_utils.EVMFeeTokenAddresses([]fees.FeeTokenWithdrawal{
			{Token: linkAddr, Amount: big.NewInt(100)},
		}, testChainSelector)
		require.Error(t, err)
		require.Contains(t, err.Error(), "partial withdrawals are not supported on EVM")
	})
}
