package weth

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"
)

func TestCreateNativeETHTransfer(t *testing.T) {
	validChainSelector := uint64(5009297550715157269)
	validAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	validAmount := big.NewInt(1000000000000000000) // 1 ETH in wei

	tests := []struct {
		name          string
		chainSelector uint64
		to            common.Address
		amount        *big.Int
		expectedErr   string
		validate      func(t *testing.T, result mcms_types.BatchOperation)
	}{
		{
			name:          "valid transfer",
			chainSelector: validChainSelector,
			to:            validAddress,
			amount:        validAmount,
			validate: func(t *testing.T, result mcms_types.BatchOperation) {
				require.Equal(t, mcms_types.ChainSelector(validChainSelector), result.ChainSelector)
				require.Len(t, result.Transactions, 1)
				tx := result.Transactions[0]
				require.Equal(t, validAddress.Hex(), tx.To)
				require.Equal(t, "NativeETHTransfer", tx.OperationMetadata.ContractType)
				require.Nil(t, tx.Data)

				// Verify AdditionalFields contains the correct value
				var additionalFields map[string]string
				err := json.Unmarshal(tx.AdditionalFields, &additionalFields)
				require.NoError(t, err)
				require.Equal(t, validAmount.String(), additionalFields["value"])
			},
		},
		{
			name:          "zero amount",
			chainSelector: validChainSelector,
			to:            validAddress,
			amount:        big.NewInt(0),
			expectedErr:   "amount must be positive",
		},
		{
			name:          "negative amount",
			chainSelector: validChainSelector,
			to:            validAddress,
			amount:        big.NewInt(-1),
			expectedErr:   "amount must be positive",
		},
		{
			name:          "nil amount",
			chainSelector: validChainSelector,
			to:            validAddress,
			amount:        nil,
			expectedErr:   "amount must be positive",
		},
		{
			name:          "zero address",
			chainSelector: validChainSelector,
			to:            common.Address{},
			amount:        validAmount,
			expectedErr:   "recipient address cannot be zero",
		},
		{
			name:          "large amount",
			chainSelector: validChainSelector,
			to:            validAddress,
			amount:        new(big.Int).Mul(big.NewInt(1000000), big.NewInt(1e18)), // 1M ETH
			validate: func(t *testing.T, result mcms_types.BatchOperation) {
				require.Len(t, result.Transactions, 1)
				var additionalFields map[string]string
				err := json.Unmarshal(result.Transactions[0].AdditionalFields, &additionalFields)
				require.NoError(t, err)
				expectedAmount := new(big.Int).Mul(big.NewInt(1000000), big.NewInt(1e18))
				require.Equal(t, expectedAmount.String(), additionalFields["value"])
			},
		},
		{
			name:          "different chain selector",
			chainSelector: uint64(4340886533089894000), // Different chain
			to:            validAddress,
			amount:        validAmount,
			validate: func(t *testing.T, result mcms_types.BatchOperation) {
				require.Equal(t, mcms_types.ChainSelector(4340886533089894000), result.ChainSelector)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := CreateNativeETHTransfer(tc.chainSelector, tc.to, tc.amount)

			if tc.expectedErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
				return
			}

			require.NoError(t, err)
			if tc.validate != nil {
				tc.validate(t, result)
			}
		})
	}
}

func TestCreateNativeETHTransferTx(t *testing.T) {
	validAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	validAmount := big.NewInt(1000000000000000000) // 1 ETH in wei

	tests := []struct {
		name        string
		to          common.Address
		amount      *big.Int
		expectedErr string
		validate    func(t *testing.T, result mcms_types.Transaction)
	}{
		{
			name:   "valid transfer returns Transaction",
			to:     validAddress,
			amount: validAmount,
			validate: func(t *testing.T, result mcms_types.Transaction) {
				require.Equal(t, validAddress.Hex(), result.To)
				require.Nil(t, result.Data)
			},
		},
		{
			name:   "ContractType is NativeETHTransfer",
			to:     validAddress,
			amount: validAmount,
			validate: func(t *testing.T, result mcms_types.Transaction) {
				require.Equal(t, "NativeETHTransfer", result.OperationMetadata.ContractType)
			},
		},
		{
			name:   "To field is hex-encoded address",
			to:     validAddress,
			amount: validAmount,
			validate: func(t *testing.T, result mcms_types.Transaction) {
				// Verify it's a valid hex address format
				require.Equal(t, validAddress.Hex(), result.To)
				// Should start with 0x
				require.True(t, len(result.To) > 2 && result.To[:2] == "0x")
			},
		},
		{
			name:   "AdditionalFields has correct value encoding",
			to:     validAddress,
			amount: validAmount,
			validate: func(t *testing.T, result mcms_types.Transaction) {
				var additionalFields map[string]string
				err := json.Unmarshal(result.AdditionalFields, &additionalFields)
				require.NoError(t, err)
				require.Contains(t, additionalFields, "value")
				require.Equal(t, validAmount.String(), additionalFields["value"])
			},
		},
		{
			name:        "zero amount",
			to:          validAddress,
			amount:      big.NewInt(0),
			expectedErr: "amount must be positive",
		},
		{
			name:        "nil amount",
			to:          validAddress,
			amount:      nil,
			expectedErr: "amount must be positive",
		},
		{
			name:        "zero address",
			to:          common.Address{},
			amount:      validAmount,
			expectedErr: "recipient address cannot be zero",
		},
		{
			name:        "negative amount",
			to:          validAddress,
			amount:      big.NewInt(-100),
			expectedErr: "amount must be positive",
		},
		{
			name:   "small amount (1 wei)",
			to:     validAddress,
			amount: big.NewInt(1),
			validate: func(t *testing.T, result mcms_types.Transaction) {
				var additionalFields map[string]string
				err := json.Unmarshal(result.AdditionalFields, &additionalFields)
				require.NoError(t, err)
				require.Equal(t, "1", additionalFields["value"])
			},
		},
		{
			name:   "very large amount",
			to:     validAddress,
			amount: new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil), // 10^30 wei
			validate: func(t *testing.T, result mcms_types.Transaction) {
				var additionalFields map[string]string
				err := json.Unmarshal(result.AdditionalFields, &additionalFields)
				require.NoError(t, err)
				expectedAmount := new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil)
				require.Equal(t, expectedAmount.String(), additionalFields["value"])
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := CreateNativeETHTransferTx(tc.to, tc.amount)

			if tc.expectedErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
				return
			}

			require.NoError(t, err)
			if tc.validate != nil {
				tc.validate(t, result)
			}
		})
	}
}

// TestAdditionalFieldsJSONStructure verifies the exact JSON structure of AdditionalFields
func TestAdditionalFieldsJSONStructure(t *testing.T) {
	validAddress := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
	amount := big.NewInt(123456789)

	t.Run("CreateNativeETHTransfer JSON structure", func(t *testing.T) {
		result, err := CreateNativeETHTransfer(12345, validAddress, amount)
		require.NoError(t, err)
		require.Len(t, result.Transactions, 1)

		// Unmarshal and verify structure
		var fields map[string]interface{}
		err = json.Unmarshal(result.Transactions[0].AdditionalFields, &fields)
		require.NoError(t, err)

		// Should only have "value" key
		require.Len(t, fields, 1)
		require.Contains(t, fields, "value")

		// Value should be the string representation of the amount
		require.Equal(t, "123456789", fields["value"])
	})

	t.Run("CreateNativeETHTransferTx JSON structure", func(t *testing.T) {
		result, err := CreateNativeETHTransferTx(validAddress, amount)
		require.NoError(t, err)

		// Unmarshal and verify structure
		var fields map[string]interface{}
		err = json.Unmarshal(result.AdditionalFields, &fields)
		require.NoError(t, err)

		// Should only have "value" key
		require.Len(t, fields, 1)
		require.Contains(t, fields, "value")

		// Value should be the string representation of the amount
		require.Equal(t, "123456789", fields["value"])
	})
}

// TestBatchOperationVsTransaction verifies the difference between the two functions
func TestBatchOperationVsTransaction(t *testing.T) {
	validAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	validAmount := big.NewInt(1e18)
	chainSelector := uint64(5009297550715157269)

	t.Run("CreateNativeETHTransfer returns BatchOperation with one transaction", func(t *testing.T) {
		batchOp, err := CreateNativeETHTransfer(chainSelector, validAddress, validAmount)
		require.NoError(t, err)

		// Verify it's a BatchOperation with exactly one transaction
		require.Equal(t, mcms_types.ChainSelector(chainSelector), batchOp.ChainSelector)
		require.Len(t, batchOp.Transactions, 1)
	})

	t.Run("CreateNativeETHTransferTx returns single Transaction", func(t *testing.T) {
		tx, err := CreateNativeETHTransferTx(validAddress, validAmount)
		require.NoError(t, err)

		// Verify it returns a Transaction (not wrapped in BatchOperation)
		require.Equal(t, validAddress.Hex(), tx.To)
		require.Equal(t, "NativeETHTransfer", tx.OperationMetadata.ContractType)
	})

	t.Run("transactions from both functions have same structure", func(t *testing.T) {
		batchOp, err := CreateNativeETHTransfer(chainSelector, validAddress, validAmount)
		require.NoError(t, err)

		tx, err := CreateNativeETHTransferTx(validAddress, validAmount)
		require.NoError(t, err)

		// The transaction inside BatchOperation should match the standalone Transaction
		batchTx := batchOp.Transactions[0]
		require.Equal(t, batchTx.To, tx.To)
		require.Equal(t, batchTx.Data, tx.Data)
		require.Equal(t, batchTx.OperationMetadata, tx.OperationMetadata)
		require.Equal(t, batchTx.AdditionalFields, tx.AdditionalFields)
	})
}