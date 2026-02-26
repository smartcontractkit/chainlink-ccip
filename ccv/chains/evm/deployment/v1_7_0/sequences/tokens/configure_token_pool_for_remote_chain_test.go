package tokens_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func makeFirstPassInput(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
	return tokens.ConfigureTokenPoolForRemoteChainInput{
		ChainSelector:       chainSel,
		TokenPoolAddress:    tokenPoolAddress,
		AdvancedPoolHooks:   advancedPoolHooksAddress,
		RemoteChainSelector: remoteChainSel,
		RemoteChainConfig: tokens_core.RemoteChainConfig[[]byte, string]{
			RemoteToken:                              common.LeftPadBytes(common.FromHex("0x123"), 32),
			RemotePool:                               common.LeftPadBytes(common.FromHex("0x456"), 32),
			DefaultFinalityInboundRateLimiterConfig:  testsetup.CreateRateLimiterConfigFloatInput(200, 2000),
			DefaultFinalityOutboundRateLimiterConfig: testsetup.CreateRateLimiterConfigFloatInput(100, 1000),
			CustomFinalityInboundRateLimiterConfig:   testsetup.CreateRateLimiterConfigFloatInput(300, 3000),
			CustomFinalityOutboundRateLimiterConfig:  testsetup.CreateRateLimiterConfigFloatInput(200, 2000),
			OutboundCCVs:                             []string{"0x789"},
			InboundCCVs:                              []string{"0xabc"},
			OutboundCCVsToAddAboveThreshold:          []string{"0xdef"},
			InboundCCVsToAddAboveThreshold:           []string{"0xace"},
			TokenTransferFeeConfig:                   testsetup.CreateBasicTokenTransferFeeConfig(),
		},
	}
}

func checkTokenPoolConfigForRemoteChain(t *testing.T, e *deployment.Environment, chainSel uint64, remoteChainSel uint64, input tokens.ConfigureTokenPoolForRemoteChainInput) {
	tp, err := tp_bindings.NewTokenPool(input.TokenPoolAddress, e.BlockChains.EVMChains()[chainSel].Client)
	require.NoError(t, err, "Failed to instantiate token pool contract")
	supportedChains, err := tp.GetSupportedChains(nil)
	require.NoError(t, err, "Failed to get supported chains from token pool")
	require.Len(t, supportedChains, 1, "There should be 1 supported remote chain in the token pool")
	require.Equal(t, remoteChainSel, supportedChains[0], "Remote chain in token pool should match expected")

	// Token decimals from basicDeployTokenAndPoolInput; on-chain rate/capacity are scaled by 10^decimals (inbound also 1.1x).
	const decimals = 18
	currentRateLimiterState, err := tp.GetCurrentRateLimiterState(nil, remoteChainSel, false)
	require.NoError(t, err, "Failed to get current rate limiter state from token pool")
	inboundRateLimiterReport := currentRateLimiterState.InboundRateLimiterState
	require.Equal(t, input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.IsEnabled, inboundRateLimiterReport.IsEnabled, "Inbound rate limiter enabled state should match")
	expectedInboundRate := tokens_core.ScaleFloatToBigInt(input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Rate, decimals, 0.10)
	expectedInboundCapacity := tokens_core.ScaleFloatToBigInt(input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Capacity, decimals, 0.10)
	requireScaledRateLimiterMatch(t, expectedInboundRate, expectedInboundCapacity, inboundRateLimiterReport.Rate, inboundRateLimiterReport.Capacity, "Inbound")

	outboundRateLimiterReport := currentRateLimiterState.OutboundRateLimiterState
	require.Equal(t, input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.IsEnabled, outboundRateLimiterReport.IsEnabled, "Outbound rate limiter enabled state should match")
	expectedOutboundRate := tokens_core.ScaleFloatToBigInt(input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Rate, decimals, 0)
	expectedOutboundCapacity := tokens_core.ScaleFloatToBigInt(input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Capacity, decimals, 0)
	requireScaledRateLimiterMatch(t, expectedOutboundRate, expectedOutboundCapacity, outboundRateLimiterReport.Rate, outboundRateLimiterReport.Capacity, "Outbound")

	inboundCCVs, err := tp.GetRequiredCCVs(nil, common.Address{}, remoteChainSel, big.NewInt(0), 0, []byte{}, inbound)
	require.NoError(t, err, "Failed to get inbound CCVs from token pool")
	for _, ccv := range input.RemoteChainConfig.InboundCCVs {
		require.Contains(t, inboundCCVs, common.HexToAddress(ccv), "Inbound CCV should be in the list of required inbound CCVs")
	}

	outboundCCVs, err := tp.GetRequiredCCVs(nil, common.Address{}, remoteChainSel, big.NewInt(0), 0, []byte{}, outbound)
	require.NoError(t, err, "Failed to get outbound CCVs from token pool")
	for _, ccv := range input.RemoteChainConfig.OutboundCCVs {
		require.Contains(t, outboundCCVs, common.HexToAddress(ccv), "Outbound CCV should be in the list of required outbound CCVs")
	}

	inboundCCVsToAddAboveThreshold, err := tp.GetRequiredCCVs(nil, common.Address{}, remoteChainSel, thresholdAmountForAdditionalCCVs, 0, []byte{}, inbound)
	require.NoError(t, err, "Failed to get inbound CCVs to add above threshold from token pool")
	for _, ccv := range input.RemoteChainConfig.InboundCCVsToAddAboveThreshold {
		require.Contains(t, inboundCCVsToAddAboveThreshold, common.HexToAddress(ccv), "Inbound CCV to add above threshold should be in the list of required inbound CCVs to add above threshold")
	}

	outboundAmountForThresholdQuery := new(big.Int).Mul(thresholdAmountForAdditionalCCVs, big.NewInt(2)) // 2x so after fees we're still above threshold
	outboundCCVsToAddAboveThreshold, err := tp.GetRequiredCCVs(nil, common.Address{}, remoteChainSel, outboundAmountForThresholdQuery, 0, []byte{}, outbound)
	require.NoError(t, err, "Failed to get outbound CCVs to add above threshold from token pool")
	for _, ccv := range input.RemoteChainConfig.OutboundCCVsToAddAboveThreshold {
		require.Contains(t, outboundCCVsToAddAboveThreshold, common.HexToAddress(ccv), "Outbound CCV to add above threshold should be in the list of required outbound CCVs to add above threshold")
	}

	remoteToken, err := tp.GetRemoteToken(nil, remoteChainSel)
	require.NoError(t, err, "Failed to get remote token from token pool")
	require.Equal(t, input.RemoteChainConfig.RemoteToken, remoteToken, "Remote token should match")

	remotePools, err := tp.GetRemotePools(nil, remoteChainSel)
	require.NoError(t, err, "Failed to get remote pool from token pool")
	require.Contains(t, remotePools, input.RemoteChainConfig.RemotePool, "Remote pool should be in the list of remote pools")
}

func TestConfigureTokenPoolForRemoteChain(t *testing.T) {
	tests := []struct {
		desc                string
		makeSecondPassInput func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput
	}{
		{
			desc: "initial configuration",
		},

		{
			desc: "update rate limits on second pass",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Capacity = 6000
				secondPassInput.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Rate = 600
				secondPassInput.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Capacity = 5000
				secondPassInput.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Rate = 500
				return secondPassInput
			},
		},
		{
			desc: "update remote token on second pass",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainConfig.RemoteToken = common.LeftPadBytes(common.FromHex("0x101"), 32)
				return secondPassInput
			},
		},
		{
			desc: "update remote pool on second pass",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainConfig.RemotePool = common.LeftPadBytes(common.FromHex("0x202"), 32)
				return secondPassInput
			},
		},
		{
			desc: "update inbound CCVs on second pass",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainConfig.InboundCCVs = []string{"0x789", "0x790"}
				return secondPassInput
			},
		},
		{
			desc: "update outbound CCVs on second pass",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainConfig.OutboundCCVs = []string{"0x789", "0x790"}
				return secondPassInput
			},
		},
		{
			desc: "idempotent second pass same config",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				return makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainSel := uint64(5009297550715157269)
			remoteChainSel := uint64(4949039107694359620)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSel}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")

			// Deploy chain
			create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
				ChainSelector:  chainSel,
				Args: create2_factory.ConstructorArgs{
					AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy CREATE2Factory")
			chainReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				e.BlockChains.EVMChains()[chainSel],
				sequences.DeployChainContractsInput{
					ChainSelector:  chainSel,
					CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
					ContractParams: testsetup.CreateBasicContractParams(),
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Deploy token and token pool
			tokenAndPoolReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployTokenAndPool,
				e.BlockChains.EVMChains()[chainSel],
				basicDeployTokenAndPoolInput(chainReport, false),
			)
			require.NoError(t, err, "ExecuteSequence should not error")
			tokenPoolAddress := common.HexToAddress(tokenAndPoolReport.Output.Addresses[1].Address)
			advancedPoolHooksAddress := common.HexToAddress(tokenAndPoolReport.Output.Addresses[2].Address)

			firstPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
			_, err = operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.ConfigureTokenPoolForRemoteChain,
				e.BlockChains.EVMChains()[chainSel],
				firstPassInput,
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			checkTokenPoolConfigForRemoteChain(t, e, chainSel, remoteChainSel, firstPassInput)

			if test.makeSecondPassInput != nil {
				secondPassInput := test.makeSecondPassInput(
					chainSel,
					remoteChainSel,
					tokenPoolAddress,
					advancedPoolHooksAddress,
				)

				_, err = operations.ExecuteSequence(
					testsetup.BundleWithFreshReporter(e.OperationsBundle),
					tokens.ConfigureTokenPoolForRemoteChain,
					e.BlockChains.EVMChains()[chainSel],
					secondPassInput,
				)
				require.NoError(t, err, "ExecuteSequence should not error")

				checkTokenPoolConfigForRemoteChain(t, e, chainSel, remoteChainSel, secondPassInput)
			}
		})
	}
}
