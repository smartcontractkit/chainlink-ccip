package sequences_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_aggregator"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/message_hasher"
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

			usdPerLink, ok := new(big.Int).SetString("15000000000000000000", 10) // $15
			require.True(t, ok, "Failed to parse USDPerLINK")
			usdPerWeth, ok := new(big.Int).SetString("2000000000000000000000", 10) // $2000
			require.True(t, ok, "Failed to parse USDPerWETH")

			deploymentReport, err := operations.ExecuteSequence(
				bundle,
				sequences.DeployChainContracts,
				evmChain,
				sequences.DeployChainContractsInput{
					ChainSelector: chainSelector,
					ContractParams: sequences.ContractParams{
						RMNRemote:     sequences.RMNRemoteParams{},
						CCVAggregator: sequences.CCVAggregatorParams{},
						CommitteeVerifier: sequences.CommitteeVerifierParams{
							FeeAggregator: common.HexToAddress("0x01"),
							SignatureConfigArgs: committee_verifier.SetSignatureConfigArgs{
								Threshold: 1,
								Signers: []common.Address{
									common.HexToAddress("0x02"),
									common.HexToAddress("0x03"),
									common.HexToAddress("0x04"),
									common.HexToAddress("0x05"),
								},
							},
							StorageLocation: "https://test.chain.link.fake",
						},
						ExecutorOnRamp: sequences.ExecutorOnRampParams{
							MaxCCVsPerMsg: 10,
						},
						CCVProxy: sequences.CCVProxyParams{
							FeeAggregator: common.HexToAddress("0x01"),
						},
						FeeQuoter: sequences.FeeQuoterParams{
							MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
							TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
							LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
							WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
							USDPerLINK:                     usdPerLink,
							USDPerWETH:                     usdPerWeth,
						},
					},
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			var r common.Address
			var ccvProxy common.Address
			var feeQuoter common.Address
			var ccvAggregator common.Address
			var committeeVerifier common.Address
			var executorOnRamp common.Address
			for _, addr := range deploymentReport.Output.Addresses {
				switch addr.Type {
				case datastore.ContractType(router.ContractType):
					r = common.HexToAddress(addr.Address)
				case datastore.ContractType(ccv_proxy.ContractType):
					ccvProxy = common.HexToAddress(addr.Address)
				case datastore.ContractType(fee_quoter_v2.ContractType):
					feeQuoter = common.HexToAddress(addr.Address)
				case datastore.ContractType(ccv_aggregator.ContractType):
					ccvAggregator = common.HexToAddress(addr.Address)
				case datastore.ContractType(committee_verifier.ContractType):
					committeeVerifier = common.HexToAddress(addr.Address)
				case datastore.ContractType(executor_onramp.ContractType):
					executorOnRamp = common.HexToAddress(addr.Address)
				}
			}
			ccipMessageSource := common.HexToAddress("0x10").Bytes()
			ccipMessageDest := common.HexToAddress("0x11").Bytes()
			fqDestChainConfig := fee_quoter_v2.DestChainConfig{
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
					ChainSelector:     chainSelector,
					Router:            r,
					CCVProxy:          ccvProxy,
					CommitteeVerifier: committeeVerifier,
					FeeQuoter:         feeQuoter,
					CCVAggregator:     ccvAggregator,
					RemoteChains: map[uint64]sequences.RemoteChainConfig{
						remoteChainSelector: {
							AllowTrafficFrom:                 true,
							CCIPMessageSource:                ccipMessageSource,
							CCIPMessageDest:                  ccipMessageDest,
							DefaultCCVOffRamps:               []common.Address{committeeVerifier},
							DefaultCCVOnRamps:                []common.Address{committeeVerifier},
							DefaultExecutor:                  executorOnRamp,
							CommitteeVerifierDestChainConfig: sequences.CommitteeVerifierDestChainConfig{},
							// FeeQuoterDestChainConfig configures the FeeQuoter for this remote chain
							FeeQuoterDestChainConfig: fqDestChainConfig,
						},
					},
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Check onRamps on router
			onRampOnRouter, err := operations.ExecuteOperation(bundle, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       r,
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, ccvProxy.Hex(), onRampOnRouter.Output.Hex(), "OnRamp address on router should match CCVProxy address")

			// Check offRamps on router
			offRampsOnRouter, err := operations.ExecuteOperation(bundle, router.GetOffRamps, evmChain, contract.FunctionInput[any]{
				ChainSelector: evmChain.Selector,
				Address:       r,
				Args:          nil,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Len(t, offRampsOnRouter.Output, 1, "There should be one OffRamp on the router for the remote chain")
			require.Equal(t, ccvAggregator.Hex(), offRampsOnRouter.Output[0].OffRamp.Hex(), "OffRamp address on router should match CCVAggregator address")

			// Check sourceChainConfig on CCVAggregator
			sourceChainConfig, err := operations.ExecuteOperation(bundle, ccv_aggregator.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       ccvAggregator,
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, ccipMessageSource, sourceChainConfig.Output.OnRamp, "OnRamp in source chain config should match CCVProxy address")
			require.Len(t, sourceChainConfig.Output.DefaultCCVs, 1, "There should be one DefaultCCV in source chain config")
			require.Equal(t, committeeVerifier.Hex(), sourceChainConfig.Output.DefaultCCVs[0].Hex(), "DefaultCCV in source chain config should match CommitteeVerifier address")
			require.True(t, sourceChainConfig.Output.IsEnabled, "IsEnabled in source chain config should be true")
			require.Equal(t, r.Hex(), sourceChainConfig.Output.Router.Hex(), "Router in source chain config should match Router address")

			// Check destChainConfig on CCVProxy
			destChainConfig, err := operations.ExecuteOperation(bundle, ccv_proxy.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       ccvProxy,
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, r.Hex(), destChainConfig.Output.Router.Hex(), "Router in dest chain config should match Router address")
			require.Equal(t, ccipMessageDest, destChainConfig.Output.CcvAggregator, "CcvAggregator in dest chain config should match CCIPMessageDest")
			require.Equal(t, executorOnRamp.Hex(), destChainConfig.Output.DefaultExecutor.Hex(), "DefaultExecutor in dest chain config should match configured DefaultExecutor")
			require.Len(t, destChainConfig.Output.DefaultCCVs, 1, "There should be one DefaultCCV in dest chain config")
			require.Equal(t, committeeVerifier.Hex(), destChainConfig.Output.DefaultCCVs[0].Hex(), "DefaultCCV in dest chain config should match CommitteeVerifier address")

			// Check destChainConfig on CommitteeVerifier
			committeeVerifierDestChainConfig, err := operations.ExecuteOperation(bundle, committee_verifier.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       committeeVerifier,
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, r.Hex(), committeeVerifierDestChainConfig.Output.Router.Hex(), "Router in CommitteeVerifier dest chain config should match Router address")
			require.False(t, committeeVerifierDestChainConfig.Output.AllowlistEnabled, "AllowlistEnabled in CommitteeVerifier dest chain config should be false")

			// Check dest chains on ExecutorOnRamp
			executorOnRampDestChains, err := operations.ExecuteOperation(bundle, executor_onramp.GetDestChains, evmChain, contract.FunctionInput[any]{
				ChainSelector: evmChain.Selector,
				Address:       executorOnRamp,
				Args:          nil,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Len(t, executorOnRampDestChains.Output, 1, "There should be one dest chain on ExecutorOnRamp")
			require.Equal(t, remoteChainSelector, executorOnRampDestChains.Output[0], "Dest chain selector on ExecutorOnRamp should match remote chain selector")

			/////////////////////////////////////////
			// Try sending CCIP message /////////////
			/////////////////////////////////////////

			_, tx, msgHasher, err := message_hasher.DeployMessageHasher(evmChain.DeployerKey, evmChain.Client)
			require.NoError(t, err, "Failed to deploy MessageHasher")
			_, err = evmChain.Confirm(tx)
			require.NoError(t, err, "Failed to confirm MessageHasher deployment")

			extraArgs, err := msgHasher.EncodeGenericExtraArgsV3(
				&bind.CallOpts{Context: t.Context()},
				message_hasher.ClientEVMExtraArgsV3{
					RequiredCCV: []message_hasher.ClientCCV{
						{
							CcvAddress: committeeVerifier,
							Args:       []byte{},
						},
					},
					OptionalCCV:       []message_hasher.ClientCCV{},
					OptionalThreshold: 0,
					FinalityConfig:    0,
					Executor:          executorOnRamp,
					ExecutorArgs:      []byte{},
					TokenArgs:         []byte{},
				},
			)
			require.NoError(t, err, "EncodeGenericExtraArgsV3 should not error")

			ccipSendArgs := router.CCIPSendArgs{
				DestChainSelector: remoteChainSelector,
				EVM2AnyMessage: router.EVM2AnyMessage{
					Receiver:     common.LeftPadBytes(evmChain.DeployerKey.From.Bytes(), 32),
					Data:         []byte{},
					TokenAmounts: []router.EVMTokenAmount{},
					ExtraArgs:    extraArgs,
				},
			}

			fee, err := operations.ExecuteOperation(bundle, router.GetFee, evmChain, contract.FunctionInput[router.CCIPSendArgs]{
				ChainSelector: evmChain.Selector,
				Address:       r,
				Args:          ccipSendArgs,
			})
			require.NoError(t, err, "ExecuteOperation should not error")

			// Send CCIP message with value
			ccipSendArgs.Value = fee.Output
			_, err = operations.ExecuteOperation(bundle, router.CCIPSend, evmChain, contract.FunctionInput[router.CCIPSendArgs]{
				ChainSelector: evmChain.Selector,
				Address:       r,
				Args:          ccipSendArgs,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
		})
	}
}
