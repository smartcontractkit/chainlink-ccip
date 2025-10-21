package deployment_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_0"

	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0"
	routerops1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	rmnops1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	rmnremoteops1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
)

func TestFastCurse(t *testing.T) {
	chain1 := chainsel.TEST_90000001.Selector
	chain2 := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain1, chain2}),
	)
	require.NoError(t, err)
	lggr := logger.Test(t)
	bundle := cldf_ops.NewBundle(
		func() context.Context { return t.Context() },
		lggr,
		cldf_ops.NewMemoryReporter(),
	)
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
	evmDeployer := &v1_0_0.EVMDeployer{}
	dReg := v1_0.NewDeployerRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, v1_0.MCMSVersion, evmDeployer)
	cs := v1_0.DeployMCMS(dReg)
	evmChain1 := env.BlockChains.EVMChains()[chain1]
	evmChain2 := env.BlockChains.EVMChains()[chain2]
	output, err := cs.Apply(*env, v1_0.MCMSDeploymentConfig{
		Chains: map[uint64]v1_0.MCMSDeploymentConfigPerChain{
			chain1: {
				Canceller:        v1_0.SingleGroupMCMSV2(),
				Bypasser:         v1_0.SingleGroupMCMSV2(),
				Proposer:         v1_0.SingleGroupMCMSV2(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test"),
				TimelockAdmin:    evmChain1.DeployerKey.From,
			},
			chain2: {
				Canceller:        v1_0.SingleGroupMCMSV2(),
				Bypasser:         v1_0.SingleGroupMCMSV2(),
				Proposer:         v1_0.SingleGroupMCMSV2(),
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
	for _, addrRef := range allAddrRefs {
		require.NoError(t, ds.Addresses().Add(addrRef))
	}

	env.DataStore = ds.Seal()
}
