package changesets_test

import (
	"github.com/Masterminds/semver/v3"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	rmn_proxy_bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcmslib "github.com/smartcontractkit/mcms"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	rmnops1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	rmnremoteops1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestActivateRMN_Apply(t *testing.T) {
	testActivateRMNApplyWithLegacyContract(t, deployLegacyRMNV2AndCurse)
}

func TestActivateRMN_Apply_WithRMN15(t *testing.T) {
	testActivateRMNApplyWithLegacyContract(t, deployLegacyRMN15AndCurse)
}

func TestActivateRMN_Apply_WithRMN16(t *testing.T) {
	testActivateRMNApplyWithLegacyContract(t, deployLegacyRMN16AndCurse)
}

type legacyRMNDeployerForTest func(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	deployer common.Address,
	curseSubject [16]byte,
) common.Address

func testActivateRMNApplyWithLegacyContract(t *testing.T, deployLegacy legacyRMNDeployerForTest) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From
	b := e.OperationsBundle
	curseSubject := fastcurse.GenericSelectorToSubject(chainSelector)
	legacyARM := deployLegacy(t, b, chain, chainSelector, deployer, curseSubject)

	proxyRef, err := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		ChainSelector:  chainSelector,
		Args:           rmn_proxy.ConstructorArgs{RMN: legacyARM},
	}, nil)
	require.NoError(t, err)

	ultraFastTimelockAddr, ultraFastAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.UltraFastCurseMCMSQualifier)
	rmnTimelockAddr, rmnMCMSAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.RMNTimelockQualifier)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(proxyRef))
	for _, ref := range ultraFastAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	for _, ref := range rmnMCMSAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.ActivateRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
		MCMS: mcms.Input{
			Qualifier:      common_utils.RMNTimelockQualifier,
			TimelockAction: mcms_types.TimelockActionSchedule,
			ValidUntil:     3759765795,
		},
		Cfg: changesets.ActivateRMNCfg{
			ChainSels: []uint64{chainSelector},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, out.MCMSTimelockProposals, "accept ownership on RMNMCMS timelock should be proposed via MCMS")
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	rmnAddrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, rmnAddrs, 1)
	require.Equal(t, datastore.ContractType(rmnops.ContractType), rmnAddrs[0].Type)

	newRMN := common.HexToAddress(rmnAddrs[0].Address)

	rmnC, err := rmnops.NewRMNContract(newRMN, chain.Client)
	require.NoError(t, err)
	cursedSubjects, err := rmnC.GetCursedSubjects(nil)
	require.NoError(t, err)
	require.Contains(t, cursedSubjects, curseSubject, "new RMN must preserve active curses from the old RMN")
	callers, err := rmnC.GetAllAuthorizedCallers(nil)
	require.NoError(t, err)
	require.Contains(t, callers, ultraFastTimelockAddr, "Ultra Fast Curse timelock must be a curse admin")

	owner, err := rmnC.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, rmnTimelockAddr, owner, "RMN should be owned by RMNMCMS timelock after proposal execution")

	proxyAddr := common.HexToAddress(proxyRef.Address)
	proxyC, err := rmn_proxy_bind.NewRMNProxy(proxyAddr, chain.Client)
	require.NoError(t, err)
	proxyOwner, err := proxyC.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, deployer, proxyOwner, "ARMProxy should be deployer-owned in this test")

	armReport, err := operations.ExecuteOperation(b, rmn_proxy.GetRMN, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       proxyAddr,
	})
	require.NoError(t, err)
	require.Equal(t, newRMN, armReport.Output, "ARMProxy must point at the new RMN implementation")
	require.NotEqual(t, legacyARM, armReport.Output)
}

func deployLegacyRMNV2AndCurse(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	deployer common.Address,
	curseSubject [16]byte,
) common.Address {
	t.Helper()
	legacyRMNRef, err := contract.MaybeDeployContract(b, rmnops.Deploy, chain, contract.DeployInput[rmnops.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmnops.ContractType, *rmnops.Version),
		ChainSelector:  chainSelector,
		Args:           rmnops.ConstructorArgs{CurseAdmins: []common.Address{deployer}},
	}, nil)
	require.NoError(t, err)
	legacyRMNAddr := common.HexToAddress(legacyRMNRef.Address)
	_, err = operations.ExecuteOperation(b, rmnops.Curse0, chain, contract.FunctionInput[[][16]byte]{
		ChainSelector: chainSelector,
		Address:       legacyRMNAddr,
		Args:          [][16]byte{curseSubject},
	})
	require.NoError(t, err)
	return legacyRMNAddr
}

func deployLegacyRMN15AndCurse(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	_ common.Address,
	curseSubject [16]byte,
) common.Address {
	t.Helper()
	legacyRMNRef, err := contract.MaybeDeployContract(b, rmnops1_5.Deploy, chain, contract.DeployInput[rmnops1_5.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmnops1_5.ContractType, *semver.MustParse("1.5.0")),
		ChainSelector:  chainSelector,
		Args: rmnops1_5.ConstructorArgs{
			RMNConfig: rmn_contract.RMNConfig{
				BlessWeightThreshold: 1,
				CurseWeightThreshold: 1,
				Voters: []rmn_contract.RMNVoter{{
					BlessWeight:   1,
					CurseWeight:   1,
					BlessVoteAddr: common.HexToAddress("0x1111111111111111111111111111111111111111"),
					CurseVoteAddr: common.HexToAddress("0x2222222222222222222222222222222222222222"),
				}},
			},
		},
	}, nil)
	require.NoError(t, err)
	legacyRMNAddr := common.HexToAddress(legacyRMNRef.Address)

	_, err = operations.ExecuteOperation(b, rmnops1_5.Curse, chain, contract.FunctionInput[rmnops1_5.CurseArgs]{
		ChainSelector: chainSelector,
		Address:       legacyRMNAddr,
		Args: rmnops1_5.CurseArgs{
			CurseID: [16]byte{1},
			Subject: []fastcurse.Subject{curseSubject},
		},
	})
	require.NoError(t, err)
	return legacyRMNAddr
}

func deployLegacyRMN16AndCurse(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	_ common.Address,
	curseSubject [16]byte,
) common.Address {
	t.Helper()
	legacyRMNRef, err := contract.MaybeDeployContract(b, rmnremoteops1_6.Deploy, chain, contract.DeployInput[rmnremoteops1_6.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmnremoteops1_6.ContractType, *rmnremoteops1_6.Version),
		ChainSelector:  chainSelector,
		Args: rmnremoteops1_6.ConstructorArgs{
			LocalChainSelector: chainSelector,
			LegacyRMN:          common.HexToAddress("0x3333333333333333333333333333333333333333"),
		},
	}, nil)
	require.NoError(t, err)
	legacyRMNAddr := common.HexToAddress(legacyRMNRef.Address)

	_, err = operations.ExecuteOperation(b, rmnremoteops1_6.Curse0, chain, contract.FunctionInput[[][16]byte]{
		ChainSelector: chainSelector,
		Address:       legacyRMNAddr,
		Args:          [][16]byte{curseSubject},
	})
	require.NoError(t, err)
	return legacyRMNAddr
}

func TestActivateRMN_WithAdditionalCurseAdmins(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From
	b := e.OperationsBundle
	additionalCurser := common.HexToAddress("0x1111111111111111111111111111111111111111")

	proxyRef, err := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		ChainSelector:  chainSelector,
		Args:           rmn_proxy.ConstructorArgs{RMN: deployer},
	}, nil)
	require.NoError(t, err)

	ultraFastTimelockAddr, ultraFastAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.UltraFastCurseMCMSQualifier)
	_, rmnMCMSAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.RMNTimelockQualifier)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(proxyRef))
	for _, ref := range ultraFastAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	for _, ref := range rmnMCMSAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.ActivateRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
		MCMS: mcms.Input{
			Qualifier:      common_utils.RMNTimelockQualifier,
			TimelockAction: mcms_types.TimelockActionSchedule,
			ValidUntil:     3759765795,
		},
		Cfg: changesets.ActivateRMNCfg{
			ChainSels: []uint64{chainSelector},
			CurseAdmins: map[uint64][]common.Address{
				chainSelector: {additionalCurser},
			},
		},
	})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	rmnAddrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, rmnAddrs, 1)

	rmnC, err := rmnops.NewRMNContract(common.HexToAddress(rmnAddrs[0].Address), chain.Client)
	require.NoError(t, err)
	callers, err := rmnC.GetAllAuthorizedCallers(nil)
	require.NoError(t, err)
	require.Contains(t, callers, ultraFastTimelockAddr, "Ultra Fast Curse timelock must be a curse admin")
	require.Contains(t, callers, additionalCurser, "configured curse admin must be authorized at deploy time")
}

func TestActivateRMN_ValidateCurseAdmins(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	mcmsRegistry := cs_core.GetRegistry()
	changeset := changesets.ActivateRMN(mcmsRegistry)

	t.Run("zero address", func(t *testing.T) {
		err := changeset.VerifyPreconditions(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
			Cfg: changesets.ActivateRMNCfg{
				ChainSels: []uint64{chainSelector},
				CurseAdmins: map[uint64][]common.Address{
					chainSelector: {common.Address{}},
				},
			},
		})
		require.ErrorContains(t, err, "curse admin address cannot be zero")
	})

	t.Run("duplicate address", func(t *testing.T) {
		addr := common.HexToAddress("0x1111111111111111111111111111111111111111")
		err := changeset.VerifyPreconditions(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
			Cfg: changesets.ActivateRMNCfg{
				ChainSels: []uint64{chainSelector},
				CurseAdmins: map[uint64][]common.Address{
					chainSelector: {addr, addr},
				},
			},
		})
		require.ErrorContains(t, err, "duplicate curse admin")
	})

	t.Run("chain not in ChainSels", func(t *testing.T) {
		err := changeset.VerifyPreconditions(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
			Cfg: changesets.ActivateRMNCfg{
				ChainSels: []uint64{chainSelector},
				CurseAdmins: map[uint64][]common.Address{
					99999: {common.HexToAddress("0x1111111111111111111111111111111111111111")},
				},
			},
		})
		require.ErrorContains(t, err, "not in ChainSels")
	})
}

func TestActivateRMN_AccumulatesBatchOpsAcrossChains(t *testing.T) {
	const (
		chainSelectorA = uint64(5009297550715157269)
		chainSelectorB = uint64(4949039107694359620)
	)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelectorA, chainSelectorB}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	for _, sel := range []uint64{chainSelectorA, chainSelectorB} {
		chain := e.BlockChains.EVMChains()[sel]
		deployer := chain.DeployerKey.From
		b := e.OperationsBundle

		proxyRef, deployErr := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
			ChainSelector:  sel,
			Args:           rmn_proxy.ConstructorArgs{RMN: deployer},
		}, nil)
		require.NoError(t, deployErr)
		require.NoError(t, ds.Addresses().Add(proxyRef))

		_, ultraFastAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.UltraFastCurseMCMSQualifier)
		for _, ref := range ultraFastAddrs {
			require.NoError(t, ds.Addresses().Add(ref))
		}

		_, rmnMCMSAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.RMNTimelockQualifier)
		for _, ref := range rmnMCMSAddrs {
			require.NoError(t, ds.Addresses().Add(ref))
		}
	}
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.ActivateRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
		MCMS: mcms.Input{
			Qualifier:      common_utils.RMNTimelockQualifier,
			TimelockAction: mcms_types.TimelockActionSchedule,
			ValidUntil:     3759765795,
		},
		Cfg: changesets.ActivateRMNCfg{
			ChainSels: []uint64{chainSelectorA, chainSelectorB},
		},
	})
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)

	proposal := out.MCMSTimelockProposals[0]
	require.Len(t, proposal.Operations, 2, "proposal should include acceptOwnership for each chain")
	chainSelectors := make([]uint64, 0, len(proposal.Operations))
	for _, op := range proposal.Operations {
		chainSelectors = append(chainSelectors, uint64(op.ChainSelector))
		require.Len(t, op.Transactions, 1, "each chain should have one acceptOwnership transaction")
	}
	require.ElementsMatch(t, []uint64{chainSelectorA, chainSelectorB}, chainSelectors)
}

func TestActivateRMN_MissingRMNMCMS(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From
	b := e.OperationsBundle

	proxyRef, err := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		ChainSelector:  chainSelector,
		Args:           rmn_proxy.ConstructorArgs{RMN: deployer},
	}, nil)
	require.NoError(t, err)

	_, ultraFastAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.UltraFastCurseMCMSQualifier)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(proxyRef))
	for _, ref := range ultraFastAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	_, err = changesets.ActivateRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
		MCMS: mcms.Input{},
		Cfg:  changesets.ActivateRMNCfg{ChainSels: []uint64{chainSelector}},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, common_utils.RMNTimelockQualifier)
}

func TestActivateRMN_CLLCCIPOwnedProxyEmitsDualProposals(t *testing.T) {
	const (
		chainSelector = uint64(5009297550715157269)
		validUntil    = uint32(3759765795)
	)
	expectedDelay := mcms_types.NewDuration(0)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From
	b := e.OperationsBundle

	legacyARM := common.HexToAddress("0x2222222222222222222222222222222222222222")
	proxyRef, err := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		ChainSelector:  chainSelector,
		Args:           rmn_proxy.ConstructorArgs{RMN: legacyARM},
	}, nil)
	require.NoError(t, err)

	_, ultraFastAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.UltraFastCurseMCMSQualifier)
	_, rmnMCMSAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.RMNTimelockQualifier)
	cllTimelockAddr, cllAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.CLLQualifier)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(proxyRef))
	for _, ref := range ultraFastAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	for _, ref := range rmnMCMSAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	for _, ref := range cllAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	transferProxyOwnershipToTimelockForTest(
		t, b, chain, common.HexToAddress(proxyRef.Address), cllTimelockAddr, e, common_utils.CLLQualifier, validUntil,
	)

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.ActivateRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
		MCMS: mcms.Input{
			OverridePreviousRoot: true,
			ValidUntil:           validUntil,
		},
		Cfg: changesets.ActivateRMNCfg{
			ChainSels: []uint64{chainSelector},
		},
	})
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 2)

	rmnProposal, cllProposal := splitActivateRMNProposalsForTest(t, *e, chainSelector, out.MCMSTimelockProposals)
	require.Equal(t, mcms_types.TimelockActionSchedule, rmnProposal.Action)
	require.Equal(t, expectedDelay, rmnProposal.Delay)
	require.Equal(t, mcms_types.TimelockActionSchedule, cllProposal.Action)
	require.Equal(t, expectedDelay, cllProposal.Delay)
	require.Len(t, rmnProposal.Operations, 1)
	require.Len(t, rmnProposal.Operations[0].Transactions, 1)
	require.Len(t, cllProposal.Operations, 1)
	require.Len(t, cllProposal.Operations[0].Transactions, 1)

	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	proxyC, err := rmn_proxy_bind.NewRMNProxy(common.HexToAddress(proxyRef.Address), chain.Client)
	require.NoError(t, err)
	require.Equal(t, cllTimelockAddr, common.HexToAddress(cllProposal.TimelockAddresses[mcms_types.ChainSelector(chainSelector)]))

	rmnAddrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, rmnAddrs, 1)

	armReport, err := operations.ExecuteOperation(b, rmn_proxy.GetRMN, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(proxyRef.Address),
	})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(rmnAddrs[0].Address), armReport.Output)
	require.NotEqual(t, legacyARM, armReport.Output)

	proxyOwner, err := proxyC.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, cllTimelockAddr, proxyOwner)
}

func TestActivateRMN_DualProposalsAccumulateBatchOpsAcrossChains(t *testing.T) {
	const (
		chainSelectorA = uint64(5009297550715157269)
		chainSelectorB = uint64(4949039107694359620)
		validUntil     = uint32(3759765795)
	)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelectorA, chainSelectorB}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	cllTimelocks := make(map[uint64]common.Address, 2)
	for _, sel := range []uint64{chainSelectorA, chainSelectorB} {
		chain := e.BlockChains.EVMChains()[sel]
		deployer := chain.DeployerKey.From
		b := e.OperationsBundle

		proxyRef, deployErr := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
			ChainSelector:  sel,
			Args:           rmn_proxy.ConstructorArgs{RMN: common.HexToAddress("0x3333333333333333333333333333333333333333")},
		}, nil)
		require.NoError(t, deployErr)
		require.NoError(t, ds.Addresses().Add(proxyRef))

		_, ultraFastAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.UltraFastCurseMCMSQualifier)
		for _, ref := range ultraFastAddrs {
			require.NoError(t, ds.Addresses().Add(ref))
		}

		_, rmnMCMSAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.RMNTimelockQualifier)
		for _, ref := range rmnMCMSAddrs {
			require.NoError(t, ds.Addresses().Add(ref))
		}

		cllTimelockAddr, cllAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.CLLQualifier)
		cllTimelocks[sel] = cllTimelockAddr
		for _, ref := range cllAddrs {
			require.NoError(t, ds.Addresses().Add(ref))
		}

		e.DataStore = ds.Seal()
		transferProxyOwnershipToTimelockForTest(
			t, b, chain, common.HexToAddress(proxyRef.Address), cllTimelockAddr, e, common_utils.CLLQualifier, validUntil,
		)
	}
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.ActivateRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
		MCMS: mcms.Input{
			OverridePreviousRoot: true,
			ValidUntil:           validUntil,
		},
		Cfg: changesets.ActivateRMNCfg{
			ChainSels: []uint64{chainSelectorA, chainSelectorB},
		},
	})
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 2)

	rmnProposal, cllProposal := splitActivateRMNProposalsForTest(t, *e, chainSelectorA, out.MCMSTimelockProposals)
	require.Len(t, rmnProposal.Operations, 2)
	require.Len(t, cllProposal.Operations, 2)

	rmnChainSelectors := make([]uint64, 0, len(rmnProposal.Operations))
	for _, op := range rmnProposal.Operations {
		rmnChainSelectors = append(rmnChainSelectors, uint64(op.ChainSelector))
		require.Len(t, op.Transactions, 1)
	}
	require.ElementsMatch(t, []uint64{chainSelectorA, chainSelectorB}, rmnChainSelectors)

	cllChainSelectors := make([]uint64, 0, len(cllProposal.Operations))
	for _, op := range cllProposal.Operations {
		cllChainSelectors = append(cllChainSelectors, uint64(op.ChainSelector))
		require.Len(t, op.Transactions, 1)
		require.Equal(t, cllTimelocks[uint64(op.ChainSelector)].Hex(), cllProposal.TimelockAddresses[op.ChainSelector])
	}
	require.ElementsMatch(t, []uint64{chainSelectorA, chainSelectorB}, cllChainSelectors)
}

func TestActivateRMN_MissingCLLCCIPWhenProxyNotDeployerOwned(t *testing.T) {
	const validUntil = uint32(3759765795)

	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From
	b := e.OperationsBundle

	proxyRef, err := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		ChainSelector:  chainSelector,
		Args:           rmn_proxy.ConstructorArgs{RMN: deployer},
	}, nil)
	require.NoError(t, err)

	_, ultraFastAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.UltraFastCurseMCMSQualifier)
	rmnTimelockAddr, rmnAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.RMNTimelockQualifier)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(proxyRef))
	for _, ref := range ultraFastAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	for _, ref := range rmnAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	transferProxyOwnershipToTimelockForTest(
		t, b, chain, common.HexToAddress(proxyRef.Address), rmnTimelockAddr, e, common_utils.RMNTimelockQualifier, validUntil,
	)

	mcmsRegistry := cs_core.GetRegistry()
	_, err = changesets.ActivateRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
		MCMS: mcms.Input{ValidUntil: validUntil},
		Cfg:  changesets.ActivateRMNCfg{ChainSels: []uint64{chainSelector}},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, common_utils.CLLQualifier)
}

func transferProxyOwnershipToTimelockForTest(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	proxyAddr common.Address,
	timelockAddr common.Address,
	e *deployment.Environment,
	mcmsQualifier string,
	validUntil uint32,
) {
	t.Helper()

	batchOps, err := mcms_seq.TransferAndAcceptOwnership(b, chain, []mcms_ops.OpTransferOwnershipInput{
		{
			ChainSelector:   chain.Selector,
			Address:         proxyAddr,
			ProposedOwner:   timelockAddr,
			ContractType:    rmn_proxy.ContractType,
			TimelockAddress: timelockAddr,
		},
	})
	require.NoError(t, err)
	if len(batchOps) == 0 {
		return
	}

	mcmsRegistry := cs_core.GetRegistry()
	out, err := cs_core.NewOutputBuilder(*e, mcmsRegistry).
		WithBatchOps(batchOps).
		Build(mcms.Input{
			Qualifier:      mcmsQualifier,
			TimelockAction: mcms_types.TimelockActionSchedule,
			ValidUntil:     validUntil,
		})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)
}

func splitActivateRMNProposalsForTest(
	t *testing.T,
	e deployment.Environment,
	chainSelector uint64,
	proposals []mcmslib.TimelockProposal,
) (rmnProposal, cllProposal mcmslib.TimelockProposal) {
	t.Helper()
	require.Len(t, proposals, 2)

	mcmsReader, ok := cs_core.GetRegistry().GetMCMSReader("evm")
	require.True(t, ok)

	rmnTimelockRef, err := mcmsReader.GetTimelockRef(e, chainSelector, mcms.Input{Qualifier: common_utils.RMNTimelockQualifier})
	require.NoError(t, err)
	cllTimelockRef, err := mcmsReader.GetTimelockRef(e, chainSelector, mcms.Input{Qualifier: common_utils.CLLQualifier})
	require.NoError(t, err)

	for _, proposal := range proposals {
		timelockAddr := proposal.TimelockAddresses[mcms_types.ChainSelector(chainSelector)]
		switch timelockAddr {
		case rmnTimelockRef.Address:
			rmnProposal = proposal
		case cllTimelockRef.Address:
			cllProposal = proposal
		default:
			t.Fatalf("unexpected timelock address %s in proposal", timelockAddr)
		}
	}
	return rmnProposal, cllProposal
}

func deployMCMSInstanceForTest(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	deployer common.Address,
	qualifier string,
) (timelockAddr common.Address, addresses []datastore.AddressRef) {
	t.Helper()
	mcmsCfg := testhelpers.SingleGroupMCMS()
	qualifierPtr := &qualifier

	proposerReport, err := operations.ExecuteSequence(b, mcms_seq.SeqDeployMCMWithConfig, chain, mcms_seq.SeqMCMSDeploymentCfg{
		ChainSelector: chain.Selector,
		ContractType:  common_utils.ProposerManyChainMultisig,
		MCMConfig:     &mcmsCfg,
		Qualifier:     qualifierPtr,
	})
	require.NoError(t, err)
	require.NotEmpty(t, proposerReport.Output.Addresses)
	addresses = append(addresses, proposerReport.Output.Addresses...)

	bypasserReport, err := operations.ExecuteSequence(b, mcms_seq.SeqDeployMCMWithConfig, chain, mcms_seq.SeqMCMSDeploymentCfg{
		ChainSelector: chain.Selector,
		ContractType:  common_utils.BypasserManyChainMultisig,
		MCMConfig:     &mcmsCfg,
		Qualifier:     qualifierPtr,
	})
	require.NoError(t, err)
	addresses = append(addresses, bypasserReport.Output.Addresses...)

	cancellerReport, err := operations.ExecuteSequence(b, mcms_seq.SeqDeployMCMWithConfig, chain, mcms_seq.SeqMCMSDeploymentCfg{
		ChainSelector: chain.Selector,
		ContractType:  common_utils.CancellerManyChainMultisig,
		MCMConfig:     &mcmsCfg,
		Qualifier:     qualifierPtr,
	})
	require.NoError(t, err)
	addresses = append(addresses, cancellerReport.Output.Addresses...)

	timelockRef, err := contract.MaybeDeployContract(b, mcms_ops.OpDeployTimelock, chain, contract.DeployInput[mcms_ops.OpDeployTimelockInput]{
		ChainSelector:  chain.Selector,
		Qualifier:      qualifierPtr,
		TypeAndVersion: deployment.NewTypeAndVersion(common_utils.RBACTimelock, *mcms_ops.MCMSVersion),
		Args: mcms_ops.OpDeployTimelockInput{
			TimelockMinDelay: big.NewInt(1),
			Proposers:        []common.Address{common.HexToAddress(proposerReport.Output.Addresses[0].Address)},
			Bypassers:        []common.Address{common.HexToAddress(bypasserReport.Output.Addresses[0].Address)},
			Cancellers:       []common.Address{common.HexToAddress(cancellerReport.Output.Addresses[0].Address)},
			Admin:            deployer,
			Executors:        []common.Address{},
		},
	}, nil)
	require.NoError(t, err)
	timelockAddr = common.HexToAddress(timelockRef.Address)
	addresses = append(addresses, timelockRef)

	callProxyRef, err := contract.MaybeDeployContract(b, mcms_ops.OpDeployCallProxy, chain, contract.DeployInput[mcms_ops.OpDeployCallProxyInput]{
		ChainSelector:  chain.Selector,
		Qualifier:      qualifierPtr,
		TypeAndVersion: deployment.NewTypeAndVersion(common_utils.CallProxy, *mcms_ops.MCMSVersion),
		Args: mcms_ops.OpDeployCallProxyInput{
			TimelockAddress: timelockAddr,
		},
	}, nil)
	require.NoError(t, err)
	addresses = append(addresses, callProxyRef)
	callProxyAddr := common.HexToAddress(callProxyRef.Address)

	_, err = operations.ExecuteOperation(b, mcms_ops.OpGrantRoleTimelock, chain, contract.FunctionInput[mcms_ops.OpGrantRoleTimelockInput]{
		ChainSelector: chain.Selector,
		Address:       timelockAddr,
		Args: mcms_ops.OpGrantRoleTimelockInput{
			RoleID:  mcms_ops.EXECUTOR_ROLE.ID,
			Account: callProxyAddr,
		},
	})
	require.NoError(t, err)

	return timelockAddr, addresses
}
