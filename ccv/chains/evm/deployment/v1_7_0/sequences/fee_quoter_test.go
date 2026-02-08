package sequences_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
	"github.com/stretchr/testify/require"

	fqops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	sequence1_7 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

func TestFeeQuoterUpdate_IsEmpty(t *testing.T) {
	empty := sequence1_7.FeeQuoterUpdate{}
	isEmpty, err := empty.IsEmpty()
	require.NoError(t, err)
	require.True(t, isEmpty, "Empty FeeQuoterUpdate should return true")

	nonEmpty := sequence1_7.FeeQuoterUpdate{
		ChainSelector: 1,
	}
	isEmpty, err = nonEmpty.IsEmpty()
	require.NoError(t, err)
	require.False(t, isEmpty, "Non-empty FeeQuoterUpdate should return false")
}

func TestMergeFeeQuoterUpdateOutputs(t *testing.T) {
	addr1 := common.HexToAddress("0x1111111111111111111111111111111111111111")
	addr2 := common.HexToAddress("0x2222222222222222222222222222222222222222")
	addr3 := common.HexToAddress("0x3333333333333333333333333333333333333333")
	addr4 := common.HexToAddress("0x4444444444444444444444444444444444444444")
	addr5 := common.HexToAddress("0x5555555555555555555555555555555555555555")

	t.Run("empty outputs", func(t *testing.T) {
		output16 := sequence1_7.FeeQuoterUpdate{}
		output15 := sequence1_7.FeeQuoterUpdate{}
		result, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Equal(t, sequence1_7.FeeQuoterUpdate{}, result)
	})

	t.Run("ConstructorArgs - output15 used when output16 is empty", func(t *testing.T) {
		linkToken := common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA")
		maxFeeJuelsPerMsg := big.NewInt(1000000000000000000) // 1 LINK
		output16 := sequence1_7.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{}, // empty
		}
		output15 := sequence1_7.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg,
				},
			},
		}
		result, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Equal(t, linkToken, result.ConstructorArgs.StaticConfig.LinkToken)
		require.Equal(t, maxFeeJuelsPerMsg, result.ConstructorArgs.StaticConfig.MaxFeeJuelsPerMsg)
	})

	t.Run("ConstructorArgs - output16 takes precedence when not empty", func(t *testing.T) {
		linkToken16 := common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA")
		linkToken15 := common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")
		maxFeeJuelsPerMsg16 := big.NewInt(2000000000000000000) // 2 LINK
		maxFeeJuelsPerMsg15 := big.NewInt(1000000000000000000) // 1 LINK
		output16 := sequence1_7.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken16,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg16,
				},
			},
		}
		output15 := sequence1_7.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken15,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg15,
				},
			},
		}
		result, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		// output16's StaticConfig should be used (takes precedence)
		require.Equal(t, linkToken16, result.ConstructorArgs.StaticConfig.LinkToken)
		require.Equal(t, maxFeeJuelsPerMsg16, result.ConstructorArgs.StaticConfig.MaxFeeJuelsPerMsg)
	})

	t.Run("ConstructorArgs - merge DestChainConfig,PriceUpdaters TokenTransferFeeConfigArgs", func(t *testing.T) {
		output16 := sequence1_7.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
					MaxFeeJuelsPerMsg: big.NewInt(2000000000000000000), // 2 LINK
				},
				PriceUpdaters: []common.Address{addr1, addr2},
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100,
						TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
							{Token: addr1},
						},
					},
				},
				DestChainConfigArgs: []fqops.DestChainConfigArgs{
					{
						DestChainSelector: 100,
						DestChainConfig: adapters.FeeQuoterDestChainConfig{
							IsEnabled:    true,
							MaxDataBytes: 1000,
						},
					},
				},
			},
		}
		output15 := sequence1_7.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
					MaxFeeJuelsPerMsg: big.NewInt(1000000000000000000), // 1 LINK
				},
				PriceUpdaters: []common.Address{addr2, addr3}, // addr2 is duplicate
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100, // duplicate selector
						TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
							{Token: addr2},
						},
					},
					{
						DestChainSelector: 200, // unique selector
						TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
							{Token: addr3},
						},
					},
				},
				DestChainConfigArgs: []fqops.DestChainConfigArgs{
					{
						DestChainSelector: 100, // duplicate selector
						DestChainConfig: adapters.FeeQuoterDestChainConfig{
							IsEnabled:    false,
							MaxDataBytes: 2000,
						},
					},
					{
						DestChainSelector: 200, // unique selector
						DestChainConfig: adapters.FeeQuoterDestChainConfig{
							IsEnabled:    true,
							MaxDataBytes: 3000,
						},
					},
				},
			},
		}
		expected := sequence1_7.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
					MaxFeeJuelsPerMsg: big.NewInt(2000000000000000000), // from output16
				},
				PriceUpdaters: []common.Address{addr1, addr2, addr3}, // merged with duplicates removed
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100,
						TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
							{Token: addr1}, // from output16 (takes precedence)
						},
					},
					{
						DestChainSelector: 200,
						TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
							{Token: addr3}, // from output15
						},
					},
				},
				DestChainConfigArgs: []fqops.DestChainConfigArgs{
					{
						DestChainSelector: 100,
						DestChainConfig: adapters.FeeQuoterDestChainConfig{
							IsEnabled:    true,
							MaxDataBytes: 1000,
						},
					},
					{
						DestChainSelector: 200,
						DestChainConfig: adapters.FeeQuoterDestChainConfig{
							IsEnabled:    true,
							MaxDataBytes: 3000,
						},
					},
				},
			},
		}
		result, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Equal(t, expected.ConstructorArgs, result.ConstructorArgs)
	})

	t.Run("DestChainConfigs - output16 takes precedence for duplicates", func(t *testing.T) {
		output16 := sequence1_7.FeeQuoterUpdate{
			DestChainConfigs: []fqops.DestChainConfigArgs{
				{
					DestChainSelector: 100,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:    true,
						MaxDataBytes: 1000,
					},
				},
			},
		}
		output15 := sequence1_7.FeeQuoterUpdate{
			DestChainConfigs: []fqops.DestChainConfigArgs{
				{
					DestChainSelector: 100, // duplicate selector
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:    false,
						MaxDataBytes: 2000,
					},
				},
				{
					DestChainSelector: 200, // unique selector
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:    true,
						MaxDataBytes: 3000,
					},
				},
			},
		}
		result, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Len(t, result.DestChainConfigs, 2)
		// output16's config for selector 100 should be used
		require.Equal(t, uint64(100), result.DestChainConfigs[0].DestChainSelector)
		require.True(t, result.DestChainConfigs[0].DestChainConfig.IsEnabled)
		require.Equal(t, uint32(1000), result.DestChainConfigs[0].DestChainConfig.MaxDataBytes)
		// output15's config for selector 200 should be added
		require.Equal(t, uint64(200), result.DestChainConfigs[1].DestChainSelector)
		require.Equal(t, uint32(3000), result.DestChainConfigs[1].DestChainConfig.MaxDataBytes)
	})

	t.Run("TokenTransferFeeConfigArgs - output16 takes precedence for duplicates", func(t *testing.T) {
		output16 := sequence1_7.FeeQuoterUpdate{
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100,
						TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
							{Token: addr1},
						},
					},
				},
			},
		}
		output15 := sequence1_7.FeeQuoterUpdate{
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100, // duplicate
						TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
							{Token: addr2},
						},
					},
					{
						DestChainSelector: 200, // unique
						TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
							{Token: addr3},
						},
					},
				},
			},
		}
		result, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Len(t, result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs, 2)
		// output16's config for selector 100 should be used
		require.Equal(t, uint64(100), result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs[0].DestChainSelector)
		require.Equal(t, addr1, result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs[0].TokenTransferFeeConfigs[0].Token)
		// output15's config for selector 200 should be added
		require.Equal(t, uint64(200), result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs[1].DestChainSelector)
		require.Equal(t, addr3, result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs[1].TokenTransferFeeConfigs[0].Token)
	})

	t.Run("TokensToUseDefaultFeeConfigs - merge by DestChainSelector and Token", func(t *testing.T) {
		output16 := sequence1_7.FeeQuoterUpdate{
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokensToUseDefaultFeeConfigs: []fqops.TokenTransferFeeConfigRemoveArgs{
					{DestChainSelector: 100, Token: addr1},
					{DestChainSelector: 100, Token: addr2},
				},
			},
		}
		output15 := sequence1_7.FeeQuoterUpdate{
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokensToUseDefaultFeeConfigs: []fqops.TokenTransferFeeConfigRemoveArgs{
					{DestChainSelector: 100, Token: addr2}, // duplicate
					{DestChainSelector: 200, Token: addr3}, // unique
				},
			},
		}
		result, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Len(t, result.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs, 3)
		// Verify all expected entries are present
		require.Contains(t, result.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs,
			fqops.TokenTransferFeeConfigRemoveArgs{DestChainSelector: 100, Token: addr1})
		require.Contains(t, result.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs,
			fqops.TokenTransferFeeConfigRemoveArgs{DestChainSelector: 100, Token: addr2})
		require.Contains(t, result.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs,
			fqops.TokenTransferFeeConfigRemoveArgs{DestChainSelector: 200, Token: addr3})
	})

	t.Run("AuthorizedCallerUpdates - merge unique entries", func(t *testing.T) {
		output16 := sequence1_7.FeeQuoterUpdate{
			AuthorizedCallerUpdates: fqops.AuthorizedCallerArgs{
				AddedCallers:   []common.Address{addr1, addr2},
				RemovedCallers: []common.Address{addr3},
			},
		}
		output15 := sequence1_7.FeeQuoterUpdate{
			AuthorizedCallerUpdates: fqops.AuthorizedCallerArgs{
				AddedCallers:   []common.Address{addr2, addr4}, // addr2 is duplicate
				RemovedCallers: []common.Address{addr3, addr5}, // addr3 is duplicate
			},
		}
		result, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Len(t, result.AuthorizedCallerUpdates.AddedCallers, 3)
		require.Contains(t, result.AuthorizedCallerUpdates.AddedCallers, addr1)
		require.Contains(t, result.AuthorizedCallerUpdates.AddedCallers, addr2)
		require.Contains(t, result.AuthorizedCallerUpdates.AddedCallers, addr4)
		require.Len(t, result.AuthorizedCallerUpdates.RemovedCallers, 2)
		require.Contains(t, result.AuthorizedCallerUpdates.RemovedCallers, addr3)
		require.Contains(t, result.AuthorizedCallerUpdates.RemovedCallers, addr5)
	})

	t.Run("comprehensive merge", func(t *testing.T) {
		linkToken16 := common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA")
		linkToken15 := common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")
		maxFeeJuelsPerMsg16 := big.NewInt(2000000000000000000) // 2 LINK
		maxFeeJuelsPerMsg15 := big.NewInt(1000000000000000000) // 1 LINK
		output16 := sequence1_7.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken16,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg16,
				},
			},
			DestChainConfigs: []fqops.DestChainConfigArgs{
				{DestChainSelector: 200, DestChainConfig: adapters.FeeQuoterDestChainConfig{IsEnabled: true}},
			},
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{DestChainSelector: 400, TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{{Token: addr1}}},
				},
			},
			AuthorizedCallerUpdates: fqops.AuthorizedCallerArgs{
				AddedCallers: []common.Address{addr1},
			},
		}
		output15 := sequence1_7.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken15,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg15,
				},
			},
			DestChainConfigs: []fqops.DestChainConfigArgs{
				{DestChainSelector: 200, DestChainConfig: adapters.FeeQuoterDestChainConfig{IsEnabled: false}}, // duplicate
				{DestChainSelector: 300, DestChainConfig: adapters.FeeQuoterDestChainConfig{IsEnabled: true}},  // unique
			},
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{DestChainSelector: 400, TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{{Token: addr2}}}, // duplicate
					{DestChainSelector: 500, TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{{Token: addr3}}}, // unique
				},
			},
			AuthorizedCallerUpdates: fqops.AuthorizedCallerArgs{
				AddedCallers: []common.Address{addr2, addr3},
			},
		}
		result, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		// ConstructorArgs from output16 (not empty) - StaticConfig takes precedence
		require.Equal(t, linkToken16, result.ConstructorArgs.StaticConfig.LinkToken)
		require.Equal(t, maxFeeJuelsPerMsg16, result.ConstructorArgs.StaticConfig.MaxFeeJuelsPerMsg)
		// DestChainConfigs: output16's config for 200, plus output15's config for 300
		require.Len(t, result.DestChainConfigs, 2)
		require.True(t, result.DestChainConfigs[0].DestChainConfig.IsEnabled) // from output16
		// TokenTransferFeeConfigArgs: output16's config for 400, plus output15's config for 500
		require.Len(t, result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs, 2)
		require.Equal(t, uint64(400), result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs[0].DestChainSelector)
		require.Equal(t, addr1, result.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs[0].TokenTransferFeeConfigs[0].Token) // from output16
		// AuthorizedCallerUpdates: merged unique entries
		require.Len(t, result.AuthorizedCallerUpdates.AddedCallers, 3)
		require.Contains(t, result.AuthorizedCallerUpdates.AddedCallers, addr1)
		require.Contains(t, result.AuthorizedCallerUpdates.AddedCallers, addr2)
		require.Contains(t, result.AuthorizedCallerUpdates.AddedCallers, addr3)
	})
}
