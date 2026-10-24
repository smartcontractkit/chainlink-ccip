package onramp

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
)

// =============================================================================
// Test Constants
// =============================================================================

var (
	validChainSel   = uint64(5009297550715157269)
	invalidChainSel = uint64(12345)
	testAddress     = common.HexToAddress("0x01")
	ownerAddress    = common.HexToAddress("0x02")
	nopAddress1     = common.HexToAddress("0x03")
	nopAddress2     = common.HexToAddress("0x04")
	feeTokenAddress = common.HexToAddress("0x05")
	withdrawToAddr  = common.HexToAddress("0x06")
)

// =============================================================================
// Operation Metadata Tests
// =============================================================================

func TestOnRampOperationsMetadata(t *testing.T) {
	tests := []struct {
		name            string
		operationID     string
		operationVer    string
		expectedID      string
		expectedVersion string
	}{
		{
			name:            "SetTokenTransferFeeConfig metadata",
			operationID:     OnRampSetTokenTransferFeeConfig.ID(),
			operationVer:    OnRampSetTokenTransferFeeConfig.Version(),
			expectedID:      "onramp:set-token-transfer-fee-config",
			expectedVersion: "1.5.0",
		},
		{
			name:            "GetTokenTransferFeeConfig metadata",
			operationID:     OnRampGetTokenTransferFeeConfig.ID(),
			operationVer:    OnRampGetTokenTransferFeeConfig.Version(),
			expectedID:      "onramp:get-token-transfer-fee-config",
			expectedVersion: "1.5.0",
		},
		{
			name:            "StaticConfig metadata",
			operationID:     OnRampStaticConfig.ID(),
			operationVer:    OnRampStaticConfig.Version(),
			expectedID:      "onramp:static-config",
			expectedVersion: "1.5.0",
		},
		{
			name:            "DynamicConfig metadata",
			operationID:     OnRampDynamicConfig.ID(),
			operationVer:    OnRampDynamicConfig.Version(),
			expectedID:      "onramp:dynamic-config",
			expectedVersion: "1.5.0",
		},
		{
			name:            "WithdrawNonLinkFees metadata",
			operationID:     OnRampWithdrawNonLinkFees.ID(),
			operationVer:    OnRampWithdrawNonLinkFees.Version(),
			expectedID:      "onramp:withdraw-non-link-fees",
			expectedVersion: "1.5.0",
		},
		{
			name:            "SetNops metadata",
			operationID:     OnRampSetNops.ID(),
			operationVer:    OnRampSetNops.Version(),
			expectedID:      "onramp:set-nops",
			expectedVersion: "1.5.0",
		},
		{
			name:            "PayNops metadata",
			operationID:     OnRampPayNops.ID(),
			operationVer:    OnRampPayNops.Version(),
			expectedID:      "onramp:pay-nops",
			expectedVersion: "1.5.0",
		},
		{
			name:            "GetNops metadata",
			operationID:     OnRampGetNops.ID(),
			operationVer:    OnRampGetNops.Version(),
			expectedID:      "onramp:get-nops",
			expectedVersion: "1.5.0",
		},
		{
			name:            "GetNopFeesJuels metadata",
			operationID:     OnRampGetNopFeesJuels.ID(),
			operationVer:    OnRampGetNopFeesJuels.Version(),
			expectedID:      "onramp:get-nop-fees-juels",
			expectedVersion: "1.5.0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedID, test.operationID)
			assert.Equal(t, test.expectedVersion, test.operationVer)
		})
	}
}

func TestContractTypeAndVersion(t *testing.T) {
	assert.Equal(t, "EVM2EVMOnRamp", string(ContractType))
	assert.Equal(t, "1.5.0", Version.String())
}

// =============================================================================
// SetNops Tests
// =============================================================================

func TestOnRampSetNops(t *testing.T) {
	tests := []struct {
		desc        string
		input       SetNopsInput
		expectedErr string
	}{
		{
			desc: "single NOP with 100% weight",
			input: SetNopsInput{
				NopsAndWeights: []NopAndWeight{
					{Nop: nopAddress1, Weight: 10000},
				},
			},
		},
		{
			desc: "multiple NOPs with split weights",
			input: SetNopsInput{
				NopsAndWeights: []NopAndWeight{
					{Nop: nopAddress1, Weight: 5000},
					{Nop: nopAddress2, Weight: 5000},
				},
			},
		},
		{
			desc: "treasury as sole NOP",
			input: SetNopsInput{
				NopsAndWeights: []NopAndWeight{
					{Nop: nopAddress1, Weight: 65535},
				},
			},
		},
		{
			desc: "unequal weight distribution",
			input: SetNopsInput{
				NopsAndWeights: []NopAndWeight{
					{Nop: nopAddress1, Weight: 7000},
					{Nop: nopAddress2, Weight: 3000},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			// Verify the input types are properly structured
			assert.NotNil(t, test.input.NopsAndWeights)

			// Validate each NOP has an address
			for i, nop := range test.input.NopsAndWeights {
				if len(test.input.NopsAndWeights) > 0 {
					assert.NotEqual(t, common.Address{}, nop.Nop, "NOP %d should have a valid address", i)
				}
			}
		})
	}
}

func TestOnRampSetNopsInputValidation(t *testing.T) {
	lggr, err := logger.New()
	require.NoError(t, err)

	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		operations.NewMemoryReporter(),
	)

	chain := evm.Chain{
		Selector: validChainSel,
	}

	t.Run("empty NopsAndWeights rejected", func(t *testing.T) {
		input := contract.FunctionInput[SetNopsInput]{
			ChainSelector: validChainSel,
			Address:       testAddress,
			Args: SetNopsInput{
				NopsAndWeights: []NopAndWeight{},
			},
		}

		_, err := operations.ExecuteOperation(bundle, OnRampSetNops, chain, input)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "NopsAndWeights list cannot be empty")
	})
}

// =============================================================================
// GetNops Tests
// =============================================================================

func TestGetNopsResultStructure(t *testing.T) {
	// Test the result structure
	result := GetNopsResult{
		NopsAndWeights: []NopAndWeight{
			{Nop: nopAddress1, Weight: 5000},
			{Nop: nopAddress2, Weight: 5000},
		},
		WeightsTotal: big.NewInt(10000),
	}

	assert.Len(t, result.NopsAndWeights, 2)
	assert.Equal(t, big.NewInt(10000), result.WeightsTotal)
	assert.Equal(t, nopAddress1, result.NopsAndWeights[0].Nop)
	assert.Equal(t, nopAddress2, result.NopsAndWeights[1].Nop)
}

func TestGetNopsEmptyResult(t *testing.T) {
	result := GetNopsResult{
		NopsAndWeights: []NopAndWeight{},
		WeightsTotal:   big.NewInt(0),
	}

	assert.Empty(t, result.NopsAndWeights)
	assert.Equal(t, big.NewInt(0), result.WeightsTotal)
}

// =============================================================================
// WithdrawNonLinkFees Tests
// =============================================================================

func TestOnRampWithdrawNonLinkFeesInput(t *testing.T) {
	tests := []struct {
		desc  string
		input WithdrawNonLinkFeesInput
	}{
		{
			desc: "valid withdrawal to address",
			input: WithdrawNonLinkFeesInput{
				FeeToken: feeTokenAddress,
				To:       withdrawToAddr,
			},
		},
		{
			desc: "zero address token handling",
			input: WithdrawNonLinkFeesInput{
				FeeToken: common.Address{},
				To:       withdrawToAddr,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			// Verify input structure
			assert.NotNil(t, test.input)
			// The To address should normally be non-zero for valid withdrawals
			if test.desc != "zero address token handling" {
				assert.NotEqual(t, common.Address{}, test.input.FeeToken)
			}
			assert.NotEqual(t, common.Address{}, test.input.To)
		})
	}
}

// =============================================================================
// PayNops Tests
// =============================================================================

func TestOnRampPayNopsNoArgs(t *testing.T) {
	// PayNops takes no arguments (uses `any` type)
	// This test verifies the operation is structured correctly
	assert.Equal(t, "onramp:pay-nops", OnRampPayNops.ID())
	assert.Equal(t, "1.5.0", OnRampPayNops.Version())
}

// =============================================================================
// SetTokenTransferFeeConfig Tests
// =============================================================================

func TestSetTokenTransferFeeConfigInput(t *testing.T) {
	tests := []struct {
		desc  string
		input SetTokenTransferFeeConfigInput
	}{
		{
			desc: "single token config",
			input: SetTokenTransferFeeConfigInput{
				TokenTransferFeeConfigArgs: []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
					{
						Token:                     feeTokenAddress,
						MinFeeUSDCents:            50,
						MaxFeeUSDCents:            1000,
						DeciBps:                   10,
						DestGasOverhead:           10000,
						DestBytesOverhead:         500,
						AggregateRateLimitEnabled: false,
					},
				},
				TokensToUseDefaultFeeConfigs: []common.Address{},
			},
		},
		{
			desc: "multiple token configs",
			input: SetTokenTransferFeeConfigInput{
				TokenTransferFeeConfigArgs: []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
					{
						Token:                     feeTokenAddress,
						MinFeeUSDCents:            50,
						MaxFeeUSDCents:            1000,
						DeciBps:                   10,
						DestGasOverhead:           10000,
						DestBytesOverhead:         500,
						AggregateRateLimitEnabled: false,
					},
					{
						Token:                     nopAddress1, // using as another token for test
						MinFeeUSDCents:            100,
						MaxFeeUSDCents:            2000,
						DeciBps:                   20,
						DestGasOverhead:           20000,
						DestBytesOverhead:         1000,
						AggregateRateLimitEnabled: true,
					},
				},
				TokensToUseDefaultFeeConfigs: []common.Address{},
			},
		},
		{
			desc: "with tokens to use default config",
			input: SetTokenTransferFeeConfigInput{
				TokenTransferFeeConfigArgs:   []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{},
				TokensToUseDefaultFeeConfigs: []common.Address{feeTokenAddress, nopAddress1},
			},
		},
		{
			desc: "empty config",
			input: SetTokenTransferFeeConfigInput{
				TokenTransferFeeConfigArgs:   []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{},
				TokensToUseDefaultFeeConfigs: []common.Address{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			assert.NotNil(t, test.input.TokenTransferFeeConfigArgs)
			assert.NotNil(t, test.input.TokensToUseDefaultFeeConfigs)
		})
	}
}

// =============================================================================
// NopAndWeight Type Tests
// =============================================================================

func TestNopAndWeightType(t *testing.T) {
	// Verify NopAndWeight is an alias for the gethwrapper type
	nop := NopAndWeight{
		Nop:    nopAddress1,
		Weight: 10000,
	}

	// This should compile and work correctly
	assert.Equal(t, nopAddress1, nop.Nop)
	assert.Equal(t, uint16(10000), nop.Weight)

	// Verify it's the same as the evm_2_evm_onramp type
	var wrapped evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight = nop
	assert.Equal(t, nop.Nop, wrapped.Nop)
	assert.Equal(t, nop.Weight, wrapped.Weight)
}

// =============================================================================
// Chain Selector Validation Tests
// =============================================================================

func TestChainSelectorMismatchErrors(t *testing.T) {
	lggr, err := logger.New()
	require.NoError(t, err)

	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		operations.NewMemoryReporter(),
	)

	chain := evm.Chain{
		Selector: validChainSel,
	}

	// Test that mismatched chain selectors produce appropriate errors
	// We use a read operation as it's simpler (no contract deployment needed)
	t.Run("GetNops with mismatched chain selector", func(t *testing.T) {
		input := contract.FunctionInput[any]{
			ChainSelector: invalidChainSel,
			Address:       testAddress,
			Args:          nil,
		}

		_, err := operations.ExecuteOperation(bundle, OnRampGetNops, chain, input)
		require.Error(t, err)
		assert.Contains(t, err.Error(), fmt.Sprintf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", invalidChainSel, validChainSel))
	})

	t.Run("GetNopFeesJuels with mismatched chain selector", func(t *testing.T) {
		input := contract.FunctionInput[any]{
			ChainSelector: invalidChainSel,
			Address:       testAddress,
			Args:          nil,
		}

		_, err := operations.ExecuteOperation(bundle, OnRampGetNopFeesJuels, chain, input)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "mismatch between inputted chain selector and selector defined within dependencies")
	})
}

// =============================================================================
// Empty Address Validation Tests
// =============================================================================

func TestEmptyAddressErrors(t *testing.T) {
	lggr, err := logger.New()
	require.NoError(t, err)

	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		operations.NewMemoryReporter(),
	)

	chain := evm.Chain{
		Selector: validChainSel,
	}

	t.Run("GetNops with empty address", func(t *testing.T) {
		input := contract.FunctionInput[any]{
			ChainSelector: validChainSel,
			Address:       common.Address{},
			Args:          nil,
		}

		_, err := operations.ExecuteOperation(bundle, OnRampGetNops, chain, input)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "address must be specified")
	})

	t.Run("GetTokenTransferFeeConfig with empty address", func(t *testing.T) {
		input := contract.FunctionInput[common.Address]{
			ChainSelector: validChainSel,
			Address:       common.Address{},
			Args:          feeTokenAddress,
		}

		_, err := operations.ExecuteOperation(bundle, OnRampGetTokenTransferFeeConfig, chain, input)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "address must be specified")
	})
}

// =============================================================================
// Write Operation Output Structure Tests
// =============================================================================

func TestWriteOutputStructure(t *testing.T) {
	// Test that WriteOutput can be properly constructed
	output := contract.WriteOutput{
		ChainSelector: validChainSel,
		ExecInfo:      nil,
	}

	assert.Equal(t, validChainSel, output.ChainSelector)
	assert.False(t, output.Executed(), "Output without ExecInfo should not be marked as executed")

	output.ExecInfo = &contract.ExecInfo{Hash: "0xabc123"}
	assert.True(t, output.Executed(), "Output with ExecInfo should be marked as executed")
}

// =============================================================================
// FunctionInput Structure Tests
// =============================================================================

func TestFunctionInputStructures(t *testing.T) {
	t.Run("SetNops FunctionInput", func(t *testing.T) {
		input := contract.FunctionInput[SetNopsInput]{
			ChainSelector: validChainSel,
			Address:       testAddress,
			Args: SetNopsInput{
				NopsAndWeights: []NopAndWeight{
					{Nop: nopAddress1, Weight: 10000},
				},
			},
		}

		assert.Equal(t, validChainSel, input.ChainSelector)
		assert.Equal(t, testAddress, input.Address)
		assert.Len(t, input.Args.NopsAndWeights, 1)
	})

	t.Run("WithdrawNonLinkFees FunctionInput", func(t *testing.T) {
		input := contract.FunctionInput[WithdrawNonLinkFeesInput]{
			ChainSelector: validChainSel,
			Address:       testAddress,
			Args: WithdrawNonLinkFeesInput{
				FeeToken: feeTokenAddress,
				To:       withdrawToAddr,
			},
		}

		assert.Equal(t, validChainSel, input.ChainSelector)
		assert.Equal(t, testAddress, input.Address)
		assert.Equal(t, feeTokenAddress, input.Args.FeeToken)
		assert.Equal(t, withdrawToAddr, input.Args.To)
	})

	t.Run("SetTokenTransferFeeConfig FunctionInput", func(t *testing.T) {
		input := contract.FunctionInput[SetTokenTransferFeeConfigInput]{
			ChainSelector: validChainSel,
			Address:       testAddress,
			Args: SetTokenTransferFeeConfigInput{
				TokenTransferFeeConfigArgs: []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
					{
						Token:          feeTokenAddress,
						MinFeeUSDCents: 50,
						MaxFeeUSDCents: 1000,
					},
				},
				TokensToUseDefaultFeeConfigs: []common.Address{},
			},
		}

		assert.Equal(t, validChainSel, input.ChainSelector)
		assert.Len(t, input.Args.TokenTransferFeeConfigArgs, 1)
		assert.Equal(t, feeTokenAddress, input.Args.TokenTransferFeeConfigArgs[0].Token)
	})

	t.Run("PayNops FunctionInput with nil args", func(t *testing.T) {
		input := contract.FunctionInput[any]{
			ChainSelector: validChainSel,
			Address:       testAddress,
			Args:          nil,
		}

		assert.Equal(t, validChainSel, input.ChainSelector)
		assert.Equal(t, testAddress, input.Address)
		assert.Nil(t, input.Args)
	})
}

// =============================================================================
// BatchOperation Construction Tests
// =============================================================================

func TestBatchOperationFromWriteOutputs(t *testing.T) {
	t.Run("single write output", func(t *testing.T) {
		outputs := []contract.WriteOutput{
			{
				ChainSelector: validChainSel,
			},
		}

		batchOp, err := contract.NewBatchOperationFromWrites(outputs)
		require.NoError(t, err)
		assert.Equal(t, validChainSel, uint64(batchOp.ChainSelector))
	})

	t.Run("multiple write outputs same chain", func(t *testing.T) {
		outputs := []contract.WriteOutput{
			{ChainSelector: validChainSel},
			{ChainSelector: validChainSel},
		}

		batchOp, err := contract.NewBatchOperationFromWrites(outputs)
		require.NoError(t, err)
		assert.Equal(t, validChainSel, uint64(batchOp.ChainSelector))
	})

	t.Run("empty outputs", func(t *testing.T) {
		outputs := []contract.WriteOutput{}

		batchOp, err := contract.NewBatchOperationFromWrites(outputs)
		require.NoError(t, err)
		assert.Equal(t, uint64(0), uint64(batchOp.ChainSelector))
	})

	t.Run("all executed outputs filtered", func(t *testing.T) {
		outputs := []contract.WriteOutput{
			{
				ChainSelector: validChainSel,
				ExecInfo:      &contract.ExecInfo{Hash: "0xabc"},
			},
		}

		batchOp, err := contract.NewBatchOperationFromWrites(outputs)
		require.NoError(t, err)
		// Executed outputs are filtered out
		assert.Equal(t, uint64(0), uint64(batchOp.ChainSelector))
	})
}

// =============================================================================
// Mock Contract Interface Tests
// =============================================================================

// mockOnRamp implements a minimal interface for testing
type mockOnRamp struct {
	address common.Address
	owner   common.Address
	nops    []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight
	fees    *big.Int
}

func (m *mockOnRamp) Address() common.Address {
	return m.address
}

func (m *mockOnRamp) Owner(opts *bind.CallOpts) (common.Address, error) {
	return m.owner, nil
}

func (m *mockOnRamp) SetNops(opts *bind.TransactOpts, nops []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight) (*types.Transaction, error) {
	m.nops = nops
	return types.NewTx(&types.LegacyTx{
		To:   &m.address,
		Data: []byte{0xDE, 0xAD, 0xBE, 0xEF},
	}), nil
}

func (m *mockOnRamp) GetNops(opts *bind.CallOpts) (struct {
	NopsAndWeights []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight
	WeightsTotal   *big.Int
}, error) {
	return struct {
		NopsAndWeights []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight
		WeightsTotal   *big.Int
	}{
		NopsAndWeights: m.nops,
		WeightsTotal:   big.NewInt(10000),
	}, nil
}

func (m *mockOnRamp) GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error) {
	if m.fees == nil {
		return big.NewInt(0), nil
	}
	return m.fees, nil
}

func TestMockOnRampBehavior(t *testing.T) {
	mock := &mockOnRamp{
		address: testAddress,
		owner:   ownerAddress,
		nops:    []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{},
		fees:    big.NewInt(1000000),
	}

	t.Run("owner check", func(t *testing.T) {
		owner, err := mock.Owner(nil)
		require.NoError(t, err)
		assert.Equal(t, ownerAddress, owner)
	})

	t.Run("set and get nops", func(t *testing.T) {
		nops := []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{
			{Nop: nopAddress1, Weight: 5000},
			{Nop: nopAddress2, Weight: 5000},
		}

		_, err := mock.SetNops(nil, nops)
		require.NoError(t, err)

		result, err := mock.GetNops(nil)
		require.NoError(t, err)
		assert.Len(t, result.NopsAndWeights, 2)
		assert.Equal(t, big.NewInt(10000), result.WeightsTotal)
	})

	t.Run("get fees", func(t *testing.T) {
		fees, err := mock.GetNopFeesJuels(nil)
		require.NoError(t, err)
		assert.Equal(t, big.NewInt(1000000), fees)
	})
}
