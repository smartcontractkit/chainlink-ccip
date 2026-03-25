package changesets_test

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
)

func validMCMSInput() mcms.Input {
	return mcms.Input{
		TimelockAction: mcms_types.TimelockActionSchedule,
		ValidUntil:     uint32(time.Now().Add(24 * time.Hour).UTC().Unix()),
		TimelockDelay:  mcms_types.MustParseDuration("1h"),
	}
}

func validMigrateConfig() changesets.MigrateHybridLockReleaseLiquidityConfig {
	return changesets.MigrateHybridLockReleaseLiquidityConfig{
		ChainSelector:              chainsel.ETHEREUM_TESTNET_SEPOLIA.Selector,
		HybridLockReleaseTokenPool: common.HexToAddress("0x01").Hex(),
		SiloedUSDCTokenPool:        common.HexToAddress("0x02").Hex(),
		USDCToken:                  common.HexToAddress("0x03").Hex(),
		WithdrawAmounts:            map[uint64]uint64{chainsel.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector: 1000},
		MCMS:                       validMCMSInput(),
	}
}

func migrateCS() deployment.ChangeSetV2[changesets.MigrateHybridLockReleaseLiquidityConfig] {
	return changesets.MigrateHybridLockReleaseLiquidity(cs_core.GetRegistry())
}

func TestMigrateHybridLockReleaseLiquidity_Validate_Valid(t *testing.T) {
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, validMigrateConfig())
	require.NoError(t, err)
}

func TestMigrateHybridLockReleaseLiquidity_Validate_InvalidMCMS(t *testing.T) {
	cfg := validMigrateConfig()
	cfg.MCMS = mcms.Input{} // zero value: invalid TimelockAction
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "MCMS config is required")
}

func TestMigrateHybridLockReleaseLiquidity_Validate_InvalidChainSelector(t *testing.T) {
	cfg := validMigrateConfig()
	cfg.ChainSelector = 99999
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid chain selector")
}

func TestMigrateHybridLockReleaseLiquidity_Validate_UnsupportedChain(t *testing.T) {
	// Only Ethereum mainnet and Sepolia are allowed; Polygon is a valid selector but wrong chain.
	cfg := validMigrateConfig()
	cfg.ChainSelector = chainsel.POLYGON_MAINNET.Selector
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "only supported on Ethereum mainnet or Sepolia")
}

func TestMigrateHybridLockReleaseLiquidity_Validate_InvalidHybridPoolAddress(t *testing.T) {
	cfg := validMigrateConfig()
	cfg.HybridLockReleaseTokenPool = "not-an-address"
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid HybridLockReleaseTokenPool address")
}

func TestMigrateHybridLockReleaseLiquidity_Validate_InvalidSiloedPoolAddress(t *testing.T) {
	cfg := validMigrateConfig()
	cfg.SiloedUSDCTokenPool = "not-an-address"
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid SiloedUSDCTokenPool address")
}

func TestMigrateHybridLockReleaseLiquidity_Validate_InvalidUSDCTokenAddress(t *testing.T) {
	cfg := validMigrateConfig()
	cfg.USDCToken = "not-an-address"
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid USDCToken address")
}

func TestMigrateHybridLockReleaseLiquidity_Validate_EmptyWithdrawAmounts(t *testing.T) {
	cfg := validMigrateConfig()
	cfg.WithdrawAmounts = map[uint64]uint64{}
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "at least one withdraw amount must be provided")
}

func TestMigrateHybridLockReleaseLiquidity_Validate_ZeroWithdrawAmount(t *testing.T) {
	cfg := validMigrateConfig()
	cfg.WithdrawAmounts = map[uint64]uint64{chainsel.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector: 0}
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "must be greater than zero")
}

func TestMigrateHybridLockReleaseLiquidity_Validate_InvalidWithdrawChainSelector(t *testing.T) {
	cfg := validMigrateConfig()
	cfg.WithdrawAmounts = map[uint64]uint64{99999: 1000}
	err := migrateCS().VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid chain selector")
}
