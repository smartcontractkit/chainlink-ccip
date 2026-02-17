package sequences_test

import (
	"math/big"
	"sort"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	evmadapter "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"

	fqops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	sequence1_7 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	fee_quoter_v1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	evm_2_evm_onramp_v1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"

	seq1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	seq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
)

// dummyAddressRefs is hardcoded address refs (previously from address_refs.json).
// Chain selectors must match dummyContractMetadata so metadata lookup succeeds.
var dummyAddressRefs = []datastore.AddressRef{
	{Address: "0x1111111111111111111111111111111111111111", ChainSelector: 5009297550715157269, Type: datastore.ContractType("FeeQuoter"), Version: semver.MustParse("1.6.3")},
	{Address: "0x6666666666666666666666666666666666666666", ChainSelector: 5009297550715157269, Type: datastore.ContractType("EVM2EVMOnRamp"), Version: semver.MustParse("1.5.0")},
	{Address: "0x2222222222222222222222222222222222222221", ChainSelector: 5009297550715157269, Type: datastore.ContractType("CommitStore"), Version: semver.MustParse("1.5.0")},
	{Address: "0x9999999999999999999999999999999999999999", ChainSelector: 4949039107694359620, Type: datastore.ContractType("CommitStore"), Version: semver.MustParse("1.5.0"), Qualifier: "commitstore1"},
	{Address: "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", ChainSelector: 4949039107694359620, Type: datastore.ContractType("FeeQuoter"), Version: semver.MustParse("1.6.3")},
	{Address: "0x1010101010101010101010101010101010101010", ChainSelector: 4949039107694359620, Type: datastore.ContractType("EVM2EVMOnRamp"), Version: semver.MustParse("1.5.0")},
	{Address: "0x3333333333333333333333333333333333333333", ChainSelector: 4949039107694359620, Type: datastore.ContractType("CommitStore"), Version: semver.MustParse("1.5.0"), Qualifier: "commitstore2"},
	{Address: "0x5050505050505050505050505050505050505050", ChainSelector: 15971525489660198786, Type: datastore.ContractType("EVM2EVMOnRamp"), Version: semver.MustParse("1.5.0")},
	{Address: "0x4444444444444444444444444444444444444444", ChainSelector: 15971525489660198786, Type: datastore.ContractType("CommitStore"), Version: semver.MustParse("1.5.0")},
	{Address: "0x6060606060606060606060606060606060606060", ChainSelector: 5936861837188149645, Type: datastore.ContractType("FeeQuoter"), Version: semver.MustParse("1.6.3")},
	{Address: "0x7070707070707070707070707070707070707070", ChainSelector: 5936861837188149645, Type: datastore.ContractType("EVM2EVMOnRamp"), Version: semver.MustParse("1.5.0")},
	{Address: "0x5555555555555555555555555555555555555551", ChainSelector: 5936861837188149645, Type: datastore.ContractType("CommitStore"), Version: semver.MustParse("1.5.0")},
}

var dummyContractMetadata = []datastore.ContractMetadata{
	{
		Address:       "0x1111111111111111111111111111111111111111",
		ChainSelector: 5009297550715157269,
		Metadata: seq1_6.FeeQuoterImportConfigSequenceOutput{
			RemoteChainCfgs: map[uint64]seq1_6.FeeQuoterImportConfigSequenceOutputPerRemoteChain{
				15971525489660198786: {
					DestChainCfg: fee_quoter_v1_6_0.DestChainConfig{
						IsEnabled:                         true,
						MaxNumberOfTokensPerMsg:           3,
						MaxDataBytes:                      8000,
						MaxPerMsgGasLimit:                 4000000,
						DestGasOverhead:                   80000,
						DestGasPerPayloadByteBase:         14,
						DestGasPerPayloadByteHigh:         28,
						DestGasPerPayloadByteThreshold:    800,
						DestDataAvailabilityOverheadGas:   40000,
						DestGasPerDataAvailabilityByte:    8,
						DestDataAvailabilityMultiplierBps: 900,
						ChainFamilySelector:               [4]byte(utils.GetSelectorHex(15971525489660198786)),
						EnforceOutOfOrder:                 false,
						DefaultTokenFeeUSDCents:           8,
						DefaultTokenDestGasOverhead:       40000,
						DefaultTxGasLimit:                 180000,
						GasMultiplierWeiPerEth:            0,
						GasPriceStalenessThreshold:        0,
						NetworkFeeUSDCents:                10,
					},
					TokenTransferFeeCfgs: map[common.Address]fee_quoter_v1_6_0.TokenTransferFeeConfig{
						common.HexToAddress("0x2222222222222222222222222222222222222222"): {
							MinFeeUSDCents:    4,
							MaxFeeUSDCents:    40,
							DeciBps:           90,
							DestGasOverhead:   25000,
							DestBytesOverhead: 80,
							IsEnabled:         true,
						},
					},
				},
			},
			StaticCfg: fee_quoter_v1_6_0.StaticConfig{
				MaxFeeJuelsPerMsg:            big.NewInt(1000000000000000000),
				LinkToken:                    common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
				TokenPriceStalenessThreshold: 3600,
			},
			PriceUpdaters: []common.Address{
				common.HexToAddress("0x4444444444444444444444444444444444444444"),
				common.HexToAddress("0x5555555555555555555555555555555555555555"),
			},
		},
	},
	{
		Address:       "0x6666666666666666666666666666666666666666",
		ChainSelector: 5009297550715157269,
		Metadata: seq1_5.OnRampImportConfigSequenceOutput{
			RemoteChainSelector: 4949039107694359620,
			StaticConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampStaticConfig{
				LinkToken:          common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
				ChainSelector:      5009297550715157269,
				DestChainSelector:  4949039107694359620,
				DefaultTxGasLimit:  200000,
				MaxNopFeesJuels:    big.NewInt(1000000000000000000),
				PrevOnRamp:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
				RmnProxy:           common.HexToAddress("0x7777777777777777777777777777777777777777"),
				TokenAdminRegistry: common.HexToAddress("0x8888888888888888888888888888888888888888"),
			},
			DynamicConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampDynamicConfig{
				Router:                            common.HexToAddress("0x9999999999999999999999999999999999999999"),
				MaxNumberOfTokensPerMsg:           5,
				DestGasOverhead:                   100000,
				DestGasPerPayloadByte:             16,
				DestDataAvailabilityOverheadGas:   50000,
				DestGasPerDataAvailabilityByte:    10,
				DestDataAvailabilityMultiplierBps: 1000,
				PriceRegistry:                     common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
				MaxDataBytes:                      10000,
				MaxPerMsgGasLimit:                 5000000,
				DefaultTokenFeeUSDCents:           0,
				DefaultTokenDestGasOverhead:       0,
				EnforceOutOfOrder:                 false,
			},
			TokenTransferFeeConfig: map[common.Address]evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampTokenTransferFeeConfig{
				common.HexToAddress("0x2222222222222222222222222222222222222222"): {
					MinFeeUSDCents:            5,
					MaxFeeUSDCents:            50,
					DeciBps:                   100,
					DestGasOverhead:           30000,
					DestBytesOverhead:         100,
					AggregateRateLimitEnabled: false,
					IsEnabled:                 true,
				},
				common.HexToAddress("0x3333333333333333333333333333333333333333"): {
					MinFeeUSDCents:            10,
					MaxFeeUSDCents:            100,
					DeciBps:                   200,
					DestGasOverhead:           40000,
					DestBytesOverhead:         200,
					AggregateRateLimitEnabled: false,
					IsEnabled:                 true,
				},
			},
			FeeTokenConfig: map[common.Address]evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampFeeTokenConfig{
				common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"): {
					NetworkFeeUSDCents:         10,
					GasMultiplierWeiPerEth:     1000000000000000000,
					PremiumMultiplierWeiPerEth: 1100000000000000000,
					Enabled:                    true,
				},
				common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"): {
					NetworkFeeUSDCents:         10,
					GasMultiplierWeiPerEth:     900000000000000000,
					PremiumMultiplierWeiPerEth: 1000000000000000000,
					Enabled:                    true,
				},
			},
		},
	},
	{
		Address:       "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		ChainSelector: 4949039107694359620,
		Metadata: seq1_6.FeeQuoterImportConfigSequenceOutput{
			RemoteChainCfgs: map[uint64]seq1_6.FeeQuoterImportConfigSequenceOutputPerRemoteChain{
				15971525489660198786: {
					DestChainCfg: fee_quoter_v1_6_0.DestChainConfig{
						IsEnabled:                         true,
						MaxNumberOfTokensPerMsg:           3,
						MaxDataBytes:                      8000,
						MaxPerMsgGasLimit:                 4000000,
						DestGasOverhead:                   80000,
						DestGasPerPayloadByteBase:         14,
						DestGasPerPayloadByteHigh:         28,
						DestGasPerPayloadByteThreshold:    800,
						DestDataAvailabilityOverheadGas:   40000,
						DestGasPerDataAvailabilityByte:    8,
						DestDataAvailabilityMultiplierBps: 900,
						ChainFamilySelector:               [4]byte(utils.GetSelectorHex(15971525489660198786)),
						EnforceOutOfOrder:                 false,
						DefaultTokenFeeUSDCents:           8,
						DefaultTokenDestGasOverhead:       40000,
						DefaultTxGasLimit:                 180000,
						GasMultiplierWeiPerEth:            0,
						GasPriceStalenessThreshold:        0,
						NetworkFeeUSDCents:                10,
					},
					TokenTransferFeeCfgs: map[common.Address]fee_quoter_v1_6_0.TokenTransferFeeConfig{
						common.HexToAddress("0x2222222222222222222222222222222222222222"): {
							MinFeeUSDCents:    4,
							MaxFeeUSDCents:    40,
							DeciBps:           90,
							DestGasOverhead:   25000,
							DestBytesOverhead: 80,
							IsEnabled:         true,
						},
					},
				},
			},
			StaticCfg: fee_quoter_v1_6_0.StaticConfig{
				MaxFeeJuelsPerMsg:            big.NewInt(1000000000000000000),
				LinkToken:                    common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
				TokenPriceStalenessThreshold: 3600,
			},
			PriceUpdaters: []common.Address{
				common.HexToAddress("0x4444444444444444444444444444444444444444"),
				common.HexToAddress("0x5555555555555555555555555555555555555555"),
			},
		},
	},
	{
		Address:       "0x1010101010101010101010101010101010101010",
		ChainSelector: 4949039107694359620,
		Metadata: seq1_5.OnRampImportConfigSequenceOutput{
			RemoteChainSelector: 5009297550715157269,
			StaticConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampStaticConfig{
				LinkToken:          common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
				ChainSelector:      4949039107694359620,
				DestChainSelector:  5009297550715157269,
				DefaultTxGasLimit:  200000,
				MaxNopFeesJuels:    big.NewInt(1000000000000000000),
				PrevOnRamp:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
				RmnProxy:           common.HexToAddress("0x7777777777777777777777777777777777777777"),
				TokenAdminRegistry: common.HexToAddress("0x8888888888888888888888888888888888888888"),
			},
			DynamicConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampDynamicConfig{
				Router:                            common.HexToAddress("0x9999999999999999999999999999999999999999"),
				MaxNumberOfTokensPerMsg:           5,
				DestGasOverhead:                   100000,
				DestGasPerPayloadByte:             16,
				DestDataAvailabilityOverheadGas:   50000,
				DestGasPerDataAvailabilityByte:    10,
				DestDataAvailabilityMultiplierBps: 1000,
				PriceRegistry:                     common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
				MaxDataBytes:                      10000,
				MaxPerMsgGasLimit:                 5000000,
				DefaultTokenFeeUSDCents:           0,
				DefaultTokenDestGasOverhead:       0,
				EnforceOutOfOrder:                 false,
			},
			TokenTransferFeeConfig: map[common.Address]evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampTokenTransferFeeConfig{
				common.HexToAddress("0x2222222222222222222222222222222222222222"): {
					MinFeeUSDCents:            5,
					MaxFeeUSDCents:            50,
					DeciBps:                   100,
					DestGasOverhead:           30000,
					DestBytesOverhead:         100,
					AggregateRateLimitEnabled: false,
					IsEnabled:                 true,
				},
			},
			FeeTokenConfig: map[common.Address]evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampFeeTokenConfig{
				common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"): {
					NetworkFeeUSDCents:         10,
					GasMultiplierWeiPerEth:     1000000000000000000,
					PremiumMultiplierWeiPerEth: 1100000000000000000,
					Enabled:                    true,
				},
			},
		},
	},
	{
		Address:       "0x5050505050505050505050505050505050505050",
		ChainSelector: 15971525489660198786,
		Metadata: seq1_5.OnRampImportConfigSequenceOutput{
			RemoteChainSelector: 5009297550715157269,
			StaticConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampStaticConfig{
				LinkToken:          common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
				ChainSelector:      15971525489660198786,
				DestChainSelector:  5009297550715157269,
				DefaultTxGasLimit:  200000,
				MaxNopFeesJuels:    big.NewInt(1000000000000000000),
				PrevOnRamp:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
				RmnProxy:           common.HexToAddress("0x7777777777777777777777777777777777777777"),
				TokenAdminRegistry: common.HexToAddress("0x8888888888888888888888888888888888888888"),
			},
			DynamicConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampDynamicConfig{
				Router:                            common.HexToAddress("0x9999999999999999999999999999999999999999"),
				MaxNumberOfTokensPerMsg:           5,
				DestGasOverhead:                   100000,
				DestGasPerPayloadByte:             16,
				DestDataAvailabilityOverheadGas:   50000,
				DestGasPerDataAvailabilityByte:    10,
				DestDataAvailabilityMultiplierBps: 1000,
				PriceRegistry:                     common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
				MaxDataBytes:                      10000,
				MaxPerMsgGasLimit:                 5000000,
				DefaultTokenFeeUSDCents:           0,
				DefaultTokenDestGasOverhead:       0,
				EnforceOutOfOrder:                 false,
			},
			TokenTransferFeeConfig: map[common.Address]evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampTokenTransferFeeConfig{
				common.HexToAddress("0x2222222222222222222222222222222222222222"): {
					MinFeeUSDCents:            5,
					MaxFeeUSDCents:            50,
					DeciBps:                   100,
					DestGasOverhead:           30000,
					DestBytesOverhead:         100,
					AggregateRateLimitEnabled: false,
					IsEnabled:                 true,
				},
				common.HexToAddress("0x3333333333333333333333333333333333333333"): {
					MinFeeUSDCents:            10,
					MaxFeeUSDCents:            100,
					DeciBps:                   200,
					DestGasOverhead:           40000,
					DestBytesOverhead:         200,
					AggregateRateLimitEnabled: false,
					IsEnabled:                 true,
				},
			},
			FeeTokenConfig: map[common.Address]evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampFeeTokenConfig{
				common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"): {
					NetworkFeeUSDCents:         10,
					GasMultiplierWeiPerEth:     1000000000000000000,
					PremiumMultiplierWeiPerEth: 1100000000000000000,
					Enabled:                    true,
				},
				common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"): {
					NetworkFeeUSDCents:         10,
					GasMultiplierWeiPerEth:     900000000000000000,
					PremiumMultiplierWeiPerEth: 1000000000000000000,
					Enabled:                    true,
				},
			},
		},
	},
	{
		Address:       "0x6060606060606060606060606060606060606060",
		ChainSelector: 5936861837188149645,
		Metadata: seq1_6.FeeQuoterImportConfigSequenceOutput{
			RemoteChainCfgs: map[uint64]seq1_6.FeeQuoterImportConfigSequenceOutputPerRemoteChain{
				5009297550715157269: {
					DestChainCfg: fee_quoter_v1_6_0.DestChainConfig{
						IsEnabled:                         true,
						MaxNumberOfTokensPerMsg:           5,
						MaxDataBytes:                      10000,
						MaxPerMsgGasLimit:                 5000000,
						DestGasOverhead:                   100000,
						DestGasPerPayloadByteBase:         16,
						DestGasPerPayloadByteHigh:         32,
						DestGasPerPayloadByteThreshold:    1000,
						DestDataAvailabilityOverheadGas:   50000,
						DestGasPerDataAvailabilityByte:    10,
						DestDataAvailabilityMultiplierBps: 1000,
						ChainFamilySelector:               [4]byte(utils.GetSelectorHex(5009297550715157269)),
						EnforceOutOfOrder:                 false,
						DefaultTokenFeeUSDCents:           10,
						DefaultTokenDestGasOverhead:       50000,
						DefaultTxGasLimit:                 200000,
						GasMultiplierWeiPerEth:            0,
						GasPriceStalenessThreshold:        0,
						NetworkFeeUSDCents:                10,
					},
					TokenTransferFeeCfgs: map[common.Address]fee_quoter_v1_6_0.TokenTransferFeeConfig{},
				},
				4949039107694359620: {
					DestChainCfg: fee_quoter_v1_6_0.DestChainConfig{
						IsEnabled:                         true,
						MaxNumberOfTokensPerMsg:           4,
						MaxDataBytes:                      9000,
						MaxPerMsgGasLimit:                 4500000,
						DestGasOverhead:                   90000,
						DestGasPerPayloadByteBase:         15,
						DestGasPerPayloadByteHigh:         30,
						DestGasPerPayloadByteThreshold:    900,
						DestDataAvailabilityOverheadGas:   45000,
						DestGasPerDataAvailabilityByte:    9,
						DestDataAvailabilityMultiplierBps: 950,
						ChainFamilySelector:               [4]byte(utils.GetSelectorHex(4949039107694359620)),
						EnforceOutOfOrder:                 false,
						DefaultTokenFeeUSDCents:           9,
						DefaultTokenDestGasOverhead:       45000,
						DefaultTxGasLimit:                 190000,
						GasMultiplierWeiPerEth:            0,
						GasPriceStalenessThreshold:        0,
						NetworkFeeUSDCents:                10,
					},
					TokenTransferFeeCfgs: map[common.Address]fee_quoter_v1_6_0.TokenTransferFeeConfig{},
				},
			},
			StaticCfg: fee_quoter_v1_6_0.StaticConfig{
				MaxFeeJuelsPerMsg:            big.NewInt(1000000000000000000),
				LinkToken:                    common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
				TokenPriceStalenessThreshold: 3600,
			},
			PriceUpdaters: []common.Address{
				common.HexToAddress("0x4444444444444444444444444444444444444444"),
				common.HexToAddress("0x5555555555555555555555555555555555555555"),
			},
		},
	},
	{
		Address:       "0x7070707070707070707070707070707070707070",
		ChainSelector: 5936861837188149645,
		Metadata: seq1_5.OnRampImportConfigSequenceOutput{
			RemoteChainSelector: 15971525489660198786,
			StaticConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampStaticConfig{
				LinkToken:          common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
				ChainSelector:      5936861837188149645,
				DestChainSelector:  15971525489660198786,
				DefaultTxGasLimit:  180000,
				MaxNopFeesJuels:    big.NewInt(900000000000000000),
				PrevOnRamp:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
				RmnProxy:           common.HexToAddress("0x7777777777777777777777777777777777777777"),
				TokenAdminRegistry: common.HexToAddress("0x8888888888888888888888888888888888888888"),
			},
			DynamicConfig: evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampDynamicConfig{
				Router:                            common.HexToAddress("0x9999999999999999999999999999999999999999"),
				MaxNumberOfTokensPerMsg:           3,
				DestGasOverhead:                   80000,
				DestGasPerPayloadByte:             14,
				DestDataAvailabilityOverheadGas:   40000,
				DestGasPerDataAvailabilityByte:    8,
				DestDataAvailabilityMultiplierBps: 900,
				PriceRegistry:                     common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
				MaxDataBytes:                      8000,
				MaxPerMsgGasLimit:                 4000000,
				DefaultTokenFeeUSDCents:           0,
				DefaultTokenDestGasOverhead:       0,
				EnforceOutOfOrder:                 false,
			},
			TokenTransferFeeConfig: map[common.Address]evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampTokenTransferFeeConfig{},
			FeeTokenConfig: map[common.Address]evm_2_evm_onramp_v1_5_0.EVM2EVMOnRampFeeTokenConfig{
				common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"): {
					NetworkFeeUSDCents:         10,
					GasMultiplierWeiPerEth:     950000000000000000,
					PremiumMultiplierWeiPerEth: 1050000000000000000,
					Enabled:                    true,
				},
			},
		},
	},
}

// validMaxFeeJuelsPerMsgFromMetadata returns the set of valid MaxFeeJuelsPerMsg for a chain
// by collecting values from contract metadata. The sequence merges v1.6 (FeeQuoter) and v1.5 (OnRamp)
// outputs; when the v1.6 path finds a FeeQuoter ref it uses FeeQuoter's StaticCfg.MaxFeeJuelsPerMsg,
// otherwise it may use only v1.5's ConstructorArgs which take MaxFeeJuelsPerMsg from OnRamp's
// MaxNopFeesJuels. So for chains that have both FeeQuoter and OnRamp metadata, either value is valid
// depending on ref version matching the sequence's dependency. Keys are big.Int.String() for set lookup.
func validMaxFeeJuelsPerMsgFromMetadata(chainSelector uint64, contractMetadata []datastore.ContractMetadata) map[string]bool {
	valid := make(map[string]bool)
	for _, meta := range contractMetadata {
		if meta.ChainSelector != chainSelector {
			continue
		}
		if fq, ok := meta.Metadata.(seq1_6.FeeQuoterImportConfigSequenceOutput); ok {
			if fq.StaticCfg.MaxFeeJuelsPerMsg != nil {
				valid[fq.StaticCfg.MaxFeeJuelsPerMsg.String()] = true
			}
		}
		if onr, ok := meta.Metadata.(seq1_5.OnRampImportConfigSequenceOutput); ok {
			if onr.StaticConfig.MaxNopFeesJuels != nil {
				valid[onr.StaticConfig.MaxNopFeesJuels.String()] = true
			}
		}
	}
	return valid
}

// getExpectedOutput returns hardcoded expected FeeQuoterUpdate values based on contract_metadata.json
func getExpectedOutput() map[uint64]sequence1_7.FeeQuoterUpdate {
	linkToken := common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA")
	maxFeeJuels, _ := new(big.Int).SetString("1000000000000000000", 10)

	expected := make(map[uint64]sequence1_7.FeeQuoterUpdate)

	// Chain 5009297550715157269: Has FeeQuoter v1.6.3 + OnRamp v1.5.0
	// Since no FeeQuoter v1.7.0 exists, it's a new deployment (ConstructorArgs populated)
	expected[5009297550715157269] = sequence1_7.FeeQuoterUpdate{
		ChainSelector: 5009297550715157269,
		ConstructorArgs: fqops.ConstructorArgs{
			StaticConfig: fqops.StaticConfig{
				LinkToken:         linkToken,
				MaxFeeJuelsPerMsg: maxFeeJuels,
			},
			PriceUpdaters: []common.Address{
				common.HexToAddress("0x4444444444444444444444444444444444444444"),
				common.HexToAddress("0x5555555555555555555555555555555555555555"),
				common.HexToAddress("0x2222222222222222222222222222222222222221"),
			},
			DestChainConfigArgs: []fqops.DestChainConfigArgs{
				{
					DestChainSelector: 15971525489660198786,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true,
						MaxDataBytes:                8000,
						MaxPerMsgGasLimit:           4000000,
						DestGasOverhead:             80000,
						DestGasPerPayloadByteBase:   14,
						ChainFamilySelector:         [4]byte(utils.GetSelectorHex(15971525489660198786)),
						DefaultTokenFeeUSDCents:     8,
						DefaultTokenDestGasOverhead: 40000,
						DefaultTxGasLimit:           180000,
						NetworkFeeUSDCents:          10,
						LinkFeeMultiplierPercent:    90,
					},
				},
				{
					DestChainSelector: 4949039107694359620,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true,
						MaxDataBytes:                10000,
						MaxPerMsgGasLimit:           5000000,
						DestGasOverhead:             100000,
						DestGasPerPayloadByteBase:   16,
						ChainFamilySelector:         [4]byte(utils.GetSelectorHex(4949039107694359620)),
						DefaultTokenFeeUSDCents:     0, // Not in OnRamp DynamicConfig in test data
						DefaultTokenDestGasOverhead: 0, // Not in OnRamp DynamicConfig in test data
						DefaultTxGasLimit:           200000,
						NetworkFeeUSDCents:          10,
						LinkFeeMultiplierPercent:    90,
					},
				},
			},
			TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
				{
					DestChainSelector: 15971525489660198786,
					TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
						{
							Token: common.HexToAddress("0x2222222222222222222222222222222222222222"),
							TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
								FeeUSDCents:       4,
								DestGasOverhead:   25000,
								DestBytesOverhead: 80,
								IsEnabled:         true,
							},
						},
					},
				},
				{
					DestChainSelector: 4949039107694359620,
					TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
						{
							Token: common.HexToAddress("0x2222222222222222222222222222222222222222"),
							TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
								FeeUSDCents:       5,
								DestGasOverhead:   30000,
								DestBytesOverhead: 100,
								IsEnabled:         true,
							},
						},
						{
							Token: common.HexToAddress("0x3333333333333333333333333333333333333333"),
							TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
								FeeUSDCents:       10,
								DestGasOverhead:   40000,
								DestBytesOverhead: 200,
								IsEnabled:         true,
							},
						},
					},
				},
			},
		},
	}

	// Chain 4949039107694359620: Has FeeQuoter v1.6.3 + OnRamp v1.5.0
	expected[4949039107694359620] = sequence1_7.FeeQuoterUpdate{
		ChainSelector: 4949039107694359620,
		ConstructorArgs: fqops.ConstructorArgs{
			StaticConfig: fqops.StaticConfig{
				LinkToken:         linkToken,
				MaxFeeJuelsPerMsg: maxFeeJuels,
			},
			PriceUpdaters: []common.Address{
				common.HexToAddress("0x4444444444444444444444444444444444444444"),
				common.HexToAddress("0x5555555555555555555555555555555555555555"),
				common.HexToAddress("0x3333333333333333333333333333333333333333"),
				common.HexToAddress("0x9999999999999999999999999999999999999999"),
			},
			DestChainConfigArgs: []fqops.DestChainConfigArgs{
				{
					DestChainSelector: 15971525489660198786,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true,
						MaxDataBytes:                8000,
						MaxPerMsgGasLimit:           4000000,
						DestGasOverhead:             80000,
						DestGasPerPayloadByteBase:   14,
						ChainFamilySelector:         [4]byte(utils.GetSelectorHex(15971525489660198786)),
						DefaultTokenFeeUSDCents:     8,
						DefaultTokenDestGasOverhead: 40000,
						DefaultTxGasLimit:           180000,
						NetworkFeeUSDCents:          10,
						LinkFeeMultiplierPercent:    90,
					},
				},
				{
					DestChainSelector: 5009297550715157269,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true,
						MaxDataBytes:                10000,
						MaxPerMsgGasLimit:           5000000,
						DestGasOverhead:             100000,
						DestGasPerPayloadByteBase:   16,
						ChainFamilySelector:         [4]byte(utils.GetSelectorHex(5009297550715157269)),
						DefaultTokenFeeUSDCents:     0, // Not in OnRamp DynamicConfig in test data
						DefaultTokenDestGasOverhead: 0, // Not in OnRamp DynamicConfig in test data
						DefaultTxGasLimit:           200000,
						NetworkFeeUSDCents:          10,
						LinkFeeMultiplierPercent:    90,
					},
				},
			},
			TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
				{
					DestChainSelector: 15971525489660198786,
					TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
						{
							Token: common.HexToAddress("0x2222222222222222222222222222222222222222"),
							TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
								FeeUSDCents:       4,
								DestGasOverhead:   25000,
								DestBytesOverhead: 80,
								IsEnabled:         true,
							},
						},
					},
				},
				{
					DestChainSelector: 5009297550715157269,
					TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
						{
							Token: common.HexToAddress("0x2222222222222222222222222222222222222222"),
							TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
								FeeUSDCents:       5,
								DestGasOverhead:   30000,
								DestBytesOverhead: 100,
								IsEnabled:         true,
							},
						},
					},
				},
			},
		},
	}

	// Chain 15971525489660198786: Only has OnRamp v1.5.0
	maxFeeJuels159, _ := new(big.Int).SetString("1000000000000000000", 10)
	expected[15971525489660198786] = sequence1_7.FeeQuoterUpdate{
		ChainSelector: 15971525489660198786,
		ConstructorArgs: fqops.ConstructorArgs{
			StaticConfig: fqops.StaticConfig{
				LinkToken:         linkToken,
				MaxFeeJuelsPerMsg: maxFeeJuels159,
			},
			DestChainConfigArgs: []fqops.DestChainConfigArgs{
				{
					DestChainSelector: 5009297550715157269,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true,
						MaxDataBytes:                10000,
						MaxPerMsgGasLimit:           5000000,
						DestGasOverhead:             100000,
						DestGasPerPayloadByteBase:   16,
						ChainFamilySelector:         [4]byte(utils.GetSelectorHex(5009297550715157269)),
						DefaultTokenFeeUSDCents:     0, // Not in OnRamp DynamicConfig in test data
						DefaultTokenDestGasOverhead: 0, // Not in OnRamp DynamicConfig in test data
						DefaultTxGasLimit:           200000,
						NetworkFeeUSDCents:          10,
						LinkFeeMultiplierPercent:    90,
					},
				},
			},
			TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
				{
					DestChainSelector: 5009297550715157269,
					TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{
						{
							Token: common.HexToAddress("0x2222222222222222222222222222222222222222"),
							TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
								FeeUSDCents:       5,
								DestGasOverhead:   30000,
								DestBytesOverhead: 100,
								IsEnabled:         true,
							},
						},
						{
							Token: common.HexToAddress("0x3333333333333333333333333333333333333333"),
							TokenTransferFeeConfig: fee_quoter.FeeQuoterTokenTransferFeeConfig{
								FeeUSDCents:       10,
								DestGasOverhead:   40000,
								DestBytesOverhead: 200,
								IsEnabled:         true,
							},
						},
					},
				},
			},
			PriceUpdaters: []common.Address{
				common.HexToAddress("0x4444444444444444444444444444444444444444"),
			},
		},
	}

	// Chain 5936861837188149645: Has FeeQuoter v1.6.3 + OnRamp v1.5.0
	expected[5936861837188149645] = sequence1_7.FeeQuoterUpdate{
		ChainSelector: 5936861837188149645,
		ConstructorArgs: fqops.ConstructorArgs{
			StaticConfig: fqops.StaticConfig{
				LinkToken:         linkToken,
				MaxFeeJuelsPerMsg: maxFeeJuels,
			},
			PriceUpdaters: []common.Address{
				common.HexToAddress("0x4444444444444444444444444444444444444444"),
				common.HexToAddress("0x5555555555555555555555555555555555555555"),
				common.HexToAddress("0x5555555555555555555555555555555555555551"),
			},
			DestChainConfigArgs: []fqops.DestChainConfigArgs{
				{
					DestChainSelector: 5009297550715157269,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true,
						MaxDataBytes:                10000,
						MaxPerMsgGasLimit:           5000000,
						DestGasOverhead:             100000,
						DestGasPerPayloadByteBase:   16,
						ChainFamilySelector:         [4]byte(utils.GetSelectorHex(5009297550715157269)),
						DefaultTokenFeeUSDCents:     10,
						DefaultTokenDestGasOverhead: 50000,
						DefaultTxGasLimit:           200000,
						NetworkFeeUSDCents:          10,
						LinkFeeMultiplierPercent:    90,
					},
				},
				{
					DestChainSelector: 4949039107694359620,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true,
						MaxDataBytes:                9000,
						MaxPerMsgGasLimit:           4500000,
						DestGasOverhead:             90000,
						DestGasPerPayloadByteBase:   15,
						ChainFamilySelector:         [4]byte(utils.GetSelectorHex(4949039107694359620)),
						DefaultTokenFeeUSDCents:     9,
						DefaultTokenDestGasOverhead: 45000,
						DefaultTxGasLimit:           190000,
						NetworkFeeUSDCents:          10,
						LinkFeeMultiplierPercent:    90,
					},
				},
				{
					DestChainSelector: 15971525489660198786,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   true,
						MaxDataBytes:                8000,
						MaxPerMsgGasLimit:           4000000,
						DestGasOverhead:             80000,
						DestGasPerPayloadByteBase:   14,
						ChainFamilySelector:         [4]byte(utils.GetSelectorHex(15971525489660198786)),
						DefaultTokenFeeUSDCents:     0, // Not in OnRamp DynamicConfig in test data
						DefaultTokenDestGasOverhead: 0, // Not in OnRamp DynamicConfig in test data
						DefaultTxGasLimit:           180000,
						NetworkFeeUSDCents:          10,
						LinkFeeMultiplierPercent:    90,
					},
				},
			},
			TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{
				{
					DestChainSelector:       5009297550715157269,
					TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{},
				},
				{
					DestChainSelector:       4949039107694359620,
					TokenTransferFeeConfigs: []fee_quoter.FeeQuoterTokenTransferFeeConfigSingleTokenArgs{},
				},
			},
		},
	}

	return expected
}

func TestSequenceFeeQuoterInputCreation(t *testing.T) {
	contractMetadata := dummyContractMetadata
	addressRefs := dummyAddressRefs

	// Collect unique chain selectors from address refs
	chainSelectors := make(map[uint64]bool)
	for _, ref := range addressRefs {
		chainSelectors[ref.ChainSelector] = true
	}

	// Convert map keys to slice and sort for deterministic test order (avoids flakiness from map iteration)
	chainSelectorList := make([]uint64, 0, len(chainSelectors))
	for selector := range chainSelectors {
		chainSelectorList = append(chainSelectorList, selector)
	}
	sort.Slice(chainSelectorList, func(i, j int) bool { return chainSelectorList[i] < chainSelectorList[j] })

	// Create environment with simulated EVM chains
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chainSelectorList),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	// Load address refs into a new datastore
	// Note: The environment's datastore is sealed, so we'll use our own datastore
	// and pass the data directly in the input to the sequence
	ds := datastore.NewMemoryDataStore()
	for _, ref := range addressRefs {
		err := ds.Addresses().Add(ref)
		require.NoError(t, err, "Failed to add address ref %+v to datastore", ref)
	}

	// Load contract metadata into the datastore
	err = sequences.WriteMetadataToDatastore(ds, sequences.Metadata{
		Contracts: contractMetadata,
	})
	require.NoError(t, err, "Failed to write contract metadata to datastore")

	// Seal the datastore for use in the test
	e.DataStore = ds.Seal()

	// Get the FeeQuoterUpdater adapter (use concrete type so report.Output is sequence1_7.FeeQuoterUpdate)
	fquUpdater := evmadapter.FeeQuoterUpdater[sequence1_7.FeeQuoterUpdate]{}

	// Test the sequence for each chain selector that has a FeeQuoter
	for _, chainSelector := range chainSelectorList {
		_, ok := e.BlockChains.EVMChains()[chainSelector]
		require.True(t, ok, "Chain with selector %d should exist", chainSelector)

		// Build input from original slices so Version/Type match exactly (sealed datastore
		// can alter refs and break GetAddressRef lookup for FeeQuoter 1.6.0).
		existingAddresses := make([]datastore.AddressRef, 0)
		for _, ref := range addressRefs {
			if ref.ChainSelector == chainSelector {
				existingAddresses = append(existingAddresses, ref)
			}
		}
		contractMeta := make([]datastore.ContractMetadata, 0)
		for _, meta := range contractMetadata {
			if meta.ChainSelector == chainSelector {
				contractMeta = append(contractMeta, meta)
			}
		}
		sort.Slice(existingAddresses, func(i, j int) bool {
			if existingAddresses[i].Type != existingAddresses[j].Type {
				return string(existingAddresses[i].Type) < string(existingAddresses[j].Type)
			}
			return existingAddresses[i].Address < existingAddresses[j].Address
		})
		sort.Slice(contractMeta, func(i, j int) bool {
			return contractMeta[i].Address < contractMeta[j].Address
		})

		// Create input for SequenceFeeQuoterInputCreation
		input := deploy.FeeQuoterUpdateInput{
			ChainSelector:     chainSelector,
			ExistingAddresses: existingAddresses,
			ContractMeta:      contractMeta,
		}

		// Execute the sequence
		report, err := cldf_ops.ExecuteSequence(
			e.OperationsBundle,
			fquUpdater.SequenceFeeQuoterInputCreation(),
			e.BlockChains,
			input,
		)

		// Verify the sequence executed successfully
		require.NoError(t, err, "SequenceFeeQuoterInputCreation should not error for chain %d", chainSelector)
		require.NotNil(t, report, "Report should not be nil for chain %d", chainSelector)

		// Verify the output is not empty
		output := report.Output
		isEmpty, err := output.IsEmpty()
		require.NoError(t, err, "IsEmpty check should not error")
		require.False(t, isEmpty, "Output should not be empty for chain %d", chainSelector)

		// Verify basic output structure
		require.Equal(t, chainSelector, output.ChainSelector, "Chain selector should match input")
		require.Equal(t, existingAddresses, output.ExistingAddresses, "Existing addresses should match input")

		// Get expected output (hardcoded based on contract_metadata.json)
		expectedMap := getExpectedOutput()
		expected, hasExpected := expectedMap[chainSelector]
		require.True(t, hasExpected, "Expected output should exist for chain %d", chainSelector)

		// Verify that the output has meaningful data
		// At least one of these should be populated:
		// - ConstructorArgs
		// - DestChainConfigs
		// - TokenTransferFeeConfigUpdates
		// - AuthorizedCallerUpdates
		hasData := !output.ConstructorArgs.IsEmpty() ||
			len(output.DestChainConfigs) > 0 ||
			len(output.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs) > 0 ||
			len(output.AuthorizedCallerUpdates.AddedCallers) > 0 ||
			len(output.AuthorizedCallerUpdates.RemovedCallers) > 0

		require.True(t, hasData, "Output should have at least some configuration data for chain %d", chainSelector)
		// Assert against expected values
		if !expected.ConstructorArgs.IsEmpty() {
			require.False(t, output.ConstructorArgs.IsEmpty(), "ConstructorArgs should be present for new deployment on chain %d", chainSelector)
			require.Equal(t, expected.ConstructorArgs.StaticConfig.LinkToken, output.ConstructorArgs.StaticConfig.LinkToken,
				"LinkToken should match expected value on chain %d", chainSelector)
			// MaxFeeJuelsPerMsg must be one of the values present in contract metadata for this chain
			// (FeeQuoter StaticCfg or OnRamp MaxNopFeesJuels); the sequence uses one or the other depending
			// on whether the v1.6 path finds a FeeQuoter ref.
			validMaxFee := validMaxFeeJuelsPerMsgFromMetadata(chainSelector, contractMetadata)
			require.NotEmpty(t, validMaxFee, "contract metadata for chain %d should define at least one MaxFeeJuelsPerMsg source", chainSelector)
			require.True(t, validMaxFee[output.ConstructorArgs.StaticConfig.MaxFeeJuelsPerMsg.String()],
				"MaxFeeJuelsPerMsg should be one of the values from contract metadata (FeeQuoter or OnRamp) on chain %d", chainSelector)
			require.ElementsMatch(t, expected.ConstructorArgs.PriceUpdaters, output.ConstructorArgs.PriceUpdaters,
				"PriceUpdaters should match expected value on chain %d", chainSelector)
		} else {
			// For existing deployments, ConstructorArgs should be empty
			require.True(t, output.ConstructorArgs.IsEmpty(), "ConstructorArgs should be empty for existing deployment on chain %d", chainSelector)
		}

		// Assert specific values based on the sequence logic in feequoterupdater.go
		// The sequence merges outputs from CreateFeeQuoterUpdateInputFromV163 and CreateFeeQuoterUpdateInputFromV150

		// Verify DestChainConfigs against expected values
		// Build a map of expected dest chain configs for easier lookup
		expectedDestChainConfigsMap := make(map[uint64]fqops.DestChainConfigArgs)
		for _, cfg := range expected.DestChainConfigs {
			expectedDestChainConfigsMap[cfg.DestChainSelector] = cfg
		}
		require.Len(t, output.DestChainConfigs, len(expectedDestChainConfigsMap),
			"Number of DestChainConfigs should match expected value on chain %d", chainSelector)

		for _, destChainCfg := range output.DestChainConfigs {
			if expectedCfg, exists := expectedDestChainConfigsMap[destChainCfg.DestChainSelector]; exists {
				require.Equal(t, expectedCfg, destChainCfg, "DestChainConfig should match expected value for "+
					"DestChainSelector %d on chain %d", destChainCfg.DestChainSelector, chainSelector)
			}
		}
		for _, cfg := range expected.ConstructorArgs.DestChainConfigArgs {
			expectedDestChainConfigsMap[cfg.DestChainSelector] = cfg
		}
		require.Len(t, output.ConstructorArgs.DestChainConfigArgs, len(expectedDestChainConfigsMap),
			"Number of Constructor DestChainConfigArgs should match expected value for chain %d", chainSelector)

		for _, destChainCfg := range output.ConstructorArgs.DestChainConfigArgs {
			if expectedCfg, exists := expectedDestChainConfigsMap[destChainCfg.DestChainSelector]; exists {
				require.Equal(t, expectedCfg, destChainCfg, "Constructor DestChainConfig should match expected value for "+
					"DestChainSelector %d on chain %d", destChainCfg.DestChainSelector, chainSelector)
			}
		}

		require.Len(t, output.ConstructorArgs.TokenTransferFeeConfigArgs, len(expected.ConstructorArgs.TokenTransferFeeConfigArgs),
			"Number of TokenTransferFeeConfigArgs should match expected value for chain %d", chainSelector)
		require.Len(t, output.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs, len(expected.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs),
			"Number of TokenTransferFeeConfigUpdates should match expected value for chain %d", chainSelector)
		for _, tokenTransferFeeConfig := range output.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs {
			found := false
			for _, expectedCfg := range expected.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs {
				if tokenTransferFeeConfig.DestChainSelector == expectedCfg.DestChainSelector {
					require.ElementsMatch(t, expectedCfg.TokenTransferFeeConfigs, tokenTransferFeeConfig.TokenTransferFeeConfigs,
						"TokenTransferFeeConfigs should match expected value for DestChainSelector %d on chain %d",
						tokenTransferFeeConfig.DestChainSelector, chainSelector)
					found = true
					break
				}
			}
			require.True(t, found, "Unexpected TokenTransferFeeConfig for DestChainSelector %d on chain %d",
				tokenTransferFeeConfig.DestChainSelector, chainSelector)
		}

		for _, tokenTransferFeeConfig := range output.ConstructorArgs.TokenTransferFeeConfigArgs {
			found := false
			for _, expectedCfg := range expected.ConstructorArgs.TokenTransferFeeConfigArgs {
				if tokenTransferFeeConfig.DestChainSelector == expectedCfg.DestChainSelector {
					require.ElementsMatch(t, expectedCfg.TokenTransferFeeConfigs, tokenTransferFeeConfig.TokenTransferFeeConfigs,
						"Constructor TokenTransferFeeConfigs should match expected value for DestChainSelector %d on chain %d",
						tokenTransferFeeConfig.DestChainSelector, chainSelector)
					found = true
					break
				}
			}
			require.True(t, found, "Unexpected Constructor TokenTransferFeeConfig for DestChainSelector %d on chain %d",
				tokenTransferFeeConfig.DestChainSelector, chainSelector)
		}

		// Verify AuthorizedCallerUpdates if present (for existing deployments)
		require.ElementsMatch(t, expected.AuthorizedCallerUpdates.AddedCallers, output.AuthorizedCallerUpdates.AddedCallers,
			"AuthorizedCallerUpdates.AddedCallers should match expected value on chain %d", chainSelector)
		require.ElementsMatch(t, expected.AuthorizedCallerUpdates.RemovedCallers, output.AuthorizedCallerUpdates.RemovedCallers,
			"AuthorizedCallerUpdates.RemovedCallers should match expected value on chain %d", chainSelector)

		t.Logf("Successfully executed SequenceFeeQuoterInputCreation for chain %d", chainSelector)
	}
}
