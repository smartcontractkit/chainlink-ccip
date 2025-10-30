package deployment_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	adaptersv1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	routerops1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	rmnops1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	adaptersv1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	rmnremoteops1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
)

func TestFastCurse(t *testing.T) {
	chain1 := chainsel.TEST_90000001.Selector
	chain2 := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain1, chain2}),
	)
	require.NoError(t, err)
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
	ds := datastore.NewMemoryDataStore()
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
				},
				OffRampAdds: []routerops1_2.OffRamp{
					{
						SourceChainSelector: destChainSelector,
						OffRamp:             offRamp,
					},
				},
			},
		})
		require.NoError(t, err)
	}

	// deploy mcms
	evmDeployer := &adapters.EVMDeployer{}
	dReg := deploy.GetRegistry()
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
				Qualifier:        ptr.String("test"),
				TimelockAdmin:    evmChain1.DeployerKey.From,
			},
			chain2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test"),
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
			Qualifier:            "test",
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
		},
		Force: false,
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Curse proposal for fast curse test",
		},
	}
	curseReg := fastcurse.GetCurseRegistry()
	adv1_6_0 := adaptersv1_6_0.NewCurseAdapter()
	adv1_5_0 := adaptersv1_5_0.NewCurseAdapter()
	crInput1_6_0 := fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.6.0"),
		SubjectFamily:       chainsel.FamilyEVM,
		CurseAdapter:        adaptersv1_6_0.NewCurseAdapter(),
		CurseSubjectAdapter: adaptersv1_6_0.NewCurseAdapter(),
	}
	crInput1_5_0 := fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.5.0"),
		SubjectFamily:       chainsel.FamilyEVM,
		CurseAdapter:        adaptersv1_5_0.NewCurseAdapter(),
		CurseSubjectAdapter: adaptersv1_5_0.NewCurseAdapter(),
	}
	curseReg.RegisterNewCurse(crInput1_6_0)
	curseReg.RegisterNewCurse(crInput1_5_0)
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
}

func TestFastCurseGlobalCurseOnChain(t *testing.T) {
	chain1 := chainsel.TEST_90000004.Selector
	chain2 := chainsel.TEST_90000005.Selector
	chain3 := chainsel.TEST_90000006.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain1, chain2, chain3}),
	)
	require.NoError(t, err)
	bundle := env.OperationsBundle
	var rmnAddress common.Address
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
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnops1_5.ContractType),
		Version:       semver.MustParse("1.5.0"),
		ChainSelector: chain1,
		Address:       deployRMNOp.Output.Address,
	}))
	rmnAddress = common.HexToAddress(deployRMNOp.Output.Address)
	rmnRemoteAddresses := make(map[uint64]common.Address)
	// deploy RMNRemote 1.6 on chain2 and chain 3
	for _, chainSel := range []uint64{chain2, chain3} {
		chain = env.BlockChains.EVMChains()[chainSel]
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
			ChainSelector: chainSel,
			Address:       deployRMNRemoteOp.Output.Address,
		}))
		rmnRemoteAddresses[chainSel] = common.HexToAddress(deployRMNRemoteOp.Output.Address)
	}
	// deploy router and rmnProxy in all chains
	for _, sel := range []uint64{chain1, chain2, chain3} {
		evmChain := env.BlockChains.EVMChains()[sel]
		// mock wrapped native and rmnproxy address
		wNative := utils.RandomAddress()
		var rmnAddr common.Address
		if sel == chain1 {
			rmnAddr = rmnAddress
		} else {
			rmnAddr = rmnRemoteAddresses[sel]
		}
		deployRMNProxyOp, err := cldf_ops.ExecuteOperation(bundle, rmnproxyops.Deploy, evmChain, contract.DeployInput[rmnproxyops.ConstructorArgs]{
			ChainSelector:  evmChain.Selector,
			TypeAndVersion: deployment.NewTypeAndVersion(rmnproxyops.ContractType, *semver.MustParse("1.0.0")),
			Args: rmnproxyops.ConstructorArgs{
				RMN: rmnAddr,
			},
		})
		require.NoError(t, err)
		rmnProxy := deployRMNProxyOp.Output.Address
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			Type:          datastore.ContractType(rmnproxyops.ContractType),
			Version:       semver.MustParse("1.0.0"),
			ChainSelector: sel,
			Address:       rmnProxy,
		}))
		deployRouterOp, err := cldf_ops.ExecuteOperation(bundle, routerops1_2.Deploy, evmChain, contract.DeployInput[routerops1_2.ConstructorArgs]{
			ChainSelector:  evmChain.Selector,
			TypeAndVersion: deployment.NewTypeAndVersion(routerops1_2.ContractType, *semver.MustParse("1.2.0")),
			Args: routerops1_2.ConstructorArgs{
				WrappedNative: wNative,
				RMNProxy:      common.HexToAddress(rmnProxy),
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
		var onRampUpdates []routerops1_2.OnRamp
		var offRampAdds []routerops1_2.OffRamp
		for _, otherSel := range []uint64{chain1, chain2, chain3} {
			if sel != otherSel {
				onRamp := utils.RandomAddress()
				offRamp := utils.RandomAddress()
				onRampUpdates = append(onRampUpdates, routerops1_2.OnRamp{
					DestChainSelector: otherSel,
					OnRamp:            onRamp,
				})
				offRampAdds = append(offRampAdds, routerops1_2.OffRamp{
					SourceChainSelector: otherSel,
					OffRamp:             offRamp,
				})
			}
		}
		_, err = cldf_ops.ExecuteOperation(bundle, routerops1_2.ApplyRampUpdates, evmChain, contract.FunctionInput[routerops1_2.ApplyRampsUpdatesArgs]{
			Address:       common.HexToAddress(routerAddr),
			ChainSelector: evmChain.Selector,
			Args: routerops1_2.ApplyRampsUpdatesArgs{
				OnRampUpdates: onRampUpdates,
				OffRampAdds:   offRampAdds,
			},
		})
		require.NoError(t, err)
	}

	// deploy mcms
	evmDeployer := &adapters.EVMDeployer{}
	dReg := deploy.GetRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, evmDeployer)
	cs := deploy.DeployMCMS(dReg)
	mcmsChainInput := make(map[uint64]deploy.MCMSDeploymentConfigPerChain)
	for _, sel := range []uint64{chain1, chain2, chain3} {
		evmChain := env.BlockChains.EVMChains()[sel]
		mcmsChainInput[sel] = deploy.MCMSDeploymentConfigPerChain{
			Canceller:        testhelpers.SingleGroupMCMS(),
			Bypasser:         testhelpers.SingleGroupMCMS(),
			Proposer:         testhelpers.SingleGroupMCMS(),
			TimelockMinDelay: big.NewInt(0),
			Qualifier:        ptr.String("test"),
			TimelockAdmin:    evmChain.DeployerKey.From,
		}
	}
	output, err := cs.Apply(*env, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains:         mcmsChainInput,
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
			{
				ChainSelector: chain3,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(rmnremoteops1_6.ContractType),
						Version: semver.MustParse("1.6.0"),
					},
				},
				ProposedOwner: timelockAddrs[chain3],
			},
		},
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
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
	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(cr, mcmsRegistry)
	output, err = transferOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
	t.Logf("Transferred ownership of RMN and RMNRemote to respective MCMS")
	// now generate a global curse on chain 3
	curseCfg := fastcurse.GlobalCurseOnNetworkInput{
		ChainSelectors: map[uint64]*semver.Version{
			chain3: semver.MustParse("1.6.0"),
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Curse proposal for fast curse test",
		},
	}
	curseReg := fastcurse.GetCurseRegistry()
	adv1_6_0 := adaptersv1_6_0.NewCurseAdapter()
	adv1_5_0 := adaptersv1_5_0.NewCurseAdapter()
	crInput1_6_0 := fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.6.0"),
		SubjectFamily:       chainsel.FamilyEVM,
		CurseAdapter:        adaptersv1_6_0.NewCurseAdapter(),
		CurseSubjectAdapter: adaptersv1_6_0.NewCurseAdapter(),
	}
	crInput1_5_0 := fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.5.0"),
		SubjectFamily:       chainsel.FamilyEVM,
		CurseAdapter:        adaptersv1_5_0.NewCurseAdapter(),
		CurseSubjectAdapter: adaptersv1_5_0.NewCurseAdapter(),
	}
	curseReg.RegisterNewCurse(crInput1_6_0)
	curseReg.RegisterNewCurse(crInput1_5_0)
	curseChangeset := fastcurse.GloballyCurseChainChangeset(curseReg, mcmsRegistry)
	output, err = curseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the subjects were actually cursed
	evmChain1 := env.BlockChains.EVMChains()[chain1]
	evmChain2 := env.BlockChains.EVMChains()[chain2]
	evmChain3 := env.BlockChains.EVMChains()[chain3]
	rmnC, err := rmn_contract.NewRMNContract(rmnAddress, evmChain1.Client)
	require.NoError(t, err)
	// chain2 subject should not be cursed
	isCursed, err := rmnC.IsCursed(nil, adv1_5_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should be cursed on rmn in chain1")
	// chain3 subject should be cursed
	isCursed, err = rmnC.IsCursed(nil, adv1_5_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain3 should be cursed on rmn in chain1")

	rmnRemoteC, err := rmn_remote.NewRMNRemote(rmnRemoteAddresses[chain2], evmChain2.Client)
	require.NoError(t, err)
	// chain1 subject should not be cursed on chain2
	isCursed, err = rmnRemoteC.IsCursed(nil, adv1_6_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain1 should be cursed on rmnremote in chain2")
	// chain3 subject should be cursed on chain2
	isCursed, err = rmnRemoteC.IsCursed(nil, adv1_6_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain3 should be cursed on rmnremote in chain2")

	// chain3 should have a global curse on itself
	rmnRemoteC3, err := rmn_remote.NewRMNRemote(rmnRemoteAddresses[chain3], evmChain3.Client)
	require.NoError(t, err)
	isCursed, err = rmnRemoteC3.IsCursed(nil, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain3 should be cursed on rmnremote in chain3")

	// Now uncurse the subjects
	// reset the operation bundle to clear any cached values
	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	uncurseChangeset := fastcurse.GloballyUncurseChainChangeset(curseReg, mcmsRegistry)
	output, err = uncurseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the subjects were actually uncursed
	isCursed, err = rmnC.IsCursed(nil, adv1_5_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should be uncursed on rmn in chain1")

	isCursed, err = rmnC.IsCursed(nil, adv1_5_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain3 should be uncursed on rmn in chain1")

	isCursed, err = rmnRemoteC.IsCursed(nil, adv1_6_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain1 should be uncursed on rmnremote in chain2")
	isCursed, err = rmnRemoteC.IsCursed(nil, adv1_6_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain3 should be uncursed on rmnremote in chain2")

	isCursed, err = rmnRemoteC3.IsCursed(nil, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.False(t, isCursed, "global curse on chain3 should be uncursed on rmnremote in chain3")
}
