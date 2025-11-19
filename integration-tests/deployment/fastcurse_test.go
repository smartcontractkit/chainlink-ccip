package deployment

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	routerops1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	adaptersv1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/adapters"
	rmnops1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	adaptersv1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	rmnremoteops1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_remote"
	soladapterv1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestFastCurseSolanaAndEVM(t *testing.T) {
	chain1 := chainsel.TEST_90000001.Selector
	chain2 := chainsel.TEST_90000002.Selector
	programsPath, dstr, err := PreloadSolanaEnvironment(t, chainsel.SOLANA_MAINNET.Selector)
	require.NoError(t, err, "Failed to set up Solana environment")
	require.NotNil(t, dstr, "Datastore should be created")
	solanaChains := []uint64{
		chainsel.SOLANA_MAINNET.Selector,
	}
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain1, chain2}),
		environment.WithSolanaContainer(t, solanaChains, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err)
	require.NotNil(t, env, "Environment should be created")
	env.DataStore = dstr.Seal() // Add preloaded contracts to env datastore
	mint, _ := solana.NewRandomPrivateKey()

	dReg := deploy.GetRegistry()
	version := semver.MustParse("1.6.0")
	_, err = deploy.DeployContracts(dReg).Apply(*env, deploy.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
			chainsel.SOLANA_MAINNET.Selector: {
				Version: version,
				// LINK TOKEN CONFIG
				// token private key used to deploy the LINK token. Solana: base58 encoded private key
				TokenPrivKey: mint.String(),
				// token decimals used to deploy the LINK token
				TokenDecimals: 9,
				// FEE QUOTER CONFIG
				MaxFeeJuelsPerMsg: big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				// OFFRAMP CONFIG
				PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
			},
		},
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
	DeployMCMS(t, env, chainsel.SOLANA_MAINNET.Selector)
	SolanaTransferOwnership(t, env, chainsel.SOLANA_MAINNET.Selector)
	ds := datastore.NewMemoryDataStore()
	bundle := env.OperationsBundle
	var rmnAddress, rmnRemoteAddress common.Address
	// deploy RMN 1.5 on chain1 and RMN 1.6 on chain2, set up routers, etc.
	chain := env.BlockChains.EVMChains()[chain1]
	deployRMNOp, err := cldf_ops.ExecuteOperation(bundle, rmnops1_5.Deploy, chain, contract.DeployInput[rmnops1_5.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmnops1_5.ContractType, *semver.MustParse("1.5.0")),
		ChainSelector:  chain.Selector,
		Args: rmnops1_5.ConstructorArgs{
			RMNConfig: rmn_contract.RMNConfig{
				BlessWeightThreshold: 2,
				CurseWeightThreshold: 2,
				// setting dummy voters
				Voters: []rmn_contract.RMNVoter{
					{
						BlessWeight:   2,
						CurseWeight:   2,
						BlessVoteAddr: utils.RandomAddress(),
						CurseVoteAddr: utils.RandomAddress(),
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnops1_5.ContractType),
		Version:       semver.MustParse("1.5.0"),
		ChainSelector: chain1,
		Address:       deployRMNOp.Output.Address,
	}))
	rmnAddress = common.HexToAddress(deployRMNOp.Output.Address)
	// deploy RMNRemote 1.6 on chain2
	chain = env.BlockChains.EVMChains()[chain2]
	deployRMNRemoteOp, err := cldf_ops.ExecuteOperation(bundle, rmnremoteops1_6.Deploy, chain, contract.DeployInput[rmnremoteops1_6.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmnremoteops1_6.ContractType, *semver.MustParse("1.6.0")),
		ChainSelector:  chain.Selector,
		Args: rmnremoteops1_6.ConstructorArgs{
			LocalChainSelector: chain.Selector,
			LegacyRMN:          utils.RandomAddress(),
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnremoteops1_6.ContractType),
		Version:       semver.MustParse("1.6.0"),
		ChainSelector: chain2,
		Address:       deployRMNRemoteOp.Output.Address,
	}))
	rmnRemoteAddress = common.HexToAddress(deployRMNRemoteOp.Output.Address)
	// deploy router in both chains
	for _, sel := range []uint64{chain1, chain2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		// mock wrapped native and rmnproxy address
		wNative := utils.RandomAddress()
		rmnProxy := utils.RandomAddress()

		deployRouterOp, err := cldf_ops.ExecuteOperation(bundle, routerops1_2.Deploy, evmChain, contract.DeployInput[routerops1_2.ConstructorArgs]{
			ChainSelector:  evmChain.Selector,
			TypeAndVersion: deployment.NewTypeAndVersion(routerops1_2.ContractType, *semver.MustParse("1.2.0")),
			Args: routerops1_2.ConstructorArgs{
				WrappedNative: wNative,
				RMNProxy:      rmnProxy,
			},
		})
		require.NoError(t, err)
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			Type:          datastore.ContractType(routerops1_2.ContractType),
			Version:       semver.MustParse("1.2.0"),
			ChainSelector: sel,
			Address:       deployRouterOp.Output.Address,
		}))
		routerAddr := deployRouterOp.Output.Address
		// add some dummy onramps to the router so that chain is supported,
		// on chain1, add chain2 as supported dest chain and vice versa
		onRamp := utils.RandomAddress()
		offRamp := utils.RandomAddress()
		var destChainSelector uint64
		if sel == chain1 {
			destChainSelector = chain2
		} else {
			destChainSelector = chain1
		}
		_, err = cldf_ops.ExecuteOperation(bundle, routerops1_2.ApplyRampUpdates, evmChain, contract.FunctionInput[routerops1_2.ApplyRampsUpdatesArgs]{
			Address:       common.HexToAddress(routerAddr),
			ChainSelector: evmChain.Selector,
			Args: routerops1_2.ApplyRampsUpdatesArgs{
				OnRampUpdates: []routerops1_2.OnRamp{
					{
						DestChainSelector: destChainSelector,
						OnRamp:            onRamp,
					},
					{
						DestChainSelector: chainsel.SOLANA_MAINNET.Selector,
						OnRamp:            common.BytesToAddress([]byte(solanaProgramIDs["ccip_router"])),
					},
				},
				OffRampAdds: []routerops1_2.OffRamp{
					{
						SourceChainSelector: destChainSelector,
						OffRamp:             offRamp,
					},
					{
						SourceChainSelector: chainsel.SOLANA_MAINNET.Selector,
						OffRamp:             common.BytesToAddress([]byte(solanaProgramIDs["ccip_offramp"])),
					},
				},
			},
		})
		require.NoError(t, err)
	}

	// deploy mcms
	evmDeployer := &adapters.EVMDeployer{}
	dReg.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, evmDeployer)
	cs := deploy.DeployMCMS(dReg)
	evmChain1 := env.BlockChains.EVMChains()[chain1]
	evmChain2 := env.BlockChains.EVMChains()[chain2]
	output, err := cs.Apply(*env, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			chain1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String(deploymentutils.CLLQualifier),
				TimelockAdmin:    evmChain1.DeployerKey.From,
			},
			chain2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String(deploymentutils.CLLQualifier),
				TimelockAdmin:    evmChain2.DeployerKey.From,
			},
		},
	})
	require.NoError(t, err)
	// store addresses in ds
	allAddrRefs, err := output.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	timelockAddrs := make(map[uint64]string)
	for _, addrRef := range allAddrRefs {
		require.NoError(t, ds.Addresses().Add(addrRef))
		if addrRef.Type == datastore.ContractType(deploymentutils.RBACTimelock) {
			timelockAddrs[addrRef.ChainSelector] = addrRef.Address
		}
	}
	// update env datastore
	require.NoError(t, ds.Merge(env.DataStore))
	env.DataStore = ds.Seal()
	// transfer ownership of RMN and RMNRemote to respective MCMS
	transferOwnershipInput := deploy.TransferOwnershipInput{
		ChainInputs: []deploy.TransferOwnershipPerChainInput{
			{
				ChainSelector: chain1,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(rmnops1_5.ContractType),
						Version: semver.MustParse("1.5.0"),
					},
				},
				ProposedOwner: timelockAddrs[chain1],
			},
			{
				ChainSelector: chain2,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(rmnremoteops1_6.ContractType),
						Version: semver.MustParse("1.6.0"),
					},
				},
				ProposedOwner: timelockAddrs[chain2],
			},
		},
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Transfer ownership to timelock for fast curse test",
		},
	}

	// register chain adapter
	cr := deploy.GetTransferOwnershipRegistry()
	evmAdapter := &adapters.EVMTransferOwnershipAdapter{}
	cr.RegisterAdapter(chainsel.FamilyEVM, transferOwnershipInput.AdapterVersion, evmAdapter)
	mcmsRegistry := changesets.GetRegistry()
	evmMCMSReader := &adapters.EVMMCMSReader{}
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, evmMCMSReader)
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilySolana, &sequences.SolanaAdapter{})
	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(cr, mcmsRegistry)
	output, err = transferOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
	t.Logf("Transferred ownership of RMN and RMNRemote to respective MCMS")
	// now generate a curse proposal
	curseCfg := fastcurse.RMNCurseConfig{
		CurseActions: []fastcurse.CurseActionInput{
			{
				IsGlobalCurse:        false,
				ChainSelector:        chain1,
				SubjectChainSelector: chain2,
				Version:              semver.MustParse("1.5.0"),
			},
			{
				IsGlobalCurse:        false,
				ChainSelector:        chain2,
				SubjectChainSelector: chain1,
				Version:              semver.MustParse("1.6.0"),
			},
			{
				IsGlobalCurse:        false,
				ChainSelector:        chain2,
				SubjectChainSelector: chainsel.SOLANA_MAINNET.Selector,
				Version:              semver.MustParse("1.6.0"),
			},
		},
		Force: false,
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Curse proposal for fast curse test",
		},
	}
	curseReg := fastcurse.GetCurseRegistry()
	adv1_6_0 := adaptersv1_6_0.NewCurseAdapter()
	adv1_5_0 := adaptersv1_5_0.NewCurseAdapter()
	crInput1_6_0 := fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.6.0"),
		CurseAdapter:        adaptersv1_6_0.NewCurseAdapter(),
		CurseSubjectAdapter: adaptersv1_6_0.NewCurseAdapter(),
	}
	crInput1_5_0 := fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.5.0"),
		CurseAdapter:        adaptersv1_5_0.NewCurseAdapter(),
		CurseSubjectAdapter: adaptersv1_5_0.NewCurseAdapter(),
	}
	crInputSol_1_6_0 := fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilySolana,
		CursingVersion:      semver.MustParse("1.6.0"),
		CurseAdapter:        soladapterv1_6_0.NewCurseAdapter(),
		CurseSubjectAdapter: soladapterv1_6_0.NewCurseAdapter(),
	}

	curseReg.RegisterNewCurse(crInput1_6_0)
	curseReg.RegisterNewCurse(crInput1_5_0)
	curseReg.RegisterNewCurse(crInputSol_1_6_0)

	curseChangeset := fastcurse.CurseChangeset(curseReg, mcmsRegistry)
	output, err = curseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the subjects were actually cursed
	rmnC, err := rmn_contract.NewRMNContract(rmnAddress, evmChain1.Client)
	require.NoError(t, err)
	isCursed, err := rmnC.IsCursed(nil, adv1_5_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain2 should be cursed on rmn in chain1")

	rmnRemoteC, err := rmn_remote.NewRMNRemote(rmnRemoteAddress, evmChain2.Client)
	require.NoError(t, err)
	isCursed, err = rmnRemoteC.IsCursed(nil, adv1_6_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain1 should be cursed on rmnremote in chain2")
	t.Logf("Subjects successfully cursed %x on chain1 %d and %x on chain2 %d", adv1_5_0.SelectorToSubject(chain2), chain1, adv1_6_0.SelectorToSubject(chain1), chain2)

	isCursed, err = rmnRemoteC.IsCursed(nil, adv1_6_0.SelectorToSubject(chainsel.SOLANA_MAINNET.Selector))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on solana chain should be cursed on rmnremote in chain2")
	t.Logf("Subject successfully cursed %x on solana chain %d in rmnremote on chain2 %d", soladapterv1_6_0.NewCurseAdapter().SelectorToSubject(chainsel.SOLANA_MAINNET.Selector), chainsel.SOLANA_MAINNET.Selector, chain2)

	/*
		 Enable Solana checks later
		 // find rmn_remote address for solana chain
		solanaRMNRemoteAddrRef, err := datastore_utils.FindAndFormatRef(
			env.DataStore,
			datastore.AddressRef{
				ChainSelector: chainsel.SOLANA_MAINNET.Selector,
				Type:          datastore.ContractType(solrmnremoteops.ContractType),
				Version:       solrmnremoteops.Version,
			},
			chain2,
			solutils.ToAddress,
		)
		require.NoError(t, err)

		isCursed, err = solrmnremoteops.IsSubjectCursed(
			env.BlockChains.SolanaChains()[chainsel.SOLANA_MAINNET.Selector],
			solanaRMNRemoteAddrRef,
			solrmn_remote.CurseSubject{
				Value: soladapterv1_6_0.NewCurseAdapter().SelectorToSubject(chain2),
			},
		)
		require.NoError(t, err)
		require.True(t, isCursed, "subject on chain1 should be cursed on solana rmnremote")
		t.Logf("Subject successfully cursed %x on chain1 %d in solana rmnremote on solana chain %d", adv1_5_0.SelectorToSubject(chain2), chain1, chainsel.SOLANA_MAINNET.Selector)
	*/
	// Now uncurse the subjects
	// reset the operation bundle to clear any cached values
	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	uncurseChangeset := fastcurse.UncurseChangeset(curseReg, mcmsRegistry)
	output, err = uncurseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the subjects were actually uncursed
	isCursed, err = rmnC.IsCursed(nil, adv1_5_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should be uncursed on rmn in chain1")

	isCursed, err = rmnRemoteC.IsCursed(nil, adv1_6_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain1 should be uncursed on rmnremote in chain2")
	t.Logf("Subjects successfully uncursed %x on chain1 %d and %x on chain2 %d", adv1_5_0.SelectorToSubject(chain2), chain1, adv1_6_0.SelectorToSubject(chain1), chain2)
	/*
			 Enable Solana checks later
		isCursed, err = solrmnremoteops.IsSubjectCursed(
			env.BlockChains.SolanaChains()[chainsel.SOLANA_MAINNET.Selector],
			solanaRMNRemoteAddrRef,
			solrmn_remote.CurseSubject{
				Value: soladapterv1_6_0.NewCurseAdapter().SelectorToSubject(chain2),
			},
		)
		require.NoError(t, err)
		require.False(t, isCursed, "subject on chain1 should be cursed on solana rmnremote")
		t.Logf("Subject successfully uncursed %x on chain1 %d in solana rmnremote on solana chain %d", adv1_5_0.SelectorToSubject(chain2), chain1, chainsel.SOLANA_MAINNET.Selector)
	*/
}
