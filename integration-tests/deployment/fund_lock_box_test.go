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
	testsetupV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	erc20LockBoxBindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	bnmERC20DripBindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"
)

func TestFundLockBox_SiloedAndUnsiloed(t *testing.T) {
	const (
		symbol      = "FUNDLB"
		remoteChain = uint64(4949039107694359620)
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

	// Deploy a siloed pool with one silo group (the remote chain) plus the unsiloed bucket.
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
					LockBoxGroups:      [][]uint64{{remoteChain}},
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
	tokenAddr := common.HexToAddress(tokenRef.Address)

	lockBoxRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType("ERC20LockBox"),
		Version:       cciputils.Version_2_0_0,
		Qualifier:     "FUNDLB-silo(4949039107694359620)",
	}, chainSel, datastore_utils.FullRef)
	require.NoError(t, err)
	lockBoxAddr := common.HexToAddress(lockBoxRef.Address)

	mcmsReader, ok := changesets.GetRegistry().GetMCMSReader(chainsel.FamilyEVM)
	require.True(t, ok)
	timelockRef, err := mcmsReader.GetTimelockRef(*e, chainSel, mcms.Input{Qualifier: cciputils.CLLQualifier})
	require.NoError(t, err)
	timelockAddr := common.HexToAddress(timelockRef.Address)

	bnmERC20Drip, err := bnmERC20DripBindings.NewBurnMintERC20WithDrip(tokenAddr, chain.Client)
	require.NoError(t, err)

	siloedAmount := big.NewInt(1e18)
	unsiloedAmount := big.NewInt(2e18)

	for range 3 {
		tx, err := bnmERC20Drip.Drip(chain.DeployerKey, timelockAddr)
		require.NoError(t, err)
		_, err = chain.Confirm(tx)
		require.NoError(t, err)
	}

	fundChangeset := v2changesets.FundLockBox(changesets.GetRegistry())
	fundOut, err := fundChangeset.Apply(*e, changesets.WithMCMS[v2changesets.FundLockBoxCfg]{
		MCMS: NewDefaultInputForMCMS("Fund lockbox"),
		Cfg: v2changesets.FundLockBoxCfg{
			Input: []v2changesets.ChainLockBoxFunding{
				{
					Selector:       chainSel,
					LockBoxAddress: lockBoxAddr,
					TokenAddress:   tokenAddr,
					Deposits: []v2changesets.LockBoxDeposit{
						{RemoteChainSelector: remoteChain, Amount: siloedAmount},
						{RemoteChainSelector: 0, Amount: unsiloedAmount},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *e, fundOut.MCMSTimelockProposals, false)
	MergeAddresses(t, e, fundOut.DataStore)

	lockBoxBindings, err := erc20LockBoxBindings.NewERC20LockBox(lockBoxAddr, chain.Client)
	require.NoError(t, err)
	lockBoxBal, err := bnmERC20Drip.BalanceOf(&bind.CallOpts{Context: t.Context()}, lockBoxAddr)
	require.NoError(t, err)
	require.Equal(t, new(big.Int).Add(siloedAmount, unsiloedAmount), lockBoxBal,
		"lockbox should hold the sum of the siloed and unsiloed deposits")

	authCallers, err := lockBoxBindings.GetAllAuthorizedCallers(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Contains(t, authCallers, timelockAddr, "timelock should be an authorized caller after funding")

	e.OperationsBundle = testsetupV2_0_0.BundleWithFreshReporter(e.OperationsBundle)
	fundOut2, err := fundChangeset.Apply(*e, changesets.WithMCMS[v2changesets.FundLockBoxCfg]{
		MCMS: NewDefaultInputForMCMS("Fund lockbox again"),
		Cfg: v2changesets.FundLockBoxCfg{
			Input: []v2changesets.ChainLockBoxFunding{
				{
					Selector:       chainSel,
					LockBoxAddress: lockBoxAddr,
					TokenAddress:   tokenAddr,
					Deposits: []v2changesets.LockBoxDeposit{
						{RemoteChainSelector: remoteChain, Amount: siloedAmount},
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

func TestFundLockBox_VerifyPreconditions(t *testing.T) {
	chainSel := chainsel.TEST_90000001.Selector
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{chainSel}))
	require.NoError(t, err)

	cs := v2changesets.FundLockBox(changesets.GetRegistry())
	lockBoxAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	tokenAddr := common.HexToAddress("0x2222222222222222222222222222222222222222")

	validEntry := func() v2changesets.ChainLockBoxFunding {
		return v2changesets.ChainLockBoxFunding{
			Selector:       chainSel,
			LockBoxAddress: lockBoxAddr,
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
			name: "rejects_zero_lockbox_address",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, TokenAddress: tokenAddr, Deposits: []v2changesets.LockBoxDeposit{{Amount: big.NewInt(1)}}},
				},
			}},
			errors: []string{"zero lockbox address"},
		},
		{
			name: "rejects_zero_token_address",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, LockBoxAddress: lockBoxAddr, Deposits: []v2changesets.LockBoxDeposit{{Amount: big.NewInt(1)}}},
				},
			}},
			errors: []string{"zero token address"},
		},
		{
			name: "rejects_no_deposits",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, LockBoxAddress: lockBoxAddr, TokenAddress: tokenAddr},
				},
			}},
			errors: []string{"no deposits listed"},
		},
		{
			name: "rejects_non_positive_amount",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, LockBoxAddress: lockBoxAddr, TokenAddress: tokenAddr, Deposits: []v2changesets.LockBoxDeposit{{Amount: big.NewInt(0)}}},
				},
			}},
			errors: []string{"amount must be positive"},
		},
		{
			name: "rejects_duplicate_bucket",
			input: changesets.WithMCMS[v2changesets.FundLockBoxCfg]{MCMS: mcms.Input{}, Cfg: v2changesets.FundLockBoxCfg{
				Input: []v2changesets.ChainLockBoxFunding{
					{Selector: chainSel, LockBoxAddress: lockBoxAddr, TokenAddress: tokenAddr, Deposits: []v2changesets.LockBoxDeposit{
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
