package tokens_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
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
			DefaultFinalityInboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(200, 2000),
			DefaultFinalityOutboundRateLimiterConfig: testsetup.CreateRateLimiterConfig(100, 1000),
			CustomFinalityInboundRateLimiterConfig:   testsetup.CreateRateLimiterConfig(300, 3000),
			CustomFinalityOutboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(200, 2000),
			OutboundCCVs:                             []string{"0x789"},
			InboundCCVs:                              []string{"0xabc"},
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

	currentRateLimiterState, err := tp.GetCurrentRateLimiterState(nil, remoteChainSel, false)
	require.NoError(t, err, "Failed to get current rate limiter state from token pool")
	inboundRateLimiterReport := currentRateLimiterState.InboundRateLimiterState
	require.NoError(t, err, "Failed to get inbound rate limiter config from token pool")
	require.Equal(t, input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.IsEnabled, inboundRateLimiterReport.IsEnabled, "Inbound rate limiter enabled state should match")
	require.Equal(t, input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Rate, inboundRateLimiterReport.Rate, "Inbound rate limiter rate should match")
	require.Equal(t, input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Capacity, inboundRateLimiterReport.Capacity, "Inbound rate limiter capacity should match")

	outboundRateLimiterReport := currentRateLimiterState.OutboundRateLimiterState
	require.NoError(t, err, "Failed to get outbound rate limiter config from token pool")
	require.Equal(t, input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.IsEnabled, outboundRateLimiterReport.IsEnabled, "Outbound rate limiter enabled state should match")
	require.Equal(t, input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Rate, outboundRateLimiterReport.Rate, "Outbound rate limiter rate should match")
	require.Equal(t, input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Capacity, outboundRateLimiterReport.Capacity, "Outbound rate limiter capacity should match")

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
				secondPassInput.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Capacity = big.NewInt(6000)
				secondPassInput.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Rate = big.NewInt(600)
				secondPassInput.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Capacity = big.NewInt(5000)
				secondPassInput.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Rate = big.NewInt(500)
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
			chainReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				e.BlockChains.EVMChains()[chainSel],
				sequences.DeployChainContractsInput{
					ChainSelector:  chainSel,
					ContractParams: testsetup.CreateBasicContractParams(),
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Deploy token and token pool
			tokenAndPoolReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployTokenAndPool,
				e.BlockChains.EVMChains()[chainSel],
				basicDeployTokenAndPoolInput(chainReport),
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
