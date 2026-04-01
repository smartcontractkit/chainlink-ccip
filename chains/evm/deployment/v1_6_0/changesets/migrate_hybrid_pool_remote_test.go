package changesets_test

import (
	"bytes"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	burn_mint_with_external_minter_token_pool_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/burn_mint_with_external_minter_token_pool"
	hybrid_with_external_minter_token_pool_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/hybrid_with_external_minter_token_pool"
	token_governor_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/token_governor"
	chainsel "github.com/smartcontractkit/chain-selectors"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
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
	_, err = hubChain.Confirm(tx)
	require.NoError(t, err)

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
	_, err = remoteChain.Confirm(tx)
	require.NoError(t, err)

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
	_, err = remoteChain.Confirm(tx)
	require.NoError(t, err)

	remoteTARAddress, tx, remoteTAR, err := token_admin_registry_bindings.DeployTokenAdminRegistry(remoteChain.DeployerKey, remoteChain.Client)
	require.NoError(t, err)
	_, err = remoteChain.Confirm(tx)
	require.NoError(t, err)

	tx, err = remoteTAR.ProposeAdministrator(remoteChain.DeployerKey, remoteTokenAddress, remoteChain.DeployerKey.From)
	require.NoError(t, err)
	_, err = remoteChain.Confirm(tx)
	require.NoError(t, err)
	tx, err = remoteTAR.AcceptAdminRole(remoteChain.DeployerKey, remoteTokenAddress)
	require.NoError(t, err)
	_, err = remoteChain.Confirm(tx)
	require.NoError(t, err)

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

	// Register the TAR in the datastore so the changeset can resolve it.
	ds := cldf_datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(cldf_datastore.AddressRef{
		ChainSelector: remoteSelector,
		Type:          cldf_datastore.ContractType(tar_ops.ContractType),
		Version:       tar_ops.Version,
		Address:       remoteTARAddress.Hex(),
	}))
	e.DataStore = ds.Seal()

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
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	tokenGovernorAddress, tx, _, err := token_governor_bindings.DeployTokenGovernor(
		chain.DeployerKey,
		chain.Client,
		tokenAddress,
		big.NewInt(0),
		chain.DeployerKey.From,
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	return tokenAddress, tokenGovernorAddress
}

func (f *migrateHybridPoolRemoteFixture) changeset(registry *core_changesets.MCMSReaderRegistry) cldf.ChangeSetV2[v1_6_0_changesets.MigrateHybridPoolRemoteConfig] {
	return v1_6_0_changesets.MigrateHybridPoolRemote(registry)
}

func (f *migrateHybridPoolRemoteFixture) validConfig(targetGroup uint8) v1_6_0_changesets.MigrateHybridPoolRemoteConfig {
	return v1_6_0_changesets.MigrateHybridPoolRemoteConfig{
		HubChainSelector:     f.hubSelector,
		HubPoolAddress:       f.hubPoolAddress,
		RemoteChainSelector:  f.remoteSelector,
		NewRemotePoolAddress: f.newRemotePoolAddress,
		OldRemotePoolAddress: f.oldRemotePoolAddress,
		TargetGroup:          targetGroup,
		RemoteTokenAddress:   f.remoteTokenAddress,
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
		require.NoError(t, err)
		_, err = f.hubChain.Confirm(tx)
		require.NoError(t, err)
		return
	}

	tx, err := f.hubPool.AddRemotePool(
		f.hubChain.DeployerKey,
		f.remoteSelector,
		common.LeftPadBytes(pool.Bytes(), 32),
	)
	require.NoError(t, err)
	_, err = f.hubChain.Confirm(tx)
	require.NoError(t, err)
}

func (f *migrateHybridPoolRemoteFixture) setTARPool(t *testing.T, pool common.Address) {
	t.Helper()
	tx, err := f.remoteTAR.SetPool(f.remoteChain.DeployerKey, f.remoteTokenAddress, pool)
	require.NoError(t, err)
	_, err = f.remoteChain.Confirm(tx)
	require.NoError(t, err)
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

	t.Run("old equals new pool address", func(t *testing.T) {
		cfg := fixture.validConfig(0)
		cfg.OldRemotePoolAddress = cfg.NewRemotePoolAddress
		err := cs.VerifyPreconditions(*fixture.env, cfg)
		require.ErrorContains(t, err, "must be different")
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

	t.Run("zero hub pool address", func(t *testing.T) {
		cfg := fixture.validConfig(0)
		cfg.HubPoolAddress = common.Address{}
		err := cs.VerifyPreconditions(*fixture.env, cfg)
		require.ErrorContains(t, err, "hub pool address cannot be the zero address")
	})
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_OnChainStateChecks(t *testing.T) {
	tests := []struct {
		name         string
		mutateConfig func(*v1_6_0_changesets.MigrateHybridPoolRemoteConfig, *migrateHybridPoolRemoteFixture)
		expectedErr  string
	}{
		{
			name: "hub pool type and version mismatch",
			mutateConfig: func(cfg *v1_6_0_changesets.MigrateHybridPoolRemoteConfig, f *migrateHybridPoolRemoteFixture) {
				cfg.HubPoolAddress = f.hubTokenAddress
			},
			expectedErr: "failed to read typeAndVersion for hub pool",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture := newMigrateHybridPoolRemoteFixture(t)
			fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
			fixture.setTARPool(t, fixture.oldRemotePoolAddress)

			cs := fixture.changeset(fixture.registry)
			cfg := fixture.validConfig(0)
			tt.mutateConfig(&cfg, fixture)
			err := cs.VerifyPreconditions(*fixture.env, cfg)
			require.ErrorContains(t, err, tt.expectedErr)
		})
	}
}

func TestMigrateHybridPoolRemote_VerifyPreconditions_MCMSReaderErrors(t *testing.T) {
	tests := []struct {
		name         string
		makeReader   func(*migrateHybridPoolRemoteFixture) *migrateHybridPoolRemoteMockReader
		expectedErrs []string
	}{
		{
			name: "missing remote MCMS ref",
			makeReader: func(f *migrateHybridPoolRemoteFixture) *migrateHybridPoolRemoteMockReader {
				return &migrateHybridPoolRemoteMockReader{
					timelockByChain: f.reader.timelockByChain,
					mcmByChain: map[uint64]string{
						f.hubSelector: f.reader.mcmByChain[f.hubSelector],
					},
				}
			},
			expectedErrs: []string{"missing MCMS for remote chain"},
		},
		{
			name: "hub owner mismatch",
			makeReader: func(f *migrateHybridPoolRemoteFixture) *migrateHybridPoolRemoteMockReader {
				return &migrateHybridPoolRemoteMockReader{
					timelockByChain: map[uint64]string{
						f.hubSelector:    common.HexToAddress("0x0000000000000000000000000000000000000c01").Hex(),
						f.remoteSelector: f.reader.timelockByChain[f.remoteSelector],
					},
					mcmByChain: f.reader.mcmByChain,
				}
			},
			expectedErrs: []string{"owner", "does not match timelock"},
		},
		{
			name: "TAR admin mismatch",
			makeReader: func(f *migrateHybridPoolRemoteFixture) *migrateHybridPoolRemoteMockReader {
				return &migrateHybridPoolRemoteMockReader{
					timelockByChain: map[uint64]string{
						f.hubSelector:    f.reader.timelockByChain[f.hubSelector],
						f.remoteSelector: common.HexToAddress("0x0000000000000000000000000000000000000d01").Hex(),
					},
					mcmByChain: f.reader.mcmByChain,
				}
			},
			expectedErrs: []string{"TAR administrator", "does not match timelock"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture := newMigrateHybridPoolRemoteFixture(t)
			fixture.addHubRemotePool(t, fixture.oldRemotePoolAddress)
			fixture.setTARPool(t, fixture.oldRemotePoolAddress)

			badReader := tt.makeReader(fixture)
			badRegistry := &core_changesets.MCMSReaderRegistry{}
			badRegistry.RegisterMCMSReader(chainsel.FamilyEVM, badReader)

			cs := fixture.changeset(badRegistry)
			err := cs.VerifyPreconditions(*fixture.env, fixture.validConfig(0))
			for _, expected := range tt.expectedErrs {
				require.ErrorContains(t, err, expected)
			}
		})
	}
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
	fixture.setTARPool(t, fixture.oldRemotePoolAddress)

	tx, err := fixture.hubPool.AddRemotePool(
		fixture.hubChain.DeployerKey,
		fixture.remoteSelector,
		common.LeftPadBytes(common.HexToAddress("0x0000000000000000000000000000000000000abc").Bytes(), 32),
	)
	require.NoError(t, err)
	_, err = fixture.hubChain.Confirm(tx)
	require.NoError(t, err)

	cs := fixture.changeset(fixture.registry)
	err = cs.VerifyPreconditions(*fixture.env, fixture.validConfig(0))
	require.ErrorContains(t, err, "unexpected pool")
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
	sealedDS := out.DataStore.Seal()
	require.True(t, v1_6_0_changesets.AddressRefExistsWithTypeVersion(
		sealedDS,
		fixture.hubSelector,
		fixture.hubPoolAddress,
		cldf_datastore.ContractType(v1_6_0_hybrid_pool_ops.ContractType),
		v1_6_0_hybrid_pool_ops.Version,
	))
	require.True(t, v1_6_0_changesets.AddressRefExistsWithTypeVersion(
		sealedDS,
		fixture.remoteSelector,
		fixture.newRemotePoolAddress,
		cldf_datastore.ContractType(v1_6_0_burn_mint_with_external_minter_token_pool_ops.ContractType),
		v1_6_0_burn_mint_with_external_minter_token_pool_ops.Version,
	))

	// Proposal-only writes should not mutate state during Apply.
	group, err := fixture.hubPool.GetGroup(&bind.CallOpts{Context: t.Context()}, fixture.remoteSelector)
	require.NoError(t, err)
	require.Equal(t, uint8(0), group)

	remotePools, err := fixture.hubPool.GetRemotePools(&bind.CallOpts{Context: t.Context()}, fixture.remoteSelector)
	require.NoError(t, err)
	require.True(t, containsPoolBytes(remotePools, common.LeftPadBytes(fixture.oldRemotePoolAddress.Bytes(), 32)))
	require.False(t, containsPoolBytes(remotePools, common.LeftPadBytes(fixture.newRemotePoolAddress.Bytes(), 32)))

	tarCfg, err := fixture.remoteTAR.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, fixture.remoteTokenAddress)
	require.NoError(t, err)
	require.Equal(t, fixture.oldRemotePoolAddress, tarCfg.TokenPool)
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

	group, err := fixture.hubPool.GetGroup(&bind.CallOpts{Context: t.Context()}, fixture.remoteSelector)
	require.NoError(t, err)
	require.Equal(t, uint8(0), group)

	remotePools, err := fixture.hubPool.GetRemotePools(&bind.CallOpts{Context: t.Context()}, fixture.remoteSelector)
	require.NoError(t, err)
	require.True(t, containsPoolBytes(remotePools, common.LeftPadBytes(fixture.newRemotePoolAddress.Bytes(), 32)))
	require.False(t, containsPoolBytes(remotePools, common.LeftPadBytes(fixture.oldRemotePoolAddress.Bytes(), 32)))

	tarCfg, err := fixture.remoteTAR.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, fixture.remoteTokenAddress)
	require.NoError(t, err)
	require.Equal(t, fixture.newRemotePoolAddress, tarCfg.TokenPool)
}

func txCountByChain(ops []mcms_types.BatchOperation) map[uint64]int {
	out := make(map[uint64]int)
	for _, op := range ops {
		out[uint64(op.ChainSelector)] += len(op.Transactions)
	}
	return out
}

func containsPoolBytes(pools [][]byte, target []byte) bool {
	for _, pool := range pools {
		if bytes.Equal(pool, target) {
			return true
		}
	}
	return false
}
