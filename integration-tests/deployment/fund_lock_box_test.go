package deployment

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	bnmERC20DripOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	v2changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	evmtokensseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	testsetupV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	erc20LockBoxBindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	bnmERC20DripBindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"
)

// fundLockBoxTestEnv is the shared fixture for the FundLockBox tests. The pool is deployed with two
// silo groups so both lockbox kinds exist: one group holds a single remote chain (a silo), the
// other holds two (the shared/unsiloed bucket).
type fundLockBoxTestEnv struct {
	env             *cldf_deployment.Environment
	chainSel        uint64
	tokenAddr       common.Address
	poolAddr        common.Address
	siloedLockBox   common.Address
	unsiloedLockBox common.Address
	timelockAddr    common.Address
	remoteChainA    uint64
	remoteChainB    uint64
	remoteChainC    uint64
}

func setupFundLockBoxEnv(t *testing.T) fundLockBoxTestEnv {
	t.Helper()

	const (
		symbol       = "FUNDLB"
		remoteChainA = uint64(4949039107694359620)
		remoteChainB = uint64(6433500567565415381)
		remoteChainC = uint64(4051577828743386545)
	)
	chainSel := chainsel.TEST_90000001.Selector

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{chainSel}))
	require.NoError(t, err)

	SeedUltraFastCurseMCMS(t, e)
	cumulative := datastore.NewMemoryDataStore()
	DeployChainContractsV2_0_0(t, e, cumulative, chainSel)
	e.DataStore = cumulative.Seal()
	DeployMCMS(t, e, chainSel, []string{cciputils.CLLQualifier})

	chain := e.BlockChains.EVMChains()[chainSel]
	deployer := chain.DeployerKey.From
	preMint := uint64(100_000)

	// Two silo groups: {A} is a silo (one chain), {B, C} is the shared bucket (two chains).
	output, err := tokens.TokenExpansion().Apply(*e, tokens.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokens.TokenExpansionInputPerChain{
			chainSel: {
				TokenPoolVersion:      cciputils.Version_2_0_0,
				SkipOwnershipTransfer: true,
				DeployTokenInput: &tokens.DeployTokenInput{
					Name:          "Fund LockBox Token",
					Symbol:        symbol,
					Decimals:      18,
					PreMint:       &preMint,
					ExternalAdmin: deployer.Hex(),
					CCIPAdmin:     deployer.Hex(),
					Type:          bnmERC20DripOps.ContractType,
				},
				DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
					PoolType:           cciputils.SiloedLockReleaseTokenPool.String(),
					TokenPoolQualifier: symbol,
					LockBoxGroups:      [][]uint64{{remoteChainA}, {remoteChainB, remoteChainC}},
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, e, output.DataStore)

	tokenRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(bnmERC20DripOps.ContractType),
		Qualifier:     symbol,
	}, chainSel, datastore_utils.FullRef)
	require.NoError(t, err)

	poolRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(cciputils.SiloedLockReleaseTokenPool.String()),
		Version:       cciputils.Version_2_0_0,
		Qualifier:     symbol,
	}, chainSel, datastore_utils.FullRef)
	require.NoError(t, err)

	siloedLockBoxRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType("ERC20LockBox"),
		Version:       cciputils.Version_2_0_0,
		Qualifier:     evmtokensseq.LockBoxQualifier(symbol, []uint64{remoteChainA}),
	}, chainSel, datastore_utils.FullRef)
	require.NoError(t, err)

	unsiloedLockBoxRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType("ERC20LockBox"),
		Version:       cciputils.Version_2_0_0,
		Qualifier:     evmtokensseq.LockBoxQualifier(symbol, []uint64{remoteChainB, remoteChainC}),
	}, chainSel, datastore_utils.FullRef)
	require.NoError(t, err)

	mcmsReader, ok := changesets.GetRegistry().GetMCMSReader(chainsel.FamilyEVM)
	require.True(t, ok)
	timelockRef, err := mcmsReader.GetTimelockRef(*e, chainSel, mcms.Input{Qualifier: cciputils.CLLQualifier})
	require.NoError(t, err)

	return fundLockBoxTestEnv{
		env:             e,
		chainSel:        chainSel,
		tokenAddr:       common.HexToAddress(tokenRef.Address),
		poolAddr:        common.HexToAddress(poolRef.Address),
		siloedLockBox:   common.HexToAddress(siloedLockBoxRef.Address),
		unsiloedLockBox: common.HexToAddress(unsiloedLockBoxRef.Address),
		timelockAddr:    common.HexToAddress(timelockRef.Address),
		remoteChainA:    remoteChainA,
		remoteChainB:    remoteChainB,
		remoteChainC:    remoteChainC,
	}
}

// dripTimelock mints n * 1e18 of the token to the timelock, which is where the deposit pulls from.
func (f fundLockBoxTestEnv) dripTimelock(t *testing.T, n int) {
	t.Helper()

	chain := f.env.BlockChains.EVMChains()[f.chainSel]
	bnmERC20Drip, err := bnmERC20DripBindings.NewBurnMintERC20WithDrip(f.tokenAddr, chain.Client)
	require.NoError(t, err)

	for range n {
		tx, err := bnmERC20Drip.Drip(chain.DeployerKey, f.timelockAddr)
		require.NoError(t, err)
		_, err = chain.Confirm(tx)
		require.NoError(t, err)
	}
}

// TestFundLockBox_SiloedAndUnsiloed funds both lockbox kinds directly (no legacy pool to migrate
// from) and verifies each receives its deposit. It also confirms the timelock is authorized as a
// caller and that a second run is idempotent (no duplicate authorize op).
func TestFundLockBox_SiloedAndUnsiloed(t *testing.T) {
	f := setupFundLockBoxEnv(t)
	chain := f.env.BlockChains.EVMChains()[f.chainSel]

	siloedAmount := big.NewInt(1e18)
	unsiloedAmount := big.NewInt(2e18)
	f.dripTimelock(t, 3)

	fundChangeset := v2changesets.FundLockBox(changesets.GetRegistry())
	fundOut, err := fundChangeset.Apply(*f.env, changesets.WithMCMS[v2changesets.FundLockBoxCfg]{
		MCMS: NewDefaultInputForMCMS("Fund lockboxes"),
		Cfg: v2changesets.FundLockBoxCfg{
			Input: []v2changesets.ChainLockBoxFunding{
				{
					Selector:       f.chainSel,
					PoolAddress:    f.poolAddr,
					LockBoxAddress: f.siloedLockBox,
					Kind:           evmtokensseq.LockBoxKindSiloed,
					TokenAddress:   f.tokenAddr,
					Deposits: []v2changesets.LockBoxDeposit{
						{RemoteChainSelector: f.remoteChainA, Amount: siloedAmount},
					},
				},
				{
					Selector:       f.chainSel,
					PoolAddress:    f.poolAddr,
					LockBoxAddress: f.unsiloedLockBox,
					Kind:           evmtokensseq.LockBoxKindUnsiloed,
					TokenAddress:   f.tokenAddr,
					Deposits: []v2changesets.LockBoxDeposit{
						{RemoteChainSelector: f.remoteChainB, Amount: unsiloedAmount},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *f.env, fundOut.MCMSTimelockProposals, false)
	MergeAddresses(t, f.env, fundOut.DataStore)

	bnmERC20Drip, err := bnmERC20DripBindings.NewBurnMintERC20WithDrip(f.tokenAddr, chain.Client)
	require.NoError(t, err)

	siloedBal, err := bnmERC20Drip.BalanceOf(&bind.CallOpts{Context: t.Context()}, f.siloedLockBox)
	require.NoError(t, err)
	require.Equal(t, siloedAmount, siloedBal, "siloed lockbox should hold its deposit")

	unsiloedBal, err := bnmERC20Drip.BalanceOf(&bind.CallOpts{Context: t.Context()}, f.unsiloedLockBox)
	require.NoError(t, err)
	require.Equal(t, unsiloedAmount, unsiloedBal, "unsiloed lockbox should hold its deposit")

	// The timelock should be an authorized caller on both lockboxes after the deposits.
	for _, lockBox := range []common.Address{f.siloedLockBox, f.unsiloedLockBox} {
		lockBoxBindings, err := erc20LockBoxBindings.NewERC20LockBox(lockBox, chain.Client)
		require.NoError(t, err)
		authCallers, err := lockBoxBindings.GetAllAuthorizedCallers(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Contains(t, authCallers, f.timelockAddr, "timelock should be an authorized caller on %s", lockBox)
	}

	// Re-running should be idempotent: the authorize step is skipped because the timelock is
	// already an authorized caller, so the second proposal contains only the approve + deposit
	// transactions (2) rather than the full authorize + approve + deposit (3). A fresh reporter is
	// required so the authorized-caller read reflects the first run's write rather than a cached
	// report.
	f.env.OperationsBundle = testsetupV2_0_0.BundleWithFreshReporter(f.env.OperationsBundle)
	fundOut2, err := fundChangeset.Apply(*f.env, changesets.WithMCMS[v2changesets.FundLockBoxCfg]{
		MCMS: NewDefaultInputForMCMS("Fund lockbox again"),
		Cfg: v2changesets.FundLockBoxCfg{
			Input: []v2changesets.ChainLockBoxFunding{
				{
					Selector:       f.chainSel,
					PoolAddress:    f.poolAddr,
					LockBoxAddress: f.siloedLockBox,
					Kind:           evmtokensseq.LockBoxKindSiloed,
					TokenAddress:   f.tokenAddr,
					Deposits: []v2changesets.LockBoxDeposit{
						{RemoteChainSelector: f.remoteChainA, Amount: siloedAmount},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, fundOut2.MCMSTimelockProposals, 1)
	require.Len(t, fundOut2.MCMSTimelockProposals[0].Operations, 1)
	require.Len(t, fundOut2.MCMSTimelockProposals[0].Operations[0].Transactions, 2,
		"second funding run should skip the authorize step, leaving only approve + deposit")
}

func TestFundLockBox_KindValidation(t *testing.T) {
	f := setupFundLockBoxEnv(t)
	f.dripTimelock(t, 3)

	fundChangeset := v2changesets.FundLockBox(changesets.GetRegistry())

	entry := func(lockBox common.Address, kind evmtokensseq.LockBoxKind, remoteChain uint64) v2changesets.ChainLockBoxFunding {
		return v2changesets.ChainLockBoxFunding{
			Selector:       f.chainSel,
			PoolAddress:    f.poolAddr,
			LockBoxAddress: lockBox,
			Kind:           kind,
			TokenAddress:   f.tokenAddr,
			Deposits:       []v2changesets.LockBoxDeposit{{RemoteChainSelector: remoteChain, Amount: big.NewInt(1e18)}},
		}
	}

	cases := []struct {
		name   string
		entry  v2changesets.ChainLockBoxFunding
		errors []string
	}{
		{
			name:   "rejects_siloed_lockbox_declared_unsiloed",
			entry:  entry(f.siloedLockBox, evmtokensseq.LockBoxKindUnsiloed, f.remoteChainA),
			errors: []string{"declared unsiloed", "maps it to only 1 chain selector"},
		},
		{
			name:   "rejects_unsiloed_lockbox_declared_siloed",
			entry:  entry(f.unsiloedLockBox, evmtokensseq.LockBoxKindSiloed, f.remoteChainB),
			errors: []string{"declared siloed", "maps it to 2 chain selectors"},
		},
		{
			name:   "rejects_deposit_targeting_wrong_chain_for_silo",
			entry:  entry(f.siloedLockBox, evmtokensseq.LockBoxKindSiloed, f.remoteChainB),
			errors: []string{"deposit targets remote chain", "is mapped to chain"},
		},
		{
			name: "rejects_lockbox_not_mapped_on_pool",
			entry: entry(
				common.HexToAddress("0x000000000000000000000000000000000000dEaD"),
				evmtokensseq.LockBoxKindSiloed,
				f.remoteChainA,
			),
			errors: []string{"is not configured on pool"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fundChangeset.Apply(*f.env, changesets.WithMCMS[v2changesets.FundLockBoxCfg]{
				MCMS: NewDefaultInputForMCMS(tt.name),
				Cfg:  v2changesets.FundLockBoxCfg{Input: []v2changesets.ChainLockBoxFunding{tt.entry}},
			})
			require.Error(t, err)
			for _, substr := range tt.errors {
				require.Contains(t, err.Error(), substr)
			}
		})
	}
}

// TestFundLockBox_VerifyPreconditions covers the input validation rules.
func TestFundLockBox_VerifyPreconditions(t *testing.T) {
	chainSel := chainsel.TEST_90000001.Selector
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{chainSel}))
	require.NoError(t, err)

	cs := v2changesets.FundLockBox(changesets.GetRegistry())
	poolAddr := common.HexToAddress("0x3333333333333333333333333333333333333333")
	lockBoxAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	tokenAddr := common.HexToAddress("0x2222222222222222222222222222222222222222")

	validEntry := func() v2changesets.ChainLockBoxFunding {
		return v2changesets.ChainLockBoxFunding{
			Selector:       chainSel,
			PoolAddress:    poolAddr,
			LockBoxAddress: lockBoxAddr,
			Kind:           evmtokensseq.LockBoxKindSiloed,
			TokenAddress:   tokenAddr,
			Deposits:       []v2changesets.LockBoxDeposit{{RemoteChainSelector: 1, Amount: big.NewInt(1)}},
		}
	}

	cases := []struct {
		name   string
		input  changesets.WithMCMS[v2changesets.FundLockBoxCfg]
		errors []string
	}{
		{
			name:   "rejects_empty_input",
			input:  changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}},
			errors: []string{"at least one entry is required"},
		},
		{
			name: "rejects_zero_pool_address",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, LockBoxAddress: lockBoxAddr, Kind: evmtokensseq.LockBoxKindSiloed, TokenAddress: tokenAddr, Deposits: []v2changesets.LockBoxDeposit{{Amount: big.NewInt(1)}}},
				},
			}},
			errors: []string{"zero pool address"},
		},
		{
			name: "rejects_zero_lockbox_address",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, PoolAddress: poolAddr, Kind: evmtokensseq.LockBoxKindSiloed, TokenAddress: tokenAddr, Deposits: []v2changesets.LockBoxDeposit{{Amount: big.NewInt(1)}}},
				},
			}},
			errors: []string{"zero lockbox address"},
		},
		{
			name: "rejects_unknown_kind",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, PoolAddress: poolAddr, LockBoxAddress: lockBoxAddr, Kind: "bogus", TokenAddress: tokenAddr, Deposits: []v2changesets.LockBoxDeposit{{Amount: big.NewInt(1)}}},
				},
			}},
			errors: []string{"kind for chain", "must be"},
		},
		{
			name: "rejects_zero_token_address",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, PoolAddress: poolAddr, LockBoxAddress: lockBoxAddr, Kind: evmtokensseq.LockBoxKindSiloed, Deposits: []v2changesets.LockBoxDeposit{{Amount: big.NewInt(1)}}},
				},
			}},
			errors: []string{"zero token address"},
		},
		{
			name: "rejects_no_deposits",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, PoolAddress: poolAddr, LockBoxAddress: lockBoxAddr, Kind: evmtokensseq.LockBoxKindSiloed, TokenAddress: tokenAddr},
				},
			}},
			errors: []string{"no deposits listed"},
		},
		{
			name: "rejects_non_positive_amount",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, PoolAddress: poolAddr, LockBoxAddress: lockBoxAddr, Kind: evmtokensseq.LockBoxKindSiloed, TokenAddress: tokenAddr, Deposits: []v2changesets.LockBoxDeposit{{Amount: big.NewInt(0)}}},
				},
			}},
			errors: []string{"amount must be positive"},
		},
		{
			name: "rejects_duplicate_bucket",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, PoolAddress: poolAddr, LockBoxAddress: lockBoxAddr, Kind: evmtokensseq.LockBoxKindSiloed, TokenAddress: tokenAddr, Deposits: []v2changesets.LockBoxDeposit{
						{RemoteChainSelector: 1, Amount: big.NewInt(1)},
						{RemoteChainSelector: 1, Amount: big.NewInt(1)},
					}},
				},
			}},
			errors: []string{"duplicate remote chain selector"},
		},
		{
			name: "rejects_duplicate_chain_entry",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					validEntry(),
					validEntry(),
				},
			}},
			errors: []string{"duplicate entry for chain"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := cs.VerifyPreconditions(*e, tt.input)
			require.Error(t, err)
			for _, substr := range tt.errors {
				require.Contains(t, err.Error(), substr)
			}
		})
	}
}
