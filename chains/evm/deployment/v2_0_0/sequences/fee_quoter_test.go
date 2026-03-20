package sequences_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
)

func TestFeeQuoterUpdate_IsEmpty(t *testing.T) {
	empty := sequences.FeeQuoterUpdate{}
	isEmpty, err := empty.IsEmpty()
	require.NoError(t, err)
	require.True(t, isEmpty, "Empty FeeQuoterUpdate should return true")

	nonEmpty := sequences.FeeQuoterUpdate{
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
		output16 := sequences.FeeQuoterUpdate{}
		output15 := sequences.FeeQuoterUpdate{}
		result, err := sequences.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Equal(t, sequences.FeeQuoterUpdate{}, result)
	})

	t.Run("ConstructorArgs - output15 used when output16 is empty", func(t *testing.T) {
		linkToken := common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA")
		maxFeeJuelsPerMsg := big.NewInt(1000000000000000000) // 1 LINK
		output16 := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{}, // empty
		}
		output15 := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg,
				},
			},
		}
		result, err := sequences.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Equal(t, linkToken, result.ConstructorArgs.StaticConfig.LinkToken)
		require.Equal(t, maxFeeJuelsPerMsg, result.ConstructorArgs.StaticConfig.MaxFeeJuelsPerMsg)
	})

	t.Run("ConstructorArgs - output16 takes precedence when not empty", func(t *testing.T) {
		linkToken16 := common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA")
		linkToken15 := common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")
		maxFeeJuelsPerMsg16 := big.NewInt(2000000000000000000) // 2 LINK
		maxFeeJuelsPerMsg15 := big.NewInt(1000000000000000000) // 1 LINK
		output16 := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken16,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg16,
				},
			},
		}
		output15 := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken15,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg15,
				},
			},
		}
		result, err := sequences.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		// output16's StaticConfig should be used (takes precedence)
		require.Equal(t, linkToken16, result.ConstructorArgs.StaticConfig.LinkToken)
		require.Equal(t, maxFeeJuelsPerMsg16, result.ConstructorArgs.StaticConfig.MaxFeeJuelsPerMsg)
	})

	t.Run("ConstructorArgs - merge DestChainConfig,PriceUpdaters TokenTransferFeeConfigArgs", func(t *testing.T) {
		output16 := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
					MaxFeeJuelsPerMsg: big.NewInt(2000000000000000000), // 2 LINK
				},
				PriceUpdaters: []common.Address{addr1, addr2},
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100,
						TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{
							{Token: addr1},
						},
					},
				},
				DestChainConfigArgs: []fqops.DestChainConfigArgs{
					{
						DestChainSelector: 100,
						DestChainConfig: fqops.DestChainConfig{
							IsEnabled:    true,
							MaxDataBytes: 1000,
						},
					},
				},
			},
		}
		output15 := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
					MaxFeeJuelsPerMsg: big.NewInt(1000000000000000000), // 1 LINK
				},
				PriceUpdaters: []common.Address{addr2, addr3}, // addr2 is duplicate
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100, // duplicate selector
						TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{
							{Token: addr2},
						},
					},
					{
						DestChainSelector: 200, // unique selector
						TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{
							{Token: addr3},
						},
					},
				},
				DestChainConfigArgs: []fqops.DestChainConfigArgs{
					{
						DestChainSelector: 100, // duplicate selector
						DestChainConfig: fqops.DestChainConfig{
							IsEnabled:    false,
							MaxDataBytes: 2000,
						},
					},
					{
						DestChainSelector: 200, // unique selector
						DestChainConfig: fqops.DestChainConfig{
							IsEnabled:    true,
							MaxDataBytes: 3000,
						},
					},
				},
			},
		}
		expected := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
					MaxFeeJuelsPerMsg: big.NewInt(2000000000000000000), // from output16
				},
				PriceUpdaters: []common.Address{addr1, addr2, addr3}, // merged with duplicates removed
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100,
						TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{
							{Token: addr1}, // from output16 (takes precedence)
						},
					},
					{
						DestChainSelector: 200,
						TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{
							{Token: addr3}, // from output15
						},
					},
				},
				DestChainConfigArgs: []fqops.DestChainConfigArgs{
					{
						DestChainSelector: 100,
						DestChainConfig: fqops.DestChainConfig{
							IsEnabled:    true,
							MaxDataBytes: 1000,
						},
					},
					{
						DestChainSelector: 200,
						DestChainConfig: fqops.DestChainConfig{
							IsEnabled:    true,
							MaxDataBytes: 3000,
						},
					},
				},
			},
		}
		result, err := sequences.MergeFeeQuoterUpdateOutputs(output16, output15)
		require.NoError(t, err)
		require.Equal(t, expected.ConstructorArgs.StaticConfig, result.ConstructorArgs.StaticConfig)
		require.ElementsMatch(t, expected.ConstructorArgs.PriceUpdaters, result.ConstructorArgs.PriceUpdaters)
		require.Len(t, result.ConstructorArgs.TokenTransferFeeConfigArgs, 2)
		require.Equal(t, uint64(100), result.ConstructorArgs.TokenTransferFeeConfigArgs[0].DestChainSelector)
		require.Equal(t, addr1, result.ConstructorArgs.TokenTransferFeeConfigArgs[0].TokenTransferFeeConfigs[0].Token) // from output16
		require.Equal(t, uint64(200), result.ConstructorArgs.TokenTransferFeeConfigArgs[1].DestChainSelector)
		require.Equal(t, addr3, result.ConstructorArgs.TokenTransferFeeConfigArgs[1].TokenTransferFeeConfigs[0].Token) // from output15
		require.Len(t, result.ConstructorArgs.DestChainConfigArgs, 2)
		require.Equal(t, uint64(100), result.ConstructorArgs.DestChainConfigArgs[0].DestChainSelector)
		require.True(t, result.ConstructorArgs.DestChainConfigArgs[0].DestChainConfig.IsEnabled)                   // from output16
		require.Equal(t, uint32(1000), result.ConstructorArgs.DestChainConfigArgs[0].DestChainConfig.MaxDataBytes) // from output16
		require.Equal(t, uint64(200), result.ConstructorArgs.DestChainConfigArgs[1].DestChainSelector)
		require.True(t, result.ConstructorArgs.DestChainConfigArgs[1].DestChainConfig.IsEnabled)                   // from output15
		require.Equal(t, uint32(3000), result.ConstructorArgs.DestChainConfigArgs[1].DestChainConfig.MaxDataBytes) // from output15
	})

	t.Run("DestChainConfigs - output16 takes precedence for duplicates", func(t *testing.T) {
		output16 := sequences.FeeQuoterUpdate{
			DestChainConfigs: []fqops.DestChainConfigArgs{
				{
					DestChainSelector: 100,
					DestChainConfig: fqops.DestChainConfig{
						IsEnabled:    true,
						MaxDataBytes: 1000,
					},
				},
			},
		}
		output15 := sequences.FeeQuoterUpdate{
			DestChainConfigs: []fqops.DestChainConfigArgs{
				{
					DestChainSelector: 100, // duplicate selector
					DestChainConfig: fqops.DestChainConfig{
						IsEnabled:    false,
						MaxDataBytes: 2000,
					},
				},
				{
					DestChainSelector: 200, // unique selector
					DestChainConfig: fqops.DestChainConfig{
						IsEnabled:    true,
						MaxDataBytes: 3000,
					},
				},
			},
		}
		result, err := sequences.MergeFeeQuoterUpdateOutputs(output16, output15)
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
		output16 := sequences.FeeQuoterUpdate{
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100,
						TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{
							{Token: addr1},
						},
					},
				},
			},
		}
		output15 := sequences.FeeQuoterUpdate{
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{
						DestChainSelector: 100, // duplicate
						TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{
							{Token: addr2},
						},
					},
					{
						DestChainSelector: 200, // unique
						TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{
							{Token: addr3},
						},
					},
				},
			},
		}
		result, err := sequences.MergeFeeQuoterUpdateOutputs(output16, output15)
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
		output16 := sequences.FeeQuoterUpdate{
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokensToUseDefaultFeeConfigs: []fqops.TokenTransferFeeConfigRemoveArgs{
					{DestChainSelector: 100, Token: addr1},
					{DestChainSelector: 100, Token: addr2},
				},
			},
		}
		output15 := sequences.FeeQuoterUpdate{
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokensToUseDefaultFeeConfigs: []fqops.TokenTransferFeeConfigRemoveArgs{
					{DestChainSelector: 100, Token: addr2}, // duplicate
					{DestChainSelector: 200, Token: addr3}, // unique
				},
			},
		}
		result, err := sequences.MergeFeeQuoterUpdateOutputs(output16, output15)
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
		output16 := sequences.FeeQuoterUpdate{
			AuthorizedCallerUpdates: fqops.AuthorizedCallerArgs{
				AddedCallers:   []common.Address{addr1, addr2},
				RemovedCallers: []common.Address{addr3},
			},
		}
		output15 := sequences.FeeQuoterUpdate{
			AuthorizedCallerUpdates: fqops.AuthorizedCallerArgs{
				AddedCallers:   []common.Address{addr2, addr4}, // addr2 is duplicate
				RemovedCallers: []common.Address{addr3, addr5}, // addr3 is duplicate
			},
		}
		result, err := sequences.MergeFeeQuoterUpdateOutputs(output16, output15)
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
		output16 := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken16,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg16,
				},
			},
			DestChainConfigs: []fqops.DestChainConfigArgs{
				{DestChainSelector: 200, DestChainConfig: fqops.DestChainConfig{IsEnabled: true}},
			},
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{DestChainSelector: 400, TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{{Token: addr1}}},
				},
			},
			AuthorizedCallerUpdates: fqops.AuthorizedCallerArgs{
				AddedCallers: []common.Address{addr1},
			},
		}
		output15 := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				StaticConfig: fqops.StaticConfig{
					LinkToken:         linkToken15,
					MaxFeeJuelsPerMsg: maxFeeJuelsPerMsg15,
				},
			},
			DestChainConfigs: []fqops.DestChainConfigArgs{
				{DestChainSelector: 200, DestChainConfig: fqops.DestChainConfig{IsEnabled: false}}, // duplicate
				{DestChainSelector: 300, DestChainConfig: fqops.DestChainConfig{IsEnabled: true}},  // unique
			},
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
					{DestChainSelector: 400, TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{{Token: addr2}}}, // duplicate
					{DestChainSelector: 500, TokenTransferFeeConfigs: []fqops.TokenTransferFeeConfigSingleTokenArgs{{Token: addr3}}}, // unique
				},
			},
			AuthorizedCallerUpdates: fqops.AuthorizedCallerArgs{
				AddedCallers: []common.Address{addr2, addr3},
			},
		}
		result, err := sequences.MergeFeeQuoterUpdateOutputs(output16, output15)
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

func destChainConfigsFromSelectors(selectors ...uint64) []fqops.DestChainConfigArgs {
	out := make([]fqops.DestChainConfigArgs, len(selectors))
	for i, s := range selectors {
		out[i] = fqops.DestChainConfigArgs{
			DestChainSelector: s,
			DestChainConfig:   fqops.DestChainConfig{IsEnabled: true},
		}
	}
	return out
}

func tokenTransferFeeConfigArgsFromSelectors(selectors ...uint64) []fqops.TokenTransferFeeConfigArgs {
	out := make([]fqops.TokenTransferFeeConfigArgs, len(selectors))
	for i, s := range selectors {
		out[i] = fqops.TokenTransferFeeConfigArgs{DestChainSelector: s}
	}
	return out
}

func TestBatchedInputForSequenceFeeQuoterUpdate(t *testing.T) {
	t.Run("empty input returns nil batches", func(t *testing.T) {
		input := sequences.FeeQuoterUpdate{}
		destBatches, tokenBatches := sequences.BatchedInputForSequenceFeeQuoterUpdate(&input)
		require.Nil(t, destBatches)
		require.Nil(t, tokenBatches)
	})

	t.Run("constructor dest chain configs within batch size stay in constructor and one apply batch", func(t *testing.T) {
		cfgs := destChainConfigsFromSelectors(1, 2, 3)
		input := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				DestChainConfigArgs: append([]fqops.DestChainConfigArgs(nil), cfgs...),
			},
		}
		destBatches, tokenBatches := sequences.BatchedInputForSequenceFeeQuoterUpdate(&input)
		require.Len(t, input.ConstructorArgs.DestChainConfigArgs, 3)
		require.Equal(t, cfgs, input.ConstructorArgs.DestChainConfigArgs)
		require.Empty(t, destBatches)
		require.Nil(t, tokenBatches)
	})

	t.Run("constructor dest chain configs over batch size keeps first batch in constructor rest in dest batches", func(t *testing.T) {
		selectors := make([]uint64, 9)
		for i := range selectors {
			selectors[i] = uint64(i + 1)
		}
		cfgs := destChainConfigsFromSelectors(selectors...)
		input := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				DestChainConfigArgs: append([]fqops.DestChainConfigArgs(nil), cfgs...),
			},
		}
		destBatches, tokenBatches := sequences.BatchedInputForSequenceFeeQuoterUpdate(&input)
		require.Len(t, input.ConstructorArgs.DestChainConfigArgs, 8)
		require.Equal(t, cfgs[:8], input.ConstructorArgs.DestChainConfigArgs)
		require.Len(t, destBatches, 1)
		require.Equal(t, []fqops.DestChainConfigArgs{cfgs[8]}, destBatches[0])
		require.Nil(t, tokenBatches)
	})

	t.Run("update DestChainConfigs only are batched by size 8", func(t *testing.T) {
		selectors := make([]uint64, 10)
		for i := range selectors {
			selectors[i] = uint64(100 + i)
		}
		cfgs := destChainConfigsFromSelectors(selectors...)
		input := sequences.FeeQuoterUpdate{
			DestChainConfigs: append([]fqops.DestChainConfigArgs(nil), cfgs...),
		}
		destBatches, tokenBatches := sequences.BatchedInputForSequenceFeeQuoterUpdate(&input)
		require.Len(t, destBatches, 2)
		require.Equal(t, cfgs[:sequences.DestChainConfigUpdateBatchLen], destBatches[0])
		require.Equal(t, cfgs[sequences.DestChainConfigUpdateBatchLen:], destBatches[1])
		require.Nil(t, tokenBatches)
	})

	t.Run("constructor remainder and update DestChainConfigs are concatenated", func(t *testing.T) {
		// 9 constructor dest configs -> 8 in constructor, 1 in first dest batch
		cons := destChainConfigsFromSelectors(1, 2, 3, 4, 5, 6, 7, 8, 9)
		updates := destChainConfigsFromSelectors(10, 11, 12)
		input := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				DestChainConfigArgs: append([]fqops.DestChainConfigArgs(nil), cons...),
			},
			DestChainConfigs: append([]fqops.DestChainConfigArgs(nil), updates...),
		}
		destBatches, _ := sequences.BatchedInputForSequenceFeeQuoterUpdate(&input)
		require.Len(t, destBatches, 2)
		require.Equal(t, []fqops.DestChainConfigArgs{cons[sequences.DestChainConfigUpdateBatchLen]}, destBatches[0])
		require.Equal(t, updates, destBatches[1])
	})

	t.Run("constructor token transfer fee configs over batch size splits first batch into constructor", func(t *testing.T) {
		selectors := make([]uint64, 6)
		for i := range selectors {
			selectors[i] = uint64(i + 1)
		}
		args := tokenTransferFeeConfigArgsFromSelectors(selectors...)
		input := sequences.FeeQuoterUpdate{
			ConstructorArgs: fqops.ConstructorArgs{
				TokenTransferFeeConfigArgs: append([]fqops.TokenTransferFeeConfigArgs(nil), args...),
			},
		}
		destBatches, tokenBatches := sequences.BatchedInputForSequenceFeeQuoterUpdate(&input)
		require.Nil(t, destBatches)
		require.Len(t, input.ConstructorArgs.TokenTransferFeeConfigArgs, sequences.TokenTransferFeeConfigUpdateBatchLen)
		require.Equal(t, args[:sequences.TokenTransferFeeConfigUpdateBatchLen], input.ConstructorArgs.TokenTransferFeeConfigArgs)
		require.Len(t, tokenBatches, 1)
		require.Equal(t, []fqops.TokenTransferFeeConfigArgs{args[sequences.TokenTransferFeeConfigUpdateBatchLen]}, tokenBatches[0])
	})

	t.Run("TokenTransferFeeConfigUpdates only are batched by size 5", func(t *testing.T) {
		selectors := make([]uint64, 7)
		for i := range selectors {
			selectors[i] = uint64(200 + i)
		}
		args := tokenTransferFeeConfigArgsFromSelectors(selectors...)
		input := sequences.FeeQuoterUpdate{
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: append([]fqops.TokenTransferFeeConfigArgs(nil), args...),
			},
		}
		destBatches, tokenBatches := sequences.BatchedInputForSequenceFeeQuoterUpdate(&input)
		require.Nil(t, destBatches)
		require.Len(t, tokenBatches, 2)
		require.Equal(t, args[:5], tokenBatches[0])
		require.Equal(t, args[5:], tokenBatches[1])
	})

	t.Run("DestChainConfigs and TokenTransferFeeConfigUpdates batched together", func(t *testing.T) {
		destSelectors := make([]uint64, 10)
		for i := range destSelectors {
			destSelectors[i] = uint64(300 + i)
		}
		destCfgs := destChainConfigsFromSelectors(destSelectors...)

		tokenSelectors := make([]uint64, 7)
		for i := range tokenSelectors {
			tokenSelectors[i] = uint64(400 + i)
		}
		tokenArgs := tokenTransferFeeConfigArgsFromSelectors(tokenSelectors...)

		input := sequences.FeeQuoterUpdate{
			DestChainConfigs: append([]fqops.DestChainConfigArgs(nil), destCfgs...),
			TokenTransferFeeConfigUpdates: fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: append([]fqops.TokenTransferFeeConfigArgs(nil), tokenArgs...),
			},
		}
		destBatches, tokenBatches := sequences.BatchedInputForSequenceFeeQuoterUpdate(&input)

		require.Len(t, destBatches, 2)
		require.Equal(t, destCfgs[:sequences.DestChainConfigUpdateBatchLen], destBatches[0])
		require.Equal(t, destCfgs[sequences.DestChainConfigUpdateBatchLen:], destBatches[1])

		require.Len(t, tokenBatches, 2)
		require.Equal(t, tokenArgs[:sequences.TokenTransferFeeConfigUpdateBatchLen], tokenBatches[0])
		require.Equal(t, tokenArgs[sequences.TokenTransferFeeConfigUpdateBatchLen:], tokenBatches[1])
	})
}
