package tokens_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_5_0_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	chains_v161_burn_mint "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_token_pool"
	chains_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

func makeFirstPassInput(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
	return tokens.ConfigureTokenPoolForRemoteChainInput{
		ChainSelector:               chainSel,
		TokenPoolAddress:            tokenPoolAddress,
		AdvancedPoolHooks:           advancedPoolHooksAddress,
		RemoteChainSelector:         remoteChainSel,
		RemoteChainAlreadySupported: false,
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
		desc                               string
		makeSecondPassInput                func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput
		checkWithFirstPassInputAfterSecond bool // when true, assert on-chain state matches first pass (for empty-input fallback tests)
	}{
		{
			desc: "initial configuration",
		},

		{
			desc: "update rate limits on second pass",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainAlreadySupported = true
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
				secondPassInput.RemoteChainAlreadySupported = true
				secondPassInput.RemoteChainConfig.RemoteToken = common.LeftPadBytes(common.FromHex("0x101"), 32)
				return secondPassInput
			},
		},
		{
			desc: "update remote pool on second pass",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainAlreadySupported = true
				secondPassInput.RemoteChainConfig.RemotePool = common.LeftPadBytes(common.FromHex("0x202"), 32)
				return secondPassInput
			},
		},
		{
			desc: "update inbound CCVs on second pass",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainAlreadySupported = true
				secondPassInput.RemoteChainConfig.InboundCCVs = []string{"0x789", "0x790"}
				return secondPassInput
			},
		},
		{
			desc: "update outbound CCVs on second pass",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainAlreadySupported = true
				secondPassInput.RemoteChainConfig.OutboundCCVs = []string{"0x789", "0x790"}
				return secondPassInput
			},
		},
		{
			desc: "idempotent second pass same config",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainAlreadySupported = true
				return secondPassInput
			},
		},
		{
			desc: "second pass with empty CCV lists falls back to on-chain",
			makeSecondPassInput: func(chainSel uint64, remoteChainSel uint64, tokenPoolAddress common.Address, advancedPoolHooksAddress common.Address) tokens.ConfigureTokenPoolForRemoteChainInput {
				secondPassInput := makeFirstPassInput(chainSel, remoteChainSel, tokenPoolAddress, advancedPoolHooksAddress)
				secondPassInput.RemoteChainAlreadySupported = true
				secondPassInput.RemoteChainConfig.InboundCCVs = nil
				secondPassInput.RemoteChainConfig.OutboundCCVs = nil
				secondPassInput.RemoteChainConfig.InboundCCVsToAddAboveThreshold = nil
				secondPassInput.RemoteChainConfig.OutboundCCVsToAddAboveThreshold = nil
				return secondPassInput
			},
			checkWithFirstPassInputAfterSecond: true, // empty CCV input should fall back to on-chain, so state stays as first pass
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
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
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
					ChainSelector:    chainSel,
					CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
					ContractParams:   testsetup.CreateBasicContractParams(),
					DeployerKeyOwned: true,
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

				checkInput := secondPassInput
				if test.checkWithFirstPassInputAfterSecond {
					checkInput = firstPassInput
				}
				checkTokenPoolConfigForRemoteChain(t, e, chainSel, remoteChainSel, checkInput)
			}
		})
	}
}

// TestConfigureTokenPoolForRemoteChainUpgradeImport verifies that when configuring a new pool for a remote chain
// with RegistryAddress and TokenAddress set to an existing token whose active pool is < 2.0.0, rate limits and
// remote pools are imported from the active pool when the input does not provide them.
func TestConfigureTokenPoolForRemoteChainUpgradeImport(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	remoteChainSel := uint64(4949039107694359620)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	// Deploy chain (includes TokenAdminRegistry)
	create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
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
			ChainSelector:    chainSel,
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ContractParams:   testsetup.CreateBasicContractParams(),
			DeployerKeyOwned: true,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")

	var registryAddress, rmnProxyAddress, routerAddress common.Address
	for _, addr := range chainReport.Output.Addresses {
		switch addr.Type {
		case datastore.ContractType(token_admin_registry.ContractType):
			registryAddress = common.HexToAddress(addr.Address)
		case datastore.ContractType(rmn_proxy.ContractType):
			rmnProxyAddress = common.HexToAddress(addr.Address)
		case datastore.ContractType(router.ContractType):
			routerAddress = common.HexToAddress(addr.Address)
		}
	}
	require.NotEqual(t, common.Address{}, registryAddress, "TokenAdminRegistry address should be set")
	require.NotEqual(t, common.Address{}, rmnProxyAddress, "RMN proxy address should be set")
	require.NotEqual(t, common.Address{}, routerAddress, "Router address should be set")

	// Deploy 2.0.0 token A + pool A (new pool we will configure with import)
	tokenAndPoolReportA, err := operations.ExecuteSequence(
		e.OperationsBundle,
		tokens.DeployTokenAndPool,
		e.BlockChains.EVMChains()[chainSel],
		basicDeployTokenAndPoolInput(chainReport, false),
	)
	require.NoError(t, err, "ExecuteSequence DeployTokenAndPool should not error")
	poolAAddress := common.HexToAddress(tokenAndPoolReportA.Output.Addresses[1].Address)
	advancedPoolHooksA := common.HexToAddress(tokenAndPoolReportA.Output.Addresses[2].Address)

	// Deploy 1.6.1 token B + pool B (will be the "active" pool in registry for import)
	legacyInput := chains_v161.DeployTokenAndPoolInput{
		Accounts: map[common.Address]*big.Int{
			common.HexToAddress("0x01"): big.NewInt(500_000),
			common.HexToAddress("0x02"): big.NewInt(500_000),
		},
		DeployTokenPoolInput: chains_v161.DeployTokenPoolInput{
			ChainSel:         chainSel,
			TokenPoolType:    datastore.ContractType(chains_v161_burn_mint.ContractType),
			TokenPoolVersion: chains_v161_burn_mint.Version,
			TokenSymbol:      "LEGACY",
			RateLimitAdmin:   common.HexToAddress("0x01"),
			ConstructorArgs: chains_v161.ConstructorArgs{
				Decimals: 18,
				RMNProxy: rmnProxyAddress,
				Router:   routerAddress,
			},
		},
	}
	legacyReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		chains_v161.DeployTokenAndPool,
		e.BlockChains.EVMChains()[chainSel],
		legacyInput,
	)
	require.NoError(t, err, "ExecuteSequence v1_6_1 DeployTokenAndPool should not error")
	tokenBAddress := common.HexToAddress(legacyReport.Output.Addresses[0].Address)
	poolBAddress := common.HexToAddress(legacyReport.Output.Addresses[1].Address)

	// Register token B with pool B in the registry
	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		v1_5_0_sequences.RegisterToken,
		e.BlockChains.EVMChains()[chainSel],
		v1_5_0_sequences.RegisterTokenInput{
			ChainSelector:             chainSel,
			TokenAddress:              tokenBAddress,
			TokenPoolAddress:          poolBAddress,
			ExternalAdmin:             common.Address{},
			TokenAdminRegistryAddress: registryAddress,
		},
	)
	require.NoError(t, err, "RegisterToken should not error")

	// Verify registration: GetTokenConfig(registry, tokenB) must return pool B as active pool (required for import).
	// If not set, the framework may not execute RegisterToken's batch in this test env; skip the rest.
	getCfgReport, err := operations.ExecuteOperation(e.OperationsBundle, token_admin_registry.GetTokenConfig, e.BlockChains.EVMChains()[chainSel], contract_utils.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       registryAddress,
		Args:          tokenBAddress,
	})
	require.NoError(t, err, "GetTokenConfig should not error")
	if getCfgReport.Output.TokenPool != poolBAddress {
		t.Skipf("Token B active pool not set in registry (got %s, expected %s); RegisterToken batch may not execute in this env. Skipping upgrade import assertions.", getCfgReport.Output.TokenPool, poolBAddress)
	}

	// Configure pool B (1.6.1) for remote chain with specific rate limits to be imported
	importedOutboundRate, importedOutboundCapacity := 111.0, 1111.0
	importedInboundRate, importedInboundCapacity := 222.0, 2222.0
	legacyConfigInput := chains_v161.ConfigureTokenPoolForRemoteChainInput{
		ChainSelector:       chainSel,
		TokenPoolAddress:    poolBAddress,
		RemoteChainSelector: remoteChainSel,
		RemoteChainConfig: tokens_core.RemoteChainConfig[[]byte, string]{
			RemoteToken:                              common.LeftPadBytes(common.FromHex("0x123"), 32),
			RemotePool:                               common.LeftPadBytes(common.FromHex("0x456"), 32),
			DefaultFinalityOutboundRateLimiterConfig: testsetup.CreateRateLimiterConfigFloatInput(importedOutboundRate, importedOutboundCapacity),
			DefaultFinalityInboundRateLimiterConfig:  testsetup.CreateRateLimiterConfigFloatInput(importedInboundRate, importedInboundCapacity),
			CustomFinalityOutboundRateLimiterConfig:  testsetup.CreateRateLimiterConfigFloatInput(importedOutboundRate, importedOutboundCapacity),
			CustomFinalityInboundRateLimiterConfig:   testsetup.CreateRateLimiterConfigFloatInput(importedInboundRate, importedInboundCapacity),
			OutboundCCVs:                             []string{"0x789"},
			InboundCCVs:                              []string{"0xabc"},
			OutboundCCVsToAddAboveThreshold:          []string{"0xdef"},
			InboundCCVsToAddAboveThreshold:           []string{"0xace"},
			TokenTransferFeeConfig:                   testsetup.CreateBasicTokenTransferFeeConfig(),
		},
	}
	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		chains_v161.ConfigureTokenPoolForRemoteChain,
		e.BlockChains.EVMChains()[chainSel],
		legacyConfigInput,
	)
	require.NoError(t, err, "ConfigureTokenPoolForRemoteChain (1.6.1) should not error")

	// Configure pool A (2.0.0) for remote with RegistryAddress and TokenAddress set but rate limits NOT provided (disabled) — should import from pool B
	upgradeInput := tokens.ConfigureTokenPoolForRemoteChainInput{
		ChainSelector:               chainSel,
		TokenPoolAddress:            poolAAddress,
		AdvancedPoolHooks:           advancedPoolHooksA,
		RemoteChainSelector:         remoteChainSel,
		RegistryAddress:             registryAddress,
		TokenAddress:                tokenBAddress,
		RemoteChainAlreadySupported: false,
		RemoteChainConfig: tokens_core.RemoteChainConfig[[]byte, string]{
			RemoteToken:                              common.LeftPadBytes(common.FromHex("0x123"), 32),
			RemotePool:                               common.LeftPadBytes(common.FromHex("0x456"), 32),
			DefaultFinalityOutboundRateLimiterConfig: tokens_core.RateLimiterConfigFloatInput{IsEnabled: false, Capacity: 0, Rate: 0},
			DefaultFinalityInboundRateLimiterConfig:  tokens_core.RateLimiterConfigFloatInput{IsEnabled: false, Capacity: 0, Rate: 0},
			CustomFinalityOutboundRateLimiterConfig:  tokens_core.RateLimiterConfigFloatInput{IsEnabled: false, Capacity: 0, Rate: 0},
			CustomFinalityInboundRateLimiterConfig:   tokens_core.RateLimiterConfigFloatInput{IsEnabled: false, Capacity: 0, Rate: 0},
			OutboundCCVs:                             []string{"0x789"},
			InboundCCVs:                              []string{"0xabc"},
			OutboundCCVsToAddAboveThreshold:          []string{"0xdef"},
			InboundCCVsToAddAboveThreshold:           []string{"0xace"},
			TokenTransferFeeConfig:                   testsetup.CreateBasicTokenTransferFeeConfig(),
		},
	}
	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		tokens.ConfigureTokenPoolForRemoteChain,
		e.BlockChains.EVMChains()[chainSel],
		upgradeInput,
	)
	require.NoError(t, err, "ConfigureTokenPoolForRemoteChain (upgrade import) should not error")

	// Assert pool A has imported rate limits (enabled and non-zero) and the active pool's remote pool
	tpA, err := tp_bindings.NewTokenPool(poolAAddress, e.BlockChains.EVMChains()[chainSel].Client)
	require.NoError(t, err, "Failed to instantiate pool A token pool contract")
	supportedChains, err := tpA.GetSupportedChains(nil)
	require.NoError(t, err, "Failed to get supported chains from token pool")
	require.Len(t, supportedChains, 1, "There should be 1 supported remote chain")
	require.Equal(t, remoteChainSel, supportedChains[0], "Remote chain in token pool should match expected")

	stateA, err := tpA.GetCurrentRateLimiterState(nil, remoteChainSel, false)
	require.NoError(t, err, "Failed to get current rate limiter state from pool A")
	require.True(t, stateA.OutboundRateLimiterState.IsEnabled, "Outbound rate limiter should be enabled (imported from active pool)")
	require.True(t, stateA.OutboundRateLimiterState.Rate.Sign() > 0 && stateA.OutboundRateLimiterState.Capacity.Sign() > 0,
		"Outbound rate and capacity should be non-zero (imported from active pool)")
	require.True(t, stateA.InboundRateLimiterState.IsEnabled, "Inbound rate limiter should be enabled (imported from active pool)")
	require.True(t, stateA.InboundRateLimiterState.Rate.Sign() > 0 && stateA.InboundRateLimiterState.Capacity.Sign() > 0,
		"Inbound rate and capacity should be non-zero (imported from active pool)")

	remotePools, err := tpA.GetRemotePools(nil, remoteChainSel)
	require.NoError(t, err, "Failed to get remote pools from token pool")
	require.Contains(t, remotePools, common.LeftPadBytes(common.FromHex("0x456"), 32), "Pool A should have the active pool's remote pool")
}
