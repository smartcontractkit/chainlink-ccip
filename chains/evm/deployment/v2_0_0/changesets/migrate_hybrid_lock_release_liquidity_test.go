package changesets_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
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

// ---- Apply tests ----

// migrateTest_MockReader is a controllable MCMSReader for Apply tests.
type migrateTest_MockReader struct {
	timelockErr error
}

var _ cs_core.MCMSReader = (*migrateTest_MockReader)(nil)

func (r *migrateTest_MockReader) GetMCMSRef(_ deployment.Environment, _ uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{}, nil
}

func (r *migrateTest_MockReader) GetChainMetadata(_ deployment.Environment, _ uint64, _ mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{}, nil
}

func (r *migrateTest_MockReader) GetTimelockRef(_ deployment.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	if r.timelockErr != nil {
		return datastore.AddressRef{}, r.timelockErr
	}
	return datastore.AddressRef{
		ChainSelector: selector,
		Type:          "Timelock",
		Address:       "0x5555555555555555555555555555555555555555",
	}, nil
}

func newMigrateApplyEnv(t *testing.T) deployment.Environment {
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

func TestMigrateHybridLockReleaseLiquidity_Apply_NoMCMSReaderRegistered(t *testing.T) {
	env := newMigrateApplyEnv(t)
	var registry cs_core.MCMSReaderRegistry // no reader registered
	cs := changesets.MigrateHybridLockReleaseLiquidity(&registry)

	_, err := cs.Apply(env, validMigrateConfig())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no MCMS reader registered for chain family 'evm'")
}

func TestMigrateHybridLockReleaseLiquidity_Apply_TimelockRefFails(t *testing.T) {
	env := newMigrateApplyEnv(t)
	var registry cs_core.MCMSReaderRegistry
	registry.RegisterMCMSReader("evm", &migrateTest_MockReader{timelockErr: errors.New("timelock not found")})
	cs := changesets.MigrateHybridLockReleaseLiquidity(&registry)

	_, err := cs.Apply(env, validMigrateConfig())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to resolve timelock address")
}

func TestMigrateHybridLockReleaseLiquidity_Apply_SequenceExecutionError(t *testing.T) {
	// The sequence fails because no EVM chain for the Sepolia selector is present in BlockChains.
	env := newMigrateApplyEnv(t)
	var registry cs_core.MCMSReaderRegistry
	registry.RegisterMCMSReader("evm", &migrateTest_MockReader{})
	cs := changesets.MigrateHybridLockReleaseLiquidity(&registry)

	_, err := cs.Apply(env, validMigrateConfig())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to migrate hybrid lock-release liquidity")
}
