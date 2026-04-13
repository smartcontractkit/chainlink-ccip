package usdc_post_migration_validation_test

import (
	"context"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets/usdc_post_migration_validation"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func validateUSDCCLDCS() deployment.ChangeSetV2[changesets.ValidateUSDCCLDRolloutConfig] {
	return changesets.ValidateUSDCCLDRollout()
}

func validValidateUSDCCLDRolloutConfig() changesets.ValidateUSDCCLDRolloutConfig {
	return changesets.ValidateUSDCCLDRolloutConfig{
		Chains: map[uint64]changesets.ValidateUSDCCLDRolloutChainConfig{
			chainsel.ETHEREUM_MAINNET.Selector: {
				USDCToken: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				TokenAdminRegistryRef: datastore.AddressRef{
					Type:    datastore.ContractType("TokenAdminRegistry"),
					Version: semver.MustParse("1.5.0"),
				},
				ExpectedTokenPool:  "0x2222222222222222222222222222222222222222",
				USDCTokenPoolProxy: "0x3333333333333333333333333333333333333333",
				ExpectedProxyPools: &changesets.ValidateUSDCCLDRolloutProxyPoolAddresses{
					CCTPV1Pool:            "0x4444444444444444444444444444444444444444",
					CCTPV2Pool:            "0x5555555555555555555555555555555555555555",
					CCTPV2PoolWithCCV:     "0x6666666666666666666666666666666666666666",
					SiloedLockReleasePool: "0x7777777777777777777777777777777777777777",
				},
				RemoteChains: map[uint64]changesets.ValidateUSDCCLDRolloutRemoteChainConfig{
					chainsel.POLYGON_MAINNET.Selector: {
						ExpectedMechanism:   "CCTP_V2",
						ExpectedRemoteToken: "0x9999999999999999999999999999999999999999",
						ExpectedRemotePools: []string{"0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa"},
					},
				},
			},
		},
	}
}

func newValidateUSDCCLDTestEnv(t *testing.T) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	return deployment.Environment{
		Name:        "test",
		BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{}),
		DataStore:   datastore.NewMemoryDataStore().Seal(),
		Logger:      lggr,
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
	}
}

func TestValidateUSDCCLDRollout_VerifyPreconditions_MissingChainInEnvironment(t *testing.T) {
	env := newValidateUSDCCLDTestEnv(t)
	cfg := validValidateUSDCCLDRolloutConfig()

	err := validateUSDCCLDCS().VerifyPreconditions(env, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "chain selector is not available in environment")
}

func TestValidateUSDCCLDRollout_VerifyPreconditions_InvalidMechanism(t *testing.T) {
	env := newValidateUSDCCLDTestEnv(t)
	cfg := validValidateUSDCCLDRolloutConfig()
	chainCfg := cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector]
	chainCfg.RemoteChains[chainsel.POLYGON_MAINNET.Selector] = changesets.ValidateUSDCCLDRolloutRemoteChainConfig{
		ExpectedMechanism: "BAD_MECHANISM",
	}
	cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector] = chainCfg

	err := validateUSDCCLDCS().VerifyPreconditions(env, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported mechanism")
}

func TestValidateUSDCCLDRollout_VerifyPreconditions_TokenAdminRegistryRef_Required(t *testing.T) {
	env := newValidateUSDCCLDTestEnv(t)
	cfg := validValidateUSDCCLDRolloutConfig()
	chainCfg := cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector]
	chainCfg.TokenAdminRegistryRef = datastore.AddressRef{}
	cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector] = chainCfg

	err := validateUSDCCLDCS().VerifyPreconditions(env, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "TokenAdminRegistryRef: value is required")
}

func TestValidateUSDCCLDRollout_VerifyPreconditions_TokenAdminRegistryRef_WrongVersion(t *testing.T) {
	env := newValidateUSDCCLDTestEnv(t)
	cfg := validValidateUSDCCLDRolloutConfig()
	chainCfg := cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector]
	chainCfg.TokenAdminRegistryRef = datastore.AddressRef{
		Type:    datastore.ContractType("TokenAdminRegistry"),
		Version: semver.MustParse("1.5.1"),
	}
	cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector] = chainCfg

	err := validateUSDCCLDCS().VerifyPreconditions(env, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "TokenAdminRegistryRef.Version: expected")
}

func TestValidateUSDCCLDRollout_VerifyPreconditions_InvalidTokenPoolKind(t *testing.T) {
	env := newValidateUSDCCLDTestEnv(t)
	cfg := validValidateUSDCCLDRolloutConfig()
	chainCfg := cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector]
	chainCfg.ExpectedTokenPoolKind = &changesets.ValidateUSDCCLDRolloutPoolKindConfig{Kind: "BAD_POOL_KIND"}
	cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector] = chainCfg

	err := validateUSDCCLDCS().VerifyPreconditions(env, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported pool kind")
}

func TestValidateUSDCCLDRollout_VerifyPreconditions_InvalidRemotePoolKind(t *testing.T) {
	env := newValidateUSDCCLDTestEnv(t)
	cfg := validValidateUSDCCLDRolloutConfig()
	chainCfg := cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector]
	remoteCfg := chainCfg.RemoteChains[chainsel.POLYGON_MAINNET.Selector]
	remoteCfg.ExpectedRemotePoolKinds = []changesets.ValidateUSDCCLDRolloutPoolKindConfig{
		{Kind: "BAD_POOL_KIND"},
	}
	chainCfg.RemoteChains[chainsel.POLYGON_MAINNET.Selector] = remoteCfg
	cfg.Chains[chainsel.ETHEREUM_MAINNET.Selector] = chainCfg

	err := validateUSDCCLDCS().VerifyPreconditions(env, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported pool kind")
}

func TestValidateUSDCCLDRollout_VerifyPreconditions_InvalidLiquidityAmount(t *testing.T) {
	env := newValidateUSDCCLDTestEnv(t)
	cfg := changesets.ValidateUSDCCLDRolloutConfig{
		HomeChainLiquidity: &changesets.ValidateUSDCCLDRolloutHomeChainLiquidityConfig{
			ChainSelector:              chainsel.ETHEREUM_MAINNET.Selector,
			USDCToken:                  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			HybridLockReleaseTokenPool: "0x1111111111111111111111111111111111111111",
			SiloedUSDCTokenPool:        "0x2222222222222222222222222222222222222222",
			Checks: map[uint64]changesets.ValidateUSDCCLDRolloutLiquidityLaneCheck{
				chainsel.POLYGON_MAINNET.Selector: {
					ExpectedWithdrawAmount: "not-a-number",
				},
			},
		},
	}

	err := validateUSDCCLDCS().VerifyPreconditions(env, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid decimal integer")
}

func TestValidateUSDCCLDRollout_Apply_MissingEVMChain(t *testing.T) {
	env := newValidateUSDCCLDTestEnv(t)
	cfg := validValidateUSDCCLDRolloutConfig()

	_, err := validateUSDCCLDCS().Apply(env, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "evm chain not available in environment")
}
