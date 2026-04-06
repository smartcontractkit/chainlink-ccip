package offchain

import (
	"strconv"
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ptr[T any](v T) *T { return &v }

func evmDefault() FeeQuoterDefaultConfig {
	return FeeQuoterDefaultConfig{
		IsEnabled:                   true,
		MaxDataBytes:                30_000,
		MaxPerMsgGasLimit:           3_000_000,
		DestGasOverhead:             300_000,
		DestGasPerPayloadByteBase:   16,
		DefaultTokenFeeUSDCents:     25,
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		NetworkFeeUSDCents:          10,
		LinkFeeMultiplierPercent:    90,
		USDPerUnitGas:               1_000_000,
	}
}

func TestFeeQuoterDefaultConfig_ApplyOverride(t *testing.T) {
	tests := []struct {
		name     string
		base     FeeQuoterDefaultConfig
		override *FeeQuoterOverrideConfig
		expected FeeQuoterDefaultConfig
	}{
		{
			name:     "nil override returns base unchanged",
			base:     evmDefault(),
			override: nil,
			expected: evmDefault(),
		},
		{
			name:     "empty override returns base unchanged",
			base:     evmDefault(),
			override: &FeeQuoterOverrideConfig{},
			expected: evmDefault(),
		},
		{
			name: "single field override",
			base: evmDefault(),
			override: &FeeQuoterOverrideConfig{
				MaxDataBytes: ptr(uint32(50_000)),
			},
			expected: func() FeeQuoterDefaultConfig {
				c := evmDefault()
				c.MaxDataBytes = 50_000
				return c
			}(),
		},
		{
			name: "multiple fields override",
			base: evmDefault(),
			override: &FeeQuoterOverrideConfig{
				MaxDataBytes:      ptr(uint32(50_000)),
				MaxPerMsgGasLimit: ptr(uint32(5_000_000)),
				IsEnabled:         ptr(false),
				USDPerUnitGas:     ptr(int64(2_000_000)),
			},
			expected: func() FeeQuoterDefaultConfig {
				c := evmDefault()
				c.MaxDataBytes = 50_000
				c.MaxPerMsgGasLimit = 5_000_000
				c.IsEnabled = false
				c.USDPerUnitGas = 2_000_000
				return c
			}(),
		},
		{
			name: "all fields override",
			base: evmDefault(),
			override: &FeeQuoterOverrideConfig{
				IsEnabled:                   ptr(false),
				MaxDataBytes:                ptr(uint32(1)),
				MaxPerMsgGasLimit:           ptr(uint32(2)),
				DestGasOverhead:             ptr(uint32(3)),
				DestGasPerPayloadByteBase:   ptr(uint8(4)),
				DefaultTokenFeeUSDCents:     ptr(uint16(5)),
				DefaultTokenDestGasOverhead: ptr(uint32(6)),
				DefaultTxGasLimit:           ptr(uint32(7)),
				NetworkFeeUSDCents:          ptr(uint16(8)),
				LinkFeeMultiplierPercent:    ptr(uint8(9)),
				USDPerUnitGas:               ptr(int64(10)),
			},
			expected: FeeQuoterDefaultConfig{
				IsEnabled:                   false,
				MaxDataBytes:                1,
				MaxPerMsgGasLimit:           2,
				DestGasOverhead:             3,
				DestGasPerPayloadByteBase:   4,
				DefaultTokenFeeUSDCents:     5,
				DefaultTokenDestGasOverhead: 6,
				DefaultTxGasLimit:           7,
				NetworkFeeUSDCents:          8,
				LinkFeeMultiplierPercent:    9,
				USDPerUnitGas:               10,
			},
		},
		{
			name: "base is not mutated by override",
			base: evmDefault(),
			override: &FeeQuoterOverrideConfig{
				MaxDataBytes: ptr(uint32(99_999)),
			},
			expected: func() FeeQuoterDefaultConfig {
				c := evmDefault()
				c.MaxDataBytes = 99_999
				return c
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			originalBase := tc.base
			result := tc.base.ApplyOverride(tc.override)
			assert.Equal(t, tc.expected, result)
			assert.Equal(t, originalBase, tc.base, "ApplyOverride must not mutate the receiver")
		})
	}
}

func TestFeeQuoterTopology_ResolveFeeQuoterConfigForLane(t *testing.T) {
	evmSelector1 := chainsel.TEST_90000001.Selector
	evmSelector2 := chainsel.TEST_90000002.Selector
	evmSelector3 := chainsel.TEST_90000003.Selector

	tests := []struct {
		name         string
		topology     *FeeQuoterTopology
		srcSelector  uint64
		destSelector uint64
		expected     FeeQuoterDefaultConfig
		expectErr    string
	}{
		{
			name: "family default only",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: evmDefault(),
				},
			},
			srcSelector:  evmSelector1,
			destSelector: evmSelector2,
			expected:     evmDefault(),
		},
		{
			name: "dest chain override applied on top of family default",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: evmDefault(),
				},
				DestChainDefaults: map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector2): {
						MaxDataBytes: ptr(uint32(50_000)),
					},
				},
			},
			srcSelector:  evmSelector1,
			destSelector: evmSelector2,
			expected: func() FeeQuoterDefaultConfig {
				c := evmDefault()
				c.MaxDataBytes = 50_000
				return c
			}(),
		},
		{
			name: "lane override applied on top of dest chain override and family default",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: evmDefault(),
				},
				DestChainDefaults: map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector2): {
						MaxDataBytes: ptr(uint32(50_000)),
					},
				},
				LaneDefaults: map[string]map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector1): {
						selectorKey(evmSelector2): {
							MaxPerMsgGasLimit: ptr(uint32(5_000_000)),
						},
					},
				},
			},
			srcSelector:  evmSelector1,
			destSelector: evmSelector2,
			expected: func() FeeQuoterDefaultConfig {
				c := evmDefault()
				c.MaxDataBytes = 50_000
				c.MaxPerMsgGasLimit = 5_000_000
				return c
			}(),
		},
		{
			name: "source_chain_defaults overrides dest for network and default token fee",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: evmDefault(),
				},
				DestChainDefaults: map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector2): {
						NetworkFeeUSDCents:      ptr(uint16(50)),
						DefaultTokenFeeUSDCents: ptr(uint16(40)),
					},
				},
				SourceChainDefaults: map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector1): {
						NetworkFeeUSDCents:      ptr(uint16(77)),
						DefaultTokenFeeUSDCents: ptr(uint16(88)),
					},
				},
			},
			srcSelector:  evmSelector1,
			destSelector: evmSelector2,
			expected: func() FeeQuoterDefaultConfig {
				c := evmDefault()
				c.NetworkFeeUSDCents = 77
				c.DefaultTokenFeeUSDCents = 88
				return c
			}(),
		},
		{
			name: "lane override wins over source_chain_defaults for same field",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: evmDefault(),
				},
				SourceChainDefaults: map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector1): {
						NetworkFeeUSDCents: ptr(uint16(77)),
					},
				},
				LaneDefaults: map[string]map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector1): {
						selectorKey(evmSelector2): {
							NetworkFeeUSDCents: ptr(uint16(99)),
						},
					},
				},
			},
			srcSelector:  evmSelector1,
			destSelector: evmSelector2,
			expected: func() FeeQuoterDefaultConfig {
				c := evmDefault()
				c.NetworkFeeUSDCents = 99
				return c
			}(),
		},
		{
			name: "lane override without dest chain override",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: evmDefault(),
				},
				LaneDefaults: map[string]map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector1): {
						selectorKey(evmSelector2): {
							NetworkFeeUSDCents: ptr(uint16(99)),
						},
					},
				},
			},
			srcSelector:  evmSelector1,
			destSelector: evmSelector2,
			expected: func() FeeQuoterDefaultConfig {
				c := evmDefault()
				c.NetworkFeeUSDCents = 99
				return c
			}(),
		},
		{
			name: "different src selector does not pick up lane override",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: evmDefault(),
				},
				LaneDefaults: map[string]map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector1): {
						selectorKey(evmSelector2): {
							NetworkFeeUSDCents: ptr(uint16(99)),
						},
					},
				},
			},
			srcSelector:  evmSelector3,
			destSelector: evmSelector2,
			expected:     evmDefault(),
		},
		{
			name: "different src selector does not pick up source_chain_defaults",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: evmDefault(),
				},
				DestChainDefaults: map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector2): {
						NetworkFeeUSDCents: ptr(uint16(50)),
					},
				},
				SourceChainDefaults: map[string]*FeeQuoterOverrideConfig{
					selectorKey(evmSelector1): {
						NetworkFeeUSDCents: ptr(uint16(77)),
					},
				},
			},
			srcSelector:  evmSelector3,
			destSelector: evmSelector2,
			expected: func() FeeQuoterDefaultConfig {
				c := evmDefault()
				c.NetworkFeeUSDCents = 50
				return c
			}(),
		},
		{
			name: "missing family default returns error",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{},
			},
			srcSelector:  evmSelector1,
			destSelector: evmSelector2,
			expectErr:    "no fee quoter family default",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.topology.ResolveFeeQuoterConfigForLane(tc.srcSelector, tc.destSelector)
			if tc.expectErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFeeQuoterTopology_Validate(t *testing.T) {
	validDefault := evmDefault()

	tests := []struct {
		name      string
		topology  *FeeQuoterTopology
		expectErr string
	}{
		{
			name: "valid topology",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: validDefault,
				},
			},
		},
		{
			name: "empty family defaults rejected",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{},
			},
			expectErr: "must contain at least one entry",
		},
		{
			name: "empty family key rejected",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					"": validDefault,
				},
			},
			expectErr: "empty family key",
		},
		{
			name: "zero max_data_bytes rejected",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: {
						MaxPerMsgGasLimit: 1,
						DefaultTxGasLimit: 1,
					},
				},
			},
			expectErr: "max_data_bytes must be > 0",
		},
		{
			name: "zero max_per_msg_gas_limit rejected",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: {
						MaxDataBytes:      1,
						DefaultTxGasLimit: 1,
					},
				},
			},
			expectErr: "max_per_msg_gas_limit must be > 0",
		},
		{
			name: "zero default_tx_gas_limit rejected",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: {
						MaxDataBytes:      1,
						MaxPerMsgGasLimit: 1,
					},
				},
			},
			expectErr: "default_tx_gas_limit must be > 0",
		},
		{
			name: "invalid dest chain selector key rejected",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: validDefault,
				},
				DestChainDefaults: map[string]*FeeQuoterOverrideConfig{
					"not-a-number": {},
				},
			},
			expectErr: "fee_quoter.dest_chain_defaults",
		},
		{
			name: "invalid source chain selector key rejected",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: validDefault,
				},
				SourceChainDefaults: map[string]*FeeQuoterOverrideConfig{
					"not-a-number": {},
				},
			},
			expectErr: "fee_quoter.source_chain_defaults",
		},
		{
			name: "invalid lane source key rejected",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: validDefault,
				},
				LaneDefaults: map[string]map[string]*FeeQuoterOverrideConfig{
					"bad": {
						"12345": {},
					},
				},
			},
			expectErr: "invalid source chain selector key",
		},
		{
			name: "invalid lane dest key rejected",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: validDefault,
				},
				LaneDefaults: map[string]map[string]*FeeQuoterOverrideConfig{
					"12345": {
						"bad": {},
					},
				},
			},
			expectErr: "invalid dest chain selector key",
		},
		{
			name: "valid with all sections",
			topology: &FeeQuoterTopology{
				FamilyDefaults: map[string]FeeQuoterDefaultConfig{
					chainsel.FamilyEVM: validDefault,
				},
				DestChainDefaults: map[string]*FeeQuoterOverrideConfig{
					"12345": {MaxDataBytes: ptr(uint32(50_000))},
				},
				SourceChainDefaults: map[string]*FeeQuoterOverrideConfig{
					"11111": {NetworkFeeUSDCents: ptr(uint16(15))},
				},
				LaneDefaults: map[string]map[string]*FeeQuoterOverrideConfig{
					"12345": {
						"67890": {MaxPerMsgGasLimit: ptr(uint32(5_000_000))},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.topology.Validate()
			if tc.expectErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestEnvironmentTopology_ResolveFeeQuoterConfigForLane_NilFeeQuoterReturnsError(t *testing.T) {
	topo := &EnvironmentTopology{}
	_, err := topo.ResolveFeeQuoterConfigForLane(1, 2)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "fee_quoter topology is not configured")
}

func selectorKey(selector uint64) string {
	return strconv.FormatUint(selector, 10)
}
