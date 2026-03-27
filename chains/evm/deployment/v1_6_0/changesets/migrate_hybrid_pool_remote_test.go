package changesets_test

import (
	"bytes"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	burn_mint_with_external_minter_token_pool_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/burn_mint_with_external_minter_token_pool"
	hybrid_with_external_minter_token_pool_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/hybrid_with_external_minter_token_pool"
	token_governor_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/token_governor"
	chainsel "github.com/smartcontractkit/chain-selectors"
	v1_6_0_changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/changesets"
	v1_6_0_burn_mint_with_external_minter_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/burn_mint_with_external_minter_token_pool"
	v1_6_0_hybrid_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/hybrid_with_external_minter_token_pool"
	token_admin_registry_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	lock_release_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/lock_release_token_pool"
	core_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	burn_mint_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type migrateHybridPoolRemoteMockReader struct {
	timelockByChain map[uint64]string
	mcmByChain      map[uint64]string
}

func (m *migrateHybridPoolRemoteMockReader) GetTimelockRef(_ cldf.Environment, chainSelector uint64, _ mcms.Input) (cldf_datastore.AddressRef, error) {
	return cldf_datastore.AddressRef{
		ChainSelector: chainSelector,
		Address:       m.timelockByChain[chainSelector],
		Type:          "Timelock",
	}, nil
}

func (m *migrateHybridPoolRemoteMockReader) GetMCMSRef(_ cldf.Environment, chainSelector uint64, _ mcms.Input) (cldf_datastore.AddressRef, error) {
	return cldf_datastore.AddressRef{
		ChainSelector: chainSelector,
		Address:       m.mcmByChain[chainSelector],
		Type:          "MCM",
	}, nil
}

func (m *migrateHybridPoolRemoteMockReader) GetChainMetadata(_ cldf.Environment, chainSelector uint64, _ mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{
		StartingOpCount: 1,
		MCMAddress:      m.mcmByChain[chainSelector],
	}, nil
}

type migrateHybridPoolRemoteFixture struct {
	env            *cldf.Environment
	hubSelector    uint64
	remoteSelector uint64
	hubChain       evm.Chain
	remoteChain    evm.Chain

	hubTokenAddress      common.Address
	hubPoolAddress       common.Address
	hubPool              *hybrid_with_external_minter_token_pool_bindings.HybridWithExternalMinterTokenPool
	oldRemotePoolAddress common.Address
	newRemotePoolAddress common.Address
	remoteTokenAddress   common.Address
	remoteChainSupply    *big.Int
	remoteTARAddress     common.Address
	remoteTAR            *token_admin_registry_bindings.TokenAdminRegistry

	reader   *migrateHybridPoolRemoteMockReader
	registry *core_changesets.MCMSReaderRegistry
}

func newMigrateHybridPoolRemoteFixture(t *testing.T) *migrateHybridPoolRemoteFixture {
	return newMigrateHybridPoolRemoteFixtureWithPreMint(t, big.NewInt(0))
}

func newMigrateHybridPoolRemoteFixtureWithPreMint(t *testing.T, remotePreMint *big.Int) *migrateHybridPoolRemoteFixture {
	t.Helper()

	hubSelector := chainsel.ETHEREUM_MAINNET.Selector
	remoteSelector := chainsel.ETHEREUM_MAINNET_BASE_1.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{hubSelector, remoteSelector}),
	)
	require.NoError(t, err)

	hubChain := e.BlockChains.EVMChains()[hubSelector]
	remoteChain := e.BlockChains.EVMChains()[remoteSelector]

	hubTokenAddress, hubTokenGovernorAddress := deployTokenAndGovernor(t, hubChain, "HUB", big.NewInt(0))
	remoteTokenAddress, remoteTokenGovernorAddress := deployTokenAndGovernor(t, remoteChain, "REM", remotePreMint)

	hubPoolAddress, tx, hubPool, err := hybrid_with_external_minter_token_pool_bindings.DeployHybridWithExternalMinterTokenPool(
		hubChain.DeployerKey,
		hubChain.Client,
		hubTokenGovernorAddress,
		hubTokenAddress,
		18,
		nil,
		common.HexToAddress("0x0000000000000000000000000000000000000011"),
		common.HexToAddress("0x0000000000000000000000000000000000000022"),
	)
	require.NoError(t, err)
	confirmTx(t, hubChain, tx, err)

	oldRemotePoolAddress, tx, _, err := lock_release_token_pool_bindings.DeployLockReleaseTokenPool(
		remoteChain.DeployerKey,
		remoteChain.Client,
		remoteTokenAddress,
		18,
		nil,
		common.HexToAddress("0x0000000000000000000000000000000000000011"),
		true,
		common.HexToAddress("0x0000000000000000000000000000000000000022"),
	)
	require.NoError(t, err)
	confirmTx(t, remoteChain, tx, err)

	newRemotePoolAddress, tx, _, err := burn_mint_with_external_minter_token_pool_bindings.DeployBurnMintWithExternalMinterTokenPool(
		remoteChain.DeployerKey,
		remoteChain.Client,
		remoteTokenGovernorAddress,
		remoteTokenAddress,
		18,
		nil,
		common.HexToAddress("0x0000000000000000000000000000000000000011"),
		common.HexToAddress("0x0000000000000000000000000000000000000022"),
	)
	require.NoError(t, err)
	confirmTx(t, remoteChain, tx, err)

	remoteTARAddress, tx, remoteTAR, err := token_admin_registry_bindings.DeployTokenAdminRegistry(remoteChain.DeployerKey, remoteChain.Client)
	require.NoError(t, err)
	confirmTx(t, remoteChain, tx, err)

	tx, err = remoteTAR.ProposeAdministrator(remoteChain.DeployerKey, remoteTokenAddress, remoteChain.DeployerKey.From)
	confirmTx(t, remoteChain, tx, err)
	tx, err = remoteTAR.AcceptAdminRole(remoteChain.DeployerKey, remoteTokenAddress)
	confirmTx(t, remoteChain, tx, err)

	reader := &migrateHybridPoolRemoteMockReader{
		timelockByChain: map[uint64]string{
			hubSelector:    hubChain.DeployerKey.From.Hex(),
			remoteSelector: remoteChain.DeployerKey.From.Hex(),
		},
		mcmByChain: map[uint64]string{
			hubSelector:    "0x00000000000000000000000000000000000000a1",
			remoteSelector: "0x00000000000000000000000000000000000000b2",
		},
	}
	registry := &core_changesets.MCMSReaderRegistry{}
	registry.RegisterMCMSReader(chainsel.FamilyEVM, reader)

	return &migrateHybridPoolRemoteFixture{
		env:                  e,
		hubSelector:          hubSelector,
		remoteSelector:       remoteSelector,
		hubChain:             hubChain,
		remoteChain:          remoteChain,
		hubTokenAddress:      hubTokenAddress,
		hubPoolAddress:       hubPoolAddress,
		hubPool:              hubPool,
		oldRemotePoolAddress: oldRemotePoolAddress,
		newRemotePoolAddress: newRemotePoolAddress,
		remoteTokenAddress:   remoteTokenAddress,
		remoteChainSupply:    new(big.Int).Set(remotePreMint),
		remoteTARAddress:     remoteTARAddress,
		remoteTAR:            remoteTAR,
		reader:               reader,
		registry:             registry,
	}
}

func deployTokenAndGovernor(t *testing.T, chain evm.Chain, symbolSuffix string, preMint *big.Int) (common.Address, common.Address) {
	t.Helper()

	preMintCopy := new(big.Int).Set(preMint)
	tokenAddress, tx, _, err := burn_mint_erc20_bindings.DeployBurnMintERC20(
		chain.DeployerKey,
		chain.Client,
		"Token"+symbolSuffix,
		"T"+symbolSuffix,
		18,
		big.NewInt(1_000_000_000_000_000_000),
		preMintCopy,
	)
	require.NoError(t, err)
	confirmTx(t, chain, tx, err)

	tokenGovernorAddress, tx, _, err := token_governor_bindings.DeployTokenGovernor(
		chain.DeployerKey,
		chain.Client,
		tokenAddress,
		big.NewInt(0),
		chain.DeployerKey.From,
	)
	require.NoError(t, err)
	confirmTx(t, chain, tx, err)

	return tokenAddress, tokenGovernorAddress
}

func confirmTx(t *testing.T, chain evm.Chain, tx *types.Transaction, callErr error) {
	t.Helper()
	_, err := cldf.ConfirmIfNoError(chain, tx, callErr)
	require.NoError(t, err)
}

func (f *migrateHybridPoolRemoteFixture) changeset(registry *core_changesets.MCMSReaderRegistry) cldf.ChangeSetV2[v1_6_0_changesets.MigrateHybridPoolRemoteConfig] {
	return v1_6_0_changesets.MigrateHybridPoolRemote(registry)
}

func (f *migrateHybridPoolRemoteFixture) validConfig(targetGroup uint8) v1_6_0_changesets.MigrateHybridPoolRemoteConfig {
	return v1_6_0_changesets.MigrateHybridPoolRemoteConfig{
		HubChainSelector:     f.hubSelector,
		HubPoolAddress:       f.hubPoolAddress.Hex(),
		RemoteChainSelector:  f.remoteSelector,
		NewRemotePoolAddress: f.newRemotePoolAddress.Hex(),
		OldRemotePoolAddress: f.oldRemotePoolAddress.Hex(),
		RemoteChainSupply:    new(big.Int).Set(f.remoteChainSupply),
		TargetGroup:          targetGroup,
		RemoteTARAddress:     f.remoteTARAddress.Hex(),
		RemoteTokenAddress:   f.remoteTokenAddress.Hex(),
		MCMS: mcms.Input{
			TimelockAction: mcms_types.TimelockActionSchedule,
			ValidUntil:     uint32(time.Now().UTC().Add(24 * time.Hour).Unix()),
			TimelockDelay:  mcms_types.MustParseDuration("1h"),
			Description:    "migrate token pool",
		},
	}
}

func (f *migrateHybridPoolRemoteFixture) addHubRemotePool(t *testing.T, pool common.Address) {
	t.Helper()

	isSupported, err := f.hubPool.IsSupportedChain(&bind.CallOpts{Context: t.Context()}, f.remoteSelector)
	require.NoError(t, err)
	if !isSupported {
		tx, err := f.hubPool.ApplyChainUpdates(
			f.hubChain.DeployerKey,
			nil,
			[]hybrid_with_external_minter_token_pool_bindings.TokenPoolChainUpdate{
				{
					RemoteChainSelector: f.remoteSelector,
					RemotePoolAddresses: [][]byte{common.LeftPadBytes(pool.Bytes(), 32)},
					RemoteTokenAddress:  common.LeftPadBytes(f.remoteTokenAddress.Bytes(), 32),
					OutboundRateLimiterConfig: hybrid_with_external_minter_token_pool_bindings.RateLimiterConfig{
						IsEnabled: false,
						Capacity:  big.NewInt(0),
						Rate:      big.NewInt(0),
					},
					InboundRateLimiterConfig: hybrid_with_external_minter_token_pool_bindings.RateLimiterConfig{
						IsEnabled: false,
						Capacity:  big.NewInt(0),
						Rate:      big.NewInt(0),
					},
				},
			},
		)
		confirmTx(t, f.hubChain, tx, err)
		return
	}

	tx, err := f.hubPool.AddRemotePool(
		f.hubChain.DeployerKey,
		f.remoteSelector,
		common.LeftPadBytes(pool.Bytes(), 32),
	)
	confirmTx(t, f.hubChain, tx, err)
}

func (f *migrateHybridPoolRemoteFixture) setTARPool(t *testing.T, pool common.Address) {
	t.Helper()
	tx, err := f.remoteTAR.SetPool(f.remoteChain.DeployerKey, f.remoteTokenAddress, pool)
	confirmTx(t, f.remoteChain, tx, err)
}

func (f *migrateHybridPoolRemoteFixture) addRawHubRemotePool(t *testing.T, pool common.Address) {
	t.Helper()
	tx, err := f.hubPool.AddRemotePool(
		f.hubChain.DeployerKey,
		f.remoteSelector,
		common.LeftPadBytes(pool.Bytes(), 32),
	)
	confirmTx(t, f.hubChain, tx, err)
}

func (f *migrateHybridPoolRemoteFixture) currentHubRemotePools(t *testing.T) [][]byte {
	t.Helper()
	remotePools, err := f.hubPool.GetRemotePools(&bind.CallOpts{Context: t.Context()}, f.remoteSelector)
	require.NoError(t, err)
	return remotePools
}

func (f *migrateHybridPoolRemoteFixture) currentHubGroup(t *testing.T) uint8 {
	t.Helper()
	group, err := f.hubPool.GetGroup(&bind.CallOpts{Context: t.Context()}, f.remoteSelector)
	require.NoError(t, err)
	return group
}

func (f *migrateHybridPoolRemoteFixture) currentTARPool(t *testing.T) common.Address {
	t.Helper()
	cfg, err := f.remoteTAR.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, f.remoteTokenAddress)
	require.NoError(t, err)
	return cfg.TokenPool
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_Valid(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	require.NoError(t, cs.VerifyPreconditions(*fixture.env, fixture.validConfig(0)))
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_InvalidInputs(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)
	cs := fixture.changeset(fixture.registry)

	t.Run("invalid hub pool address", func(t *testing.T) {
		cfg := fixture.validConfig(0)
		cfg.HubPoolAddress = "not-an-address"
		err := cs.VerifyPreconditions(*fixture.env, cfg)
		require.ErrorContains(t, err, "hubPoolAddress is not a valid hex address")
	})

	t.Run("old equals new pool address", func(t *testing.T) {
		cfg := fixture.validConfig(0)
		cfg.OldRemotePoolAddress = cfg.NewRemotePoolAddress
		err := cs.VerifyPreconditions(*fixture.env, cfg)
		require.ErrorContains(t, err, "must be different")
	})

	t.Run("nil remote chain supply", func(t *testing.T) {
		cfg := fixture.validConfig(0)
		cfg.RemoteChainSupply = nil
		err := cs.VerifyPreconditions(*fixture.env, cfg)
		require.ErrorContains(t, err, "remote chain supply must be provided")
	})

	t.Run("invalid target group", func(t *testing.T) {
		cfg := fixture.validConfig(5)
		err := cs.VerifyPreconditions(*fixture.env, cfg)
		require.ErrorContains(t, err, "target group must be 0 or 1")
	})

	t.Run("same hub and remote selector", func(t *testing.T) {
		cfg := fixture.validConfig(0)
		cfg.RemoteChainSelector = cfg.HubChainSelector
		err := cs.VerifyPreconditions(*fixture.env, cfg)
		require.ErrorContains(t, err, "must be different")
	})
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_MissingMCMSRef(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	badReader := &migrateHybridPoolRemoteMockReader{
		timelockByChain: fixture.reader.timelockByChain,
		mcmByChain: map[uint64]string{
			fixture.hubSelector: fixture.reader.mcmByChain[fixture.hubSelector],
		},
	}
	badRegistry := &core_changesets.MCMSReaderRegistry{}
	badRegistry.RegisterMCMSReader(chainsel.FamilyEVM, badReader)

	cs := fixture.changeset(badRegistry)
	err := cs.VerifyPreconditions(*fixture.env, fixture.validConfig(0))
	require.ErrorContains(t, err, "missing MCMS for remote chain")
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_UnsupportedRemoteChainOnHub(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	err := cs.VerifyPreconditions(*fixture.env, fixture.validConfig(0))
	require.ErrorContains(t, err, "is not supported on hub pool")
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_UnexpectedHubRemotePool(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.addRawHubRemotePool(t, common.HexToAddress("0x0000000000000000000000000000000000000abc"))
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	err := cs.VerifyPreconditions(*fixture.env, fixture.validConfig(0))
	require.ErrorContains(t, err, "unexpected pool")
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_HubPoolTypeAndVersionMismatch(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	cfg := fixture.validConfig(0)
	cfg.HubPoolAddress = fixture.hubTokenAddress.Hex()
	err := cs.VerifyPreconditions(*fixture.env, cfg)
	require.ErrorContains(t, err, "failed to read typeAndVersion for hub pool")
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_HubOwnerMismatch(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	badReader := &migrateHybridPoolRemoteMockReader{
		timelockByChain: map[uint64]string{
			fixture.hubSelector:    common.HexToAddress("0x0000000000000000000000000000000000000c01").Hex(),
			fixture.remoteSelector: fixture.reader.timelockByChain[fixture.remoteSelector],
		},
		mcmByChain: fixture.reader.mcmByChain,
	}
	badRegistry := &core_changesets.MCMSReaderRegistry{}
	badRegistry.RegisterMCMSReader(chainsel.FamilyEVM, badReader)

	cs := fixture.changeset(badRegistry)
	err := cs.VerifyPreconditions(*fixture.env, fixture.validConfig(0))
	require.ErrorContains(t, err, "owner")
	require.ErrorContains(t, err, "does not match timelock")
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_TARAdminMismatch(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	badReader := &migrateHybridPoolRemoteMockReader{
		timelockByChain: map[uint64]string{
			fixture.hubSelector:    fixture.reader.timelockByChain[fixture.hubSelector],
			fixture.remoteSelector: common.HexToAddress("0x0000000000000000000000000000000000000d01").Hex(),
		},
		mcmByChain: fixture.reader.mcmByChain,
	}
	badRegistry := &core_changesets.MCMSReaderRegistry{}
	badRegistry.RegisterMCMSReader(chainsel.FamilyEVM, badReader)

	cs := fixture.changeset(badRegistry)
	err := cs.VerifyPreconditions(*fixture.env, fixture.validConfig(0))
	require.ErrorContains(t, err, "TAR administrator")
	require.ErrorContains(t, err, "does not match timelock")
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_SupplyMismatch(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	cfg := fixture.validConfig(1)
	cfg.RemoteChainSupply = big.NewInt(1)
	err := cs.VerifyPreconditions(*fixture.env, cfg)
	require.ErrorContains(t, err, "does not match remote token totalSupply")
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_SupplyMismatch_WhenGroupAlreadyTarget(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	cfg := fixture.validConfig(0)
	cfg.RemoteChainSupply = big.NewInt(1)
	err := cs.VerifyPreconditions(*fixture.env, cfg)
	require.ErrorContains(t, err, "does not match remote token totalSupply")
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_LockedTokensBound(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixtureWithPreMint(t, big.NewInt(100))
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	err := cs.VerifyPreconditions(*fixture.env, fixture.validConfig(1))
	require.ErrorContains(t, err, "exceeds locked token accounting")
}

func TestMigrateHybridPoolRemote_Apply_FreshState_ProposalOnly(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	cfg := fixture.validConfig(1)
	require.NoError(t, cs.VerifyPreconditions(*fixture.env, cfg))

	out, err := cs.Apply(*fixture.env, cfg)
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)
	require.NotNil(t, out.DataStore)

	txCounts := txCountByChain(out.MCMSTimelockProposals[0].Operations)
	require.Equal(t, 3, txCounts[fixture.hubSelector])
	require.Equal(t, 1, txCounts[fixture.remoteSelector])
	require.True(t, hasAddressRefWithTypeVersion(
		out.DataStore,
		fixture.hubSelector,
		fixture.hubPoolAddress,
		cldf_datastore.ContractType(v1_6_0_hybrid_pool_ops.ContractType),
		v1_6_0_hybrid_pool_ops.Version.String(),
	))
	require.True(t, hasAddressRefWithTypeVersion(
		out.DataStore,
		fixture.remoteSelector,
		fixture.newRemotePoolAddress,
		cldf_datastore.ContractType(v1_6_0_burn_mint_with_external_minter_token_pool_ops.ContractType),
		v1_6_0_burn_mint_with_external_minter_token_pool_ops.Version.String(),
	))

	// Proposal-only writes should not mutate state during Apply.
	require.Equal(t, uint8(0), fixture.currentHubGroup(t))
	require.True(t, containsPoolBytes(fixture.currentHubRemotePools(t), common.LeftPadBytes(fixture.oldRemotePoolAddress.Bytes(), 32)))
	require.False(t, containsPoolBytes(fixture.currentHubRemotePools(t), common.LeftPadBytes(fixture.newRemotePoolAddress.Bytes(), 32)))
	require.Equal(t, fixture.oldRemotePoolAddress, fixture.currentTARPool(t))
}

func TestMigrateHybridPoolRemote_Apply_PartialState(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
	fixture.addHubRemotePool(t, fixture.newRemotePoolAddress)
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	cfg := fixture.validConfig(1)
	require.NoError(t, cs.VerifyPreconditions(*fixture.env, cfg))

	out, err := cs.Apply(*fixture.env, cfg)
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)
	require.NotNil(t, out.DataStore)

	txCounts := txCountByChain(out.MCMSTimelockProposals[0].Operations)
	require.Equal(t, 2, txCounts[fixture.hubSelector])
	require.Equal(t, 1, txCounts[fixture.remoteSelector])
}

func TestMigrateHybridPoolRemote_Apply_AlreadyComplete_NoProposal(t *testing.T) {
	fixture := newMigrateHybridPoolRemoteFixture(t)
	fixture.addHubRemotePool(t, fixture.newRemotePoolAddress)
	fixture.setTARPool(t, fixture.newRemotePoolAddress)

	cs := fixture.changeset(fixture.registry)
	cfg := fixture.validConfig(0)
	require.NoError(t, cs.VerifyPreconditions(*fixture.env, cfg))

	out, err := cs.Apply(*fixture.env, cfg)
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 0)
	require.Equal(t, uint8(0), fixture.currentHubGroup(t))
	require.True(t, containsPoolBytes(fixture.currentHubRemotePools(t), common.LeftPadBytes(fixture.newRemotePoolAddress.Bytes(), 32)))
	require.False(t, containsPoolBytes(fixture.currentHubRemotePools(t), common.LeftPadBytes(fixture.oldRemotePoolAddress.Bytes(), 32)))
	require.Equal(t, fixture.newRemotePoolAddress, fixture.currentTARPool(t))
}

func txCountByChain(ops []mcms_types.BatchOperation) map[uint64]int {
	out := make(map[uint64]int)
	for _, op := range ops {
		out[uint64(op.ChainSelector)] += len(op.Transactions)
	}
	return out
}

func hasAddressRefWithTypeVersion(
	ds cldf_datastore.MutableDataStore,
	chainSelector uint64,
	address common.Address,
	expectedType cldf_datastore.ContractType,
	expectedVersion string,
) bool {
	refs := ds.Addresses().Filter(
		cldf_datastore.AddressRefByChainSelector(chainSelector),
		cldf_datastore.AddressRefByAddress(address.Hex()),
	)
	for _, ref := range refs {
		if ref.Type == expectedType && ref.Version != nil && ref.Version.String() == expectedVersion {
			return true
		}
	}
	return false
}

func containsPoolBytes(pools [][]byte, target []byte) bool {
	for _, pool := range pools {
		if bytes.Equal(pool, target) {
			return true
		}
	}
	return false
}
