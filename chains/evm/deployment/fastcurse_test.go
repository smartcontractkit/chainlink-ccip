package deployment_test

import (
	"math/big"
	"sync"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	routerops1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	adaptersv2_1_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/adapters"
	rmnops2_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
)

// deployRMN2_1WithProxy deploys RMN 2.1.0 on the given chain together with the RMNProxy that
// fronts it, and registers both in the datastore. The proxy is mandatory for the 2.1.0 fastcurse
// adapter: it resolves the active RMN via utils.ActiveRMNAddress -> RMNProxy.getARM(). It returns
// the RMN address and the RMNProxy address.
func deployRMN2_1WithProxy(
	t *testing.T,
	bundle cldf_ops.Bundle,
	evmChain cldf_evm.Chain,
	ds *datastore.MemoryDataStore,
) (common.Address, common.Address) {
	t.Helper()
	deployRMNOp, err := cldf_ops.ExecuteOperation(bundle, rmnops2_1.Deploy, evmChain, contract.DeployInput[rmnops2_1.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmnops2_1.ContractType, *rmnops2_1.Version),
		ChainSelector:  evmChain.Selector,
		Args:           rmnops2_1.ConstructorArgs{CurseAdmins: []common.Address{}},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnops2_1.ContractType),
		Version:       rmnops2_1.Version,
		ChainSelector: evmChain.Selector,
		Address:       deployRMNOp.Output.Address,
	}))
	rmnAddress := common.HexToAddress(deployRMNOp.Output.Address)

	deployRMNProxyOp, err := cldf_ops.ExecuteOperation(bundle, rmnproxyops.Deploy, evmChain, contract.DeployInput[rmnproxyops.ConstructorArgs]{
		ChainSelector:  evmChain.Selector,
		TypeAndVersion: deployment.NewTypeAndVersion(rmnproxyops.ContractType, *semver.MustParse("1.0.0")),
		Args:           rmnproxyops.ConstructorArgs{RMN: rmnAddress},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		ChainSelector: evmChain.Selector,
		Address:       deployRMNProxyOp.Output.Address,
	}))

	return rmnAddress, common.HexToAddress(deployRMNProxyOp.Output.Address)
}

// rmnOwnershipTransferInput builds the transfer-ownership input that hands the RMN 2.1.0
// deployment on each of the given chains over to that chain's timelock.
func rmnOwnershipTransferInput(selectors []uint64, timelockAddrs map[uint64]string) []deploy.TransferOwnershipPerChainInput {
	inputs := make([]deploy.TransferOwnershipPerChainInput, 0, len(selectors))
	for _, sel := range selectors {
		inputs = append(inputs, deploy.TransferOwnershipPerChainInput{
			ChainSelector: sel,
			ContractRef: []datastore.AddressRef{
				{
					Type:    datastore.ContractType(rmnops2_1.ContractType),
					Version: rmnops2_1.Version,
				},
			},
			ProposedOwner: timelockAddrs[sel],
		})
	}
	return inputs
}

func TestFastCurse(t *testing.T) {
	chain1 := chainsel.TEST_90000001.Selector
	chain2 := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain1, chain2}),
	)
	require.NoError(t, err)
	bundle := env.OperationsBundle
	ds := datastore.NewMemoryDataStore()
	rmnAddresses := make(map[uint64]common.Address)
	// deploy RMN 2.1 (behind its RMNProxy) and a router on both chains
	for _, sel := range []uint64{chain1, chain2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		rmnAddr, rmnProxy := deployRMN2_1WithProxy(t, bundle, evmChain, ds)
		rmnAddresses[sel] = rmnAddr
		// mock wrapped native
		wNative := utils.RandomAddress()

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
	cs := deploy.DeployMCMS(deploy.GetRegistry(), nil)
	evmChain1 := env.BlockChains.EVMChains()[chain1]
	evmChain2 := env.BlockChains.EVMChains()[chain2]
	output, err := cs.Apply(*env, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS:           testhelpers.MCMSInputForQualifier(deploymentutils.CLLQualifier),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			chain1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(1),
				Qualifier:        deploymentutils.StringPtr(deploymentutils.CLLQualifier),
			},
			chain2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(1),
				Qualifier:        deploymentutils.StringPtr(deploymentutils.CLLQualifier),
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
	// transfer ownership of the RMN deployments to respective MCMS
	transferOwnershipInput := deploy.TransferOwnershipInput{
		ChainInputs:    rmnOwnershipTransferInput([]uint64{chain1, chain2}, timelockAddrs),
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Transfer ownership to timelock for fast curse test",
		},
	}

	// register chain adapter
	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(deploy.GetTransferOwnershipRegistry(), changesets.GetRegistry())
	output, err = transferOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
	t.Logf("Transferred ownership of RMN to respective MCMS")
	// now generate a curse proposal
	curseCfg := fastcurse.RMNCurseConfig{
		CurseActions: []fastcurse.CurseActionInput{
			{
				IsGlobalCurse:        false,
				ChainSelector:        chain1,
				SubjectChainSelector: chain2,
				Version:              rmnops2_1.Version,
			},
			{
				IsGlobalCurse:        false,
				ChainSelector:        chain2,
				SubjectChainSelector: chain1,
				Version:              rmnops2_1.Version,
			},
		},
		Force: false,
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Curse proposal for fast curse test",
		},
	}

	curseChangeset := fastcurse.CurseChangeset(fastcurse.GetCurseRegistry(), changesets.GetRegistry())
	output, err = curseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the subjects were actually cursed
	adv2_1_0 := adaptersv2_1_0.NewCurseAdapter()

	rmnC1, err := rmnops2_1.NewRMNContract(rmnAddresses[chain1], evmChain1.Client)
	require.NoError(t, err)
	isCursed, err := rmnC1.IsCursed(nil, adv2_1_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain2 should be cursed on rmn in chain1")

	rmnC2, err := rmnops2_1.NewRMNContract(rmnAddresses[chain2], evmChain2.Client)
	require.NoError(t, err)
	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain1 should be cursed on rmn in chain2")
	t.Logf("Subjects successfully cursed %x on chain1 %d and %x on chain2 %d", adv2_1_0.SelectorToSubject(chain2), chain1, adv2_1_0.SelectorToSubject(chain1), chain2)

	// Now uncurse the subjects
	// reset the operation bundle to clear any cached values
	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	uncurseChangeset := fastcurse.UncurseChangeset(fastcurse.GetCurseRegistry(), changesets.GetRegistry())
	output, err = uncurseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the subjects were actually uncursed
	isCursed, err = rmnC1.IsCursed(nil, adv2_1_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should be uncursed on rmn in chain1")

	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain1 should be uncursed on rmn in chain2")
	t.Logf("Subjects successfully uncursed %x on chain1 %d and %x on chain2 %d", adv2_1_0.SelectorToSubject(chain2), chain1, adv2_1_0.SelectorToSubject(chain1), chain2)

	// Uncursing again should fail because nothing is cursed on chain.
	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	_, err = uncurseChangeset.Apply(*env, curseCfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "uncurse skipped all actions")
}

func TestFastCurseFiredrill(t *testing.T) {
	chain1 := chainsel.TEST_90000007.Selector
	chain2 := chainsel.TEST_90000008.Selector
	chain3 := chainsel.TEST_90000009.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain1, chain2, chain3}),
	)
	require.NoError(t, err)
	bundle := env.OperationsBundle
	ds := datastore.NewMemoryDataStore()
	rmnAddresses := make(map[uint64]common.Address)
	// deploy RMN 2.1 (behind its RMNProxy) and a router on chain1 and chain2
	for _, sel := range []uint64{chain1, chain2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		rmnAddr, rmnProxy := deployRMN2_1WithProxy(t, bundle, evmChain, ds)
		rmnAddresses[sel] = rmnAddr
		// mock wrapped native
		wNative := utils.RandomAddress()

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

	// deploy RMN 2.1.0 on chain3, along with an RMNProxy pointing at it (required for the
	// v2.1.0 fastcurse adapter, which resolves the active RMN via the proxy) and a router
	// (required by CurseAdapter.Initialize, same as for chain1/chain2 above).
	evmChain3Setup := env.BlockChains.EVMChains()[chain3]
	rmnAddress3, rmnProxy3 := deployRMN2_1WithProxy(t, bundle, evmChain3Setup, ds)
	rmnAddresses[chain3] = rmnAddress3

	deployRouter3Op, err := cldf_ops.ExecuteOperation(bundle, routerops1_2.Deploy, evmChain3Setup, contract.DeployInput[routerops1_2.ConstructorArgs]{
		ChainSelector:  evmChain3Setup.Selector,
		TypeAndVersion: deployment.NewTypeAndVersion(routerops1_2.ContractType, *semver.MustParse("1.2.0")),
		Args: routerops1_2.ConstructorArgs{
			WrappedNative: utils.RandomAddress(),
			RMNProxy:      rmnProxy3,
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(routerops1_2.ContractType),
		Version:       semver.MustParse("1.2.0"),
		ChainSelector: chain3,
		Address:       deployRouter3Op.Output.Address,
	}))

	// deploy mcms
	cs := deploy.DeployMCMS(deploy.GetRegistry(), nil)
	evmChain2 := env.BlockChains.EVMChains()[chain2]
	evmChain3 := env.BlockChains.EVMChains()[chain3]
	output, err := cs.Apply(*env, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS:           testhelpers.MCMSInputForQualifier(deploymentutils.CLLQualifier),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			chain1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(1),
				Qualifier:        deploymentutils.StringPtr(deploymentutils.CLLQualifier),
			},
			chain2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(1),
				Qualifier:        deploymentutils.StringPtr(deploymentutils.CLLQualifier),
			},
			chain3: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(1),
				Qualifier:        deploymentutils.StringPtr(deploymentutils.CLLQualifier),
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
	// transfer ownership of the RMN deployments to respective MCMS
	transferOwnershipInput := deploy.TransferOwnershipInput{
		ChainInputs:    rmnOwnershipTransferInput([]uint64{chain1, chain2, chain3}, timelockAddrs),
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Transfer ownership to timelock for fast curse firedrill test",
		},
	}

	// register chain adapter
	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(deploy.GetTransferOwnershipRegistry(), changesets.GetRegistry())
	output, err = transferOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
	t.Logf("Transferred ownership of RMN to respective MCMS")

	// curse the hardcoded, inert firedrill subject on chain2 - it doesn't correspond to
	// any real chain selector and isn't the global curse subject, so it has no effect on
	// real lane curses.
	firedrillSubject := fastcurse.FiredrillSubject()
	curseCfg := fastcurse.RMNCurseConfig{
		CurseActions: []fastcurse.CurseActionInput{
			{
				ChainSelector: chain2,
				Subject:       &firedrillSubject,
				Version:       rmnops2_1.Version,
			},
			{
				ChainSelector: chain3,
				Subject:       &firedrillSubject,
				Version:       rmnops2_1.Version,
			},
		},
		Force: false,
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Firedrill curse proposal for fast curse test",
		},
	}

	curseChangeset := fastcurse.CurseChangeset(fastcurse.GetCurseRegistry(), changesets.GetRegistry())
	output, err = curseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the firedrill subject was actually cursed, and nothing else was affected
	adv2_1_0 := adaptersv2_1_0.NewCurseAdapter()
	rmnC2, err := rmnops2_1.NewRMNContract(rmnAddresses[chain2], evmChain2.Client)
	require.NoError(t, err)
	isCursed, err := rmnC2.IsCursed(nil, firedrillSubject)
	require.NoError(t, err)
	require.True(t, isCursed, "firedrill subject should be cursed on rmn in chain2")

	isCursed, err = rmnC2.IsCursed(nil, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.False(t, isCursed, "global curse subject should not be cursed on rmn in chain2")

	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "real subject for chain1 should not be cursed on rmn in chain2")
	t.Logf("Firedrill subject %x successfully cursed on chain2 %d", firedrillSubject, chain2)

	// check that the firedrill subject was also cursed on the RMN 2.1.0 deployment on chain3,
	// and nothing else was affected there either
	rmn2_1C, err := rmnops2_1.NewRMNContract(rmnAddresses[chain3], evmChain3.Client)
	require.NoError(t, err)
	isCursed, err = rmn2_1C.IsCursed(nil, firedrillSubject)
	require.NoError(t, err)
	require.True(t, isCursed, "firedrill subject should be cursed on rmn 2.1.0 in chain3")

	isCursed, err = rmn2_1C.IsCursed(nil, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.False(t, isCursed, "global curse subject should not be cursed on rmn 2.1.0 in chain3")

	isCursed, err = rmn2_1C.IsCursed(nil, adv2_1_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "real subject for chain1 should not be cursed on rmn 2.1.0 in chain3")
	t.Logf("Firedrill subject %x successfully cursed on chain3 %d", firedrillSubject, chain3)

	// Now uncurse the firedrill subject
	// reset the operation bundle to clear any cached values
	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	uncurseChangeset := fastcurse.UncurseChangeset(fastcurse.GetCurseRegistry(), changesets.GetRegistry())
	output, err = uncurseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that the firedrill subject was actually uncursed
	isCursed, err = rmnC2.IsCursed(nil, firedrillSubject)
	require.NoError(t, err)
	require.False(t, isCursed, "firedrill subject should be uncursed on rmn in chain2")
	t.Logf("Firedrill subject %x successfully uncursed on chain2 %d", firedrillSubject, chain2)

	isCursed, err = rmn2_1C.IsCursed(nil, firedrillSubject)
	require.NoError(t, err)
	require.False(t, isCursed, "firedrill subject should be uncursed on rmn 2.1.0 in chain3")
	t.Logf("Firedrill subject %x successfully uncursed on chain3 %d", firedrillSubject, chain3)

	// Uncursing again should fail because nothing is cursed on chain.
	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	_, err = uncurseChangeset.Apply(*env, curseCfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "uncurse skipped all actions")
}

// TestFastCurseSubjectOverridesLaneAndGlobal shows that the explicit Subject field is a
// general-purpose replacement for SubjectChainSelector (lane curses) and IsGlobalCurse
// (the well-known global subject), not just a firedrill-only mechanism. It curses a real
// lane subject and a global curse using the "normal" fields, then uncurses both purely by
// passing their literal Subject bytes - proving the same on-chain state can be reached and
// reversed through Subject alone.
func TestFastCurseSubjectOverridesLaneAndGlobal(t *testing.T) {
	chain1 := chainsel.TEST_90000010.Selector
	chain2 := chainsel.TEST_90000011.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain1, chain2}),
	)
	require.NoError(t, err)
	bundle := env.OperationsBundle
	ds := datastore.NewMemoryDataStore()
	rmnAddresses := make(map[uint64]common.Address)
	// deploy RMN 2.1 (behind its RMNProxy) and a router on both chains. No ramp updates are
	// configured, so neither chain is "connected" to the other from the router's point of
	// view - that's the point: the Subject field bypasses IsChainConnectedToTargetChain
	// entirely, so the lane curse below still succeeds even without any ramps set up.
	for _, sel := range []uint64{chain1, chain2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		rmnAddr, rmnProxy := deployRMN2_1WithProxy(t, bundle, evmChain, ds)
		rmnAddresses[sel] = rmnAddr
		deployRouterOp, err := cldf_ops.ExecuteOperation(bundle, routerops1_2.Deploy, evmChain, contract.DeployInput[routerops1_2.ConstructorArgs]{
			ChainSelector:  evmChain.Selector,
			TypeAndVersion: deployment.NewTypeAndVersion(routerops1_2.ContractType, *semver.MustParse("1.2.0")),
			Args: routerops1_2.ConstructorArgs{
				WrappedNative: utils.RandomAddress(),
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
	}

	// deploy mcms
	cs := deploy.DeployMCMS(deploy.GetRegistry(), nil)
	evmChain1 := env.BlockChains.EVMChains()[chain1]
	evmChain2 := env.BlockChains.EVMChains()[chain2]
	output, err := cs.Apply(*env, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS:           testhelpers.MCMSInputForQualifier(deploymentutils.CLLQualifier),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			chain1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(1),
				Qualifier:        deploymentutils.StringPtr(deploymentutils.CLLQualifier),
			},
			chain2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(1),
				Qualifier:        deploymentutils.StringPtr(deploymentutils.CLLQualifier),
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
	// transfer ownership of the RMN deployments to respective MCMS
	transferOwnershipInput := deploy.TransferOwnershipInput{
		ChainInputs:    rmnOwnershipTransferInput([]uint64{chain1, chain2}, timelockAddrs),
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Transfer ownership to timelock for fast curse subject-override test",
		},
	}

	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(deploy.GetTransferOwnershipRegistry(), changesets.GetRegistry())
	output, err = transferOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
	t.Logf("Transferred ownership of RMN to respective MCMS")

	// Curse a real lane subject on chain1 via the literal Subject field instead of
	// SubjectChainSelector, and put a global curse on chain2 via the normal IsGlobalCurse
	// flag. Note the lane action doesn't need a reverse-direction pair the way
	// SubjectChainSelector-based lane curses do: Subject actions are exempt from the
	// bidirectional-lane validation entirely.
	globalSubject := fastcurse.GlobalCurseSubject()
	laneSubject := fastcurse.GenericSelectorToSubject(chain2)
	curseCfg := fastcurse.RMNCurseConfig{
		CurseActions: []fastcurse.CurseActionInput{
			{
				ChainSelector: chain1,
				Subject:       &laneSubject,
				Version:       rmnops2_1.Version,
			},
			{
				ChainSelector: chain2,
				Subject:       &globalSubject,
				Version:       rmnops2_1.Version,
			},
		},
		Force: false,
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Subject-override curse proposal for fast curse test",
		},
	}

	curseChangeset := fastcurse.CurseChangeset(fastcurse.GetCurseRegistry(), changesets.GetRegistry())
	output, err = curseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	rmnC, err := rmnops2_1.NewRMNContract(rmnAddresses[chain1], evmChain1.Client)
	require.NoError(t, err)
	isCursed, err := rmnC.IsCursed(nil, laneSubject)
	require.NoError(t, err)
	require.True(t, isCursed, "lane subject cursed via Subject field should be cursed on rmn in chain1")

	rmnC2, err := rmnops2_1.NewRMNContract(rmnAddresses[chain2], evmChain2.Client)
	require.NoError(t, err)
	isCursed, err = rmnC2.IsCursed(nil, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.True(t, isCursed, "global curse should be cursed on rmn in chain2")
	t.Logf("Lane subject %x cursed via Subject field on chain1 %d, global curse applied on chain2 %d", laneSubject, chain1, chain2)

	// Now uncurse both purely through the Subject field: the lane subject bytes for
	// chain1, and the well-known GlobalCurseSubject() bytes (instead of IsGlobalCurse) for
	// chain2 - showing Subject can address and remove the global curse too.
	uncurseCfg := fastcurse.RMNCurseConfig{
		CurseActions: []fastcurse.CurseActionInput{
			{
				ChainSelector: chain1,
				Subject:       &laneSubject,
				Version:       rmnops2_1.Version,
			},
			{
				ChainSelector: chain2,
				Subject:       &globalSubject,
				Version:       rmnops2_1.Version,
			},
		},
		Force: false,
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Subject-override uncurse proposal for fast curse test",
		},
	}

	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	uncurseChangeset := fastcurse.UncurseChangeset(fastcurse.GetCurseRegistry(), changesets.GetRegistry())
	output, err = uncurseChangeset.Apply(*env, uncurseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	isCursed, err = rmnC.IsCursed(nil, laneSubject)
	require.NoError(t, err)
	require.False(t, isCursed, "lane subject should be uncursed on rmn in chain1")

	isCursed, err = rmnC2.IsCursed(nil, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.False(t, isCursed, "global curse should be uncursed on rmn in chain2, having been removed via the Subject field")
	t.Logf("Lane subject %x and global curse on chain2 %d both uncursed via the Subject field", laneSubject, chain2)

	// Uncursing again should fail because nothing is cursed on chain.
	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	_, err = uncurseChangeset.Apply(*env, uncurseCfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "uncurse skipped all actions")
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
	ds := datastore.NewMemoryDataStore()
	rmnAddresses := make(map[uint64]common.Address)
	// deploy RMN 2.1 (behind its RMNProxy) and a router in all chains
	for _, sel := range []uint64{chain1, chain2, chain3} {
		evmChain := env.BlockChains.EVMChains()[sel]
		rmnAddr, rmnProxy := deployRMN2_1WithProxy(t, bundle, evmChain, ds)
		rmnAddresses[sel] = rmnAddr
		// mock wrapped native
		wNative := utils.RandomAddress()
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
	cs := deploy.DeployMCMS(dReg, nil)
	mcmsChainInput := make(map[uint64]deploy.MCMSDeploymentConfigPerChain)
	for _, sel := range []uint64{chain1, chain2, chain3} {
		mcmsChainInput[sel] = deploy.MCMSDeploymentConfigPerChain{
			Canceller:        testhelpers.SingleGroupMCMS(),
			Bypasser:         testhelpers.SingleGroupMCMS(),
			Proposer:         testhelpers.SingleGroupMCMS(),
			TimelockMinDelay: big.NewInt(1),
			Qualifier:        deploymentutils.StringPtr(deploymentutils.CLLQualifier),
		}
	}
	output, err := cs.Apply(*env, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS:           testhelpers.MCMSInputForQualifier(deploymentutils.CLLQualifier),
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
	// transfer ownership of the RMN deployments to respective MCMS
	transferOwnershipInput := deploy.TransferOwnershipInput{
		ChainInputs:    rmnOwnershipTransferInput([]uint64{chain1, chain2, chain3}, timelockAddrs),
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
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
	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(cr, mcmsRegistry)
	output, err = transferOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
	t.Logf("Transferred ownership of RMN to respective MCMS")
	// now generate a global curse on chain 3
	curseCfg := fastcurse.GlobalCurseOnNetworkInput{
		ChainSelectors: map[uint64]*semver.Version{
			chain3: rmnops2_1.Version,
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Curse proposal for fast curse test",
		},
	}
	// The RMN 2.1.0 curse adapter is already registered against rmnops2_1.Version by the
	// v2_0_0 adapters package init(), and RegisterNewCurse is first-write-wins, so there is
	// nothing to register here - just read the registered adapter back below.
	curseReg := fastcurse.GetCurseRegistry()
	adv2_1_0 := adaptersv2_1_0.NewCurseAdapter()
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
	rmnC, err := rmnops2_1.NewRMNContract(rmnAddresses[chain1], evmChain1.Client)
	require.NoError(t, err)
	// chain2 subject should not be cursed
	isCursed, err := rmnC.IsCursed(nil, adv2_1_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should be cursed on rmn in chain1")
	// chain3 subject should be cursed
	isCursed, err = rmnC.IsCursed(nil, adv2_1_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain3 should be cursed on rmn in chain1")

	rmnC2, err := rmnops2_1.NewRMNContract(rmnAddresses[chain2], evmChain2.Client)
	require.NoError(t, err)
	// chain1 subject should not be cursed on chain2
	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain1 should be cursed on rmn in chain2")
	// chain3 subject should be cursed on chain2
	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain3 should be cursed on rmn in chain2")

	// chain3 should have a global curse on itself
	rmnC3, err := rmnops2_1.NewRMNContract(rmnAddresses[chain3], evmChain3.Client)
	require.NoError(t, err)
	isCursed, err = rmnC3.IsCursed(nil, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.True(t, isCursed, "subject on chain3 should be cursed on rmn in chain3")

	curseAdapter, ok := curseReg.GetCurseAdapter(chainsel.FamilyEVM, rmnops2_1.Version)
	require.True(t, ok)
	// Check that adapter does not return true for specific subjects when global curse is present
	isCursed, err = curseAdapter.IsSubjectCursedOnChain(*env, chain3, adv2_1_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should not be cursed on rmn in chain3")

	isCursed, err = curseAdapter.IsSubjectCursedOnChain(*env, chain3, adv2_1_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain1 should not be cursed on rmn in chain3")

	isCursed, err = curseAdapter.IsSubjectCursedOnChain(*env, chain3, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.True(t, isCursed, "global curse on chain3 should be cursed on rmnremote in chain3")

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
	isCursed, err = rmnC.IsCursed(nil, adv2_1_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should be uncursed on rmn in chain1")

	isCursed, err = rmnC.IsCursed(nil, adv2_1_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain3 should be uncursed on rmn in chain1")

	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chain1))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain1 should be uncursed on rmn in chain2")
	isCursed, err = rmnC2.IsCursed(nil, adv2_1_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain3 should be uncursed on rmn in chain2")

	isCursed, err = rmnC3.IsCursed(nil, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.False(t, isCursed, "global curse on chain3 should be uncursed on rmn in chain3")

	// Now do a global curse on chain1
	curseCfg = fastcurse.GlobalCurseOnNetworkInput{
		ChainSelectors: map[uint64]*semver.Version{
			chain1: rmnops2_1.Version,
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "Curse proposal for fast curse test",
		},
	}
	curseChangeset = fastcurse.GloballyCurseChainChangeset(curseReg, mcmsRegistry)
	output, err = curseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	curseAdapter, ok = curseReg.GetCurseAdapter(chainsel.FamilyEVM, rmnops2_1.Version)
	require.True(t, ok)

	// Check that adapter does not return true for specific subjects when global curse is present
	isCursed, err = curseAdapter.IsSubjectCursedOnChain(*env, chain1, adv2_1_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should not be cursed on rmn in chain1")

	isCursed, err = curseAdapter.IsSubjectCursedOnChain(*env, chain1, adv2_1_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain3 should not be cursed on rmn in chain1")

	isCursed, err = curseAdapter.IsSubjectCursedOnChain(*env, chain1, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.True(t, isCursed, "global curse on chain1 should be cursed on rmn in chain1")

	// Now uncurse the global curse
	uncurseChangeset = fastcurse.GloballyUncurseChainChangeset(curseReg, mcmsRegistry)
	output, err = uncurseChangeset.Apply(*env, curseCfg)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	isCursed, err = curseAdapter.IsSubjectCursedOnChain(*env, chain1, fastcurse.GlobalCurseSubject())
	require.NoError(t, err)
	require.False(t, isCursed, "global curse on chain1 should be uncursed on rmn in chain1")

	isCursed, err = curseAdapter.IsSubjectCursedOnChain(*env, chain1, adv2_1_0.SelectorToSubject(chain2))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain2 should not be cursed on rmn in chain1")

	isCursed, err = curseAdapter.IsSubjectCursedOnChain(*env, chain1, adv2_1_0.SelectorToSubject(chain3))
	require.NoError(t, err)
	require.False(t, isCursed, "subject on chain3 should not be cursed on rmn in chain1")
}

// TestCurseAdapterConcurrentAccess exercises a single shared CurseAdapter instance
// under concurrent Initialize and Is* calls from many goroutines.
// Run with -race to detect unsynchronized map access.
func TestCurseAdapterConcurrentAccess(t *testing.T) {
	selectors := []uint64{
		chainsel.TEST_90000015.Selector,
		chainsel.TEST_90000016.Selector,
		chainsel.TEST_90000017.Selector,
	}

	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, selectors),
	)
	require.NoError(t, err)

	bundle := env.OperationsBundle
	ds := datastore.NewMemoryDataStore()

	for _, sel := range selectors {
		evmChain := env.BlockChains.EVMChains()[sel]
		_, rmnProxy := deployRMN2_1WithProxy(t, bundle, evmChain, ds)

		deployRouterOp, err := cldf_ops.ExecuteOperation(bundle, routerops1_2.Deploy, evmChain, contract.DeployInput[routerops1_2.ConstructorArgs]{
			ChainSelector:  evmChain.Selector,
			TypeAndVersion: deployment.NewTypeAndVersion(routerops1_2.ContractType, *semver.MustParse("1.2.0")),
			Args: routerops1_2.ConstructorArgs{
				WrappedNative: utils.RandomAddress(),
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
	}
	env.DataStore = ds.Seal()

	adapter := adaptersv2_1_0.NewCurseAdapter()

	// Phase 1: concurrent Initialize for all selectors — the original race site.
	var wg sync.WaitGroup
	for _, sel := range selectors {
		wg.Add(1)
		go func(s uint64) {
			defer wg.Done()
			_ = adapter.Initialize(*env, s)
		}(sel)
	}
	wg.Wait()

	for _, sel := range selectors {
		enabled, err := adapter.IsCurseEnabledForChain(*env, sel)
		require.NoError(t, err)
		require.True(t, enabled, "adapter should be initialized for selector %d", sel)
	}

	// Phase 2: hammer Is* methods concurrently — exercises RWMutex-protected cache
	// reads alongside RPC calls through the shared ethclient.
	const goroutinesPerSelector = 20
	for _, sel := range selectors {
		for i := 0; i < goroutinesPerSelector; i++ {
			wg.Add(2)
			go func(s uint64) {
				defer wg.Done()
				_, _ = adapter.IsCurseEnabledForChain(*env, s)
			}(sel)
			go func(s uint64) {
				defer wg.Done()
				_, _ = adapter.IsSubjectCursedOnChain(*env, s, fastcurse.GlobalCurseSubject())
			}(sel)
		}
	}
	wg.Wait()
}
