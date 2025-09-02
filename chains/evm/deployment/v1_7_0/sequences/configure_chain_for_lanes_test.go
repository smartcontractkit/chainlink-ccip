package sequences_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/call"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_aggregator"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestConfigureChainForLanes(t *testing.T) {
	tests := []struct {
		desc string
	}{
		{
			desc: "valid input",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainSelector := uint64(5009297550715157269)

			lggr, err := logger.New()
			require.NoError(t, err, "Failed to create logger")

			bundle := operations.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				operations.NewMemoryReporter(),
			)

			chain, err := cldf_evm_provider.NewSimChainProvider(t, chainSelector,
				cldf_evm_provider.SimChainProviderConfig{
					NumAdditionalAccounts: 1,
				},
			).Initialize(t.Context())
			require.NoError(t, err, "Failed to create SimChainProvider")

			chains := cldf_chain.NewBlockChainsFromSlice(
				[]cldf_chain.BlockChain{chain},
			)
			evmChain := chains.EVMChains()[chainSelector]

			deploymentReport, err := operations.ExecuteSequence(
				bundle,
				sequences.DeployChain,
				evmChain,
				sequences.DeployChainInput{
					ChainSelector: chainSelector,
					ContractParams: sequences.ContractParams{
						RMNRemote:     sequences.RMNRemoteParams{},
						CCVAggregator: sequences.CCVAggregatorParams{},
						CommitOnRamp: sequences.CommitOnRampParams{
							FeeAggregator: common.HexToAddress("0x01"),
						},
						CCVProxy: sequences.CCVProxyParams{
							FeeAggregator: common.HexToAddress("0x01"),
						},
						FeeQuoter: sequences.FeeQuoterParams{
							MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
							TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
							LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
							WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
						},
						CommitOffRamp: sequences.CommitOffRampParams{
							SignatureConfigArgs: commit_offramp.SignatureConfigArgs{{
								ConfigDigest: [32]byte{0x01},
								F:            1,
								Signers: []common.Address{
									common.HexToAddress("0x02"),
									common.HexToAddress("0x03"),
									common.HexToAddress("0x04"),
									common.HexToAddress("0x05"),
								},
							}},
						},
					},
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			var r common.Address
			var ccvProxy common.Address
			var commitOnRamp common.Address
			var feeQuoter common.Address
			var ccvAggregator common.Address
			var commitOffRamp common.Address
			for _, addr := range deploymentReport.Output.Addresses {
				switch addr.Type {
				case datastore.ContractType(router.ContractType):
					r = common.HexToAddress(addr.Address)
				case datastore.ContractType(ccv_proxy.ContractType):
					ccvProxy = common.HexToAddress(addr.Address)
				case datastore.ContractType(commit_onramp.ContractType):
					commitOnRamp = common.HexToAddress(addr.Address)
				case datastore.ContractType(fee_quoter.ContractType):
					feeQuoter = common.HexToAddress(addr.Address)
				case datastore.ContractType(ccv_aggregator.ContractType):
					ccvAggregator = common.HexToAddress(addr.Address)
				case datastore.ContractType(commit_offramp.ContractType):
					commitOffRamp = common.HexToAddress(addr.Address)
				}
			}
			ccipMessageSource := common.HexToAddress("0x10").Bytes()
			defaultExecutor := common.HexToAddress("0x11")
			fqDestChainConfig := fee_quoter.DestChainConfig{
				IsEnabled:                         true,
				MaxNumberOfTokensPerMsg:           10,
				MaxDataBytes:                      30_000,
				MaxPerMsgGasLimit:                 3_000_000,
				DestGasOverhead:                   300_000,
				DefaultTokenFeeUSDCents:           25,
				DestGasPerPayloadByteBase:         16,
				DestGasPerPayloadByteHigh:         40,
				DestGasPerPayloadByteThreshold:    3000,
				DestDataAvailabilityOverheadGas:   100,
				DestGasPerDataAvailabilityByte:    16,
				DestDataAvailabilityMultiplierBps: 1,
				DefaultTokenDestGasOverhead:       90_000,
				DefaultTxGasLimit:                 200_000,
				GasMultiplierWeiPerEth:            11e17, // Gas multiplier in wei per eth is scaled by 1e18, so 11e17 is 1.1 = 110%
				NetworkFeeUSDCents:                10,
				ChainFamilySelector:               [4]byte{0x28, 0x12, 0xd5, 0x2c}, // EVM
			}
			remoteChainSelector := uint64(4356164186791070119)

			_, err = operations.ExecuteSequence(
				bundle,
				sequences.ConfigureChainForLanes,
				evmChain,
				sequences.ConfigureChainForLanesInput{
					ChainSelector: chainSelector,
					Router:        r,
					CCVProxy:      ccvProxy,
					CommitOnRamp:  commitOnRamp,
					FeeQuoter:     feeQuoter,
					CCVAggregator: ccvAggregator,
					RemoteChains: map[uint64]sequences.RemoteChainConfig{
						remoteChainSelector: {
							AllowTrafficFrom:            true,
							CCIPMessageSource:           ccipMessageSource,
							DefaultCCVOffRamps:          []common.Address{commitOffRamp},
							LaneMandatedCCVOffRamps:     []common.Address{commitOffRamp},
							DefaultCCVOnRamp:            commitOnRamp,
							RequiredCCVOnRamp:           commitOnRamp,
							DefaultExecutor:             defaultExecutor,
							CommitOnRampDestChainConfig: sequences.CommitOnRampDestChainConfig{},
							// FeeQuoterDestChainConfig configures the FeeQuoter for this remote chain
							FeeQuoterDestChainConfig: fqDestChainConfig,
						},
					},
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Check onRamps on router
			onRampOnRouter, err := operations.ExecuteOperation(bundle, router.GetOnRamp, evmChain, call.Input[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       r,
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, ccvProxy.Hex(), onRampOnRouter.Output.Hex(), "OnRamp address on router should match CCVProxy address")

			// Check offRamps on router
			offRampsOnRouter, err := operations.ExecuteOperation(bundle, router.GetOffRamps, evmChain, call.Input[any]{
				ChainSelector: evmChain.Selector,
				Address:       r,
				Args:          nil,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Len(t, offRampsOnRouter.Output, 1, "There should be one OffRamp on the router for the remote chain")
			require.Equal(t, ccvAggregator.Hex(), offRampsOnRouter.Output[0].OffRamp.Hex(), "OffRamp address on router should match CCVAggregator address")

			// Check sourceChainConfig on CCVAggregator
			sourceChainConfig, err := operations.ExecuteOperation(bundle, ccv_aggregator.GetSourceChainConfig, evmChain, call.Input[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       ccvAggregator,
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, ccipMessageSource, sourceChainConfig.Output.OnRamp, "OnRamp in source chain config should match CCVProxy address")
			require.Len(t, sourceChainConfig.Output.DefaultCCVs, 1, "There should be one DefaultCCV in source chain config")
			require.Equal(t, commitOffRamp.Hex(), sourceChainConfig.Output.DefaultCCVs[0].Hex(), "DefaultCCV in source chain config should match CommitOffRamp address")
			require.Len(t, sourceChainConfig.Output.LaneMandatedCCVs, 1, "There should be one LaneMandatedCCV in source chain config")
			require.Equal(t, commitOffRamp.Hex(), sourceChainConfig.Output.LaneMandatedCCVs[0].Hex(), "LaneMandatedCCV in source chain config should match CommitOffRamp address")
			require.True(t, sourceChainConfig.Output.IsEnabled, "IsEnabled in source chain config should be true")
			require.Equal(t, r.Hex(), sourceChainConfig.Output.Router.Hex(), "Router in source chain config should match Router address")

			// Check destChainConfig on CCVProxy
			destChainConfig, err := operations.ExecuteOperation(bundle, ccv_proxy.GetDestChainConfig, evmChain, call.Input[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       ccvProxy,
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, r.Hex(), destChainConfig.Output.Router.Hex(), "Router in dest chain config should match Router address")

			// Check destChainConfig on CommitOnRamp
			commitOnRampDestChainConfig, err := operations.ExecuteOperation(bundle, commit_onramp.GetDestChainConfig, evmChain, call.Input[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       commitOnRamp,
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, ccvProxy.Hex(), commitOnRampDestChainConfig.Output.CcvProxy.Hex(), "CcvProxy in CommitOnRamp dest chain config should match CommitOnRamp address")
			require.False(t, commitOnRampDestChainConfig.Output.AllowlistEnabled, "AllowlistEnabled in CommitOnRamp dest chain config should be false")
		})
	}
}
