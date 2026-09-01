package sequences_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/stretchr/testify/require"
)

var (
	testTreasury      = common.HexToAddress("0x1111111111111111111111111111111111111111")
	testOtherAddr     = common.HexToAddress("0x2222222222222222222222222222222222222222")
	testOnRampAddr    = common.HexToAddress("0x3333333333333333333333333333333333333333")
	testWETH9Addr     = common.HexToAddress("0x4444444444444444444444444444444444444444")
	testMCMSAddr      = common.HexToAddress("0x5555555555555555555555555555555555555555")
	testChainSelector = uint64(5009297550715157269)
	otherChainSel     = uint64(4340886533089894000)
)

// =============================================================================
// validateTreasuryAddress Tests (Security Critical)
// =============================================================================

func TestValidateTreasuryAddress(t *testing.T) {
	chain := evm.Chain{Selector: testChainSelector}

	tests := []struct {
		desc              string
		allowedRecipients map[uint64]common.Address
		chainSelector     uint64
		address           common.Address
		expectedErr       string
	}{
		{
			desc:              "zero address rejected",
			allowedRecipients: map[uint64]common.Address{testChainSelector: testTreasury},
			chainSelector:     testChainSelector,
			address:           common.Address{},
			expectedErr:       "cannot be zero",
		},
		{
			desc:              "nil map rejected",
			allowedRecipients: nil,
			chainSelector:     testChainSelector,
			address:           testTreasury,
			expectedErr:       "allowedRecipients map is nil",
		},
		{
			desc:              "chain not in allowed list",
			allowedRecipients: map[uint64]common.Address{otherChainSel: testTreasury},
			chainSelector:     testChainSelector,
			address:           testTreasury,
			expectedErr:       "not in the allowed recipients list",
		},
		{
			desc:              "wrong address for chain",
			allowedRecipients: map[uint64]common.Address{testChainSelector: testTreasury},
			chainSelector:     testChainSelector,
			address:           testOtherAddr,
			expectedErr:       "is not the approved treasury",
		},
		{
			desc:              "happy path - address matches allowed",
			allowedRecipients: map[uint64]common.Address{testChainSelector: testTreasury},
			chainSelector:     testChainSelector,
			address:           testTreasury,
			expectedErr:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			input := sequences.ConfigureFeeSweepV150Input{
				ChainSelector:     test.chainSelector,
				OnRampAddress:     testOnRampAddr,
				Treasury:          test.address,
				AllowedRecipients: test.allowedRecipients,
			}

			err := input.Validate(chain)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), test.expectedErr)
			}
		})
	}
}

// =============================================================================
// ConfigureFeeSweepInput Validation Tests
// =============================================================================

func TestConfigureFeeSweepInputValidate(t *testing.T) {
	chain := evm.Chain{Selector: testChainSelector}
	validRecipients := map[uint64]common.Address{testChainSelector: testTreasury}

	tests := []struct {
		desc        string
		makeInput   func() sequences.ConfigureFeeSweepV150Input
		expectedErr string
	}{
		{
			desc: "happy path",
			makeInput: func() sequences.ConfigureFeeSweepV150Input {
				return sequences.ConfigureFeeSweepV150Input{
					ChainSelector:     testChainSelector,
					OnRampAddress:     testOnRampAddr,
					Treasury:          testTreasury,
					AllowedRecipients: validRecipients,
				}
			},
		},
		{
			desc: "chain selector mismatch",
			makeInput: func() sequences.ConfigureFeeSweepV150Input {
				return sequences.ConfigureFeeSweepV150Input{
					ChainSelector:     otherChainSel,
					OnRampAddress:     testOnRampAddr,
					Treasury:          testTreasury,
					AllowedRecipients: map[uint64]common.Address{otherChainSel: testTreasury},
				}
			},
			expectedErr: "chain selector",
		},
		{
			desc: "treasury not in allowed list",
			makeInput: func() sequences.ConfigureFeeSweepV150Input {
				return sequences.ConfigureFeeSweepV150Input{
					ChainSelector:     testChainSelector,
					OnRampAddress:     testOnRampAddr,
					Treasury:          testOtherAddr,
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "is not the approved treasury",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := test.makeInput().Validate(chain)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), test.expectedErr)
			}
		})
	}
}

// =============================================================================
// SweepLinkFeesInput Validation Tests
// =============================================================================

func TestSweepLinkFeesInputValidate(t *testing.T) {
	chain := evm.Chain{Selector: testChainSelector}
	validRecipients := map[uint64]common.Address{testChainSelector: testTreasury}

	tests := []struct {
		desc        string
		makeInput   func() sequences.SweepLinkFeesV150Input
		expectedErr string
	}{
		{
			desc: "happy path",
			makeInput: func() sequences.SweepLinkFeesV150Input {
				return sequences.SweepLinkFeesV150Input{
					ChainSelector:     testChainSelector,
					OnRampAddress:     testOnRampAddr,
					ExpectedTreasury:  testTreasury,
					AllowedRecipients: validRecipients,
				}
			},
		},
		{
			desc: "chain selector mismatch",
			makeInput: func() sequences.SweepLinkFeesV150Input {
				return sequences.SweepLinkFeesV150Input{
					ChainSelector:     otherChainSel,
					OnRampAddress:     testOnRampAddr,
					ExpectedTreasury:  testTreasury,
					AllowedRecipients: map[uint64]common.Address{otherChainSel: testTreasury},
				}
			},
			expectedErr: "chain selector",
		},
		{
			desc: "expected treasury not approved",
			makeInput: func() sequences.SweepLinkFeesV150Input {
				return sequences.SweepLinkFeesV150Input{
					ChainSelector:     testChainSelector,
					OnRampAddress:     testOnRampAddr,
					ExpectedTreasury:  testOtherAddr,
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "expected treasury validation failed",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := test.makeInput().Validate(chain)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), test.expectedErr)
			}
		})
	}
}

// =============================================================================
// SweepNonLinkFeesInput Validation Tests (with WETH auto-detection fields)
// =============================================================================

func TestSweepNonLinkFeesInputValidate(t *testing.T) {
	chain := evm.Chain{Selector: testChainSelector}
	validRecipients := map[uint64]common.Address{testChainSelector: testTreasury}
	feeTokens := []common.Address{testWETH9Addr}

	tests := []struct {
		desc        string
		makeInput   func() sequences.SweepNonLinkFeesV150Input
		expectedErr string
	}{
		{
			desc: "happy path - no WETH",
			makeInput: func() sequences.SweepNonLinkFeesV150Input {
				return sequences.SweepNonLinkFeesV150Input{
					ChainSelector:     testChainSelector,
					OnRampAddress:     testOnRampAddr,
					FeeTokens:         feeTokens,
					Treasury:          testTreasury,
					AllowedRecipients: validRecipients,
				}
			},
		},
		{
			desc: "happy path - with WETH",
			makeInput: func() sequences.SweepNonLinkFeesV150Input {
				return sequences.SweepNonLinkFeesV150Input{
					ChainSelector:     testChainSelector,
					OnRampAddress:     testOnRampAddr,
					FeeTokens:         feeTokens,
					Treasury:          testTreasury,
					WETH9Address:      testWETH9Addr,
					MCMSAddress:       testMCMSAddr,
					WETHBalance:       big.NewInt(1000),
					AllowedRecipients: validRecipients,
				}
			},
		},
		{
			desc: "WETH9 set but MCMSAddress missing",
			makeInput: func() sequences.SweepNonLinkFeesV150Input {
				return sequences.SweepNonLinkFeesV150Input{
					ChainSelector:     testChainSelector,
					OnRampAddress:     testOnRampAddr,
					FeeTokens:         feeTokens,
					Treasury:          testTreasury,
					WETH9Address:      testWETH9Addr,
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "MCMSAddress must be set",
		},
		{
			desc: "chain selector mismatch",
			makeInput: func() sequences.SweepNonLinkFeesV150Input {
				return sequences.SweepNonLinkFeesV150Input{
					ChainSelector:     otherChainSel,
					OnRampAddress:     testOnRampAddr,
					FeeTokens:         feeTokens,
					Treasury:          testTreasury,
					AllowedRecipients: map[uint64]common.Address{otherChainSel: testTreasury},
				}
			},
			expectedErr: "chain selector",
		},
		{
			desc: "treasury not approved",
			makeInput: func() sequences.SweepNonLinkFeesV150Input {
				return sequences.SweepNonLinkFeesV150Input{
					ChainSelector:     testChainSelector,
					OnRampAddress:     testOnRampAddr,
					FeeTokens:         feeTokens,
					Treasury:          testOtherAddr,
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "treasury validation failed",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := test.makeInput().Validate(chain)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), test.expectedErr)
			}
		})
	}
}

// =============================================================================
// SweepAllOnRampsV150Input Validation Tests (Mega Flow)
// =============================================================================

func TestSweepAllOnRampsV150InputValidate(t *testing.T) {
	chain := evm.Chain{Selector: testChainSelector}
	validRecipients := map[uint64]common.Address{testChainSelector: testTreasury}

	tests := []struct {
		desc        string
		makeInput   func() sequences.SweepAllOnRampsV150Input
		expectedErr string
	}{
		{
			desc: "happy path - with WETH",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     testChainSelector,
					Treasury:          testTreasury,
					OnRamps:           []common.Address{testOnRampAddr},
					NonLinkFeeTokens:  []common.Address{testWETH9Addr},
					WETH9Address:      testWETH9Addr,
					MCMSAddress:       testMCMSAddr,
					MinSweepAmount:    big.NewInt(0),
					AllowedRecipients: validRecipients,
				}
			},
		},
		{
			desc: "happy path - no WETH",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     testChainSelector,
					Treasury:          testTreasury,
					OnRamps:           []common.Address{testOnRampAddr},
					NonLinkFeeTokens:  []common.Address{testOtherAddr},
					MinSweepAmount:    big.NewInt(0),
					AllowedRecipients: validRecipients,
				}
			},
		},
		{
			desc: "chain selector mismatch",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     otherChainSel,
					Treasury:          testTreasury,
					OnRamps:           []common.Address{testOnRampAddr},
					MinSweepAmount:    big.NewInt(0),
					AllowedRecipients: map[uint64]common.Address{otherChainSel: testTreasury},
				}
			},
			expectedErr: "chain selector",
		},
		{
			desc: "treasury not approved",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     testChainSelector,
					Treasury:          testOtherAddr,
					OnRamps:           []common.Address{testOnRampAddr},
					MinSweepAmount:    big.NewInt(0),
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "treasury validation failed",
		},
		{
			desc: "empty OnRamps",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     testChainSelector,
					Treasury:          testTreasury,
					OnRamps:           []common.Address{},
					MinSweepAmount:    big.NewInt(0),
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "no OnRamps provided",
		},
		{
			desc: "WETH9 set but MCMSAddress missing",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     testChainSelector,
					Treasury:          testTreasury,
					OnRamps:           []common.Address{testOnRampAddr},
					WETH9Address:      testWETH9Addr,
					MinSweepAmount:    big.NewInt(0),
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "MCMSAddress must be set",
		},
		{
			desc: "nil MinSweepAmount",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     testChainSelector,
					Treasury:          testTreasury,
					OnRamps:           []common.Address{testOnRampAddr},
					MinSweepAmount:    nil,
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "MinSweepAmount must not be nil",
		},
		{
			desc: "nil allowed recipients",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     testChainSelector,
					Treasury:          testTreasury,
					OnRamps:           []common.Address{testOnRampAddr},
					MinSweepAmount:    big.NewInt(0),
					AllowedRecipients: nil,
				}
			},
			expectedErr: "allowedRecipients map is nil",
		},
		{
			desc: "negative MinSweepAmount",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     testChainSelector,
					Treasury:          testTreasury,
					OnRamps:           []common.Address{testOnRampAddr},
					MinSweepAmount:    big.NewInt(-100),
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "MinSweepAmount must be non-negative",
		},
		{
			desc: "zero address in OnRamps",
			makeInput: func() sequences.SweepAllOnRampsV150Input {
				return sequences.SweepAllOnRampsV150Input{
					ChainSelector:     testChainSelector,
					Treasury:          testTreasury,
					OnRamps:           []common.Address{common.Address{}},
					MinSweepAmount:    big.NewInt(0),
					AllowedRecipients: validRecipients,
				}
			},
			expectedErr: "OnRamps[0] cannot be zero address",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := test.makeInput().Validate(chain)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), test.expectedErr)
			}
		})
	}
}

// =============================================================================
// Validation Order Tests
// =============================================================================

func TestSweepAllOnRampsV150InputValidationOrder(t *testing.T) {
	chain := evm.Chain{Selector: testChainSelector}

	// Chain selector checked before treasury
	input := sequences.SweepAllOnRampsV150Input{
		ChainSelector:     otherChainSel,
		Treasury:          common.Address{},
		OnRamps:           nil,
		MinSweepAmount:    nil,
		AllowedRecipients: nil,
	}

	err := input.Validate(chain)
	require.Error(t, err)
	require.Contains(t, err.Error(), "chain selector")
}

// =============================================================================
// Edge Case Tests
// =============================================================================

func TestInputValidationEdgeCases(t *testing.T) {
	chain := evm.Chain{Selector: testChainSelector}
	validRecipients := map[uint64]common.Address{testChainSelector: testTreasury}

	t.Run("multiple chains in allowed recipients", func(t *testing.T) {
		multiChainRecipients := map[uint64]common.Address{
			testChainSelector: testTreasury,
			otherChainSel:     testOtherAddr,
		}

		input := sequences.ConfigureFeeSweepV150Input{
			ChainSelector:     testChainSelector,
			OnRampAddress:     testOnRampAddr,
			Treasury:          testTreasury,
			AllowedRecipients: multiChainRecipients,
		}

		err := input.Validate(chain)
		require.NoError(t, err)
	})

	t.Run("empty allowed recipients map", func(t *testing.T) {
		input := sequences.ConfigureFeeSweepV150Input{
			ChainSelector:     testChainSelector,
			OnRampAddress:     testOnRampAddr,
			Treasury:          testTreasury,
			AllowedRecipients: map[uint64]common.Address{},
		}

		err := input.Validate(chain)
		require.Error(t, err)
		require.Contains(t, err.Error(), "not in the allowed recipients list")
	})

	t.Run("mega flow with zero MinSweepAmount sweeps everything", func(t *testing.T) {
		input := sequences.SweepAllOnRampsV150Input{
			ChainSelector:     testChainSelector,
			Treasury:          testTreasury,
			OnRamps:           []common.Address{testOnRampAddr},
			MinSweepAmount:    big.NewInt(0),
			AllowedRecipients: validRecipients,
		}

		err := input.Validate(chain)
		require.NoError(t, err)
	})

	t.Run("mega flow with SkipNopsCheck true is valid", func(t *testing.T) {
		input := sequences.SweepAllOnRampsV150Input{
			ChainSelector:     testChainSelector,
			Treasury:          testTreasury,
			OnRamps:           []common.Address{testOnRampAddr},
			MinSweepAmount:    big.NewInt(0),
			AllowedRecipients: validRecipients,
			SkipNopsCheck:     true,
		}

		err := input.Validate(chain)
		require.NoError(t, err)
	})

	t.Run("mega flow with SkipNopsCheck false is valid", func(t *testing.T) {
		input := sequences.SweepAllOnRampsV150Input{
			ChainSelector:     testChainSelector,
			Treasury:          testTreasury,
			OnRamps:           []common.Address{testOnRampAddr},
			MinSweepAmount:    big.NewInt(0),
			AllowedRecipients: validRecipients,
			SkipNopsCheck:     false,
		}

		err := input.Validate(chain)
		require.NoError(t, err)
	})

	t.Run("mega flow with high MinSweepAmount is valid", func(t *testing.T) {
		input := sequences.SweepAllOnRampsV150Input{
			ChainSelector:     testChainSelector,
			Treasury:          testTreasury,
			OnRamps:           []common.Address{testOnRampAddr},
			MinSweepAmount:    new(big.Int).SetUint64(1e18),
			AllowedRecipients: validRecipients,
		}

		err := input.Validate(chain)
		require.NoError(t, err)
	})

	t.Run("negative MinSweepAmount rejected", func(t *testing.T) {
		input := sequences.SweepAllOnRampsV150Input{
			ChainSelector:     testChainSelector,
			Treasury:          testTreasury,
			OnRamps:           []common.Address{testOnRampAddr},
			MinSweepAmount:    big.NewInt(-1),
			AllowedRecipients: validRecipients,
		}

		err := input.Validate(chain)
		require.Error(t, err)
		require.Contains(t, err.Error(), "MinSweepAmount must be non-negative")
	})

	t.Run("zero address in OnRamps slice rejected", func(t *testing.T) {
		input := sequences.SweepAllOnRampsV150Input{
			ChainSelector:     testChainSelector,
			Treasury:          testTreasury,
			OnRamps:           []common.Address{testOnRampAddr, {}},
			MinSweepAmount:    big.NewInt(0),
			AllowedRecipients: validRecipients,
		}

		err := input.Validate(chain)
		require.Error(t, err)
		require.Contains(t, err.Error(), "OnRamps[1] cannot be zero address")
	})
}

// =============================================================================
// OnRampAddress Zero-Check Tests
// =============================================================================

func TestOnRampAddressZeroCheck(t *testing.T) {
	chain := evm.Chain{Selector: testChainSelector}
	validRecipients := map[uint64]common.Address{testChainSelector: testTreasury}

	t.Run("ConfigureFeeSweep rejects zero OnRampAddress", func(t *testing.T) {
		input := sequences.ConfigureFeeSweepV150Input{
			ChainSelector:     testChainSelector,
			OnRampAddress:     common.Address{},
			Treasury:          testTreasury,
			AllowedRecipients: validRecipients,
		}

		err := input.Validate(chain)
		require.Error(t, err)
		require.Contains(t, err.Error(), "OnRampAddress cannot be zero address")
	})

	t.Run("SweepLinkFees rejects zero OnRampAddress", func(t *testing.T) {
		input := sequences.SweepLinkFeesV150Input{
			ChainSelector:     testChainSelector,
			OnRampAddress:     common.Address{},
			ExpectedTreasury:  testTreasury,
			AllowedRecipients: validRecipients,
		}

		err := input.Validate(chain)
		require.Error(t, err)
		require.Contains(t, err.Error(), "OnRampAddress cannot be zero address")
	})

	t.Run("SweepNonLinkFees rejects zero OnRampAddress", func(t *testing.T) {
		input := sequences.SweepNonLinkFeesV150Input{
			ChainSelector:     testChainSelector,
			OnRampAddress:     common.Address{},
			FeeTokens:         []common.Address{testWETH9Addr},
			Treasury:          testTreasury,
			AllowedRecipients: validRecipients,
		}

		err := input.Validate(chain)
		require.Error(t, err)
		require.Contains(t, err.Error(), "OnRampAddress cannot be zero address")
	})
}
