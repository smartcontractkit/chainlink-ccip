package tokens_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	outbound uint8 = 0
	inbound  uint8 = 1
)

func TestConfigureTokenForTransfers(t *testing.T) {
	t.Run("happy path - configure token for transfers with 2 remote chains", func(t *testing.T) {
		chainSel := uint64(5009297550715157269)        // Main chain
		remoteChainSel1 := uint64(4949039107694359620) // First remote chain
		remoteChainSel2 := uint64(6433500567565415381) // Second remote chain

		e, err := environment.New(t.Context(),
			environment.WithEVMSimulated(t, []uint64{chainSel}),
		)
		require.NoError(t, err, "Failed to create environment")
		require.NotNil(t, e, "Environment should be created")

		// Deploy chain contracts
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
			basicDeployTokenAndPoolInput(chainReport, false),
		)
		require.NoError(t, err, "ExecuteSequence should not error")
		require.Len(t, tokenAndPoolReport.Output.Addresses, 3, "Expected 3 addresses (token, pool, advanced pool hooks)")

		tokenAddress := tokenAndPoolReport.Output.Addresses[0].Address
		tokenPoolAddress := tokenAndPoolReport.Output.Addresses[1].Address

		// Find token admin registry address
		var tokenAdminRegistryAddress string
		for _, addr := range chainReport.Output.Addresses {
			if addr.Type == datastore.ContractType(token_admin_registry.ContractType) {
				tokenAdminRegistryAddress = addr.Address
				break
			}
		}
		require.NotEmpty(t, tokenAdminRegistryAddress, "Token admin registry address should be found")

		// Prepare input for configuring token for transfers
		input := tokens_core.ConfigureTokenForTransfersInput{
			ChainSelector:    chainSel,
			TokenPoolAddress: tokenPoolAddress,
			RemoteChains: map[uint64]tokens_core.RemoteChainConfig[[]byte, string]{
				remoteChainSel1: {
					RemoteToken:                              common.LeftPadBytes(common.FromHex("0x123"), 32),
					RemotePool:                               common.LeftPadBytes(common.FromHex("0x456"), 32),
					DefaultFinalityInboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(100, 1000),
					DefaultFinalityOutboundRateLimiterConfig: testsetup.CreateRateLimiterConfig(150, 1500),
					OutboundCCVs:                             []string{"0x789"},
					InboundCCVs:                              []string{"0xabc"},
					CustomFinalityInboundRateLimiterConfig:   testsetup.CreateRateLimiterConfig(60, 600),
					CustomFinalityOutboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(70, 700),
					TokenTransferFeeConfig:                   testsetup.CreateBasicTokenTransferFeeConfig(),
				},
				remoteChainSel2: {
					RemoteToken:                              common.LeftPadBytes(common.FromHex("0x321"), 32),
					RemotePool:                               common.LeftPadBytes(common.FromHex("0x654"), 32),
					DefaultFinalityInboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(200, 2000),
					DefaultFinalityOutboundRateLimiterConfig: testsetup.CreateRateLimiterConfig(250, 2500),
					OutboundCCVs:                             []string{"0xdef"},
					InboundCCVs:                              []string{"0x012"},
					CustomFinalityInboundRateLimiterConfig:   testsetup.CreateRateLimiterConfig(80, 800),
					CustomFinalityOutboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(90, 900),
					TokenTransferFeeConfig:                   testsetup.CreateBasicTokenTransferFeeConfig(),
				},
			},
			ExternalAdmin:    "", // Use internal admin
			RegistryAddress:  tokenAdminRegistryAddress,
			MinFinalityValue: 12,
		}

		// Execute the configure token for transfers sequence
		configureReport, err := operations.ExecuteSequence(
			e.OperationsBundle,
			tokens.ConfigureTokenForTransfers,
			e.BlockChains,
			input,
		)
		require.NoError(t, err, "ExecuteSequence should not error")
		require.NotEmpty(t, configureReport.Output.BatchOps, "Expected batch operations")

		// Verify token pool configuration for remote chains
		tp, err := tp_bindings.NewTokenPool(common.HexToAddress(tokenPoolAddress), e.BlockChains.EVMChains()[chainSel].Client)
		require.NoError(t, err, "Failed to instantiate token pool contract")

		// Check supported chains
		supportedChains, err := tp.GetSupportedChains(nil)
		require.NoError(t, err, "Failed to get supported chains from token pool")
		require.Len(t, supportedChains, 2, "There should be 2 supported remote chains in the token pool")
		require.Contains(t, supportedChains, remoteChainSel1, "First remote chain should be supported")
		require.Contains(t, supportedChains, remoteChainSel2, "Second remote chain should be supported")

		// Verify configuration for first remote chain
		checkRemoteChainConfiguration(t, tp, remoteChainSel1, input.RemoteChains[remoteChainSel1])

		// Verify configuration for second remote chain
		checkRemoteChainConfiguration(t, tp, remoteChainSel2, input.RemoteChains[remoteChainSel2])

		minBlockConfirmation, err := tp.GetMinBlockConfirmation(nil)
		require.NoError(t, err, "Failed to get configured min block confirmation")
		require.Equal(t, input.MinFinalityValue, minBlockConfirmation, "Min block confirmation should match input")

		customFinalityInboundRateLimiterConfig := input.RemoteChains[remoteChainSel1].CustomFinalityInboundRateLimiterConfig
		customFinalityOutboundRateLimiterConfig := input.RemoteChains[remoteChainSel1].CustomFinalityOutboundRateLimiterConfig
		assertCustomBlockConfirmationBucket(t, tp, remoteChainSel1, &customFinalityInboundRateLimiterConfig, &customFinalityOutboundRateLimiterConfig)

		customFinalityInboundRateLimiterConfig = input.RemoteChains[remoteChainSel2].CustomFinalityInboundRateLimiterConfig
		customFinalityOutboundRateLimiterConfig = input.RemoteChains[remoteChainSel2].CustomFinalityOutboundRateLimiterConfig
		assertCustomBlockConfirmationBucket(t, tp, remoteChainSel2, &customFinalityInboundRateLimiterConfig, &customFinalityOutboundRateLimiterConfig)

		// Verify token registration in token admin registry
		tokenConfigReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			token_admin_registry.GetTokenConfig,
			e.BlockChains.EVMChains()[chainSel],
			evm_contract.FunctionInput[common.Address]{
				ChainSelector: chainSel,
				Address:       common.HexToAddress(tokenAdminRegistryAddress),
				Args:          common.HexToAddress(tokenAddress),
			},
		)
		require.NoError(t, err, "ExecuteOperation should not error")

		// Since we're using internal admin (empty ExternalAdmin), the deployer key should be the admin
		require.Equal(t, common.Address{}, tokenConfigReport.Output.PendingAdministrator, "No pending administrator should be set")
		require.Equal(t, e.BlockChains.EVMChains()[chainSel].DeployerKey.From, tokenConfigReport.Output.Administrator, "Deployer key should be the administrator")
		require.Equal(t, common.HexToAddress(tokenPoolAddress), tokenConfigReport.Output.TokenPool, "Token pool address should be set correctly")

		// Verify token address from token pool
		actualTokenAddress, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			token_pool.GetToken,
			e.BlockChains.EVMChains()[chainSel],
			evm_contract.FunctionInput[any]{
				ChainSelector: chainSel,
				Address:       common.HexToAddress(tokenPoolAddress),
			},
		)
		require.NoError(t, err, "ExecuteOperation should not error")
		require.Equal(t, common.HexToAddress(tokenAddress), actualTokenAddress.Output, "Token address should match")
	})

	t.Run("applies custom finality rate limit config when global finality unchanged", func(t *testing.T) {
		chainSel := uint64(5009297550715157269)
		remoteChainSel := uint64(4949039107694359620)

		e, err := environment.New(t.Context(),
			environment.WithEVMSimulated(t, []uint64{chainSel}),
		)
		require.NoError(t, err, "Failed to create environment")
		require.NotNil(t, e, "Environment should be created")

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

		tokenAndPoolReport, err := operations.ExecuteSequence(
			e.OperationsBundle,
			tokens.DeployTokenAndPool,
			e.BlockChains.EVMChains()[chainSel],
			basicDeployTokenAndPoolInput(chainReport, false),
		)
		require.NoError(t, err, "ExecuteSequence should not error")
		require.Len(t, tokenAndPoolReport.Output.Addresses, 3, "Expected 3 addresses (token, pool, advanced pool hooks)")

		tokenPoolAddress := tokenAndPoolReport.Output.Addresses[1].Address

		var tokenAdminRegistryAddress string
		for _, addr := range chainReport.Output.Addresses {
			if addr.Type == datastore.ContractType(token_admin_registry.ContractType) {
				tokenAdminRegistryAddress = addr.Address
				break
			}
		}
		require.NotEmpty(t, tokenAdminRegistryAddress, "Token admin registry address should be found")

		input := tokens_core.ConfigureTokenForTransfersInput{
			ChainSelector:    chainSel,
			TokenPoolAddress: tokenPoolAddress,
			RemoteChains: map[uint64]tokens_core.RemoteChainConfig[[]byte, string]{
				remoteChainSel: {
					RemoteToken:                              common.LeftPadBytes(common.FromHex("0x777"), 32),
					RemotePool:                               common.LeftPadBytes(common.FromHex("0x888"), 32),
					DefaultFinalityInboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(500, 5000),
					DefaultFinalityOutboundRateLimiterConfig: testsetup.CreateRateLimiterConfig(600, 6000),
					OutboundCCVs:                             []string{"0x999"},
					InboundCCVs:                              []string{"0xaa0"},
					CustomFinalityInboundRateLimiterConfig:   testsetup.CreateRateLimiterConfig(11, 111),
					CustomFinalityOutboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(22, 222),
					TokenTransferFeeConfig:                   testsetup.CreateBasicTokenTransferFeeConfig(),
				},
			},
			ExternalAdmin:   "",
			RegistryAddress: tokenAdminRegistryAddress,
		}

		_, err = operations.ExecuteSequence(
			e.OperationsBundle,
			tokens.ConfigureTokenForTransfers,
			e.BlockChains,
			input,
		)
		require.NoError(t, err, "ExecuteSequence should not error")

		tp, err := tp_bindings.NewTokenPool(common.HexToAddress(tokenPoolAddress), e.BlockChains.EVMChains()[chainSel].Client)
		require.NoError(t, err, "Failed to instantiate token pool contract")

		minBlockConfirmation, err := tp.GetMinBlockConfirmation(nil)
		require.NoError(t, err, "Failed to get configured min block confirmation")
		require.Equal(t, uint16(0), minBlockConfirmation, "Min block confirmation should remain default")
		customFinalityInboundRateLimiterConfig := input.RemoteChains[remoteChainSel].CustomFinalityInboundRateLimiterConfig
		customFinalityOutboundRateLimiterConfig := input.RemoteChains[remoteChainSel].CustomFinalityOutboundRateLimiterConfig
		assertCustomBlockConfirmationBucket(t, tp, remoteChainSel, &customFinalityInboundRateLimiterConfig, &customFinalityOutboundRateLimiterConfig)
	})
}

// checkRemoteChainConfiguration verifies the configuration for a remote chain on the token pool
func checkRemoteChainConfiguration(t *testing.T, tp *tp_bindings.TokenPool, remoteChainSel uint64, config tokens_core.RemoteChainConfig[[]byte, string]) {
	rateLimiterStates, err := tp.GetCurrentRateLimiterState(nil, remoteChainSel, false)
	require.NoError(t, err, "Failed to get rate limiter state")

	// Check inbound rate limiter
	inboundRateLimiter := rateLimiterStates.InboundRateLimiterState
	require.Equal(t, config.DefaultFinalityInboundRateLimiterConfig.IsEnabled, inboundRateLimiter.IsEnabled, "Inbound rate limiter enabled state should match")
	require.Equal(t, config.DefaultFinalityInboundRateLimiterConfig.Rate, inboundRateLimiter.Rate, "Inbound rate limiter rate should match")
	require.Equal(t, config.DefaultFinalityInboundRateLimiterConfig.Capacity, inboundRateLimiter.Capacity, "Inbound rate limiter capacity should match")

	// Check outbound rate limiter
	outboundRateLimiter := rateLimiterStates.OutboundRateLimiterState
	require.Equal(t, config.DefaultFinalityOutboundRateLimiterConfig.IsEnabled, outboundRateLimiter.IsEnabled, "Outbound rate limiter enabled state should match")
	require.Equal(t, config.DefaultFinalityOutboundRateLimiterConfig.Rate, outboundRateLimiter.Rate, "Outbound rate limiter rate should match")
	require.Equal(t, config.DefaultFinalityOutboundRateLimiterConfig.Capacity, outboundRateLimiter.Capacity, "Outbound rate limiter capacity should match")

	// Check remote token
	remoteToken, err := tp.GetRemoteToken(nil, remoteChainSel)
	require.NoError(t, err, "Failed to get remote token")
	require.Equal(t, config.RemoteToken, remoteToken, "Remote token should match")

	// Check remote pools
	remotePools, err := tp.GetRemotePools(nil, remoteChainSel)
	require.NoError(t, err, "Failed to get remote pools")
	require.Contains(t, remotePools, config.RemotePool, "Remote pool should be in the list of remote pools")

	// Check inbound CCVs
	inboundCCVs, err := tp.GetRequiredCCVs(nil, common.Address{}, remoteChainSel, big.NewInt(0), 0, []byte{}, inbound)
	require.NoError(t, err, "Failed to get inbound CCVs")
	for _, ccv := range config.InboundCCVs {
		require.Contains(t, inboundCCVs, common.HexToAddress(ccv), "Inbound CCV should be in the list of required inbound CCVs")
	}

	// Check outbound CCVs
	outboundCCVs, err := tp.GetRequiredCCVs(nil, common.Address{}, remoteChainSel, big.NewInt(0), 0, []byte{}, outbound)
	require.NoError(t, err, "Failed to get outbound CCVs")
	for _, ccv := range config.OutboundCCVs {
		require.Contains(t, outboundCCVs, common.HexToAddress(ccv), "Outbound CCV should be in the list of required outbound CCVs")
	}
}

func assertCustomBlockConfirmationBucket(
	t *testing.T,
	tp *tp_bindings.TokenPool,
	remoteChainSel uint64,
	expectedInbound *tokens_core.RateLimiterConfig,
	expectedOutbound *tokens_core.RateLimiterConfig,
) {
	require.NotNil(t, expectedOutbound, "expected outbound rate limiter config must be provided for selector %d", remoteChainSel)
	require.NotNil(t, expectedInbound, "expected inbound rate limiter config must be provided for selector %d", remoteChainSel)

	states, err := tp.GetCurrentRateLimiterState(nil, remoteChainSel, true)
	require.NoError(t, err, "Failed to get custom block confirmation buckets for selector %d", remoteChainSel)

	assertBucketMatchesConfig(t, states.OutboundRateLimiterState, *expectedOutbound)
	assertBucketMatchesConfig(t, states.InboundRateLimiterState, *expectedInbound)
}

func assertBucketMatchesConfig(t *testing.T, actual tp_bindings.RateLimiterTokenBucket, expected tokens_core.RateLimiterConfig) {
	require.Equal(t, expected.IsEnabled, actual.IsEnabled, "Rate limiter enabled mismatch")

	expectedCapacity := big.NewInt(0)
	if expected.Capacity != nil {
		expectedCapacity = expected.Capacity
	}
	if actual.Capacity == nil {
		actual.Capacity = big.NewInt(0)
	}
	require.Zero(t, expectedCapacity.Cmp(actual.Capacity), "Rate limiter capacity mismatch")

	expectedRate := big.NewInt(0)
	if expected.Rate != nil {
		expectedRate = expected.Rate
	}
	if actual.Rate == nil {
		actual.Rate = big.NewInt(0)
	}
	require.Zero(t, expectedRate.Cmp(actual.Rate), "Rate limiter rate mismatch")
}
